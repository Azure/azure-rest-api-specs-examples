package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBSqlContainerRedistributeThroughput.json
func ExampleSQLResourcesClient_BeginSQLContainerRedistributeThroughput() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewSQLResourcesClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginSQLContainerRedistributeThroughput(ctx, "rg1", "ddb1", "databaseName", "containerName", armcosmos.RedistributeThroughputParameters{
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
	// TODO: use response item
	_ = res
}
