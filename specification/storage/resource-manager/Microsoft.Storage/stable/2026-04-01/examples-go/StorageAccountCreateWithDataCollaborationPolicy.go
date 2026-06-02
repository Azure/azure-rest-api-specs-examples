package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-04-01/StorageAccountCreateWithDataCollaborationPolicy.json
func ExampleAccountsClient_BeginCreate_storageAccountCreateWithDataCollaborationPolicy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCreate(ctx, "res9101", "sto4445", armstorage.AccountCreateParameters{
		Kind:     to.Ptr(armstorage.KindStorage),
		Location: to.Ptr("eastus"),
		Properties: &armstorage.AccountPropertiesCreateParameters{
			DataCollaborationPolicyProperties: &armstorage.DataCollaborationPolicyProperties{
				AllowStorageConnectors:      to.Ptr(true),
				AllowStorageDataShares:      to.Ptr(true),
				AllowCrossTenantDataSharing: to.Ptr(false),
			},
		},
		SKU: &armstorage.SKU{
			Name: to.Ptr(armstorage.SKUNameStandardGRS),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstorage.AccountsClientCreateResponse{
	// 	Account: armstorage.Account{
	// 		Name: to.Ptr("sto4445"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9101/providers/Microsoft.Storage/storageAccounts/sto4445"),
	// 		Kind: to.Ptr(armstorage.KindStorage),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armstorage.AccountProperties{
	// 			CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T13:25:33.4863236Z"); return t}()),
	// 			PrimaryEndpoints: &armstorage.Endpoints{
	// 				Blob: to.Ptr("https://sto4445.blob.core.windows.net/"),
	// 				Dfs: to.Ptr("https://sto4445.dfs.core.windows.net/"),
	// 				File: to.Ptr("https://sto4445.file.core.windows.net/"),
	// 				Queue: to.Ptr("https://sto4445.queue.core.windows.net/"),
	// 				Table: to.Ptr("https://sto4445.table.core.windows.net/"),
	// 				Web: to.Ptr("https://sto4445.web.core.windows.net/"),
	// 			},
	// 			PrimaryLocation: to.Ptr("eastus2euap"),
	// 			ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
	// 			DataCollaborationPolicyProperties: &armstorage.DataCollaborationPolicyProperties{
	// 				AllowStorageConnectors: to.Ptr(true),
	// 				AllowStorageDataShares: to.Ptr(true),
	// 				AllowCrossTenantDataSharing: to.Ptr(false),
	// 			},
	// 		},
	// 		SKU: &armstorage.SKU{
	// 			Name: to.Ptr(armstorage.SKUNameStandardGRS),
	// 			Tier: to.Ptr(armstorage.SKUTierStandard),
	// 		},
	// 	},
	// }
}
