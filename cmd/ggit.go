package main

import (
	"fmt"
	"log"

	git_repository "git_commit_assistant/internal/git"
	"git_commit_assistant/internal/handler"
	"git_commit_assistant/internal/model"
	"git_commit_assistant/internal/parser"
	"git_commit_assistant/internal/ui"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	newInstace := model.Application{}

	fmt.Println("\nHi! I'm a git commit assistant\n")
	if git_repository.Exist_repository() {

		//------
		unadded, err := git_repository.Get_unadded_changes()
		if err != nil {
			fmt.Printf("\nERROR : %s", err.Error())
		}
		newInstace.UnAdded = unadded

		//------
		if newInstace.UnAdded != "" {

			fmt.Println("I noticed there are changes outside the stage;")
			prompt := ui.Select("Would you like me to add them ? [Y/N]")
			option := "y"
			c := prompt[3]
			for _, char := range c {
				if char == 'x' {
					option = "n"
				}
			}

			if option == "y" {
				if err := git_repository.Add_changes(); err != nil {
					fmt.Printf("\nERROR : %s", err.Error())
				}
			}
		}
		//------
		uncommitted, err := git_repository.Get_uncommitted_changes()
		if err != nil {
			fmt.Printf("\nERROR : %s", err.Error())
		}

		if uncommitted == "" {
			fmt.Println("There are no changes to be committed.")
			return
		}

		newInstace.UnCommitted = uncommitted

		//------

		fmt.Println("What did you do?")
		fmt.Print(" > ")
		fmt.Scan(&newInstace.Description)

		// --------
		data := parser.Message(newInstace)

		stop := make(chan bool)

		go ui.Loading(stop)

		resp, _ := handler.Get_commit_message(data)

		stop <- true

		//---------

		fmt.Println(parser.Get_commit_message(resp.Text))
	}
}
