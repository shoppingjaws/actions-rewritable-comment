FROM golang:1.18.0-alpine3.14
WORKDIR /app
COPY ./ /app/
CMD go mod init \
    go mod tidy \
    go run src/main.go -owner nakamuloud -commentId default -issueNumber 1 -repository actions-dashboard-comment -key abcd -body Goodbye
# RUN ["chmod", "+x", "./entrypoint.sh"]
# ENTRYPOINT ["./entrypoint.sh"]
