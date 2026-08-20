package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-12-03/communityGalleryExamples/CommunityGalleryImageVersion_List.json
func ExampleCommunityGalleryImageVersionsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCommunityGalleryImageVersionsClient().NewListPager("myLocation", "publicGalleryName", "myGalleryImageName", nil)
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
		// page = armcompute.CommunityGalleryImageVersionsClientListResponse{
		// 	CommunityGalleryImageVersionList: armcompute.CommunityGalleryImageVersionList{
		// 		Value: []*armcompute.CommunityGalleryImageVersion{
		// 			{
		// 				Properties: &armcompute.CommunityGalleryImageVersionProperties{
		// 					PublishedDate: to.Ptr(time.Date(2018, time.March, 20, 9, 12, 28, 0, time.UTC)),
		// 					EndOfLifeDate: to.Ptr(time.Date(2022, time.March, 20, 9, 12, 28, 0, time.UTC)),
		// 					ExcludeFromLatest: to.Ptr(false),
		// 					StorageProfile: &armcompute.SharedGalleryImageVersionStorageProfile{
		// 						OSDiskImage: &armcompute.SharedGalleryOSDiskImage{
		// 							DiskSizeGB: to.Ptr[int32](29),
		// 							HostCaching: to.Ptr(armcompute.SharedGalleryHostCachingNone),
		// 						},
		// 					},
		// 					ArtifactTags: map[string]*string{
		// 						"ShareTag-CommunityGallery": to.Ptr("CommunityGallery"),
		// 					},
		// 					Disclaimer: to.Ptr("https://test-uri.com"),
		// 				},
		// 				Location: to.Ptr("myLocation"),
		// 				Name: to.Ptr("myGalleryImageVersionName"),
		// 				Identifier: &armcompute.CommunityGalleryIdentifier{
		// 					UniqueID: to.Ptr("/CommunityGalleries/publicGalleryName/Images/myGalleryImageName/Versions/myGalleryImageVersionName"),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("http://svchost:99/subscriptions/subscriptionId/providers/Microsoft.Compute/communityGalleries/publicGalleryName/images/myGalleryImageName/versions?$skiptoken=token/communityGalleries/publicGalleryName/images/myGalleryImageName/versions/myGalleryImageVersionName"),
		// 	},
		// }
	}
}
