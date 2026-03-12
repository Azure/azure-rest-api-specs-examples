package armazurestackhci_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci/v3"
)

// Generated from example definition: 2026-03-01-preview/ListEdgeDevices.json
func ExampleEdgeDevicesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazurestackhci.NewClientFactory("<subscriptionID>", cred, nil)
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
		// page = armazurestackhci.EdgeDevicesClientListResponse{
		// 	EdgeDeviceListResult: armazurestackhci.EdgeDeviceListResult{
		// 		Value: []armazurestackhci.EdgeDeviceClassification{
		// 			&armazurestackhci.HciEdgeDevice{
		// 				ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/ArcInstance-rg/providers/Microsoft.HybridCompute/machines/Node-1/providers/Microsoft.AzureStackHCI/edgeDevices/default"),
		// 				Name: to.Ptr("default"),
		// 				Type: to.Ptr("Microsoft.AzureStackHCI/edgeDevices"),
		// 				Kind: to.Ptr(armazurestackhci.DeviceKindHCI),
		// 				SystemData: &armazurestackhci.SystemData{
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armazurestackhci.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-02T17:18:19.1234567Z"); return t}()),
		// 				},
		// 				Properties: &armazurestackhci.HciEdgeDeviceProperties{
		// 					ReportedProperties: &armazurestackhci.HciReportedProperties{
		// 						NetworkProfile: &armazurestackhci.HciNetworkProfile{
		// 							NicDetails: []*armazurestackhci.HciNicDetail{
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
		// 							SdnProperties: &armazurestackhci.SdnProperties{
		// 								SdnStatus: to.Ptr(armazurestackhci.SdnStatusEnabled),
		// 								SdnDomainName: to.Ptr("cl-nc.fqdn"),
		// 								SdnAPIAddress: to.Ptr("192.0.2.10"),
		// 							},
		// 						},
		// 						OSProfile: &armazurestackhci.HciOsProfile{
		// 							BootType: to.Ptr("UEFI"),
		// 							AssemblyVersion: to.Ptr("2402.1"),
		// 						},
		// 						SbeDeploymentPackageInfo: &armazurestackhci.SbeDeploymentPackageInfo{
		// 							Code: to.Ptr("NewerThanLatestPublished"),
		// 							Message: to.Ptr("The SBE package at path 'C:\\SBE' with version 4.1.2312.10 is published later than the latest SBE manifest published for online discovery. "),
		// 							SbeManifest: to.Ptr("PEFwcGxpY2Fi"),
		// 						},
		// 						LastSyncTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-01T17:18:19.1234567Z"); return t}()),
		// 						ConfidentialVMProfile: &armazurestackhci.ConfidentialVMProfile{
		// 							IgvmStatus: to.Ptr(armazurestackhci.IgvmStatusEnabled),
		// 							StatusDetails: []*armazurestackhci.IgvmStatusDetail{
		// 								{
		// 									Code: to.Ptr("IgvmAgentDeployed"),
		// 									Message: to.Ptr("IgvmAgent is deployed"),
		// 								},
		// 								{
		// 									Code: to.Ptr("ConfidentialVmHardwareCapabilityFound"),
		// 									Message: to.Ptr("Confidential VM supported hardware detected"),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr(armazurestackhci.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
