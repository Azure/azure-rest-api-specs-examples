package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/PartnerTopicEventSubscriptions_CreateOrUpdate.json
func ExamplePartnerTopicEventSubscriptionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPartnerTopicEventSubscriptionsClient().BeginCreateOrUpdate(ctx, "examplerg", "examplePartnerTopic1", "exampleEventSubscriptionName1", armeventgrid.EventSubscription{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EventSubscription = armeventgrid.EventSubscription{
	// 	Name: to.Ptr("exampleEventSubscriptionName1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/partnerTopics/eventSubscriptions"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerTopics/examplePartnerTopic1/eventSubscriptions/exampleEventSubscriptionName1"),
	// 	Properties: &armeventgrid.EventSubscriptionProperties{
	// 		Destination: &armeventgrid.WebHookEventSubscriptionDestination{
	// 			EndpointType: to.Ptr(armeventgrid.EndpointTypeWebHook),
	// 			Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
	// 				EndpointBaseURL: to.Ptr("https://requestb.in/15ksip71"),
	// 			},
	// 		},
	// 		EventDeliverySchema: to.Ptr(armeventgrid.EventDeliverySchemaEventGridSchema),
	// 		Filter: &armeventgrid.EventSubscriptionFilter{
	// 			IsSubjectCaseSensitive: to.Ptr(false),
	// 			SubjectBeginsWith: to.Ptr("ExamplePrefix"),
	// 			SubjectEndsWith: to.Ptr("ExampleSuffix"),
	// 		},
	// 		ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
	// 		RetryPolicy: &armeventgrid.RetryPolicy{
	// 			EventTimeToLiveInMinutes: to.Ptr[int32](1440),
	// 			MaxDeliveryAttempts: to.Ptr[int32](30),
	// 		},
	// 		Topic: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/partnerTopics/examplePartnerTopic1"),
	// 	},
	// }
}
