package armpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2018-06-01/examples/PrivateLinkResourcesGet.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.PrivateLinkResource = armpostgresql.PrivateLinkResource{
	// 	Name: to.Ptr("plr"),
	// 	Type: to.Ptr("Microsoft.DBforPostgreSQL/servers/privateLinkResources"),
	// 	ID: to.Ptr("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.DBforPostgreSQL/servers/test-svr/privateLinkResources/plr"),
	// 	Properties: &armpostgresql.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("postgresqlServer"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("postgresqlServer")},
	// 		},
	// 	}
}
