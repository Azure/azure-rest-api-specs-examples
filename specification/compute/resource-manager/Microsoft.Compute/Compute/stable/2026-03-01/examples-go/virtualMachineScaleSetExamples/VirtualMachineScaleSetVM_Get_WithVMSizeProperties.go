package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_WithVMSizeProperties.json
func ExampleVirtualMachineScaleSetVMsClient_Get_getVMScaleSetVMWithVMSizeProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetVMsClient().Get(ctx, "myResourceGroup", "{vmss-name}", "0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachineScaleSetVMsClientGetResponse{
	// 	VirtualMachineScaleSetVM: armcompute.VirtualMachineScaleSetVM{
	// 		Name: to.Ptr("{vmss-vm-name}"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/0"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachines"),
	// 		Location: to.Ptr("westus"),
	// 		Tags: map[string]*string{
	// 			"myTag1": to.Ptr("tagValue1"),
	// 		},
	// 		Etag: to.Ptr("\"1\""),
	// 		Properties: &armcompute.VirtualMachineScaleSetVMProperties{
	// 			LatestModelApplied: to.Ptr(true),
	// 			ModelDefinitionApplied: to.Ptr("VirtualMachineScaleSet"),
	// 			NetworkProfileConfiguration: &armcompute.VirtualMachineScaleSetVMNetworkProfileConfiguration{
	// 				NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
	// 					{
	// 						Name: to.Ptr("vmsstestnetconfig5415"),
	// 						Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 							Primary: to.Ptr(true),
	// 							EnableAcceleratedNetworking: to.Ptr(false),
	// 							DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
	// 								DNSServers: []*string{
	// 								},
	// 							},
	// 							EnableIPForwarding: to.Ptr(false),
	// 							IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
	// 								{
	// 									Name: to.Ptr("vmsstestnetconfig9693"),
	// 									Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
	// 										Subnet: &armcompute.APIEntityReference{
	// 											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/vn4071/subnets/sn5503"),
	// 										},
	// 										PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			VMID: to.Ptr("42af9fdf-b906-4ad7-9905-8316209ff619"),
	// 			HardwareProfile: &armcompute.HardwareProfile{
	// 				VMSizeProperties: &armcompute.VMSizeProperties{
	// 					VCPUsAvailable: to.Ptr[int32](1),
	// 					VCPUsPerCore: to.Ptr[int32](1),
	// 				},
	// 			},
	// 			StorageProfile: &armcompute.StorageProfile{
	// 				ImageReference: &armcompute.ImageReference{
	// 					Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 					Offer: to.Ptr("WindowsServer"),
	// 					SKU: to.Ptr("2012-R2-Datacenter"),
	// 					Version: to.Ptr("4.127.20180315"),
	// 					ExactVersion: to.Ptr("4.127.20180315"),
	// 				},
	// 				OSDisk: &armcompute.OSDisk{
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 					Name: to.Ptr("vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc"),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					Caching: to.Ptr(armcompute.CachingTypesNone),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc"),
	// 					},
	// 					DiskSizeGB: to.Ptr[int32](127),
	// 				},
	// 				DataDisks: []*armcompute.DataDisk{
	// 					{
	// 						Lun: to.Ptr[int32](1),
	// 						Name: to.Ptr("vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
	// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
	// 						Caching: to.Ptr(armcompute.CachingTypesNone),
	// 						ManagedDisk: &armcompute.ManagedDiskParameters{
	// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
	// 						},
	// 						DiskSizeGB: to.Ptr[int32](128),
	// 						ToBeDetached: to.Ptr(false),
	// 					},
	// 				},
	// 			},
	// 			OSProfile: &armcompute.OSProfile{
	// 				ComputerName: to.Ptr("test000000"),
	// 				AdminUsername: to.Ptr("Foo12"),
	// 				WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 					ProvisionVMAgent: to.Ptr(true),
	// 					EnableAutomaticUpdates: to.Ptr(true),
	// 				},
	// 				Secrets: []*armcompute.VaultSecretGroup{
	// 				},
	// 				AllowExtensionOperations: to.Ptr(true),
	// 				RequireGuestProvisionSignal: to.Ptr(true),
	// 			},
	// 			UserData: to.Ptr("RXhhbXBsZSBVc2VyRGF0YQ=="),
	// 			NetworkProfile: &armcompute.NetworkProfile{
	// 				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/0/networkInterfaces/vmsstestnetconfig5415"),
	// 					},
	// 				},
	// 			},
	// 			SecurityProfile: &armcompute.SecurityProfile{
	// 				SecurityType: to.Ptr(armcompute.SecurityTypesStandard),
	// 			},
	// 			DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 				BootDiagnostics: &armcompute.BootDiagnostics{
	// 					Enabled: to.Ptr(true),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 		Resources: []*armcompute.VirtualMachineExtension{
	// 			{
	// 				Name: to.Ptr("CustomScriptExtension-DSC"),
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/extensions/CustomScriptExtension-DSC"),
	// 				Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
	// 				Location: to.Ptr("westus"),
	// 				Tags: map[string]*string{
	// 					"displayName": to.Ptr("CustomScriptExtension-DSC"),
	// 				},
	// 				Properties: &armcompute.VirtualMachineExtensionProperties{
	// 					AutoUpgradeMinorVersion: to.Ptr(true),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					Publisher: to.Ptr("Microsoft.Compute"),
	// 					Type: to.Ptr("CustomScriptExtension"),
	// 					TypeHandlerVersion: to.Ptr("1.9"),
	// 					Settings: map[string]any{
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
