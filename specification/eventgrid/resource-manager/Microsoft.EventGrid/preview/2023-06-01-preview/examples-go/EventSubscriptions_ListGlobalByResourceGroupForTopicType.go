package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/EventSubscriptions_ListGlobalByResourceGroupForTopicType.json
func ExampleEventSubscriptionsClient_NewListGlobalByResourceGroupForTopicTypePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEventSubscriptionsClient().NewListGlobalByResourceGroupForTopicTypePager("examplerg", "Microsoft.Resources.ResourceGroups", &armeventgrid.EventSubscriptionsClientListGlobalByResourceGroupForTopicTypeOptions{Filter: nil,
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
		// 			Name: to.Ptr("examplesubscription3"),
		// 			Type: to.Ptr("Microsoft.EventGrid/eventSubscriptions"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/eventSubscriptions/examplesubscription3"),
		// 			Properties: &armeventgrid.EventSubscriptionProperties{
		// 				Destination: &armeventgrid.WebHookEventSubscriptionDestination{
		// 					EndpointType: to.Ptr(armeventgrid.EndpointTypeWebHook),
		// 					Properties: &armeventgrid.WebHookEventSubscriptionDestinationProperties{
		// 						EndpointBaseURL: to.Ptr("https://requestb.in/15ksip71"),
		// 					},
		// 				},
		// 				Filter: &armeventgrid.EventSubscriptionFilter{
		// 					IsSubjectCaseSensitive: to.Ptr(false),
		// 					SubjectBeginsWith: to.Ptr("ExamplePrefix"),
		// 					SubjectEndsWith: to.Ptr("ExampleSuffix"),
		// 				},
		// 				Labels: []*string{
		// 					to.Ptr("Finance"),
		// 					to.Ptr("HR")},
		// 					ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
		// 					Topic: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg"),
		// 				},
		// 		}},
		// 	}
	}
}
