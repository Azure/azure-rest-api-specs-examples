package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/EventSubscriptions_GetForResource.json
func ExampleEventSubscriptionsClient_Get_eventSubscriptionsGetForResource() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEventSubscriptionsClient().Get(ctx, "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventHub/namespaces/examplenamespace1", "examplesubscription1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.EventSubscription = armeventgrid.EventSubscription{
	// 	Name: to.Ptr("examplesubscription1"),
	// 	Type: to.Ptr("Microsoft.EventGrid/eventSubscriptions"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventHub/namespaces/examplenamespace1/providers/Microsoft.EventGrid/eventSubscriptions/examplesubscription1"),
	// 	Properties: &armeventgrid.EventSubscriptionProperties{
	// 		Destination: &armeventgrid.EventHubEventSubscriptionDestination{
	// 			EndpointType: to.Ptr(armeventgrid.EndpointTypeEventHub),
	// 			Properties: &armeventgrid.EventHubEventSubscriptionDestinationProperties{
	// 				ResourceID: to.Ptr("/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.EventHub/namespaces/ContosoNamespace/eventhubs/EH1"),
	// 			},
	// 		},
	// 		Filter: &armeventgrid.EventSubscriptionFilter{
	// 			IsSubjectCaseSensitive: to.Ptr(false),
	// 			SubjectBeginsWith: to.Ptr("ExamplePrefix"),
	// 			SubjectEndsWith: to.Ptr("ExampleSuffix"),
	// 		},
	// 		Labels: []*string{
	// 			to.Ptr("label1"),
	// 			to.Ptr("label2")},
	// 			ProvisioningState: to.Ptr(armeventgrid.EventSubscriptionProvisioningStateSucceeded),
	// 			Topic: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventHub/namespaces/examplenamespace1"),
	// 		},
	// 	}
}
