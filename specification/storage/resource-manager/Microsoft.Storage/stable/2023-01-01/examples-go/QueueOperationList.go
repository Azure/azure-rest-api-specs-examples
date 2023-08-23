package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/QueueOperationList.json
func ExampleQueueClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewQueueClient().NewListPager("res9290", "sto328", &armstorage.QueueClientListOptions{Maxpagesize: nil,
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
		// page.ListQueueResource = armstorage.ListQueueResource{
		// 	Value: []*armstorage.ListQueue{
		// 		{
		// 			Name: to.Ptr("queue6185"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/queueServices/queues"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/queueServices/default/queues/queue6185"),
		// 			QueueProperties: &armstorage.ListQueueProperties{
		// 				Metadata: map[string]*string{
		// 					"sample1": to.Ptr("meta1"),
		// 					"sample2": to.Ptr("meta2"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("queue6186"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts/queueServices/queues"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res3376/providers/Microsoft.Storage/storageAccounts/sto328/queueServices/default/queues/queue6186"),
		// 			QueueProperties: &armstorage.ListQueueProperties{
		// 				Metadata: map[string]*string{
		// 					"sample1": to.Ptr("meta1"),
		// 					"sample2": to.Ptr("meta2"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
