package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/galleryExamples/Gallery_Get.json
func ExampleGalleriesClient_Get_getAGallery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleriesClient().Get(ctx, "myResourceGroup", "myGalleryName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.GalleriesClientGetResponse{
	// 	Gallery: armcompute.Gallery{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName"),
	// 		Properties: &armcompute.GalleryProperties{
	// 			Description: to.Ptr("This is the gallery description."),
	// 			Identifier: &armcompute.GalleryIdentifier{
	// 				UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Name: to.Ptr("myGalleryName"),
	// 	},
	// }
}
