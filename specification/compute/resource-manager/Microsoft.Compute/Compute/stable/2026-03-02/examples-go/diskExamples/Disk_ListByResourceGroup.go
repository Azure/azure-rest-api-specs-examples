package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-02/diskExamples/Disk_ListByResourceGroup.json
func ExampleDisksClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDisksClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page = armcompute.DisksClientListByResourceGroupResponse{
		// 	DiskList: armcompute.DiskList{
		// 		Value: []*armcompute.Disk{
		// 			{
		// 				Properties: &armcompute.DiskProperties{
		// 					OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 					CreationData: &armcompute.CreationData{
		// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionCopy),
		// 						SourceResourceID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
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
		// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-28T04:41:35.9278721+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 				Type: to.Ptr("Microsoft.Compute/disks"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 					"department": to.Ptr("Development"),
		// 					"project": to.Ptr("ManagedDisks"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
		// 				Name: to.Ptr("myManagedDisk1"),
		// 			},
		// 			{
		// 				Properties: &armcompute.DiskProperties{
		// 					OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 					CreationData: &armcompute.CreationData{
		// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionEmpty),
		// 					},
		// 					DiskSizeGB: to.Ptr[int32](10),
		// 					Encryption: &armcompute.Encryption{
		// 						Type: to.Ptr(armcompute.EncryptionTypeEncryptionAtRestWithPlatformKey),
		// 					},
		// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-28T04:41:36.872242+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 				Type: to.Ptr("Microsoft.Compute/disks"),
		// 				Location: to.Ptr("westus"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
		// 				Name: to.Ptr("myManagedDisk2"),
		// 			},
		// 			{
		// 				Properties: &armcompute.DiskProperties{
		// 					OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 					CreationData: &armcompute.CreationData{
		// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionFromImage),
		// 						ImageReference: &armcompute.ImageDiskReference{
		// 							ID: to.Ptr("/Subscriptions/{subscriptionId}/Providers/Microsoft.Compute/Locations/uswest/Publishers/Microsoft/ArtifactTypes/VMImage/Offers/{offer}"),
		// 						},
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
		// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-28T04:41:36.3973934+00:00"); return t}()),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 				Type: to.Ptr("Microsoft.Compute/disks"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 					"department": to.Ptr("Development"),
		// 					"project": to.Ptr("ManagedDisks"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
		// 				Name: to.Ptr("myManagedDisk3"),
		// 			},
		// 		},
		// 		NextLink: to.Ptr("http://disksvchost:99/subscriptions/subscriptionId/providers/Microsoft.Compute/disks?$skiptoken=token/Subscriptions/subscriptionId/ResourceGroups/myResourceGroup/Disks/myManagedDisk"),
		// 	},
		// }
	}
}
