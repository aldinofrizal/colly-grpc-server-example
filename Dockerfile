# Build Stage

FROM golang:1.19-alpine AS build

WORKDIR /usr/app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

EXPOSE 9000

RUN CGO_ENABLED=0 G8OOS=linux go build -o /app


# Deploy Stage

FROM alpine:3.16

WORKDIR /

COPY --from=build /app /app

EXPOSE 9000

ENTRYPOINT [ "/app" ]