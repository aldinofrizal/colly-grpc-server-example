# Colly gRPC Example

to generate go grpc code, please use
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    news/news.proto
```

for running on local/development, just run
and feel free to change the compose file to suit your current dev env setting
```bash
docker-compose -f docker.yaml up
```

app example that consume using Remote Procedure Call
[gin-app](https://github.com/aldinofrizal/gin-rest-api-example)