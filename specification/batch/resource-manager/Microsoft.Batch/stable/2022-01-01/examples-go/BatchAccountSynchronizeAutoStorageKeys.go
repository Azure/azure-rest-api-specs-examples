package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/BatchAccountSynchronizeAutoStorageKeys.json
func ExampleAccountClient_SynchronizeAutoStorageKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewAccountClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.SynchronizeAutoStorageKeys(ctx,
		"default-azurebatch-japaneast",
		"sampleacct",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
