package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d18f451b13796bded0808d508c6297f2227271d5/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2025-07-01/examples/autoImportJobs_CreateOrUpdate.json
func ExampleAutoImportJobsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAutoImportJobsClient().BeginCreateOrUpdate(ctx, "scgroup", "fs1", "autojob1", armstoragecache.AutoImportJob{
		Location: to.Ptr("eastus"),
		Tags: map[string]*string{
			"Dept": to.Ptr("ContosoAds"),
		},
		Properties: &armstoragecache.AutoImportJobProperties{
			AdminStatus: to.Ptr(armstoragecache.AutoImportJobPropertiesAdminStatusEnable),
			AutoImportPrefixes: []*string{
				to.Ptr("/")},
			ConflictResolutionMode: to.Ptr(armstoragecache.ConflictResolutionModeSkip),
			EnableDeletions:        to.Ptr(false),
			MaximumErrors:          to.Ptr[int64](0),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AutoImportJob = armstoragecache.AutoImportJob{
	// 	Name: to.Ptr("autojob1"),
	// 	Type: to.Ptr("Microsoft.StorageCache/amlFilesystem/autoImportJob"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlfilesystems/fs1/autoImportJobs/autojob1"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"Dept": to.Ptr("ContosoAds"),
	// 	},
	// 	Properties: &armstoragecache.AutoImportJobProperties{
	// 		AdminStatus: to.Ptr(armstoragecache.AutoImportJobPropertiesAdminStatusEnable),
	// 		AutoImportPrefixes: []*string{
	// 			to.Ptr("/")},
	// 			ConflictResolutionMode: to.Ptr(armstoragecache.ConflictResolutionModeSkip),
	// 			EnableDeletions: to.Ptr(false),
	// 			MaximumErrors: to.Ptr[int64](0),
	// 			ProvisioningState: to.Ptr(armstoragecache.AutoImportJobPropertiesProvisioningStateSucceeded),
	// 			Status: &armstoragecache.AutoImportJobPropertiesStatus{
	// 				BlobSyncEvents: &armstoragecache.AutoImportJobPropertiesStatusBlobSyncEvents{
	// 					Deletions: to.Ptr[int64](20),
	// 					ImportedDirectories: to.Ptr[int64](100),
	// 					ImportedFiles: to.Ptr[int64](500),
	// 					ImportedSymlinks: to.Ptr[int64](50),
	// 					LastChangeFeedEventConsumedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-21T18:35:43.511Z"); return t}()),
	// 					LastTimeFullySynchronized: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-21T18:25:43.511Z"); return t}()),
	// 					PreexistingDirectories: to.Ptr[int64](1000),
	// 					PreexistingFiles: to.Ptr[int64](10000),
	// 					PreexistingSymlinks: to.Ptr[int64](200),
	// 					RateOfBlobImport: to.Ptr[int64](120),
	// 					TotalBlobsImported: to.Ptr[int64](11850),
	// 					TotalConflicts: to.Ptr[int64](1),
	// 					TotalErrors: to.Ptr[int64](3),
	// 				},
	// 				ImportedDirectories: to.Ptr[int64](1000),
	// 				ImportedFiles: to.Ptr[int64](8000),
	// 				ImportedSymlinks: to.Ptr[int64](1000),
	// 				LastStartedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-21T17:25:43.511Z"); return t}()),
	// 				PreexistingDirectories: to.Ptr[int64](80000),
	// 				PreexistingFiles: to.Ptr[int64](800000),
	// 				PreexistingSymlinks: to.Ptr[int64](10000),
	// 				RateOfBlobImport: to.Ptr[int64](4000),
	// 				RateOfBlobWalk: to.Ptr[int64](10000),
	// 				ScanEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-21T18:25:43.511Z"); return t}()),
	// 				ScanStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-03-21T17:25:43.511Z"); return t}()),
	// 				State: to.Ptr(armstoragecache.AutoImportJobStateInProgress),
	// 				TotalBlobsImported: to.Ptr[int64](900000),
	// 				TotalBlobsWalked: to.Ptr[int64](1000000),
	// 				TotalConflicts: to.Ptr[int64](5),
	// 				TotalErrors: to.Ptr[int64](10),
	// 			},
	// 		},
	// 	}
}
