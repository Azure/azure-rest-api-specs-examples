package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/galleryExamples/GalleryImage_Create.json
func ExampleGalleryImagesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryImagesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", "myGalleryImageName", armcompute.GalleryImage{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryImageProperties{
			OSType:           to.Ptr(armcompute.OperatingSystemTypesWindows),
			OSState:          to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
			HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
			Identifier: &armcompute.GalleryImageIdentifier{
				Publisher: to.Ptr("myPublisherName"),
				Offer:     to.Ptr("myOfferName"),
				SKU:       to.Ptr("mySkuName"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.GalleryImagesClientCreateOrUpdateResponse{
	// 	GalleryImage: &armcompute.GalleryImage{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGallery/Images/myGalleryImageName"),
	// 		Properties: &armcompute.GalleryImageProperties{
	// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
	// 			HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
	// 			Identifier: &armcompute.GalleryImageIdentifier{
	// 				Publisher: to.Ptr("myPublisherName"),
	// 				Offer: to.Ptr("myOfferName"),
	// 				SKU: to.Ptr("mySkuName"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateUpdating),
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Name: to.Ptr("myGalleryImageName"),
	// 	},
	// }
}
