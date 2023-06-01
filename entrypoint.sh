#!/bin/sh
cd /app

if [ -z "$OWNER" ]; then
  OWNER=${GITHUB_REPOSITORY_OWNER}
fi

if [ -z "$REPOSITORY" ]; then
  REPOSITORY=${GITHUB_REPOSITORY#*/}
fi

if [ -z "$ISSUE_NUMBER" ]; then
  ISSUE_NUMBER=${GITHUB_REF%/*}
fi

echo OWNER: "${OWNER}"
echo REPOSITORY: "${REPOSITORY}"
echo ISSUE_NUMBER: "${ISSUE_NUMBER}"

go run src/main.go -owner ${OWNER} -repository ${REPOSITORY} -commentId rewritable-comment -issueNumber ${ISSUE_NUMBER} -key ${KEY:-default} -value ${VALUE}
