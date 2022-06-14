//go:build go || all
// +build go all

package examples

import (
	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
	"path/filepath"
	"testing"
)

func TestWebAppGo(t *testing.T) {
	test := getGoBaseOptions(t).With(integration.ProgramTestOptions{
		Dir: filepath.Join(getCwd(),, "golangwebapp"),
	})
	integration.ProgramTest(t, &test)
}

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseGo := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "sdk", "go"),
		},
	})

	return baseGo
}
