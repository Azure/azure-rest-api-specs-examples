package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-get-by-name.json
func ExampleStreamingLocatorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewStreamingLocatorsClient().Get(ctx, "contoso", "contosomedia", "clearStreamingLocator", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.StreamingLocator = armmediaservices.StreamingLocator{
	// 	Name: to.Ptr("clearStreamingLocator"),
	// 	Type: to.Ptr("Microsoft.Media/mediaservices/streamingLocators"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosomedia/streamingLocators/clearStreamingLocator"),
	// 	Properties: &armmediaservices.StreamingLocatorProperties{
	// 		AssetName: to.Ptr("ClimbingMountRainier"),
	// 		ContentKeys: []*armmediaservices.StreamingLocatorContentKey{
	// 		},
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:32.115Z"); return t}()),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T23:59:59.999Z"); return t}()),
	// 		StreamingLocatorID: to.Ptr("7684043b-f6d1-44a7-8bed-8a4aa474c5a4"),
	// 		StreamingPolicyName: to.Ptr("clearStreamingPolicy"),
	// 	},
	// }
}
