package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e2b5f9323c4214408969a6e953b4075cfdc693b6/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2023-05-01/examples/amlFilesystems_ListByResourceGroup.json
func ExampleAmlFilesystemsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAmlFilesystemsClient().NewListByResourceGroupPager("scgroup", nil)
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
		// page.AmlFilesystemsListResult = armstoragecache.AmlFilesystemsListResult{
		// 	Value: []*armstoragecache.AmlFilesystem{
		// 		{
		// 			Name: to.Ptr("fs1"),
		// 			Type: to.Ptr("Microsoft.StorageCache/amlFilesystem"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlFilesystems/fs1"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"Dept": to.Ptr("ContosoAds"),
		// 			},
		// 			Identity: &armstoragecache.AmlFilesystemIdentity{
		// 				Type: to.Ptr(armstoragecache.AmlFilesystemIdentityTypeUserAssigned),
		// 				UserAssignedIdentities: map[string]*armstoragecache.UserAssignedIdentitiesValue{
		// 					"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armstoragecache.UserAssignedIdentitiesValue{
		// 					},
		// 				},
		// 			},
		// 			Properties: &armstoragecache.AmlFilesystemProperties{
		// 				ClientInfo: &armstoragecache.AmlFilesystemClientInfo{
		// 					ContainerStorageInterface: &armstoragecache.AmlFilesystemContainerStorageInterface{
		// 						PersistentVolume: to.Ptr("<Base64 encoded YAML>"),
		// 						PersistentVolumeClaim: to.Ptr("<Base64 encoded YAML>"),
		// 						StorageClass: to.Ptr("<Base64 encoded YAML>"),
		// 					},
		// 					LustreVersion: to.Ptr("2.15.0"),
		// 					MgsAddress: to.Ptr("10.0.0.4"),
		// 					MountCommand: to.Ptr("mount -t lustre 10.0.0.4@tcp:/lustrefs /lustre/lustrefs"),
		// 				},
		// 				EncryptionSettings: &armstoragecache.AmlFilesystemEncryptionSettings{
		// 					KeyEncryptionKey: &armstoragecache.KeyVaultKeyReference{
		// 						KeyURL: to.Ptr("https://examplekv.vault.azure.net/keys/kvk/3540a47df75541378d3518c6a4bdf5af"),
		// 						SourceVault: &armstoragecache.KeyVaultKeyReferenceSourceVault{
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.KeyVault/vaults/keyvault-cmk"),
		// 						},
		// 					},
		// 				},
		// 				FilesystemSubnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/fsSub1"),
		// 				Health: &armstoragecache.AmlFilesystemHealth{
		// 					State: to.Ptr(armstoragecache.AmlFilesystemHealthStateTypeAvailable),
		// 					StatusDescription: to.Ptr("amlFilesystem is ok."),
		// 				},
		// 				Hsm: &armstoragecache.AmlFilesystemPropertiesHsm{
		// 					ArchiveStatus: []*armstoragecache.AmlFilesystemArchive{
		// 						{
		// 							FilesystemPath: to.Ptr("/"),
		// 							Status: &armstoragecache.AmlFilesystemArchiveStatus{
		// 								LastCompletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T18:25:43.511Z"); return t}()),
		// 								LastStartedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T17:25:43.511Z"); return t}()),
		// 								State: to.Ptr(armstoragecache.ArchiveStatusTypeCompleted),
		// 							},
		// 					}},
		// 					Settings: &armstoragecache.AmlFilesystemHsmSettings{
		// 						Container: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/storageaccountname/blobServices/default/containers/containername"),
		// 						ImportPrefix: to.Ptr("/"),
		// 						LoggingContainer: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/storageaccountname/blobServices/default/containers/loggingcontainername"),
		// 					},
		// 				},
		// 				MaintenanceWindow: &armstoragecache.AmlFilesystemPropertiesMaintenanceWindow{
		// 					DayOfWeek: to.Ptr(armstoragecache.MaintenanceDayOfWeekTypeFriday),
		// 					TimeOfDayUTC: to.Ptr("22:00"),
		// 				},
		// 				ProvisioningState: to.Ptr(armstoragecache.AmlFilesystemProvisioningStateTypeSucceeded),
		// 				StorageCapacityTiB: to.Ptr[float32](16),
		// 				ThroughputProvisionedMBps: to.Ptr[int32](500),
		// 			},
		// 			SKU: &armstoragecache.SKUName{
		// 				Name: to.Ptr("AMLFS-Durable-Premium-250"),
		// 			},
		// 			Zones: []*string{
		// 				to.Ptr("1")},
		// 			},
		// 			{
		// 				Name: to.Ptr("fs2"),
		// 				Type: to.Ptr("Microsoft.StorageCache/amlFilesystem"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlFilesystems/fs2"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("ContosoAds"),
		// 				},
		// 				Identity: &armstoragecache.AmlFilesystemIdentity{
		// 					Type: to.Ptr(armstoragecache.AmlFilesystemIdentityTypeUserAssigned),
		// 					UserAssignedIdentities: map[string]*armstoragecache.UserAssignedIdentitiesValue{
		// 						"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armstoragecache.UserAssignedIdentitiesValue{
		// 						},
		// 					},
		// 				},
		// 				Properties: &armstoragecache.AmlFilesystemProperties{
		// 					ClientInfo: &armstoragecache.AmlFilesystemClientInfo{
		// 						ContainerStorageInterface: &armstoragecache.AmlFilesystemContainerStorageInterface{
		// 							PersistentVolume: to.Ptr("<Base64 encoded YAML>"),
		// 							PersistentVolumeClaim: to.Ptr("<Base64 encoded YAML>"),
		// 							StorageClass: to.Ptr("<Base64 encoded YAML>"),
		// 						},
		// 						LustreVersion: to.Ptr("2.15.0"),
		// 						MgsAddress: to.Ptr("10.0.0.4"),
		// 						MountCommand: to.Ptr("mount -t lustre 10.0.0.4@tcp:/lustrefs /lustre/lustrefs"),
		// 					},
		// 					EncryptionSettings: &armstoragecache.AmlFilesystemEncryptionSettings{
		// 						KeyEncryptionKey: &armstoragecache.KeyVaultKeyReference{
		// 							KeyURL: to.Ptr("https://examplekv.vault.azure.net/keys/kvk/3540a47df75541378d3518c6a4bdf5af"),
		// 							SourceVault: &armstoragecache.KeyVaultKeyReferenceSourceVault{
		// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.KeyVault/vaults/keyvault-cmk"),
		// 							},
		// 						},
		// 					},
		// 					FilesystemSubnet: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Network/virtualNetworks/scvnet/subnets/fsSub2"),
		// 					Health: &armstoragecache.AmlFilesystemHealth{
		// 						State: to.Ptr(armstoragecache.AmlFilesystemHealthStateTypeAvailable),
		// 						StatusDescription: to.Ptr("amlFilesystem is ok."),
		// 					},
		// 					Hsm: &armstoragecache.AmlFilesystemPropertiesHsm{
		// 						ArchiveStatus: []*armstoragecache.AmlFilesystemArchive{
		// 							{
		// 								FilesystemPath: to.Ptr("/"),
		// 								Status: &armstoragecache.AmlFilesystemArchiveStatus{
		// 									LastCompletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T18:25:43.511Z"); return t}()),
		// 									LastStartedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T17:25:43.511Z"); return t}()),
		// 									State: to.Ptr(armstoragecache.ArchiveStatusTypeCompleted),
		// 								},
		// 						}},
		// 						Settings: &armstoragecache.AmlFilesystemHsmSettings{
		// 							Container: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/storageaccountname/blobServices/default/containers/containername"),
		// 							ImportPrefix: to.Ptr("/"),
		// 							LoggingContainer: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/storageaccountname/blobServices/default/containers/loggingcontainername"),
		// 						},
		// 					},
		// 					MaintenanceWindow: &armstoragecache.AmlFilesystemPropertiesMaintenanceWindow{
		// 						DayOfWeek: to.Ptr(armstoragecache.MaintenanceDayOfWeekTypeFriday),
		// 						TimeOfDayUTC: to.Ptr("22:00"),
		// 					},
		// 					ProvisioningState: to.Ptr(armstoragecache.AmlFilesystemProvisioningStateTypeSucceeded),
		// 					StorageCapacityTiB: to.Ptr[float32](16),
		// 					ThroughputProvisionedMBps: to.Ptr[int32](500),
		// 				},
		// 				SKU: &armstoragecache.SKUName{
		// 					Name: to.Ptr("AMLFS-Durable-Premium-250"),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("1")},
		// 			}},
		// 		}
	}
}
