package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/galleryExamples/GalleryImage_Get.json
func ExampleGalleryImagesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleryImagesClient().Get(ctx, "myResourceGroup", "myGalleryName", "myGalleryImageName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.GalleryImagesClientGetResponse{
	// 	GalleryImage: armcompute.GalleryImage{
	// 		ID: to.Ptr("/providers/Microsoft.Compute/galleries/myGallery/Images/myGalleryImageName"),
	// 		Properties: &armcompute.GalleryImageProperties{
	// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
	// 			HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
	// 			Identifier: &armcompute.GalleryImageIdentifier{
	// 				Publisher: to.Ptr("myPublisherName"),
	// 				Offer: to.Ptr("myOfferName"),
	// 				SKU: to.Ptr("mySkuName"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Name: to.Ptr("myGalleryImageName"),
	// 	},
	// }
}
