FROM golang:1.20.5-alpine

WORKDIR /app

COPY * ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortener

EXPOSE 52520

CMD ["/url-shortener"]
