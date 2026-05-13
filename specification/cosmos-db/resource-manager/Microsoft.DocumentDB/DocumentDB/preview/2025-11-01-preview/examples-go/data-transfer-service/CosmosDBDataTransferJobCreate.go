package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/data-transfer-service/CosmosDBDataTransferJobCreate.json
func ExampleDataTransferJobsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDataTransferJobsClient().Create(ctx, "rg1", "ddb1", "j1", armcosmos.CreateJobRequest{
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
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcosmos.DataTransferJobsClientCreateResponse{
	// 	DataTransferJobGetResults: &armcosmos.DataTransferJobGetResults{
	// 		Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/dataTransferJobs"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/dataTransferJobs/j1"),
	// 		Properties: &armcosmos.DataTransferJobProperties{
	// 			Destination: &armcosmos.AzureBlobDataTransferDataSourceSink{
	// 				Component: to.Ptr(armcosmos.DataTransferComponentAzureBlobStorage),
	// 				ContainerName: to.Ptr("blob_container"),
	// 				EndpointURL: to.Ptr("https://blob.windows.net"),
	// 			},
	// 			Duration: to.Ptr("00:00:00"),
	// 			JobName: to.Ptr("j1"),
	// 			LastUpdatedUTCTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-12T11:10:55.2780000Z"); return t}()),
	// 			ProcessedCount: to.Ptr[int64](0),
	// 			Source: &armcosmos.CassandraDataTransferDataSourceSink{
	// 				Component: to.Ptr(armcosmos.DataTransferComponentCosmosDBCassandra),
	// 				KeyspaceName: to.Ptr("keyspace"),
	// 				TableName: to.Ptr("table"),
	// 			},
	// 			Status: to.Ptr("Running"),
	// 			TotalCount: to.Ptr[int64](0),
	// 		},
	// 	},
	// }
}
