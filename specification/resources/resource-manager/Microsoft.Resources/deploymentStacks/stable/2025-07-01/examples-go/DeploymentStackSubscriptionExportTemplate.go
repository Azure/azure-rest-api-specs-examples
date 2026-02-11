package armdeploymentstacks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armdeploymentstacks/v2"
)

// Generated from example definition: 2025-07-01/DeploymentStackSubscriptionExportTemplate.json
func ExampleClient_ExportTemplateAtSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeploymentstacks.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewClient().ExportTemplateAtSubscription(ctx, "simpleDeploymentStack", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdeploymentstacks.ClientExportTemplateAtSubscriptionResponse{
	// 	DeploymentStackTemplateDefinition: &armdeploymentstacks.DeploymentStackTemplateDefinition{
	// 		TemplateLink: &armdeploymentstacks.TemplateLink{
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Resources/templateSpecs/templateSpec/versions/1.0"),
	// 		},
	// 		Template: map[string]any{
	// 			"$schema": "https://schema.management.azure.com/schemas/2019-04-01/deploymentTemplate.json#",
	// 			"contentVersion": "1.0.0.0",
	// 			"metadata": map[string]any{
	// 				"_generator": map[string]any{
	// 					"name": "bicep",
	// 					"version": "0.4.1008.15138",
	// 					"templateHash": "1201162276450656794",
	// 				},
	// 			},
	// 			"functions": []any{
	// 			},
	// 			"resources": []any{
	// 			},
	// 			"outputs": map[string]any{
	// 				"myOut": map[string]any{
	// 					"type": "int",
	// 					"value": 1,
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
