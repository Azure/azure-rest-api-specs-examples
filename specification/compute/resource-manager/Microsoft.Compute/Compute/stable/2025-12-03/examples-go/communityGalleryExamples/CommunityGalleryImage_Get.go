package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-12-03/communityGalleryExamples/CommunityGalleryImage_Get.json
func ExampleCommunityGalleryImagesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCommunityGalleryImagesClient().Get(ctx, "myLocation", "publicGalleryName", "myGalleryImageName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.CommunityGalleryImagesClientGetResponse{
	// 	CommunityGalleryImage: armcompute.CommunityGalleryImage{
	// 		Properties: &armcompute.CommunityGalleryImageProperties{
	// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
	// 			HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
	// 			Identifier: &armcompute.CommunityGalleryImageIdentifier{
	// 				Publisher: to.Ptr("myPublisherName"),
	// 				Offer: to.Ptr("myOfferName"),
	// 				SKU: to.Ptr("mySkuName"),
	// 			},
	// 			PrivacyStatementURI: to.Ptr("https://test-uri.com"),
	// 			Eula: to.Ptr("https://test-uri.com"),
	// 			ArtifactTags: map[string]*string{
	// 				"ShareTag-CommunityGallery": to.Ptr("CommunityGallery"),
	// 			},
	// 			Disclaimer: to.Ptr("https://test-uri.com"),
	// 		},
	// 		Location: to.Ptr("myLocation"),
	// 		Name: to.Ptr("myGalleryImageName"),
	// 		Type: to.Ptr("Microsoft.Compute/locations/communityGalleryImage"),
	// 		Identifier: &armcompute.CommunityGalleryIdentifier{
	// 			UniqueID: to.Ptr("/CommunityGalleries/publicGalleryName/Images/myGalleryImageName"),
	// 		},
	// 	},
	// }
}
