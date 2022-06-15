//go:build nodejs || all
// +build nodejs all

package examples

import (
	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestWebAppNodeJS(t *testing.T) {
	test := getJSBaseOptions(t).With(integration.ProgramTestOptions{
		Dir: filepath.Join(getCwd(t), "nodejswebapp"),
	})
	integration.ProgramTest(t, &test)
}

func TestContainerAppNodeJS(t *testing.T) {
	test := getJSBaseOptions(t).With(integration.ProgramTestOptions{
		Dir: filepath.Join(getCwd(t), "nodejscontainerapp"),
	})
	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{},
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			cwd := stack.Outputs["cwd"].(string)
			assert.NotEmpty(t, cwd)

			exec, err := executable.FindExecutable("yarn")
			assert.NoError(t, err)

			err = integration.RunCommand(t, "Yarn Link", []string{exec, "link", "@pulumi/azure-justrun"}, cwd, &integration.ProgramTestOptions{})
			assert.NoError(t, err)
		},
	})

	return baseJS
}
