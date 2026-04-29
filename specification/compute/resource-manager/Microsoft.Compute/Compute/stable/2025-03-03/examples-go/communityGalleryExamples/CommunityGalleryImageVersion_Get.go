package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/communityGalleryExamples/CommunityGalleryImageVersion_Get.json
func ExampleCommunityGalleryImageVersionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCommunityGalleryImageVersionsClient().Get(ctx, "myLocation", "publicGalleryName", "myGalleryImageName", "myGalleryImageVersionName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.CommunityGalleryImageVersionsClientGetResponse{
	// 	CommunityGalleryImageVersion: &armcompute.CommunityGalleryImageVersion{
	// 		Properties: &armcompute.CommunityGalleryImageVersionProperties{
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
	// 				"ShareTag-CommunityGallery": to.Ptr("CommunityGallery"),
	// 			},
	// 			Disclaimer: to.Ptr("https://test-uri.com"),
	// 		},
	// 		Location: to.Ptr("myLocation"),
	// 		Name: to.Ptr("myGalleryImageVersionName"),
	// 		Type: to.Ptr("Microsoft.Compute/locations/communityGalleryImageVersion"),
	// 		Identifier: &armcompute.CommunityGalleryIdentifier{
	// 			UniqueID: to.Ptr("/CommunityGalleries/publicGalleryName/Images/myGalleryImageName/Versions/myGalleryImageVersionName"),
	// 		},
	// 	},
	// }
}
