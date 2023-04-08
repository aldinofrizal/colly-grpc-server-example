FROM golang:alpine

WORKDIR /usr/app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 G8OOS=linux go build -o /app

EXPOSE 9000

CMD ["/app"]