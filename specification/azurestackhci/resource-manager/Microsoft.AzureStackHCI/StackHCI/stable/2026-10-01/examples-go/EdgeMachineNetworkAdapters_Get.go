package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-10-01/EdgeMachineNetworkAdapters_Get.json
func ExampleEdgeMachineNetworkAdaptersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("fd3c3665-1729-4b7b-9a38-238e83b0f98b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEdgeMachineNetworkAdaptersClient().Get(ctx, "test-rg", "EdgeMachine01", "ethernet0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armazurestackhci.EdgeMachineNetworkAdaptersClientGetResponse{
	// 	EdgeMachineNetworkAdapter: armazurestackhci.EdgeMachineNetworkAdapter{
	// 		Properties: &armazurestackhci.EdgeMachineNetworkAdapterProperties{
	// 			NetworkConfiguration: &armazurestackhci.NetworkAdapterConfiguration{
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
	// 			ReportedProperties: &armazurestackhci.NetworkAdapterReportedProperties{
	// 				AdapterName: to.Ptr("ethernet0"),
	// 				Ip4Address: to.Ptr("192.168.1.100"),
	// 				SubnetMask: to.Ptr("255.255.255.0"),
	// 				DefaultGateway: to.Ptr("192.168.1.1"),
	// 				DNSServers: []*string{
	// 					to.Ptr("8.8.8.8"),
	// 					to.Ptr("8.8.4.4"),
	// 				},
	// 				InterfaceDescription: to.Ptr("Intel(R) Ethernet Controller X710"),
	// 				ComponentID: to.Ptr("PCI\\VEN_8086&DEV_1572"),
	// 				MacAddress: to.Ptr("00:1A:2B:3C:4D:5E"),
	// 				InterfaceSpeed: to.Ptr[int32](10000),
	// 				IPInterfaceType: to.Ptr(armazurestackhci.IPInterfaceTypeStatic),
	// 				VlanID: to.Ptr[int32](100),
	// 				InterfaceState: to.Ptr(armazurestackhci.InterfaceStateUp),
	// 			},
	// 			ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/test-rg/providers/Microsoft.AzureStackHCI/edgeMachines/EdgeMachine01/networkAdapters/ethernet0"),
	// 		Name: to.Ptr("ethernet0"),
	// 		Type: to.Ptr("Microsoft.AzureStackHCI/edgeMachines/networkAdapters"),
	// 	},
	// }
}
