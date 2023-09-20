package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/galleryExamples/Gallery_Get_WithSelectPermissions.json
func ExampleGalleriesClient_Get_getAGalleryWithSelectPermissions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleriesClient().Get(ctx, "myResourceGroup", "myGalleryName", &armcompute.GalleriesClientGetOptions{Select: to.Ptr(armcompute.SelectPermissionsPermissions),
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
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryProperties{
	// 		SharingProfile: &armcompute.SharingProfile{
	// 			Groups: []*armcompute.SharingProfileGroup{
	// 				{
	// 					Type: to.Ptr(armcompute.SharingProfileGroupTypesSubscriptions),
	// 					IDs: []*string{
	// 						to.Ptr("34a4ab42-0d72-47d9-bd1a-aed207386dac"),
	// 						to.Ptr("380fd389-260b-41aa-bad9-0a83108c370b")},
	// 					},
	// 					{
	// 						Type: to.Ptr(armcompute.SharingProfileGroupTypesAADTenants),
	// 						IDs: []*string{
	// 							to.Ptr("c24c76aa-8897-4027-9b03-8f7928b54ff6")},
	// 					}},
	// 					Permissions: to.Ptr(armcompute.GallerySharingPermissionTypesGroups),
	// 				},
	// 			},
	// 		}
}
