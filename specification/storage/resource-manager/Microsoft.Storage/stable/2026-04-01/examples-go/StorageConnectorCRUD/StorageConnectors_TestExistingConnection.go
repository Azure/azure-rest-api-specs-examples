package armstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage/v4"
)

// Generated from example definition: 2026-04-01/StorageConnectorCRUD/StorageConnectors_TestExistingConnection.json
func ExampleConnectorsClient_BeginTestExistingConnection() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorage.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewConnectorsClient().BeginTestExistingConnection(ctx, "testrg", "teststorageaccount", "testconnector", armstorage.TestExistingConnectionRequest{
		UniqueID: to.Ptr("12345678-1234-1234-1234-12345678912"),
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
	// res = armstorage.ConnectorsClientTestExistingConnectionResponse{
	// 	TestConnectionResponse: armstorage.TestConnectionResponse{
	// 		StorageConnectorMethodName: to.Ptr("TestExistingConnection"),
	// 		StorageConnectorRequestID: to.Ptr("request-id-123"),
	// 	},
	// }
}
