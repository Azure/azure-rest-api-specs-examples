package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowVersions_Get.json
func ExampleWorkflowVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowVersionsClient().Get(ctx, "test-resource-group", "test-name", "test-workflow", "08586676824806722526", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowVersion = armappservice.WorkflowVersion{
	// 	Name: to.Ptr("test-workflow"),
	// 	Type: to.Ptr("/workflows/versions"),
	// 	ID: to.Ptr("/workflows/test-workflow"),
	// 	Properties: &armappservice.WorkflowVersionProperties{
	// 		ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-09T22:54:54.353Z"); return t}()),
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-25T01:39:21.436Z"); return t}()),
	// 		Definition: map[string]any{
	// 			"$schema": "https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#",
	// 			"actions":map[string]any{
	// 			},
	// 			"contentVersion": "1.0.0.0",
	// 			"outputs":map[string]any{
	// 			},
	// 			"parameters":map[string]any{
	// 			},
	// 			"triggers":map[string]any{
	// 			},
	// 		},
	// 		EndpointsConfiguration: &armappservice.FlowEndpointsConfiguration{
	// 			Connector: &armappservice.FlowEndpoints{
	// 				OutgoingIPAddresses: []*armappservice.IPAddress{
	// 				},
	// 			},
	// 			Workflow: &armappservice.FlowEndpoints{
	// 				AccessEndpointIPAddresses: []*armappservice.IPAddress{
	// 				},
	// 				OutgoingIPAddresses: []*armappservice.IPAddress{
	// 				},
	// 			},
	// 		},
	// 		Parameters: map[string]*armappservice.WorkflowParameter{
	// 		},
	// 		Version: to.Ptr("08586677515911718341"),
	// 	},
	// }
}
