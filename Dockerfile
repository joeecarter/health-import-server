FROM golang:1.16 AS builder
ARG VERSION=0.0.0
WORKDIR /go/src/github.com/joeecarter/health-import-server/
COPY . .
RUN CGO_ENABLED=0 go build \
	-ldflags="-X 'main.Version=$VERSION'" \
	-o ./server cmd/server/main.go


FROM alpine:latest  
COPY --from=builder /go/src/github.com/joeecarter/health-import-server/server /server
RUN apk --no-cache add ca-certificates && mkdir /config
EXPOSE 8080
ENTRYPOINT [ "/server", "--config=/config/config.json" ]
