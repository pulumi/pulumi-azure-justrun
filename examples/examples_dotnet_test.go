//go:build dotnet || all
// +build dotnet all

package examples

import (
    "testing"

    "github.com/pulumi/pulumi/pkg/v2/testing/integration"
)


func TestWebAppDotNet(t *testing.T) {
	test := integration.ProgramTestOptions{
		Quick:       true,
        SkipRefresh: true,
        Dir:         path.Join(cwd, "cswebapp"),
        Config: map[string]string{
            "azure-native:location": "WestUS",
        },
	}

	integration.ProgramTest(t, &test)
}