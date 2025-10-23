package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4e9df3afd38a1cfa00a5d49419dce51bd014601f/specification/storage/resource-manager/Microsoft.Storage/stable/2025-06-01/examples/StorageAccountPostMigration.json
func ExampleAccountsClient_BeginCustomerInitiatedMigration() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAccountsClient().BeginCustomerInitiatedMigration(ctx, "resource-group-name", "accountname", armstorage.AccountMigration{
		StorageAccountMigrationDetails: &armstorage.AccountMigrationProperties{
			TargetSKUName: to.Ptr(armstorage.SKUNameStandardZRS),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
