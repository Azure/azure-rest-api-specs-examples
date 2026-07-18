package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2026-03-15/CosmosDBSqlDatabaseMigrateToAutoscale.json
func ExampleSQLResourcesClient_BeginMigrateSQLDatabaseToAutoscale() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLResourcesClient().BeginMigrateSQLDatabaseToAutoscale(ctx, "rg1", "ddb1", "databaseName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.SQLResourcesClientMigrateSQLDatabaseToAutoscaleResponse{
	// 	ThroughputSettingsGetResults: armcosmos.ThroughputSettingsGetResults{
	// 		Properties: &armcosmos.ThroughputSettingsGetProperties{
	// 			Resource: &armcosmos.ThroughputSettingsGetPropertiesResource{
	// 				Throughput: to.Ptr[int32](400),
	// 				AutoscaleSettings: &armcosmos.AutoscaleSettingsResource{
	// 					MaxThroughput: to.Ptr[int32](4000),
	// 				},
	// 				MinimumThroughput: to.Ptr("4000"),
	// 				OfferReplacePending: to.Ptr("false"),
	// 				Rid: to.Ptr("PD5DALigDgw="),
	// 				Ts: to.Ptr[float32](1459200611),
	// 				Etag: to.Ptr("\"00005900-0000-0000-0000-56f9a2630000\""),
	// 			},
	// 		},
	// 	},
	// }
}
