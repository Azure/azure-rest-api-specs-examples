package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3ca2244e6becf79bdd1198e9d66f0e8e8d3a0977/specification/storagecache/resource-manager/Microsoft.StorageCache/StorageCache/stable/2026-01-01/examples/expansionJobs_ListByAmlFilesystem.json
func ExampleExpansionJobsClient_NewListByAmlFilesystemPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ExpansionJobsListResult = armstoragecache.ExpansionJobsListResult{
		// 	Value: []*armstoragecache.ExpansionJob{
		// 		{
		// 			Name: to.Ptr("expansionjob1"),
		// 			Type: to.Ptr("Microsoft.StorageCache/amlFilesystem/expansionJob"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlfilesystems/fs1/expansionJobs/expansionjob1"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"Dept": to.Ptr("ContosoAds"),
		// 			},
		// 			Properties: &armstoragecache.ExpansionJobProperties{
		// 				NewStorageCapacityTiB: to.Ptr[float32](16),
		// 				ProvisioningState: to.Ptr(armstoragecache.ExpansionJobPropertiesProvisioningStateSucceeded),
		// 				Status: &armstoragecache.ExpansionJobPropertiesStatus{
		// 					CompletionTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-21T19:15:43.511Z"); return t}()),
		// 					PercentComplete: to.Ptr[float32](100),
		// 					StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-21T17:25:43.511Z"); return t}()),
		// 					State: to.Ptr(armstoragecache.ExpansionJobStatusTypeCompleted),
		// 					StatusMessage: to.Ptr("Expansion completed successfully"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("expansionjob2"),
		// 			Type: to.Ptr("Microsoft.StorageCache/amlFilesystem/expansionJob"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlfilesystems/fs1/expansionJobs/expansionjob2"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"Dept": to.Ptr("ContosoFinance"),
		// 			},
		// 			Properties: &armstoragecache.ExpansionJobProperties{
		// 				NewStorageCapacityTiB: to.Ptr[float32](32),
		// 				ProvisioningState: to.Ptr(armstoragecache.ExpansionJobPropertiesProvisioningStateSucceeded),
		// 				Status: &armstoragecache.ExpansionJobPropertiesStatus{
		// 					PercentComplete: to.Ptr[float32](45),
		// 					StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-22T10:30:15.120Z"); return t}()),
		// 					State: to.Ptr(armstoragecache.ExpansionJobStatusTypeInProgress),
		// 					StatusMessage: to.Ptr("Expansion is in progress"),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("expansionjob3"),
		// 			Type: to.Ptr("Microsoft.StorageCache/amlFilesystem/expansionJob"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlfilesystems/fs1/expansionJobs/expansionjob3"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"Dept": to.Ptr("ContosoHR"),
		// 			},
		// 			Properties: &armstoragecache.ExpansionJobProperties{
		// 				NewStorageCapacityTiB: to.Ptr[float32](24),
		// 				ProvisioningState: to.Ptr(armstoragecache.ExpansionJobPropertiesProvisioningStateSucceeded),
		// 				Status: &armstoragecache.ExpansionJobPropertiesStatus{
		// 					CompletionTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-21T08:50:22.334Z"); return t}()),
		// 					PercentComplete: to.Ptr[float32](0),
		// 					StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-21T08:45:22.334Z"); return t}()),
		// 					State: to.Ptr(armstoragecache.ExpansionJobStatusTypeFailed),
		// 					StatusCode: to.Ptr("InsufficientCapacity"),
		// 					StatusMessage: to.Ptr("Insufficient capacity available in the region for the requested expansion"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
