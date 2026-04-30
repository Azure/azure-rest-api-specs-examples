package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/galleryExamples/GalleryApplicationVersion_Get.json
func ExampleGalleryApplicationVersionsClient_Get_getAGalleryApplicationVersion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleryApplicationVersionsClient().Get(ctx, "myResourceGroup", "myGalleryName", "myGalleryApplicationName", "1.0.0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.GalleryApplicationVersionsClientGetResponse{
	// 	GalleryApplicationVersion: &armcompute.GalleryApplicationVersion{
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
	// 				CustomActions: []*armcompute.GalleryApplicationCustomAction{
	// 					{
	// 						Name: to.Ptr("myCustomAction"),
	// 						Script: to.Ptr("myCustomActionScript"),
	// 						Description: to.Ptr("This is the custom action description."),
	// 						Parameters: []*armcompute.GalleryApplicationCustomActionParameter{
	// 							{
	// 								Name: to.Ptr("myCustomActionParameter"),
	// 								Required: to.Ptr(false),
	// 								Type: to.Ptr(armcompute.GalleryApplicationCustomActionParameterTypeString),
	// 								DefaultValue: to.Ptr("default value of parameter."),
	// 								Description: to.Ptr("This is the description of the parameter"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			SafetyProfile: &armcompute.GalleryApplicationVersionSafetyProfile{
	// 				AllowDeletionOfReplicatedLocations: to.Ptr(false),
	// 			},
	// 			ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
