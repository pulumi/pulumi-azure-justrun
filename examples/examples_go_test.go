//go:build go || all
// +build go all

package examples

import (
	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/executable"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestWebAppGo(t *testing.T) {
	test := getGoBaseOptions(t).With(integration.ProgramTestOptions{
		Dir: filepath.Join(getCwd(t), "golangwebapp"),
	})
	integration.ProgramTest(t, &test)
}

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseGo := base.With(integration.ProgramTestOptions{
		Dependencies: []string{},
		ExtraRuntimeValidation: func(t *testing.T, stack integration.RuntimeValidationStackInfo) {
			cwd := stack.Outputs["cwd"].(string)
			assert.NotEmpty(t, cwd)

			exec, err := executable.FindExecutable("go")
			assert.NoError(t, err)

			err := integration.RunCommand(t, "Go Mod Replace", []string{exec, "mod", "edit", "replace", "github.com/pulumi/pulumi-azure-justrun/sdk/go/azure-justrun=../../sdk/go/azure-justrun"}, cwd, &integration.ProgramTestOptions{})
			assert.NoError(t, err)
		},
	})

	return baseGo
}
