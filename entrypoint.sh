#!/bin/sh
env
ls /app/src
pwd
time=$(date)
go run src/main.go -owner nakamuloud -commentId default -issueNumber 1 -repository actions-dashboard-comment -key abcd -body Goodbye
