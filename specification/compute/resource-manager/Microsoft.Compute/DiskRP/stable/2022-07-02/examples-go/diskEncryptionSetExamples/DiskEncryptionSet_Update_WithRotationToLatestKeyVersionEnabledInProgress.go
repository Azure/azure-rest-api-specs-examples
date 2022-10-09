package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Update_WithRotationToLatestKeyVersionEnabledInProgress.json
func ExampleDiskEncryptionSetsClient_BeginUpdate_updateADiskEncryptionSetWithRotationToLatestKeyVersionEnabledSetToTrueUpdating() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewDiskEncryptionSetsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "myResourceGroup", "myDiskEncryptionSet", armcompute.DiskEncryptionSetUpdate{
		Identity: &armcompute.EncryptionSetIdentity{
			Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
		},
		Properties: &armcompute.DiskEncryptionSetUpdateProperties{
			ActiveKey: &armcompute.KeyForDiskEncryptionSet{
				KeyURL: to.Ptr("https://myvaultdifferentsub.vault-int.azure-int.net/keys/keyName/keyVersion1"),
			},
			EncryptionType:                    to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
			RotationToLatestKeyVersionEnabled: to.Ptr(true),
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
