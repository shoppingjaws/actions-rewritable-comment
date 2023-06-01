#!/bin/sh
env
cd /app

go run src/main.go -owner ${GITHUB_REPOSITORY_OWNER} -repository ${GITHUB_REPOSITORY#*/} -commentId rewritable-comment -issueNumber ${INPUT_ISSUE_NUMBER} -key ${KEY:-default} -value ${BODY}
