# Contributing
When contributing to this package, make sure to bump the version in the schema.json as well as the makefile when preparing a new release, and then regenerate the SDKs before pushing. 
Add tags of the form v0.0.0 AND sdk/v0.0.0 to trigger a build.

## Prerequisites
- Pulumi CLI
- Node.js
- Yarn
- Go 1.17 (to regenerate the SDKs)
- Python 3.6+ (to build the Python SDK)
- .NET Core SDK (to build the .NET SDK)
