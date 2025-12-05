package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e96c24570a484cff13d153fb472f812878866a39/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2025-08-01/examples/AdministratorsMicrosoftEntraGet.json
func ExampleAdministratorsMicrosoftEntraClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAdministratorsMicrosoftEntraClient().Get(ctx, "exampleresourcegroup", "exampleserver", "oooooooo-oooo-oooo-oooo-oooooooooooo", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AdministratorMicrosoftEntra = armpostgresqlflexibleservers.AdministratorMicrosoftEntra{
	// 	Name: to.Ptr("exampleuser@contoso.com"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/administrators"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/exampleresourcegroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/exampleserver/administrators/oooooooo-oooo-oooo-oooo-oooooooooooo"),
	// 	Properties: &armpostgresqlflexibleservers.AdministratorMicrosoftEntraProperties{
	// 		ObjectID: to.Ptr("oooooooo-oooo-oooo-oooo-oooooooooooo"),
	// 		PrincipalName: to.Ptr("exampleuser@contoso.com"),
	// 		PrincipalType: to.Ptr(armpostgresqlflexibleservers.PrincipalTypeUser),
	// 		TenantID: to.Ptr("tttttttt-tttt-tttt-tttt-tttttttttttt"),
	// 	},
	// }
}
