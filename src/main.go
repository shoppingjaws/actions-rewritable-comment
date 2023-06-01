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
	value        *string
	commentId   *string
	key         *string
}

func getClient(ctx context.Context) (*github.Client, error) {
	token := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client, nil
}

// Get adc-comment list
func getAdcComment(ctx context.Context, client *github.Client, metadata metadata, page int) (*github.IssueComment, int, error) {
	r, _ := regexp.Compile("<!-- " + *metadata.key + ":" + *metadata.commentId + " -->")
	opt := &github.IssueListCommentsOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
			Page:    page,
		},
	}

	cmts, res, err := client.Issues.ListComments(ctx, *metadata.owner, *metadata.repository, *metadata.issueNumber, opt)
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
		value:        flag.String("value", "", "rewritable comment value"),
		issueNumber: flag.Int("issueNumber", 0, "rewritable comment Issue / PR number"),
		repository:  flag.String("repository", "", "target repository name"),
		key:         flag.String("key", "actions-dashboard-comment", "rewritable comment key"),
	}
	flag.Parse()
	if *metadata.owner == "" || *metadata.repository == "" || *metadata.issueNumber == 0 || *metadata.value == "" || *metadata.commentId == "" {
		fmt.Printf("Error: %v\n", "invalid argument")
		return
	}
	fmt.Printf("owner = %s\ncommentId = %s\nvalue = %s\nissueNumber = %d\nrepository%s\nkey = %s\n", *metadata.owner, *metadata.commentId, *metadata.value, *metadata.issueNumber, *metadata.repository, *metadata.key)
	ctx := context.Background()
	client, err := getClient(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	c, _, _ := getAdcComment(ctx, client, metadata, 0)
	if c != nil {
		// case: Comment Hit
		fmt.Printf("⏪Comment edited from\n" + c.GetBody()+"\n")
		cmt, _, err := client.Issues.EditComment(ctx, *metadata.owner, *metadata.repository, c.GetID(), &github.IssueComment{Body: github.String(*metadata.value + "\n<!-- " + *metadata.key + ":" + *metadata.commentId + " -->")})
		if err != nil {
			panic(err)
		}
		fmt.Printf("⏩Comment edited to\n" + cmt.GetBody()+"\n")
	} else {
		// case: Comment Not Hit
		cmt, _, err := client.Issues.CreateComment(ctx, *metadata.owner, *metadata.repository, *metadata.issueNumber, &github.IssueComment{Body: github.String(*metadata.value + "\n<!-- " + *metadata.key + ":" + *metadata.commentId + " -->")})
		if err != nil {
			panic(err)
		}
		fmt.Printf("🆕New comment posted\n" + cmt.GetBody()+"\n")
	}
}
