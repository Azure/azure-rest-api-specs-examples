package armdeploymentstacks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentstacks"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f4c6c8697c59f966db0d1e36b62df3af3bca9065/specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/DeploymentStackExportTemplate.json
func ExampleClient_ExportTemplateAtResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentstacks.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ExportTemplateAtResourceGroup(ctx, "deploymentStacksRG", "simpleDeploymentStack", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeploymentStackTemplateDefinition = armdeploymentstacks.DeploymentStackTemplateDefinition{
	// 	Template: map[string]any{
	// 		"$schema": "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
	// 		"contentVersion": "1.0.0.0",
	// 		"functions":[]any{
	// 		},
	// 		"metadata":map[string]any{
	// 			"_generator":map[string]any{
	// 				"name": "bicep",
	// 				"templateHash": "1201162276450656794",
	// 				"version": "0.4.1008.15138",
	// 			},
	// 		},
	// 		"outputs":map[string]any{
	// 			"myOut":map[string]any{
	// 				"type": "int",
	// 				"value": float64(1),
	// 			},
	// 		},
	// 		"resources":[]any{
	// 		},
	// 	},
	// 	TemplateLink: &armdeploymentstacks.TemplateLink{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/deploymentStacksRG/providers/Microsoft.Resources/templateSpecs/templateSpec/versions/1.0"),
	// 	},
	// }
}
