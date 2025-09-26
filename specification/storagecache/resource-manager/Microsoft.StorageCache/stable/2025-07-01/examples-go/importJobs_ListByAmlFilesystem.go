package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d18f451b13796bded0808d508c6297f2227271d5/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2025-07-01/examples/importJobs_ListByAmlFilesystem.json
func ExampleImportJobsClient_NewListByAmlFilesystemPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewImportJobsClient().NewListByAmlFilesystemPager("scgroup", "fs1", nil)
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
		// page.ImportJobsListResult = armstoragecache.ImportJobsListResult{
		// 	Value: []*armstoragecache.ImportJob{
		// 		{
		// 			Name: to.Ptr("job1"),
		// 			Type: to.Ptr("Microsoft.StorageCache/amlFilesystem/importJob"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlfilesystems/fs1/importJobs/job1"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armstoragecache.ImportJobProperties{
		// 				AdminStatus: to.Ptr(armstoragecache.ImportJobAdminStatusActive),
		// 				ConflictResolutionMode: to.Ptr(armstoragecache.ConflictResolutionModeOverwriteAlways),
		// 				ImportPrefixes: []*string{
		// 					to.Ptr("/")},
		// 					MaximumErrors: to.Ptr[int32](0),
		// 					ProvisioningState: to.Ptr(armstoragecache.ImportJobProvisioningStateTypeSucceeded),
		// 					Status: &armstoragecache.ImportJobPropertiesStatus{
		// 						BlobsImportedPerSecond: to.Ptr[int64](4000),
		// 						BlobsWalkedPerSecond: to.Ptr[int64](10000),
		// 						ImportedDirectories: to.Ptr[int64](500),
		// 						ImportedFiles: to.Ptr[int64](1000),
		// 						ImportedSymlinks: to.Ptr[int64](1000),
		// 						LastCompletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T18:25:43.511Z"); return t}()),
		// 						LastStartedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T17:25:43.511Z"); return t}()),
		// 						PreexistingDirectories: to.Ptr[int64](100),
		// 						PreexistingFiles: to.Ptr[int64](500),
		// 						PreexistingSymlinks: to.Ptr[int64](200),
		// 						State: to.Ptr(armstoragecache.ImportStatusTypeCompleted),
		// 						StatusMessage: to.Ptr("Import job completed successfully"),
		// 						TotalBlobsImported: to.Ptr[int64](1000000),
		// 						TotalBlobsWalked: to.Ptr[int64](1000000),
		// 						TotalConflicts: to.Ptr[int32](1),
		// 						TotalErrors: to.Ptr[int32](1),
		// 					},
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("job2"),
		// 				Type: to.Ptr("Microsoft.StorageCache/amlFilesystem/importJob"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlfilesystems/fs1/importJobs/job2"),
		// 				Location: to.Ptr("eastus"),
		// 				Properties: &armstoragecache.ImportJobProperties{
		// 					AdminStatus: to.Ptr(armstoragecache.ImportJobAdminStatusActive),
		// 					ConflictResolutionMode: to.Ptr(armstoragecache.ConflictResolutionModeSkip),
		// 					ImportPrefixes: []*string{
		// 						to.Ptr("/dir1")},
		// 						MaximumErrors: to.Ptr[int32](0),
		// 						ProvisioningState: to.Ptr(armstoragecache.ImportJobProvisioningStateTypeSucceeded),
		// 						Status: &armstoragecache.ImportJobPropertiesStatus{
		// 							BlobsImportedPerSecond: to.Ptr[int64](4000),
		// 							BlobsWalkedPerSecond: to.Ptr[int64](10000),
		// 							ImportedDirectories: to.Ptr[int64](500),
		// 							ImportedFiles: to.Ptr[int64](1000),
		// 							ImportedSymlinks: to.Ptr[int64](1000),
		// 							LastCompletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T18:25:43.511Z"); return t}()),
		// 							LastStartedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T17:25:43.511Z"); return t}()),
		// 							PreexistingDirectories: to.Ptr[int64](100),
		// 							PreexistingFiles: to.Ptr[int64](500),
		// 							PreexistingSymlinks: to.Ptr[int64](200),
		// 							State: to.Ptr(armstoragecache.ImportStatusTypeCompleted),
		// 							StatusMessage: to.Ptr("Import job completed successfully"),
		// 							TotalBlobsImported: to.Ptr[int64](1000000),
		// 							TotalBlobsWalked: to.Ptr[int64](1000000),
		// 							TotalConflicts: to.Ptr[int32](1),
		// 							TotalErrors: to.Ptr[int32](1),
		// 						},
		// 					},
		// 			}},
		// 		}
	}
}
