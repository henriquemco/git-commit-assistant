package main

import (
	"fmt"

	git_repository "git_commit_assistant/internal/git"
)

func main() {
	data, err := git_repository.Verify()
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Println(data)
}
