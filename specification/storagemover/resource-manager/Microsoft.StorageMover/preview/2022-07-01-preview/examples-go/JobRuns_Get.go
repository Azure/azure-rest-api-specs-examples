package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3d7a3848106b831a4a7f46976fe38aa605c4f44d/specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2022-07-01-preview/examples/JobRuns_Get.json
func ExampleJobRunsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragemover.NewJobRunsClient("11111111-2222-3333-4444-555555555555", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "examples-rg", "examples-storageMoverName", "examples-projectName", "examples-jobDefinitionName", "examples-jobRunName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobRun = armstoragemover.JobRun{
	// 	Name: to.Ptr("examples-jobRunName"),
	// 	Type: to.Ptr("Microsoft.StorageMover/storageMovers/projects/jobDefinitions/jobRuns"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/projects/examples-projectName/jobDefinitions/examples-jobDefinitionName/jobRuns/examples-jobRunName"),
	// 	Properties: &armstoragemover.JobRunProperties{
	// 		AgentName: to.Ptr("migration-agent"),
	// 		AgentResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/agents/migration-agent"),
	// 		BytesExcluded: to.Ptr[int64](995116277760),
	// 		BytesFailed: to.Ptr[int64](5116277760),
	// 		BytesNoTransferNeeded: to.Ptr[int64](2995116277760),
	// 		BytesScanned: to.Ptr[int64](49951162777600),
	// 		BytesTransferred: to.Ptr[int64](1995116277760),
	// 		BytesUnsupported: to.Ptr[int64](495116277760),
	// 		ExecutionStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:01:01.1075056Z"); return t}()),
	// 		ItemsExcluded: to.Ptr[int64](50),
	// 		ItemsFailed: to.Ptr[int64](3),
	// 		ItemsNoTransferNeeded: to.Ptr[int64](150),
	// 		ItemsScanned: to.Ptr[int64](351),
	// 		ItemsTransferred: to.Ptr[int64](100),
	// 		ItemsUnsupported: to.Ptr[int64](27),
	// 		JobDefinitionProperties: map[string]any{
	// 		},
	// 		LastStatusUpdate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-01T01:31:01.1075056Z"); return t}()),
	// 		ScanStatus: to.Ptr(armstoragemover.JobRunScanStatusScanning),
	// 		SourceName: to.Ptr("sourceEndpoint"),
	// 		SourceProperties: map[string]any{
	// 		},
	// 		SourceResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/sourceEndpoint"),
	// 		Status: to.Ptr(armstoragemover.JobRunStatusRunning),
	// 		TargetName: to.Ptr("targetEndpoint"),
	// 		TargetProperties: map[string]any{
	// 		},
	// 		TargetResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.StorageMover/storageMovers/examples-storageMoverName/endpoints/targetEndpoint"),
	// 	},
	// }
}
