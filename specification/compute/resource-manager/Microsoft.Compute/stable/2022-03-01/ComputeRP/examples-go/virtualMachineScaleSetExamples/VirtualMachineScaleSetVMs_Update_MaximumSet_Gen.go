package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMs_Update_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetVMsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetVMsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"rgcompute",
		"aaaaaaaaaaaaaa",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		armcompute.VirtualMachineScaleSetVM{
			Location: to.Ptr("westus"),
			Tags:     map[string]*string{},
			Plan: &armcompute.Plan{
				Name:          to.Ptr("aaaaaaaaaa"),
				Product:       to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
				PromotionCode: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
				Publisher:     to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
			},
			Properties: &armcompute.VirtualMachineScaleSetVMProperties{
				AdditionalCapabilities: &armcompute.AdditionalCapabilities{
					HibernationEnabled: to.Ptr(true),
					UltraSSDEnabled:    to.Ptr(true),
				},
				AvailabilitySet: &armcompute.SubResource{
					ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
				},
				DiagnosticsProfile: &armcompute.DiagnosticsProfile{
					BootDiagnostics: &armcompute.BootDiagnostics{
						Enabled:    to.Ptr(true),
						StorageURI: to.Ptr("aaaaaaaaaaaaa"),
					},
				},
				HardwareProfile: &armcompute.HardwareProfile{
					VMSize: to.Ptr(armcompute.VirtualMachineSizeTypesBasicA0),
					VMSizeProperties: &armcompute.VMSizeProperties{
						VCPUsAvailable: to.Ptr[int32](9),
						VCPUsPerCore:   to.Ptr[int32](12),
					},
				},
				InstanceView: &armcompute.VirtualMachineScaleSetVMInstanceView{
					BootDiagnostics: &armcompute.BootDiagnosticsInstanceView{
						Status: &armcompute.InstanceViewStatus{
							Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
							DisplayStatus: to.Ptr("aaaaaa"),
							Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
							Message:       to.Ptr("a"),
							Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
						},
					},
					Disks: []*armcompute.DiskInstanceView{
						{
							Name: to.Ptr("aaaaaaaaaaa"),
							EncryptionSettings: []*armcompute.DiskEncryptionSettings{
								{
									DiskEncryptionKey: &armcompute.KeyVaultSecretReference{
										SecretURL: to.Ptr("aaaaaaaa"),
										SourceVault: &armcompute.SubResource{
											ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
										},
									},
									Enabled: to.Ptr(true),
									KeyEncryptionKey: &armcompute.KeyVaultKeyReference{
										KeyURL: to.Ptr("aaaaaaaaaaaaaa"),
										SourceVault: &armcompute.SubResource{
											ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
										},
									},
								}},
							Statuses: []*armcompute.InstanceViewStatus{
								{
									Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
									DisplayStatus: to.Ptr("aaaaaa"),
									Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
									Message:       to.Ptr("a"),
									Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
								}},
						}},
					MaintenanceRedeployStatus: &armcompute.MaintenanceRedeployStatus{
						IsCustomerInitiatedMaintenanceAllowed: to.Ptr(true),
						LastOperationMessage:                  to.Ptr("aaaaaa"),
						LastOperationResultCode:               to.Ptr(armcompute.MaintenanceOperationResultCodeTypesNone),
						MaintenanceWindowEndTime:              to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.531Z"); return t }()),
						MaintenanceWindowStartTime:            to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.531Z"); return t }()),
						PreMaintenanceWindowEndTime:           to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.531Z"); return t }()),
						PreMaintenanceWindowStartTime:         to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.531Z"); return t }()),
					},
					PlacementGroupID:     to.Ptr("aaa"),
					PlatformFaultDomain:  to.Ptr[int32](14),
					PlatformUpdateDomain: to.Ptr[int32](23),
					RdpThumbPrint:        to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
					Statuses: []*armcompute.InstanceViewStatus{
						{
							Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
							DisplayStatus: to.Ptr("aaaaaa"),
							Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
							Message:       to.Ptr("a"),
							Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
						}},
					VMAgent: &armcompute.VirtualMachineAgentInstanceView{
						ExtensionHandlers: []*armcompute.VirtualMachineExtensionHandlerInstanceView{
							{
								Type: to.Ptr("aaaaaaaaaaaaa"),
								Status: &armcompute.InstanceViewStatus{
									Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
									DisplayStatus: to.Ptr("aaaaaa"),
									Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
									Message:       to.Ptr("a"),
									Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
								},
								TypeHandlerVersion: to.Ptr("aaaaa"),
							}},
						Statuses: []*armcompute.InstanceViewStatus{
							{
								Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
								DisplayStatus: to.Ptr("aaaaaa"),
								Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
								Message:       to.Ptr("a"),
								Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
							}},
						VMAgentVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
					},
					VMHealth: &armcompute.VirtualMachineHealthStatus{
						Status: &armcompute.InstanceViewStatus{
							Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
							DisplayStatus: to.Ptr("aaaaaa"),
							Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
							Message:       to.Ptr("a"),
							Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
						},
					},
					Extensions: []*armcompute.VirtualMachineExtensionInstanceView{
						{
							Name: to.Ptr("aaaaaaaaaaaaaaaaa"),
							Type: to.Ptr("aaaaaaaaa"),
							Statuses: []*armcompute.InstanceViewStatus{
								{
									Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
									DisplayStatus: to.Ptr("aaaaaa"),
									Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
									Message:       to.Ptr("a"),
									Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
								}},
							Substatuses: []*armcompute.InstanceViewStatus{
								{
									Code:          to.Ptr("aaaaaaaaaaaaaaaaaaaaaaa"),
									DisplayStatus: to.Ptr("aaaaaa"),
									Level:         to.Ptr(armcompute.StatusLevelTypesInfo),
									Message:       to.Ptr("a"),
									Time:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-30T12:58:26.522Z"); return t }()),
								}},
							TypeHandlerVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
						}},
				},
				LicenseType: to.Ptr("aaaaaaaaaa"),
				NetworkProfile: &armcompute.NetworkProfile{
					NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineNetworkInterfaceConfiguration{
						{
							Name: to.Ptr("aaaaaaaaaaa"),
							Properties: &armcompute.VirtualMachineNetworkInterfaceConfigurationProperties{
								DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
								DNSSettings: &armcompute.VirtualMachineNetworkInterfaceDNSSettingsConfiguration{
									DNSServers: []*string{
										to.Ptr("aaaaaa")},
								},
								DscpConfiguration: &armcompute.SubResource{
									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
								},
								EnableAcceleratedNetworking: to.Ptr(true),
								EnableFpga:                  to.Ptr(true),
								EnableIPForwarding:          to.Ptr(true),
								IPConfigurations: []*armcompute.VirtualMachineNetworkInterfaceIPConfiguration{
									{
										Name: to.Ptr("aa"),
										Properties: &armcompute.VirtualMachineNetworkInterfaceIPConfigurationProperties{
											ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
												}},
											ApplicationSecurityGroups: []*armcompute.SubResource{
												{
													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
												}},
											LoadBalancerBackendAddressPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
												}},
											Primary:                 to.Ptr(true),
											PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionsIPv4),
											PublicIPAddressConfiguration: &armcompute.VirtualMachinePublicIPAddressConfiguration{
												Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
												Properties: &armcompute.VirtualMachinePublicIPAddressConfigurationProperties{
													DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
													DNSSettings: &armcompute.VirtualMachinePublicIPAddressDNSSettingsConfiguration{
														DomainNameLabel: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
													},
													IdleTimeoutInMinutes: to.Ptr[int32](2),
													IPTags: []*armcompute.VirtualMachineIPTag{
														{
															IPTagType: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
															Tag:       to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
														}},
													PublicIPAddressVersion:   to.Ptr(armcompute.IPVersionsIPv4),
													PublicIPAllocationMethod: to.Ptr(armcompute.PublicIPAllocationMethodDynamic),
													PublicIPPrefix: &armcompute.SubResource{
														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
													},
												},
												SKU: &armcompute.PublicIPAddressSKU{
													Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
													Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
												},
											},
											Subnet: &armcompute.SubResource{
												ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
											},
										},
									}},
								NetworkSecurityGroup: &armcompute.SubResource{
									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
								},
								Primary: to.Ptr(true),
							},
						}},
					NetworkInterfaces: []*armcompute.NetworkInterfaceReference{
						{
							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}/virtualMachines/0/networkInterfaces/vmsstestnetconfig5415"),
							Properties: &armcompute.NetworkInterfaceReferenceProperties{
								DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
								Primary:      to.Ptr(true),
							},
						}},
				},
				NetworkProfileConfiguration: &armcompute.VirtualMachineScaleSetVMNetworkProfileConfiguration{
					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
						{
							ID:   to.Ptr("aaaaaaaa"),
							Name: to.Ptr("vmsstestnetconfig5415"),
							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
								DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
								DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
									DNSServers: []*string{},
								},
								EnableAcceleratedNetworking: to.Ptr(true),
								EnableFpga:                  to.Ptr(true),
								EnableIPForwarding:          to.Ptr(true),
								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
									{
										ID:   to.Ptr("aaaaaaaaa"),
										Name: to.Ptr("vmsstestnetconfig9693"),
										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
											ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
												}},
											ApplicationSecurityGroups: []*armcompute.SubResource{
												{
													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
												}},
											LoadBalancerBackendAddressPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
												}},
											LoadBalancerInboundNatPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
												}},
											Primary:                 to.Ptr(true),
											PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
											PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
												Name: to.Ptr("aaaaaaaaaaaaaaaaaa"),
												Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
													DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
													DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
														DomainNameLabel: to.Ptr("aaaaaaaaaaaaaaaaaa"),
													},
													IdleTimeoutInMinutes: to.Ptr[int32](18),
													IPTags: []*armcompute.VirtualMachineScaleSetIPTag{
														{
															IPTagType: to.Ptr("aaaaaaa"),
															Tag:       to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
														}},
													PublicIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
													PublicIPPrefix: &armcompute.SubResource{
														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
													},
												},
												SKU: &armcompute.PublicIPAddressSKU{
													Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
													Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
												},
											},
											Subnet: &armcompute.APIEntityReference{
												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/vn4071/subnets/sn5503"),
											},
										},
									}},
								NetworkSecurityGroup: &armcompute.SubResource{
									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
								},
								Primary: to.Ptr(true),
							},
						}},
				},
				OSProfile: &armcompute.OSProfile{
					AdminPassword:            to.Ptr("aaaaaaaaaaaaaaaa"),
					AdminUsername:            to.Ptr("Foo12"),
					AllowExtensionOperations: to.Ptr(true),
					ComputerName:             to.Ptr("test000000"),
					CustomData:               to.Ptr("aaaa"),
					LinuxConfiguration: &armcompute.LinuxConfiguration{
						DisablePasswordAuthentication: to.Ptr(true),
						PatchSettings: &armcompute.LinuxPatchSettings{
							AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
							PatchMode:      to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
						},
						ProvisionVMAgent: to.Ptr(true),
						SSH: &armcompute.SSHConfiguration{
							PublicKeys: []*armcompute.SSHPublicKey{
								{
									Path:    to.Ptr("aaa"),
									KeyData: to.Ptr("aaaaaa"),
								}},
						},
					},
					RequireGuestProvisionSignal: to.Ptr(true),
					Secrets:                     []*armcompute.VaultSecretGroup{},
					WindowsConfiguration: &armcompute.WindowsConfiguration{
						AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
							{
								ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
								Content:       to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
								PassName:      to.Ptr("OobeSystem"),
								SettingName:   to.Ptr(armcompute.SettingNamesAutoLogon),
							}},
						EnableAutomaticUpdates: to.Ptr(true),
						PatchSettings: &armcompute.PatchSettings{
							AssessmentMode:    to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
							EnableHotpatching: to.Ptr(true),
							PatchMode:         to.Ptr(armcompute.WindowsVMGuestPatchModeManual),
						},
						ProvisionVMAgent: to.Ptr(true),
						TimeZone:         to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
						WinRM: &armcompute.WinRMConfiguration{
							Listeners: []*armcompute.WinRMListener{
								{
									CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
									Protocol:       to.Ptr(armcompute.ProtocolTypesHTTP),
								}},
						},
					},
				},
				ProtectionPolicy: &armcompute.VirtualMachineScaleSetVMProtectionPolicy{
					ProtectFromScaleIn:         to.Ptr(true),
					ProtectFromScaleSetActions: to.Ptr(true),
				},
				SecurityProfile: &armcompute.SecurityProfile{
					EncryptionAtHost: to.Ptr(true),
					SecurityType:     to.Ptr(armcompute.SecurityTypesTrustedLaunch),
					UefiSettings: &armcompute.UefiSettings{
						SecureBootEnabled: to.Ptr(true),
						VTpmEnabled:       to.Ptr(true),
					},
				},
				StorageProfile: &armcompute.StorageProfile{
					DataDisks: []*armcompute.DataDisk{
						{
							Name:         to.Ptr("vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
							Caching:      to.Ptr(armcompute.CachingTypesNone),
							CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
							DeleteOption: to.Ptr(armcompute.DiskDeleteOptionTypesDelete),
							DetachOption: to.Ptr(armcompute.DiskDetachOptionTypesForceDetach),
							DiskSizeGB:   to.Ptr[int32](128),
							Image: &armcompute.VirtualHardDisk{
								URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
							},
							Lun: to.Ptr[int32](1),
							ManagedDisk: &armcompute.ManagedDiskParameters{
								ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_disk2_6c4f554bdafa49baa780eb2d128ff39d"),
								DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
									ID: to.Ptr("aaaaaaaaaaaa"),
								},
								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
							},
							ToBeDetached: to.Ptr(true),
							Vhd: &armcompute.VirtualHardDisk{
								URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
							},
							WriteAcceleratorEnabled: to.Ptr(true),
						}},
					ImageReference: &armcompute.ImageReference{
						ID:                   to.Ptr("a"),
						Offer:                to.Ptr("WindowsServer"),
						Publisher:            to.Ptr("MicrosoftWindowsServer"),
						SharedGalleryImageID: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
						SKU:                  to.Ptr("2012-R2-Datacenter"),
						Version:              to.Ptr("4.127.20180315"),
					},
					OSDisk: &armcompute.OSDisk{
						Name:         to.Ptr("vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc"),
						Caching:      to.Ptr(armcompute.CachingTypesNone),
						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
						DeleteOption: to.Ptr(armcompute.DiskDeleteOptionTypesDelete),
						DiffDiskSettings: &armcompute.DiffDiskSettings{
							Option:    to.Ptr(armcompute.DiffDiskOptionsLocal),
							Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
						},
						DiskSizeGB: to.Ptr[int32](127),
						EncryptionSettings: &armcompute.DiskEncryptionSettings{
							DiskEncryptionKey: &armcompute.KeyVaultSecretReference{
								SecretURL: to.Ptr("aaaaaaaa"),
								SourceVault: &armcompute.SubResource{
									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
								},
							},
							Enabled: to.Ptr(true),
							KeyEncryptionKey: &armcompute.KeyVaultKeyReference{
								KeyURL: to.Ptr("aaaaaaaaaaaaaa"),
								SourceVault: &armcompute.SubResource{
									ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
								},
							},
						},
						Image: &armcompute.VirtualHardDisk{
							URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
						},
						ManagedDisk: &armcompute.ManagedDiskParameters{
							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vmss3176_vmss3176_0_OsDisk_1_6d72b805e50e4de6830303c5055077fc"),
							DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
								ID: to.Ptr("aaaaaaaaaaaa"),
							},
							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
						},
						OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
						Vhd: &armcompute.VirtualHardDisk{
							URI: to.Ptr("https://{storageAccountName}.blob.core.windows.net/{containerName}/{vhdName}.vhd"),
						},
						WriteAcceleratorEnabled: to.Ptr(true),
					},
				},
				UserData: to.Ptr("RXhhbXBsZSBVc2VyRGF0YQ=="),
			},
			SKU: &armcompute.SKU{
				Name:     to.Ptr("Classic"),
				Capacity: to.Ptr[int64](29),
				Tier:     to.Ptr("aaaaaaaaaaaaaa"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
