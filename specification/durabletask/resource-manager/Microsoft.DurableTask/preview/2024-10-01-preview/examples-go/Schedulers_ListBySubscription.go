package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2024-10-01-preview/Schedulers_ListBySubscription.json
func ExampleSchedulersClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("EE9BD735-67CE-4A90-89C4-439D3F6A4C93", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSchedulersClient().NewListBySubscriptionPager(nil)
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
		// page = armdurabletask.SchedulersClientListBySubscriptionResponse{
		// 	SchedulerListResult: armdurabletask.SchedulerListResult{
		// 		Value: []*armdurabletask.Scheduler{
		// 			{
		// 				ID: to.Ptr("/subscriptions/EE9BD735-67CE-4A90-89C4-439D3F6A4C93/resourceGroups/rgopenapi/providers/Microsoft.DurableTask/schedulers/testscheduler"),
		// 				Location: to.Ptr("northcentralus"),
		// 				Properties: &armdurabletask.SchedulerProperties{
		// 					ProvisioningState: to.Ptr(armdurabletask.ProvisioningStateSucceeded),
		// 					Endpoint: to.Ptr("https://test.northcentralus.1.durabletask.io"),
		// 					IPAllowlist: []*string{
		// 						to.Ptr("10.0.0.0/8"),
		// 					},
		// 					SKU: &armdurabletask.SchedulerSKU{
		// 						Name: to.Ptr("Dedicated"),
		// 						Capacity: to.Ptr[int32](3),
		// 						RedundancyState: to.Ptr(armdurabletask.RedundancyStateZone),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key2138": to.Ptr("fjaeecgnvqd"),
		// 				},
		// 				Name: to.Ptr("testscheduler"),
		// 				Type: to.Ptr("tglszobjxdhauglzxuqozeg"),
		// 				SystemData: &armdurabletask.SystemData{
		// 					CreatedBy: to.Ptr("tenmbevaunjzikxowqexrsx"),
		// 					CreatedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-17T15:34:17.365Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("xfvdcegtj"),
		// 					LastModifiedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-17T15:34:17.366Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
