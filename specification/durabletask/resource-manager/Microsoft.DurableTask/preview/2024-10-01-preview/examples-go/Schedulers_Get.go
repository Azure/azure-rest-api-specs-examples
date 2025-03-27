package armdurabletask_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/durabletask/armdurabletask"
)

// Generated from example definition: 2024-10-01-preview/Schedulers_Get.json
func ExampleSchedulersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdurabletask.NewClientFactory("EE9BD735-67CE-4A90-89C4-439D3F6A4C93", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSchedulersClient().Get(ctx, "rgopenapi", "testscheduler", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdurabletask.SchedulersClientGetResponse{
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
	// 			"key7131": to.Ptr("ryohwcoiccwsnewjigfmijz"),
	// 			"key2138": to.Ptr("fjaeecgnvqd"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/EE9BD735-67CE-4A90-89C4-439D3F6A4C93/resourceGroups/rgopenapi/providers/Microsoft.DurableTask/schedulers/testscheduler"),
	// 		Name: to.Ptr("testscheduler"),
	// 		Type: to.Ptr("vwqdbpynxwfhiopdypuabwvfohnr"),
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
