package armkusto_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8adce17dc500f338f86f18af30aac61dcb71c5f/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoDatabasesCheckNameAvailability.json
func ExampleDatabasesClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkusto.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDatabasesClient().CheckNameAvailability(ctx, "kustorptest", "kustoCluster", armkusto.CheckNameRequest{
		Name: to.Ptr("database1"),
		Type: to.Ptr(armkusto.TypeMicrosoftKustoClustersDatabases),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameResult = armkusto.CheckNameResult{
	// 	Name: to.Ptr("database1"),
	// 	Message: to.Ptr("Name 'database1' is already taken. Please specify a different name"),
	// 	NameAvailable: to.Ptr(false),
	// }
}
