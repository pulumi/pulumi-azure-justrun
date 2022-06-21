// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func getEnvRegion() string {
	/*
		envRegion := os.Getenv("AWS_REGION")
		if envRegion == "" {
			t.Skipf("Skipping test due to missing AWS_REGION environment variable")
		}*/

	return "WestUS"
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		ExpectRefreshChanges: true,
		SkipRefresh:          true,
		Quick:                true,
		Config: map[string]string{
			"azure-native:location": getEnvRegion(),
		},
	}
}
