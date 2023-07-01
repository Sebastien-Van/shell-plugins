package docker

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func DockerAccessToken() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.DockerAccessToken,
		DocsURL:       sdk.URL("https://docs.docker.com/docker-hub/access-tokens/"),
		ManagementURL: sdk.URL("https://hub.docker.com/settings/security"),
		Fields: []schema.CredentialField{
			{
				Name:                fieldname.Token,
				MarkdownDescription: "Token used to authenticate to Docker.",
				Secret:              true,
				Composition: &schema.ValueComposition{
					Length: 23,
					Charset: schema.Charset{
						Uppercase: true,
						Lowercase: true,
						Digits:    true,
					},
				},
			},
		},
		DefaultProvisioner: provision.EnvVars(defaultEnvVarMapping),
	}
}

var defaultEnvVarMapping = map[string]sdk.FieldName{
	"DOCKER_TOKEN": fieldname.Token,
}



// The platform does not store the Docker Access Token in a local config file