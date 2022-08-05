package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-08-01/examples/BeginGetAccessDiskRestorePoint.json
func ExampleDiskRestorePointClient_BeginGrantAccess() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewDiskRestorePointClient("<subscription-id>", cred, nil)
	poller, err := client.BeginGrantAccess(ctx,
		"<resource-group-name>",
		"<restore-point-collection-name>",
		"<vm-restore-point-name>",
		"<disk-restore-point-name>",
		armcompute.GrantAccessData{
			Access:            armcompute.AccessLevel("Read").ToPtr(),
			DurationInSeconds: to.Int32Ptr(300),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DiskRestorePointClientGrantAccessResult)
}
