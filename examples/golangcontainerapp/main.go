package main

import (
	justrun "github.com/pulumi/pulumi-azure-justrun/sdk/go/azure-justrun"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		containerapp, err := justrun.NewContainerapp(ctx, "containerapp", &justrun.containerappArgs{
			ImageDirectory: pulumi.String("node-app"),
		})
		if err != nil {
			return err
		}

		ctx.Export("url", containerapp.Url)

		return nil
	})
}
