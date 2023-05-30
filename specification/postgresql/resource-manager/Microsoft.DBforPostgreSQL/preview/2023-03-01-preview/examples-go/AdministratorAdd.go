package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/66a59c94238bf973680355fb179fade4c9405710/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-03-01-preview/examples/AdministratorAdd.json
func ExampleAdministratorsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAdministratorsClient().BeginCreate(ctx, "testrg", "testserver", "oooooooo-oooo-oooo-oooo-oooooooooooo", armpostgresqlflexibleservers.ActiveDirectoryAdministratorAdd{
		Properties: &armpostgresqlflexibleservers.AdministratorPropertiesForAdd{
			PrincipalName: to.Ptr("testuser1@microsoft.com"),
			PrincipalType: to.Ptr(armpostgresqlflexibleservers.PrincipalTypeUser),
			TenantID:      to.Ptr("tttttttt-tttt-tttt-tttt-tttttttttttt"),
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
	// res.ActiveDirectoryAdministrator = armpostgresqlflexibleservers.ActiveDirectoryAdministrator{
	// 	Name: to.Ptr("testuser1@microsoft.com"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/administrators"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/flexibleServers/testserver/administrators/oooooooo-oooo-oooo-oooo-oooooooooooo"),
	// 	Properties: &armpostgresqlflexibleservers.AdministratorProperties{
	// 		ObjectID: to.Ptr("oooooooo-oooo-oooo-oooo-oooooooooooo"),
	// 		PrincipalName: to.Ptr("testuser1@microsoft.com"),
	// 		PrincipalType: to.Ptr(armpostgresqlflexibleservers.PrincipalTypeUser),
	// 	},
	// }
}
