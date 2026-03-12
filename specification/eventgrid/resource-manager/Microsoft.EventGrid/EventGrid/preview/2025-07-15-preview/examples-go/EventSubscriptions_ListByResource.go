package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: 2025-07-15-preview/EventSubscriptions_ListByResource.json
func ExampleEventSubscriptionsClient_NewListByResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("8f6b6269-84f2-4d09-9e31-1127efcd1e40", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEventSubscriptionsClient().NewListByResourcePager("examplerg", "Microsoft.EventGrid", "topics", "exampletopic2", nil)
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
		// page = armeventgrid.EventSubscriptionsClientListByResourceResponse{
		// 	EventSubscriptionsListResult: armeventgrid.EventSubscriptionsListResult{
		// 		Value: []*armeventgrid.EventSubscription{
		// 			{
		// 				Properties: &armeventgrid.EventSubscriptionProperties{
		// 					Destination: &armeventgrid.WebHookEventSubscriptionDestination{
		// 						Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
		// 							EndpointBaseURL: to.Ptr("https://requestb.in/15ksip71"),
		// 						},
		// 						EndpointType: to.Ptr(armeventgrid.EndpointTypeWebHook),
		// 					},
		// 					Filter: &armeventgrid.EventSubscriptionFilter{
		// 						IsSubjectCaseSensitive: to.Ptr(false),
		// 						SubjectBeginsWith: to.Ptr(""),
		// 						SubjectEndsWith: to.Ptr(""),
		// 					},
		// 					ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
		// 					Topic: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/microsoft.eventgrid/topics/exampletopic2"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2/providers/Microsoft.EventGrid/eventSubscriptions/examplesubscription1"),
		// 				Name: to.Ptr("examplesubscription1"),
		// 				Type: to.Ptr("Microsoft.EventGrid/eventSubscriptions"),
		// 			},
		// 			{
		// 				Properties: &armeventgrid.EventSubscriptionProperties{
		// 					Destination: &armeventgrid.WebHookEventSubscriptionDestination{
		// 						Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
		// 							EndpointBaseURL: to.Ptr("https://requestb.in/15ksip71"),
		// 						},
		// 						EndpointType: to.Ptr(armeventgrid.EndpointTypeWebHook),
		// 					},
		// 					Filter: &armeventgrid.EventSubscriptionFilter{
		// 						IsSubjectCaseSensitive: to.Ptr(false),
		// 						SubjectBeginsWith: to.Ptr(""),
		// 						SubjectEndsWith: to.Ptr(""),
		// 					},
		// 					ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
		// 					Topic: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/microsoft.eventgrid/topics/exampletopic2"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2/providers/Microsoft.EventGrid/eventSubscriptions/examplesubscription2"),
		// 				Name: to.Ptr("examplesubscription2"),
		// 				Type: to.Ptr("Microsoft.EventGrid/eventSubscriptions"),
		// 			},
		// 			{
		// 				Properties: &armeventgrid.EventSubscriptionProperties{
		// 					Destination: &armeventgrid.WebHookEventSubscriptionDestination{
		// 						Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
		// 							EndpointBaseURL: to.Ptr("https://requestb.in/15ksip71"),
		// 						},
		// 						EndpointType: to.Ptr(armeventgrid.EndpointTypeWebHook),
		// 					},
		// 					Filter: &armeventgrid.EventSubscriptionFilter{
		// 						IsSubjectCaseSensitive: to.Ptr(false),
		// 						SubjectBeginsWith: to.Ptr(""),
		// 						SubjectEndsWith: to.Ptr(""),
		// 					},
		// 					Labels: []*string{
		// 					},
		// 					ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
		// 					Topic: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/microsoft.eventgrid/topics/exampletopic2"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2/providers/Microsoft.EventGrid/eventSubscriptions/examplesubscription3"),
		// 				Name: to.Ptr("examplesubscription3"),
		// 				Type: to.Ptr("Microsoft.EventGrid/eventSubscriptions"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
