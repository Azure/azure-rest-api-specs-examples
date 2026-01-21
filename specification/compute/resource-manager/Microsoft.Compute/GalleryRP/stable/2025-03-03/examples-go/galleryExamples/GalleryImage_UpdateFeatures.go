package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/GalleryImage_UpdateFeatures.json
func ExampleGalleryImagesClient_BeginUpdate_updateAGalleryImageFeature() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryImagesClient().BeginUpdate(ctx, "myResourceGroup", "myGalleryName", "myGalleryImageName", armcompute.GalleryImageUpdate{
		Properties: &armcompute.GalleryImageProperties{
			AllowUpdateImage: to.Ptr(true),
			Features: []*armcompute.GalleryImageFeature{
				{
					Name:            to.Ptr("SecurityType"),
					StartsAtVersion: to.Ptr("2.0.0"),
					Value:           to.Ptr("TrustedLaunch"),
				}},
			HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV2),
			Identifier: &armcompute.GalleryImageIdentifier{
				Offer:     to.Ptr("myOfferName"),
				Publisher: to.Ptr("myPublisherName"),
				SKU:       to.Ptr("mySkuName"),
			},
			OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
			OSType:  to.Ptr(armcompute.OperatingSystemTypesWindows),
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
	// res.GalleryImage = armcompute.GalleryImage{
	// 	Name: to.Ptr("myGalleryImageName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryImageProperties{
	// 		Features: []*armcompute.GalleryImageFeature{
	// 			{
	// 				Name: to.Ptr("SecurityType"),
	// 				StartsAtVersion: to.Ptr("2.0.0"),
	// 				Value: to.Ptr("TrustedLaunch"),
	// 		}},
	// 		HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV2),
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
