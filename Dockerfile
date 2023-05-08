FROM golang:1.18.0-alpine3.14

COPY ./ /app/
WORKDIR /app
CMD go mod tidy
ENTRYPOINT ["/app/entrypoint.sh"]
