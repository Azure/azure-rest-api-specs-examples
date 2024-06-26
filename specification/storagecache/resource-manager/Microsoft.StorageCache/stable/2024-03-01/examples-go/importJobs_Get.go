package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/33c4457b1d13f83965f4fe3367dca4a6df898100/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2024-03-01/examples/importJobs_Get.json
func ExampleImportJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewImportJobsClient().Get(ctx, "scgroup", "fs1", "job1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ImportJob = armstoragecache.ImportJob{
	// 	Name: to.Ptr("job1"),
	// 	Type: to.Ptr("Microsoft.StorageCache/amlFilesystem/importJob"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlfilesystems/fs1/storagetargets/job1"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armstoragecache.ImportJobProperties{
	// 		ConflictResolutionMode: to.Ptr(armstoragecache.ConflictResolutionModeOverwriteAlways),
	// 		ImportPrefixes: []*string{
	// 			to.Ptr("/")},
	// 			MaximumErrors: to.Ptr[int32](0),
	// 			ProvisioningState: to.Ptr(armstoragecache.ImportJobProvisioningStateTypeSucceeded),
	// 			Status: &armstoragecache.ImportJobPropertiesStatus{
	// 				BlobsImportedPerSecond: to.Ptr[int64](4000),
	// 				BlobsWalkedPerSecond: to.Ptr[int64](10000),
	// 				LastCompletionTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T18:25:43.511Z"); return t}()),
	// 				LastStartedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-04-21T17:25:43.511Z"); return t}()),
	// 				State: to.Ptr(armstoragecache.ImportStatusTypeCompleted),
	// 				StatusMessage: to.Ptr("Import job completed successfully"),
	// 				TotalBlobsImported: to.Ptr[int64](1000000),
	// 				TotalBlobsWalked: to.Ptr[int64](1000000),
	// 				TotalConflicts: to.Ptr[int32](1),
	// 				TotalErrors: to.Ptr[int32](1),
	// 			},
	// 		},
	// 	}
}
