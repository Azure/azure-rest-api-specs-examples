package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CancelElasticPoolOperation.json
func ExampleElasticPoolOperationsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewElasticPoolOperationsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Cancel(ctx,
		"sqlcrudtest-7398",
		"sqlcrudtest-6661",
		"testpool",
		"f779414b-e748-4925-8cfe-c8598f7660ae",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
