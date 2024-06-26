package armpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2017-12-01/examples/ServerAdminGet.json
func ExampleServerAdministratorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServerAdministratorsClient().Get(ctx, "testrg", "pgtestsvc4", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ServerAdministratorResource = armpostgresql.ServerAdministratorResource{
	// 	Name: to.Ptr("activeDirectory"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/servers/administrators"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/servers/pgtestsvc4/administrators/activeDirectory"),
	// 	Properties: &armpostgresql.ServerAdministratorProperties{
	// 		AdministratorType: to.Ptr("ActiveDirectory"),
	// 		Login: to.Ptr("bob@contoso.com"),
	// 		Sid: to.Ptr("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
	// 		TenantID: to.Ptr("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
	// 	},
	// }
}
