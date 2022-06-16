package armanalysisservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/analysisservices/resource-manager/Microsoft.AnalysisServices/stable/2017-08-01/examples/operationResults.json
func ExampleServersClient_ListOperationResults() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armanalysisservices.NewServersClient("613192d7-503f-477a-9cfe-4efc3ee2bd60", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.ListOperationResults(ctx,
		"West US",
		"00000000000000000000000000000000",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
