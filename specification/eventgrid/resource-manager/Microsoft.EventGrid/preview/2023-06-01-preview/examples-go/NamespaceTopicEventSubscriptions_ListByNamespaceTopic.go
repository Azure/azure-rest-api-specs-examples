package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/NamespaceTopicEventSubscriptions_ListByNamespaceTopic.json
func ExampleNamespaceTopicEventSubscriptionsClient_NewListByNamespaceTopicPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNamespaceTopicEventSubscriptionsClient().NewListByNamespaceTopicPager("examplerg", "examplenamespace2", "examplenamespacetopic2", &armeventgrid.NamespaceTopicEventSubscriptionsClientListByNamespaceTopicOptions{Filter: nil,
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
		// page.SubscriptionsListResult = armeventgrid.SubscriptionsListResult{
		// 	Value: []*armeventgrid.Subscription{
		// 		{
		// 			Name: to.Ptr("examplenamespacetopicEventSub1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/namespaces/topics/eventsubscriptions"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/examplenamespace2/topics/examplenamespacetopic2/eventSubscriptions/examplenamespacetopicEventSub1"),
		// 			Properties: &armeventgrid.SubscriptionProperties{
		// 				DeliveryConfiguration: &armeventgrid.DeliveryConfiguration{
		// 					DeliveryMode: to.Ptr(armeventgrid.DeliveryModeQueue),
		// 					Queue: &armeventgrid.QueueInfo{
		// 						EventTimeToLive: to.Ptr("P1D"),
		// 						MaxDeliveryCount: to.Ptr[int32](5),
		// 						ReceiveLockDurationInSeconds: to.Ptr[int32](50),
		// 					},
		// 				},
		// 				EventDeliverySchema: to.Ptr(armeventgrid.DeliverySchemaCloudEventSchemaV10),
		// 				ProvisioningState: to.Ptr(armeventgrid.SubscriptionProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examplenamespacetopicEventSub2"),
		// 			Type: to.Ptr("Microsoft.EventGrid/namespaces/topics/eventsubscriptions"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/examplenamespace2/topics/examplenamespacetopic2/eventSubscriptions/examplenamespacetopicEventSub2"),
		// 			Properties: &armeventgrid.SubscriptionProperties{
		// 				DeliveryConfiguration: &armeventgrid.DeliveryConfiguration{
		// 					DeliveryMode: to.Ptr(armeventgrid.DeliveryModeQueue),
		// 					Queue: &armeventgrid.QueueInfo{
		// 						EventTimeToLive: to.Ptr("P1D"),
		// 						MaxDeliveryCount: to.Ptr[int32](4),
		// 						ReceiveLockDurationInSeconds: to.Ptr[int32](60),
		// 					},
		// 				},
		// 				EventDeliverySchema: to.Ptr(armeventgrid.DeliverySchemaCloudEventSchemaV10),
		// 				ProvisioningState: to.Ptr(armeventgrid.SubscriptionProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
