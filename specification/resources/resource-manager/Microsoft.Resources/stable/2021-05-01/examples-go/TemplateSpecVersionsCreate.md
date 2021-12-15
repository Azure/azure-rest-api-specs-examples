Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmtemplatespecs%2Fv0.1.1/sdk/resourcemanager/resources/armtemplatespecs/README.md) on how to add the SDK to your project and authenticate.

```go
package armtemplatespecs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armtemplatespecs"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-05-01/examples/TemplateSpecVersionsCreate.json
func ExampleTemplateSpecVersionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armtemplatespecs.NewTemplateSpecVersionsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<template-spec-name>",
		"<template-spec-version>",
		armtemplatespecs.TemplateSpecVersion{
			Location: to.StringPtr("<location>"),
			Properties: &armtemplatespecs.TemplateSpecVersionProperties{
				Description: to.StringPtr("<description>"),
				MainTemplate: map[string]interface{}{
					"$schema":        "http://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
					"contentVersion": "1.0.0.0",
					"parameters":     map[string]interface{}{},
					"resources":      []interface{}{},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("TemplateSpecVersion.ID: %s\n", *res.ID)
}
```
