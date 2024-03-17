FROM golang:1.19.3
  
  
WORKDIR /app

 
COPY go.mod go.sum ./
  
RUN go mod download

 
COPY . .

 
RUN go build -o main .

 
EXPOSE 9000

 
CMD ["./main"]


#************************

# FROM golang:latest as build

# WORKDIR /app

# # Copy the Go module files
# COPY go.mod .
# COPY go.sum .

# # Download the Go module dependencies
# RUN go mod download

# COPY . .

# RUN go build -o /out .
 
# FROM alpine:latest as run

# # Copy the application executable from the build image
# COPY --from=build /out /out

# WORKDIR /app
# EXPOSE 9000
# CMD ["/out"]


#************************

# FROM golang:latest as build

# WORKDIR /app

# # Copy the Go module files
# COPY go.mod .
# COPY go.sum .

# # Download the Go module dependencies
# RUN go mod download

# COPY . .

# RUN go build -o /myapp ./cmd/web
 
# FROM alpine:latest as run

# # Copy the application executable from the build image
# COPY --from=build /myapp /myapp

# WORKDIR /app
# EXPOSE 8080
# CMD ["/myapp"]