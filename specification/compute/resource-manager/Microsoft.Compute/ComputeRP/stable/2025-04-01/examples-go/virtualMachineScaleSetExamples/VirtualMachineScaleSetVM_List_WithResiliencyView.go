package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_List_WithResiliencyView.json
func ExampleVirtualMachineScaleSetVMsClient_NewListPager_listVmssVMsWithResilientVmDeletionStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineScaleSetVMsClient().NewListPager("resourceGroupname", "vmssName", &armcompute.VirtualMachineScaleSetVMsClientListOptions{Filter: nil,
		Select: nil,
		Expand: nil,
	})
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
		// page.VirtualMachineScaleSetVMListResult = armcompute.VirtualMachineScaleSetVMListResult{
		// 	Value: []*armcompute.VirtualMachineScaleSetVM{
		// 		{
		// 			Name: to.Ptr("{vmss-vm-name}_1"),
		// 			Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/1"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Etag: to.Ptr("\"4\""),
		// 			Identity: &armcompute.VirtualMachineIdentity{
		// 				Type: to.Ptr(armcompute.ResourceIdentityTypeUserAssigned),
		// 				UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
		// 					"/subscriptions/{subscription-id}/resourceGroups/AzSecPackAutoConfigRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/AzSecPackAutoConfigUA-eastus2euap": &armcompute.UserAssignedIdentitiesValue{
		// 						ClientID: to.Ptr("215414c9-8a82-4439-95ea-d09e3543a6e2"),
		// 						PrincipalID: to.Ptr("f31e5089-a1e5-44a6-9048-a767ce07d26c"),
		// 					},
		// 				},
		// 			},
		// 			InstanceID: to.Ptr("1"),
		// 			Properties: &armcompute.VirtualMachineScaleSetVMProperties{
		// 				DiagnosticsProfile: &armcompute.DiagnosticsProfile{
		// 					BootDiagnostics: &armcompute.BootDiagnostics{
		// 						Enabled: to.Ptr(true),
		// 					},
		// 				},
		// 				HardwareProfile: &armcompute.HardwareProfile{
		// 					VMSize: to.Ptr(armcompute.VirtualMachineSizeTypes("Standard_D2ls_v5")),
		// 				},
		// 				LatestModelApplied: to.Ptr(true),
		// 				ModelDefinitionApplied: to.Ptr("VirtualMachineScaleSet"),
		// 				NetworkProfile: &armcompute.NetworkProfile{
		// 					NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/1/networkInterfaces/vnet-eastus2euap-2-nic01"),
		// 					}},
		// 				},
		// 				NetworkProfileConfiguration: &armcompute.VirtualMachineScaleSetVMNetworkProfileConfiguration{
		// 					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
		// 						{
		// 							Name: to.Ptr("vnet-eastus2euap-2-nic01"),
		// 							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
		// 								DisableTCPStateTracking: to.Ptr(false),
		// 								DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
		// 									DNSServers: []*string{
		// 									},
		// 								},
		// 								EnableAcceleratedNetworking: to.Ptr(true),
		// 								EnableIPForwarding: to.Ptr(false),
		// 								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
		// 									{
		// 										Name: to.Ptr("vnet-eastus2euap-2-nic01-defaultIpConfiguration"),
		// 										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
		// 											Primary: to.Ptr(true),
		// 											PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
		// 											Subnet: &armcompute.APIEntityReference{
		// 												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet-eastus2euap-2/subnets/snet-eastus2euap-1"),
		// 											},
		// 										},
		// 								}},
		// 								NetworkSecurityGroup: &armcompute.SubResource{
		// 									ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkSecurityGroups/basicNsgvnet-eastus2euap-2-nic01"),
		// 								},
		// 								Primary: to.Ptr(true),
		// 							},
		// 					}},
		// 				},
		// 				OSProfile: &armcompute.OSProfile{
		// 					AdminUsername: to.Ptr("SomeRandomUser"),
		// 					AllowExtensionOperations: to.Ptr(true),
		// 					ComputerName: to.Ptr("statustes000001"),
		// 					LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 						DisablePasswordAuthentication: to.Ptr(false),
		// 						EnableVMAgentPlatformUpdates: to.Ptr(true),
		// 						PatchSettings: &armcompute.LinuxPatchSettings{
		// 							AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
		// 							PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
		// 						},
		// 						ProvisionVMAgent: to.Ptr(true),
		// 					},
		// 					RequireGuestProvisionSignal: to.Ptr(true),
		// 					Secrets: []*armcompute.VaultSecretGroup{
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ResilientVMDeletionStatus: to.Ptr(armcompute.ResilientVMDeletionStatusEnabled),
		// 				SecurityProfile: &armcompute.SecurityProfile{
		// 					SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
		// 					UefiSettings: &armcompute.UefiSettings{
		// 						SecureBootEnabled: to.Ptr(true),
		// 						VTpmEnabled: to.Ptr(true),
		// 					},
		// 				},
		// 				StorageProfile: &armcompute.StorageProfile{
		// 					DataDisks: []*armcompute.DataDisk{
		// 					},
		// 					DiskControllerType: to.Ptr(armcompute.DiskControllerTypesSCSI),
		// 					ImageReference: &armcompute.ImageReference{
		// 						ExactVersion: to.Ptr("20.04.202501110"),
		// 						Offer: to.Ptr("0001-com-ubuntu-server-focal"),
		// 						Publisher: to.Ptr("canonical"),
		// 						SKU: to.Ptr("20_04-lts-gen2"),
		// 						Version: to.Ptr("latest"),
		// 					},
		// 					OSDisk: &armcompute.OSDisk{
		// 						Name: to.Ptr("{vmss-name}_{vmss-vm-name}_1_OsDisk_1_8e93ddcf18be4b0f9815910b3a0f8182"),
		// 						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
		// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 						DiskSizeGB: to.Ptr[int32](30),
		// 						ManagedDisk: &armcompute.ManagedDiskParameters{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/{vmss-name}_{vmss-vm-name}_1_OsDisk_1_8e93ddcf18be4b0f9815910b3a0f8182"),
		// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
		// 						},
		// 						OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
		// 					},
		// 				},
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-05T20:11:32.572Z"); return t}()),
		// 				VMID: to.Ptr("eb282db2-12d4-4fc6-8bd5-0c6473a4078c"),
		// 			},
		// 			Resources: []*armcompute.VirtualMachineExtension{
		// 				{
		// 					Name: to.Ptr("Microsoft.Azure.Security.Monitoring.AzureSecurityLinuxAgent"),
		// 					Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/{vmss-vm-name}_1/extensions/Microsoft.Azure.Security.Monitoring.AzureSecurityLinuxAgent"),
		// 					Location: to.Ptr("eastus2euap"),
		// 					Properties: &armcompute.VirtualMachineExtensionProperties{
		// 						Type: to.Ptr("AzureSecurityLinuxAgent"),
		// 						AutoUpgradeMinorVersion: to.Ptr(true),
		// 						EnableAutomaticUpgrade: to.Ptr(true),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 						Publisher: to.Ptr("Microsoft.Azure.Security.Monitoring"),
		// 						Settings: map[string]any{
		// 							"enableAutoConfig": true,
		// 							"enableGenevaUpload": true,
		// 							"reportSuccessOnUnsupportedDistro": true,
		// 						},
		// 						TypeHandlerVersion: to.Ptr("2.0"),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("Microsoft.Azure.Monitor.AzureMonitorLinuxAgent"),
		// 					Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/{vmss-vm-name}_1/extensions/Microsoft.Azure.Monitor.AzureMonitorLinuxAgent"),
		// 					Location: to.Ptr("eastus2euap"),
		// 					Properties: &armcompute.VirtualMachineExtensionProperties{
		// 						Type: to.Ptr("AzureMonitorLinuxAgent"),
		// 						AutoUpgradeMinorVersion: to.Ptr(true),
		// 						EnableAutomaticUpgrade: to.Ptr(true),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 						Publisher: to.Ptr("Microsoft.Azure.Monitor"),
		// 						Settings: map[string]any{
		// 							"GCS_AUTO_CONFIG": true,
		// 						},
		// 						TypeHandlerVersion: to.Ptr("1.0"),
		// 					},
		// 			}},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Standard_D2ls_v5"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("{vmss-vm-name}_2"),
		// 			Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/2"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Tags: map[string]*string{
		// 				"FailForResilientVMDeletionAtDiskDetach": to.Ptr("true"),
		// 				"azsecpack": to.Ptr("nonprod"),
		// 				"platformsettings.host_environment.service.platform_optedin_for_rootcerts": to.Ptr("true"),
		// 			},
		// 			Etag: to.Ptr("\"4\""),
		// 			Identity: &armcompute.VirtualMachineIdentity{
		// 				Type: to.Ptr(armcompute.ResourceIdentityTypeUserAssigned),
		// 				UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
		// 					"/subscriptions/{subscription-id}/resourceGroups/AzSecPackAutoConfigRG/providers/Microsoft.ManagedIdentity/userAssignedIdentities/AzSecPackAutoConfigUA-eastus2euap": &armcompute.UserAssignedIdentitiesValue{
		// 						ClientID: to.Ptr("215414c9-8a82-4439-95ea-d09e3543a6e2"),
		// 						PrincipalID: to.Ptr("f31e5089-a1e5-44a6-9048-a767ce07d26c"),
		// 					},
		// 				},
		// 			},
		// 			InstanceID: to.Ptr("2"),
		// 			Properties: &armcompute.VirtualMachineScaleSetVMProperties{
		// 				DiagnosticsProfile: &armcompute.DiagnosticsProfile{
		// 					BootDiagnostics: &armcompute.BootDiagnostics{
		// 						Enabled: to.Ptr(true),
		// 					},
		// 				},
		// 				HardwareProfile: &armcompute.HardwareProfile{
		// 					VMSize: to.Ptr(armcompute.VirtualMachineSizeTypes("Standard_D2ls_v5")),
		// 				},
		// 				LatestModelApplied: to.Ptr(true),
		// 				ModelDefinitionApplied: to.Ptr("VirtualMachineScaleSet"),
		// 				NetworkProfile: &armcompute.NetworkProfile{
		// 					NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
		// 						{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/2/networkInterfaces/vnet-eastus2euap-2-nic01"),
		// 					}},
		// 				},
		// 				NetworkProfileConfiguration: &armcompute.VirtualMachineScaleSetVMNetworkProfileConfiguration{
		// 					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
		// 						{
		// 							Name: to.Ptr("vnet-eastus2euap-2-nic01"),
		// 							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
		// 								DisableTCPStateTracking: to.Ptr(false),
		// 								DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
		// 									DNSServers: []*string{
		// 									},
		// 								},
		// 								EnableAcceleratedNetworking: to.Ptr(true),
		// 								EnableIPForwarding: to.Ptr(false),
		// 								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
		// 									{
		// 										Name: to.Ptr("vnet-eastus2euap-2-nic01-defaultIpConfiguration"),
		// 										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
		// 											Primary: to.Ptr(true),
		// 											PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
		// 											Subnet: &armcompute.APIEntityReference{
		// 												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet-eastus2euap-2/subnets/snet-eastus2euap-1"),
		// 											},
		// 										},
		// 								}},
		// 								NetworkSecurityGroup: &armcompute.SubResource{
		// 									ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkSecurityGroups/basicNsgvnet-eastus2euap-2-nic01"),
		// 								},
		// 								Primary: to.Ptr(true),
		// 							},
		// 					}},
		// 				},
		// 				OSProfile: &armcompute.OSProfile{
		// 					AdminUsername: to.Ptr("SomeRandomUser"),
		// 					AllowExtensionOperations: to.Ptr(true),
		// 					ComputerName: to.Ptr("statustes000002"),
		// 					LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 						DisablePasswordAuthentication: to.Ptr(false),
		// 						EnableVMAgentPlatformUpdates: to.Ptr(true),
		// 						PatchSettings: &armcompute.LinuxPatchSettings{
		// 							AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
		// 							PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
		// 						},
		// 						ProvisionVMAgent: to.Ptr(true),
		// 					},
		// 					RequireGuestProvisionSignal: to.Ptr(true),
		// 					Secrets: []*armcompute.VaultSecretGroup{
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ResilientVMDeletionStatus: to.Ptr(armcompute.ResilientVMDeletionStatusEnabled),
		// 				SecurityProfile: &armcompute.SecurityProfile{
		// 					SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
		// 					UefiSettings: &armcompute.UefiSettings{
		// 						SecureBootEnabled: to.Ptr(true),
		// 						VTpmEnabled: to.Ptr(true),
		// 					},
		// 				},
		// 				StorageProfile: &armcompute.StorageProfile{
		// 					DataDisks: []*armcompute.DataDisk{
		// 					},
		// 					DiskControllerType: to.Ptr(armcompute.DiskControllerTypesSCSI),
		// 					ImageReference: &armcompute.ImageReference{
		// 						ExactVersion: to.Ptr("20.04.202501110"),
		// 						Offer: to.Ptr("0001-com-ubuntu-server-focal"),
		// 						Publisher: to.Ptr("canonical"),
		// 						SKU: to.Ptr("20_04-lts-gen2"),
		// 						Version: to.Ptr("latest"),
		// 					},
		// 					OSDisk: &armcompute.OSDisk{
		// 						Name: to.Ptr("{vmss-name}_{vmss-vm-name}_2_OsDisk_1_fb3bbb00f81e465ba963e493bc9b30fa"),
		// 						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
		// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 						DiskSizeGB: to.Ptr[int32](30),
		// 						ManagedDisk: &armcompute.ManagedDiskParameters{
		// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/{vmss-name}_{vmss-vm-name}_2_OsDisk_1_fb3bbb00f81e465ba963e493bc9b30fa"),
		// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesPremiumLRS),
		// 						},
		// 						OSType: to.Ptr(armcompute.OperatingSystemTypesLinux),
		// 					},
		// 				},
		// 				TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-02-05T20:25:40.022Z"); return t}()),
		// 				VMID: to.Ptr("8cc48891-30a3-4f5a-9197-cb92168b7cb3"),
		// 			},
		// 			Resources: []*armcompute.VirtualMachineExtension{
		// 				{
		// 					Name: to.Ptr("Microsoft.Azure.Security.Monitoring.AzureSecurityLinuxAgent"),
		// 					Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/{vmss-vm-name}_2/extensions/Microsoft.Azure.Security.Monitoring.AzureSecurityLinuxAgent"),
		// 					Location: to.Ptr("eastus2euap"),
		// 					Properties: &armcompute.VirtualMachineExtensionProperties{
		// 						Type: to.Ptr("AzureSecurityLinuxAgent"),
		// 						AutoUpgradeMinorVersion: to.Ptr(true),
		// 						EnableAutomaticUpgrade: to.Ptr(true),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 						Publisher: to.Ptr("Microsoft.Azure.Security.Monitoring"),
		// 						Settings: map[string]any{
		// 							"enableAutoConfig": true,
		// 							"enableGenevaUpload": true,
		// 							"reportSuccessOnUnsupportedDistro": true,
		// 						},
		// 						TypeHandlerVersion: to.Ptr("2.0"),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("Microsoft.Azure.Monitor.AzureMonitorLinuxAgent"),
		// 					Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
		// 					ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/{vmss-vm-name}_2/extensions/Microsoft.Azure.Monitor.AzureMonitorLinuxAgent"),
		// 					Location: to.Ptr("eastus2euap"),
		// 					Properties: &armcompute.VirtualMachineExtensionProperties{
		// 						Type: to.Ptr("AzureMonitorLinuxAgent"),
		// 						AutoUpgradeMinorVersion: to.Ptr(true),
		// 						EnableAutomaticUpgrade: to.Ptr(true),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 						Publisher: to.Ptr("Microsoft.Azure.Monitor"),
		// 						Settings: map[string]any{
		// 							"GCS_AUTO_CONFIG": true,
		// 						},
		// 						TypeHandlerVersion: to.Ptr("1.0"),
		// 					},
		// 			}},
		// 			SKU: &armcompute.SKU{
		// 				Name: to.Ptr("Standard_D2ls_v5"),
		// 				Tier: to.Ptr("Standard"),
		// 			},
		// 	}},
		// }
	}
}
