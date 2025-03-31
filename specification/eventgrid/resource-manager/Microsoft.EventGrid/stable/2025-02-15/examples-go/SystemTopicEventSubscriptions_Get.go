package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/SystemTopicEventSubscriptions_Get.json
func ExampleSystemTopicEventSubscriptionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSystemTopicEventSubscriptionsClient().Get(ctx, "examplerg", "exampleSystemTopic1", "examplesubscription1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EventSubscription = armeventgrid.EventSubscription{
	// 	Name: to.Ptr("examplesubscription1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/systemTopics/eventSubscriptions"),
	// 	ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/systemTopics/exampleSystemTopic1/eventSubscriptions/examplesubscription1"),
	// 	Properties: &armeventgrid.EventSubscriptionProperties{
	// 		Destination: &armeventgrid.StorageQueueEventSubscriptionDestination{
	// 			EndpointType: to.Ptr(armeventgrid.EndpointTypeStorageQueue),
	// 			Properties: &armeventgrid.StorageQueueEventSubscriptionDestinationProperties{
	// 				QueueName: to.Ptr("que"),
	// 				ResourceID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.Storage/storageAccounts/testtrackedsource"),
	// 			},
	// 		},
	// 		EventDeliverySchema: to.Ptr(armeventgrid.EventDeliverySchemaEventGridSchema),
	// 		Filter: &armeventgrid.EventSubscriptionFilter{
	// 			IncludedEventTypes: []*string{
	// 				to.Ptr("Microsoft.Storage.BlobCreated"),
	// 				to.Ptr("Microsoft.Storage.BlobDeleted")},
	// 				IsSubjectCaseSensitive: to.Ptr(false),
	// 				SubjectBeginsWith: to.Ptr("ExamplePrefix"),
	// 				SubjectEndsWith: to.Ptr("ExampleSuffix"),
	// 			},
	// 			Labels: []*string{
	// 				to.Ptr("label1"),
	// 				to.Ptr("label2")},
	// 				ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
	// 				RetryPolicy: &armeventgrid.RetryPolicy{
	// 					EventTimeToLiveInMinutes: to.Ptr[int32](1440),
	// 					MaxDeliveryAttempts: to.Ptr[int32](30),
	// 				},
	// 				Topic: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/systemTopics/exampleSystemTopic1"),
	// 			},
	// 		}
}
