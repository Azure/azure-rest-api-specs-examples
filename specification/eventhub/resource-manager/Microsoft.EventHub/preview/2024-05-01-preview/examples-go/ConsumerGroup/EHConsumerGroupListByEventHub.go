package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5759c77eee2d57bdb9e47aa1805d0ffb61704f2d/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2024-05-01-preview/examples/ConsumerGroup/EHConsumerGroupListByEventHub.json
func ExampleConsumerGroupsClient_NewListByEventHubPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConsumerGroupsClient().NewListByEventHubPager("ArunMonocle", "sdk-Namespace-2661", "sdk-EventHub-6681", &armeventhub.ConsumerGroupsClientListByEventHubOptions{Skip: nil,
		Top: nil,
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
		// page.ConsumerGroupListResult = armeventhub.ConsumerGroupListResult{
		// 	Value: []*armeventhub.ConsumerGroup{
		// 		{
		// 			Name: to.Ptr("$Default"),
		// 			Type: to.Ptr("Microsoft.EventHub/Namespaces/EventHubs/ConsumerGroups"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.EventHub/namespaces/sdk-Namespace-2661/eventhubs/sdk-EventHub-6681/consumergroups/$Default"),
		// 			Properties: &armeventhub.ConsumerGroupProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-25T03:42:52.287Z"); return t}()),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-25T03:42:52.287Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sdk-consumergroup-5563"),
		// 			Type: to.Ptr("Microsoft.EventHub/Namespaces/EventHubs/ConsumerGroups"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.EventHub/namespaces/sdk-Namespace-2661/eventhubs/sdk-EventHub-6681/consumergroups/sdk-consumergroup-5563"),
		// 			Properties: &armeventhub.ConsumerGroupProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-25T03:43:09.314Z"); return t}()),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-25T03:43:09.314Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
