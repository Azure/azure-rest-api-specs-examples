package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-12-03/galleryExamples/GalleryImage_ListByGallery.json
func ExampleGalleryImagesClient_NewListByGalleryPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
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
		// page = armcompute.GalleryImagesClientListByGalleryResponse{
		// 	GalleryImageList: armcompute.GalleryImageList{
		// 		Value: []*armcompute.GalleryImage{
		// 			{
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/locations/westus/Galleries/myGalleryName/Images/myGalleryImageName"),
		// 				Properties: &armcompute.GalleryImageProperties{
		// 					OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 					OSState: to.Ptr(armcompute.OperatingSystemStateTypesGeneralized),
		// 					HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
		// 					Identifier: &armcompute.GalleryImageIdentifier{
		// 						Publisher: to.Ptr("myPublisherName"),
		// 						Offer: to.Ptr("myOfferName"),
		// 						SKU: to.Ptr("mySkuName"),
		// 					},
		// 					ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
		// 				},
		// 				Location: to.Ptr("West US"),
		// 				Name: to.Ptr("myGalleryImageName"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("http://svchost:99/subscriptions/subscriptionId/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/images?$skiptoken=token/Subscriptions/subscriptionId/ResourceGroups/myResourceGroup/galleries/myGalleryName/images/myGalleryImageName"),
		// 	},
		// }
	}
}
