package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-03-03/galleryExamples/GalleryImageVersion_Create_WithStorageAccountStrategy.json
func ExampleGalleryImageVersionsClient_BeginCreateOrUpdate_createOrUpdateASimpleGalleryImageVersionWithStorageAccountStrategyAndRegionalStorageAccountTypeOverride() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewGalleryImageVersionsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myGalleryName", "myGalleryImageName", "1.0.0", armcompute.GalleryImageVersion{
		Location: to.Ptr("West US"),
		Properties: &armcompute.GalleryImageVersionProperties{
			PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
				TargetRegions: []*armcompute.TargetRegion{
					{
						Name: to.Ptr("West US"),
					},
					{
						Name: to.Ptr("East US"),
					},
					{
						Name:               to.Ptr("East US 2"),
						StorageAccountType: to.Ptr(armcompute.StorageAccountTypePremiumLRS),
					},
				},
				StorageAccountStrategy: to.Ptr(armcompute.StorageAccountStrategyPreferStandardZRS),
			},
			StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
				Source: &armcompute.GalleryArtifactVersionFullSource{
					VirtualMachineID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/virtualMachines/{vmName}"),
				},
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
	// res = armcompute.GalleryImageVersionsClientCreateOrUpdateResponse{
	// 	GalleryImageVersion: &armcompute.GalleryImageVersion{
	// 		ID: to.Ptr("/providers/Microsoft.Compute/locations/westus/Galleries/myGalleryName/Images/myGalleryImageName/Versions/1.0.0"),
	// 		Properties: &armcompute.GalleryImageVersionProperties{
	// 			PublishingProfile: &armcompute.GalleryImageVersionPublishingProfile{
	// 				TargetRegions: []*armcompute.TargetRegion{
	// 					{
	// 						Name: to.Ptr("West US"),
	// 						RegionalReplicaCount: to.Ptr[int32](1),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardLRS),
	// 					},
	// 					{
	// 						Name: to.Ptr("East US"),
	// 						RegionalReplicaCount: to.Ptr[int32](1),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypeStandardZRS),
	// 					},
	// 					{
	// 						Name: to.Ptr("East US 2"),
	// 						RegionalReplicaCount: to.Ptr[int32](1),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypePremiumLRS),
	// 					},
	// 				},
	// 				ReplicaCount: to.Ptr[int32](1),
	// 				PublishedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-01T00:00:00Z"); return t}()),
	// 				StorageAccountStrategy: to.Ptr(armcompute.StorageAccountStrategyPreferStandardZRS),
	// 				ExcludeFromLatest: to.Ptr(false),
	// 			},
	// 			StorageProfile: &armcompute.GalleryImageVersionStorageProfile{
	// 				Source: &armcompute.GalleryArtifactVersionFullSource{
	// 					VirtualMachineID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/virtualMachines/{vmName}"),
	// 				},
	// 			},
	// 			SafetyProfile: &armcompute.GalleryImageVersionSafetyProfile{
	// 				ReportedForPolicyViolation: to.Ptr(false),
	// 				AllowDeletionOfReplicatedLocations: to.Ptr(false),
	// 			},
	// 			ProvisioningState: to.Ptr(armcompute.GalleryProvisioningStateUpdating),
	// 		},
	// 		Location: to.Ptr("West US"),
	// 		Name: to.Ptr("1.0.0"),
	// 	},
	// }
}
