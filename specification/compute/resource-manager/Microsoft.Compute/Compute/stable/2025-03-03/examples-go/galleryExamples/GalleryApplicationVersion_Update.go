package armcompute_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/galleryExamples/GalleryApplicationVersion_Update.json
func ExampleGalleryApplicationVersionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryApplicationVersionsClient().BeginUpdate(ctx, "myResourceGroup", "myGalleryName", "myGalleryApplicationName", "1.0.0", armcompute.GalleryApplicationVersionUpdate{
		Properties: &armcompute.GalleryApplicationVersionProperties{
			PublishingProfile: &armcompute.GalleryApplicationVersionPublishingProfile{
				Source: &armcompute.UserArtifactSource{
					MediaLink: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/package.zip?{sasKey}"),
				},
				ManageActions: &armcompute.UserArtifactManage{
					Install: to.Ptr("powershell -command \"Expand-Archive -Path package.zip -DestinationPath C:\\package\""),
					Remove:  to.Ptr("del C:\\package "),
				},
				TargetRegions: []*armcompute.TargetRegion{
					{
						Name:                 to.Ptr("West US"),
						RegionalReplicaCount: to.Ptr[int32](1),
						StorageAccountType:   to.Ptr(armcompute.StorageAccountTypeStandardLRS),
						ExcludeFromLatest:    to.Ptr(false),
					},
				},
				ReplicaCount:       to.Ptr[int32](1),
				EndOfLifeDate:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-01T07:00:00Z"); return t }()),
				StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
			},
			SafetyProfile: &armcompute.GalleryApplicationVersionSafetyProfile{
				AllowDeletionOfReplicatedLocations: to.Ptr(false),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.GalleryApplicationVersionsClientUpdateResponse{
	// 	GalleryApplicationVersion: armcompute.GalleryApplicationVersion{
	// 		Location: to.Ptr("West US"),
	// 		Name: to.Ptr("1.0.0"),
	// 		Type: to.Ptr("Microsoft.Compute/galleries/applications/versions"),
	// 		ID: to.Ptr("/subscriptions/01523d7c-60da-455e-adef-521b547922c4/resourceGroups/galleryPsTestRg98/providers/Microsoft.Compute/galleries/galleryPsTestGallery6165/applications/galleryPsTestGalleryApplication7825/versions/1.0.0"),
	// 		Properties: &armcompute.GalleryApplicationVersionProperties{
	// 			PublishingProfile: &armcompute.GalleryApplicationVersionPublishingProfile{
	// 				Source: &armcompute.UserArtifactSource{
	// 					MediaLink: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/package.zip?{sasKey}"),
	// 				},
	// 				ManageActions: &armcompute.UserArtifactManage{
	// 					Install: to.Ptr("powershell -command \"Expand-Archive -Path package.zip -DestinationPath C:\\package\""),
	// 					Remove: to.Ptr("del C:\\package "),
	// 				},
	// 				EnableHealthCheck: to.Ptr(false),
	// 				TargetRegions: []*armcompute.TargetRegion{
	// 					{
	// 						Name: to.Ptr("West US"),
	// 						RegionalReplicaCount: to.Ptr[int32](1),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 						ExcludeFromLatest: to.Ptr(false),
	// 					},
	// 				},
	// 				ReplicaCount: to.Ptr[int32](1),
	// 				ExcludeFromLatest: to.Ptr(false),
	// 				PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-21T17:13:57.5972568+00:00"); return t}()),
	// 				EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-01T07:00:00+00:00"); return t}()),
	// 				StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 			},
	// 			SafetyProfile: &armcompute.GalleryApplicationVersionSafetyProfile{
	// 				AllowDeletionOfReplicatedLocations: to.Ptr(false),
	// 			},
	// 			ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateUpdating),
	// 		},
	// 	},
	// }
}
