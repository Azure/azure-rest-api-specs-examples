package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/CommunityGallery_Get.json
func ExampleGalleriesClient_Get_getACommunityGallery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleriesClient().Get(ctx, "myResourceGroup", "myGalleryName", &armcompute.GalleriesClientGetOptions{Select: nil,
		Expand: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Gallery = armcompute.Gallery{
	// 	Name: to.Ptr("myGalleryName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/communityGalleries/myGalleryName"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryProperties{
	// 		Description: to.Ptr("This is the gallery description."),
	// 		Identifier: &armcompute.GalleryIdentifier{
	// 			UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		SharingProfile: &armcompute.SharingProfile{
	// 			CommunityGalleryInfo: &armcompute.CommunityGalleryInfo{
	// 				CommunityGalleryEnabled: to.Ptr(true),
	// 				Eula: to.Ptr("eula"),
	// 				PublicNames: []*string{
	// 					to.Ptr("GalleryPublicName")},
	// 					PublisherContact: to.Ptr("pir@microsoft.com"),
	// 					PublisherURI: to.Ptr("uri"),
	// 				},
	// 				Permissions: to.Ptr(armcompute.GallerySharingPermissionTypesCommunity),
	// 			},
	// 			SharingStatus: &armcompute.SharingStatus{
	// 				AggregatedState: to.Ptr(armcompute.SharingStateSucceeded),
	// 				Summary: []*armcompute.RegionalSharingStatus{
	// 					{
	// 						Region: to.Ptr("westus"),
	// 						State: to.Ptr(armcompute.SharingStateSucceeded),
	// 						Details: to.Ptr(""),
	// 				}},
	// 			},
	// 		},
	// 	}
}
