package armcomputebulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/computebulkactions/armcomputebulkactions"
)

// Generated from example definition: 2026-02-01-preview/BulkActions_ListBySubscription_MinimumSet_Gen.json
func ExampleBulkActionsClient_NewListBySubscriptionPager_bulkActionsListBySubscriptionGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcomputebulkactions.NewClientFactory("50352BBD-59F1-4EE2-BA9C-A6E51B0B1B2B", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBulkActionsClient().NewListBySubscriptionPager("eastus2euap", nil)
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
		// page = armcomputebulkactions.BulkActionsClientListBySubscriptionResponse{
		// 	LaunchBulkInstancesOperationListResult: armcomputebulkactions.LaunchBulkInstancesOperationListResult{
		// 		Value: []*armcomputebulkactions.LocationBasedLaunchBulkInstancesOperation{
		// 			{
		// 				ID: to.Ptr("/subscriptions/50352BBD-59F1-4EE2-BA9C-A6E51B0B1B2B/resourceGroups/rgcomputebulkactions/providers/Microsoft.Compute/bulkActions/3ec2ab23-9f13-4328-85c8-21928acbc7b8"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
