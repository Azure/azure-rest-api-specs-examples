package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bf204aab860f2eb58a9d346b00d44760f2a9b0a2/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2023-06-01-preview/examples/DatabasesListByServer.json
func ExampleDatabasesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDatabasesClient().NewListByServerPager("TestGroup", "testserver", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.DatabaseListResult = armpostgresqlflexibleservers.DatabaseListResult{
		// 	Value: []*armpostgresqlflexibleservers.Database{
		// 		{
		// 			Name: to.Ptr("db1"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/databases"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/testserver/databases/db1"),
		// 			Properties: &armpostgresqlflexibleservers.DatabaseProperties{
		// 				Charset: to.Ptr("utf8"),
		// 				Collation: to.Ptr("en_US.utf8"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("db2"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/flexibleServers/databases"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforPostgreSQL/flexibleServers/testserver/databases/db2"),
		// 			Properties: &armpostgresqlflexibleservers.DatabaseProperties{
		// 				Charset: to.Ptr("utf8"),
		// 				Collation: to.Ptr("en_US.utf8"),
		// 			},
		// 	}},
		// }
	}
}
