package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/EdgeMachineJobs_CreateOrUpdate_ProvisionOs.json
func ExampleEdgeMachineJobsClient_BeginCreateOrUpdate_edgeMachineJobsCreateOrUpdateProvisionOS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEdgeMachineJobsClient().BeginCreateOrUpdate(ctx, "ArcInstance-rg", "machine1", "ProvisionOs", armazurestackhci.EdgeMachineJob{
		Properties: &armazurestackhci.ProvisionOsJobProperties{
			ProvisioningRequest: &armazurestackhci.ProvisioningRequest{
				OSProfile: &armazurestackhci.OsProvisionProfile{
					OSName:          to.Ptr("AzureLinux"),
					OSType:          to.Ptr("AzureLinux"),
					OSVersion:       to.Ptr("3.0"),
					OSImageLocation: to.Ptr("https://aka.ms/aep/azlinux3.0"),
					VsrVersion:      to.Ptr("1.0.0"),
					ImageHash:       to.Ptr("sha256:a1b2c3d4e5f6789012345678901234567890abcdef1234567890abcdef123456"),
					GpgPubKey:       to.Ptr("LS0tLS1CRUdJTiBQR1AgUFVCTElDIEtFWSBCTE9DSy0tLS0tXG5WZXJzaW9uOiBHbnVQRyB2MlxuXG5tUUVOQkZYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYXG4tLS0tLUVORCBQR1AgUFVCTElDIEtFWSBCTE9DSy0tLS0t"),
					OperationType:   to.Ptr(armazurestackhci.OSOperationTypeProvision),
				},
				UserDetails: []*armazurestackhci.UserDetails{
					{
						UserName:       to.Ptr("edgeuser"),
						SecretType:     to.Ptr(armazurestackhci.SecretTypeKeyVault),
						SecretLocation: to.Ptr("https://bhukumar-test-kv.vault.azure.net/secrets/edgeuser/7b2d7db11bad4e1599cb6a0f4d2b2e00"),
						SSHPubKey: []*string{
							to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC7... edgeuser@example.com"),
						},
					},
				},
				OnboardingConfiguration: &armazurestackhci.OnboardingConfiguration{
					ResourceID:          to.Ptr("/subscriptions/ff0aa6da-20f8-44fe-9aee-381c8e8a4aeb/resourceGroups/bhukumar-test-rg/providers/Microsoft.HybridCompute/machines/bkumar-t1"),
					Location:            to.Ptr("eastus"),
					TenantID:            to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
					ArcVirtualMachineID: to.Ptr("634b9db8-83e1-46ed-b391-c1614e2d0097"),
					Type:                to.Ptr(armazurestackhci.OnboardingResourceTypeHybridComputeMachine),
				},
				DeviceConfiguration: &armazurestackhci.TargetDeviceConfiguration{
					Network: &armazurestackhci.NetworkConfiguration{
						NetworkAdapters: []*armazurestackhci.NetworkAdapter{
							{
								IPAssignmentType: to.Ptr(armazurestackhci.IPAssignmentTypeAutomatic),
								IPAddress:        to.Ptr(""),
								IPAddressRange: &armazurestackhci.IPAddressRange{
									StartIP: to.Ptr(""),
									EndIP:   to.Ptr(""),
								},
								Gateway:    to.Ptr(""),
								SubnetMask: to.Ptr(""),
								DNSAddressArray: []*string{
									to.Ptr("8.8.8.8"),
								},
								VlanID: to.Ptr("0"),
							},
						},
					},
					HostName: to.Ptr("634b9db8-83e1-46ed-b391-c1614e2d0097"),
					WebProxy: &armazurestackhci.WebProxyConfiguration{
						ConnectionURI: to.Ptr("https://microsoft.com/a"),
						Port:          to.Ptr(""),
						BypassList:    []*string{},
					},
					Time: &armazurestackhci.TimeConfiguration{
						PrimaryTimeServer:   to.Ptr(""),
						SecondaryTimeServer: to.Ptr(""),
						TimeZone:            to.Ptr("UTC"),
					},
					Storage: &armazurestackhci.StorageConfiguration{
						PartitionSize: to.Ptr("30GB"),
					},
				},
				Target:              to.Ptr(armazurestackhci.ProvisioningOsTypeAzureLinux),
				CustomConfiguration: to.Ptr("eyJjdXN0b21Db25maWciOiAiZXhhbXBsZSBiYXNlNjQgZW5jb2RlZCBjb25maWcifQ=="),
			},
			JobType:        to.Ptr(armazurestackhci.EdgeMachineJobTypeProvisionOs),
			DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
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
	// 		ID: to.Ptr("/subscriptions/ff0aa6da-20f8-44fe-9aee-381c8e8a4aeb/resourceGroups/bhukumar-test-rg/providers/Microsoft.AzureStackHCI/EdgeMachines/em1/jobs/ProvisionOs"),
	// 		Name: to.Ptr("ProvisionOs"),
	// 		Type: to.Ptr("microsoft.azurestackhci/edgemachines/jobs"),
	// 		Properties: &armazurestackhci.ProvisionOsJobProperties{
	// 			ProvisioningRequest: &armazurestackhci.ProvisioningRequest{
	// 				OSProfile: &armazurestackhci.OsProvisionProfile{
	// 					OSName: to.Ptr("AzureLinux"),
	// 					OSType: to.Ptr("AzureLinux"),
	// 					OSVersion: to.Ptr("3.0"),
	// 					OSImageLocation: to.Ptr("https://aka.ms/aep/azlinux3.0"),
	// 					VsrVersion: to.Ptr("1.0.0"),
	// 					ImageHash: to.Ptr("sha256:a1b2c3d4e5f6789012345678901234567890abcdef1234567890abcdef123456"),
	// 					GpgPubKey: to.Ptr("LS0tLS1CRUdJTiBQR1AgUFVCTElDIEtFWSBCTE9DSy0tLS0tXG5WZXJzaW9uOiBHbnVQRyB2MlxuXG5tUUVOQkZYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYWFhYXG4tLS0tLUVORCBQR1AgUFVCTElDIEtFWSBCTE9DSy0tLS0t"),
	// 					OperationType: to.Ptr(armazurestackhci.OSOperationTypeProvision),
	// 				},
	// 				UserDetails: []*armazurestackhci.UserDetails{
	// 					{
	// 						UserName: to.Ptr("edgeuser"),
	// 						SecretType: to.Ptr(armazurestackhci.SecretTypeKeyVault),
	// 						SecretLocation: to.Ptr("https://bhukumar-test-kv.vault.azure.net/secrets/edgeuser/7b2d7db11bad4e1599cb6a0f4d2b2e00"),
	// 						SSHPubKey: []*string{
	// 							to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQC7... edgeuser@example.com"),
	// 						},
	// 					},
	// 				},
	// 				OnboardingConfiguration: &armazurestackhci.OnboardingConfiguration{
	// 					ResourceID: to.Ptr("/subscriptions/ff0aa6da-20f8-44fe-9aee-381c8e8a4aeb/resourceGroups/bhukumar-test-rg/providers/Microsoft.HybridCompute/machines/bkumar-t1"),
	// 					Location: to.Ptr("eastus"),
	// 					TenantID: to.Ptr("72f988bf-86f1-41af-91ab-2d7cd011db47"),
	// 					ArcVirtualMachineID: to.Ptr("634b9db8-83e1-46ed-b391-c1614e2d0097"),
	// 					Type: to.Ptr(armazurestackhci.OnboardingResourceTypeHybridComputeMachine),
	// 				},
	// 				DeviceConfiguration: &armazurestackhci.TargetDeviceConfiguration{
	// 					Network: &armazurestackhci.NetworkConfiguration{
	// 						NetworkAdapters: []*armazurestackhci.NetworkAdapter{
	// 							{
	// 								IPAssignmentType: to.Ptr(armazurestackhci.IPAssignmentTypeAutomatic),
	// 								IPAddress: to.Ptr(""),
	// 								IPAddressRange: &armazurestackhci.IPAddressRange{
	// 									StartIP: to.Ptr(""),
	// 									EndIP: to.Ptr(""),
	// 								},
	// 								Gateway: to.Ptr(""),
	// 								SubnetMask: to.Ptr(""),
	// 								DNSAddressArray: []*string{
	// 									to.Ptr("8.8.8.8"),
	// 								},
	// 								VlanID: to.Ptr("0"),
	// 							},
	// 						},
	// 					},
	// 					HostName: to.Ptr("634b9db8-83e1-46ed-b391-c1614e2d0097"),
	// 					WebProxy: &armazurestackhci.WebProxyConfiguration{
	// 						ConnectionURI: to.Ptr("https://microsoft.com/a"),
	// 						Port: to.Ptr(""),
	// 						BypassList: []*string{
	// 						},
	// 					},
	// 					Time: &armazurestackhci.TimeConfiguration{
	// 						PrimaryTimeServer: to.Ptr(""),
	// 						SecondaryTimeServer: to.Ptr(""),
	// 						TimeZone: to.Ptr("UTC"),
	// 					},
	// 					Storage: &armazurestackhci.StorageConfiguration{
	// 						PartitionSize: to.Ptr("30GB"),
	// 					},
	// 				},
	// 				Target: to.Ptr(armazurestackhci.ProvisioningOsTypeAzureLinux),
	// 				CustomConfiguration: to.Ptr("eyJjdXN0b21Db25maWciOiAiZXhhbXBsZSBiYXNlNjQgZW5jb2RlZCBjb25maWcifQ=="),
	// 			},
	// 			ReportedProperties: &armazurestackhci.ProvisionOsReportedProperties{
	// 				PercentComplete: to.Ptr[int32](100),
	// 				DeploymentStatus: &armazurestackhci.EceActionStatus{
	// 					Status: to.Ptr("Succeeded"),
	// 					Steps: []*armazurestackhci.DeploymentStep{
	// 						{
	// 							Name: to.Ptr("Initialize"),
	// 							StartTimeUTC: to.Ptr("03/11/2025 15:11:01"),
	// 							EndTimeUTC: to.Ptr("03/11/2025 15:11:01"),
	// 							Status: to.Ptr("Succeeded"),
	// 						},
	// 						{
	// 							Name: to.Ptr("ResolveConfiguration"),
	// 							StartTimeUTC: to.Ptr("03/11/2025 15:11:02"),
	// 							EndTimeUTC: to.Ptr("03/11/2025 15:11:02"),
	// 							Status: to.Ptr("Succeeded"),
	// 						},
	// 						{
	// 							Name: to.Ptr("DownloadOSArtifacts"),
	// 							StartTimeUTC: to.Ptr("03/11/2025 15:11:03"),
	// 							EndTimeUTC: to.Ptr("03/11/2025 15:11:23"),
	// 							Status: to.Ptr("Succeeded"),
	// 						},
	// 						{
	// 							Name: to.Ptr("SetupHybridComputeRP"),
	// 							StartTimeUTC: to.Ptr("03/11/2025 15:11:23"),
	// 							EndTimeUTC: to.Ptr("03/11/2025 15:13:24"),
	// 							Status: to.Ptr("Succeeded"),
	// 						},
	// 						{
	// 							Name: to.Ptr("PrepareTargetOSProvisioning"),
	// 							Description: to.Ptr("NotStarted"),
	// 							StartTimeUTC: to.Ptr("01/01/0001 00:00:00"),
	// 							EndTimeUTC: to.Ptr("01/01/0001 00:00:00"),
	// 							Status: to.Ptr("Succeeded"),
	// 						},
	// 						{
	// 							Name: to.Ptr("InitiateTargetOSProvisioning"),
	// 							Description: to.Ptr("NotStarted"),
	// 							StartTimeUTC: to.Ptr("01/01/0001 00:00:00"),
	// 							EndTimeUTC: to.Ptr("01/01/0001 00:00:00"),
	// 							Status: to.Ptr("Succeeded"),
	// 						},
	// 						{
	// 							Name: to.Ptr("BootIntoTargetOS"),
	// 							Description: to.Ptr("NotStarted"),
	// 							StartTimeUTC: to.Ptr("01/01/0001 00:00:00"),
	// 							EndTimeUTC: to.Ptr("01/01/0001 00:00:00"),
	// 							Status: to.Ptr("Succeeded"),
	// 						},
	// 					},
	// 				},
	// 				ValidationStatus: &armazurestackhci.EceActionStatus{
	// 				},
	// 			},
	// 			JobType: to.Ptr(armazurestackhci.EdgeMachineJobTypeProvisionOs),
	// 			DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 			JobID: to.Ptr("8f94f7e2-00c8-4751-986c-3802143c7900"),
	// 			StartTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-12T10:42:06.0716967Z"); return t}()),
	// 			EndTimeUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-03-12T10:45:32.7360665Z"); return t}()),
	// 			Status: to.Ptr(armazurestackhci.JobStatusDeploymentSuccess),
	// 		},
	// 	},
	// }
}
