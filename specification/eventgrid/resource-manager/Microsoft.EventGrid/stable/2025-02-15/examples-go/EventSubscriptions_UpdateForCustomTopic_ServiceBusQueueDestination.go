package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/EventSubscriptions_UpdateForCustomTopic_ServiceBusQueueDestination.json
func ExampleEventSubscriptionsClient_BeginUpdate_eventSubscriptionsUpdateForCustomTopicServiceBusQueueDestination() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEventSubscriptionsClient().BeginUpdate(ctx, "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1", "examplesubscription1", armeventgrid.EventSubscriptionUpdateParameters{
		DeadLetterDestination: &armeventgrid.StorageBlobDeadLetterDestination{
			EndpointType: to.Ptr(armeventgrid.DeadLetterEndPointTypeStorageBlob),
			Properties: &armeventgrid.StorageBlobDeadLetterDestinationProperties{
				BlobContainerName: to.Ptr("contosocontainer"),
				ResourceID:        to.Ptr("/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg"),
			},
		},
		Destination: &armeventgrid.ServiceBusQueueEventSubscriptionDestination{
			EndpointType: to.Ptr(armeventgrid.EndpointTypeServiceBusQueue),
			Properties: &armeventgrid.ServiceBusQueueEventSubscriptionDestinationProperties{
				ResourceID: to.Ptr("/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.ServiceBus/namespaces/ContosoNamespace/queues/SBQ"),
			},
		},
		Filter: &armeventgrid.EventSubscriptionFilter{
			IsSubjectCaseSensitive: to.Ptr(false),
			SubjectBeginsWith:      to.Ptr("ExamplePrefix"),
			SubjectEndsWith:        to.Ptr("ExampleSuffix"),
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
