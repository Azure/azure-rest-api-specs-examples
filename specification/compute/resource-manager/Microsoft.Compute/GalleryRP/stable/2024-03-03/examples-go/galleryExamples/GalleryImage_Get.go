package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d6d0798c6f5eb196fba7bd1924db2b145a94f58c/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/GalleryImage_Get.json
func ExampleGalleryImagesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.GalleryImage = armcompute.GalleryImage{
	// 	Name: to.Ptr("myGalleryImageName"),
	// 	ID: to.Ptr("/providers/Microsoft.Compute/galleries/myGallery/Images/myGalleryImageName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryImageProperties{
	// 		HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
	// 		Identifier: &armcompute.GalleryImageIdentifier{
	// 			Offer: to.Ptr("myOfferName"),
	// 			Publisher: to.Ptr("myPublisherName"),
	// 			SKU: to.Ptr("mySkuName"),
	// 		},
	// 		OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
	// 		OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 	},
	// }
}
