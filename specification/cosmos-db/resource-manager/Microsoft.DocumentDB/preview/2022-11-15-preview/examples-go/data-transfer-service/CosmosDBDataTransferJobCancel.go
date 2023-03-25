package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-11-15-preview/examples/data-transfer-service/CosmosDBDataTransferJobCancel.json
func ExampleDataTransferJobsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataTransferJobsClient().Cancel(ctx, "rg1", "ddb1", "j1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DataTransferJobGetResults = armcosmos.DataTransferJobGetResults{
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/dataTransferJobs"),
	// 	ID: to.Ptr("ab1b6f34-b33c-46b1-98c7-3a0a63fd0d16"),
	// 	Properties: &armcosmos.DataTransferJobProperties{
	// 		Destination: &armcosmos.AzureBlobDataTransferDataSourceSink{
	// 			Component: to.Ptr(armcosmos.DataTransferComponentAzureBlobStorage),
	// 			ContainerName: to.Ptr("blob_container"),
	// 			EndpointURL: to.Ptr("https://blob.windows.net"),
	// 		},
	// 		JobName: to.Ptr("j1"),
	// 		LastUpdatedUTCTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-12T11:10:55.2780000Z"); return t}()),
	// 		ProcessedCount: to.Ptr[int64](20),
	// 		Source: &armcosmos.CassandraDataTransferDataSourceSink{
	// 			Component: to.Ptr(armcosmos.DataTransferComponentCosmosDBCassandra),
	// 			KeyspaceName: to.Ptr("keyspace"),
	// 			TableName: to.Ptr("table"),
	// 		},
	// 		Status: to.Ptr("Cancelled"),
	// 		TotalCount: to.Ptr[int64](100),
	// 	},
	// }
}
