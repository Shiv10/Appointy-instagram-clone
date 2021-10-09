FROM golang:latest

LABEL maintainer="Shivansh Sharma"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download 

COPY . .

RUN go build

CMD ["./Appointy-instagram-clone"]