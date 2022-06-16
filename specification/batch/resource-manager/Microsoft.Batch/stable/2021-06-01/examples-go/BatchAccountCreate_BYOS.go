package armbatch_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/BatchAccountCreate_BYOS.json
func ExampleAccountClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewAccountClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armbatch.AccountCreateParameters{
			Location: to.StringPtr("<location>"),
			Properties: &armbatch.AccountCreateProperties{
				AutoStorage: &armbatch.AutoStorageBaseProperties{
					StorageAccountID: to.StringPtr("<storage-account-id>"),
				},
				KeyVaultReference: &armbatch.KeyVaultReference{
					ID:  to.StringPtr("<id>"),
					URL: to.StringPtr("<url>"),
				},
				PoolAllocationMode: armbatch.PoolAllocationModeUserSubscription.ToPtr(),
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
	log.Printf("Response result: %#v\n", res.AccountClientCreateResult)
}
