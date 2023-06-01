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
  ISSUE_NUMBER=${ISSUE_NUMBER##*/}
fi

echo OWNER: "${OWNER}"
echo REPOSITORY: "${REPOSITORY}"
echo ISSUE_NUMBER: "${ISSUE_NUMBER}"
echo "<!-- KEY : COMMENT_ID --> : " "<!-- ${KEY} : ${COMMENT_ID} -->"
echo VALUE: "${VALUE}"

go run src/main.go -owner ${OWNER} -repository ${REPOSITORY} -commentId ${COMMENT_ID} -issueNumber ${ISSUE_NUMBER} -key ${KEY:-default} -value ${VALUE}
