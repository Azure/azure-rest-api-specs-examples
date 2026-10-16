package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-10-01/EdgeMachines_ListByResourceGroup.json
func ExampleEdgeMachinesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEdgeMachinesClient().NewListByResourceGroupPager("ArcInstance-rg", nil)
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
		// page = armazurestackhci.EdgeMachinesClientListByResourceGroupResponse{
		// 	EdgeMachineListResult: armazurestackhci.EdgeMachineListResult{
		// 		Value: []*armazurestackhci.EdgeMachine{
		// 			{
		// 				Properties: &armazurestackhci.EdgeMachineProperties{
		// 					ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
		// 					ArcMachineResourceGroupID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg"),
		// 					ArcMachineResourceID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1"),
		// 					DevicePoolResourceID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.AzureStackHCI/devicePools/pool-1"),
		// 					ReportedProperties: &armazurestackhci.EdgeMachineReportedProperties{
		// 						NetworkProfile: &armazurestackhci.EdgeMachineNetworkProfile{
		// 							NicDetails: []*armazurestackhci.EdgeMachineNicDetail{
		// 								{
		// 									AdapterName: to.Ptr("vmanagement"),
		// 									InterfaceDescription: to.Ptr("Hyper-V Virtual Ethernet Adapter"),
		// 									ComponentID: to.Ptr("vms_mp"),
		// 									DriverVersion: to.Ptr("10.0.25398.1"),
		// 									Ip4Address: to.Ptr("192.168.200.92"),
		// 									SubnetMask: to.Ptr("255.255.255.0"),
		// 									DefaultGateway: to.Ptr("192.168.200.1"),
		// 									DNSServers: []*string{
		// 										to.Ptr("192.168.200.222"),
		// 									},
		// 									DefaultIsolationID: to.Ptr("0"),
		// 									Slot: to.Ptr("0"),
		// 									MacAddress: to.Ptr("000000000041"),
		// 									NicType: to.Ptr("Virtual"),
		// 									VlanID: to.Ptr("0"),
		// 									NicStatus: to.Ptr("Up"),
		// 								},
		// 							},
		// 							SwitchDetails: []*armazurestackhci.SwitchDetail{
		// 								{
		// 									SwitchName: to.Ptr("vmanagement"),
		// 									SwitchType: to.Ptr("External"),
		// 								},
		// 							},
		// 							HostNetwork: &armazurestackhci.EdgeMachineHostNetwork{
		// 								Intents: []*armazurestackhci.EdgeMachineHostNetworkIntent{
		// 									{
		// 										Scope: to.Ptr[int64](1),
		// 										IntentType: to.Ptr[int64](3),
		// 										IsComputeIntentSet: to.Ptr(true),
		// 										IsStorageIntentSet: to.Ptr(true),
		// 										IsOnlyStorage: to.Ptr(false),
		// 										IsManagementIntentSet: to.Ptr(true),
		// 										IsStretchIntentSet: to.Ptr(false),
		// 										IsOnlyStretch: to.Ptr(false),
		// 										IsNetworkIntentType: to.Ptr(true),
		// 										IntentName: to.Ptr("ManagementComputeStorage"),
		// 										IntentAdapters: []*string{
		// 											to.Ptr("port1"),
		// 											to.Ptr("port2"),
		// 										},
		// 										OverrideVirtualSwitchConfiguration: to.Ptr(false),
		// 										VirtualSwitchConfigurationOverrides: &armazurestackhci.EdgeMachineVirtualSwitchConfigurationOverrides{
		// 											EnableIov: to.Ptr("true"),
		// 											LoadBalancingAlgorithm: to.Ptr("Dynamic"),
		// 										},
		// 										OverrideQosPolicy: to.Ptr(false),
		// 										QosPolicyOverrides: &armazurestackhci.QosPolicyOverrides{
		// 											PriorityValue8021ActionCluster: to.Ptr("7"),
		// 											PriorityValue8021ActionSMB: to.Ptr("3"),
		// 											BandwidthPercentageSMB: to.Ptr("50"),
		// 										},
		// 										OverrideAdapterProperty: to.Ptr(false),
		// 										AdapterPropertyOverrides: &armazurestackhci.EdgeMachineAdapterPropertyOverrides{
		// 											JumboPacket: to.Ptr("9014"),
		// 											NetworkDirect: to.Ptr("Enabled"),
		// 											NetworkDirectTechnology: to.Ptr("RoCEv2"),
		// 										},
		// 									},
		// 								},
		// 								StorageNetworks: []*armazurestackhci.EdgeMachineStorageNetwork{
		// 									{
		// 										Name: to.Ptr("Storage1Network"),
		// 										NetworkAdapterName: to.Ptr("port3"),
		// 										StorageVlanID: to.Ptr("711"),
		// 										StorageAdapterIPInfo: []*armazurestackhci.EdgeMachineStorageAdapterIPInfo{
		// 											{
		// 												PhysicalNode: to.Ptr("Node-1"),
		// 												IPv4Address: to.Ptr("10.0.0.1"),
		// 												SubnetMask: to.Ptr("255.255.255.0"),
		// 											},
		// 										},
		// 									},
		// 								},
		// 								StorageConnectivitySwitchless: to.Ptr(true),
		// 								EnableStorageAutoIP: to.Ptr(true),
		// 							},
		// 							SdnProperties: &armazurestackhci.SdnProperties{
		// 								SdnStatus: to.Ptr(armazurestackhci.SdnStatusEnabled),
		// 								SdnDomainName: to.Ptr("sdn-nc.contoso.com"),
		// 								SdnAPIAddress: to.Ptr("10.0.0.10"),
		// 							},
		// 						},
		// 						StorageProfile: &armazurestackhci.StorageProfile{
		// 							PoolableDisksCount: to.Ptr[int64](4),
		// 						},
		// 						OSProfile: &armazurestackhci.OsProfile{
		// 							BootType: to.Ptr("UEFI"),
		// 							AssemblyVersion: to.Ptr("2402.1"),
		// 						},
		// 						HardwareProfile: &armazurestackhci.HardwareProfile{
		// 							CPUCores: to.Ptr[int64](4),
		// 							CPUSockets: to.Ptr[int64](14),
		// 							MemoryCapacityInGb: to.Ptr[int64](25),
		// 							SerialNumber: to.Ptr("5855-7104-5665-0328-3864-8600-21"),
		// 						},
		// 						SbeDeploymentPackageInfo: &armazurestackhci.SbeDeploymentPackageInfo{
		// 							Code: to.Ptr("NewerThanLatestPublished"),
		// 							Message: to.Ptr("The SBE package at path 'C:\\SBE' with version 4.1.2312.10 is published later than the latest SBE manifest published for online discovery. "),
		// 							SbeManifest: to.Ptr("PEFwcGxpY2Fi"),
		// 						},
		// 						ExtensionProfile: &armazurestackhci.ExtensionProfile{
		// 							Extensions: []*armazurestackhci.HciEdgeDeviceArcExtension{
		// 								{
		// 									ExtensionName: to.Ptr("DeviceManagementExtension"),
		// 									State: to.Ptr(armazurestackhci.ArcExtensionState("Connected")),
		// 									ExtensionResourceID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1/Extensions/DeviceManagementExtension"),
		// 									TypeHandlerVersion: to.Ptr("1.10.3"),
		// 									ManagedBy: to.Ptr(armazurestackhci.ExtensionManagedByAzure),
		// 								},
		// 							},
		// 						},
		// 					},
		// 				},
		// 				Location: to.Ptr("eastus"),
		// 				ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.AzureStackHCI/edgeMachines/machine-1"),
		// 				Name: to.Ptr("machine-1"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/edgeMachines"),
		// 				SystemData: &armazurestackhci.SystemData{
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2024, time.November, 14, 10, 46, 55, 167000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2024, time.November, 14, 10, 46, 55, 167000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
