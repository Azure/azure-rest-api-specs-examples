package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/EdgeMachines_CreateOrUpdate.json
func ExampleEdgeMachinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewEdgeMachinesClient().BeginCreateOrUpdate(ctx, "ArcInstance-rg", "machine-1", armazurestackhci.EdgeMachine{
		Properties: &armazurestackhci.EdgeMachineProperties{
			ArcMachineResourceGroupID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg"),
			ArcMachineResourceID:      to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1"),
		},
		Location: to.Ptr("eastus"),
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
	// res = armazurestackhci.EdgeMachinesClientCreateOrUpdateResponse{
	// 	EdgeMachine: &armazurestackhci.EdgeMachine{
	// 		Properties: &armazurestackhci.EdgeMachineProperties{
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 			ArcMachineResourceGroupID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg"),
	// 			ArcMachineResourceID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1"),
	// 			DevicePoolResourceID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.AzureStackHCI/devicePools/pool-1"),
	// 			ReportedProperties: &armazurestackhci.EdgeMachineReportedProperties{
	// 				NetworkProfile: &armazurestackhci.EdgeMachineNetworkProfile{
	// 					NicDetails: []*armazurestackhci.EdgeMachineNicDetail{
	// 						{
	// 							AdapterName: to.Ptr("vmanagement"),
	// 							InterfaceDescription: to.Ptr("Hyper-V Virtual Ethernet Adapter"),
	// 							ComponentID: to.Ptr("vms_mp"),
	// 							DriverVersion: to.Ptr("10.0.25398.1"),
	// 							Ip4Address: to.Ptr("192.168.200.92"),
	// 							SubnetMask: to.Ptr("255.255.255.0"),
	// 							DefaultGateway: to.Ptr("192.168.200.1"),
	// 							DNSServers: []*string{
	// 								to.Ptr("192.168.200.222"),
	// 							},
	// 							DefaultIsolationID: to.Ptr("0"),
	// 							Slot: to.Ptr("0"),
	// 							MacAddress: to.Ptr("000000000041"),
	// 							NicType: to.Ptr("Virtual"),
	// 							VlanID: to.Ptr("0"),
	// 							NicStatus: to.Ptr("Up"),
	// 						},
	// 					},
	// 					SwitchDetails: []*armazurestackhci.SwitchDetail{
	// 						{
	// 							SwitchName: to.Ptr("vmanagement"),
	// 							SwitchType: to.Ptr("External"),
	// 						},
	// 					},
	// 				},
	// 				OSProfile: &armazurestackhci.OsProfile{
	// 					BootType: to.Ptr("UEFI"),
	// 					AssemblyVersion: to.Ptr("2402.1"),
	// 				},
	// 				HardwareProfile: &armazurestackhci.HardwareProfile{
	// 					CPUCores: to.Ptr[int64](4),
	// 					CPUSockets: to.Ptr[int64](14),
	// 					MemoryCapacityInGb: to.Ptr[int64](25),
	// 					SerialNumber: to.Ptr("5855-7104-5665-0328-3864-8600-21"),
	// 				},
	// 				SbeDeploymentPackageInfo: &armazurestackhci.SbeDeploymentPackageInfo{
	// 					Code: to.Ptr("NewerThanLatestPublished"),
	// 					Message: to.Ptr("The SBE package at path 'C:\\SBE' with version 4.1.2312.10 is published later than the latest SBE manifest published for online discovery. "),
	// 					SbeManifest: to.Ptr("PEFwcGxpY2Fi"),
	// 				},
	// 				ExtensionProfile: &armazurestackhci.ExtensionProfile{
	// 					Extensions: []*armazurestackhci.HciEdgeDeviceArcExtension{
	// 						{
	// 							ExtensionName: to.Ptr("DeviceManagementExtension"),
	// 							State: to.Ptr(armazurestackhci.ArcExtensionState("Connected")),
	// 							ExtensionResourceID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1/Extensions/DeviceManagementExtension"),
	// 							TypeHandlerVersion: to.Ptr("1.10.3"),
	// 							ManagedBy: to.Ptr(armazurestackhci.ExtensionManagedByAzure),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("eastus"),
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.AzureStackHCI/edgeMachines/machine-1"),
	// 		Name: to.Ptr("machine-1"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/edgeMachines"),
	// 		SystemData: &armazurestackhci.SystemData{
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-14T10:46:55.167Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-11-14T10:46:55.167Z"); return t}()),
	// 		},
	// 	},
	// }
}
