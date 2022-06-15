module golangwebapp

replace github.com/pulumi/pulumi-azure-justrun/sdk/go/azure-justrun => ../../sdk/go/azure-justrun

go 1.17

require (
	github.com/pulumi/pulumi-azure-justrun v0.1.14
	github.com/pulumi/pulumi/sdk/v3 v3.34.1
)

require (
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pulumi/pulumi-azure-native/sdk v1.65.0 // indirect
)
