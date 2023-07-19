package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/BareMetalMachines_Create.json
func ExampleBareMetalMachinesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBareMetalMachinesClient().BeginCreateOrUpdate(ctx, "resourceGroupName", "bareMetalMachineName", armnetworkcloud.BareMetalMachine{
		Location: to.Ptr("location"),
		Tags: map[string]*string{
			"key1": to.Ptr("myvalue1"),
			"key2": to.Ptr("myvalue2"),
		},
		ExtendedLocation: &armnetworkcloud.ExtendedLocation{
			Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
			Type: to.Ptr("CustomLocation"),
		},
		Properties: &armnetworkcloud.BareMetalMachineProperties{
			BmcConnectionString: to.Ptr("bmcconnectionstring"),
			BmcCredentials: &armnetworkcloud.AdministrativeCredentials{
				Password: to.Ptr("{password}"),
				Username: to.Ptr("bmcuser"),
			},
			BmcMacAddress:  to.Ptr("00:00:4f:00:57:00"),
			BootMacAddress: to.Ptr("00:00:4e:00:58:af"),
			MachineDetails: to.Ptr("User-provided machine details."),
			MachineName:    to.Ptr("r01c001"),
			MachineSKUID:   to.Ptr("684E-3B16-399E"),
			RackID:         to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName"),
			RackSlot:       to.Ptr[int64](1),
			SerialNumber:   to.Ptr("BM1219XXX"),
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
	// res.BareMetalMachine = armnetworkcloud.BareMetalMachine{
	// 	Name: to.Ptr("bareMetalMachineName"),
	// 	Type: to.Ptr("Microsoft.NetworkCloud/bareMetalMachines"),
	// 	ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
	// 	SystemData: &armnetworkcloud.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
	// 		CreatedBy: to.Ptr("identityA"),
	// 		CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("identityB"),
	// 		LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("myvalue1"),
	// 		"key2": to.Ptr("myvalue2"),
	// 	},
	// 	ExtendedLocation: &armnetworkcloud.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
	// 		Type: to.Ptr("CustomLocation"),
	// 	},
	// 	Properties: &armnetworkcloud.BareMetalMachineProperties{
	// 		AssociatedResourceIDs: []*string{
	// 			to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/virtualMachines/virtualMachineName")},
	// 			BmcConnectionString: to.Ptr("bmcconnectionstring"),
	// 			BmcCredentials: &armnetworkcloud.AdministrativeCredentials{
	// 				Username: to.Ptr("bmcuser"),
	// 			},
	// 			BmcMacAddress: to.Ptr("00:00:4f:00:57:00"),
	// 			BootMacAddress: to.Ptr("00:00:4e:00:58:af"),
	// 			ClusterID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName"),
	// 			CordonStatus: to.Ptr(armnetworkcloud.BareMetalMachineCordonStatusUncordoned),
	// 			DetailedStatus: to.Ptr(armnetworkcloud.BareMetalMachineDetailedStatusAvailable),
	// 			DetailedStatusMessage: to.Ptr("DetailedStatusMessage"),
	// 			HardwareInventory: &armnetworkcloud.HardwareInventory{
	// 				AdditionalHostInformation: to.Ptr("Machine specific information..."),
	// 				Interfaces: []*armnetworkcloud.HardwareInventoryNetworkInterface{
	// 					{
	// 						Name: to.Ptr("networkInterfaceName"),
	// 						LinkStatus: to.Ptr("Up"),
	// 						MacAddress: to.Ptr("2C:54:91:88:C9:E3"),
	// 						NetworkInterfaceID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/networkDevices/networkDeviceName/networkInterfaces/networkInterfaceName"),
	// 				}},
	// 			},
	// 			HardwareValidationStatus: &armnetworkcloud.HardwareValidationStatus{
	// 				LastValidationTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-09-30T13:27:03.008Z"); return t}()),
	// 				Result: to.Ptr(armnetworkcloud.BareMetalMachineHardwareValidationResultPass),
	// 			},
	// 			KubernetesNodeName: to.Ptr("node01"),
	// 			KubernetesVersion: to.Ptr("1.21"),
	// 			MachineDetails: to.Ptr("User-provided machine details."),
	// 			MachineName: to.Ptr("r01c001"),
	// 			MachineSKUID: to.Ptr("684E-3B16-399E"),
	// 			OamIPv4Address: to.Ptr("192.0.2.1"),
	// 			OamIPv6Address: to.Ptr("0:0:0:0:0:FFFF:7F00:0001"),
	// 			OSImage: to.Ptr("v20220805e"),
	// 			PowerState: to.Ptr(armnetworkcloud.BareMetalMachinePowerStateOn),
	// 			ProvisioningState: to.Ptr(armnetworkcloud.BareMetalMachineProvisioningStateSucceeded),
	// 			RackID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/racks/rackName"),
	// 			RackSlot: to.Ptr[int64](1),
	// 			ReadyState: to.Ptr(armnetworkcloud.BareMetalMachineReadyStateTrue),
	// 			SerialNumber: to.Ptr("BM1219XXX"),
	// 			ServiceTag: to.Ptr("ST1219XXX"),
	// 		},
	// 	}
}
