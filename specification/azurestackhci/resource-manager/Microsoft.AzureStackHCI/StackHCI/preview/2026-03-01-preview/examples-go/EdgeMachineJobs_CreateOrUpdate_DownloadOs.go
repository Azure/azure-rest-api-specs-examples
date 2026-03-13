package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/EdgeMachineJobs_CreateOrUpdate_DownloadOs.json
func ExampleEdgeMachineJobsClient_BeginCreateOrUpdate_edgeMachineJobsCreateOrUpdateDownloadOS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEdgeMachineJobsClient().BeginCreateOrUpdate(ctx, "ArcInstance-rg", "machine1", "DownloadOs", armazurestackhci.EdgeMachineJob{
		Properties: &armazurestackhci.DownloadOsJobProperties{
			JobType:        to.Ptr(armazurestackhci.EdgeMachineJobTypeDownloadOs),
			DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
			DownloadRequest: &armazurestackhci.DownloadRequest{
				Target: to.Ptr(armazurestackhci.ProvisioningOsTypeAzureLinux),
				OSProfile: &armazurestackhci.DownloadOsProfile{
					OSName:          to.Ptr("AzureLinux"),
					OSType:          to.Ptr("AzureLinux"),
					OSVersion:       to.Ptr("3.0"),
					OSImageLocation: to.Ptr("https://aka.ms/aep/azlinux3.0"),
					VsrVersion:      to.Ptr("1.0.0"),
					ImageHash:       to.Ptr("sha256:a8b1c2d3e4f5a6b7c8d9e0f1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b1"),
					GpgPubKey:       to.Ptr("LS0tLS1CRUdJTiBQR1AgUFVCTElDIEtFWSBCTE9DSy0tLS0tXG5WZXJzaW9uOiBHbnVQRyB2MlxuXG5tUUVOQkZYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYXG4tLS0tLUVORCBQR1AgUFVCTElDIEtFWSBCTE9DSy0tLS0t"),
				},
			},
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
	// res = armazurestackhci.EdgeMachineJobsClientCreateOrUpdateResponse{
	// 	EdgeMachineJob: &armazurestackhci.EdgeMachineJob{
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.AzureStackHCI/edgeMachines/machine1/jobs/DownloadOs"),
	// 		Name: to.Ptr("DownloadOs"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/edgeMachines/jobs"),
	// 		Properties: &armazurestackhci.DownloadOsJobProperties{
	// 			JobType: to.Ptr(armazurestackhci.EdgeMachineJobTypeDownloadOs),
	// 			DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
	// 			DownloadRequest: &armazurestackhci.DownloadRequest{
	// 				Target: to.Ptr(armazurestackhci.ProvisioningOsTypeAzureLinux),
	// 				OSProfile: &armazurestackhci.DownloadOsProfile{
	// 					OSName: to.Ptr("AzureLinux"),
	// 					OSType: to.Ptr("AzureLinux"),
	// 					OSVersion: to.Ptr("3.0"),
	// 					OSImageLocation: to.Ptr("https://aka.ms/aep/azlinux3.0"),
	// 					VsrVersion: to.Ptr("1.0.0"),
	// 					ImageHash: to.Ptr("sha256:a8b1c2d3e4f5a6b7c8d9e0f1a2b3c4d5e6f7a8b9c0d1e2f3a4b5c6d7e8f9a0b1"),
	// 					GpgPubKey: to.Ptr("LS0tLS1CRUdJTiBQR1AgUFVCTElDIEtFWSBCTE9DSy0tLS0tXG5WZXJzaW9uOiBHbnVQRyB2MlxuXG5tUUVOQkZYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYXG4tLS0tLUVORCBQR1AgUFVCTElDIEtFWSBCTE9DSy0tLS0t"),
	// 				},
	// 			},
	// 			ReportedProperties: &armazurestackhci.ProvisionOsReportedProperties{
	// 				PercentComplete: to.Ptr[int32](75),
	// 				ValidationStatus: &armazurestackhci.EceActionStatus{
	// 					Status: to.Ptr("Succeeded"),
	// 					Steps: []*armazurestackhci.DeploymentStep{
	// 						{
	// 							Name: to.Ptr("ValidateOsProfile"),
	// 							StartTimeUTC: to.Ptr("2024-08-19T16:45:01"),
	// 							EndTimeUTC: to.Ptr("2024-08-19T16:45:05"),
	// 							Status: to.Ptr("Succeeded"),
	// 						},
	// 						{
	// 							Name: to.Ptr("ValidateImageLocation"),
	// 							StartTimeUTC: to.Ptr("2024-08-19T16:45:05"),
	// 							EndTimeUTC: to.Ptr("2024-08-19T16:45:10"),
	// 							Status: to.Ptr("Succeeded"),
	// 						},
	// 					},
	// 				},
	// 				DeploymentStatus: &armazurestackhci.EceActionStatus{
	// 					Status: to.Ptr("InProgress"),
	// 					Steps: []*armazurestackhci.DeploymentStep{
	// 						{
	// 							Name: to.Ptr("Initialize"),
	// 							StartTimeUTC: to.Ptr("2024-08-19T16:45:15"),
	// 							EndTimeUTC: to.Ptr("2024-08-19T16:45:16"),
	// 							Status: to.Ptr("Succeeded"),
	// 						},
	// 						{
	// 							Name: to.Ptr("DownloadOSPackage"),
	// 							StartTimeUTC: to.Ptr("2024-08-19T16:45:16"),
	// 							EndTimeUTC: to.Ptr("2024-08-19T16:50:30"),
	// 							Status: to.Ptr("Succeeded"),
	// 						},
	// 						{
	// 							Name: to.Ptr("VerifyPackageIntegrity"),
	// 							StartTimeUTC: to.Ptr("2024-08-19T16:50:30"),
	// 							EndTimeUTC: to.Ptr(""),
	// 							Status: to.Ptr("InProgress"),
	// 						},
	// 						{
	// 							Name: to.Ptr("StorePackageLocally"),
	// 							StartTimeUTC: to.Ptr(""),
	// 							EndTimeUTC: to.Ptr(""),
	// 							Status: to.Ptr("NotStarted"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateInProgress),
	// 			JobID: to.Ptr("download-os-job-abc123def456"),
	// 			StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-19T16:45:00.000Z"); return t}()),
	// 			Status: to.Ptr(armazurestackhci.JobStatusDeploymentInProgress),
	// 		},
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-19T16:44:55.167Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user1"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-19T16:50:30.167Z"); return t}()),
	// 		},
	// 	},
	// }
}
