package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: 2026-08-01/RebalanceJobs_Update.json
func ExampleRebalanceJobsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRebalanceJobsClient().BeginUpdate(ctx, "scgroup", "fs1", "expansionjob1-rebalance", armstoragecache.RebalanceJobUpdate{
		Properties: &armstoragecache.RebalanceJobUpdateProperties{
			AdminStatus: to.Ptr(armstoragecache.RebalanceJobAdminStatusCancel),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragecache.RebalanceJobsClientUpdateResponse{
	// 	RebalanceJob: armstoragecache.RebalanceJob{
	// 		Name: to.Ptr("expansionjob1-rebalance"),
	// 		Type: to.Ptr("Microsoft.StorageCache/amlFilesystems/rebalanceJobs"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlFilesystems/fs1/rebalanceJobs/expansionjob1-rebalance"),
	// 		Properties: &armstoragecache.RebalanceJobProperties{
	// 			ProvisioningState: to.Ptr(armstoragecache.RebalanceJobPropertiesProvisioningStateSucceeded),
	// 			AdminStatus: to.Ptr(armstoragecache.RebalanceJobAdminStatusCancel),
	// 			ExpansionJobID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlFilesystems/fs1/expansionJobs/expansionjob1"),
	// 			Status: &armstoragecache.RebalanceJobPropertiesStatus{
	// 				State: to.Ptr(armstoragecache.RebalanceJobStatusTypeInProgress),
	// 				StatusCode: to.Ptr("CancelRequested"),
	// 				StatusMessage: to.Ptr("Rebalance cancellation has been requested."),
	// 				PercentComplete: to.Ptr[float32](70),
	// 				BalancePercent: to.Ptr[float64](67.8),
	// 				FilesMigrated: to.Ptr[int64](320000),
	// 				DirsMigrated: to.Ptr[int64](25000),
	// 				BytesMoved: to.Ptr[int64](3435973836),
	// 				FilesMovedPerSecond: to.Ptr[float64](0),
	// 				ThroughputMiBps: to.Ptr[float64](0),
	// 				TotalErrors: to.Ptr[int32](2),
	// 				TotalSkipped: to.Ptr[int64](150),
	// 				StartTimeUTC: to.Ptr(time.Date(2025, time.September, 12, 10, 0, 0, 0, time.UTC)),
	// 			},
	// 		},
	// 	},
	// }
}
