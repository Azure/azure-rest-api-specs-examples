package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/173bb3b6fd5b1809fdbf347f67fccfa0440ac126/specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Topics_ListByResourceGroup.json
func ExampleTopicsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopicsClient().NewListByResourceGroupPager("examplerg", &armeventgrid.TopicsClientListByResourceGroupOptions{Filter: nil,
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
		// page.TopicsListResult = armeventgrid.TopicsListResult{
		// 	Value: []*armeventgrid.Topic{
		// 		{
		// 			Name: to.Ptr("exampletopic1"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topics"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic1"),
		// 			Location: to.Ptr("westus2"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("value1"),
		// 				"tag2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armeventgrid.TopicProperties{
		// 				Endpoint: to.Ptr("https://exampletopic1.westus2-1.eventgrid.azure.net/api/events"),
		// 				ProvisioningState: to.Ptr(armeventgrid.TopicProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("exampletopic2"),
		// 			Type: to.Ptr("Microsoft.EventGrid/topics"),
		// 			ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2"),
		// 			Location: to.Ptr("westcentralus"),
		// 			Tags: map[string]*string{
		// 				"tag1": to.Ptr("value1"),
		// 				"tag2": to.Ptr("value2"),
		// 			},
		// 			Properties: &armeventgrid.TopicProperties{
		// 				Endpoint: to.Ptr("https://exampletopic2.westcentralus-1.eventgrid.azure.net/api/events"),
		// 				ProvisioningState: to.Ptr(armeventgrid.TopicProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
