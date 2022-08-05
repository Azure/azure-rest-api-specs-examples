package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-08-01/examples/UpdateADiskEncryptionSetWithRotationToLatestKeyVersionEnabled.json
func ExampleDiskEncryptionSetsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewDiskEncryptionSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<disk-encryption-set-name>",
		armcompute.DiskEncryptionSetUpdate{
			Identity: &armcompute.EncryptionSetIdentity{
				Type: armcompute.DiskEncryptionSetIdentityType("SystemAssigned").ToPtr(),
			},
			Properties: &armcompute.DiskEncryptionSetUpdateProperties{
				ActiveKey: &armcompute.KeyForDiskEncryptionSet{
					KeyURL: to.StringPtr("<key-url>"),
				},
				EncryptionType:                    armcompute.DiskEncryptionSetType("EncryptionAtRestWithCustomerKey").ToPtr(),
				RotationToLatestKeyVersionEnabled: to.BoolPtr(true),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DiskEncryptionSetsClientUpdateResult)
}
