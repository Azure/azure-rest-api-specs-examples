package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ee1eec42dcc710ff88db2d1bf574b2f9afe3d654/specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/EventSubscriptions_ListByDomainTopic.json
func ExampleEventSubscriptionsClient_NewListByDomainTopicPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEventSubscriptionsClient().NewListByDomainTopicPager("examplerg", "domain1", "topic1", &armeventgrid.EventSubscriptionsClientListByDomainTopicOptions{Filter: nil,
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
		// 			Type: to.Ptr("Microsoft.EventGrid/eventSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/domains/domain1/topics/topic1/providers/Microsoft.EventGrid/eventSubscriptions/examplesubscription1"),
		// 			Properties: &armeventgrid.EventSubscriptionProperties{
		// 				Destination: &armeventgrid.WebHookEventSubscriptionDestination{
		// 					EndpointType: to.Ptr(armeventgrid.EndpointTypeWebHook),
		// 					Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
		// 						EndpointBaseURL: to.Ptr("https://requestb.in/15ksip71"),
		// 					},
		// 				},
		// 				Filter: &armeventgrid.EventSubscriptionFilter{
		// 					IsSubjectCaseSensitive: to.Ptr(false),
		// 					SubjectBeginsWith: to.Ptr(""),
		// 					SubjectEndsWith: to.Ptr(""),
		// 				},
		// 				ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
		// 				Topic: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/microsoft.eventgrid/domains/domain1/topics/topic1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examplesubscription2"),
		// 			Type: to.Ptr("Microsoft.EventGrid/eventSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/domains/domain1/topics/topic1/providers/Microsoft.EventGrid/eventSubscriptions/examplesubscription2"),
		// 			Properties: &armeventgrid.EventSubscriptionProperties{
		// 				Destination: &armeventgrid.WebHookEventSubscriptionDestination{
		// 					EndpointType: to.Ptr(armeventgrid.EndpointTypeWebHook),
		// 					Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
		// 						EndpointBaseURL: to.Ptr("https://requestb.in/15ksip71"),
		// 					},
		// 				},
		// 				Filter: &armeventgrid.EventSubscriptionFilter{
		// 					IsSubjectCaseSensitive: to.Ptr(false),
		// 					SubjectBeginsWith: to.Ptr(""),
		// 					SubjectEndsWith: to.Ptr(""),
		// 				},
		// 				ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
		// 				Topic: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/microsoft.eventgrid/domains/domain1/topics/topic1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("examplesubscription3"),
		// 			Type: to.Ptr("Microsoft.EventGrid/eventSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/domains/domain1/topics/topic1/providers/Microsoft.EventGrid/eventSubscriptions/examplesubscription3"),
		// 			Properties: &armeventgrid.EventSubscriptionProperties{
		// 				Destination: &armeventgrid.WebHookEventSubscriptionDestination{
		// 					EndpointType: to.Ptr(armeventgrid.EndpointTypeWebHook),
		// 					Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
		// 						EndpointBaseURL: to.Ptr("https://requestb.in/15ksip71"),
		// 					},
		// 				},
		// 				Filter: &armeventgrid.EventSubscriptionFilter{
		// 					IsSubjectCaseSensitive: to.Ptr(false),
		// 					SubjectBeginsWith: to.Ptr(""),
		// 					SubjectEndsWith: to.Ptr(""),
		// 				},
		// 				Labels: []*string{
		// 				},
		// 				ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
		// 				Topic: to.Ptr("/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/microsoft.eventgrid/domains/domain1/topics/topic1"),
		// 			},
		// 	}},
		// }
	}
}
