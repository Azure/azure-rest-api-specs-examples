package armappservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appservice/armappservice/v6"
)

// Generated from example definition: 2025-05-01/WorkflowTriggers_List.json
func ExampleWorkflowTriggersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappservice.NewClientFactory("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowTriggersClient().NewListPager("test-resource-group", "test-name", "test-workflow", nil)
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
		// page = armappservice.WorkflowTriggersClientListResponse{
		// 	WorkflowTriggerListResult: armappservice.WorkflowTriggerListResult{
		// 		Value: []*armappservice.WorkflowTrigger{
		// 			{
		// 				Name: to.Ptr("manual"),
		// 				Type: to.Ptr("Microsoft.Web/sites/workflows/triggers"),
		// 				ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-resource-group/providers/Microsoft.Web/sites/test-name/workflows/test-workflow/triggers/manual"),
		// 				Properties: &armappservice.WorkflowTriggerProperties{
		// 					ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T18:47:49.5288666Z"); return t}()),
		// 					CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-10T17:32:30.2496336Z"); return t}()),
		// 					ProvisioningState: to.Ptr(armappservice.WorkflowTriggerProvisioningStateSucceeded),
		// 					State: to.Ptr(armappservice.WorkflowStateEnabled),
		// 					Workflow: &armappservice.ResourceReference{
		// 						Name: to.Ptr("08586676800160476478"),
		// 						Type: to.Ptr("Microsoft.Web/sites/workflows/versions"),
		// 						ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-resource-group/providers/Microsoft.Web/sites/test-name/workflows/test-workflow/versions/08586676800160476478"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
