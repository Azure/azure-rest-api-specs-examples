package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/EventSubscriptions_UpdateForResource.json
func ExampleEventSubscriptionsClient_BeginUpdate_eventSubscriptionsUpdateForResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEventSubscriptionsClient().BeginUpdate(ctx, "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventHub/namespaces/examplenamespace1", "examplesubscription1", armeventgrid.EventSubscriptionUpdateParameters{
		Destination: &armeventgrid.WebHookEventSubscriptionDestination{
			EndpointType: to.Ptr(armeventgrid.EndpointTypeWebHook),
			Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
				EndpointURL: to.Ptr("https://requestb.in/15ksip71"),
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
