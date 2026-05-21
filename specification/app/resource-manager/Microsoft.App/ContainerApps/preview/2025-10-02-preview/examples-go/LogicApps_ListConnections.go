package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v5"
)

// Generated from example definition: 2025-10-02-preview/LogicApps_ListConnections.json
func ExampleLogicAppsClient_ListWorkflowsConnections() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
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
	// res = armappcontainers.LogicAppsClientListWorkflowsConnectionsResponse{
	// 	WorkflowEnvelope: armappcontainers.WorkflowEnvelope{
	// 		Name: to.Ptr("testapp2/connections"),
	// 		Type: to.Ptr("Microsoft.App/logicApps/workflowsconfiguration"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/examplerg/providers/Microsoft.App/containerApps/testapp2/providers/Microsoft.App/logicApps/testapp2/workflowconfigurations/connections"),
	// 		Properties: &armappcontainers.WorkflowEnvelopeProperties{
	// 			Files: map[string]any{
	// 				"connections.json": map[string]any{
	// 					"managedApiConnections": map[string]any{
	// 						"office365": map[string]any{
	// 							"api": map[string]any{
	// 								"id": "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Web/locations/centraluseuap/managedApis/arm",
	// 							},
	// 							"authentication": map[string]any{
	// 								"type": "Raw",
	// 								"parameter": "@appsetting('office365-connectionKey')",
	// 								"scheme": "Key",
	// 							},
	// 							"connection": map[string]any{
	// 								"id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/examplerg/providers/Microsoft.Web/connections/arm",
	// 							},
	// 							"connectionRuntimeUrl": "https://9d51d1ffc9f77572.00.common.logic.azure-apihub.net/apim/arm/connectionruntimeid",
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Health: &armappcontainers.WorkflowHealth{
	// 				State: to.Ptr(armappcontainers.WorkflowHealthStateHealthy),
	// 			},
	// 		},
	// 	},
	// }
}
