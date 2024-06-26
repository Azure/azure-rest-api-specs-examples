package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_Get.json
func ExampleWorkflowsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkflowsClient().Get(ctx, "test-resource-group", "test-workflow", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Workflow = armlogic.Workflow{
	// 	Name: to.Ptr("test-workflow"),
	// 	Type: to.Ptr("Microsoft.Logic/workflows"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-resource-group/providers/Microsoft.Logic/workflows/test-workflow"),
	// 	Location: to.Ptr("brazilsouth"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armlogic.WorkflowProperties{
	// 		AccessEndpoint: to.Ptr("http://tempuri.org"),
	// 		ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-09T22:54:54.353Z"); return t}()),
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-25T01:39:21.436Z"); return t}()),
	// 		Definition: map[string]any{
	// 			"$schema": "https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2018-07-01-preview/workflowdefinition.json#",
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
	// 		EndpointsConfiguration: &armlogic.FlowEndpointsConfiguration{
	// 			Connector: &armlogic.FlowEndpoints{
	// 				OutgoingIPAddresses: []*armlogic.IPAddress{
	// 					{
	// 						Address: to.Ptr("40.84.145.61"),
	// 				}},
	// 			},
	// 			Workflow: &armlogic.FlowEndpoints{
	// 				AccessEndpointIPAddresses: []*armlogic.IPAddress{
	// 					{
	// 						Address: to.Ptr("104.210.153.89"),
	// 					},
	// 					{
	// 						Address: to.Ptr("13.85.79.155"),
	// 					},
	// 					{
	// 						Address: to.Ptr("13.65.39.247"),
	// 				}},
	// 				OutgoingIPAddresses: []*armlogic.IPAddress{
	// 					{
	// 						Address: to.Ptr("13.84.159.168"),
	// 					},
	// 					{
	// 						Address: to.Ptr("13.65.86.56"),
	// 					},
	// 					{
	// 						Address: to.Ptr("13.65.82.190"),
	// 				}},
	// 			},
	// 		},
	// 		IntegrationAccount: &armlogic.ResourceReference{
	// 			Name: to.Ptr("test-integration-account"),
	// 			Type: to.Ptr("Microsoft.Logic/integrationAccounts"),
	// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-resource-group/providers/Microsoft.Logic/integrationAccounts/test-integration-account"),
	// 		},
	// 		IntegrationServiceEnvironment: &armlogic.ResourceReference{
	// 			Name: to.Ptr("test-integration-service-environment"),
	// 			Type: to.Ptr("Microsoft.Logic/integrationServiceEnvironments"),
	// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-resource-group/providers/Microsoft.Logic/integrationServiceEnvironments/test-integration-service-environment"),
	// 		},
	// 		Parameters: map[string]*armlogic.WorkflowParameter{
	// 		},
	// 		ProvisioningState: to.Ptr(armlogic.WorkflowProvisioningStateSucceeded),
	// 		State: to.Ptr(armlogic.WorkflowStateEnabled),
	// 		Version: to.Ptr("08586677515911718341"),
	// 	},
	// }
}
