package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/assets-list-streaming-locators.json
func ExampleAssetsClient_ListStreamingLocators() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssetsClient().ListStreamingLocators(ctx, "contoso", "contosomedia", "ClimbingMountSaintHelens", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ListStreamingLocatorsResponse = armmediaservices.ListStreamingLocatorsResponse{
	// 	StreamingLocators: []*armmediaservices.AssetStreamingLocator{
	// 		{
	// 			Name: to.Ptr("secureStreamingLocator"),
	// 			AssetName: to.Ptr("ClimbingMountSaintHelens"),
	// 			Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:26.972Z"); return t}()),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T23:59:59.999Z"); return t}()),
	// 			StreamingLocatorID: to.Ptr("36b74ce3-20b4-4de0-84f1-97e9138e886c"),
	// 			StreamingPolicyName: to.Ptr("secureStreamingPolicy"),
	// 		},
	// 		{
	// 			Name: to.Ptr("clearStreamingLocator"),
	// 			AssetName: to.Ptr("ClimbingMountSaintHelens"),
	// 			Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-08T18:29:26.948Z"); return t}()),
	// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "9999-12-31T23:59:59.999Z"); return t}()),
	// 			StreamingLocatorID: to.Ptr("3e8d9ac3-50f6-4f6d-8482-078ceb56f23a"),
	// 			StreamingPolicyName: to.Ptr("clearStreamingPolicy"),
	// 	}},
	// }
}
