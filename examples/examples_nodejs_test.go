//go:build nodejs || all
// +build nodejs all

package examples

import (
    "testing"

    "github.com/pulumi/pulumi/pkg/v2/testing/integration"
)


func TestWebAppNodeJS(t *testing.T) {
	test := integration.ProgramTestOptions{
		Quick:       true,
        SkipRefresh: true,
        Dir:         path.Join(cwd, "nodejswebapp"),
        Config: map[string]string{
            "azure-native:location": "WestUS",
        },
	}

	integration.ProgramTest(t, &test)
}