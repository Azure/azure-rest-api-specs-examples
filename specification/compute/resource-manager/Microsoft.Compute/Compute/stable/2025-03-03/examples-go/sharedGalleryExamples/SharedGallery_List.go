package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/sharedGalleryExamples/SharedGallery_List.json
func ExampleSharedGalleriesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSharedGalleriesClient().NewListPager("myLocation", nil)
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
		// page = armcompute.SharedGalleriesClientListResponse{
		// 	SharedGalleryList: armcompute.SharedGalleryList{
		// 		Value: []*armcompute.SharedGallery{
		// 			{
		// 				Location: to.Ptr("myLocation"),
		// 				Name: to.Ptr("galleryUniqueName"),
		// 				Identifier: &armcompute.SharedGalleryIdentifier{
		// 					UniqueID: to.Ptr("/SharedGalleries/galleryUniqueName"),
		// 				},
		// 				Properties: &armcompute.SharedGalleryProperties{
		// 					ArtifactTags: map[string]*string{
		// 						"ShareTag-Official1PGallery": to.Ptr("Official1PGallery"),
		// 					},
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("http://svchost:99/subscriptions/subscriptionId/providers/Microsoft.Compute/sharedGalleries?$skiptoken=token/Subscriptions/subscriptionId/galleries/galleryUniqueName"),
		// 	},
		// }
	}
}
