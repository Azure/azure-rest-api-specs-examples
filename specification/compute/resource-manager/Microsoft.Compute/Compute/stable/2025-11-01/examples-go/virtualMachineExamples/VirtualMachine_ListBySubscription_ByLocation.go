package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/virtualMachineExamples/VirtualMachine_ListBySubscription_ByLocation.json
func ExampleVirtualMachinesClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscriptionId}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachinesClient().NewListByLocationPager("eastus", nil)
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
		// page = armcompute.VirtualMachinesClientListByLocationResponse{
		// 	VirtualMachineListResult: armcompute.VirtualMachineListResult{
		// 		Value: []*armcompute.VirtualMachine{
		// 			{
		// 				Properties: &armcompute.VirtualMachineProperties{
		// 					VMID: to.Ptr("{vmId}"),
		// 					AvailabilitySet: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					HardwareProfile: &armcompute.HardwareProfile{
		// 						VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardA0),
		// 					},
		// 					StorageProfile: &armcompute.StorageProfile{
		// 						ImageReference: &armcompute.ImageReference{
		// 							Publisher: to.Ptr("MicrosoftWindowsServer"),
		// 							Offer: to.Ptr("WindowsServer"),
		// 							SKU: to.Ptr("2012-R2-Datacenter"),
		// 							Version: to.Ptr("4.127.20170406"),
		// 						},
		// 						OSDisk: &armcompute.OSDisk{
		// 							OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 							Name: to.Ptr("test"),
		// 							CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 							Vhd: &armcompute.VirtualHardDisk{
		// 								URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 							},
		// 							Caching: to.Ptr(armcompute.CachingTypesNone),
		// 							DiskSizeGB: to.Ptr[int32](127),
		// 						},
		// 						DataDisks: []*armcompute.DataDisk{
		// 						},
		// 					},
		// 					OSProfile: &armcompute.OSProfile{
		// 						ComputerName: to.Ptr("Test"),
		// 						AdminUsername: to.Ptr("Foo12"),
		// 						WindowsConfiguration: &armcompute.WindowsConfiguration{
		// 							ProvisionVMAgent: to.Ptr(true),
		// 							EnableAutomaticUpdates: to.Ptr(true),
		// 						},
		// 						Secrets: []*armcompute.VaultSecretGroup{
		// 						},
		// 						AllowExtensionOperations: to.Ptr(true),
		// 					},
		// 					NetworkProfile: &armcompute.NetworkProfile{
		// 						NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
		// 							{
		// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}"),
		// 							},
		// 						},
		// 					},
		// 					SecurityProfile: &armcompute.SecurityProfile{
		// 						SecurityType: to.Ptr(armcompute.SecurityTypesStandard),
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 					"RG": to.Ptr("rg"),
		// 					"testTag": to.Ptr("1"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}"),
		// 				Name: to.Ptr("{virtualMachineName}"),
		// 			},
		// 			{
		// 				Properties: &armcompute.VirtualMachineProperties{
		// 					VMID: to.Ptr("{vmId}"),
		// 					AvailabilitySet: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					HardwareProfile: &armcompute.HardwareProfile{
		// 						VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardA0),
		// 					},
		// 					StorageProfile: &armcompute.StorageProfile{
		// 						ImageReference: &armcompute.ImageReference{
		// 							Publisher: to.Ptr("MicrosoftWindowsServer"),
		// 							Offer: to.Ptr("WindowsServer"),
		// 							SKU: to.Ptr("2012-R2-Datacenter"),
		// 							Version: to.Ptr("4.127.20170406"),
		// 						},
		// 						OSDisk: &armcompute.OSDisk{
		// 							OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 							Name: to.Ptr("test"),
		// 							CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 							Vhd: &armcompute.VirtualHardDisk{
		// 								URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 							},
		// 							Caching: to.Ptr(armcompute.CachingTypesNone),
		// 							DiskSizeGB: to.Ptr[int32](127),
		// 						},
		// 						DataDisks: []*armcompute.DataDisk{
		// 						},
		// 					},
		// 					OSProfile: &armcompute.OSProfile{
		// 						ComputerName: to.Ptr("Test"),
		// 						AdminUsername: to.Ptr("Foo12"),
		// 						WindowsConfiguration: &armcompute.WindowsConfiguration{
		// 							ProvisionVMAgent: to.Ptr(true),
		// 							EnableAutomaticUpdates: to.Ptr(true),
		// 						},
		// 						Secrets: []*armcompute.VaultSecretGroup{
		// 						},
		// 						AllowExtensionOperations: to.Ptr(true),
		// 					},
		// 					NetworkProfile: &armcompute.NetworkProfile{
		// 						NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
		// 							{
		// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}"),
		// 							},
		// 						},
		// 					},
		// 					SecurityProfile: &armcompute.SecurityProfile{
		// 						SecurityType: to.Ptr(armcompute.SecurityTypesStandard),
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 				Location: to.Ptr("eastus"),
		// 				Tags: map[string]*string{
		// 					"RG": to.Ptr("rg"),
		// 					"testTag": to.Ptr("1"),
		// 				},
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}"),
		// 				Name: to.Ptr("{virtualMachineName}"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
