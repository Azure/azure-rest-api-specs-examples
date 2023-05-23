package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5d2adf9b7fda669b4a2538c65e937ee74fe3f966/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/snapshotExamples/Snapshot_ListByResourceGroup.json
func ExampleSnapshotsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSnapshotsClient().NewListByResourceGroupPager("myResourceGroup", nil)
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
		// page.SnapshotList = armcompute.SnapshotList{
		// 	Value: []*armcompute.Snapshot{
		// 		{
		// 			Name: to.Ptr("mySnapshot"),
		// 			Type: to.Ptr("Microsoft.Compute/snapshots"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"department": to.Ptr("Development"),
		// 				"project": to.Ptr("Snapshots"),
		// 			},
		// 			Properties: &armcompute.SnapshotProperties{
		// 				CreationData: &armcompute.CreationData{
		// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionCopy),
		// 					SourceResourceID: to.Ptr("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
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
		// 	}},
		// }
	}
}
