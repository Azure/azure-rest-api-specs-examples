package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2025-08-01/StorageConnectorCRUD/StorageConnectors_ListByStorageAccount.json
func ExampleConnectorsClient_NewListByStorageAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewConnectorsClient().NewListByStorageAccountPager("testrg", "teststorageaccount", nil)
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
		// page = armstorage.ConnectorsClientListByStorageAccountResponse{
		// 	ConnectorListResult: armstorage.ConnectorListResult{
		// 		Value: []*armstorage.Connector{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/teststorageaccount/connectors/testconnector"),
		// 				Name: to.Ptr("testconnector1"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/connectors"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armstorage.ConnectorProperties{
		// 					UniqueID: to.Ptr("12345678-1234-1234-1234-12345678912"),
		// 					State: to.Ptr(armstorage.StorageConnectorStateActive),
		// 					CreationTime: to.Ptr("2023-04-01T00:00:00Z"),
		// 					Description: to.Ptr("Example connector"),
		// 					DataSourceType: to.Ptr(armstorage.StorageConnectorDataSourceTypeAzureDataShare),
		// 					Source: &armstorage.DataShareSource{
		// 						Type: to.Ptr(armstorage.StorageConnectorSourceTypeDataShare),
		// 						Connection: &armstorage.DataShareConnection{
		// 							Type: to.Ptr(armstorage.StorageConnectorConnectionTypeDataShare),
		// 							DataShareURI: to.Ptr("azds://eastus:datashare1:12345678-1234-1234-1234-123456789123"),
		// 						},
		// 						AuthProperties: &armstorage.ManagedIdentityAuthProperties{
		// 							Type: to.Ptr(armstorage.StorageConnectorAuthTypeManagedIdentity),
		// 							IdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testIdentity"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armstorage.NativeDataSharingProvisioningStateSucceeded),
		// 				},
		// 			},
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/teststorageaccount/connectors/Jill"),
		// 				Name: to.Ptr("testconnector2"),
		// 				Type: to.Ptr("Microsoft.Storage/storageAccounts/connectors"),
		// 				Location: to.Ptr("westus"),
		// 				Properties: &armstorage.ConnectorProperties{
		// 					UniqueID: to.Ptr("12345678-1234-1234-1234-12345678912"),
		// 					State: to.Ptr(armstorage.StorageConnectorStateActive),
		// 					CreationTime: to.Ptr("2023-04-01T00:00:00Z"),
		// 					Description: to.Ptr("Example connector"),
		// 					DataSourceType: to.Ptr(armstorage.StorageConnectorDataSourceTypeAzureDataShare),
		// 					Source: &armstorage.DataShareSource{
		// 						Type: to.Ptr(armstorage.StorageConnectorSourceTypeDataShare),
		// 						Connection: &armstorage.DataShareConnection{
		// 							Type: to.Ptr(armstorage.StorageConnectorConnectionTypeDataShare),
		// 							DataShareURI: to.Ptr("azds://eastus:datashare1:12345678-1234-1234-1234-123456789123"),
		// 						},
		// 						AuthProperties: &armstorage.ManagedIdentityAuthProperties{
		// 							Type: to.Ptr(armstorage.StorageConnectorAuthTypeManagedIdentity),
		// 							IdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testIdentity"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armstorage.NativeDataSharingProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
