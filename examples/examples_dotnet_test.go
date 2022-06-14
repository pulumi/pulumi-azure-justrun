//go:build dotnet || all
// +build dotnet all

package examples

import (
	"cwd"
	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
	"path"
	"path/filepath"
	"testing"
)

func TestWebAppDotNet(t *testing.T) {
	test := getCsharpBaseOptions(t).With(integration.ProgramTestOptions{
		Dir: path.Join(cwd, "cswebapp"),
	})
	integration.ProgramTest(t, &test)
}

func getCsharpBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseCsharp := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "sdk", "dotnet"),
		},
	})

	return baseCsharp
}
