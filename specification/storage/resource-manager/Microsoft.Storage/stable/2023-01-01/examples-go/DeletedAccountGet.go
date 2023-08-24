package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0baf811c3c76c87b3c127d098519bd97141222dd/specification/storage/resource-manager/Microsoft.Storage/stable/2023-01-01/examples/DeletedAccountGet.json
func ExampleDeletedAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeletedAccountsClient().Get(ctx, "sto1125", "eastus", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DeletedAccount = armstorage.DeletedAccount{
	// 	Name: to.Ptr("sto1125"),
	// 	Type: to.Ptr("Microsoft.Storage/deletedAccounts"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.Storage/locations/eastus/deletedAccounts/sto1125"),
	// 	Properties: &armstorage.DeletedAccountProperties{
	// 		CreationTime: to.Ptr("2020-08-17T03:35:37.4588848Z"),
	// 		DeletionTime: to.Ptr("2020-08-17T04:41:37.3442475Z"),
	// 		Location: to.Ptr("eastus"),
	// 		RestoreReference: to.Ptr("sto1125|2020-08-17T03:35:37.4588848Z"),
	// 		StorageAccountResourceID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/sto/providers/Microsoft.Storage/storageAccounts/sto1125"),
	// 	},
	// }
}
