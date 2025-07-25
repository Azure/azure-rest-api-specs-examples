package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/NamespaceTopicEventSubscriptions_CreateOrUpdate.json
func ExampleNamespaceTopicEventSubscriptionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNamespaceTopicEventSubscriptionsClient().BeginCreateOrUpdate(ctx, "examplerg", "examplenamespace2", "examplenamespacetopic2", "examplenamespacetopicEventSub2", armeventgrid.Subscription{
		Properties: &armeventgrid.SubscriptionProperties{
			DeliveryConfiguration: &armeventgrid.DeliveryConfiguration{
				DeliveryMode: to.Ptr(armeventgrid.DeliveryModeQueue),
				Queue: &armeventgrid.QueueInfo{
					EventTimeToLive:              to.Ptr("P1D"),
					MaxDeliveryCount:             to.Ptr[int32](4),
					ReceiveLockDurationInSeconds: to.Ptr[int32](60),
				},
			},
			EventDeliverySchema: to.Ptr(armeventgrid.DeliverySchemaCloudEventSchemaV10),
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
	// res.Subscription = armeventgrid.Subscription{
	// 	Name: to.Ptr("examplenamespacetopicEventSub2"),
	// 	Type: to.Ptr("Microsoft.EventGrid/namespaces/topics/eventsubscriptions"),
	// 	ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/namespaces/examplenamespace2/topics/examplenamespacetopic2/eventSubscriptions/examplenamespacetopicEventSub2"),
	// 	Properties: &armeventgrid.SubscriptionProperties{
	// 		DeliveryConfiguration: &armeventgrid.DeliveryConfiguration{
	// 			DeliveryMode: to.Ptr(armeventgrid.DeliveryModeQueue),
	// 			Queue: &armeventgrid.QueueInfo{
	// 				EventTimeToLive: to.Ptr("P1D"),
	// 				MaxDeliveryCount: to.Ptr[int32](4),
	// 				ReceiveLockDurationInSeconds: to.Ptr[int32](60),
	// 			},
	// 		},
	// 		EventDeliverySchema: to.Ptr(armeventgrid.DeliverySchemaCloudEventSchemaV10),
	// 		ProvisioningState: to.Ptr(armeventgrid.SubscriptionProvisioningStateSucceeded),
	// 	},
	// }
}
