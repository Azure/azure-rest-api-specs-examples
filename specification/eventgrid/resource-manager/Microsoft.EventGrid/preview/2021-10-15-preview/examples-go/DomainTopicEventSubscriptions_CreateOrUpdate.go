package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2021-10-15-preview/examples/DomainTopicEventSubscriptions_CreateOrUpdate.json
func ExampleDomainTopicEventSubscriptionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeventgrid.NewDomainTopicEventSubscriptionsClient("5b4b650e-28b9-4790-b3ab-ddbd88d727c4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "examplerg", "exampleDomain1", "exampleDomainTopic1", "exampleEventSubscriptionName1", armeventgrid.EventSubscription{
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
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
