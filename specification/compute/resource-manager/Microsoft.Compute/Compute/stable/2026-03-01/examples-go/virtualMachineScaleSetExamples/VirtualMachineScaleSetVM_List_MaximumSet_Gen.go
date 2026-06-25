package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_List_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetVMsClient_NewListPager_virtualMachineScaleSetVMListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualMachineScaleSetVMsClient().NewListPager("rgcompute", "aaaaaaaaaaaaaaaaaaaaaa", &armcompute.VirtualMachineScaleSetVMsClientListOptions{
		Filter: to.Ptr("aaaaaaaaaaaaaa"),
		Select: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
		Expand: to.Ptr("aaaaaaaaaaaaa")})
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
		// page = armcompute.VirtualMachineScaleSetVMsClientListResponse{
		// 	VirtualMachineScaleSetVMListResult: armcompute.VirtualMachineScaleSetVMListResult{
		// 		Value: []*armcompute.VirtualMachineScaleSetVM{
		// 			{
		// 				Name: to.Ptr("{vmss-vm-name}"),
		// 				ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/0"),
		// 				Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets/virtualMachines"),
		// 				Location: to.Ptr("westus"),
		// 				Tags: map[string]*string{
		// 				},
		// 				Properties: &armcompute.VirtualMachineScaleSetVMProperties{
		// 					LatestModelApplied: to.Ptr(true),
		// 					ModelDefinitionApplied: to.Ptr("VirtualMachineScaleSet"),
		// 					NetworkProfileConfiguration: &armcompute.VirtualMachineScaleSetVMNetworkProfileConfiguration{
		// 						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
		// 							{
		// 								Name: to.Ptr("vmsstestnetconfig5415"),
		// 								Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
		// 									Primary: to.Ptr(true),
		// 									EnableAcceleratedNetworking: to.Ptr(true),
		// 									DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
		// 										DNSServers: []*string{
		// 										},
		// 									},
		// 									EnableIPForwarding: to.Ptr(true),
		// 									IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
		// 										{
		// 											Name: to.Ptr("vmsstestnetconfig9693"),
		// 											Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
		// 												Subnet: &armcompute.APIEntityReference{
		// 													ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/vn4071/subnets/sn5503"),
		// 												},
		// 												PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
		// 												Primary: to.Ptr(true),
		// 												PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
		// 													Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 													Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
		// 														IdleTimeoutInMinutes: to.Ptr[int32](18),
		// 														DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
		// 															DomainNameLabel: to.Ptr("aaaaaaaaaaaaaaaaaa"),
		// 														},
		// 														IPTags: []*armcompute.VirtualMachineScaleSetIPTag{
		// 															{
		// 																IPTagType: to.Ptr("aaaaaaa"),
		// 																Tag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 															},
		// 														},
		// 														PublicIPPrefix: &armcompute.SubResource{
		// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														},
		// 														PublicIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
		// 														DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 													},
		// 													SKU: &armcompute.PublicIPAddressSKU{
		// 														Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
		// 														Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
		// 													},
		// 												},
		// 												ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 													},
		// 												},
		// 												ApplicationSecurityGroups: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 													},
		// 												},
		// 												LoadBalancerBackendAddressPools: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 													},
		// 												},
		// 												LoadBalancerInboundNatPools: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 													},
		// 												},
		// 											},
		// 										},
		// 									},
		// 									EnableFpga: to.Ptr(true),
		// 									NetworkSecurityGroup: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 									DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 								},
		// 							},
		// 						},
		// 					},
		// 					VMID: to.Ptr("42af9fdf-b906-4ad7-9905-8316209ff619"),
		// 					HardwareProfile: &armcompute.HardwareProfile{
		// 						VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesBasicA0),
		// 						VMSizeProperties: &armcompute.VMSizeProperties{
		// 							VCPUsAvailable: to.Ptr[int32](9),
		// 							VCPUsPerCore: to.Ptr[int32](12),
		// 						},
		// 					},
		// 					StorageProfile: &armcompute.StorageProfile{
		// 						ImageReference: &armcompute.ImageReference{
		// 							Publisher: to.Ptr("MicrosoftWindowsServer"),
		// 							Offer: to.Ptr("WindowsServer"),
		// 							SKU: to.Ptr("2012-R2-Datacenter"),
		// 							Version: to.Ptr("4.127.20180315"),
		// 							ExactVersion: to.Ptr("4.127.20180315"),
		// 							SharedGalleryImageID: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 							ID: to.Ptr("a"),
		// 						},
		// 						OSDisk: &armcompute.OSDisk{
		// 							OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
		// 							Name: to.Ptr("vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc"),
		// 							CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
		// 							Caching: to.Ptr(armcompute.CachingTypesNone),
		// 							ManagedDisk: &armcompute.ManagedDiskParameters{
		// 								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
		// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc"),
		// 								DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 									ID: to.Ptr("aaaaaaaaaaaa"),
		// 								},
		// 							},
		// 							DiskSizeGB: to.Ptr[int32](127),
		// 							EncryptionSettings: &armcompute.DiskEncryptionSettings{
		// 								DiskEncryptionKey: &armcompute.KeyVaultSecretReference{
		// 									SecretURL: to.Ptr("aaaaaaaa"),
		// 									SourceVault: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 								},
		// 								KeyEncryptionKey: &armcompute.KeyVaultKeyReference{
		// 									KeyURL: to.Ptr("aaaaaaaaaaaaaa"),
		// 									SourceVault: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 								},
		// 								Enabled: to.Ptr(true),
		// 							},
		// 							Vhd: &armcompute.VirtualHardDisk{
		// 								URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 							},
		// 							Image: &armcompute.VirtualHardDisk{
		// 								URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 							},
		// 							WriteAcceleratorEnabled: to.Ptr(true),
		// 							DiffDiskSettings: &armcompute.DiffDiskSettings{
		// 								Option: to.Ptr(armcompute.DiffDiskOptionsLocal),
		// 								Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
		// 							},
		// 							DeleteOption: to.Ptr(armcompute.DiskDeleteOptionTypesDelete),
		// 						},
		// 						DataDisks: []*armcompute.DataDisk{
		// 							{
		// 								Lun: to.Ptr[int32](1),
		// 								Name: to.Ptr("vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
		// 								CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
		// 								Caching: to.Ptr(armcompute.CachingTypesNone),
		// 								ManagedDisk: &armcompute.ManagedDiskParameters{
		// 									StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
		// 									ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
		// 									DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
		// 										ID: to.Ptr("aaaaaaaaaaaa"),
		// 									},
		// 								},
		// 								DiskSizeGB: to.Ptr[int32](128),
		// 								ToBeDetached: to.Ptr(true),
		// 								Vhd: &armcompute.VirtualHardDisk{
		// 									URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 								},
		// 								Image: &armcompute.VirtualHardDisk{
		// 									URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
		// 								},
		// 								WriteAcceleratorEnabled: to.Ptr(true),
		// 								DiskIOPSReadWrite: to.Ptr[int64](18),
		// 								DiskMBpsReadWrite: to.Ptr[int64](29),
		// 								DetachOption: to.Ptr(armcompute.DiskDetachOptionTypesForceDetach),
		// 								DeleteOption: to.Ptr(armcompute.DiskDeleteOptionTypesDelete),
		// 							},
		// 						},
		// 					},
		// 					OSProfile: &armcompute.OSProfile{
		// 						ComputerName: to.Ptr("test000000"),
		// 						AdminUsername: to.Ptr("Foo12"),
		// 						WindowsConfiguration: &armcompute.WindowsConfiguration{
		// 							ProvisionVMAgent: to.Ptr(true),
		// 							EnableAutomaticUpdates: to.Ptr(true),
		// 							TimeZone: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
		// 								{
		// 									PassName: to.Ptr("OobeSystem"),
		// 									ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
		// 									SettingName: to.Ptr(armcompute.SettingNamesAutoLogon),
		// 									Content: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 								},
		// 							},
		// 							PatchSettings: &armcompute.PatchSettings{
		// 								PatchMode: to.Ptr(armcompute.WindowsVMGuestPatchModeManual),
		// 								EnableHotpatching: to.Ptr(true),
		// 								AssessmentMode: to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
		// 							},
		// 							WinRM: &armcompute.WinRMConfiguration{
		// 								Listeners: []*armcompute.WinRMListener{
		// 									{
		// 										Protocol: to.Ptr(armcompute.ProtocolTypesHTTP),
		// 										CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 						Secrets: []*armcompute.VaultSecretGroup{
		// 						},
		// 						AllowExtensionOperations: to.Ptr(true),
		// 						RequireGuestProvisionSignal: to.Ptr(true),
		// 						CustomData: to.Ptr("aaaa"),
		// 						LinuxConfiguration: &armcompute.LinuxConfiguration{
		// 							DisablePasswordAuthentication: to.Ptr(true),
		// 							SSH: &armcompute.SSHConfiguration{
		// 								PublicKeys: []*armcompute.SSHPublicKey{
		// 									{
		// 										Path: to.Ptr("aaa"),
		// 										KeyData: to.Ptr("aaaaaa"),
		// 									},
		// 								},
		// 							},
		// 							ProvisionVMAgent: to.Ptr(true),
		// 							PatchSettings: &armcompute.LinuxPatchSettings{
		// 								PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
		// 								AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
		// 							},
		// 						},
		// 					},
		// 					UserData: to.Ptr("RXhhbXBsZSBVc2VyRGF0YQ=="),
		// 					NetworkProfile: &armcompute.NetworkProfile{
		// 						NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
		// 							{
		// 								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/0/networkInterfaces/vmsstestnetconfig5415"),
		// 								Properties: &armcompute.NetworkInterfaceReferenceProperties{
		// 									Primary: to.Ptr(true),
		// 									DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 								},
		// 							},
		// 						},
		// 						NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
		// 						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineNetworkInterfaceConfiguration{
		// 							{
		// 								Name: to.Ptr("aaaaaaaaaaa"),
		// 								Properties: &armcompute.VirtualMachineNetworkInterfaceConfigurationProperties{
		// 									Primary: to.Ptr(true),
		// 									DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 									EnableAcceleratedNetworking: to.Ptr(true),
		// 									EnableFpga: to.Ptr(true),
		// 									EnableIPForwarding: to.Ptr(true),
		// 									NetworkSecurityGroup: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 									DNSSettings: &armcompute.VirtualMachineNetworkInterfaceDNSSettingsConfiguration{
		// 										DNSServers: []*string{
		// 											to.Ptr("aaaaaa"),
		// 										},
		// 									},
		// 									IPConfigurations: []*armcompute.VirtualMachineNetworkInterfaceIPConfiguration{
		// 										{
		// 											Name: to.Ptr("aa"),
		// 											Properties: &armcompute.VirtualMachineNetworkInterfaceIPConfigurationProperties{
		// 												Subnet: &armcompute.SubResource{
		// 													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 												},
		// 												Primary: to.Ptr(true),
		// 												PublicIPAddressConfiguration: &armcompute.VirtualMachinePublicIPAddressConfiguration{
		// 													Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 													Properties: &armcompute.VirtualMachinePublicIPAddressConfigurationProperties{
		// 														IdleTimeoutInMinutes: to.Ptr[int32](2),
		// 														DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
		// 														DNSSettings: &armcompute.VirtualMachinePublicIPAddressDNSSettingsConfiguration{
		// 															DomainNameLabel: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 														},
		// 														IPTags: []*armcompute.VirtualMachineIPTag{
		// 															{
		// 																IPTagType: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 																Tag: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 															},
		// 														},
		// 														PublicIPPrefix: &armcompute.SubResource{
		// 															ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 														},
		// 														PublicIPAddressVersion: to.Ptr(armcompute.IPVersionsIPv4),
		// 														PublicIPAllocationMethod: to.Ptr(armcompute.PublicIPAllocationMethodDynamic),
		// 													},
		// 													SKU: &armcompute.PublicIPAddressSKU{
		// 														Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
		// 														Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
		// 													},
		// 												},
		// 												PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionsIPv4),
		// 												ApplicationSecurityGroups: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 													},
		// 												},
		// 												ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 													},
		// 												},
		// 												LoadBalancerBackendAddressPools: []*armcompute.SubResource{
		// 													{
		// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 													},
		// 												},
		// 											},
		// 										},
		// 									},
		// 									DscpConfiguration: &armcompute.SubResource{
		// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 									},
		// 								},
		// 							},
		// 						},
		// 					},
		// 					DiagnosticsProfile: &armcompute.DiagnosticsProfile{
		// 						BootDiagnostics: &armcompute.BootDiagnostics{
		// 							Enabled: to.Ptr(true),
		// 							StorageURI: to.Ptr("aaaaaaaaaaaaa"),
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					InstanceView: &armcompute.VirtualMachineScaleSetVMInstanceView{
		// 						PlatformUpdateDomain: to.Ptr[int32](23),
		// 						PlatformFaultDomain: to.Ptr[int32](14),
		// 						RdpThumbPrint: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 						VMAgent: &armcompute.VirtualMachineAgentInstanceView{
		// 							VMAgentVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 							ExtensionHandlers: []*armcompute.VirtualMachineExtensionHandlerInstanceView{
		// 								{
		// 									Type: to.Ptr("aaaaaaaaaaaaa"),
		// 									TypeHandlerVersion: to.Ptr("aaaaa"),
		// 									Status: &armcompute.InstanceViewStatus{
		// 										Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 										DisplayStatus: to.Ptr("aaaaaa"),
		// 										Message: to.Ptr("a"),
		// 										Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 									},
		// 								},
		// 							},
		// 							Statuses: []*armcompute.InstanceViewStatus{
		// 								{
		// 									Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 									Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 									DisplayStatus: to.Ptr("aaaaaa"),
		// 									Message: to.Ptr("a"),
		// 									Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 								},
		// 							},
		// 						},
		// 						MaintenanceRedeployStatus: &armcompute.MaintenanceRedeployStatus{
		// 							IsCustomerInitiatedMaintenanceAllowed: to.Ptr(true),
		// 							PreMaintenanceWindowStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.531Z"); return t}()),
		// 							PreMaintenanceWindowEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.531Z"); return t}()),
		// 							MaintenanceWindowStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.531Z"); return t}()),
		// 							MaintenanceWindowEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.531Z"); return t}()),
		// 							LastOperationResultCode: to.Ptr(armcompute.MaintenanceOperationResultCodeTypesNone),
		// 							LastOperationMessage: to.Ptr("aaaaaa"),
		// 						},
		// 						Disks: []*armcompute.DiskInstanceView{
		// 							{
		// 								Name: to.Ptr("aaaaaaaaaaa"),
		// 								EncryptionSettings: []*armcompute.DiskEncryptionSettings{
		// 									{
		// 										DiskEncryptionKey: &armcompute.KeyVaultSecretReference{
		// 											SecretURL: to.Ptr("aaaaaaaa"),
		// 											SourceVault: &armcompute.SubResource{
		// 												ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 											},
		// 										},
		// 										KeyEncryptionKey: &armcompute.KeyVaultKeyReference{
		// 											KeyURL: to.Ptr("aaaaaaaaaaaaaa"),
		// 											SourceVault: &armcompute.SubResource{
		// 												ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 											},
		// 										},
		// 										Enabled: to.Ptr(true),
		// 									},
		// 								},
		// 								Statuses: []*armcompute.InstanceViewStatus{
		// 									{
		// 										Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 										DisplayStatus: to.Ptr("aaaaaa"),
		// 										Message: to.Ptr("a"),
		// 										Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 									},
		// 								},
		// 							},
		// 						},
		// 						Extensions: []*armcompute.VirtualMachineExtensionInstanceView{
		// 							{
		// 								Name: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 								Type: to.Ptr("aaaaaaaaa"),
		// 								TypeHandlerVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								Substatuses: []*armcompute.InstanceViewStatus{
		// 									{
		// 										Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 										DisplayStatus: to.Ptr("aaaaaa"),
		// 										Message: to.Ptr("a"),
		// 										Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 									},
		// 								},
		// 								Statuses: []*armcompute.InstanceViewStatus{
		// 									{
		// 										Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 										DisplayStatus: to.Ptr("aaaaaa"),
		// 										Message: to.Ptr("a"),
		// 										Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 									},
		// 								},
		// 							},
		// 						},
		// 						VMHealth: &armcompute.VirtualMachineHealthStatus{
		// 							Status: &armcompute.InstanceViewStatus{
		// 								Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 								Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 								DisplayStatus: to.Ptr("aaaaaa"),
		// 								Message: to.Ptr("a"),
		// 								Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 							},
		// 						},
		// 						BootDiagnostics: &armcompute.BootDiagnosticsInstanceView{
		// 							ConsoleScreenshotBlobURI: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 							SerialConsoleLogBlobURI: to.Ptr("aaaaaaaa"),
		// 							Status: &armcompute.InstanceViewStatus{
		// 								Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 								Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 								DisplayStatus: to.Ptr("aaaaaa"),
		// 								Message: to.Ptr("a"),
		// 								Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 							},
		// 						},
		// 						Statuses: []*armcompute.InstanceViewStatus{
		// 							{
		// 								Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 								Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 								DisplayStatus: to.Ptr("aaaaaa"),
		// 								Message: to.Ptr("a"),
		// 								Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 							},
		// 						},
		// 						AssignedHost: to.Ptr("aaaaaaa"),
		// 						PlacementGroupID: to.Ptr("aaa"),
		// 					},
		// 					AdditionalCapabilities: &armcompute.AdditionalCapabilities{
		// 						UltraSSDEnabled: to.Ptr(true),
		// 						HibernationEnabled: to.Ptr(true),
		// 					},
		// 					SecurityProfile: &armcompute.SecurityProfile{
		// 						UefiSettings: &armcompute.UefiSettings{
		// 							SecureBootEnabled: to.Ptr(true),
		// 							VTpmEnabled: to.Ptr(true),
		// 						},
		// 						EncryptionAtHost: to.Ptr(true),
		// 						SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
		// 					},
		// 					AvailabilitySet: &armcompute.SubResource{
		// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
		// 					},
		// 					LicenseType: to.Ptr("aaaaaaaaaa"),
		// 					ProtectionPolicy: &armcompute.VirtualMachineScaleSetVMProtectionPolicy{
		// 						ProtectFromScaleIn: to.Ptr(true),
		// 						ProtectFromScaleSetActions: to.Ptr(true),
		// 					},
		// 					TimeCreated: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-06-27T01:02:38.3138469+00:00"); return t}()),
		// 				},
		// 				Resources: []*armcompute.VirtualMachineExtension{
		// 					{
		// 						Name: to.Ptr("CustomScriptExtension-DSC"),
		// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM/extensions/CustomScriptExtension-DSC"),
		// 						Type: to.Ptr("Microsoft.Compute/virtualMachines/extensions"),
		// 						Location: to.Ptr("westus"),
		// 						Tags: map[string]*string{
		// 						},
		// 						Properties: &armcompute.VirtualMachineExtensionProperties{
		// 							AutoUpgradeMinorVersion: to.Ptr(true),
		// 							ProvisioningState: to.Ptr("Succeeded"),
		// 							Publisher: to.Ptr("Microsoft.Compute"),
		// 							Type: to.Ptr("CustomScriptExtension"),
		// 							TypeHandlerVersion: to.Ptr("1.9"),
		// 							Settings: map[string]any{
		// 							},
		// 							ForceUpdateTag: to.Ptr("aaaaaaa"),
		// 							EnableAutomaticUpgrade: to.Ptr(true),
		// 							ProtectedSettings: map[string]any{
		// 							},
		// 							InstanceView: &armcompute.VirtualMachineExtensionInstanceView{
		// 								Name: to.Ptr("aaaaaaaaaaaaaaaaa"),
		// 								Type: to.Ptr("aaaaaaaaa"),
		// 								TypeHandlerVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
		// 								Substatuses: []*armcompute.InstanceViewStatus{
		// 									{
		// 										Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 										DisplayStatus: to.Ptr("aaaaaa"),
		// 										Message: to.Ptr("a"),
		// 										Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 									},
		// 								},
		// 								Statuses: []*armcompute.InstanceViewStatus{
		// 									{
		// 										Code: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
		// 										Level: to.Ptr(armcompute.StatusLevelTypesInfo),
		// 										DisplayStatus: to.Ptr("aaaaaa"),
		// 										Message: to.Ptr("a"),
		// 										Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t}()),
		// 									},
		// 								},
		// 							},
		// 							SuppressFailures: to.Ptr(true),
		// 						},
		// 					},
		// 				},
		// 				InstanceID: to.Ptr("aaaaaaaaaaaa"),
		// 				SKU: &armcompute.SKU{
		// 					Name: to.Ptr("Classic"),
		// 					Tier: to.Ptr("aaaaaaaaaaaaaa"),
		// 					Capacity: to.Ptr[int64](29),
		// 				},
		// 				Plan: &armcompute.Plan{
		// 					Name: to.Ptr("aaaaaaaaaa"),
		// 					Publisher: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
		// 					Product: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 					PromotionCode: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("a"),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("a://example.com/aaaaaaaaaaaaaa"),
		// 	},
		// }
	}
}
