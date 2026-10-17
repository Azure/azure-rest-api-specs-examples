package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-10-01/EdgeMachineDiskPrivilegedJobs_List.json
func ExampleEdgeMachineDiskPrivilegedJobsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEdgeMachineDiskPrivilegedJobsClient().NewListPager("test-rg", "EdgeMachine01", "disk-001", nil)
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
		// page = armazurestackhci.EdgeMachineDiskPrivilegedJobsClientListResponse{
		// 	EdgeMachineDiskPrivilegedJobListResult: armazurestackhci.EdgeMachineDiskPrivilegedJobListResult{
		// 		Value: []*armazurestackhci.EdgeMachineDiskPrivilegedJob{
		// 			{
		// 				Properties: &armazurestackhci.DeleteVolumeJobProperties{
		// 					JobType: to.Ptr(armazurestackhci.PrivilegedJobTypeDeleteVolume),
		// 					DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
		// 					ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
		// 					JobID: to.Ptr("job-pvol-001"),
		// 					StartTimeUTC: to.Ptr(time.Date(2026, time.January, 15, 10, 0, 0, 0, time.UTC)),
		// 					EndTimeUTC: to.Ptr(time.Date(2026, time.January, 15, 10, 5, 0, 0, time.UTC)),
		// 					Status: to.Ptr(armazurestackhci.JobStatusSucceeded),
		// 					VolumePath: to.Ptr("/data"),
		// 					ConfirmDeletion: to.Ptr(true),
		// 					ReportedProperties: &armazurestackhci.DiskJobReportedProperties{
		// 						PercentComplete: to.Ptr[int32](100),
		// 						ValidationStatus: &armazurestackhci.DiskActionStatus{
		// 							Status: to.Ptr("Success"),
		// 							Steps: []*armazurestackhci.DiskDeploymentStep{
		// 							},
		// 						},
		// 						DeploymentStatus: &armazurestackhci.DiskActionStatus{
		// 							Status: to.Ptr("Success"),
		// 							Steps: []*armazurestackhci.DiskDeploymentStep{
		// 							},
		// 						},
		// 					},
		// 				},
		// 				ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/edgeMachines/EdgeMachine01/disks/disk-001/privilegedJobs/DeleteVolume"),
		// 				Name: to.Ptr("DeleteVolume"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/edgeMachines/disks/privilegedJobs"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
