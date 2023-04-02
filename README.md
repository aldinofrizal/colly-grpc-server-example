# Colly gRPC Example

to generate go grpc code, please use
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    news/news.proto
```