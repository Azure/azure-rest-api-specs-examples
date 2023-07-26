package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/sharedGalleryExamples/SharedGalleryImageVersions_List.json
func ExampleSharedGalleryImageVersionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSharedGalleryImageVersionsClient().NewListPager("myLocation", "galleryUniqueName", "myGalleryImageName", &armcompute.SharedGalleryImageVersionsClientListOptions{SharedTo: nil})
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
		// page.SharedGalleryImageVersionList = armcompute.SharedGalleryImageVersionList{
		// 	Value: []*armcompute.SharedGalleryImageVersion{
		// 		{
		// 			Name: to.Ptr("myGalleryImageVersionName"),
		// 			Location: to.Ptr("myLocation"),
		// 			Identifier: &armcompute.SharedGalleryIdentifier{
		// 				UniqueID: to.Ptr("/SharedGalleries/galleryUniqueName/Images/myGalleryImageName/Versions/myGalleryImageVersionName"),
		// 			},
		// 			Properties: &armcompute.SharedGalleryImageVersionProperties{
		// 				EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-20T09:12:28Z"); return t}()),
		// 				ExcludeFromLatest: to.Ptr(false),
		// 				PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-20T09:12:28Z"); return t}()),
		// 				StorageProfile: &armcompute.SharedGalleryImageVersionStorageProfile{
		// 					OSDiskImage: &armcompute.SharedGalleryOSDiskImage{
		// 						DiskSizeGB: to.Ptr[int32](29),
		// 						HostCaching: to.Ptr(armcompute.SharedGalleryHostCachingNone),
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
