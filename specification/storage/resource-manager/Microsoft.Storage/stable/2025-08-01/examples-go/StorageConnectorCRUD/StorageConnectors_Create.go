package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2025-08-01/StorageConnectorCRUD/StorageConnectors_Create.json
func ExampleConnectorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectorsClient().BeginCreate(ctx, "testrg", "teststorageaccount", "testconnector", armstorage.Connector{
		Location: to.Ptr("eastus"),
		Properties: &armstorage.ConnectorProperties{
			State:          to.Ptr(armstorage.StorageConnectorStateActive),
			Description:    to.Ptr("Example connector"),
			DataSourceType: to.Ptr(armstorage.StorageConnectorDataSourceTypeAzureDataShare),
			Source: &armstorage.DataShareSource{
				Type: to.Ptr(armstorage.StorageConnectorSourceTypeDataShare),
				Connection: &armstorage.DataShareConnection{
					Type:         to.Ptr(armstorage.StorageConnectorConnectionTypeDataShare),
					DataShareURI: to.Ptr("azds://eastus:datashare1:12345678-1234-1234-1234-123456789123"),
				},
				AuthProperties: &armstorage.ManagedIdentityAuthProperties{
					Type:               to.Ptr(armstorage.StorageConnectorAuthTypeManagedIdentity),
					IdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testIdentity"),
				},
			},
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
	// res = armstorage.ConnectorsClientCreateResponse{
	// 	Connector: armstorage.Connector{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.Storage/storageAccounts/teststorageaccount/connectors/testconnector"),
	// 		Name: to.Ptr("testconnector"),
	// 		Type: to.Ptr("Microsoft.Storage/storageAccounts/connectors"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armstorage.ConnectorProperties{
	// 			UniqueID: to.Ptr("12345678-1234-1234-1234-12345678912"),
	// 			State: to.Ptr(armstorage.StorageConnectorStateActive),
	// 			CreationTime: to.Ptr("2023-04-01T00:00:00Z"),
	// 			Description: to.Ptr("Example connector"),
	// 			DataSourceType: to.Ptr(armstorage.StorageConnectorDataSourceTypeAzureDataShare),
	// 			Source: &armstorage.DataShareSource{
	// 				Type: to.Ptr(armstorage.StorageConnectorSourceTypeDataShare),
	// 				Connection: &armstorage.DataShareConnection{
	// 					Type: to.Ptr(armstorage.StorageConnectorConnectionTypeDataShare),
	// 					DataShareURI: to.Ptr("azds://eastus:datashare1:12345678-1234-1234-1234-123456789123"),
	// 				},
	// 				AuthProperties: &armstorage.ManagedIdentityAuthProperties{
	// 					Type: to.Ptr(armstorage.StorageConnectorAuthTypeManagedIdentity),
	// 					IdentityResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testrg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testIdentity"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armstorage.NativeDataSharingProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
