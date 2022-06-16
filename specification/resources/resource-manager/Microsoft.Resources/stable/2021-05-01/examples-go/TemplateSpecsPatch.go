package armtemplatespecs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armtemplatespecs"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/resources/resource-manager/Microsoft.Resources/stable/2021-05-01/examples/TemplateSpecsPatch.json
func ExampleClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armtemplatespecs.NewClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<template-spec-name>",
		&armtemplatespecs.ClientUpdateOptions{TemplateSpec: &armtemplatespecs.TemplateSpecUpdateModel{
			Tags: map[string]*string{
				"myTag": to.Ptr("My Value"),
			},
		},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
