package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6f8faf5da91b5b9af5f3512fe609e22e99383d41/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-11-15-preview/examples/CosmosDBCassandraViewCreateUpdate.json
func ExampleCassandraResourcesClient_BeginCreateUpdateCassandraView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraResourcesClient().BeginCreateUpdateCassandraView(ctx, "rg1", "ddb1", "keyspacename", "viewname", armcosmos.CassandraViewCreateUpdateParameters{
		Tags: map[string]*string{},
		Properties: &armcosmos.CassandraViewCreateUpdateProperties{
			Options: &armcosmos.CreateUpdateOptions{},
			Resource: &armcosmos.CassandraViewResource{
				ID:             to.Ptr("viewname"),
				ViewDefinition: to.Ptr("SELECT columna, columnb, columnc FROM keyspacename.srctablename WHERE columna IS NOT NULL AND columnc IS NOT NULL PRIMARY (columnc, columna)"),
			},
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
	// res.CassandraViewGetResults = armcosmos.CassandraViewGetResults{
	// 	Name: to.Ptr("viewname"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/cassandraKeyspaces/views"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/cassandraKeyspaces/keyspacename/views/viewname"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armcosmos.CassandraViewGetProperties{
	// 		Resource: &armcosmos.CassandraViewGetPropertiesResource{
	// 			ID: to.Ptr("viewname"),
	// 			ViewDefinition: to.Ptr("SELECT columna, columnb, columnc FROM keyspacename.srctablename WHERE columna IS NOT NULL AND columnc IS NOT NULL PRIMARY KEY (columnc, columna)"),
	// 		},
	// 	},
	// }
}
