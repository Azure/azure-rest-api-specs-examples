package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-04-15/examples/CosmosDBCassandraKeyspaceList.json
func ExampleCassandraResourcesClient_NewListCassandraKeyspacesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCassandraResourcesClient().NewListCassandraKeyspacesPager("rgName", "ddb1", nil)
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
		// page.CassandraKeyspaceListResult = armcosmos.CassandraKeyspaceListResult{
		// 	Value: []*armcosmos.CassandraKeyspaceGetResults{
		// 		{
		// 			Name: to.Ptr("keyspaceName"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/cassandraKeyspaces"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/cassandraKeyspaces/keyspaceName"),
		// 			Location: to.Ptr("West US"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armcosmos.CassandraKeyspaceGetProperties{
		// 				Resource: &armcosmos.CassandraKeyspaceGetPropertiesResource{
		// 					ID: to.Ptr("keyspaceName"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
