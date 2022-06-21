// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gen

import (
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
)

const (
	WebApp         string = "azure-justrun:index:webapp"
	PublicAccess   string = "azure-justrun:index:PublicAccess"
	StorageSkuName string = "azure-justrun:index:StorageSkuName"
	ContainerApp   string = "azure-justrun:index:containerapp"
)

func generateSchema(azureNativeSpec schema.PackageSpec, dockerSpec schema.PackageSpec) schema.PackageSpec {
	return schema.PackageSpec{
		Resources: map[string]schema.ResourceSpec{
			WebApp:       webappResourceSpec(azureNativeSpec),
			ContainerApp: containerappResourceSpec(azureNativeSpec),
		},
		Types: map[string]schema.ComplexTypeSpec{
			PublicAccess:   publicAccessTypeSpec(azureNativeSpec),
			StorageSkuName: skuNameTypeSpec(azureNativeSpec),
		},
		Functions: map[string]schema.FunctionSpec{},
	}
}

func publicAccessTypeSpec(azureNativeSpec schema.PackageSpec) schema.ComplexTypeSpec {
	return azureNativeSpec.Types["azure-native:storage:PublicAccess"]
}

func skuNameTypeSpec(azureNativeSpec schema.PackageSpec) schema.ComplexTypeSpec {
	return azureNativeSpec.Types["azure-native:storage:SkuName"]
}

func webappResourceSpec(azureNativeSpec schema.PackageSpec) schema.ResourceSpec {
	spec := schema.ResourceSpec{
		IsComponent: true,
		InputProperties: map[string]schema.PropertySpec{
			"appSkuName": stringProperty("The tier of the compute instance running the server. Also see appSkuName"),
			"appSkuTier": stringProperty("The name of the compute instance running the server. Also see appSkuTier"),
			"filePath":   stringPropertyDefault("The relative file path to the folder containing web files.", "./www"),
			"containerPublicAccess": schema.PropertySpec{
				TypeSpec: schema.TypeSpec{
					Ref: selfTypeRef(PublicAccess),
				},
				Description: "The public access level of the BlobContainer containg the website data.",
			},
			"storageSkuName": schema.PropertySpec{
				TypeSpec: schema.TypeSpec{
					Ref: selfTypeRef(StorageSkuName),
				},
				Description: "The SKU name of the storage account created, if storageAccount is not provided",
			},
			"storageAccountName": stringProperty("The name of the storage account to use. One will be created if not provided."),
			"resourceGroupName":  stringProperty("The resource group to use. One will be created if not provided."),
			"namePrefix":         stringProperty("The name prefix given to child resources of this component. Should not contain dashes."),
		},
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "This represents a web app component resource",
			Properties: map[string]schema.PropertySpec{
				"url": stringProperty("The URL of the web app"),
			},
		},
	}
	return spec
}

func containerappResourceSpec(azureNativeSpec schema.PackageSpec) schema.ResourceSpec {
	spec := schema.ResourceSpec{
		IsComponent: true,
		InputProperties: map[string]schema.PropertySpec{
			"registryName":       stringProperty("The name of the image registry. Must belong to the resource group specified in ResourceGroupName. One will be created if not provided."),
			"dockerImageName":    stringProperty("The name of the docker image. One will be created if not provided"),
			"imageDirectory":     stringProperty("The name of the directory where the docker image to be created is. NOT the actual directory, i.e. 'nodeapp' instead of './nodeapp'"),
			"version":            stringProperty("The version of the created docker image"),
			"storageAccountName": stringProperty("The name of the storage account to use. One will be created if not provided."),
			"resourceGroupName":  stringProperty("The resource group to use. One will be created if not provided."),
			"namePrefix":         stringProperty("The name prefix given to child resources of this component. Should not contain dashes."),
		},
		ObjectTypeSpec: schema.ObjectTypeSpec{
			Description: "This represents a container app component resource",
			Properties: map[string]schema.PropertySpec{
				"url": stringProperty("The URL of the container app"),
			},
		},
	}
	return spec
}
func stringProperty(description string) schema.PropertySpec {
	return schema.PropertySpec{
		TypeSpec: schema.TypeSpec{
			Type: "string",
		},
		Description: description,
	}
}

func stringPropertyDefault(description string, def string) schema.PropertySpec {
	return schema.PropertySpec{
		TypeSpec: schema.TypeSpec{
			Type: "string",
		},
		Description: description,
		Default:     def,
	}
}

func selfTypeRef(token string) string {
	return ref("", "types", token)
}

func ref(pathtoschema string, prefix string, token string) string {
	return pathtoschema + "#/" + prefix + "/" + token
}
