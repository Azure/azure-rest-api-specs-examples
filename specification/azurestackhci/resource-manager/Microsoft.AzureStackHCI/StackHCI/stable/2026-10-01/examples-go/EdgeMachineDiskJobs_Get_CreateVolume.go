package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-10-01/EdgeMachineDiskJobs_Get_CreateVolume.json
func ExampleEdgeMachineDiskJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEdgeMachineDiskJobsClient().Get(ctx, "test-rg", "EdgeMachine01", "disk-001", "CreateVolume", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.EdgeMachineDiskJobsClientGetResponse{
	// 	EdgeMachineDiskJob: armazurestackhci.EdgeMachineDiskJob{
	// 		Properties: &armazurestackhci.CreateVolumeJobProperties{
	// 			JobType: to.Ptr(armazurestackhci.DiskJobTypeCreateVolume),
	// 			DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 			JobID: to.Ptr("job-vol-001"),
	// 			StartTimeUTC: to.Ptr(time.Date(2026, time.January, 15, 10, 0, 0, 0, time.UTC)),
	// 			EndTimeUTC: to.Ptr(time.Date(2026, time.January, 15, 10, 10, 0, 0, time.UTC)),
	// 			Status: to.Ptr(armazurestackhci.JobStatusSucceeded),
	// 			SizeInBytes: to.Ptr("107374182400"),
	// 			Path: to.Ptr("/data"),
	// 			FileSystem: to.Ptr(armazurestackhci.DiskFileSystemTypeExt4),
	// 			CreatedVolumeID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/edgeMachines/EdgeMachine01/volumes/vol-001"),
	// 			ReportedProperties: &armazurestackhci.DiskJobReportedProperties{
	// 				PercentComplete: to.Ptr[int32](100),
	// 				ValidationStatus: &armazurestackhci.DiskActionStatus{
	// 					Status: to.Ptr("Success"),
	// 					Steps: []*armazurestackhci.DiskDeploymentStep{
	// 						{
	// 							Name: to.Ptr("ValidateVolumeCreation"),
	// 							Description: to.Ptr("Validate volume creation prerequisites."),
	// 							FullStepIndex: to.Ptr("0"),
	// 							StartTimeUTC: to.Ptr("2026-01-15T10:00:00.000Z"),
	// 							EndTimeUTC: to.Ptr("2026-01-15T10:01:00.000Z"),
	// 							Status: to.Ptr("Success"),
	// 							Steps: []*armazurestackhci.DiskDeploymentStep{
	// 							},
	// 						},
	// 					},
	// 				},
	// 				DeploymentStatus: &armazurestackhci.DiskActionStatus{
	// 					Status: to.Ptr("Success"),
	// 					Steps: []*armazurestackhci.DiskDeploymentStep{
	// 						{
	// 							Name: to.Ptr("CreateVolume"),
	// 							Description: to.Ptr("Create volume on disk."),
	// 							FullStepIndex: to.Ptr("0"),
	// 							StartTimeUTC: to.Ptr("2026-01-15T10:01:00.000Z"),
	// 							EndTimeUTC: to.Ptr("2026-01-15T10:10:00.000Z"),
	// 							Status: to.Ptr("Success"),
	// 							Steps: []*armazurestackhci.DiskDeploymentStep{
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/edgeMachines/EdgeMachine01/disks/disk-001/jobs/CreateVolume"),
	// 		Name: to.Ptr("CreateVolume"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/edgeMachines/disks/jobs"),
	// 	},
	// }
}
