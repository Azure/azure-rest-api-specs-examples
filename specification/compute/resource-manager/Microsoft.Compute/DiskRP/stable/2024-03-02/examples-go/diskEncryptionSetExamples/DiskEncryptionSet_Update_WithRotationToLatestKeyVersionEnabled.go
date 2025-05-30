package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/250861bb6a886b75255edfa0aa5ee2dd0d6e7a11/specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_Update_WithRotationToLatestKeyVersionEnabled.json
func ExampleDiskEncryptionSetsClient_BeginUpdate_updateADiskEncryptionSetWithRotationToLatestKeyVersionEnabledSetToTrueSucceeded() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewDiskEncryptionSetsClient().BeginUpdate(ctx, "myResourceGroup", "myDiskEncryptionSet", armcompute.DiskEncryptionSetUpdate{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DiskEncryptionSet = armcompute.DiskEncryptionSet{
	// 	Name: to.Ptr("myDiskEncryptionSet"),
	// 	Type: to.Ptr("Microsoft.Compute/diskEncryptionSets"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"),
	// 	Location: to.Ptr("West US"),
	// 	Identity: &armcompute.EncryptionSetIdentity{
	// 		Type: to.Ptr(armcompute.DiskEncryptionSetIdentityTypeSystemAssigned),
	// 	},
	// 	Properties: &armcompute.EncryptionSetProperties{
	// 		ActiveKey: &armcompute.KeyForDiskEncryptionSet{
	// 			KeyURL: to.Ptr("https://myvaultdifferentsub.vault-int.azure-int.net/keys/keyName/KeyVersion2"),
	// 		},
	// 		EncryptionType: to.Ptr(armcompute.DiskEncryptionSetTypeEncryptionAtRestWithCustomerKey),
	// 		LastKeyRotationTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-01T04:41:35.079Z"); return t}()),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RotationToLatestKeyVersionEnabled: to.Ptr(true),
	// 	},
	// }
}
