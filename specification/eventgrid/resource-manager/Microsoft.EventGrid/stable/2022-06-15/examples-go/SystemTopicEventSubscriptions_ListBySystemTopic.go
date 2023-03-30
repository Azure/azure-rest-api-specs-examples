package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f88928d723133dc392e3297e6d61b7f6d10501fd/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/SystemTopicEventSubscriptions_ListBySystemTopic.json
func ExampleSystemTopicEventSubscriptionsClient_NewListBySystemTopicPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSystemTopicEventSubscriptionsClient().NewListBySystemTopicPager("examplerg", "exampleSystemTopic1", &armeventgrid.SystemTopicEventSubscriptionsClientListBySystemTopicOptions{Filter: nil,
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
		// page.EventSubscriptionsListResult = armeventgrid.EventSubscriptionsListResult{
		// 	Value: []*armeventgrid.EventSubscription{
		// 		{
		// 			Name: to.Ptr("examplesubscription1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/systemTopics/eventSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/systemTopics/exampleSystemTopic1/eventSubscriptions/examplesubscription1"),
		// 			Properties: &armeventgrid.EventSubscriptionProperties{
		// 				Destination: &armeventgrid.StorageQueueEventSubscriptionDestination{
		// 					EndpointType: to.Ptr(armeventgrid.EndpointTypeStorageQueue),
		// 					Properties: &armeventgrid.StorageQueueEventSubscriptionDestinationProperties{
		// 						QueueName: to.Ptr("que"),
		// 						ResourceID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.Storage/storageAccounts/testtrackedsource"),
		// 					},
		// 				},
		// 				EventDeliverySchema: to.Ptr(armeventgrid.EventDeliverySchemaEventGridSchema),
		// 				Filter: &armeventgrid.EventSubscriptionFilter{
		// 					SubjectBeginsWith: to.Ptr(""),
		// 					SubjectEndsWith: to.Ptr(""),
		// 				},
		// 				Labels: []*string{
		// 				},
		// 				ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
		// 				RetryPolicy: &armeventgrid.RetryPolicy{
		// 					EventTimeToLiveInMinutes: to.Ptr[int32](1440),
		// 					MaxDeliveryAttempts: to.Ptr[int32](10),
		// 				},
		// 				Topic: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/systemTopics/exampleSystemTopic1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examplesubscription2"),
		// 			Type: to.Ptr("Microsoft.EventGrid/systemTopics/eventSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/systemTopics/exampleSystemTopic1/eventSubscriptions/examplesubscription2"),
		// 			Properties: &armeventgrid.EventSubscriptionProperties{
		// 				Destination: &armeventgrid.StorageQueueEventSubscriptionDestination{
		// 					EndpointType: to.Ptr(armeventgrid.EndpointTypeStorageQueue),
		// 					Properties: &armeventgrid.StorageQueueEventSubscriptionDestinationProperties{
		// 						QueueName: to.Ptr("que"),
		// 						ResourceID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.Storage/storageAccounts/testtrackedsource"),
		// 					},
		// 				},
		// 				EventDeliverySchema: to.Ptr(armeventgrid.EventDeliverySchemaEventGridSchema),
		// 				Filter: &armeventgrid.EventSubscriptionFilter{
		// 					IncludedEventTypes: []*string{
		// 						to.Ptr("Microsoft.Storage.BlobCreated"),
		// 						to.Ptr("Microsoft.Storage.BlobDeleted")},
		// 						IsSubjectCaseSensitive: to.Ptr(false),
		// 						SubjectBeginsWith: to.Ptr("ExamplePrefix"),
		// 						SubjectEndsWith: to.Ptr("ExampleSuffix"),
		// 					},
		// 					Labels: []*string{
		// 						to.Ptr("label1"),
		// 						to.Ptr("label2")},
		// 						ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
		// 						RetryPolicy: &armeventgrid.RetryPolicy{
		// 							EventTimeToLiveInMinutes: to.Ptr[int32](1440),
		// 							MaxDeliveryAttempts: to.Ptr[int32](30),
		// 						},
		// 						Topic: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/systemTopics/exampleSystemTopic1"),
		// 					},
		// 			}},
		// 		}
	}
}
