package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ecee919199a39cc0d864410f540aa105bf7cdb64/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2024-08-01/examples/PrivateLinkResourcesGet.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "Default", "test-svr", "plr", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResource = armpostgresqlflexibleservers.PrivateLinkResource{
	// 	Name: to.Ptr("plr"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/privateLinkResources"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default/providers/Microsoft.DBforPostgreSQL/flexibleServers/test-svr/privateLinkResources/plr"),
	// 	Properties: &armpostgresqlflexibleservers.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("postgresqlServer"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("postgresqlServer")},
	// 			RequiredZoneNames: []*string{
	// 				to.Ptr("privatelink.meru.com")},
	// 			},
	// 		}
}
