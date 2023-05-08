FROM golang:1.18.0-alpine3.14
WORKDIR /app
COPY ./ /app/
CMD go mod init \
    go mod tidy
RUN ["chmod", "+x", "./entrypoint.sh"]
RUN ["chmod", "+x", "./src/main.go"]
ENTRYPOINT ["./entrypoint.sh"]
