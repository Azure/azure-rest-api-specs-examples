package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-01-02/snapshotExamples/Snapshot_Get.json
func ExampleSnapshotsClient_Get_getInformationAboutASnapshot() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSnapshotsClient().Get(ctx, "myResourceGroup", "mySnapshot", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.SnapshotsClientGetResponse{
	// 	Snapshot: armcompute.Snapshot{
	// 		Properties: &armcompute.SnapshotProperties{
	// 			OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			HyperVGeneration: to.Ptr(armcompute.HyperVGenerationV1),
	// 			PurchasePlan: &armcompute.DiskPurchasePlan{
	// 				Name: to.Ptr("test_sku"),
	// 				Publisher: to.Ptr("test_test_pmc2pc1"),
	// 				Product: to.Ptr("marketplace_vm_test"),
	// 			},
	// 			SupportedCapabilities: &armcompute.SupportedCapabilities{
	// 				AcceleratedNetwork: to.Ptr(true),
	// 			},
	// 			SupportsHibernation: to.Ptr(true),
	// 			CreationData: &armcompute.CreationData{
	// 				CreateOption: to.Ptr(armcompute.DiskCreateOptionCopy),
	// 				SourceResourceID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myDisk"),
	// 				SourceUniqueID: to.Ptr("d633885d-d102-4481-901e-5b2413d1a7be"),
	// 			},
	// 			DiskSizeGB: to.Ptr[int32](100),
	// 			EncryptionSettingsCollection: &armcompute.EncryptionSettingsCollection{
	// 				Enabled: to.Ptr(true),
	// 				EncryptionSettings: []*armcompute.EncryptionSettingsElement{
	// 					{
	// 						DiskEncryptionKey: &armcompute.KeyVaultAndSecretReference{
	// 							SourceVault: &armcompute.SourceVault{
	// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
	// 							},
	// 							SecretURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/secrets/{secret}"),
	// 						},
	// 						KeyEncryptionKey: &armcompute.KeyVaultAndKeyReference{
	// 							SourceVault: &armcompute.SourceVault{
	// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
	// 							},
	// 							KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			Encryption: &armcompute.Encryption{
	// 				Type: to.Ptr(armcompute.EncryptionTypeEncryptionAtRestWithPlatformKey),
	// 			},
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2016-12-28T04:41:35.079872+00:00"); return t}()),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 		Type: to.Ptr("Microsoft.Compute/snapshots"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"department": to.Ptr("Development"),
	// 			"project": to.Ptr("Snapshots"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
	// 		Name: to.Ptr("mySnapshot"),
	// 	},
	// }
}
