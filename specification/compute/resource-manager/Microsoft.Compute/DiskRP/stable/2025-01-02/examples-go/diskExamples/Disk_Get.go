package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskExamples/Disk_Get.json
func ExampleDisksClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDisksClient().Get(ctx, "myResourceGroup", "myManagedDisk", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Disk = armcompute.Disk{
	// 	Name: to.Ptr("myManagedDisk"),
	// 	Type: to.Ptr("Microsoft.Compute/disks"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("Development"),
	// 		"project": to.Ptr("ManagedDisks"),
	// 	},
	// 	ManagedBy: to.Ptr("/subscriptions/123caaa-123v-v211-a49f-f88ccac5bf88/resourceGroups/ResourceGroupName/providers/Microsoft.Compute/virtualMachines/TestVM414689371c88843d65ec"),
	// 	Properties: &armcompute.DiskProperties{
	// 		LastOwnershipUpdateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-28T04:41:35.079Z"); return t}()),
	// 		CreationData: &armcompute.CreationData{
	// 			CreateOption: to.Ptr(armcompute.DiskCreateOptionFromImage),
	// 			ImageReference: &armcompute.ImageDiskReference{
	// 				ID: to.Ptr("/Subscriptions/{subscription-id}/Providers/Microsoft.Compute/Locations/westus/Publishers/test_test_pmc2pc1/ArtifactTypes/VMImage/Offers/marketplace_vm_test/Skus/test_sku/Versions/1.0.0"),
	// 			},
	// 		},
	// 		DiskSizeGB: to.Ptr[int32](10),
	// 		Encryption: &armcompute.Encryption{
	// 			Type: to.Ptr(armcompute.EncryptionTypeEncryptionAtRestWithPlatformKey),
	// 		},
	// 		EncryptionSettingsCollection: &armcompute.EncryptionSettingsCollection{
	// 			Enabled: to.Ptr(true),
	// 			EncryptionSettings: []*armcompute.EncryptionSettingsElement{
	// 				{
	// 					DiskEncryptionKey: &armcompute.KeyVaultAndSecretReference{
	// 						SecretURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/secrets/{secret}"),
	// 						SourceVault: &armcompute.SourceVault{
	// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
	// 						},
	// 					},
	// 					KeyEncryptionKey: &armcompute.KeyVaultAndKeyReference{
	// 						KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
	// 						SourceVault: &armcompute.SourceVault{
	// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
	// 						},
	// 					},
	// 			}},
	// 		},
	// 		HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
	// 		OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		PurchasePlan: &armcompute.DiskPurchasePlan{
	// 			Name: to.Ptr("test_sku"),
	// 			Product: to.Ptr("marketplace_vm_test"),
	// 			Publisher: to.Ptr("test_test_pmc2pc1"),
	// 		},
	// 		SecurityProfile: &armcompute.DiskSecurityProfile{
	// 			SecurityType: to.Ptr(armcompute.DiskSecurityTypesTrustedLaunch),
	// 		},
	// 		SupportedCapabilities: &armcompute.SupportedCapabilities{
	// 			AcceleratedNetwork: to.Ptr(true),
	// 		},
	// 		SupportsHibernation: to.Ptr(true),
	// 		TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-28T04:41:35.079Z"); return t}()),
	// 	},
	// 	SKU: &armcompute.DiskSKU{
	// 		Name: to.Ptr(armcompute.DiskStorageAccountTypesStandardLRS),
	// 	},
	// }
}
