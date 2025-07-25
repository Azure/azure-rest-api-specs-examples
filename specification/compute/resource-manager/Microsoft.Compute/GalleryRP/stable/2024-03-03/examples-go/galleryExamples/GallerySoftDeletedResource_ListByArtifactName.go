package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4517f89a8ebd2f6a94e107e5ee60fff9886f3612/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/GallerySoftDeletedResource_ListByArtifactName.json
func ExampleSoftDeletedResourceClient_NewListByArtifactNamePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSoftDeletedResourceClient().NewListByArtifactNamePager("myResourceGroup", "myGalleryName", "images", "myGalleryImageName", nil)
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
		// page.GallerySoftDeletedResourceList = armcompute.GallerySoftDeletedResourceList{
		// 	Value: []*armcompute.GallerySoftDeletedResource{
		// 		{
		// 			Name: to.Ptr("1.0.0"),
		// 			Type: to.Ptr("Microsoft.Compute/galleries/softDeletedArtifactTypes/artifacts/versions"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/softDeletedArtifactTypes/Images/artifacts/myGalleryImageName/versions/1.0.0"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcompute.GallerySoftDeletedResourceProperties{
		// 				ResourceArmID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/images/myGalleryImageName/versions/1.0.0"),
		// 				SoftDeletedArtifactType: to.Ptr(armcompute.SoftDeletedArtifactTypesImages),
		// 				SoftDeletedTime: to.Ptr("2024-10-17T13:01:05+00:00"),
		// 			},
		// 	}},
		// }
	}
}
