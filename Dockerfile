FROM golang:1.18.0-alpine3.14
WORKDIR /github/workspace
COPY ./ /github/workspace/
CMD go mod init \
    go mod tidy
RUN ["chmod", "+x", "./entrypoint.sh"]
ENTRYPOINT ["./entrypoint.sh"]
