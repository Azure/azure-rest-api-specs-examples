package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountGetAsyncSkuConversionStatus.json
func ExampleAccountsClient_GetProperties_storageAccountGetAsyncSkuConversionStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().GetProperties(ctx, "res9407", "sto8596", &armstorage.AccountsClientGetPropertiesOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Account = armstorage.Account{
	// 	Name: to.Ptr("sto8596"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/res9407/providers/Microsoft.Storage/storageAccounts/sto8596"),
	// 	Location: to.Ptr("eastus"),
	// 	Kind: to.Ptr(armstorage.KindStorageV2),
	// 	Properties: &armstorage.AccountProperties{
	// 		AllowBlobPublicAccess: to.Ptr(false),
	// 		MinimumTLSVersion: to.Ptr(armstorage.MinimumTLSVersionTLS12),
	// 		StorageAccountSKUConversionStatus: &armstorage.AccountSKUConversionStatus{
	// 			EndTime: to.Ptr("2021-09-02T02:53:39.0932539Z"),
	// 			SKUConversionStatus: to.Ptr(armstorage.SKUConversionStatusInProgress),
	// 			StartTime: to.Ptr("2022-09-01T02:53:39.0932539Z"),
	// 			TargetSKUName: to.Ptr(armstorage.SKUNameStandardGRS),
	// 		},
	// 	},
	// 	SKU: &armstorage.SKU{
	// 		Name: to.Ptr(armstorage.SKUNameStandardGRS),
	// 		Tier: to.Ptr(armstorage.SKUTierStandard),
	// 	},
	// }
}
