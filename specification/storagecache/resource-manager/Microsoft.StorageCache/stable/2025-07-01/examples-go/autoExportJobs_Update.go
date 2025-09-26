package armstoragecache_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d18f451b13796bded0808d508c6297f2227271d5/specification/storagecache/resource-manager/Microsoft.StorageCache/stable/2025-07-01/examples/autoExportJobs_Update.json
func ExampleAutoExportJobsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragecache.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAutoExportJobsClient().BeginUpdate(ctx, "scgroup", "fs1", "job1", armstoragecache.AutoExportJobUpdate{
		Tags: map[string]*string{
			"Dept": to.Ptr("ContosoAds"),
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
	// res.AutoExportJob = armstoragecache.AutoExportJob{
	// 	Name: to.Ptr("job1"),
	// 	Type: to.Ptr("Microsoft.StorageCache/amlFilesystem/autoExportJob"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.StorageCache/amlfilesystems/fs1/autoExportJob/job1"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"Dept": to.Ptr("ContosoAds"),
	// 	},
	// 	Properties: &armstoragecache.AutoExportJobProperties{
	// 		AdminStatus: to.Ptr(armstoragecache.AutoExportJobAdminStatusEnable),
	// 		AutoExportPrefixes: []*string{
	// 			to.Ptr("/")},
	// 			ProvisioningState: to.Ptr(armstoragecache.AutoExportJobProvisioningStateTypeSucceeded),
	// 			Status: &armstoragecache.AutoExportJobPropertiesStatus{
	// 				CurrentIterationFilesDiscovered: to.Ptr[int64](10),
	// 				CurrentIterationFilesExported: to.Ptr[int64](5),
	// 				CurrentIterationFilesFailed: to.Ptr[int64](1),
	// 				CurrentIterationMiBDiscovered: to.Ptr[int64](4000),
	// 				CurrentIterationMiBExported: to.Ptr[int64](500),
	// 				ExportIterationCount: to.Ptr[int32](100),
	// 				LastStartedTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-21T17:25:43.511Z"); return t}()),
	// 				LastSuccessfulIterationCompletionTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-04-21T19:28:43.511Z"); return t}()),
	// 				State: to.Ptr(armstoragecache.AutoExportStatusTypeInProgress),
	// 				StatusMessage: to.Ptr("Auto Export is in progress"),
	// 				TotalFilesExported: to.Ptr[int64](1000000),
	// 				TotalFilesFailed: to.Ptr[int64](5),
	// 				TotalMiBExported: to.Ptr[int64](10000),
	// 			},
	// 		},
	// 	}
}
