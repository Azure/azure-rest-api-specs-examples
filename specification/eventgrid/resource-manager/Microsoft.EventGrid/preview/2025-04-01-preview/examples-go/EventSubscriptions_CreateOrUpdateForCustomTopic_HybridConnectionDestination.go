package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/EventSubscriptions_CreateOrUpdateForCustomTopic_HybridConnectionDestination.json
func ExampleEventSubscriptionsClient_BeginCreateOrUpdate_eventSubscriptionsCreateOrUpdateForCustomTopicHybridConnectionDestination() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEventSubscriptionsClient().BeginCreateOrUpdate(ctx, "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1", "examplesubscription1", armeventgrid.EventSubscription{
		Properties: &armeventgrid.EventSubscriptionProperties{
			DeadLetterDestination: &armeventgrid.StorageBlobDeadLetterDestination{
				EndpointType: to.Ptr(armeventgrid.DeadLetterEndPointTypeStorageBlob),
				Properties: &armeventgrid.StorageBlobDeadLetterDestinationProperties{
					BlobContainerName: to.Ptr("contosocontainer"),
					ResourceID:        to.Ptr("/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.Storage/storageAccounts/contosostg"),
				},
			},
			Destination: &armeventgrid.HybridConnectionEventSubscriptionDestination{
				EndpointType: to.Ptr(armeventgrid.EndpointTypeHybridConnection),
				Properties: &armeventgrid.HybridConnectionEventSubscriptionDestinationProperties{
					ResourceID: to.Ptr("/subscriptions/d33c5f7a-02ea-40f4-bf52-07f17e84d6a8/resourceGroups/TestRG/providers/Microsoft.Relay/namespaces/ContosoNamespace/hybridConnections/HC1"),
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
