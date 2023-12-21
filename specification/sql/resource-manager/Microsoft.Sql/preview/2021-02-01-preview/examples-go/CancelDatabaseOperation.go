package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/CancelDatabaseOperation.json
func ExampleDatabaseOperationsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDatabaseOperationsClient().Cancel(ctx, "sqlcrudtest-7398", "sqlcrudtest-6661", "testdb", "f779414b-e748-4925-8cfe-c8598f7660ae", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
