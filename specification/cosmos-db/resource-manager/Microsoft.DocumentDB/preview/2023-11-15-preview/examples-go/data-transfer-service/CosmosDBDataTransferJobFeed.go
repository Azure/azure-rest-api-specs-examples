package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6f8faf5da91b5b9af5f3512fe609e22e99383d41/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2023-11-15-preview/examples/data-transfer-service/CosmosDBDataTransferJobFeed.json
func ExampleDataTransferJobsClient_NewListByDatabaseAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDataTransferJobsClient().NewListByDatabaseAccountPager("rg1", "ddb1", nil)
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
		// page.DataTransferJobFeedResults = armcosmos.DataTransferJobFeedResults{
		// 	Value: []*armcosmos.DataTransferJobGetResults{
		// 		{
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/dataTransferJobs"),
		// 			ID: to.Ptr("ab1b6f34-b33c-46b1-98c7-3a0a63fd0d16"),
		// 			Properties: &armcosmos.DataTransferJobProperties{
		// 				Destination: &armcosmos.AzureBlobDataTransferDataSourceSink{
		// 					Component: to.Ptr(armcosmos.DataTransferComponentAzureBlobStorage),
		// 					ContainerName: to.Ptr("blob_container"),
		// 					EndpointURL: to.Ptr("https://blob.windows.net"),
		// 				},
		// 				Duration: to.Ptr("00:00:00"),
		// 				JobName: to.Ptr("j1"),
		// 				LastUpdatedUTCTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-12T12:04:23.153Z"); return t}()),
		// 				ProcessedCount: to.Ptr[int64](100),
		// 				Source: &armcosmos.CassandraDataTransferDataSourceSink{
		// 					Component: to.Ptr(armcosmos.DataTransferComponentCosmosDBCassandra),
		// 					KeyspaceName: to.Ptr("keyspace"),
		// 					TableName: to.Ptr("table"),
		// 				},
		// 				Status: to.Ptr("Completed"),
		// 				TotalCount: to.Ptr[int64](100),
		// 			},
		// 	}},
		// }
	}
}
