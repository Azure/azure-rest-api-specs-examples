package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-10-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Get.json
func ExampleDiskEncryptionSetsClient_Get_getInformationAboutADiskEncryptionSet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.DiskEncryptionSet = armcompute.DiskEncryptionSet{
	// 	Name: to.Ptr("myDiskEncryptionSet"),
	// 	Type: to.Ptr("Microsoft.Compute/diskEncryptionSets"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"department": to.Ptr("Development"),
	// 		"project": to.Ptr("Encryption"),
	// 	},
	// 	Identity: &armcompute.EncryptionSetIdentity{
	// 		Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
	// 	},
	// 	Properties: &armcompute.EncryptionSetProperties{
	// 		ActiveKey: &armcompute.KeyForDiskEncryptionSet{
	// 			KeyURL: to.Ptr("https://myvmvault.vault-int.azure-int.net/keys/{key}"),
	// 			SourceVault: &armcompute.SourceVault{
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myVMVault"),
	// 			},
	// 		},
	// 		EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
	// 		PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 	},
	// }
}
