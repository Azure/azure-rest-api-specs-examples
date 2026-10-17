package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-09-01/Operations_List.json
func ExampleOperationsClient_NewListPager_listOperations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("<subscriptionID>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page = armstorage.OperationsClientListResponse{
		// 	OperationListResult: armstorage.OperationListResult{
		// 		Value: []*armstorage.Operation{
		// 			{
		// 				Name: to.Ptr("Microsoft.Storage/storageAccounts/connectors/read"),
		// 				Display: &armstorage.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Storage"),
		// 					Resource: to.Ptr("Storage Connector"),
		// 					Operation: to.Ptr("Read Storage Connector"),
		// 					Description: to.Ptr("Reads the properties of a storage connector."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Storage/storageAccounts/connectors/write"),
		// 				Display: &armstorage.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Storage"),
		// 					Resource: to.Ptr("Storage Connector"),
		// 					Operation: to.Ptr("Write Storage Connector"),
		// 					Description: to.Ptr("Creates or updates a storage connector."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Storage/storageAccounts/connectors/delete"),
		// 				Display: &armstorage.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Storage"),
		// 					Resource: to.Ptr("Storage Connector"),
		// 					Operation: to.Ptr("Delete Storage Connector"),
		// 					Description: to.Ptr("Deletes a storage connector."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Storage/storageAccounts/connectors/testExistingConnection/action"),
		// 				Display: &armstorage.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Storage"),
		// 					Resource: to.Ptr("Storage Connector"),
		// 					Operation: to.Ptr("Test Existing Connection"),
		// 					Description: to.Ptr("Tests connection to an existing storage connector."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Storage/storageAccounts/dataShares/read"),
		// 				Display: &armstorage.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Storage"),
		// 					Resource: to.Ptr("Storage DataShare"),
		// 					Operation: to.Ptr("Read Storage DataShare"),
		// 					Description: to.Ptr("Reads the properties of a storage data share."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Storage/storageAccounts/dataShares/write"),
		// 				Display: &armstorage.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Storage"),
		// 					Resource: to.Ptr("Storage DataShare"),
		// 					Operation: to.Ptr("Write Storage DataShare"),
		// 					Description: to.Ptr("Reads the properties of a storage data share."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Storage/storageAccounts/dataShares/delete"),
		// 				Display: &armstorage.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Storage"),
		// 					Resource: to.Ptr("Storage DataShare"),
		// 					Operation: to.Ptr("Delete Storage DataShare"),
		// 					Description: to.Ptr("Reads the properties of a storage data share."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Storage/storageAccounts/blobAccessPointConfigurations/read"),
		// 				Display: &armstorage.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Storage"),
		// 					Resource: to.Ptr("Blob Access Point Configuration"),
		// 					Operation: to.Ptr("Read Blob Access Point Configuration"),
		// 					Description: to.Ptr("Reads Blob Access Point configuration properties."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Storage/storageAccounts/blobAccessPointConfigurations/write"),
		// 				Display: &armstorage.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Storage"),
		// 					Resource: to.Ptr("Blob Access Point Configuration"),
		// 					Operation: to.Ptr("Write Blob Access Point Configuration"),
		// 					Description: to.Ptr("Creates or updates a Blob Access Point configuration."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Storage/storageAccounts/blobAccessPointConfigurations/delete"),
		// 				Display: &armstorage.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Storage"),
		// 					Resource: to.Ptr("Blob Access Point Configuration"),
		// 					Operation: to.Ptr("Delete Blob Access Point Configuration"),
		// 					Description: to.Ptr("Deletes a Blob Access Point configuration."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Storage/storageAccounts/blobAccessPointConfigurations/testExistingConnection/action"),
		// 				Display: &armstorage.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Storage"),
		// 					Resource: to.Ptr("Blob Access Point Configuration"),
		// 					Operation: to.Ptr("Test Existing Blob Access Point Connection"),
		// 					Description: to.Ptr("Tests the connection on an existing Blob Access Point configuration."),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("Microsoft.Storage/storageAccounts/testBlobAccessPointConfigurationProposedConnection/action"),
		// 				Display: &armstorage.OperationDisplay{
		// 					Provider: to.Ptr("Microsoft Storage"),
		// 					Resource: to.Ptr("Blob Access Point Connection Test"),
		// 					Operation: to.Ptr("Test Proposed Blob Access Point Connection"),
		// 					Description: to.Ptr("Tests a proposed Blob Access Point connection before the configuration is created."),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
