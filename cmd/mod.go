package cmd

import (
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
)

func Mod() {
	// TODO
	// 1. read configure file
	// 1. parse it as tag or release
	// 1. clone it into correct path
	fmt.Println("hi world, this is the gh-manage-extension extension!")
	client, err := api.DefaultRESTClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	response := struct{ Login string }{}
	err = client.Get("user", &response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("running as %s\n", response.Login)

	// Clone the given repository to the given directory
	fmt.Println("git clone https://github.com/go-git/go-git")

	clone("https://github.com/go-git/go-git")
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
