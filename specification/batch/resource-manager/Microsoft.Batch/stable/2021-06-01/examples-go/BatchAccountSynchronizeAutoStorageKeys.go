package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/BatchAccountSynchronizeAutoStorageKeys.json
func ExampleAccountClient_SynchronizeAutoStorageKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewAccountClient("<subscription-id>", cred, nil)
	_, err = client.SynchronizeAutoStorageKeys(ctx,
		"<resource-group-name>",
		"<account-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
