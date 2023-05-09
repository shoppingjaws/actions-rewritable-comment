#!/bin/sh
env
cd /app
if ${GITHUB_REF_TYPE}
go run src/main.go -owner ${GITHUB_REPOSITORY_OWNER} -commentId rewritable-comment -issueNumber 1 -repository ${GITHUB_REPOSITORY#*/} -key abcd -body Goodbye
