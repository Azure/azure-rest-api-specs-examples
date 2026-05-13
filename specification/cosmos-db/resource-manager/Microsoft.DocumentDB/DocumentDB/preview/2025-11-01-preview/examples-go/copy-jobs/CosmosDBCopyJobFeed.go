package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
)

// Generated from example definition: 2025-11-01-preview/copy-jobs/CosmosDBCopyJobFeed.json
func ExampleCopyJobsClient_NewListByDatabaseAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCopyJobsClient().NewListByDatabaseAccountPager("rg1", "ddb1", nil)
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
		// page = armcosmos.CopyJobsClientListByDatabaseAccountResponse{
		// 	CopyJobFeedResults: armcosmos.CopyJobFeedResults{
		// 		Value: []*armcosmos.CopyJobGetResults{
		// 			{
		// 				Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/copyJobs"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/copyJobs/j1"),
		// 				Properties: &armcosmos.CopyJobProperties{
		// 					Duration: to.Ptr("00:00:00"),
		// 					JobProperties: &armcosmos.NoSQLRUToNoSQLRUCopyJobProperties{
		// 						JobType: to.Ptr(armcosmos.CopyJobTypeNoSQLRUToNoSQLRU),
		// 						Tasks: []*armcosmos.NoSQLRUToNoSQLRUCopyJobTask{
		// 							{
		// 								Destination: &armcosmos.DBNoSQLContainer{
		// 									ContainerName: to.Ptr("destTable1"),
		// 									DatabaseName: to.Ptr("destDB1"),
		// 								},
		// 								Source: &armcosmos.DBNoSQLContainer{
		// 									ContainerName: to.Ptr("sourceTable1"),
		// 									DatabaseName: to.Ptr("sourceDb1"),
		// 								},
		// 							},
		// 							{
		// 								Destination: &armcosmos.DBNoSQLContainer{
		// 									ContainerName: to.Ptr("destTable2"),
		// 									DatabaseName: to.Ptr("destDB2"),
		// 								},
		// 								Source: &armcosmos.DBNoSQLContainer{
		// 									ContainerName: to.Ptr("sourceTable2"),
		// 									DatabaseName: to.Ptr("sourceDb2"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					LastUpdatedUTCTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-12T11:10:55.2780000Z"); return t}()),
		// 					Mode: to.Ptr(armcosmos.CopyJobModeOffline),
		// 					ProcessedCount: to.Ptr[int64](0),
		// 					Status: to.Ptr(armcosmos.CopyJobStatusRunning),
		// 					TotalCount: to.Ptr[int64](0),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
