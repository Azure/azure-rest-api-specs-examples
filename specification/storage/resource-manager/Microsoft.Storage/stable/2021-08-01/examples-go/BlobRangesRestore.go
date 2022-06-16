package armstorage_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-08-01/examples/BlobRangesRestore.json
func ExampleAccountsClient_BeginRestoreBlobRanges() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstorage.NewAccountsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginRestoreBlobRanges(ctx,
		"<resource-group-name>",
		"<account-name>",
		armstorage.BlobRestoreParameters{
			BlobRanges: []*armstorage.BlobRestoreRange{
				{
					EndRange:   to.StringPtr("<end-range>"),
					StartRange: to.StringPtr("<start-range>"),
				},
				{
					EndRange:   to.StringPtr("<end-range>"),
					StartRange: to.StringPtr("<start-range>"),
				}},
			TimeToRestore: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-20T15:30:00.0000000Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientRestoreBlobRangesResult)
}
