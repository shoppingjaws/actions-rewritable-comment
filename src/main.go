package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type metadata struct {
	owner       *string
	repository  *string
	issueNumber *int
	body        *string
	commentId   *string
	key         *string
}

func getClient(ctx context.Context) (*github.Client, error) {
	token := os.Getenv("GH_TOKEN")
	print(token)
	// user := os.Getenv("GH_REPO")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client, nil
}

// Get adc-comment list
func getAdcComment(ctx context.Context, client *github.Client, metadata metadata, page int) (*github.IssueComment, int, error) {
	print("getAdcComment", page, metadata.commentId, "\n")
	r, _ := regexp.Compile("<!-- " + *metadata.key + ":" + *metadata.commentId + " -->")
	opt := &github.IssueListCommentsOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
			Page:    page,
		},
	}

	cmts, res, err := client.Issues.ListComments(ctx, *metadata.owner, *metadata.repository, 1, opt)
	if err != nil {
		panic(err)
	}
	for _, c := range cmts {
		if r.Find([]byte(c.GetBody())) != nil {
			return c, res.NextPage, nil
		}
	}
	if res.NextPage == 0 {
		return nil, 0, nil
	}
	return getAdcComment(ctx, client, metadata, res.NextPage)
}

func main() {
	var metadata metadata = metadata{
		owner:       flag.String("owner", "", "target repository owner"),
		commentId:   flag.String("commentId", "", "rewritable comment id"),
		body:        flag.String("body", "", "rewritable comment body"),
		issueNumber: flag.Int("issueNumber", 0, "rewritable comment Issue / PR number"),
		repository:  flag.String("repository", "", "target repository name"),
		key:         flag.String("key", "actions-dashboard-comment", "rewritable comment key"),
	}
	flag.Parse()
	if *metadata.owner == "" || *metadata.repository == "" || *metadata.issueNumber == 0 || *metadata.body == "" || *metadata.commentId == "" {
		fmt.Printf("Error: %v\n", "invalid argument")
		return
	}
	ctx := context.Background()
	client, err := getClient(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	c, _, _ := getAdcComment(ctx, client, metadata, 0)
	if c != nil {
		// case: Comment Hit
		client.Issues.EditComment(ctx, *metadata.owner, *metadata.repository, c.GetID(), &github.IssueComment{Body: github.String(*metadata.body + "\n<!-- " + *metadata.key + ":" + *metadata.commentId + " -->")})
	} else {
		// case: Comment Not Hit
		client.Issues.CreateComment(ctx, *metadata.owner, *metadata.repository, *metadata.issueNumber, &github.IssueComment{Body: github.String(*metadata.body + "\n<!-- " + *metadata.key + ":" + *metadata.commentId + " -->")})
	}
}
