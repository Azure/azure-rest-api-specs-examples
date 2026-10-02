package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: 2026-08-01/RebalanceJobs_Get.json
func ExampleRebalanceJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRebalanceJobsClient().Get(ctx, "scgroup", "fs1", "expansionjob1-rebalance", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragecache.RebalanceJobsClientGetResponse{
	// 	RebalanceJob: armstoragecache.RebalanceJob{
	// 		Name: to.Ptr("expansionjob1-rebalance"),
	// 		Type: to.Ptr("Microsoft.StorageCache/amlFilesystems/rebalanceJobs"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlFilesystems/fs1/rebalanceJobs/expansionjob1-rebalance"),
	// 		Properties: &armstoragecache.RebalanceJobProperties{
	// 			ProvisioningState: to.Ptr(armstoragecache.RebalanceJobPropertiesProvisioningStateSucceeded),
	// 			AdminStatus: to.Ptr(armstoragecache.RebalanceJobAdminStatusActive),
	// 			ExpansionJobID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlFilesystems/fs1/expansionJobs/expansionjob1"),
	// 			Status: &armstoragecache.RebalanceJobPropertiesStatus{
	// 				State: to.Ptr(armstoragecache.RebalanceJobStatusTypeInProgress),
	// 				StatusCode: to.Ptr("RebalanceInProgress"),
	// 				StatusMessage: to.Ptr("Rebalance is 45% complete, migrating files."),
	// 				PercentComplete: to.Ptr[float32](50),
	// 				BalancePercent: to.Ptr[float64](45.2),
	// 				EstimatedRemainingSeconds: to.Ptr[int32](3600),
	// 				FilesMigrated: to.Ptr[int64](150000),
	// 				DirsMigrated: to.Ptr[int64](12000),
	// 				BytesMoved: to.Ptr[int64](1073741824),
	// 				FilesMovedPerSecond: to.Ptr[float64](250.5),
	// 				ThroughputMiBps: to.Ptr[float64](1200),
	// 				TotalErrors: to.Ptr[int32](0),
	// 				TotalSkipped: to.Ptr[int64](42),
	// 				StartTimeUTC: to.Ptr(time.Date(2025, time.September, 12, 10, 0, 0, 0, time.UTC)),
	// 			},
	// 		},
	// 	},
	// }
}
