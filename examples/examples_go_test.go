//go:build go || all
// +build go all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)


func TestWebAppGo(t *testing.T) {
	test := getGoBaseOptions(t)
	integration.ProgramTest(t, &test)
}

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseGo := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"github.com/pulumi/pulumi-azure-native/sdk/go",
		},
	})

	return baseGo
}