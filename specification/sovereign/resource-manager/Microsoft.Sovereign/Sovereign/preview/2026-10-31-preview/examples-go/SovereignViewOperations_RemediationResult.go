package armsovereign_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sovereign/armsovereign"
)

// Generated from example definition: 2026-10-31-preview/SovereignViewOperations_RemediationResult.json
func ExampleViewOperationsClient_RemediationResult() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsovereign.NewClientFactory("22b62f96-c8d5-4131-97f8-f820a5a3c55f", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewViewOperationsClient().RemediationResult(ctx, "ExampleResourceGroup", "ExampleSV", armsovereign.ViewRemediationResultContent{
		RemediationID: to.Ptr("2f30c9e8-1b4d-4b4f-a17d-54db32e132f4"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsovereign.ViewOperationsClientRemediationResultResponse{
	// 	ViewRemediationResult: armsovereign.ViewRemediationResult{
	// 		RemediationID: to.Ptr("2f30c9e8-1b4d-4b4f-a17d-54db32e132f4"),
	// 		DimensionType: to.Ptr(armsovereign.ComplianceDimensionDataResidency),
	// 		Status: to.Ptr(armsovereign.ViewRemediationStatusSucceeded),
	// 		Results: []*armsovereign.ViewRemediationResultItem{
	// 			{
	// 				ResourceIDs: []*string{
	// 					to.Ptr("/subscriptions/22b62f96-c8d5-4131-97f8-f820a5a3c55f/resourceGroups/ExampleResourceGroup/providers/Microsoft.Storage/storageAccounts/examplestorage"),
	// 				},
	// 				Type: to.Ptr(armsovereign.ViewRemediationPayloadTypeArmTemplateBundle),
	// 				Result: &armsovereign.ViewArmTemplateBundle{
	// 					Template: to.Ptr("{\"$schema\":\"https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#\",\"contentVersion\":\"1.0.0.0\",\"parameters\":{\"storageAccountName\":{\"type\":\"string\"}},\"resources\":[{\"type\":\"Microsoft.Storage/storageAccounts\",\"apiVersion\":\"2025-01-01\",\"name\":\"[parameters('storageAccountName')]\",\"location\":\"swedencentral\",\"sku\":{\"name\":\"Standard_LRS\"},\"kind\":\"StorageV2\",\"properties\":{}}]}"),
	// 					Parameters: to.Ptr("{\"storageAccountName\":{\"value\":\"examplestorage\"}}"),
	// 				},
	// 			},
	// 		},
	// 		CreatedAt: to.Ptr(time.Date(2026, time.October, 15, 10, 0, 0, 0, time.UTC)),
	// 		ExpiresAt: to.Ptr(time.Date(2026, time.October, 22, 10, 0, 0, 0, time.UTC)),
	// 	},
	// }
}
