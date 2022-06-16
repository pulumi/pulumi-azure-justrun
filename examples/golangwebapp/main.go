package main

import (
	justrun "github.com/pulumi/pulumi-azure-justrun/sdk"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		webapp, err := justrun.NewWebapp(ctx, "webapp", &justrun.WebappArgs{
			FilePath: pulumi.String("./www"),
		})
		if err != nil {
			return err
		}

		ctx.Export("url", webapp.Url)

		return nil
	})
}
