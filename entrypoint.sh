#!/bin/sh
env
cd /app

if [ "${$GITHUB_EVENT_NAME}" == "pull_request" ];
then
  echo "Pull Request"
  NUMBER=$(echo $GITHUB_REF | sed -e 's/[^0-9]//g')
else
  NUMBER=${INPUT_ISSUE_NUMBER}
fi

go run src/main.go -owner ${GITHUB_REPOSITORY_OWNER} -repository ${GITHUB_REPOSITORY#*/} -commentId rewritable-comment -issueNumber ${NUMBER} -key ${KEY:-default} -value ${BODY}
