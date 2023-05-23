package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5d2adf9b7fda669b4a2538c65e937ee74fe3f966/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-03-03/examples/galleryExamples/GalleryApplicationVersion_Update.json
func ExampleGalleryApplicationVersionsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryApplicationVersionsClient().BeginUpdate(ctx, "myResourceGroup", "myGalleryName", "myGalleryApplicationName", "1.0.0", armcompute.GalleryApplicationVersionUpdate{
		Properties: &armcompute.GalleryApplicationVersionProperties{
			PublishingProfile: &armcompute.GalleryApplicationVersionPublishingProfile{
				EndOfLifeDate:      to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-01T07:00:00Z"); return t }()),
				ReplicaCount:       to.Ptr[int32](1),
				StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
				TargetRegions: []*armcompute.TargetRegion{
					{
						Name:                 to.Ptr("West US"),
						ExcludeFromLatest:    to.Ptr(false),
						RegionalReplicaCount: to.Ptr[int32](1),
						StorageAccountType:   to.Ptr(armcompute.StorageAccountTypeStandardLRS),
					}},
				ManageActions: &armcompute.UserArtifactManage{
					Install: to.Ptr("powershell -command \"Expand-Archive -Path package.zip -DestinationPath C:\\package\""),
					Remove:  to.Ptr("del C:\\package "),
				},
				Source: &armcompute.UserArtifactSource{
					MediaLink: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/package.zip?{sasKey}"),
				},
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
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GalleryApplicationVersion = armcompute.GalleryApplicationVersion{
	// 	Name: to.Ptr("1.0.0"),
	// 	Type: to.Ptr("Microsoft.Compute/galleries/applications/versions"),
	// 	ID: to.Ptr("/subscriptions/01523d7c-60da-455e-adef-521b547922c4/resourceGroups/galleryPsTestRg98/providers/Microsoft.Compute/galleries/galleryPsTestGallery6165/applications/galleryPsTestGalleryApplication7825/versions/1.0.0"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryApplicationVersionProperties{
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		PublishingProfile: &armcompute.GalleryApplicationVersionPublishingProfile{
	// 			EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-01T07:00:00+00:00"); return t}()),
	// 			ExcludeFromLatest: to.Ptr(false),
	// 			PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-21T17:13:57.5972568+00:00"); return t}()),
	// 			ReplicaCount: to.Ptr[int32](1),
	// 			StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 			TargetRegions: []*armcompute.TargetRegion{
	// 				{
	// 					Name: to.Ptr("West US"),
	// 					ExcludeFromLatest: to.Ptr(false),
	// 					RegionalReplicaCount: to.Ptr[int32](1),
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 			}},
	// 			EnableHealthCheck: to.Ptr(false),
	// 			ManageActions: &armcompute.UserArtifactManage{
	// 				Install: to.Ptr("powershell -command \"Expand-Archive -Path package.zip -DestinationPath C:\\package\""),
	// 				Remove: to.Ptr("del C:\\package "),
	// 			},
	// 			Source: &armcompute.UserArtifactSource{
	// 				MediaLink: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/package.zip?{sasKey}"),
	// 			},
	// 		},
	// 		SafetyProfile: &armcompute.GalleryApplicationVersionSafetyProfile{
	// 			AllowDeletionOfReplicatedLocations: to.Ptr(false),
	// 		},
	// 	},
	// }
}
