//go:build python || all
// +build python all

package examples

import (
    "os"
    "path"
    "testing"

    "github.com/pulumi/pulumi/pkg/v2/testing/integration"
)


func TestWebAppPython(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}
	test := integration.ProgramTestOptions{
		Quick:       true,
        SkipRefresh: true,
        Dir:         path.Join(cwd, "pythonwebapp"),
        Config: map[string]string{
            "azure-native:location": "WestUS",
        },
	}

	integration.ProgramTest(t, &test)
}