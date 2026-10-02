package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: 2026-08-01/RebalanceJobs_ListByAmlFilesystem.json
func ExampleRebalanceJobsClient_NewListByAmlFilesystemPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRebalanceJobsClient().NewListByAmlFilesystemPager("scgroup", "fs1", nil)
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
		// page = armstoragecache.RebalanceJobsClientListByAmlFilesystemResponse{
		// 	RebalanceJobsListResult: armstoragecache.RebalanceJobsListResult{
		// 		Value: []*armstoragecache.RebalanceJob{
		// 			{
		// 				Name: to.Ptr("expansionjob1-rebalance"),
		// 				Type: to.Ptr("Microsoft.StorageCache/amlFilesystems/rebalanceJobs"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlFilesystems/fs1/rebalanceJobs/expansionjob1-rebalance"),
		// 				Properties: &armstoragecache.RebalanceJobProperties{
		// 					ProvisioningState: to.Ptr(armstoragecache.RebalanceJobPropertiesProvisioningStateSucceeded),
		// 					AdminStatus: to.Ptr(armstoragecache.RebalanceJobAdminStatusActive),
		// 					ExpansionJobID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlFilesystems/fs1/expansionJobs/expansionjob1"),
		// 					Status: &armstoragecache.RebalanceJobPropertiesStatus{
		// 						State: to.Ptr(armstoragecache.RebalanceJobStatusTypeCompleted),
		// 						StatusCode: to.Ptr("Success"),
		// 						StatusMessage: to.Ptr("Rebalance completed successfully."),
		// 						PercentComplete: to.Ptr[float32](100),
		// 						BalancePercent: to.Ptr[float64](100),
		// 						FilesMigrated: to.Ptr[int64](500000),
		// 						DirsMigrated: to.Ptr[int64](40000),
		// 						BytesMoved: to.Ptr[int64](5368709120),
		// 						FilesMovedPerSecond: to.Ptr[float64](0),
		// 						ThroughputMiBps: to.Ptr[float64](0),
		// 						TotalErrors: to.Ptr[int32](0),
		// 						TotalSkipped: to.Ptr[int64](230),
		// 						StartTimeUTC: to.Ptr(time.Date(2025, time.September, 12, 10, 0, 0, 0, time.UTC)),
		// 						CompletionTimeUTC: to.Ptr(time.Date(2025, time.September, 14, 6, 30, 0, 0, time.UTC)),
		// 					},
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
