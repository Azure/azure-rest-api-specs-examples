package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-01-02/snapshotExamples/Snapshot_ListBySubscription.json
func ExampleSnapshotsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSnapshotsClient().NewListPager(nil)
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
		// page = armcompute.SnapshotsClientListResponse{
		// 	SnapshotList: armcompute.SnapshotList{
		// 		Value: []*armcompute.Snapshot{
		// 			{
		// 				Properties: &armcompute.SnapshotProperties{
		// 					OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 					CreationData: &armcompute.CreationData{
		// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionCopy),
		// 						SourceResourceID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
		// 					},
		// 					DiskSizeGB: to.Ptr[int32](200),
		// 					EncryptionSettingsCollection: &armcompute.EncryptionSettingsCollection{
		// 						Enabled: to.Ptr(true),
		// 						EncryptionSettings: []*armcompute.EncryptionSettingsElement{
		// 							{
		// 								DiskEncryptionKey: &armcompute.KeyVaultAndSecretReference{
		// 									SourceVault: &armcompute.SourceVault{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
		// 									},
		// 									SecretURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/secrets/{secret}"),
		// 								},
		// 								KeyEncryptionKey: &armcompute.KeyVaultAndKeyReference{
		// 									SourceVault: &armcompute.SourceVault{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
		// 									},
		// 									KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					Encryption: &armcompute.Encryption{
		// 						Type: to.Ptr(armcompute.EncryptionTypeEncryptionAtRestWithPlatformKey),
		// 					},
		// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-28T04:47:30.6630569+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 				Type: to.Ptr("Microsoft.Compute/snapshots"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 					"department": to.Ptr("Development"),
		// 					"project": to.Ptr("Snapshots"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot1"),
		// 				Name: to.Ptr("mySnapshot1"),
		// 			},
		// 			{
		// 				Properties: &armcompute.SnapshotProperties{
		// 					OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 					CreationData: &armcompute.CreationData{
		// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionImport),
		// 						StorageAccountID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount"),
		// 						SourceURI: to.Ptr("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
		// 					},
		// 					DiskSizeGB: to.Ptr[int32](200),
		// 					EncryptionSettingsCollection: &armcompute.EncryptionSettingsCollection{
		// 						Enabled: to.Ptr(true),
		// 						EncryptionSettings: []*armcompute.EncryptionSettingsElement{
		// 							{
		// 								DiskEncryptionKey: &armcompute.KeyVaultAndSecretReference{
		// 									SourceVault: &armcompute.SourceVault{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
		// 									},
		// 									SecretURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/secrets/{secret}"),
		// 								},
		// 								KeyEncryptionKey: &armcompute.KeyVaultAndKeyReference{
		// 									SourceVault: &armcompute.SourceVault{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
		// 									},
		// 									KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					Encryption: &armcompute.Encryption{
		// 						Type: to.Ptr(armcompute.EncryptionTypeEncryptionAtRestWithPlatformKey),
		// 					},
		// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-28T04:47:30.3247198+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 				Type: to.Ptr("Microsoft.Compute/snapshots"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 					"department": to.Ptr("Development"),
		// 					"project": to.Ptr("Snapshots"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot2"),
		// 				Name: to.Ptr("mySnapshot2"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
