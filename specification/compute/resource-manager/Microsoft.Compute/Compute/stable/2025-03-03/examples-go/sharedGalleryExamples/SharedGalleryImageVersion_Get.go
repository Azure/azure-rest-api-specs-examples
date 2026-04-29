package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/sharedGalleryExamples/SharedGalleryImageVersion_Get.json
func ExampleSharedGalleryImageVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSharedGalleryImageVersionsClient().Get(ctx, "myLocation", "galleryUniqueName", "myGalleryImageName", "myGalleryImageVersionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.SharedGalleryImageVersionsClientGetResponse{
	// 	SharedGalleryImageVersion: &armcompute.SharedGalleryImageVersion{
	// 		Properties: &armcompute.SharedGalleryImageVersionProperties{
	// 			PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-03-20T09:12:28Z"); return t}()),
	// 			EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-03-20T09:12:28Z"); return t}()),
	// 			ExcludeFromLatest: to.Ptr(false),
	// 			StorageProfile: &armcompute.SharedGalleryImageVersionStorageProfile{
	// 				OSDiskImage: &armcompute.SharedGalleryOSDiskImage{
	// 					DiskSizeGB: to.Ptr[int32](29),
	// 					HostCaching: to.Ptr(armcompute.SharedGalleryHostCachingNone),
	// 				},
	// 			},
	// 			ArtifactTags: map[string]*string{
	// 				"ShareTag-Official1PGallery": to.Ptr("Official1PGallery"),
	// 			},
	// 		},
	// 		Location: to.Ptr("myLocation"),
	// 		Name: to.Ptr("myGalleryImageVersionName"),
	// 		Identifier: &armcompute.SharedGalleryIdentifier{
	// 			UniqueID: to.Ptr("/SharedGalleries/galleryUniqueName/Images/myGalleryImageName/Versions/myGalleryImageVersionName"),
	// 		},
	// 	},
	// }
}
