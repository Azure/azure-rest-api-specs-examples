package armcomputeschedule_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computeschedule/armcomputeschedule/v2"
)

// Generated from example definition: 2026-04-15-preview/ScheduledActionExtension_ListByVms_MinimumSet_Gen.json
func ExampleScheduledActionExtensionClient_NewListByVMsPager_scheduledActionExtensionListByVmsMinimumSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputeschedule.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScheduledActionExtensionClient().NewListByVMsPager("subscriptions/732116BD-AF31-4E74-9283-B387C44B4A44/resourceGroups/rgcomputeschedule/providers/Microsoft.Compute/virtualMachines/myVm", nil)
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
		// page = armcomputeschedule.ScheduledActionExtensionClientListByVMsResponse{
		// 	ScheduledActionResourcesListResult: armcomputeschedule.ScheduledActionResourcesListResult{
		// 		Value: []*armcomputeschedule.ScheduledActionResources{
		// 			{
		// 				ID: to.Ptr("/subscriptions/83C27AB3-A7B9-498B-B165-D9440661474F/resourceGroups/myRg/providers/Microsoft.ComputeSchedule/scheduledActions/myScheduledAction"),
		// 				Name: to.Ptr("myScheduledAction"),
		// 				Type: to.Ptr("Microsoft.ComputeSchedule/scheduledActions"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
