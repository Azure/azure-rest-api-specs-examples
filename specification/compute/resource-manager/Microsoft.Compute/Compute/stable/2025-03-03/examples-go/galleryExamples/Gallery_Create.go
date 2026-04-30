package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/galleryExamples/Gallery_Create.json
func ExampleGalleriesClient_BeginCreateOrUpdate_createOrUpdateASimpleGallery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleriesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", armcompute.Gallery{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryProperties{
			Description: to.Ptr("This is the gallery description."),
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
	// res = armcompute.GalleriesClientCreateOrUpdateResponse{
	// 	Gallery: &armcompute.Gallery{
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGallery"),
	// 		Properties: &armcompute.GalleryProperties{
	// 			Description: to.Ptr("This is the gallery description."),
	// 			Identifier: &armcompute.GalleryIdentifier{
	// 				UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
	// 			},
	// 			ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateUpdating),
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Name: to.Ptr("myGalleryName"),
	// 	},
	// }
}
