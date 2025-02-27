package armmysqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e26b89bcbec9eed5026c01416e481408b2a1ca1a/specification/mysql/resource-manager/Microsoft.DBforMySQL/Databases/stable/2023-12-30/examples/DatabasesListByServer.json
func ExampleDatabasesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysqlflexibleservers.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.DatabaseListResult = armmysqlflexibleservers.DatabaseListResult{
		// 	Value: []*armmysqlflexibleservers.Database{
		// 		{
		// 			Name: to.Ptr("db1"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/databases"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/testserver/databases/db1"),
		// 			Properties: &armmysqlflexibleservers.DatabaseProperties{
		// 				Charset: to.Ptr("utf8"),
		// 				Collation: to.Ptr("utf8_general_ci"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("db2"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/flexibleServers/databases"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DBforMySQL/flexibleServers/testserver/databases/db2"),
		// 			Properties: &armmysqlflexibleservers.DatabaseProperties{
		// 				Charset: to.Ptr("utf8"),
		// 				Collation: to.Ptr("utf8_general_ci"),
		// 			},
		// 	}},
		// }
	}
}
