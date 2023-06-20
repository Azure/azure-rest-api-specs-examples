package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBSqlDatabasePartitionMerge.json
func ExampleSQLResourcesClient_BeginSQLDatabasePartitionMerge() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLResourcesClient().BeginSQLDatabasePartitionMerge(ctx, "rgName", "ddb1", "databaseName", armcosmos.MergeParameters{
		IsDryRun: to.Ptr(false),
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
	// res.PhysicalPartitionStorageInfoCollection = armcosmos.PhysicalPartitionStorageInfoCollection{
	// 	PhysicalPartitionStorageInfoCollection: []*armcosmos.PhysicalPartitionStorageInfo{
	// 		{
	// 			ID: to.Ptr("0"),
	// 			StorageInKB: to.Ptr[float64](333),
	// 		},
	// 		{
	// 			ID: to.Ptr("1"),
	// 			StorageInKB: to.Ptr[float64](305),
	// 		},
	// 		{
	// 			ID: to.Ptr("177"),
	// 			StorageInKB: to.Ptr[float64](368),
	// 		},
	// 		{
	// 			ID: to.Ptr("178"),
	// 			StorageInKB: to.Ptr[float64](96313),
	// 		},
	// 		{
	// 			ID: to.Ptr("5"),
	// 			StorageInKB: to.Ptr[float64](194),
	// 		},
	// 		{
	// 			ID: to.Ptr("6"),
	// 			StorageInKB: to.Ptr[float64](331),
	// 	}},
	// }
}
