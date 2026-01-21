package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_ListBySubscription.json
func ExampleDiskEncryptionSetsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiskEncryptionSetsClient().NewListPager(nil)
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
		// page.DiskEncryptionSetList = armcompute.DiskEncryptionSetList{
		// 	Value: []*armcompute.DiskEncryptionSet{
		// 		{
		// 			Name: to.Ptr("myDiskEncryptionSet"),
		// 			Type: to.Ptr("Microsoft.Compute/diskEncryptionSets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"department": to.Ptr("Development"),
		// 				"project": to.Ptr("Encryption"),
		// 			},
		// 			Identity: &armcompute.EncryptionSetIdentity{
		// 				Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
		// 			},
		// 			Properties: &armcompute.EncryptionSetProperties{
		// 				ActiveKey: &armcompute.KeyForDiskEncryptionSet{
		// 					KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
		// 					SourceVault: &armcompute.SourceVault{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
		// 					},
		// 				},
		// 				EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
		// 				PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myDiskEncryptionSet2"),
		// 			Type: to.Ptr("Microsoft.Compute/diskEncryptionSets"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/mySecondResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet2"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 				"department": to.Ptr("Development"),
		// 				"project": to.Ptr("Encryption"),
		// 			},
		// 			Identity: &armcompute.EncryptionSetIdentity{
		// 				Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
		// 			},
		// 			Properties: &armcompute.EncryptionSetProperties{
		// 				ActiveKey: &armcompute.KeyForDiskEncryptionSet{
		// 					KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
		// 					SourceVault: &armcompute.SourceVault{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/mySecondResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault2"),
		// 					},
		// 				},
		// 				EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
		// 				PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 			},
		// 	}},
		// }
	}
}
