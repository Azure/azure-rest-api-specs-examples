package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f4cb2884f1948b879ecfb3f410e8cbc8805c213/specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/WorkflowTriggers_List.json
func ExampleWorkflowTriggersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowTriggersClient().NewListPager("test-resource-group", "test-name", "test-workflow", &armappservice.WorkflowTriggersClientListOptions{Top: nil,
		Filter: nil,
	})
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
		// page.WorkflowTriggerListResult = armappservice.WorkflowTriggerListResult{
		// 	Value: []*armappservice.WorkflowTrigger{
		// 		{
		// 			ID: to.Ptr("/workflows/test-workflow/triggers/manual"),
		// 			Name: to.Ptr("manual"),
		// 			Type: to.Ptr("/workflows/triggers"),
		// 			Properties: &armappservice.WorkflowTriggerProperties{
		// 				ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T18:47:49.528Z"); return t}()),
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T17:32:30.249Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armappservice.WorkflowTriggerProvisioningStateSucceeded),
		// 				State: to.Ptr(armappservice.WorkflowStateEnabled),
		// 				Workflow: &armappservice.ResourceReference{
		// 					Name: to.Ptr("08586676800160476478"),
		// 					Type: to.Ptr("/workflows/versions"),
		// 					ID: to.Ptr("/workflows/test-workflow/versions/08586676800160476478"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
