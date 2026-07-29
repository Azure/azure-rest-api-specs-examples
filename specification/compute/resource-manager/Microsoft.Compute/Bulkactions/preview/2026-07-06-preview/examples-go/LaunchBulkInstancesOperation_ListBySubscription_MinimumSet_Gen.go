package armbulkactions_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armbulkactions"
)

// Generated from example definition: 2026-07-06-preview/LaunchBulkInstancesOperation_ListBySubscription_MinimumSet_Gen.json
func ExampleLaunchBulkInstancesOperationClient_NewListBySubscriptionPager_launchBulkInstancesOperationListBySubscriptionExampleGeneratedByMinimumSetRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbulkactions.NewClientFactory("1FBA3C66-5C9C-4391-B72F-9F52735FC9F2", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLaunchBulkInstancesOperationClient().NewListBySubscriptionPager("useast2euap", nil)
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
		// page = armbulkactions.LaunchBulkInstancesOperationClientListBySubscriptionResponse{
		// 	LaunchBulkInstancesOperationListResult: armbulkactions.LaunchBulkInstancesOperationListResult{
		// 		Value: []*armbulkactions.LocationBasedLaunchBulkInstancesOperation{
		// 			{
		// 				ID: to.Ptr("/subscriptions/1FBA3C66-5C9C-4391-B72F-9F52735FC9F2/resourceGroups/rgBulkactions/providers/Microsoft.Compute/locations/useast2euap/launchBulkInstancesOperations/myBulkOperation"),
		// 				Name: to.Ptr("myBulkOperation"),
		// 				Type: to.Ptr("Microsoft.Compute/locations/launchBulkInstancesOperations"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
