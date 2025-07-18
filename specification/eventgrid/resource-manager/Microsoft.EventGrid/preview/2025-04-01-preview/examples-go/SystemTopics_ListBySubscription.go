package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9549e9fff6f7a4e4232370865bfb6fc771a1f4ac/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/SystemTopics_ListBySubscription.json
func ExampleSystemTopicsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSystemTopicsClient().NewListBySubscriptionPager(&armeventgrid.SystemTopicsClientListBySubscriptionOptions{Filter: nil,
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
		// page.SystemTopicsListResult = armeventgrid.SystemTopicsListResult{
		// 	Value: []*armeventgrid.SystemTopic{
		// 		{
		// 			Name: to.Ptr("exampleSystemTopic2"),
		// 			Type: to.Ptr("Microsoft.EventGrid/systemTopics"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/systemTopics/exampleSystemTopic2"),
		// 			Location: to.Ptr("centraluseuap"),
		// 			Properties: &armeventgrid.SystemTopicProperties{
		// 				MetricResourceID: to.Ptr("183c0fb1-17ff-47b6-ac77-5a47420ab01e"),
		// 				ProvisioningState: to.Ptr(armeventgrid.ResourceProvisioningStateSucceeded),
		// 				Source: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/azureeventgridrunnerrgcentraluseuap/providers/microsoft.storage/storageaccounts/pubstgrunnerb71cd29e"),
		// 				TopicType: to.Ptr("microsoft.storage.storageaccounts"),
		// 			},
		// 	}},
		// }
	}
}
