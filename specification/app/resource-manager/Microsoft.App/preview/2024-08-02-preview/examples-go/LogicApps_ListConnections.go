package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/LogicApps_ListConnections.json
func ExampleLogicAppsClient_ListWorkflowsConnections() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLogicAppsClient().ListWorkflowsConnections(ctx, "testrg123", "testapp2", "testapp2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowEnvelope = armappcontainers.WorkflowEnvelope{
	// 	Name: to.Ptr("testapp2/connections"),
	// 	Type: to.Ptr("Microsoft.App/logicApps/workflowsconfiguration"),
	// 	ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/containerApps/testapp2/providers/Microsoft.App/logicApps/testapp2/workflowconfigurations/connections"),
	// 	Location: to.Ptr("East US"),
	// 	Properties: &armappcontainers.WorkflowEnvelopeProperties{
	// 		Files: map[string]any{
	// 			"connections.json":map[string]any{
	// 				"managedApiConnections":map[string]any{
	// 					"office365":map[string]any{
	// 						"api":map[string]any{
	// 							"id": "string",
	// 						},
	// 						"authentication":map[string]any{
	// 							"type": "Raw",
	// 							"parameter": "@appsetting('office365-connectionKey')",
	// 							"scheme": "Key",
	// 						},
	// 						"connection":map[string]any{
	// 							"id": "string",
	// 						},
	// 						"connectionRuntimeUrl": "string",
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Health: &armappcontainers.WorkflowHealth{
	// 			State: to.Ptr(armappcontainers.WorkflowHealthStateHealthy),
	// 		},
	// 	},
	// }
}
