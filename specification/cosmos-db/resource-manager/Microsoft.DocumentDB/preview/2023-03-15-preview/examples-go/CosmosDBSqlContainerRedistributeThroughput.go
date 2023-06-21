package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1b33e81bbdc28fcd6644a1315b8d7b1b6d030590/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-03-15-preview/examples/CosmosDBSqlContainerRedistributeThroughput.json
func ExampleSQLResourcesClient_BeginSQLContainerRedistributeThroughput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLResourcesClient().BeginSQLContainerRedistributeThroughput(ctx, "rg1", "ddb1", "databaseName", "containerName", armcosmos.RedistributeThroughputParameters{
		Properties: &armcosmos.RedistributeThroughputProperties{
			Resource: &armcosmos.RedistributeThroughputPropertiesResource{
				SourcePhysicalPartitionThroughputInfo: []*armcosmos.PhysicalPartitionThroughputInfoResource{
					{
						ID:         to.Ptr("2"),
						Throughput: to.Ptr[float64](5000),
					},
					{
						ID: to.Ptr("3"),
					}},
				TargetPhysicalPartitionThroughputInfo: []*armcosmos.PhysicalPartitionThroughputInfoResource{
					{
						ID:         to.Ptr("0"),
						Throughput: to.Ptr[float64](5000),
					},
					{
						ID:         to.Ptr("1"),
						Throughput: to.Ptr[float64](5000),
					}},
				ThroughputPolicy: to.Ptr(armcosmos.ThroughputPolicyTypeCustom),
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
	// res.PhysicalPartitionThroughputInfoResult = armcosmos.PhysicalPartitionThroughputInfoResult{
	// 	Properties: &armcosmos.PhysicalPartitionThroughputInfoResultProperties{
	// 		Resource: &armcosmos.PhysicalPartitionThroughputInfoResultPropertiesResource{
	// 			PhysicalPartitionThroughputInfo: []*armcosmos.PhysicalPartitionThroughputInfoResource{
	// 				{
	// 					ID: to.Ptr("0"),
	// 					Throughput: to.Ptr[float64](5000),
	// 				},
	// 				{
	// 					ID: to.Ptr("1"),
	// 					Throughput: to.Ptr[float64](5000),
	// 				},
	// 				{
	// 					ID: to.Ptr("2"),
	// 					Throughput: to.Ptr[float64](5000),
	// 				},
	// 				{
	// 					ID: to.Ptr("3"),
	// 					Throughput: to.Ptr[float64](3000),
	// 			}},
	// 		},
	// 	},
	// }
}
