package armlogic_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggers_List.json
func ExampleWorkflowTriggersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowTriggersClient().NewListPager("test-resource-group", "test-workflow", &armlogic.WorkflowTriggersClientListOptions{Top: nil,
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
		// page.WorkflowTriggerListResult = armlogic.WorkflowTriggerListResult{
		// 	Value: []*armlogic.WorkflowTrigger{
		// 		{
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-resource-group/providers/Microsoft.Logic/workflows/test-workflow/triggers/manual"),
		// 			Name: to.Ptr("manual"),
		// 			Type: to.Ptr("Microsoft.Logic/workflows/triggers"),
		// 			Properties: &armlogic.WorkflowTriggerProperties{
		// 				ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T18:47:49.528Z"); return t}()),
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T17:32:30.249Z"); return t}()),
		// 				ProvisioningState: to.Ptr(armlogic.WorkflowTriggerProvisioningStateSucceeded),
		// 				State: to.Ptr(armlogic.WorkflowStateEnabled),
		// 				Workflow: &armlogic.ResourceReference{
		// 					Name: to.Ptr("08586676800160476478"),
		// 					Type: to.Ptr("Microsoft.Logic/workflows/versions"),
		// 					ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-resource-group/providers/Microsoft.Logic/workflows/test-workflow/versions/08586676800160476478"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
