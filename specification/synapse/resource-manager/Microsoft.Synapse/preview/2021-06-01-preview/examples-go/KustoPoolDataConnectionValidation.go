package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionValidation.json
func ExampleKustoPoolDataConnectionsClient_BeginDataConnectionValidation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewKustoPoolDataConnectionsClient().BeginDataConnectionValidation(ctx, "kustorptest", "kustorptest", "kustoclusterrptest4", "KustoDatabase8", armsynapse.DataConnectionValidation{
		DataConnectionName: to.Ptr("DataConnections8"),
		Properties: &armsynapse.EventHubDataConnection{
			Kind: to.Ptr(armsynapse.DataConnectionKindEventHub),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataConnectionValidationListResult = armsynapse.DataConnectionValidationListResult{
	// 	Value: []*armsynapse.DataConnectionValidationResult{
	// 		{
	// 			ErrorMessage: to.Ptr("Event hub's namespace does not exist"),
	// 		},
	// 		{
	// 			ErrorMessage: to.Ptr("Database does not exist"),
	// 	}},
	// }
}
