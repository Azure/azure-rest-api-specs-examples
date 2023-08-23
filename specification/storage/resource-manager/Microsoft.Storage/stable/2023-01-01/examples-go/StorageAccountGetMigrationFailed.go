package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/StorageAccountGetMigrationFailed.json
func ExampleAccountsClient_GetCustomerInitiatedMigration_storageAccountGetMigrationFailed() {
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
	// 		MigrationFailedDetailedReason: to.Ptr("ZRS is not supported for accounts with archive data."),
	// 		MigrationFailedReason: to.Ptr("ZrsNotSupportedForAccountWithArchiveData"),
	// 		MigrationStatus: to.Ptr(armstorage.MigrationStatusFailed),
	// 		TargetSKUName: to.Ptr(armstorage.SKUNameStandardZRS),
	// 	},
	// }
}
