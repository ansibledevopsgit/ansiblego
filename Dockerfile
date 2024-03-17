ARG BUILDER_IMAGE=golang:1.19.3
ARG DISTROLESS_IMAGE=gcr.io/distroless/static

FROM ${BUILDER_IMAGE} as builder
RUN update-ca-certificates
WORKDIR /myapp

EXPOSE 8080

ENV GOCACHE=$HOME/.cache/go-build
RUN --mount=type=cache,target=$GOCACHE

COPY go.mod .
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o /myapp/hello .

FROM ${DISTROLESS_IMAGE}

COPY --from=builder /myapp/hello /myapp/hello

ENTRYPOINT ["/myapp/hello"]

# FROM golang:1.19.3 as builder
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# FROM alpine:latest  

# RUN apk --no-cache add ca-certificates
# WORKDIR /root/
# COPY --from=builder /app/main .
# EXPOSE 8080

# CMD ["./main"] 

# FROM golang:1.19.3
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o main .
# EXPOSE 9000
# CMD ["./main"]

