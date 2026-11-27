#
# This is a multi target docker file
# The first state - builder - builds all the binaries
# Subsequent targets are used to build the final docker image, with the prebuild binaries
# This docker file is used to create multiple docker images
#

FROM --platform=$BUILDPLATFORM golang:alpine AS builder
ARG TARGETOS
ARG TARGETARCH

LABEL maintainer="Binu Udayakumar"

RUN apk add --no-cache ca-certificates

RUN addgroup -S nonroot -g 10001 && adduser -S nonroot -G nonroot -u 10001
RUN mkdir /build
COPY ./ /build/

## For mounting dataVolume files on distroless containers
RUN mkdir -p /dataVolume/configs
RUN mkdir -p /dataVolume/memories
RUN mkdir -p /dataVolume/static/avatarStore
RUN mkdir -p /dataVolume/static/gen/images
RUN mkdir -p /dataVolume/static/gen/tts
RUN mkdir -p /dataVolume/uploads
RUN mkdir -p /avatar-store

WORKDIR /build

ENV GO111MODULE=auto

# Build Avatar - The harness of life
# mark an env variable buildEnv as docker, so avatar knows, its running within a container
# RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -tags docker -ldflags="-w -s -X main.buildEnv=docker" -o avatar cmd/avatar/main.go

##----------------------------------------------------------------------------------------------
## Final Image
## FROM scratch
FROM binuud/avatar-mcp-servers:latest 

LABEL maintainer="Binu Udayakumar"

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT

# Copy the compiled certificate bundle from the builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

## HTTP Port
EXPOSE 9080

## GRPC Port
EXPOSE 9090 

ARG BUILD_VER
ENV BUILD_VER=$BUILD_VER
ENV BUILD_ENV="docker"
ENV PATH="$PATH:/mcpServers/bin"
ENV FILE_STORE_PATH="/dataVolume"
ENV AVATAR_STORE_DIRECTORY="/avatar-store"


## Copy non root user 
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

COPY avatar-store /avatar-store

RUN mkdir -p /tmp/dist

# USER nonroot:nonroot

## create a directory called dataVolume, using busybox
## since scratch images do not allow for directory creation
COPY --from=builder --chown=nonroot:nonroot /dataVolume /dataVolume

## COPY PREBUILT BINARIES
ARG APPS="avatar"

# copy compiled app
# Mount the local 'dist' folder temporarily and loop through the apps
RUN --mount=type=bind,source=dist,target=/tmp/dist \
    VARIANT="$TARGETVARIANT"; \
    if [ "$TARGETARCH" = "amd64" ] && [ -z "$VARIANT" ]; then \
        VARIANT="v1"; \
    fi; \
    if [ "$TARGETARCH" = "arm64" ] && [ -z "$VARIANT" ]; then \
        VARIANT="v8.0"; \
    fi; \
    for app in $APPS; do \
        cp "/tmp/dist/${app}_${TARGETOS}_${TARGETARCH}_${VARIANT:+$VARIANT}/${app}" "/usr/local/bin/${app}" && \
        chmod +x "/usr/local/bin/${app}"; \
    done

# CMD ["tail", "-f", "/dev/null"]

ENTRYPOINT ["/usr/local/bin/avatar"]