# ==== Build stage ====
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

# ==== Publish stage ====
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/aws-ecs-api-go /app

# Define port and labels, these arguments should be
# overridden during build
ARG PORT=8080
ARG NAME=sampleapi
ARG VERSION=0.1
# ENTRYPOINT can't take ARGs so use variable
ENV PORT_VAR=${PORT}
ENV GIN_MODE=release
ENTRYPOINT ./app ${PORT_VAR}

LABEL Name=${NAME} Version=${VERSION}

EXPOSE ${PORT}
