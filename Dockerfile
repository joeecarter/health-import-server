FROM alpine:latest  
COPY ./server /server
RUN apk --no-cache add ca-certificates && mkdir /config
EXPOSE 8080
ENTRYPOINT [ "/server", "--config=/config/config.json" ]
