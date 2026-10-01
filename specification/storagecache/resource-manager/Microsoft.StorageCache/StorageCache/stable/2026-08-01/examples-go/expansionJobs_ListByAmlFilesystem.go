package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: 2026-08-01/expansionJobs_ListByAmlFilesystem.json
func ExampleExpansionJobsClient_NewListByAmlFilesystemPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewExpansionJobsClient().NewListByAmlFilesystemPager("scgroup", "fs1", nil)
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
		// page = armstoragecache.ExpansionJobsClientListByAmlFilesystemResponse{
		// 	ExpansionJobsListResult: armstoragecache.ExpansionJobsListResult{
		// 		NextLink: to.Ptr("https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlfilesystems/fs1/expansionJobs?api-version=2026-08-01&$skiptoken=abc123"),
		// 		Value: []*armstoragecache.ExpansionJob{
		// 			{
		// 				Name: to.Ptr("expansionjob1"),
		// 				Type: to.Ptr("Microsoft.StorageCache/amlFilesystem/expansionJob"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlfilesystems/fs1/expansionJobs/expansionjob1"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armstoragecache.ExpansionJobProperties{
		// 					NewStorageCapacityTiB: to.Ptr[float32](16),
		// 					ProvisioningState: to.Ptr(armstoragecache.ExpansionJobPropertiesProvisioningStateSucceeded),
		// 					RunRebalanceJob: to.Ptr(true),
		// 					Status: &armstoragecache.ExpansionJobPropertiesStatus{
		// 						CompletionTimeUTC: to.Ptr(time.Date(2024, time.March, 21, 19, 15, 43, 511000000, time.UTC)),
		// 						PercentComplete: to.Ptr[float32](100),
		// 						StartTimeUTC: to.Ptr(time.Date(2024, time.March, 21, 17, 25, 43, 511000000, time.UTC)),
		// 						State: to.Ptr(armstoragecache.ExpansionJobStatusTypeCompleted),
		// 						StatusMessage: to.Ptr("Expansion completed successfully"),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("ContosoAds"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("expansionjob2"),
		// 				Type: to.Ptr("Microsoft.StorageCache/amlFilesystem/expansionJob"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlfilesystems/fs1/expansionJobs/expansionjob2"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armstoragecache.ExpansionJobProperties{
		// 					NewStorageCapacityTiB: to.Ptr[float32](32),
		// 					ProvisioningState: to.Ptr(armstoragecache.ExpansionJobPropertiesProvisioningStateSucceeded),
		// 					RunRebalanceJob: to.Ptr(true),
		// 					Status: &armstoragecache.ExpansionJobPropertiesStatus{
		// 						PercentComplete: to.Ptr[float32](45),
		// 						StartTimeUTC: to.Ptr(time.Date(2024, time.March, 22, 10, 30, 15, 120000000, time.UTC)),
		// 						State: to.Ptr(armstoragecache.ExpansionJobStatusTypeInProgress),
		// 						StatusMessage: to.Ptr("Expansion is in progress"),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("ContosoFinance"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("expansionjob3"),
		// 				Type: to.Ptr("Microsoft.StorageCache/amlFilesystem/expansionJob"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlfilesystems/fs1/expansionJobs/expansionjob3"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armstoragecache.ExpansionJobProperties{
		// 					NewStorageCapacityTiB: to.Ptr[float32](24),
		// 					ProvisioningState: to.Ptr(armstoragecache.ExpansionJobPropertiesProvisioningStateFailed),
		// 					RunRebalanceJob: to.Ptr(true),
		// 					Status: &armstoragecache.ExpansionJobPropertiesStatus{
		// 						CompletionTimeUTC: to.Ptr(time.Date(2024, time.March, 21, 8, 50, 22, 334000000, time.UTC)),
		// 						PercentComplete: to.Ptr[float32](0),
		// 						StartTimeUTC: to.Ptr(time.Date(2024, time.March, 21, 8, 45, 22, 334000000, time.UTC)),
		// 						State: to.Ptr(armstoragecache.ExpansionJobStatusTypeFailed),
		// 						StatusCode: to.Ptr("InsufficientCapacity"),
		// 						StatusMessage: to.Ptr("Insufficient capacity available in the region for the requested expansion"),
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"Dept": to.Ptr("ContosoHR"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
