package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	pbAvatar "github.com/binuud/avatar-conscious/gen/go/v1/avatar"
	pbConscious "github.com/binuud/avatar-conscious/gen/go/v1/conscious"
	pbSettings "github.com/binuud/avatar-conscious/gen/go/v1/settings"
	"github.com/binuud/avatar-conscious/pkg/bootstrap"
	"github.com/binuud/avatar-conscious/pkg/servers/avatarServer"
	llmProxyServer "github.com/binuud/avatar-conscious/pkg/servers/llmProxyServer"
	"github.com/binuud/avatar-conscious/pkg/servers/staticFileServer"
	"github.com/binuud/avatar-conscious/pkg/utils"
	spahttpserver "github.com/binuud/watchdog/pkg/spaHttpServer"

	// mqttServer "github.com/dronasys-com/brain-nodes/pkg/servers/mqttBridge"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//go:embed ui/*
var embeddedFiles embed.FS

var (
	server_grpc_port = flag.Int("grpc_port", 9090, "The Brain Backend Server GRPC port, no token required unsecured")
	server_http_port = flag.Int("http_port", 9080, "The Brain Backend Server HTTP port, no token required unsecured")
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// type spaFileSystem struct {
// 	baseFS    fs.FS
// 	indexPath string
// }

// to know that the binary is running within a docker container
var buildEnv = "local" // "local" - for normal builds || "docker" - for docker builds

func serveUI() *http.ServeMux {

	// Strip the "ui" prefix from the embedded paths
	// so we don't have to include "/ui/" in your URLs.
	uiFS, err := fs.Sub(embeddedFiles, "ui")
	if err != nil {
		panic(err)
	}

	// Initialize our SPA handler wrapper
	handler := spahttpserver.SpaFileSystem{
		BaseFS:    uiFS,
		IndexPath: "index.html",
	}

	httpMux := http.NewServeMux()

	httpMux.Handle("/", handler)

	logrus.Infof("Server directory %s -> %s", "/", ".")

	return httpMux
}

func serveStaticPaths(httpMux *http.ServeMux) *http.ServeMux {

	// OPEN_API_FILES_PATH is the folder where the swagger files for apps are stored
	SERVABLE_DIRECTORY := utils.GetEnv("SERVABLE_DIRECTORY", "/dataVolume/static")

	if !strings.Contains(SERVABLE_DIRECTORY, "/dataVolume/static") {
		logrus.Panic("Protection fail, only directory having /dataVolume/static can be served")
	}

	SERVABLE_UI_DIRECTORY := os.Getenv("SERVABLE_UI_DIRECTORY")
	if SERVABLE_UI_DIRECTORY == "" {
		SERVABLE_UI_DIRECTORY = "/ui"
	}

	if !strings.Contains(SERVABLE_UI_DIRECTORY, "/ui") {
		logrus.Panic("Protection fail, only directory having /ui can be served")
	}

	staticFileServer.StaticServeAiPath(httpMux, "/static/", SERVABLE_DIRECTORY)
	// staticFileServer.SPAStaticServePath(httpMux, "/", SERVABLE_UI_DIRECTORY)

	return httpMux
}

func runHtppServer() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.NewClient(
		fmt.Sprintf("0.0.0.0:%d", *server_grpc_port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwMux := runtime.NewServeMux()

	// consciousServer
	err = pbConscious.RegisterConsciousHandler(ctx, gwMux, conn)
	if err != nil {
		log.Fatalln(err)
	}

	// avatar store server
	err = pbAvatar.RegisterAvatarStoreServiceHandler(ctx, gwMux, conn)
	if err != nil {
		log.Fatalln(err)
	}

	// avatar server
	err = pbAvatar.RegisterAvatarServiceHandler(ctx, gwMux, conn)
	if err != nil {
		log.Fatalln(err)
	}

	// systemMetricsServer
	err = pbConscious.RegisterMetricsHandler(ctx, gwMux, conn)
	if err != nil {
		log.Fatalln(err)
	}

	// settings server
	err = pbSettings.RegisterSettingsHandler(ctx, gwMux, conn)
	if err != nil {
		log.Fatalln(err)
	}

	httpMux := serveUI()
	serveStaticPaths(httpMux)

	// for bidirectional stream on avatar live server
	httpMux.HandleFunc("/api/avatar/ws", avatarServer.WebsocketHandler)
	// mqttServer.InitMQTT()

	// handle llm proxt on path /llmHandler
	httpMux.Handle("/llmHandler/", http.StripPrefix("/llmHandler", llmProxyServer.GetLlmProxyHandler()))
	logrus.Infof("Server directory %s -> %s", "/llmHandler/", "llm proxy")

	// handle websocket
	// httpMux.HandleFunc("/ws", mqttServer.WebsocketHandler)

	// link the concious gateway server to httpmux
	httpMux.Handle("/api/", http.StripPrefix("/api", gwMux))
	logrus.Infof("Server directory %s -> %s", "/api/avatar/ws", "avatar control websocket")
	logrus.Infof("Server directory %s -> %s", "/api/", "conscious server")
	logrus.Infof("Server directory %s -> %s", "/api/", "metrics server")
	logrus.Infof("Server directory %s -> %s", "/api/", "avatar store server")
	logrus.Infof("Server directory %s -> %s", "/api/", "settings server")

	// Configure CORS options as needed
	corsEnabledMux := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Replace "*" with specific origin(s) for production
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"}, // []string{"Authorization", "Content-Type", "Accept"},
		AllowCredentials: true,
		Debug:            false, // Set to true to trace CORS errors in console
	})

	corsEnabledHandler := corsEnabledMux.Handler(httpMux)

	gwServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", *server_http_port),
		Handler: corsEnabledHandler,
	}

	logrus.Infof("HTTP and websocket server starting  :(%d)", *server_http_port)
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Fatalln(gwServer.ListenAndServe())
}

func main() {
	logrus.Println("Welcome to life")
	logrus.Printf("My App %s, commit %s, built at %s", version, commit, date)
	bootstrap.Bootstrap(buildEnv)

	go func() {
		bootstrap.Bootstrap_GRPC_Services(*server_grpc_port)
	}()

	runHtppServer()

}
