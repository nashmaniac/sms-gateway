FROM golang:latest
MAINTAINER Raju Ahmed Shetu
RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o main .

CMD ["/app/main"]
