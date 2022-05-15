## Build
FROM arm64v8/golang:1.18.2-alpine as build
WORKDIR /build
COPY go.mod ./
RUN go mod download
RUN go mod tidy
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /app cmd/myapp/main.go
RUN chmod +x /app

## Deploy
FROM alpine:3
COPY tls /etc/tls
COPY --from=build /app /bin/app
EXPOSE 443
ENTRYPOINT ["/bin/app"]
