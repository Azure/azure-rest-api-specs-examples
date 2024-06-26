package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_GenerateUpgradedDefinition.json
func ExampleWorkflowsClient_GenerateUpgradedDefinition() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowsClient().GenerateUpgradedDefinition(ctx, "test-resource-group", "test-workflow", armlogic.GenerateUpgradedDefinitionParameters{
		TargetSchemaVersion: to.Ptr("2016-06-01"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Interface = map[string]any{
	// 	"$schema": "https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#",
	// 	"actions":map[string]any{
	// 	},
	// 	"contentVersion": "1.0.0.0",
	// 	"outputs":map[string]any{
	// 	},
	// 	"parameters":map[string]any{
	// 		"$connections":map[string]any{
	// 			"type": "Object",
	// 			"defaultValue":map[string]any{
	// 			},
	// 		},
	// 	},
	// 	"triggers":map[string]any{
	// 	},
	// }
}
