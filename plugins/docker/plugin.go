package docker

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "docker",
		Platform: schema.PlatformInfo{
			Name:     "Docker",
			Homepage: sdk.URL("https://hub.docker.com/"), 
		},
		Credentials: []schema.CredentialType{
			DockerAccessToken(),
		},
		Executables: []schema.Executable{
			DockerCLI(),
		},
	}
}
