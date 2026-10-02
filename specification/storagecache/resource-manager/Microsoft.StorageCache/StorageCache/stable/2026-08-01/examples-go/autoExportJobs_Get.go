package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: 2026-08-01/autoExportJobs_Get.json
func ExampleAutoExportJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAutoExportJobsClient().Get(ctx, "scgroup", "fs1", "job1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armstoragecache.AutoExportJobsClientGetResponse{
	// 	AutoExportJob: armstoragecache.AutoExportJob{
	// 		Name: to.Ptr("job1"),
	// 		Type: to.Ptr("Microsoft.StorageCache/amlFilesystems/autoExportJobs"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlFilesystems/fs1/autoExportJobs/job1"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armstoragecache.AutoExportJobProperties{
	// 			AdminStatus: to.Ptr(armstoragecache.AutoExportJobAdminStatusEnable),
	// 			AutoExportPrefixes: []*string{
	// 				to.Ptr("/"),
	// 			},
	// 			ProvisioningState: to.Ptr(armstoragecache.AutoExportJobProvisioningStateTypeSucceeded),
	// 			Status: &armstoragecache.AutoExportJobPropertiesStatus{
	// 				CurrentIterationFilesDiscovered: to.Ptr[int64](10),
	// 				CurrentIterationFilesExported: to.Ptr[int64](5),
	// 				CurrentIterationFilesFailed: to.Ptr[int64](1),
	// 				CurrentIterationMiBDiscovered: to.Ptr[int64](4000),
	// 				CurrentIterationMiBExported: to.Ptr[int64](500),
	// 				ExportIterationCount: to.Ptr[int32](100),
	// 				LastStartedTimeUTC: to.Ptr(time.Date(2024, time.April, 21, 17, 25, 43, 511000000, time.UTC)),
	// 				LastSuccessfulIterationCompletionTimeUTC: to.Ptr(time.Date(2024, time.April, 21, 19, 28, 43, 511000000, time.UTC)),
	// 				State: to.Ptr(armstoragecache.AutoExportStatusTypeInProgress),
	// 				StatusMessage: to.Ptr("Auto Export is in progress"),
	// 				TotalFilesExported: to.Ptr[int64](1000000),
	// 				TotalFilesFailed: to.Ptr[int64](5),
	// 				TotalMiBExported: to.Ptr[int64](10000),
	// 			},
	// 		},
	// 	},
	// }
}
