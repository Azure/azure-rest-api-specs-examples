package armtemplatespecs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armtemplatespecs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-05-01/examples/TemplateSpecVersionsCreate.json
func ExampleTemplateSpecVersionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armtemplatespecs.NewTemplateSpecVersionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<template-spec-name>",
		"<template-spec-version>",
		armtemplatespecs.TemplateSpecVersion{
			Location: to.Ptr("<location>"),
			Properties: &armtemplatespecs.TemplateSpecVersionProperties{
				Description: to.Ptr("<description>"),
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
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
