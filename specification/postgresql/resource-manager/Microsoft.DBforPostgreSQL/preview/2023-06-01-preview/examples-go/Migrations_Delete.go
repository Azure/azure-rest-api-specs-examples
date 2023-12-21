package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/90115af9fda46f323e5c42c274f2b376108d1d47/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/Migrations_Delete.json
func ExampleMigrationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewMigrationsClient().Delete(ctx, "ffffffff-ffff-ffff-ffff-ffffffffffff", "testrg", "testtarget", "testmigration", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
