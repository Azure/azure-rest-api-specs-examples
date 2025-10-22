package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4e9df3afd38a1cfa00a5d49419dce51bd014601f/specification/storage/resource-manager/Microsoft.Storage/stable/2025-06-01/examples/StorageAccountGetMigrationInProgress.json
func ExampleAccountsClient_GetCustomerInitiatedMigration_storageAccountGetMigrationInProgress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.AccountMigration = armstorage.AccountMigration{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Storage/storageAccounts/accountMigrations"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/resource-group-name/providers/Microsoft.Storage/storageAccounts/accountname/accountMigrations/default"),
	// 	StorageAccountMigrationDetails: &armstorage.AccountMigrationProperties{
	// 		MigrationStatus: to.Ptr(armstorage.MigrationStatusInProgress),
	// 		TargetSKUName: to.Ptr(armstorage.SKUNameStandardZRS),
	// 	},
	// }
}
