package assistant

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"git_commit_assistant/internal/model"
)

var credentials_file_name = ".git-assistant-credentials"

func Check_credentials_files() (bool, error) {
	home_path, err := os.UserHomeDir()
	if err != nil {
		return false, fmt.Errorf("Error: %s", err.Error())
	}

	file_path := filepath.Join(home_path, credentials_file_name)

	if _, err := os.Stat(file_path); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func Create_credentials_files(credentials model.CredentialsFile) error {
	home_path, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}

	file_path := filepath.Join(home_path, credentials_file_name)

	file, err := os.Create(file_path)
	if err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}

	file.Close()

	data_file := []byte(credentials.Key + credentials.Model)

	err = os.WriteFile(file_path, data_file, 0o644)
	if err != nil {
		return fmt.Errorf("Error: %s", err.Error())
	}

	return nil
}

func Get_credentials() (model.CredentialsFile, error) {
	home_path, err := os.UserHomeDir()
	if err != nil {
		return model.CredentialsFile{}, fmt.Errorf("Error: %s", err.Error())
	}

	file_path := filepath.Join(home_path, credentials_file_name)

	file, err := os.Open(file_path)
	if err != nil {
		return model.CredentialsFile{}, fmt.Errorf("Error: %s", err.Error())
	}

	defer file.Close()

	credentials := model.CredentialsFile{}

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		credentials.Key = scanner.Text()
	}

	if scanner.Scan() {
		credentials.Model = scanner.Text()
	}

	return credentials, nil
}
