package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_WithInterconnectBlock.json
func ExampleVirtualMachineScaleSetVMsClient_Get_getVMScaleSetVMWithInterconnectBlock() {
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
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines"),
	// 		Location: to.Ptr("westus"),
	// 		InstanceID: to.Ptr("0"),
	// 		SKU: &armcompute.SKU{
	// 			Name: to.Ptr("Standard_ND128isr_GB300_v6"),
	// 			Tier: to.Ptr("Standard"),
	// 		},
	// 		Properties: &armcompute.VirtualMachineScaleSetVMProperties{
	// 			LatestModelApplied: to.Ptr(true),
	// 			ModelDefinitionApplied: to.Ptr("VirtualMachineScaleSet"),
	// 			NetworkProfileConfiguration: &armcompute.VirtualMachineScaleSetVMNetworkProfileConfiguration{
	// 				InterconnectGroupProfile: &armcompute.InterconnectGroupProfile{
	// 					InterconnectGroup: &armcompute.SubResource{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup"),
	// 					},
	// 					Subgroups: []*armcompute.SubResource{
	// 						{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup/subgroups/subgroup0"),
	// 						},
	// 					},
	// 				},
	// 				NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
	// 					{
	// 						Name: to.Ptr("{vmss-name}-nic"),
	// 						Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 							Primary: to.Ptr(true),
	// 							DisableTCPStateTracking: to.Ptr(false),
	// 							DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
	// 								DNSServers: []*string{
	// 								},
	// 							},
	// 							EnableIPForwarding: to.Ptr(false),
	// 							IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
	// 								{
	// 									Name: to.Ptr("ipconfig1"),
	// 									Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
	// 										Subnet: &armcompute.APIEntityReference{
	// 											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}"),
	// 										},
	// 										PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
	// 									},
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			HardwareProfile: &armcompute.HardwareProfile{
	// 				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypes("Standard_ND128isr_GB300_v6")),
	// 			},
	// 			ResilientVMDeletionStatus: to.Ptr(armcompute.ResilientVMDeletionStatusDisabled),
	// 			VMID: to.Ptr("bd73f1c6-1b3b-43b0-8b96-718bace3ba2e"),
	// 			InterconnectBlockProfile: &armcompute.InterconnectBlockProfile{
	// 				InterconnectBlock: &armcompute.APIEntityReference{
	// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/interconnectBlocks/myInterconnectBlock"),
	// 				},
	// 			},
	// 			StorageProfile: &armcompute.StorageProfile{
	// 				ImageReference: &armcompute.ImageReference{
	// 					Publisher: to.Ptr("microsoft-dsvm"),
	// 					Offer: to.Ptr("ubuntu-hpc"),
	// 					SKU: to.Ptr("2404-gb"),
	// 					Version: to.Ptr("latest"),
	// 					ExactVersion: to.Ptr("24.04.2025110401"),
	// 				},
	// 				OSDisk: &armcompute.OSDisk{
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
	// 					Name: to.Ptr("{vmss-name}_{vmss-vm-name}_OsDisk_1_e195e35c5b854d3ab60960e773dd2997"),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/{vmss-name}_{vmss-vm-name}_OsDisk_1_e195e35c5b854d3ab60960e773dd2997"),
	// 					},
	// 					DiskSizeGB: to.Ptr[int32](64),
	// 				},
	// 				DataDisks: []*armcompute.DataDisk{
	// 				},
	// 				DiskControllerType: to.Ptr(armcompute.DiskControllerTypesNVMe),
	// 			},
	// 			OSProfile: &armcompute.OSProfile{
	// 				ComputerName: to.Ptr("{vmss-name}000000"),
	// 				AdminUsername: to.Ptr("{your-username}"),
	// 				LinuxConfiguration: &armcompute.LinuxConfiguration{
	// 					DisablePasswordAuthentication: to.Ptr(false),
	// 					SSH: &armcompute.SSHConfiguration{
	// 						PublicKeys: []*armcompute.SSHPublicKey{
	// 							{
	// 								Path: to.Ptr("/home/{your-username}/.ssh/authorized_keys"),
	// 								KeyData: to.Ptr("{ssh-rsa public key}"),
	// 							},
	// 						},
	// 					},
	// 					ProvisionVMAgent: to.Ptr(true),
	// 					EnableVMAgentPlatformUpdates: to.Ptr(true),
	// 				},
	// 				Secrets: []*armcompute.VaultSecretGroup{
	// 				},
	// 				AllowExtensionOperations: to.Ptr(true),
	// 				RequireGuestProvisionSignal: to.Ptr(true),
	// 			},
	// 			SecurityProfile: &armcompute.SecurityProfile{
	// 				SecurityType: to.Ptr(armcompute.SecurityTypesStandard),
	// 			},
	// 			NetworkProfile: &armcompute.NetworkProfile{
	// 				InterconnectGroupProfile: &armcompute.InterconnectGroupProfile{
	// 					InterconnectGroup: &armcompute.SubResource{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup"),
	// 					},
	// 					Subgroups: []*armcompute.SubResource{
	// 						{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/interconnectGroups/myInterconnectGroup/subgroups/subgroup0"),
	// 						},
	// 					},
	// 				},
	// 				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/0/networkInterfaces/{vmss-name}-nic"),
	// 					},
	// 				},
	// 			},
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T19:59:05.5350459+00:00"); return t}()),
	// 		},
	// 		Etag: to.Ptr("\"1\""),
	// 		Zones: []*string{
	// 			to.Ptr("1"),
	// 		},
	// 	},
	// }
}
