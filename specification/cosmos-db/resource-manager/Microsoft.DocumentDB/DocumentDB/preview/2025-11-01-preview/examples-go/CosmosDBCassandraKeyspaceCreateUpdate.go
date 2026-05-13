package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/CosmosDBCassandraKeyspaceCreateUpdate.json
func ExampleCassandraResourcesClient_BeginCreateUpdateCassandraKeyspace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("d6beadc3-5469-4109-ab3c-22a036c7e175", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCassandraResourcesClient().BeginCreateUpdateCassandraKeyspace(ctx, "rg1", "ddb1", "keyspaceName", armcosmos.CassandraKeyspaceCreateUpdateParameters{
		Location: to.Ptr("West US"),
		Properties: &armcosmos.CassandraKeyspaceCreateUpdateProperties{
			Options: &armcosmos.CreateUpdateOptions{},
			Resource: &armcosmos.CassandraKeyspaceResource{
				ID: to.Ptr("keyspaceName"),
			},
		},
		Tags: map[string]*string{},
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
	// res = armcosmos.CassandraResourcesClientCreateUpdateCassandraKeyspaceResponse{
	// 	CassandraKeyspaceGetResults: &armcosmos.CassandraKeyspaceGetResults{
	// 		Name: to.Ptr("keyspaceName"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/cassandraKeyspaces"),
	// 		ID: to.Ptr("/subscriptions/d6beadc3-5469-4109-ab3c-22a036c7e175/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/cassandraKeyspaces/keyspaceName"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armcosmos.CassandraKeyspaceGetProperties{
	// 			Resource: &armcosmos.CassandraKeyspaceGetPropertiesResource{
	// 				ID: to.Ptr("keyspaceName"),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
