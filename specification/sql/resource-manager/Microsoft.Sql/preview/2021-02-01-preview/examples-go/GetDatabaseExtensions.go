package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/fe78d8f1e7bd86c778c7e1cafd52cb0e9fec67ef/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/GetDatabaseExtensions.json
func ExampleDatabaseExtensionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewDatabaseExtensionsClient().Get(ctx, "rg_a1f9d6f8-30d5-4228-9504-8a364361bca3", "srv_65858e0f-b1d1-4bdc-8351-a7da86ca4939", "11aa6c5e-58ed-4693-b303-3b8e3131deaa", "polybaseimport", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
