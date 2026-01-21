package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/CommunityGallery_Create.json
func ExampleGalleriesClient_BeginCreateOrUpdate_createACommunityGallery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleriesClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", armcompute.Gallery{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryProperties{
			Description: to.Ptr("This is the gallery description."),
			SharingProfile: &armcompute.SharingProfile{
				CommunityGalleryInfo: &armcompute.CommunityGalleryInfo{
					Eula:             to.Ptr("eula"),
					PublicNamePrefix: to.Ptr("PirPublic"),
					PublisherContact: to.Ptr("pir@microsoft.com"),
					PublisherURI:     to.Ptr("uri"),
				},
				Permissions: to.Ptr(armcompute.GallerySharingPermissionTypesCommunity),
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
	// res.Gallery = armcompute.Gallery{
	// 	Name: to.Ptr("myGalleryName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryProperties{
	// 		Description: to.Ptr("This is the gallery description."),
	// 		Identifier: &armcompute.GalleryIdentifier{
	// 			UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		SharingProfile: &armcompute.SharingProfile{
	// 			CommunityGalleryInfo: &armcompute.CommunityGalleryInfo{
	// 				Eula: to.Ptr("eula"),
	// 				PublicNamePrefix: to.Ptr("PirPublic"),
	// 				PublisherContact: to.Ptr("pir@microsoft.com"),
	// 				PublisherURI: to.Ptr("uri"),
	// 			},
	// 			Permissions: to.Ptr(armcompute.GallerySharingPermissionTypesCommunity),
	// 		},
	// 	},
	// }
}
