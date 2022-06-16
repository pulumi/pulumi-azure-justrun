//go:build python || all
// +build python all

package examples

import (
	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
	"path/filepath"
	"testing"
)

func TestWebAppPython(t *testing.T) {
	test := getPythonBaseOptions(t).With(integration.ProgramTestOptions{
		Dir: filepath.Join(getCwd(t), "pythonwebapp"),
	})
	integration.ProgramTest(t, &test)
}

func TestContainerAppPython(t *testing.T) {
	test := getPythonBaseOptions(t).With(integration.ProgramTestOptions{
		Dir: filepath.Join(getCwd(t), "pythoncontainerapp"),
	})
	integration.ProgramTest(t, &test)
}

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	pythonBase := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "..", "sdk", "python"),
		},
	})

	return pythonBase
}
