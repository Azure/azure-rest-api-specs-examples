package armmysql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/PrivateLinkResourcesGet.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.PrivateLinkResource = armmysql.PrivateLinkResource{
	// 	Name: to.Ptr("plr"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers/privateLinkResources"),
	// 	ID: to.Ptr("subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.DBforMySQL/servers/test-svr/privateLinkResources/plr"),
	// 	Properties: &armmysql.PrivateLinkResourceProperties{
	// 		GroupID: to.Ptr("mysqlServer"),
	// 		RequiredMembers: []*string{
	// 			to.Ptr("mysqlServer")},
	// 		},
	// 	}
}
