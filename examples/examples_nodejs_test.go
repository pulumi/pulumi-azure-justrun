//go:build nodejs || all
// +build nodejs all

package examples

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestWebAppNodeJS(t *testing.T) {
	test := getJSBaseOptions(t)

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join("..", "sdk", "nodejs"),
		},
	})

	return baseJS
}
