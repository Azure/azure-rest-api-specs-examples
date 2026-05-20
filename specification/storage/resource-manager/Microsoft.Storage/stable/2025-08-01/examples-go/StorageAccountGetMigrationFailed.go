package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2025-08-01/StorageAccountGetMigrationFailed.json
func ExampleAccountsClient_GetCustomerInitiatedMigration_storageAccountGetMigrationFailed() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountsClient().GetCustomerInitiatedMigration(ctx, "resource-group-name", "accountname", armstorage.MigrationNameDefault, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstorage.AccountsClientGetCustomerInitiatedMigrationResponse{
	// 	AccountMigration: armstorage.AccountMigration{
	// 		Name: to.Ptr("default"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts/accountMigrations"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resource-group-name/providers/Microsoft.Storage/storageAccounts/accountname/accountMigrations/default"),
	// 		StorageAccountMigrationDetails: &armstorage.AccountMigrationProperties{
	// 			MigrationFailedDetailedReason: to.Ptr("ZRS is not supported for accounts with archive data."),
	// 			MigrationFailedReason: to.Ptr("ZrsNotSupportedForAccountWithArchiveData"),
	// 			MigrationStatus: to.Ptr(armstorage.MigrationStatusFailed),
	// 			TargetSKUName: to.Ptr(armstorage.SKUNameStandardZRS),
	// 		},
	// 	},
	// }
}
