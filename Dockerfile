FROM golang:latest
LABEL org.opencontainers.image.authors="shetu2153@gmail.com"
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o main .

CMD ["/app/main"]
