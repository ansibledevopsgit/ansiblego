FROM golang:1.19.3

WORKDIR /app
 
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out .
 
EXPOSE 9000
CMD ["./out"]