package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/EventSubscriptions_UpdateForCustomTopic_HybridConnectionDestination.json
func ExampleEventSubscriptionsClient_BeginUpdate_eventSubscriptionsUpdateForCustomTopicHybridConnectionDestination() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEventSubscriptionsClient().BeginUpdate(ctx, "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2", "examplesubscription1", armeventgrid.EventSubscriptionUpdateParameters{
		Destination: &armeventgrid.HybridConnectionEventSubscriptionDestination{
			EndpointType: to.Ptr(armeventgrid.EndpointTypeHybridConnection),
			Properties: &armeventgrid.HybridConnectionEventSubscriptionDestinationProperties{
				ResourceID: to.Ptr("/subscriptions/d33c5f7a-02ea-40f4-bf52-07f17e84d6a8/resourceGroups/TestRG/providers/Microsoft.Relay/namespaces/ContosoNamespace/hybridConnections/HC1"),
			},
		},
		Filter: &armeventgrid.EventSubscriptionFilter{
			IsSubjectCaseSensitive: to.Ptr(true),
			SubjectBeginsWith:      to.Ptr("existingPrefix"),
			SubjectEndsWith:        to.Ptr("newSuffix"),
		},
		Labels: []*string{
			to.Ptr("label1"),
			to.Ptr("label2")},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
