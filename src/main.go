package main

import (
	"fmt"
	"log"

	"github.com/cli/go-gh/v2"
)

func main() {
	// These examples assume `gh` is installed and has been authenticated.

	// Shell out to a gh command and read its output.	
	issueList, _, err := gh.Exec("issue", "list", "--repo", "cli/cli", "--limit", "5")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(issueList.String())

	// // Use an API client to retrieve repository tags.
	// client, err := api.DefaultRESTClient()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// response := []struct{
	// 	Name string
	// }{}
	// err = client.Get("repos/cli/cli/tags", &response)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(response)
}
