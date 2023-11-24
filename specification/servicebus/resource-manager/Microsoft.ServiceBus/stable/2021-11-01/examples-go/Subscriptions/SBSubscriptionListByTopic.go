package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Subscriptions/SBSubscriptionListByTopic.json
func ExampleSubscriptionsClient_NewListByTopicPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubscriptionsClient().NewListByTopicPager("ResourceGroup", "sdk-Namespace-1349", "sdk-Topics-8740", &armservicebus.SubscriptionsClientListByTopicOptions{Skip: nil,
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
		// page.SBSubscriptionListResult = armservicebus.SBSubscriptionListResult{
		// 	Value: []*armservicebus.SBSubscription{
		// 		{
		// 			Name: to.Ptr("sdk-Subscriptions-2178"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics/Subscriptions"),
		// 			ID: to.Ptr("/subscriptions/Subscriptionid/resourceGroups/ResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-1349/topics/sdk-Topics-8740/subscriptions/sdk-Subscriptions-2178"),
		// 			Properties: &armservicebus.SBSubscriptionProperties{
		// 				AccessedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.599Z"); return t}()),
		// 				AutoDeleteOnIdle: to.Ptr("P10675199DT2H48M5.4775807S"),
		// 				CountDetails: &armservicebus.MessageCountDetails{
		// 					ActiveMessageCount: to.Ptr[int64](0),
		// 					DeadLetterMessageCount: to.Ptr[int64](0),
		// 					ScheduledMessageCount: to.Ptr[int64](0),
		// 					TransferDeadLetterMessageCount: to.Ptr[int64](0),
		// 					TransferMessageCount: to.Ptr[int64](0),
		// 				},
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.599Z"); return t}()),
		// 				DeadLetteringOnFilterEvaluationExceptions: to.Ptr(true),
		// 				DeadLetteringOnMessageExpiration: to.Ptr(true),
		// 				DefaultMessageTimeToLive: to.Ptr("P10675199DT2H48M5.4775807S"),
		// 				EnableBatchedOperations: to.Ptr(true),
		// 				ForwardDeadLetteredMessagesTo: to.Ptr("sdk-Topics-3065"),
		// 				ForwardTo: to.Ptr("sdk-Topics-3065"),
		// 				LockDuration: to.Ptr("PT1M"),
		// 				MaxDeliveryCount: to.Ptr[int32](10),
		// 				MessageCount: to.Ptr[int64](0),
		// 				RequiresSession: to.Ptr(false),
		// 				Status: to.Ptr(armservicebus.EntityStatusActive),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.599Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
