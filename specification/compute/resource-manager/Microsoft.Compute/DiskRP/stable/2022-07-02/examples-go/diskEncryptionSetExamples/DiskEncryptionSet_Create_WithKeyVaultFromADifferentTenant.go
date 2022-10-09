package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Create_WithKeyVaultFromADifferentTenant.json
func ExampleDiskEncryptionSetsClient_BeginCreateOrUpdate_createADiskEncryptionSetWithKeyVaultFromADifferentTenant() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewDiskEncryptionSetsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "myResourceGroup", "myDiskEncryptionSet", armcompute.DiskEncryptionSet{
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
	// TODO: use response item
	_ = res
}
