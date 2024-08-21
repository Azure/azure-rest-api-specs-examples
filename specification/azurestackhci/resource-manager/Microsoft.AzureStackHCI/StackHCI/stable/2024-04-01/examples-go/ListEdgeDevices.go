package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c9b146c38df5f76e2d34a3ef771979805ff4ff73/specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/ListEdgeDevices.json
func ExampleEdgeDevicesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEdgeDevicesClient().NewListPager("subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1", nil)
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
		// page.EdgeDeviceListResult = armazurestackhci.EdgeDeviceListResult{
		// 	Value: []armazurestackhci.EdgeDeviceClassification{
		// 		&armazurestackhci.HciEdgeDevice{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.AzureStackHCI/edgeDevices"),
		// 			ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1/providers/Microsoft.AzureStackHCI/edgeDevices/default"),
		// 			SystemData: &armazurestackhci.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 			},
		// 			Kind: to.Ptr(armazurestackhci.DeviceKindHCI),
		// 			Properties: &armazurestackhci.HciEdgeDeviceProperties{
		// 				ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
		// 				ReportedProperties: &armazurestackhci.HciReportedProperties{
		// 					NetworkProfile: &armazurestackhci.HciNetworkProfile{
		// 						NicDetails: []*armazurestackhci.HciNicDetail{
		// 							{
		// 								AdapterName: to.Ptr("vmanagement"),
		// 								ComponentID: to.Ptr("vms_mp"),
		// 								DefaultGateway: to.Ptr("192.168.200.1"),
		// 								DefaultIsolationID: to.Ptr("0"),
		// 								DNSServers: []*string{
		// 									to.Ptr("192.168.200.222")},
		// 									DriverVersion: to.Ptr("10.0.25398.1"),
		// 									InterfaceDescription: to.Ptr("Hyper-V Virtual Ethernet Adapter"),
		// 									Ip4Address: to.Ptr("192.168.200.92"),
		// 									MacAddress: to.Ptr("000000000041"),
		// 									NicStatus: to.Ptr("Up"),
		// 									NicType: to.Ptr("Virtual"),
		// 									Slot: to.Ptr("0"),
		// 									SubnetMask: to.Ptr("255.255.255.0"),
		// 									VlanID: to.Ptr("0"),
		// 							}},
		// 							SwitchDetails: []*armazurestackhci.SwitchDetail{
		// 								{
		// 									SwitchName: to.Ptr("vmanagement"),
		// 									SwitchType: to.Ptr("External"),
		// 							}},
		// 						},
		// 						OSProfile: &armazurestackhci.HciOsProfile{
		// 							AssemblyVersion: to.Ptr("2402.1"),
		// 							BootType: to.Ptr("UEFI"),
		// 						},
		// 						SbeDeploymentPackageInfo: &armazurestackhci.SbeDeploymentPackageInfo{
		// 							Code: to.Ptr("NewerThanLatestPublished"),
		// 							Message: to.Ptr("The SBE package at path 'C:\\SBE' with version 4.1.2312.10 is published later than the latest SBE manifest published for online discovery. "),
		// 							SbeManifest: to.Ptr("PEFwcGxpY2Fi"),
		// 						},
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}
