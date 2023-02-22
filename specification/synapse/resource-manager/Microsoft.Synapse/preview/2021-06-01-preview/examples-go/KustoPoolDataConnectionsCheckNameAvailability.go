package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsCheckNameAvailability.json
func ExampleKustoPoolDataConnectionsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsynapse.NewKustoPoolDataConnectionsClient("12345678-1234-1234-1234-123456789098", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CheckNameAvailability(ctx, "kustorptest", "synapseWorkspaceName", "kustoclusterrptest4", "Kustodatabase8", armsynapse.DataConnectionCheckNameRequest{
		Name: to.Ptr("DataConnections8"),
		Type: to.Ptr("Microsoft.Synapse/workspaces/kustoPools/databases/dataConnections"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameResult = armsynapse.CheckNameResult{
	// 	Name: to.Ptr("DataConnections8"),
	// 	Message: to.Ptr("Name 'DataConnections8' is already taken. Please specify a different name."),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr(armsynapse.ReasonAlreadyExists),
	// }
}
