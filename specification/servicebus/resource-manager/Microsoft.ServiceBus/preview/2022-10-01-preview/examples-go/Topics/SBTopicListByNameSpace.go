package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5e24fa8e30bff0cb321494a0b550b1c1282a8a3c/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Topics/SBTopicListByNameSpace.json
func ExampleTopicsClient_NewListByNamespacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armservicebus.NewTopicsClient("5f750a97-50d9-4e36-8081-c9ee4c0210d4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByNamespacePager("Default-ServiceBus-WestUS", "sdk-Namespace-1617", &armservicebus.TopicsClientListByNamespaceOptions{Skip: nil,
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
		// page.SBTopicListResult = armservicebus.SBTopicListResult{
		// 	Value: []*armservicebus.SBTopic{
		// 		{
		// 			Name: to.Ptr("sdk-topics-5488"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-1617/topics/sdk-topics-5488"),
		// 			Properties: &armservicebus.SBTopicProperties{
		// 				AccessedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00Z"); return t}()),
		// 				AutoDeleteOnIdle: to.Ptr("P10675199DT2H48M5.4775807S"),
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-26T20:50:31.4442694Z"); return t}()),
		// 				DefaultMessageTimeToLive: to.Ptr("P10675199DT2H48M5.4775807S"),
		// 				DuplicateDetectionHistoryTimeWindow: to.Ptr("PT10M"),
		// 				EnableBatchedOperations: to.Ptr(true),
		// 				EnableExpress: to.Ptr(true),
		// 				EnablePartitioning: to.Ptr(false),
		// 				MaxMessageSizeInKilobytes: to.Ptr[int64](10240),
		// 				MaxSizeInMegabytes: to.Ptr[int32](10240),
		// 				RequiresDuplicateDetection: to.Ptr(false),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 				Status: to.Ptr(armservicebus.EntityStatusActive),
		// 				SubscriptionCount: to.Ptr[int32](0),
		// 				SupportOrdering: to.Ptr(true),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-26T20:52:32.2092264Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
