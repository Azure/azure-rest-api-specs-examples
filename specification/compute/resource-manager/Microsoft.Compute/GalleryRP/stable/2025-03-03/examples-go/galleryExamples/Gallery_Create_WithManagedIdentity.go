package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/Gallery_Create_WithManagedIdentity.json
func ExampleGalleriesClient_BeginCreateOrUpdate_createOrUpdateASimpleGalleryWithSystemAssignedAndUserAssignedManagedIdentities() {
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
		Identity: &armcompute.GalleryIdentity{
			Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
				"/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity": {},
			},
		},
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
	// res.Gallery = armcompute.Gallery{
	// 	Name: to.Ptr("myGalleryName"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName"),
	// 	Location: to.Ptr("West US"),
	// 	Identity: &armcompute.GalleryIdentity{
	// 		Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssignedUserAssigned),
	// 		PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		TenantID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
	// 			"/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity": &armcompute.UserAssignedIdentitiesValue{
	// 			},
	// 		},
	// 	},
	// 	Properties: &armcompute.GalleryProperties{
	// 		Description: to.Ptr("This is the gallery description."),
	// 		Identifier: &armcompute.GalleryIdentifier{
	// 			UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
	// 		},
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 	},
	// }
}
