package armbatch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/batch/resource-manager/Microsoft.Batch/stable/2022-10-01/examples/BatchAccountCreate_BYOS.json
func ExampleAccountClient_BeginCreate_batchAccountCreateByos() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armbatch.NewAccountClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx, "default-azurebatch-japaneast", "sampleacct", armbatch.AccountCreateParameters{
		Location: to.Ptr("japaneast"),
		Properties: &armbatch.AccountCreateProperties{
			AutoStorage: &armbatch.AutoStorageBaseProperties{
				StorageAccountID: to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Storage/storageAccounts/samplestorage"),
			},
			KeyVaultReference: &armbatch.KeyVaultReference{
				ID:  to.Ptr("/subscriptions/subid/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.KeyVault/vaults/sample"),
				URL: to.Ptr("http://sample.vault.azure.net/"),
			},
			PoolAllocationMode: to.Ptr(armbatch.PoolAllocationModeUserSubscription),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
