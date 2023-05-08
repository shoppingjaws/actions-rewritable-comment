#!/bin/sh
env
ls /app
time=$(date)
# go run src/main.go -owner nakamuloud -commentId default -body Hello -issueNumber 1 -repository actions-dashboard-comment -key abc
go run src/main.go -owner nakamuloud -commentId default -issueNumber 1 -repository actions-dashboard-comment -key abcd -body Goodbye
