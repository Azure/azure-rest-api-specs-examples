package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/Topics/SBTopicGet.json
func ExampleTopicsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTopicsClient().Get(ctx, "ArunMonocle", "sdk-Namespace-1617", "sdk-Topics-5488", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SBTopic = armservicebus.SBTopic{
	// 	Name: to.Ptr("sdk-Topics-5488"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-1617/topics/sdk-Topics-5488"),
	// 	Properties: &armservicebus.SBTopicProperties{
	// 		AccessedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "1-01-01T00:00:00.000Z"); return t}()),
	// 		AutoDeleteOnIdle: to.Ptr("P10675199DT2H48M5.4775807S"),
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-26T20:50:31.444Z"); return t}()),
	// 		DefaultMessageTimeToLive: to.Ptr("P10675199DT2H48M5.4775807S"),
	// 		DuplicateDetectionHistoryTimeWindow: to.Ptr("PT10M"),
	// 		EnableBatchedOperations: to.Ptr(true),
	// 		EnableExpress: to.Ptr(true),
	// 		EnablePartitioning: to.Ptr(false),
	// 		MaxMessageSizeInKilobytes: to.Ptr[int64](10240),
	// 		MaxSizeInMegabytes: to.Ptr[int32](10240),
	// 		RequiresDuplicateDetection: to.Ptr(false),
	// 		SizeInBytes: to.Ptr[int64](0),
	// 		Status: to.Ptr(armservicebus.EntityStatusActive),
	// 		SubscriptionCount: to.Ptr[int32](0),
	// 		SupportOrdering: to.Ptr(true),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-26T20:52:32.209Z"); return t}()),
	// 	},
	// }
}
