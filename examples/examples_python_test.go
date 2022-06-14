//go:build python || all
// +build python all

package examples

import (
	"cwd"
	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
	"path"
	"path/filepath"
	"testing"
)

func TestWebAppPython(t *testing.T) {
	test := getPythonBaseOptions(t).With(integration.ProgramTestOptions{
		Dir: path.Join(cwd, "pythonwebapp"),
	})
	integration.ProgramTest(t, &test)
}

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	envRegion := getEnvRegion(t)
	base := getBaseOptions()
	pythonBase := base.With(integration.ProgramTestOptions{
		Config: map[string]string{
			"azure-native:location": envRegion,
		},
		Dependencies: []string{
			filepath.Join("..", "sdk", "python"),
		},
	})

	return pythonBase
}
