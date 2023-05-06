
#!/bin/sh -l
env
time=$(date)
go run src/main.go -owner nakamuloud -commentId default -body Hello -issueNumber 1 -repository actions-dashboard-comment -key abc
