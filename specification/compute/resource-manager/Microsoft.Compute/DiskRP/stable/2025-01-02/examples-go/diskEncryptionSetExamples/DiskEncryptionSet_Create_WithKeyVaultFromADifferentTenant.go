package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Create_WithKeyVaultFromADifferentTenant.json
func ExampleDiskEncryptionSetsClient_BeginCreateOrUpdate_createADiskEncryptionSetWithKeyVaultFromADifferentTenant() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskEncryptionSetsClient().BeginCreateOrUpdate(ctx, "myResourceGroup", "myDiskEncryptionSet", armcompute.DiskEncryptionSet{
		Location: to.Ptr("West US"),
		Identity: &armcompute.EncryptionSetIdentity{
			Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
				"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}": {},
			},
		},
		Properties: &armcompute.EncryptionSetProperties{
			ActiveKey: &armcompute.KeyForDiskEncryptionSet{
				KeyURL: to.Ptr("https://myvaultdifferenttenant.vault-int.azure-int.net/keys/{key}"),
			},
			EncryptionType:    to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
			FederatedClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiskEncryptionSet = armcompute.DiskEncryptionSet{
	// 	Name: to.Ptr("myDiskEncryptionSet"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
	// 	Location: to.Ptr("West US"),
	// 	Identity: &armcompute.EncryptionSetIdentity{
	// 		Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
	// 			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}": &armcompute.UserAssignedIdentitiesValue{
	// 			},
	// 		},
	// 	},
	// 	Properties: &armcompute.EncryptionSetProperties{
	// 		ActiveKey: &armcompute.KeyForDiskEncryptionSet{
	// 			KeyURL: to.Ptr("https://myvaultdifferenttenant.vault-int.azure-int.net/keys/{key}"),
	// 		},
	// 		EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
	// 		FederatedClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
	// 		},
	// 	},
	// }
}
