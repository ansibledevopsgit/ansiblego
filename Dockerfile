FROM golang:1.19.3 as build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /myapp .
 
FROM alpine:latest as run

# Copy the application executable from the build image
COPY --from=build /myapp /myapp

WORKDIR /app
EXPOSE 9000
CMD ["/myapp"]


# FROM golang:1.19.3
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o main .
# EXPOSE 9000
# CMD ["./main"]

