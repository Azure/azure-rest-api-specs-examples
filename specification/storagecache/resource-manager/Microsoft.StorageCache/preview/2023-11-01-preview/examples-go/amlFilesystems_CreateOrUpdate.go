package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/storagecache/resource-manager/Microsoft.StorageCache/preview/2023-11-01-preview/examples/amlFilesystems_CreateOrUpdate.json
func ExampleAmlFilesystemsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAmlFilesystemsClient().BeginCreateOrUpdate(ctx, "scgroup", "fs1", armstoragecache.AmlFilesystem{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"Dept": to.Ptr("ContosoAds"),
		},
		Identity: &armstoragecache.AmlFilesystemIdentity{
			Type: to.Ptr(armstoragecache.AmlFilesystemIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armstoragecache.UserAssignedIdentitiesValue{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": {},
			},
		},
		Properties: &armstoragecache.AmlFilesystemProperties{
			EncryptionSettings: &armstoragecache.AmlFilesystemEncryptionSettings{
				KeyEncryptionKey: &armstoragecache.KeyVaultKeyReference{
					KeyURL: to.Ptr("https://examplekv.vault.azure.net/keys/kvk/3540a47df75541378d3518c6a4bdf5af"),
					SourceVault: &armstoragecache.KeyVaultKeyReferenceSourceVault{
						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.KeyVault/vaults/keyvault-cmk"),
					},
				},
			},
			FilesystemSubnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/fsSub"),
			Hsm: &armstoragecache.AmlFilesystemPropertiesHsm{
				Settings: &armstoragecache.AmlFilesystemHsmSettings{
					Container:        to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/storageaccountname/blobServices/default/containers/containername"),
					ImportPrefix:     to.Ptr("/"),
					LoggingContainer: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/storageaccountname/blobServices/default/containers/loggingcontainername"),
				},
			},
			MaintenanceWindow: &armstoragecache.AmlFilesystemPropertiesMaintenanceWindow{
				DayOfWeek:    to.Ptr(armstoragecache.MaintenanceDayOfWeekTypeFriday),
				TimeOfDayUTC: to.Ptr("22:00"),
			},
			RootSquashSettings: &armstoragecache.AmlFilesystemRootSquashSettings{
				Mode:             to.Ptr(armstoragecache.AmlFilesystemSquashModeAll),
				NoSquashNidLists: to.Ptr("10.0.0.[5-6]@tcp;10.0.1.2@tcp"),
				SquashGID:        to.Ptr[int64](99),
				SquashUID:        to.Ptr[int64](99),
			},
			StorageCapacityTiB: to.Ptr[float32](16),
		},
		SKU: &armstoragecache.SKUName{
			Name: to.Ptr("AMLFS-Durable-Premium-250"),
		},
		Zones: []*string{
			to.Ptr("1")},
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
	// res.AmlFilesystem = armstoragecache.AmlFilesystem{
	// 	Name: to.Ptr("fs1"),
	// 	Type: to.Ptr("Microsoft.StorageCache/amlFilesystem"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlFilesystems/fs1"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"Dept": to.Ptr("ContosoAds"),
	// 	},
	// 	Identity: &armstoragecache.AmlFilesystemIdentity{
	// 		Type: to.Ptr(armstoragecache.AmlFilesystemIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armstoragecache.UserAssignedIdentitiesValue{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armstoragecache.UserAssignedIdentitiesValue{
	// 			},
	// 		},
	// 	},
	// 	Properties: &armstoragecache.AmlFilesystemProperties{
	// 		ClientInfo: &armstoragecache.AmlFilesystemClientInfo{
	// 			ContainerStorageInterface: &armstoragecache.AmlFilesystemContainerStorageInterface{
	// 				PersistentVolume: to.Ptr("<Base64 encoded YAML>"),
	// 				PersistentVolumeClaim: to.Ptr("<Base64 encoded YAML>"),
	// 				StorageClass: to.Ptr("<Base64 encoded YAML>"),
	// 			},
	// 			LustreVersion: to.Ptr("2.15.0"),
	// 			MgsAddress: to.Ptr("10.0.0.4"),
	// 			MountCommand: to.Ptr("mount -t lustre 10.0.0.4@tcp:/lustrefs /lustre/lustrefs"),
	// 		},
	// 		FilesystemSubnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/fsSub"),
	// 		Health: &armstoragecache.AmlFilesystemHealth{
	// 			State: to.Ptr(armstoragecache.AmlFilesystemHealthStateTypeAvailable),
	// 			StatusDescription: to.Ptr("amlFilesystem is ok."),
	// 		},
	// 		Hsm: &armstoragecache.AmlFilesystemPropertiesHsm{
	// 			ArchiveStatus: []*armstoragecache.AmlFilesystemArchive{
	// 				{
	// 					FilesystemPath: to.Ptr("/"),
	// 					Status: &armstoragecache.AmlFilesystemArchiveStatus{
	// 						LastCompletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T18:25:43.511Z"); return t}()),
	// 						LastStartedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T17:25:43.511Z"); return t}()),
	// 						State: to.Ptr(armstoragecache.ArchiveStatusTypeCompleted),
	// 					},
	// 			}},
	// 			Settings: &armstoragecache.AmlFilesystemHsmSettings{
	// 				Container: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/storageaccountname/blobServices/default/containers/containername"),
	// 				ImportPrefix: to.Ptr("/"),
	// 				LoggingContainer: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/storageaccountname/blobServices/default/containers/loggingcontainername"),
	// 			},
	// 		},
	// 		MaintenanceWindow: &armstoragecache.AmlFilesystemPropertiesMaintenanceWindow{
	// 			DayOfWeek: to.Ptr(armstoragecache.MaintenanceDayOfWeekTypeFriday),
	// 			TimeOfDayUTC: to.Ptr("22:00"),
	// 		},
	// 		ProvisioningState: to.Ptr(armstoragecache.AmlFilesystemProvisioningStateTypeSucceeded),
	// 		RootSquashSettings: &armstoragecache.AmlFilesystemRootSquashSettings{
	// 			Mode: to.Ptr(armstoragecache.AmlFilesystemSquashModeAll),
	// 			NoSquashNidLists: to.Ptr("10.0.0.[5-6]@tcp;10.0.1.2@tcp"),
	// 			SquashGID: to.Ptr[int64](99),
	// 			SquashUID: to.Ptr[int64](99),
	// 			Status: to.Ptr("nodemap.active=1"),
	// 		},
	// 		StorageCapacityTiB: to.Ptr[float32](16),
	// 		ThroughputProvisionedMBps: to.Ptr[int32](500),
	// 	},
	// 	SKU: &armstoragecache.SKUName{
	// 		Name: to.Ptr("AMLFS-Durable-Premium-250"),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}
