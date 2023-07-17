FROM golang:1.20.5-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o /url-shortener

EXPOSE 52520

CMD ["/url-shortener"]
