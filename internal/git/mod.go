package git_repository

import (
	"fmt"
	"os"
	"os/exec"
)

type Data struct {
	Unadded     string
	Uncommitted string
}

func Exist_repository() bool {
	if _, err := os.ReadDir(".git"); err != nil {
		return false
	}
	return true
}

// If the diff command returns anything,
// then there are changes that haven't been added yet.
func Get_diff() (string, error) {
	cmd := exec.Command("git", "diff")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// If the `diff --cached` command returns anything,
// it means that the changes haven't been committed yet.
func Get_Uncommitted_changes() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func Verify() (Data, error) {
	if Exist_repository() {

		diff, _ := Get_diff()

		uncommitted, _ := Get_Uncommitted_changes()

		data := Data{
			Unadded:     diff,
			Uncommitted: uncommitted,
		}

		return data, nil
	}
	return Data{}, fmt.Errorf("There is no git repository.")
}
