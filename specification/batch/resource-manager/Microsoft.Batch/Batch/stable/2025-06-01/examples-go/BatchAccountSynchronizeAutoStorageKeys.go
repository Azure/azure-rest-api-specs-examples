package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v4"
)

// Generated from example definition: 2025-06-01/BatchAccountSynchronizeAutoStorageKeys.json
func ExampleAccountClient_SynchronizeAutoStorageKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armbatch.NewClientFactory("12345678-1234-1234-1234-123456789012", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewAccountClient().SynchronizeAutoStorageKeys(ctx, "default-azurebatch-japaneast", "sampleacct", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
