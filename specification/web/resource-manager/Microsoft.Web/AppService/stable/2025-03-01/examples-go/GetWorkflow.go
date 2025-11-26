package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetWorkflow.json
func ExampleWebAppsClient_GetWorkflow() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWebAppsClient().GetWorkflow(ctx, "testrg123", "testsite2", "stateful1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowEnvelope = armappservice.WorkflowEnvelope{
	// 	Name: to.Ptr("testsite2/stateful1"),
	// 	Type: to.Ptr("Microsoft.Web/sites/workflows"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/testsite2/workflows/stateful1"),
	// 	Kind: to.Ptr("Stateful"),
	// 	Location: to.Ptr("USAAnywhere"),
	// 	Properties: &armappservice.WorkflowEnvelopeProperties{
	// 		Files: map[string]any{
	// 			"connections.json": map[string]any{
	// 				"managedApiConnections":map[string]any{
	// 					"office365":map[string]any{
	// 						"api":map[string]any{
	// 							"id": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.Web/locations/brazilsouth/managedApis/office365",
	// 						},
	// 						"authentication":map[string]any{
	// 							"type": "Raw",
	// 							"parameter": "@appsetting('office365-connectionKey')",
	// 							"scheme": "Key",
	// 						},
	// 						"connection":map[string]any{
	// 							"id": "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/connections/office365-1",
	// 						},
	// 						"connectionRuntimeUrl": "string",
	// 					},
	// 				},
	// 			},
	// 			"workflow.json": map[string]any{
	// 				"definition":map[string]any{
	// 					"$schema": "https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#",
	// 					"actions":map[string]any{
	// 					},
	// 					"contentVersion": "1.0.0.0",
	// 					"outputs":map[string]any{
	// 					},
	// 					"parameters":map[string]any{
	// 					},
	// 					"triggers":map[string]any{
	// 					},
	// 				},
	// 			},
	// 		},
	// 		FlowState: to.Ptr(armappservice.WorkflowStateEnabled),
	// 		Health: &armappservice.WorkflowHealth{
	// 			State: to.Ptr(armappservice.WorkflowHealthStateHealthy),
	// 		},
	// 	},
	// }
}
