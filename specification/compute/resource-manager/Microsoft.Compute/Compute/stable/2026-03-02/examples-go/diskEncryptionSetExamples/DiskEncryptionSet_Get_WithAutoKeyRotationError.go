package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-02/diskEncryptionSetExamples/DiskEncryptionSet_Get_WithAutoKeyRotationError.json
func ExampleDiskEncryptionSetsClient_Get_getInformationAboutADiskEncryptionSetWhenAutoKeyRotationFailed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDiskEncryptionSetsClient().Get(ctx, "myResourceGroup", "myDiskEncryptionSet", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.DiskEncryptionSetsClientGetResponse{
	// 	DiskEncryptionSet: armcompute.DiskEncryptionSet{
	// 		Identity: &armcompute.EncryptionSetIdentity{
	// 			Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
	// 		},
	// 		Properties: &armcompute.EncryptionSetProperties{
	// 			ActiveKey: &armcompute.KeyForDiskEncryptionSet{
	// 				SourceVault: &armcompute.SourceVault{
	// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
	// 				},
	// 				KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
	// 			},
	// 			EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
	// 			PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
	// 			},
	// 			RotationToLatestKeyVersionEnabled: to.Ptr(true),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			AutoKeyRotationError: &armcompute.APIError{
	// 				Code: to.Ptr("ManagedServiceIdentityNotFound"),
	// 				Message: to.Ptr("Auto-key rotation was disabled as managed service identity associated with DiskEncryptionSet 'myDiskEncryptionSet' was not found. Please update the resource with correct identity to re-enable auto-key rotation."),
	// 			},
	// 		},
	// 		Type: to.Ptr("Microsoft.Compute/diskEncryptionSets"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"department": to.Ptr("Development"),
	// 			"project": to.Ptr("Encryption"),
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
	// 		Name: to.Ptr("myDiskEncryptionSet"),
	// 	},
	// }
}
