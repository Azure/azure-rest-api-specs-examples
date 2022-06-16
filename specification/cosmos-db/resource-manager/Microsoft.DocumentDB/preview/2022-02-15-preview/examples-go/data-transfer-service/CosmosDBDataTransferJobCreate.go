package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/data-transfer-service/CosmosDBDataTransferJobCreate.json
func ExampleDataTransferJobsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewDataTransferJobsClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"rg1",
		"ddb1",
		"j1",
		armcosmos.CreateJobRequest{
			Properties: &armcosmos.DataTransferJobProperties{
				Destination: &armcosmos.AzureBlobDataTransferDataSourceSink{
					Component:     to.Ptr(armcosmos.DataTransferComponentAzureBlobStorage),
					ContainerName: to.Ptr("blob_container"),
					EndpointURL:   to.Ptr("https://blob.windows.net"),
				},
				Source: &armcosmos.CassandraDataTransferDataSourceSink{
					Component:    to.Ptr(armcosmos.DataTransferComponentCosmosDBCassandra),
					KeyspaceName: to.Ptr("keyspace"),
					TableName:    to.Ptr("table"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
