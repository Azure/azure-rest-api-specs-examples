package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: 2026-07-01-preview/Queues/SBQueueListByNameSpace.json
func ExampleQueuesClient_NewListByNamespacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewQueuesClient().NewListByNamespacePager("ArunMonocle", "sdk-Namespace-3174", nil)
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
		// page = armservicebus.QueuesClientListByNamespaceResponse{
		// 	SBQueueListResult: armservicebus.SBQueueListResult{
		// 		Value: []*armservicebus.SBQueue{
		// 			{
		// 				Name: to.Ptr("sdk-queues-5647"),
		// 				Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Queues"),
		// 				ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-3174/queues/sdk-queues-5647"),
		// 				Properties: &armservicebus.SBQueueProperties{
		// 					AccessedAt: to.Ptr(time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)),
		// 					AutoDeleteOnIdle: to.Ptr("P10675199DT2H48M5.4775807S"),
		// 					CreatedAt: to.Ptr(time.Date(2017, time.May, 26, 18, 7, 32, 459293100, time.UTC)),
		// 					DefaultMessageTimeToLive: to.Ptr("P10675199DT2H48M5.4775807S"),
		// 					DuplicateDetectionHistoryTimeWindow: to.Ptr("PT10M"),
		// 					EnableExpress: to.Ptr(false),
		// 					EnablePartitioning: to.Ptr(true),
		// 					LockDuration: to.Ptr("PT1M"),
		// 					MaxDeliveryCount: to.Ptr[int32](10),
		// 					MaxMessageSizeInKilobytes: to.Ptr[int64](10240),
		// 					MaxSizeInMegabytes: to.Ptr[int32](163840),
		// 					MessageCount: to.Ptr[int64](0),
		// 					RequiresDuplicateDetection: to.Ptr(false),
		// 					RequiresSession: to.Ptr(false),
		// 					SizeInBytes: to.Ptr[int64](0),
		// 					Status: to.Ptr(armservicebus.EntityStatusActive),
		// 					UpdatedAt: to.Ptr(time.Date(2017, time.May, 26, 18, 7, 34, 624376100, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
