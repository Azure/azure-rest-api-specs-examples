package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/EventSubscriptions_GetForCustomTopic_StorageQueueDestination.json
func ExampleEventSubscriptionsClient_Get_eventSubscriptionsGetForCustomTopicStorageQueueDestination() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEventSubscriptionsClient().Get(ctx, "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2", "examplesubscription1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EventSubscription = armeventgrid.EventSubscription{
	// 	Name: to.Ptr("examplesubscription1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/eventSubscriptions"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2/providers/Microsoft.EventGrid/eventSubscriptions/examplesubscription1"),
	// 	Properties: &armeventgrid.EventSubscriptionProperties{
	// 		Destination: &armeventgrid.StorageQueueEventSubscriptionDestination{
	// 			EndpointType: to.Ptr(armeventgrid.EndpointTypeStorageQueue),
	// 			Properties: &armeventgrid.StorageQueueEventSubscriptionDestinationProperties{
	// 				QueueMessageTimeToLiveInSeconds: to.Ptr[int64](300),
	// 				QueueName: to.Ptr("queue1"),
	// 				ResourceID: to.Ptr("/subscriptions/d33c5f7a-02ea-40f4-bf52-07f17e84d6a8/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg"),
	// 			},
	// 		},
	// 		Filter: &armeventgrid.EventSubscriptionFilter{
	// 			IsSubjectCaseSensitive: to.Ptr(false),
	// 			SubjectBeginsWith: to.Ptr("ExamplePrefix"),
	// 			SubjectEndsWith: to.Ptr("ExampleSuffix"),
	// 		},
	// 		Labels: []*string{
	// 			to.Ptr("label1"),
	// 			to.Ptr("label2")},
	// 			ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
	// 			Topic: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/microsoft.eventgrid/topics/exampletopic2"),
	// 		},
	// 	}
}
