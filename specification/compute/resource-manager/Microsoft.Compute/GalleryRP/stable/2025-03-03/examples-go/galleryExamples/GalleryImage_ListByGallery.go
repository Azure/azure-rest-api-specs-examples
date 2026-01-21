package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/GalleryImage_ListByGallery.json
func ExampleGalleryImagesClient_NewListByGalleryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGalleryImagesClient().NewListByGalleryPager("myResourceGroup", "myGalleryName", nil)
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
		// page.GalleryImageList = armcompute.GalleryImageList{
		// 	Value: []*armcompute.GalleryImage{
		// 		{
		// 			Name: to.Ptr("myGalleryImageName"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/locations/westus/Galleries/myGalleryName/Images/myGalleryImageName"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcompute.GalleryImageProperties{
		// 				HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
		// 				Identifier: &armcompute.GalleryImageIdentifier{
		// 					Offer: to.Ptr("myOfferName"),
		// 					Publisher: to.Ptr("myPublisherName"),
		// 					SKU: to.Ptr("mySkuName"),
		// 				},
		// 				OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 				ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
