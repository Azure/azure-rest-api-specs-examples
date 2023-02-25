package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Subscriptions/SBSubscriptionCreate.json
func ExampleSubscriptionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewSubscriptionsClient("Subscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx, "ResourceGroup", "sdk-Namespace-1349", "sdk-Topics-8740", "sdk-Subscriptions-2178", armservicebus.SBSubscription{
		Properties: &armservicebus.SBSubscriptionProperties{
			EnableBatchedOperations: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SBSubscription = armservicebus.SBSubscription{
	// 	Name: to.Ptr("sdk-Subscriptions-2178"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics/Subscriptions"),
	// 	ID: to.Ptr("/subscriptions/Subscriptionid/resourceGroups/ResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-1349/topics/sdk-Topics-8740/subscriptions/sdk-Subscriptions-2178"),
	// 	Properties: &armservicebus.SBSubscriptionProperties{
	// 		AccessedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.5992764Z"); return t}()),
	// 		AutoDeleteOnIdle: to.Ptr("P10675199DT2H48M5.4775807S"),
	// 		CountDetails: &armservicebus.MessageCountDetails{
	// 			ActiveMessageCount: to.Ptr[int64](0),
	// 			DeadLetterMessageCount: to.Ptr[int64](0),
	// 			ScheduledMessageCount: to.Ptr[int64](0),
	// 			TransferDeadLetterMessageCount: to.Ptr[int64](0),
	// 			TransferMessageCount: to.Ptr[int64](0),
	// 		},
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.5992764Z"); return t}()),
	// 		DeadLetteringOnFilterEvaluationExceptions: to.Ptr(true),
	// 		DeadLetteringOnMessageExpiration: to.Ptr(true),
	// 		DefaultMessageTimeToLive: to.Ptr("P10675199DT2H48M5.4775807S"),
	// 		EnableBatchedOperations: to.Ptr(true),
	// 		ForwardDeadLetteredMessagesTo: to.Ptr("sdk-Topics-3065"),
	// 		ForwardTo: to.Ptr("sdk-Topics-3065"),
	// 		LockDuration: to.Ptr("PT1M"),
	// 		MaxDeliveryCount: to.Ptr[int32](10),
	// 		MessageCount: to.Ptr[int64](0),
	// 		RequiresSession: to.Ptr(false),
	// 		Status: to.Ptr(armservicebus.EntityStatusActive),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.5992764Z"); return t}()),
	// 	},
	// }
}
