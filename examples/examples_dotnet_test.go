//go:build dotnet || all
// +build dotnet all

package examples

import (
	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestWebAppDotNet(t *testing.T) {
	test := getCsharpBaseOptions(t).With(integration.ProgramTestOptions{
		Dir: filepath.Join(getCwd(t), "cswebapp"),
	})
	integration.ProgramTest(t, &test)
}

func getCsharpBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseCsharp := base.With(integration.ProgramTestOptions{
		Dependencies: []string{},
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {

			cwd := stack.Outputs["cwd"].(string)
			assert.NotEmpty(t, cwd)

			err := integration.RunCommand(t, "yarn", []string{"link"}, cwd, &integration.ProgramTestOptions{})
			assert.NoError(t, err)
		},
	})

	return baseCsharp
}
