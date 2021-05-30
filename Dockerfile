FROM golang:1.16 AS builder
WORKDIR /go/src/github.com/joeecarter/health-import-server/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

FROM alpine:latest  
WORKDIR /health-import-server/
COPY --from=builder /go/src/github.com/joeecarter/health-import-server/server .
RUN apk --no-cache add ca-certificates && mkdir /config
ENV CONFIG_FILE_PATH=/config/config.json
EXPOSE 8080
ENTRYPOINT ["/health-import-server/server"]
