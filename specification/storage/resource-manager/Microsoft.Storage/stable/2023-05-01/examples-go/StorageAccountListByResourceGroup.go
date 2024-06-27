package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/220ad9c6554fc7d6d10a89bdb441c1e3b36e3285/specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountListByResourceGroup.json
func ExampleAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAccountsClient().NewListByResourceGroupPager("res6117", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AccountListResult = armstorage.AccountListResult{
		// 	Value: []*armstorage.Account{
		// 		{
		// 			Name: to.Ptr("sto4036"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6117/providers/Microsoft.Storage/storageAccounts/sto4036"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Kind: to.Ptr(armstorage.KindStorage),
		// 			Properties: &armstorage.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T13:24:47.818Z"); return t}()),
		// 				IsHnsEnabled: to.Ptr(true),
		// 				PrimaryEndpoints: &armstorage.Endpoints{
		// 					Blob: to.Ptr("https://sto4036.blob.core.windows.net/"),
		// 					Dfs: to.Ptr("https://sto4036.dfs.core.windows.net/"),
		// 					File: to.Ptr("https://sto4036.file.core.windows.net/"),
		// 					Queue: to.Ptr("https://sto4036.queue.core.windows.net/"),
		// 					Table: to.Ptr("https://sto4036.table.core.windows.net/"),
		// 					Web: to.Ptr("https://sto4036.web.core.windows.net/"),
		// 				},
		// 				PrimaryLocation: to.Ptr("eastus"),
		// 				ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
		// 				SecondaryLocation: to.Ptr("centraluseuap"),
		// 				StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				StatusOfSecondary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				EnableHTTPSTrafficOnly: to.Ptr(false),
		// 			},
		// 			SKU: &armstorage.SKU{
		// 				Name: to.Ptr(armstorage.SKUNameStandardGRS),
		// 				Tier: to.Ptr(armstorage.SKUTierStandard),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sto4452"),
		// 			Type: to.Ptr("Microsoft.Storage/storageAccounts"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res6117/providers/Microsoft.Storage/storageAccounts/sto4452"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("value1"),
		// 				"key2": to.Ptr("value2"),
		// 			},
		// 			Kind: to.Ptr(armstorage.KindStorage),
		// 			Properties: &armstorage.AccountProperties{
		// 				CreationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T13:24:15.706Z"); return t}()),
		// 				PrimaryEndpoints: &armstorage.Endpoints{
		// 					Blob: to.Ptr("https://sto4452.blob.core.windows.net/"),
		// 					File: to.Ptr("https://sto4452.file.core.windows.net/"),
		// 					Queue: to.Ptr("https://sto4452.queue.core.windows.net/"),
		// 					Table: to.Ptr("https://sto4452.table.core.windows.net/"),
		// 				},
		// 				PrimaryLocation: to.Ptr("eastus"),
		// 				ProvisioningState: to.Ptr(armstorage.ProvisioningStateSucceeded),
		// 				SecondaryLocation: to.Ptr("centraluseuap"),
		// 				StatusOfPrimary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				StatusOfSecondary: to.Ptr(armstorage.AccountStatusAvailable),
		// 				EnableHTTPSTrafficOnly: to.Ptr(false),
		// 			},
		// 			SKU: &armstorage.SKU{
		// 				Name: to.Ptr(armstorage.SKUNameStandardGRS),
		// 				Tier: to.Ptr(armstorage.SKUTierStandard),
		// 			},
		// 	}},
		// }
	}
}
