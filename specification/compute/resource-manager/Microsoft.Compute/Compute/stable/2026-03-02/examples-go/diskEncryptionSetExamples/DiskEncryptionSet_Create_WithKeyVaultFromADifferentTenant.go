package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-02/diskEncryptionSetExamples/DiskEncryptionSet_Create_WithKeyVaultFromADifferentTenant.json
func ExampleDiskEncryptionSetsClient_BeginCreateOrUpdate_createADiskEncryptionSetWithKeyVaultFromADifferentTenant() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
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
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.DiskEncryptionSetsClientCreateOrUpdateResponse{
	// 	DiskEncryptionSet: armcompute.DiskEncryptionSet{
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
	// 		Name: to.Ptr("myDiskEncryptionSet"),
	// 		Location: to.Ptr("West US"),
	// 		Identity: &armcompute.EncryptionSetIdentity{
	// 			Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
	// 				"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}": &armcompute.UserAssignedIdentitiesValue{
	// 				},
	// 			},
	// 		},
	// 		Properties: &armcompute.EncryptionSetProperties{
	// 			ActiveKey: &armcompute.KeyForDiskEncryptionSet{
	// 				KeyURL: to.Ptr("https://myvaultdifferenttenant.vault-int.azure-int.net/keys/{key}"),
	// 			},
	// 			EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
	// 			FederatedClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			PreviousKeys: []*armcompute.KeyForDiskEncryptionSet{
	// 			},
	// 		},
	// 	},
	// }
}
