//go:build go || all
// +build go all

package examples

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestWebAppGo(t *testing.T) {
	test := getGoBaseOptions(t)
	integration.ProgramTest(t, &test)
}

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions()
	baseGo := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			filepath.Join(".", "sdk", "go"),
		},
	})

	return baseGo
}
