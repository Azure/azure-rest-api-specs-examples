package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/EventSubscriptions_CreateOrUpdateForCustomTopic_StorageQueueDestination.json
func ExampleEventSubscriptionsClient_BeginCreateOrUpdate_eventSubscriptionsCreateOrUpdateForCustomTopicStorageQueueDestination() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeventgrid.NewEventSubscriptionsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1", "examplesubscription1", armeventgrid.EventSubscription{
		Properties: &armeventgrid.EventSubscriptionProperties{
			DeadLetterDestination: &armeventgrid.StorageBlobDeadLetterDestination{
				EndpointType: to.Ptr(armeventgrid.DeadLetterEndPointTypeStorageBlob),
				Properties: &armeventgrid.StorageBlobDeadLetterDestinationProperties{
					BlobContainerName: to.Ptr("contosocontainer"),
					ResourceID:        to.Ptr("/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg"),
				},
			},
			Destination: &armeventgrid.StorageQueueEventSubscriptionDestination{
				EndpointType: to.Ptr(armeventgrid.EndpointTypeStorageQueue),
				Properties: &armeventgrid.StorageQueueEventSubscriptionDestinationProperties{
					QueueMessageTimeToLiveInSeconds: to.Ptr[int64](300),
					QueueName:                       to.Ptr("queue1"),
					ResourceID:                      to.Ptr("/subscriptions/d33c5f7a-02ea-40f4-bf52-07f17e84d6a8/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg"),
				},
			},
			Filter: &armeventgrid.EventSubscriptionFilter{
				IsSubjectCaseSensitive: to.Ptr(false),
				SubjectBeginsWith:      to.Ptr("ExamplePrefix"),
				SubjectEndsWith:        to.Ptr("ExampleSuffix"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
