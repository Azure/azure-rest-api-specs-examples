package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_WithUserData.json
func ExampleVirtualMachineScaleSetVMsClient_Get_getVmScaleSetVmWithUserData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetVMsClient().Get(ctx, "myResourceGroup", "{vmss-name}", "0", &armcompute.VirtualMachineScaleSetVMsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineScaleSetVM = armcompute.VirtualMachineScaleSetVM{
	// 	Name: to.Ptr("{vmss-vm-name}"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/0"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"myTag1": to.Ptr("tagValue1"),
	// 	},
	// 	Etag: to.Ptr("\"1\""),
	// 	Properties: &armcompute.VirtualMachineScaleSetVMProperties{
	// 		DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 			BootDiagnostics: &armcompute.BootDiagnostics{
	// 				Enabled: to.Ptr(true),
	// 			},
	// 		},
	// 		HardwareProfile: &armcompute.HardwareProfile{
	// 		},
	// 		LatestModelApplied: to.Ptr(true),
	// 		ModelDefinitionApplied: to.Ptr("VirtualMachineScaleSet"),
	// 		NetworkProfile: &armcompute.NetworkProfile{
	// 			NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 				{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/0/networkInterfaces/vmsstestnetconfig5415"),
	// 			}},
	// 		},
	// 		NetworkProfileConfiguration: &armcompute.VirtualMachineScaleSetVMNetworkProfileConfiguration{
	// 			NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
	// 				{
	// 					Name: to.Ptr("vmsstestnetconfig5415"),
	// 					Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 						DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
	// 							DNSServers: []*string{
	// 							},
	// 						},
	// 						EnableAcceleratedNetworking: to.Ptr(false),
	// 						EnableIPForwarding: to.Ptr(false),
	// 						IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
	// 							{
	// 								Name: to.Ptr("vmsstestnetconfig9693"),
	// 								Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
	// 									PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
	// 									Subnet: &armcompute.APIEntityReference{
	// 										ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/vn4071/subnets/sn5503"),
	// 									},
	// 								},
	// 						}},
	// 						Primary: to.Ptr(true),
	// 					},
	// 			}},
	// 		},
	// 		OSProfile: &armcompute.OSProfile{
	// 			AdminUsername: to.Ptr("Foo12"),
	// 			AllowExtensionOperations: to.Ptr(true),
	// 			ComputerName: to.Ptr("test000000"),
	// 			RequireGuestProvisionSignal: to.Ptr(true),
	// 			Secrets: []*armcompute.VaultSecretGroup{
	// 			},
	// 			WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 				EnableAutomaticUpdates: to.Ptr(true),
	// 				ProvisionVMAgent: to.Ptr(true),
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		StorageProfile: &armcompute.StorageProfile{
	// 			DataDisks: []*armcompute.DataDisk{
	// 				{
	// 					Name: to.Ptr("vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
	// 					Caching: to.Ptr(armcompute.CachingTypesNone),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
	// 					DiskSizeGB: to.Ptr[int32](128),
	// 					Lun: to.Ptr[int32](1),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 					},
	// 					ToBeDetached: to.Ptr(false),
	// 			}},
	// 			ImageReference: &armcompute.ImageReference{
	// 				ExactVersion: to.Ptr("4.127.20180315"),
	// 				Offer: to.Ptr("WindowsServer"),
	// 				Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 				SKU: to.Ptr("2012-R2-Datacenter"),
	// 				Version: to.Ptr("4.127.20180315"),
	// 			},
	// 			OSDisk: &armcompute.OSDisk{
	// 				Name: to.Ptr("vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc"),
	// 				Caching: to.Ptr(armcompute.CachingTypesNone),
	// 				CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 				DiskSizeGB: to.Ptr[int32](127),
	// 				ManagedDisk: &armcompute.ManagedDiskParameters{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc"),
	// 					StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 				},
	// 				OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 			},
	// 		},
	// 		UserData: to.Ptr("RXhhbXBsZSBVc2VyRGF0YQ=="),
	// 		VMID: to.Ptr("42af9fdf-b906-4ad7-9905-8316209ff619"),
	// 	},
	// 	Resources: []*armcompute.VirtualMachineExtension{
	// 		{
	// 			Name: to.Ptr("CustomScriptExtension-DSC"),
	// 			Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
	// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/extensions/CustomScriptExtension-DSC"),
	// 			Location: to.Ptr("westus"),
	// 			Tags: map[string]*string{
	// 				"displayName": to.Ptr("CustomScriptExtension-DSC"),
	// 			},
	// 			Properties: &armcompute.VirtualMachineExtensionProperties{
	// 				Type: to.Ptr("CustomScriptExtension"),
	// 				AutoUpgradeMinorVersion: to.Ptr(true),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				Publisher: to.Ptr("Microsoft.Compute"),
	// 				Settings: map[string]any{
	// 				},
	// 				TypeHandlerVersion: to.Ptr("1.9"),
	// 			},
	// 	}},
	// }
}
