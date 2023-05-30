package environment

import (
	"encoding/json"
	"io"
	"os"
	"os/user"
)

type Environments struct {
	Kubernetes map[string]KubernetesOptions `json:"kubernetes"`
	// Docker     []EnvDocker                  `json:"docker"`
}

type KubernetesOptions struct {
	Cmd string `json:"cmd"`
}

// type EnvDocker struct {
// 	Name string `json:"name"`
// }

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

// GetEnvironments will return the saved environments that we
// have saved
func GetEnvironments() (*Environments, error) {
	f, err := openJson("environments.json")
	if err != nil {
		return nil, err
	}

	envs, err := unmarshallJSON(f)
	if err != nil {
		return nil, err
	}

	return envs, err
}
