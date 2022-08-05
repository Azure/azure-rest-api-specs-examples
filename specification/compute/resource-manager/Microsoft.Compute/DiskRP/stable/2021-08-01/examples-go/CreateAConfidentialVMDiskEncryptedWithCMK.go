package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-08-01/examples/CreateAConfidentialVMDiskEncryptedWithCMK.json
func ExampleDisksClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewDisksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<disk-name>",
		armcompute.Disk{
			Location: to.StringPtr("<location>"),
			Properties: &armcompute.DiskProperties{
				CreationData: &armcompute.CreationData{
					CreateOption: armcompute.DiskCreateOption("FromImage").ToPtr(),
					ImageReference: &armcompute.ImageDiskReference{
						ID: to.StringPtr("<id>"),
					},
				},
				OSType: armcompute.OperatingSystemTypesWindows.ToPtr(),
				SecurityProfile: &armcompute.DiskSecurityProfile{
					SecureVMDiskEncryptionSetID: to.StringPtr("<secure-vmdisk-encryption-set-id>"),
					SecurityType:                armcompute.DiskSecurityTypes("ConfidentialVM_DiskEncryptedWithCustomerKey").ToPtr(),
				},
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
	log.Printf("Response result: %#v\n", res.DisksClientCreateOrUpdateResult)
}
