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
RUN mkdir -p /dataVolume/avatars
RUN mkdir -p /dataVolume/avatarStore
RUN mkdir -p /dataVolume/configs
RUN mkdir -p /dataVolume/memories
RUN mkdir -p /dataVolume/static/avatarStore
RUN mkdir -p /dataVolume/static/gen/images
RUN mkdir -p /dataVolume/static/gen/tts
RUN mkdir -p /dataVolume/uploads

WORKDIR /build

ENV GO111MODULE=auto

# Build Avatar - The harness of life
# mark an env variable buildEnv as docker, so avatar knows, its running within a container
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -tags docker -ldflags="-w -s -X main.buildEnv=docker" -o avatar cmd/avatar/main.go

##----------------------------------------------------------------------------------------------
## Final Image
FROM scratch

LABEL maintainer="Binu Udayakumar"

# Copy the compiled certificate bundle from the builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

## HTTP Port
EXPOSE 9080

## GRPC Port
EXPOSE 9090 

ARG BUILD_VER
ENV BUILD_VER=$BUILD_VER

## Copy non root user 
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

USER nonroot:nonroot

## create a directory called dataVolume, using busybox
## since scratch images do not allow for directory creation
COPY --from=builder --chown=nonroot:nonroot /dataVolume /dataVolume
# copy compiled app
COPY --from=builder --chown=nonroot:nonroot /build/avatar /avatar


ENTRYPOINT ["./avatar"]
