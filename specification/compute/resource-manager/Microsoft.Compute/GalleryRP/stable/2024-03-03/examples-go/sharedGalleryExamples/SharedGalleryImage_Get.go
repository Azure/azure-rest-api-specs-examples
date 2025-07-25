package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/sharedGalleryExamples/SharedGalleryImage_Get.json
func ExampleSharedGalleryImagesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSharedGalleryImagesClient().Get(ctx, "myLocation", "galleryUniqueName", "myGalleryImageName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SharedGalleryImage = armcompute.SharedGalleryImage{
	// 	Name: to.Ptr("myGalleryImageName"),
	// 	Location: to.Ptr("myLocation"),
	// 	Identifier: &armcompute.SharedGalleryIdentifier{
	// 		UniqueID: to.Ptr("/SharedGalleries/galleryUniqueName/Images/myGalleryImageName"),
	// 	},
	// 	Properties: &armcompute.SharedGalleryImageProperties{
	// 		ArtifactTags: map[string]*string{
	// 			"ShareTag-Official1PGallery": to.Ptr("Official1PGallery"),
	// 		},
	// 		Eula: to.Ptr("https://www.microsoft.com/en-us/"),
	// 		HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
	// 		Identifier: &armcompute.GalleryImageIdentifier{
	// 			Offer: to.Ptr("myOfferName"),
	// 			Publisher: to.Ptr("myPublisherName"),
	// 			SKU: to.Ptr("mySkuName"),
	// 		},
	// 		OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
	// 		OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 	},
	// }
}
