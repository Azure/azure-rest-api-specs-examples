package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/Gallery_ListByResourceGroup.json
func ExampleGalleriesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGalleriesClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page.GalleryList = armcompute.GalleryList{
		// 	Value: []*armcompute.Gallery{
		// 		{
		// 			Name: to.Ptr("myGalleryName"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcompute.GalleryProperties{
		// 				Description: to.Ptr("This is the gallery description."),
		// 				Identifier: &armcompute.GalleryIdentifier{
		// 					UniqueName: to.Ptr("{subscription-id}-MYGALLERYNAME"),
		// 				},
		// 				ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
