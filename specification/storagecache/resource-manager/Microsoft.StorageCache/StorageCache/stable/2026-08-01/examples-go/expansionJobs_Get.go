package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: 2026-08-01/expansionJobs_Get.json
func ExampleExpansionJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpansionJobsClient().Get(ctx, "scgroup", "fs1", "expansionjob1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragecache.ExpansionJobsClientGetResponse{
	// 	ExpansionJob: armstoragecache.ExpansionJob{
	// 		Name: to.Ptr("expansionjob1"),
	// 		Type: to.Ptr("Microsoft.StorageCache/amlFilesystems/expansionJobs"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlFilesystems/fs1/expansionJobs/expansionjob1"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armstoragecache.ExpansionJobProperties{
	// 			NewStorageCapacityTiB: to.Ptr[float32](16),
	// 			ProvisioningState: to.Ptr(armstoragecache.ExpansionJobPropertiesProvisioningStateSucceeded),
	// 			RunRebalanceJob: to.Ptr(true),
	// 			RebalanceJobID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlFilesystems/fs1/rebalanceJobs/expansionjob1-rebalance"),
	// 			Status: &armstoragecache.ExpansionJobPropertiesStatus{
	// 				CompletionTimeUTC: to.Ptr(time.Date(2024, time.March, 21, 19, 15, 43, 511000000, time.UTC)),
	// 				PercentComplete: to.Ptr[float32](100),
	// 				StartTimeUTC: to.Ptr(time.Date(2024, time.March, 21, 17, 25, 43, 511000000, time.UTC)),
	// 				State: to.Ptr(armstoragecache.ExpansionJobStatusTypeCompleted),
	// 				StatusMessage: to.Ptr("Expansion completed successfully"),
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"Dept": to.Ptr("ContosoAds"),
	// 		},
	// 	},
	// }
}
