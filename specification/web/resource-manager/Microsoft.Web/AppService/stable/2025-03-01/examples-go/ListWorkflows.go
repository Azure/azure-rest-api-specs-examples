package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListWorkflows.json
func ExampleWebAppsClient_NewListWorkflowsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebAppsClient().NewListWorkflowsPager("testrg123", "testsite2", nil)
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
		// page.WorkflowEnvelopeCollection = armappservice.WorkflowEnvelopeCollection{
		// 	Value: []*armappservice.WorkflowEnvelope{
		// 		{
		// 			Name: to.Ptr("testsite2/a1"),
		// 			Type: to.Ptr("Microsoft.Web/sites/workflows"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testrg123/providers/Microsoft.Web/sites/testsite2/workflows/a1"),
		// 			Kind: to.Ptr("Stateful"),
		// 			Location: to.Ptr("USAAnywhere"),
		// 			Properties: &armappservice.WorkflowEnvelopeProperties{
		// 				FlowState: to.Ptr(armappservice.WorkflowStateEnabled),
		// 				Health: &armappservice.WorkflowHealth{
		// 					State: to.Ptr(armappservice.WorkflowHealthStateHealthy),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testsite2/stateful2"),
		// 			Type: to.Ptr("Microsoft.Web/sites/workflows"),
		// 			ID: to.Ptr("/subscriptions/testsub/resourceGroups/testrg/providers/Microsoft.Web/sites/testsite2/workflows/stateful2"),
		// 			Kind: to.Ptr("Stateful"),
		// 			Location: to.Ptr("USAAnywhere"),
		// 			Properties: &armappservice.WorkflowEnvelopeProperties{
		// 				FlowState: to.Ptr(armappservice.WorkflowStateEnabled),
		// 				Health: &armappservice.WorkflowHealth{
		// 					Error: &armappservice.ErrorEntity{
		// 						Code: to.Ptr("InvalidWorkflowJson"),
		// 						Message: to.Ptr("Invalid character after parsing property name. Expected ':' but got: \". Path '', line 2, position 2."),
		// 					},
		// 					State: to.Ptr(armappservice.WorkflowHealthStateUnhealthy),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
