package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/GalleryApplicationVersion_ListByGalleryApplication.json
func ExampleGalleryApplicationVersionsClient_NewListByGalleryApplicationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGalleryApplicationVersionsClient().NewListByGalleryApplicationPager("myResourceGroup", "myGalleryName", "myGalleryApplicationName", nil)
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
		// page.GalleryApplicationVersionList = armcompute.GalleryApplicationVersionList{
		// 	Value: []*armcompute.GalleryApplicationVersion{
		// 		{
		// 			Name: to.Ptr("1.0.0"),
		// 			Type: to.Ptr("Microsoft.Compute/galleries/applications/versions"),
		// 			ID: to.Ptr("/subscriptions/01523d7c-60da-455e-adef-521b547922c4/resourceGroups/galleryPsTestRg98/providers/Microsoft.Compute/galleries/galleryPsTestGallery6165/applications/galleryPsTestGalleryApplication7825/versions/1.0.0"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcompute.GalleryApplicationVersionProperties{
		// 				ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
		// 				PublishingProfile: &armcompute.GalleryApplicationVersionPublishingProfile{
		// 					EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-01T07:00:00.000Z"); return t}()),
		// 					ExcludeFromLatest: to.Ptr(false),
		// 					PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-21T17:13:57.597Z"); return t}()),
		// 					ReplicaCount: to.Ptr[int32](1),
		// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
		// 					TargetRegions: []*armcompute.TargetRegion{
		// 						{
		// 							Name: to.Ptr("West US"),
		// 							ExcludeFromLatest: to.Ptr(false),
		// 							RegionalReplicaCount: to.Ptr[int32](1),
		// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
		// 					}},
		// 					CustomActions: []*armcompute.GalleryApplicationCustomAction{
		// 						{
		// 							Name: to.Ptr("myCustomAction"),
		// 							Description: to.Ptr("This is the custom action description."),
		// 							Parameters: []*armcompute.GalleryApplicationCustomActionParameter{
		// 								{
		// 									Name: to.Ptr("myCustomActionParameter"),
		// 									Type: to.Ptr(armcompute.GalleryApplicationCustomActionParameterTypeString),
		// 									Description: to.Ptr("This is the description of the parameter"),
		// 									DefaultValue: to.Ptr("default value of parameter."),
		// 									Required: to.Ptr(false),
		// 							}},
		// 							Script: to.Ptr("myCustomActionScript"),
		// 					}},
		// 					EnableHealthCheck: to.Ptr(false),
		// 					ManageActions: &armcompute.UserArtifactManage{
		// 						Install: to.Ptr("powershell -command \"Expand-Archive -Path package.zip -DestinationPath C:\\package\""),
		// 						Remove: to.Ptr("del C:\\package "),
		// 					},
		// 					Source: &armcompute.UserArtifactSource{
		// 						MediaLink: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/package.zip?{sasKey}"),
		// 					},
		// 				},
		// 				SafetyProfile: &armcompute.GalleryApplicationVersionSafetyProfile{
		// 					AllowDeletionOfReplicatedLocations: to.Ptr(false),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
