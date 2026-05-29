package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: 2025-02-14/KustoDataConnectionsCheckNameAvailability.json
func ExampleDataConnectionsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataConnectionsClient().CheckNameAvailability(ctx, "kustorptest", "kustoCluster", "KustoDatabase8", armkusto.DataConnectionCheckNameRequest{
		Name: to.Ptr("DataConnections8"),
		Type: to.Ptr("Microsoft.Kusto/clusters/databases/dataConnections"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkusto.DataConnectionsClientCheckNameAvailabilityResponse{
	// 	CheckNameResult: armkusto.CheckNameResult{
	// 		Name: to.Ptr("DataConnections8"),
	// 		Message: to.Ptr("Name 'DataConnections8' is already taken. Please specify a different name."),
	// 		NameAvailable: to.Ptr(false),
	// 		Reason: to.Ptr(armkusto.ReasonAlreadyExists),
	// 	},
	// }
}
