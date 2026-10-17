package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-10-01/EdgeMachineNetworkAdapterJobs_Get.json
func ExampleEdgeMachineNetworkAdapterJobsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEdgeMachineNetworkAdapterJobsClient().Get(ctx, "test-rg", "EdgeMachine01", "ethernet0", "ApplyConfiguration", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.EdgeMachineNetworkAdapterJobsClientGetResponse{
	// 	EdgeMachineNetworkAdapterJob: armazurestackhci.EdgeMachineNetworkAdapterJob{
	// 		Properties: &armazurestackhci.ApplyConfigurationNetworkAdapterJobProperties{
	// 			JobType: to.Ptr(armazurestackhci.NetworkAdapterJobTypeApplyConfiguration),
	// 			DeploymentMode: to.Ptr(armazurestackhci.DeploymentModeDeploy),
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 			JobID: to.Ptr("job-12345"),
	// 			StartTimeUTC: to.Ptr(time.Date(2026, time.January, 15, 10, 0, 0, 0, time.UTC)),
	// 			EndTimeUTC: to.Ptr(time.Date(2026, time.January, 15, 10, 5, 0, 0, time.UTC)),
	// 			Status: to.Ptr(armazurestackhci.JobStatusSucceeded),
	// 			TargetConfiguration: &armazurestackhci.NetworkAdapterConfiguration{
	// 				AdapterName: to.Ptr("ethernet0"),
	// 				Ip4Address: to.Ptr("192.168.1.100"),
	// 				SubnetMask: to.Ptr("255.255.255.0"),
	// 				DefaultGateway: to.Ptr("192.168.1.1"),
	// 				DNSServers: []*string{
	// 					to.Ptr("8.8.8.8"),
	// 					to.Ptr("8.8.4.4"),
	// 				},
	// 				IPInterfaceType: to.Ptr(armazurestackhci.IPInterfaceTypeStatic),
	// 				VlanID: to.Ptr[int32](100),
	// 				InterfaceState: to.Ptr(armazurestackhci.InterfaceStateUp),
	// 			},
	// 			ReportedProperties: &armazurestackhci.NetworkAdapterJobReportedProperties{
	// 				PercentComplete: to.Ptr[int32](100),
	// 				ValidationStatus: &armazurestackhci.NetworkAdapterActionStatus{
	// 					Status: to.Ptr("Success"),
	// 					Steps: []*armazurestackhci.NetworkAdapterDeploymentStep{
	// 						{
	// 							Name: to.Ptr("ValidateConfiguration"),
	// 							Description: to.Ptr("Validate network adapter configuration."),
	// 							FullStepIndex: to.Ptr("0"),
	// 							StartTimeUTC: to.Ptr("2026-01-15T10:00:00.000Z"),
	// 							EndTimeUTC: to.Ptr("2026-01-15T10:01:00.000Z"),
	// 							Status: to.Ptr("Success"),
	// 							Steps: []*armazurestackhci.NetworkAdapterDeploymentStep{
	// 							},
	// 						},
	// 					},
	// 				},
	// 				DeploymentStatus: &armazurestackhci.NetworkAdapterActionStatus{
	// 					Status: to.Ptr("Success"),
	// 					Steps: []*armazurestackhci.NetworkAdapterDeploymentStep{
	// 						{
	// 							Name: to.Ptr("ApplyNetworkConfiguration"),
	// 							Description: to.Ptr("Apply network adapter configuration."),
	// 							FullStepIndex: to.Ptr("0"),
	// 							StartTimeUTC: to.Ptr("2026-01-15T10:01:00.000Z"),
	// 							EndTimeUTC: to.Ptr("2026-01-15T10:05:00.000Z"),
	// 							Status: to.Ptr("Success"),
	// 							Steps: []*armazurestackhci.NetworkAdapterDeploymentStep{
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/edgeMachines/EdgeMachine01/networkAdapters/ethernet0/jobs/ApplyConfiguration"),
	// 		Name: to.Ptr("ApplyConfiguration"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/edgeMachines/networkAdapters/jobs"),
	// 	},
	// }
}
