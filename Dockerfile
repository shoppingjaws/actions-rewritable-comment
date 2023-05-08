FROM golang:1.18.0-alpine3.14

COPY ./ ./
CMD go mod tidy
RUN ls
ENTRYPOINT ["entrypoint.sh"]
