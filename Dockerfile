ARG BUILDER_IMAGE=golang:1.19.3
ARG DISTROLESS_IMAGE=gcr.io/distroless/static

FROM ${BUILDER_IMAGE} as builder
RUN update-ca-certificates
WORKDIR /myapp



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

EXPOSE 9000

ENTRYPOINT ["/myapp/hello"]
 