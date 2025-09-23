package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1d2097f1ed03e8a61eed4fe63602a641bedd77ae/specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/LogicApps_ListWorkflows.json
func ExampleLogicAppsClient_NewListWorkflowsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLogicAppsClient().NewListWorkflowsPager("examplerg", "testcontainerApp0", "testcontainerApp0", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.WorkflowEnvelopeCollection = armappcontainers.WorkflowEnvelopeCollection{
		// 	Value: []*armappcontainers.WorkflowEnvelope{
		// 		{
		// 			Name: to.Ptr("testcontainerApp0/a1"),
		// 			Type: to.Ptr("Microsoft.App/logicApps/workflows"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/containerApps/testcontainerApp0/providers/Microsoft.App/logicApps/testcontainerApp0/workflows/a1"),
		// 			Kind: to.Ptr("Stateful"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.WorkflowEnvelopeProperties{
		// 				FlowState: to.Ptr(armappcontainers.WorkflowStateEnabled),
		// 				Health: &armappcontainers.WorkflowHealth{
		// 					State: to.Ptr(armappcontainers.WorkflowHealthStateHealthy),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testcontainerApp0/stateful2"),
		// 			Type: to.Ptr("Microsoft.App/logicApps/workflows"),
		// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/containerApps/testcontainerApp0/providers/Microsoft.App/logicApps/testcontainerApp0/workflows/stateful2"),
		// 			Kind: to.Ptr("Stateful"),
		// 			Location: to.Ptr("East US"),
		// 			Properties: &armappcontainers.WorkflowEnvelopeProperties{
		// 				FlowState: to.Ptr(armappcontainers.WorkflowStateEnabled),
		// 				Health: &armappcontainers.WorkflowHealth{
		// 					Error: &armappcontainers.ErrorEntity{
		// 						Code: to.Ptr("InvalidWorkflowJson"),
		// 						Message: to.Ptr("Invalid character after parsing property name. Expected ':' but got: \". Path '', line 2, position 2."),
		// 					},
		// 					State: to.Ptr(armappcontainers.WorkflowHealthStateUnhealthy),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
