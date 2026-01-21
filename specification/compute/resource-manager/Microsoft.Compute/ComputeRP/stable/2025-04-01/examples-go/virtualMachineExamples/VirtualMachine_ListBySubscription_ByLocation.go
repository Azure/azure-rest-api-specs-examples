package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineExamples/VirtualMachine_ListBySubscription_ByLocation.json
func ExampleVirtualMachinesClient_NewListByLocationPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.VirtualMachineListResult = armcompute.VirtualMachineListResult{
		// 	Value: []*armcompute.VirtualMachine{
		// 		{
		// 			Name: to.Ptr("{virtualMachineName}"),
		// 			Type: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"RG": to.Ptr("rg"),
		// 				"testTag": to.Ptr("1"),
		// 			},
		// 			Properties: &armcompute.VirtualMachineProperties{
		// 				AvailabilitySet: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				HardwareProfile: &armcompute.HardwareProfile{
		// 					VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardA0),
		// 				},
		// 				NetworkProfile: &armcompute.NetworkProfile{
		// 					NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}"),
		// 					}},
		// 				},
		// 				OSProfile: &armcompute.OSProfile{
		// 					AdminUsername: to.Ptr("Foo12"),
		// 					AllowExtensionOperations: to.Ptr(true),
		// 					ComputerName: to.Ptr("Test"),
		// 					Secrets: []*armcompute.VaultSecretGroup{
		// 					},
		// 					WindowsConfiguration: &armcompute.WindowsConfiguration{
		// 						EnableAutomaticUpdates: to.Ptr(true),
		// 						ProvisionVMAgent: to.Ptr(true),
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				StorageProfile: &armcompute.StorageProfile{
		// 					DataDisks: []*armcompute.DataDisk{
		// 					},
		// 					ImageReference: &armcompute.ImageReference{
		// 						Offer: to.Ptr("WindowsServer"),
		// 						Publisher: to.Ptr("MicrosoftWindowsServer"),
		// 						SKU: to.Ptr("2012-R2-Datacenter"),
		// 						Version: to.Ptr("4.127.20170406"),
		// 					},
		// 					OSDisk: &armcompute.OSDisk{
		// 						Name: to.Ptr("test"),
		// 						Caching: to.Ptr(armcompute.CachingTypesNone),
		// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 						DiskSizeGB: to.Ptr[int32](127),
		// 						OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 						Vhd: &armcompute.VirtualHardDisk{
		// 							URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 						},
		// 					},
		// 				},
		// 				VMID: to.Ptr("{vmId}"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{virtualMachineName}"),
		// 			Type: to.Ptr("Microsoft.Compute/virtualMachines"),
		// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{virtualMachineName}"),
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 				"RG": to.Ptr("rg"),
		// 				"testTag": to.Ptr("1"),
		// 			},
		// 			Properties: &armcompute.VirtualMachineProperties{
		// 				AvailabilitySet: &armcompute.SubResource{
		// 					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 				},
		// 				HardwareProfile: &armcompute.HardwareProfile{
		// 					VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesStandardA0),
		// 				},
		// 				NetworkProfile: &armcompute.NetworkProfile{
		// 					NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}"),
		// 					}},
		// 				},
		// 				OSProfile: &armcompute.OSProfile{
		// 					AdminUsername: to.Ptr("Foo12"),
		// 					AllowExtensionOperations: to.Ptr(true),
		// 					ComputerName: to.Ptr("Test"),
		// 					Secrets: []*armcompute.VaultSecretGroup{
		// 					},
		// 					WindowsConfiguration: &armcompute.WindowsConfiguration{
		// 						EnableAutomaticUpdates: to.Ptr(true),
		// 						ProvisionVMAgent: to.Ptr(true),
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				StorageProfile: &armcompute.StorageProfile{
		// 					DataDisks: []*armcompute.DataDisk{
		// 					},
		// 					ImageReference: &armcompute.ImageReference{
		// 						Offer: to.Ptr("WindowsServer"),
		// 						Publisher: to.Ptr("MicrosoftWindowsServer"),
		// 						SKU: to.Ptr("2012-R2-Datacenter"),
		// 						Version: to.Ptr("4.127.20170406"),
		// 					},
		// 					OSDisk: &armcompute.OSDisk{
		// 						Name: to.Ptr("test"),
		// 						Caching: to.Ptr(armcompute.CachingTypesNone),
		// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 						DiskSizeGB: to.Ptr[int32](127),
		// 						OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 						Vhd: &armcompute.VirtualHardDisk{
		// 							URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 						},
		// 					},
		// 				},
		// 				VMID: to.Ptr("{vmId}"),
		// 			},
		// 	}},
		// }
	}
}
