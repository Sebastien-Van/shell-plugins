package docker

import (
	"testing"
	
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)
	
func TestDockerAccessTokenProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, DockerAccessToken().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default": {
			ItemFields: map[sdk.FieldName]string{ // TODO: Check if this is correct
				fieldname.Token: "nYQnrK70BJd2H3NAEXAMPLE",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"DOCKER_TOKEN": "nYQnrK70BJd2H3NAEXAMPLE",
				},
			},
		},
	})
}

func TestDockerAccessTokenImporter(t *testing.T) {
	plugintest.TestImporter(t, DockerAccessToken().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{ // TODO: Check if this is correct
				"DOCKER_TOKEN": "nYQnrK70BJd2H3NAEXAMPLE",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Token: "nYQnrK70BJd2H3NAEXAMPLE",
					},
				},
			},
		},
		// TODO: If you implemented a config file importer, add a test file example in docker/test-fixtures
		// and fill the necessary details in the test template below.
		"config file": {
			Files: map[string]string{
				// "~/path/to/config.yml": plugintest.LoadFixture(t, "config.yml"),
			},
			ExpectedCandidates: []sdk.ImportCandidate{
			// 	{
			// 		Fields: map[sdk.FieldName]string{
			// 			fieldname.Token: "nYQnrK70BJd2H3NAEXAMPLE",
			// 		},
			// 	},
			},
		},
	})
}
