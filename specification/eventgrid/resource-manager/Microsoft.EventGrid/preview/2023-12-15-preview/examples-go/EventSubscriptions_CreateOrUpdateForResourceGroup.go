package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-12-15-preview/examples/EventSubscriptions_CreateOrUpdateForResourceGroup.json
func ExampleEventSubscriptionsClient_BeginCreateOrUpdate_eventSubscriptionsCreateOrUpdateForResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEventSubscriptionsClient().BeginCreateOrUpdate(ctx, "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg", "examplesubscription2", armeventgrid.EventSubscription{
		Properties: &armeventgrid.EventSubscriptionProperties{
			Destination: &armeventgrid.WebHookEventSubscriptionDestination{
				EndpointType: to.Ptr(armeventgrid.EndpointTypeWebHook),
				Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
					EndpointURL: to.Ptr("https://requestb.in/15ksip71"),
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
