package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-list.json
func ExampleStreamingLocatorsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStreamingLocatorsClient().NewListPager("contoso", "contosomedia", &armmediaservices.StreamingLocatorsClientListOptions{Filter: nil,
		Top:     nil,
		Orderby: nil,
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
		// page.StreamingLocatorCollection = armmediaservices.StreamingLocatorCollection{
		// 	Value: []*armmediaservices.StreamingLocator{
		// 		{
		// 			Name: to.Ptr("clearStreamingLocator"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/streamingLocators"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/streamingLocators/clearStreamingLocator"),
		// 			Properties: &armmediaservices.StreamingLocatorProperties{
		// 				AssetName: to.Ptr("ClimbingMountRainier"),
		// 				ContentKeys: []*armmediaservices.StreamingLocatorContentKey{
		// 				},
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:31.934Z"); return t}()),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T23:59:59.999Z"); return t}()),
		// 				StreamingLocatorID: to.Ptr("6a116ec6-0c85-441f-9c31-89a5bc3adf0a"),
		// 				StreamingPolicyName: to.Ptr("clearStreamingPolicy"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("secureStreamingLocator"),
		// 			Type: to.Ptr("Microsoft.Media/mediaservices/streamingLocators"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/streamingLocators/secureStreamingLocator"),
		// 			Properties: &armmediaservices.StreamingLocatorProperties{
		// 				AssetName: to.Ptr("ClimbingMountRainier"),
		// 				ContentKeys: []*armmediaservices.StreamingLocatorContentKey{
		// 				},
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:31.954Z"); return t}()),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T23:59:59.999Z"); return t}()),
		// 				StreamingLocatorID: to.Ptr("7338ef90-ffc8-42de-8bff-de8f99973300"),
		// 				StreamingPolicyName: to.Ptr("secureStreamingPolicy"),
		// 			},
		// 	}},
		// }
	}
}
