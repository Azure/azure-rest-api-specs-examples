package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_Get_WithResilientVMDeletionStatus.json
func ExampleVirtualMachineScaleSetVMsClient_Get_getVMScaleSetVMWithResiliencyView() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualMachineScaleSetVMsClient().Get(ctx, "myResourceGroup", "{vmss-name}", "1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachineScaleSetVMsClientGetResponse{
	// 	VirtualMachineScaleSetVM: &armcompute.VirtualMachineScaleSetVM{
	// 		Name: to.Ptr("{vmss-vm-name}"),
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/1"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines"),
	// 		Location: to.Ptr("eastus2euap"),
	// 		Identity: &armcompute.VirtualMachineIdentity{
	// 			Type: to.Ptr(armcompute.ResourceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
	// 				"/subscriptions/{subscription-id}/resourceGroups/AzSecPackAutoConfigRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/AzSecPackAutoConfigUA-eastus2euap": &armcompute.UserAssignedIdentitiesValue{
	// 					PrincipalID: to.Ptr("f31e5089-a1e5-44a6-9048-a767ce07d26c"),
	// 					ClientID: to.Ptr("215414c9-8a82-4439-95ea-d09e3543a6e2"),
	// 				},
	// 			},
	// 		},
	// 		InstanceID: to.Ptr("1"),
	// 		SKU: &armcompute.SKU{
	// 			Name: to.Ptr("Standard_D2ls_v5"),
	// 			Tier: to.Ptr("Standard"),
	// 		},
	// 		Properties: &armcompute.VirtualMachineScaleSetVMProperties{
	// 			LatestModelApplied: to.Ptr(true),
	// 			ModelDefinitionApplied: to.Ptr("VirtualMachineScaleSet"),
	// 			NetworkProfileConfiguration: &armcompute.VirtualMachineScaleSetVMNetworkProfileConfiguration{
	// 				NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
	// 					{
	// 						Name: to.Ptr("vnet-eastus2euap-2-nic01"),
	// 						Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 							Primary: to.Ptr(true),
	// 							EnableAcceleratedNetworking: to.Ptr(true),
	// 							DisableTCPStateTracking: to.Ptr(false),
	// 							NetworkSecurityGroup: &armcompute.SubResource{
	// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkSecurityGroups/basicNsgvnet-eastus2euap-2-nic01"),
	// 							},
	// 							DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
	// 								DNSServers: []*string{
	// 								},
	// 							},
	// 							EnableIPForwarding: to.Ptr(false),
	// 							IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
	// 								{
	// 									Name: to.Ptr("vnet-eastus2euap-2-nic01-defaultIpConfiguration"),
	// 									Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
	// 										Primary: to.Ptr(true),
	// 										Subnet: &armcompute.APIEntityReference{
	// 											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet-eastus2euap-2/subnets/snet-eastus2euap-1"),
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
	// 				VMSize: to.Ptr(armcompute.VirtualMachineSizeTypes("Standard_D2ls_v5")),
	// 			},
	// 			ResilientVMDeletionStatus: to.Ptr(armcompute.ResilientVMDeletionStatusEnabled),
	// 			VMID: to.Ptr("eb282db2-12d4-4fc6-8bd5-0c6473a4078c"),
	// 			StorageProfile: &armcompute.StorageProfile{
	// 				ImageReference: &armcompute.ImageReference{
	// 					Publisher: to.Ptr("canonical"),
	// 					Offer: to.Ptr("0001-com-ubuntu-server-focal"),
	// 					SKU: to.Ptr("20_04-lts-gen2"),
	// 					Version: to.Ptr("latest"),
	// 					ExactVersion: to.Ptr("20.04.202501110"),
	// 				},
	// 				OSDisk: &armcompute.OSDisk{
	// 					OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
	// 					Name: to.Ptr("{vmss-name}_{vmss-vm-name}_OsDisk_1_8e93ddcf18be4b0f9815910b3a0f8182"),
	// 					CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 					Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 					ManagedDisk: &armcompute.ManagedDiskParameters{
	// 						StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/{vmss-name}_{vmss-vm-name}_OsDisk_1_8e93ddcf18be4b0f9815910b3a0f8182"),
	// 					},
	// 					DiskSizeGB: to.Ptr[int32](30),
	// 				},
	// 				DataDisks: []*armcompute.DataDisk{
	// 				},
	// 				DiskControllerType: to.Ptr(armcompute.DiskControllerTypesSCSI),
	// 			},
	// 			OSProfile: &armcompute.OSProfile{
	// 				ComputerName: to.Ptr("statustes000001"),
	// 				AdminUsername: to.Ptr("SomeRandomUser"),
	// 				LinuxConfiguration: &armcompute.LinuxConfiguration{
	// 					DisablePasswordAuthentication: to.Ptr(false),
	// 					ProvisionVMAgent: to.Ptr(true),
	// 					PatchSettings: &armcompute.LinuxPatchSettings{
	// 						PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
	// 						AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
	// 					},
	// 					EnableVMAgentPlatformUpdates: to.Ptr(true),
	// 				},
	// 				Secrets: []*armcompute.VaultSecretGroup{
	// 				},
	// 				AllowExtensionOperations: to.Ptr(true),
	// 				RequireGuestProvisionSignal: to.Ptr(true),
	// 			},
	// 			SecurityProfile: &armcompute.SecurityProfile{
	// 				UefiSettings: &armcompute.UefiSettings{
	// 					SecureBootEnabled: to.Ptr(true),
	// 					VTpmEnabled: to.Ptr(true),
	// 				},
	// 				SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
	// 			},
	// 			NetworkProfile: &armcompute.NetworkProfile{
	// 				NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
	// 					{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/1/networkInterfaces/vnet-eastus2euap-2-nic01"),
	// 					},
	// 				},
	// 			},
	// 			DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 				BootDiagnostics: &armcompute.BootDiagnostics{
	// 					Enabled: to.Ptr(true),
	// 				},
	// 			},
	// 			TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-05T20:11:32.5722157+00:00"); return t}()),
	// 		},
	// 		Etag: to.Ptr("\"3\""),
	// 		Resources: []*armcompute.VirtualMachineExtension{
	// 			{
	// 				Name: to.Ptr("Microsoft.Azure.Security.Monitoring.AzureSecurityLinuxAgent"),
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/{vmss-vm-name}/extensions/Microsoft.Azure.Security.Monitoring.AzureSecurityLinuxAgent"),
	// 				Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
	// 				Location: to.Ptr("eastus2euap"),
	// 				Properties: &armcompute.VirtualMachineExtensionProperties{
	// 					AutoUpgradeMinorVersion: to.Ptr(true),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					EnableAutomaticUpgrade: to.Ptr(true),
	// 					Publisher: to.Ptr("Microsoft.Azure.Security.Monitoring"),
	// 					Type: to.Ptr("AzureSecurityLinuxAgent"),
	// 					TypeHandlerVersion: to.Ptr("2.0"),
	// 					Settings: map[string]any{
	// 						"enableGenevaUpload": true,
	// 						"enableAutoConfig": true,
	// 						"reportSuccessOnUnsupportedDistro": true,
	// 					},
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("Microsoft.Azure.Monitor.AzureMonitorLinuxAgent"),
	// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/{vmss-vm-name}/extensions/Microsoft.Azure.Monitor.AzureMonitorLinuxAgent"),
	// 				Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
	// 				Location: to.Ptr("eastus2euap"),
	// 				Properties: &armcompute.VirtualMachineExtensionProperties{
	// 					AutoUpgradeMinorVersion: to.Ptr(true),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					EnableAutomaticUpgrade: to.Ptr(true),
	// 					Publisher: to.Ptr("Microsoft.Azure.Monitor"),
	// 					Type: to.Ptr("AzureMonitorLinuxAgent"),
	// 					TypeHandlerVersion: to.Ptr("1.0"),
	// 					Settings: map[string]any{
	// 						"GCS_AUTO_CONFIG": true,
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
