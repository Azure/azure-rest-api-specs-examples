package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-08-01/examples/CreateADiskEncryptionSetWithKeyVaultFromADifferentSubscription.json
func ExampleDiskEncryptionSetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewDiskEncryptionSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<disk-encryption-set-name>",
		armcompute.DiskEncryptionSet{
			Location: to.StringPtr("<location>"),
			Identity: &armcompute.EncryptionSetIdentity{
				Type: armcompute.DiskEncryptionSetIdentityType("SystemAssigned").ToPtr(),
			},
			Properties: &armcompute.EncryptionSetProperties{
				ActiveKey: &armcompute.KeyForDiskEncryptionSet{
					KeyURL: to.StringPtr("<key-url>"),
				},
				EncryptionType: armcompute.DiskEncryptionSetType("EncryptionAtRestWithCustomerKey").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.DiskEncryptionSetsClientCreateOrUpdateResult)
}
