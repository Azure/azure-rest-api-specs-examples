package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/examples/diskExamples/Disk_ListBySubscription.json
func ExampleDisksClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDisksClient().NewListPager(nil)
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
		// page.DiskList = armcompute.DiskList{
		// 	Value: []*armcompute.Disk{
		// 		{
		// 			Name: to.Ptr("myManagedDisk1"),
		// 			Type: to.Ptr("Microsoft.Compute/disks"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk1"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"department": to.Ptr("Development"),
		// 				"project": to.Ptr("ManagedDisks"),
		// 			},
		// 			Properties: &armcompute.DiskProperties{
		// 				CreationData: &armcompute.CreationData{
		// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionCopy),
		// 					SourceResourceID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk1"),
		// 				},
		// 				DiskSizeGB: to.Ptr[int32](200),
		// 				Encryption: &armcompute.Encryption{
		// 					Type: to.Ptr(armcompute.EncryptionTypeEncryptionAtRestWithPlatformKey),
		// 				},
		// 				EncryptionSettingsCollection: &armcompute.EncryptionSettingsCollection{
		// 					Enabled: to.Ptr(true),
		// 					EncryptionSettings: []*armcompute.EncryptionSettingsElement{
		// 						{
		// 							DiskEncryptionKey: &armcompute.KeyVaultAndSecretReference{
		// 								SecretURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/secrets/{secret}"),
		// 								SourceVault: &armcompute.SourceVault{
		// 									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
		// 								},
		// 							},
		// 							KeyEncryptionKey: &armcompute.KeyVaultAndKeyReference{
		// 								KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
		// 								SourceVault: &armcompute.SourceVault{
		// 									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
		// 								},
		// 							},
		// 					}},
		// 				},
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-28T04:41:35.9278721+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myManagedDisk2"),
		// 			Type: to.Ptr("Microsoft.Compute/disks"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk2"),
		// 			Location: to.Ptr("westus"),
		// 			Properties: &armcompute.DiskProperties{
		// 				CreationData: &armcompute.CreationData{
		// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionEmpty),
		// 				},
		// 				DiskSizeGB: to.Ptr[int32](10),
		// 				Encryption: &armcompute.Encryption{
		// 					Type: to.Ptr(armcompute.EncryptionTypeEncryptionAtRestWithPlatformKey),
		// 				},
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-28T04:41:36.872242+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myManagedDisk3"),
		// 			Type: to.Ptr("Microsoft.Compute/disks"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk3"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"department": to.Ptr("Development"),
		// 				"project": to.Ptr("ManagedDisks"),
		// 			},
		// 			Properties: &armcompute.DiskProperties{
		// 				CreationData: &armcompute.CreationData{
		// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionFromImage),
		// 					ImageReference: &armcompute.ImageDiskReference{
		// 						ID: to.Ptr("/Subscriptions/{subscriptionId}/Providers/Microsoft.Compute/Locations/uswest/Publishers/Microsoft/ArtifactTypes/VMImage/Offers/{offer}"),
		// 					},
		// 				},
		// 				DiskSizeGB: to.Ptr[int32](200),
		// 				Encryption: &armcompute.Encryption{
		// 					Type: to.Ptr(armcompute.EncryptionTypeEncryptionAtRestWithPlatformKey),
		// 				},
		// 				EncryptionSettingsCollection: &armcompute.EncryptionSettingsCollection{
		// 					Enabled: to.Ptr(true),
		// 					EncryptionSettings: []*armcompute.EncryptionSettingsElement{
		// 						{
		// 							DiskEncryptionKey: &armcompute.KeyVaultAndSecretReference{
		// 								SecretURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/secrets/{secret}"),
		// 								SourceVault: &armcompute.SourceVault{
		// 									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
		// 								},
		// 							},
		// 							KeyEncryptionKey: &armcompute.KeyVaultAndKeyReference{
		// 								KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
		// 								SourceVault: &armcompute.SourceVault{
		// 									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
		// 								},
		// 							},
		// 					}},
		// 				},
		// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-28T04:41:36.3973934+00:00"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
