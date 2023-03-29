package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4fd842fb73656039ec94ce367bcedee25a57bd18/specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/ExportResourceGroupWithFiltering.json
func ExampleResourceGroupsClient_BeginExportTemplate_exportAResourceGroupWithFiltering() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armresources.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewResourceGroupsClient().BeginExportTemplate(ctx, "my-resource-group", armresources.ExportTemplateRequest{
		Options: to.Ptr("SkipResourceNameParameterization"),
		Resources: []*string{
			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/my-resource-group/providers/My.RP/myResourceType/myFirstResource")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ResourceGroupExportResult = armresources.ResourceGroupExportResult{
	// 	Template: map[string]any{
	// 		"$schema": "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
	// 		"contentVersion": "1.0.0.0",
	// 		"parameters":map[string]any{
	// 			"myResourceType_myFirstResource_secret":map[string]any{
	// 				"type": "SecureString",
	// 				"defaultValue": nil,
	// 			},
	// 		},
	// 		"resources":[]any{
	// 			map[string]any{
	// 				"name": "myFirstResource",
	// 				"type": "My.RP/myResourceType",
	// 				"apiVersion": "2019-01-01",
	// 				"location": "West US",
	// 				"properties":map[string]any{
	// 					"secret": "[parameters('myResourceType_myFirstResource_secret')]",
	// 				},
	// 			},
	// 		},
	// 		"variables":map[string]any{
	// 		},
	// 	},
	// }
}
