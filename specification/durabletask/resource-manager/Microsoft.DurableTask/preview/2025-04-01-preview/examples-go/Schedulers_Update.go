package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2025-04-01-preview/Schedulers_Update.json
func ExampleSchedulersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("EE9BD735-67CE-4A90-89C4-439D3F6A4C93", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSchedulersClient().BeginUpdate(ctx, "rgopenapi", "testscheduler", armdurabletask.SchedulerUpdate{
		Tags: map[string]*string{
			"hello": to.Ptr("world"),
		},
		Properties: &armdurabletask.SchedulerPropertiesUpdate{
			IPAllowlist: []*string{
				to.Ptr("10.0.0.0/8"),
			},
			SKU: &armdurabletask.SchedulerSKUUpdate{
				Name:     to.Ptr("Dedicated"),
				Capacity: to.Ptr[int32](3),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdurabletask.SchedulersClientUpdateResponse{
	// 	Scheduler: &armdurabletask.Scheduler{
	// 		Location: to.Ptr("northcentralus"),
	// 		Properties: &armdurabletask.SchedulerProperties{
	// 			ProvisioningState: to.Ptr(armdurabletask.ProvisioningStateSucceeded),
	// 			Endpoint: to.Ptr("https://test.northcentralus.1.durabletask.io"),
	// 			IPAllowlist: []*string{
	// 				to.Ptr("10.0.0.0/8"),
	// 			},
	// 			SKU: &armdurabletask.SchedulerSKU{
	// 				Name: to.Ptr("Dedicated"),
	// 				Capacity: to.Ptr[int32](3),
	// 				RedundancyState: to.Ptr(armdurabletask.RedundancyStateZone),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"department": to.Ptr("research"),
	// 			"development": to.Ptr("true"),
	// 			"hello": to.Ptr("world"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/EE9BD735-67CE-4A90-89C4-439D3F6A4C93/resourceGroups/rgopenapi/providers/Microsoft.DurableTask/schedulers/testscheduler"),
	// 		Name: to.Ptr("testscheduler"),
	// 		Type: to.Ptr("Microsoft.DurableTask/schedulers"),
	// 		SystemData: &armdurabletask.SystemData{
	// 			CreatedBy: to.Ptr("tenmbevaunjzikxowqexrsx"),
	// 			CreatedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-17T15:34:17.365Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("xfvdcegtj"),
	// 			LastModifiedByType: to.Ptr(armdurabletask.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-17T15:34:17.366Z"); return t}()),
	// 		},
	// 	},
	// }
}
