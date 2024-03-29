ARG TARGET_ARCH=amd64
ARG TARGET_OS=linux
ARG PLATFORM="${TARGET_OS}/${TARGET_ARCH}"

# Use the official golang alpine image
FROM --platform=${PLATFORM} golang:alpine3.19 as builder

# ARG steps are scoped so we need to repeat them
# to use them within the FROM: https://docs.docker.com/reference/dockerfile/#scope
ARG TARGET_ARCH=amd64
ARG TARGET_OS=linux

RUN apk add --no-cache make

# Create a build environment
RUN mkdir -p /tmp/external-armory
ADD . /tmp/external-armory
WORKDIR /tmp/external-armory
RUN GOOS=linux make release
RUN cp "./release/armory-server_${TARGET_OS}-${TARGET_ARCH}" ./release/external-armory

ARG TARGET_ARCH=amd64
ARG TARGET_OS=linux
ARG PLATFORM="${TARGET_OS}/${TARGET_ARCH}"
# Final layer
FROM --platform=${PLATFORM} alpine:3.19
COPY --from=builder /tmp/external-armory/release/external-armory /opt/external-armory

WORKDIR /data

VOLUME [ "/data/armory-data" ]

ENTRYPOINT [ "/opt/external-armory" ]
