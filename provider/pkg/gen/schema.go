package gen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

const (
	azureNativeTypesVersion = "1.65.0"
)

// nolint: lll
func GenerateSchema(packageDir string) schema.PackageSpec {
	//dependencies := readPackageDependencies(packageDir)
	azureNativeSpec := getPackageSpec("azure-native", "1.65.0") //getPackageSpec("azure-native", dependencies.AzureNative)
	dockerSpec := getPackageSpec("docker", "3.2.0")             //getPackageSpec("docker", dependencies.Docker)

	packageSpec := schema.PackageSpec{
		Name:        "azure-justrun",
		Version:     "v0.2.0",
		Description: "Provides some simple components to run azure apps",
		License:     "Apache-2.0",
		Publisher:   "Pulumi",
		Keywords:    []string{""},
		Homepage:    "https://pulumi.com",
		Repository:  "https://github.com/pulumi/pulumi-azure-justrun",

		Functions: map[string]schema.FunctionSpec{},
		Resources: map[string]schema.ResourceSpec{},
		Types:     map[string]schema.ComplexTypeSpec{},
		Language: map[string]schema.RawMessage{
			"csharp": rawMessage(map[string]interface{}{
				"packageReferences": map[string]string{
					"Pulumi":             "3.*",
					"Pulumi.AzureNative": "1.*",
				},
				"liftSingleValueMethodReturns": true,
				"respectSchemaVersion":         true,
			}),
			"go": rawMessage(map[string]interface{}{
				"generateResourceContainerTypes": true,
				"importBasePath":                 "github.com/pulumi/pulumi-azure-justrun/sdk/go/azure-justrun",
				"liftSingleValueMethodReturns":   true,
				"internalDependencies":           []string{"github.com/pulumi/pulumi-docker/sdk/v3/go/docker"},
				"respectSchemaVersion":           true,
			}),
			"nodejs": rawMessage(map[string]interface{}{
				"dependencies": map[string]string{
					"@pulumi/pulumi":       "^3.0.0",
					"@pulumi/azure-native": "^1.0.0",
				},
				"respectSchemaVersion": true,
				"containsEnums":        true,
			}),
			"python": rawMessage(map[string]interface{}{
				"requires": map[string]string{
					"pulumi":              ">=3.0.0,<4.0.0",
					"pulumi-azure-native": ">=1.0.0,<2.0.0",
				},
				"usesIOClasses":                true,
				"readme":                       "Pulumi Azure Justrun Components",
				"liftSingleValueMethodReturns": true,
				"respectSchemaVersion":         true,
			}),
		},
	}

	return extendSchemas(packageSpec,
		generateSchema(azureNativeSpec, dockerSpec),
	)
}

func packageRef(spec schema.PackageSpec, ref string) string {
	version := spec.Version
	refWithoutHash := strings.TrimLeft(ref, "#")
	return fmt.Sprintf("/%s/%s/schema.json#%s", spec.Name, version, refWithoutHash)
}

func getPackageSpec(name, version string) schema.PackageSpec {
	url := fmt.Sprintf("https://raw.githubusercontent.com/pulumi/pulumi-%s/v%s/provider/cmd/pulumi-resource-%s/schema.json", name, version, name)
	spec := getSpecFromUrl(url)
	if spec.Version == "" {
		// Version is rarely included, so we'll just add it.
		spec.Version = "v" + version
	}
	return spec
}

func getSpecFromUrl(url string) schema.PackageSpec {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var spec schema.PackageSpec
	err = json.Unmarshal(body, &spec)
	if err != nil {
		log.Fatal(err)
	}
	return spec
}

// Perform a simple string replacement on Refs in all sub-specs
func renameComplexRefs(spec schema.ComplexTypeSpec, old, new string) schema.ComplexTypeSpec {
	spec.Properties = renamePropertiesRefs(spec.Properties, old, new)
	return spec
}

func renameAwsPropertiesRefs(spec schema.PackageSpec, propertySpec map[string]schema.PropertySpec) map[string]schema.PropertySpec {
	return renamePropertiesRefs(propertySpec, "#/types/aws:", packageRef(spec, "/types/aws:"))
}

func renamePropertiesRefs(propertySpec map[string]schema.PropertySpec, old, new string) map[string]schema.PropertySpec {
	properties := map[string]schema.PropertySpec{}
	for k, v := range propertySpec {
		properties[k] = renamePropertyRefs(v, old, new)
	}
	return properties
}

func renamePropertyRefs(propSpec schema.PropertySpec, old, new string) schema.PropertySpec {
	if propSpec.Ref != "" {
		propSpec.Ref = strings.Replace(propSpec.Ref, old, new, 1)
	}
	if propSpec.AdditionalProperties != nil {
		additionalProperties := renameTypeSpecRefs(*propSpec.AdditionalProperties, old, new)
		propSpec.AdditionalProperties = &additionalProperties
	}
	if propSpec.Items != nil {
		items := renameTypeSpecRefs(*propSpec.Items, old, new)
		propSpec.Items = &items
	}
	if propSpec.OneOf != nil {
		propSpec.OneOf = renameTypeSpecsRefs(propSpec.OneOf, old, new)
	}
	return propSpec
}

func renameTypeSpecsRefs(typeSpec []schema.TypeSpec, old, new string) []schema.TypeSpec {
	newSpecs := make([]schema.TypeSpec, len(typeSpec))
	for i, spec := range typeSpec {
		newSpecs[i] = renameTypeSpecRefs(spec, old, new)
	}
	return newSpecs
}
func renameTypeSpecRefs(typeSpec schema.TypeSpec, old, new string) schema.TypeSpec {
	if typeSpec.Ref != "" {
		typeSpec.Ref = strings.Replace(typeSpec.Ref, old, new, 1)
	}
	if typeSpec.AdditionalProperties != nil {
		additionalProperties := renameTypeSpecRefs(*typeSpec.AdditionalProperties, old, new)
		typeSpec.AdditionalProperties = &additionalProperties
	}
	if typeSpec.Items != nil {
		items := renameTypeSpecRefs(*typeSpec.Items, old, new)
		typeSpec.Items = &items
	}
	return typeSpec
}

func extendSchemas(spec schema.PackageSpec, extensions ...schema.PackageSpec) schema.PackageSpec {
	for _, extension := range extensions {
		for k, v := range extension.Resources {
			if _, found := spec.Resources[k]; found {
				log.Fatalf("resource already defined %q", k)
			}
			spec.Resources[k] = v
		}

		for k, v := range extension.Types {
			if _, found := spec.Types[k]; found {
				log.Fatalf("type already defined %q", k)
			}
			spec.Types[k] = v
		}

		for k, v := range extension.Functions {
			if _, found := spec.Functions[k]; found {
				log.Fatalf("function already defined %q", k)
			}
			spec.Functions[k] = v
		}
	}

	return spec
}

func rawMessage(v interface{}) schema.RawMessage {
	bytes, err := json.Marshal(v)
	contract.Assert(err == nil)
	return bytes
}

type Dependencies struct {
	AzureNative string `json:"@pulumi/azure-native"`
	Docker      string `json:"@pulumi/docker"`
	Pulumi      string `json:"@pulumi/pulumi"`
}

type PackageJson struct {
	Dependencies Dependencies
}

func readPackageDependencies(packageDir string) Dependencies {
	content, err := ioutil.ReadFile(path.Join(packageDir, "package.json"))
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload PackageJson
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return payload.Dependencies
}
