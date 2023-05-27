package file_parser

import (
	"encoding/json"
	"io"
	"os"
	"os/user"
)

type Environments struct {
	Environments []Environment `json:"enviroments"`
}

type Environment struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func openJson(fileName string) (*os.File, error) {
	currentUser, err := user.Current()
	path := currentUser.HomeDir + "/.slog/" + fileName
	if err != nil {
		return nil, err
	}

	f, err := os.Open(path)
	return f, err

}

func unmarshallJSON(file *os.File) (*Environments, error) {
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var envs Environments
	json.Unmarshal(byteValue, &envs)
	return &envs, nil
}

func GetSystems() ([]Environment, error) {
	f, err := openJson("environments.json")
	if err != nil {
		return nil, err
	}

	envs, err := unmarshallJSON(f)
	if err != nil {
		return nil, err
	}

	return envs.Environments, err
}
