package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/GalleryApplicationVersion_Get_WithReplicationStatus.json
func ExampleGalleryApplicationVersionsClient_Get_getAGalleryApplicationVersionWithReplicationStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewGalleryApplicationVersionsClient().Get(ctx, "myResourceGroup", "myGalleryName", "myGalleryApplicationName", "1.0.0", &armcompute.GalleryApplicationVersionsClientGetOptions{Expand: to.Ptr(armcompute.ReplicationStatusTypesReplicationStatus)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.GalleryApplicationVersion = armcompute.GalleryApplicationVersion{
	// 	Name: to.Ptr("1.0.0"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myGalleryName/applications/myGalleryApplicationName/versions/1.0.0"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcompute.GalleryApplicationVersionProperties{
	// 		ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateSucceeded),
	// 		PublishingProfile: &armcompute.GalleryApplicationVersionPublishingProfile{
	// 			EndOfLifeDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-01T07:00:00.000Z"); return t}()),
	// 			ExcludeFromLatest: to.Ptr(false),
	// 			PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-06-21T17:13:57.597Z"); return t}()),
	// 			ReplicaCount: to.Ptr[int32](1),
	// 			StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 			TargetRegions: []*armcompute.TargetRegion{
	// 				{
	// 					Name: to.Ptr("West US"),
	// 					ExcludeFromLatest: to.Ptr(false),
	// 					RegionalReplicaCount: to.Ptr[int32](1),
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 			}},
	// 			AdvancedSettings: map[string]*string{
	// 				"timeout": to.Ptr("300"),
	// 				"user": to.Ptr("root"),
	// 			},
	// 			CustomActions: []*armcompute.GalleryApplicationCustomAction{
	// 				{
	// 					Name: to.Ptr("myCustomAction"),
	// 					Description: to.Ptr("This is the custom action description."),
	// 					Parameters: []*armcompute.GalleryApplicationCustomActionParameter{
	// 						{
	// 							Name: to.Ptr("myCustomActionParameter"),
	// 							Type: to.Ptr(armcompute.GalleryApplicationCustomActionParameterTypeString),
	// 							Description: to.Ptr("This is the description of the parameter"),
	// 							DefaultValue: to.Ptr("default value of parameter."),
	// 							Required: to.Ptr(false),
	// 					}},
	// 					Script: to.Ptr("myCustomActionScript"),
	// 			}},
	// 			EnableHealthCheck: to.Ptr(false),
	// 			ManageActions: &armcompute.UserArtifactManage{
	// 				Install: to.Ptr("powershell -command \"Expand-Archive -Path package.zip -DestinationPath C:\\package\""),
	// 				Remove: to.Ptr("del C:\\package "),
	// 			},
	// 			Settings: &armcompute.UserArtifactSettings{
	// 				ConfigFileName: to.Ptr("configuration.cfg"),
	// 				PackageFileName: to.Ptr("package.zip"),
	// 			},
	// 			Source: &armcompute.UserArtifactSource{
	// 				MediaLink: to.Ptr("https://mystorageaccount.blob.core.windows.net/mycontainer/package.zip?{sasKey}"),
	// 			},
	// 		},
	// 		ReplicationStatus: &armcompute.ReplicationStatus{
	// 			AggregatedState: to.Ptr(armcompute.AggregatedReplicationStateCompleted),
	// 			Summary: []*armcompute.RegionalReplicationStatus{
	// 				{
	// 					Progress: to.Ptr[int32](100),
	// 					Region: to.Ptr("West US"),
	// 					State: to.Ptr(armcompute.ReplicationStateCompleted),
	// 					Details: to.Ptr(""),
	// 			}},
	// 		},
	// 		SafetyProfile: &armcompute.GalleryApplicationVersionSafetyProfile{
	// 			AllowDeletionOfReplicatedLocations: to.Ptr(false),
	// 		},
	// 	},
	// }
}
