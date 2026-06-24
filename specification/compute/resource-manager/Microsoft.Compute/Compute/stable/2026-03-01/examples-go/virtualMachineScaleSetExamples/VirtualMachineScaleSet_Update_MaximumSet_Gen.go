package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v8"
)

// Generated from example definition: 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Update_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginUpdate_virtualMachineScaleSetUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetsClient().BeginUpdate(ctx, "rgcompute", "aaaaaaaaaaaaa", armcompute.VirtualMachineScaleSetUpdate{
		SKU: &armcompute.SKU{
			Name:     to.Ptr("DSv3-Type1"),
			Tier:     to.Ptr("aaa"),
			Capacity: to.Ptr[int64](7),
		},
		Plan: &armcompute.Plan{
			Publisher:     to.Ptr("microsoft-ads"),
			Product:       to.Ptr("windows-data-science-vm"),
			Name:          to.Ptr("windows2016"),
			PromotionCode: to.Ptr("aaaaaaaaaa"),
		},
		Properties: &armcompute.VirtualMachineScaleSetUpdateProperties{
			UpgradePolicy: &armcompute.UpgradePolicy{
				Mode: to.Ptr(armcompute.UpgradeModeManual),
				RollingUpgradePolicy: &armcompute.RollingUpgradePolicy{
					MaxBatchInstancePercent:               to.Ptr[int32](49),
					MaxUnhealthyInstancePercent:           to.Ptr[int32](81),
					MaxUnhealthyUpgradedInstancePercent:   to.Ptr[int32](98),
					PauseTimeBetweenBatches:               to.Ptr("aaaaaaaaaaaaaaa"),
					EnableCrossZoneUpgrade:                to.Ptr(true),
					PrioritizeUnhealthyInstances:          to.Ptr(true),
					RollbackFailedInstancesOnPolicyBreach: to.Ptr(true),
					MaxSurge:                              to.Ptr(true),
				},
				AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
					EnableAutomaticOSUpgrade: to.Ptr(true),
					DisableAutomaticRollback: to.Ptr(true),
					OSRollingUpgradeDeferral: to.Ptr(true),
				},
			},
			AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
				Enabled:     to.Ptr(true),
				GracePeriod: to.Ptr("PT30M"),
			},
			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetUpdateVMProfile{
				OSProfile: &armcompute.VirtualMachineScaleSetUpdateOSProfile{
					CustomData: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
					WindowsConfiguration: &armcompute.WindowsConfiguration{
						ProvisionVMAgent:       to.Ptr(true),
						EnableAutomaticUpdates: to.Ptr(true),
						PatchSettings: &armcompute.PatchSettings{
							PatchMode:         to.Ptr(armcompute.WindowsVMGuestPatchModeAutomaticByPlatform),
							EnableHotpatching: to.Ptr(true),
							AssessmentMode:    to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
							AutomaticByPlatformSettings: &armcompute.WindowsVMGuestPatchAutomaticByPlatformSettings{
								RebootSetting: to.Ptr(armcompute.WindowsVMGuestPatchAutomaticByPlatformRebootSettingNever),
							},
						},
						TimeZone: to.Ptr("aaaaaaaaaaaaaaaa"),
						AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
							{
								PassName:      to.Ptr("OobeSystem"),
								ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
								SettingName:   to.Ptr(armcompute.SettingNamesAutoLogon),
								Content:       to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
							},
						},
						WinRM: &armcompute.WinRMConfiguration{
							Listeners: []*armcompute.WinRMListener{
								{
									Protocol:       to.Ptr(armcompute.ProtocolTypesHTTP),
									CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
								},
							},
						},
					},
					LinuxConfiguration: &armcompute.LinuxConfiguration{
						SSH: &armcompute.SSHConfiguration{
							PublicKeys: []*armcompute.SSHPublicKey{
								{
									Path:    to.Ptr("/home/{your-username}/.ssh/authorized_keys"),
									KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCeClRAk2ipUs/l5voIsDC5q9RI+YSRd1Bvd/O+axgY4WiBzG+4FwJWZm/mLLe5DoOdHQwmU2FrKXZSW4w2sYE70KeWnrFViCOX5MTVvJgPE8ClugNl8RWth/tU849DvM9sT7vFgfVSHcAS2yDRyDlueii+8nF2ym8XWAPltFVCyLHRsyBp5YPqK8JFYIa1eybKsY3hEAxRCA+/7bq8et+Gj3coOsuRmrehav7rE6N12Pb80I6ofa6SM5XNYq4Xk0iYNx7R3kdz0Jj9XgZYWjAHjJmT0gTRoOnt6upOuxK7xI/ykWrllgpXrCPu3Ymz+c+ujaqcxDopnAl2lmf69/J1"),
								},
							},
						},
						DisablePasswordAuthentication: to.Ptr(true),
						ProvisionVMAgent:              to.Ptr(true),
						PatchSettings: &armcompute.LinuxPatchSettings{
							PatchMode:      to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
							AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
						},
					},
					Secrets: []*armcompute.VaultSecretGroup{
						{
							SourceVault: &armcompute.SubResource{
								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
							},
							VaultCertificates: []*armcompute.VaultCertificate{
								{
									CertificateURL:   to.Ptr("aaaaaaa"),
									CertificateStore: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
								},
							},
						},
					},
				},
				StorageProfile: &armcompute.VirtualMachineScaleSetUpdateStorageProfile{
					ImageReference: &armcompute.ImageReference{
						SKU:                  to.Ptr("2016-Datacenter"),
						Publisher:            to.Ptr("MicrosoftWindowsServer"),
						Version:              to.Ptr("latest"),
						Offer:                to.Ptr("WindowsServer"),
						SharedGalleryImageID: to.Ptr("aaaaaa"),
						ID:                   to.Ptr("aaaaaaaaaaaaaaaaaaa"),
					},
					OSDisk: &armcompute.VirtualMachineScaleSetUpdateOSDisk{
						Caching:                 to.Ptr(armcompute.CachingTypesReadWrite),
						WriteAcceleratorEnabled: to.Ptr(true),
						DiffDiskSettings: &armcompute.DiffDiskSettings{
							Option:    to.Ptr(armcompute.DiffDiskOptionsLocal),
							Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
						},
						DiskSizeGB: to.Ptr[int32](6),
						Image: &armcompute.VirtualHardDisk{
							URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd"),
						},
						VhdContainers: []*string{
							to.Ptr("aa"),
						},
						ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
							DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
								ID: to.Ptr("aaaaaaaaaaaa"),
							},
						},
					},
					DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
						{
							DiskSizeGB:              to.Ptr[int32](1023),
							CreateOption:            to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
							Lun:                     to.Ptr[int32](26),
							Name:                    to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
							Caching:                 to.Ptr(armcompute.CachingTypesNone),
							WriteAcceleratorEnabled: to.Ptr(true),
							ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
								DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
									ID: to.Ptr("aaaaaaaaaaaa"),
								},
							},
							DiskIOPSReadWrite: to.Ptr[int64](28),
							DiskMBpsReadWrite: to.Ptr[int64](15),
						},
					},
				},
				NetworkProfile: &armcompute.VirtualMachineScaleSetUpdateNetworkProfile{
					HealthProbe: &armcompute.APIEntityReference{
						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123"),
					},
					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetUpdateNetworkConfiguration{
						{
							Name: to.Ptr("aaaaaaaa"),
							Properties: &armcompute.VirtualMachineScaleSetUpdateNetworkConfigurationProperties{
								Primary:                     to.Ptr(true),
								EnableAcceleratedNetworking: to.Ptr(true),
								EnableFpga:                  to.Ptr(true),
								NetworkSecurityGroup: &armcompute.SubResource{
									ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
								},
								DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
									DNSServers: []*string{},
								},
								IPConfigurations: []*armcompute.VirtualMachineScaleSetUpdateIPConfiguration{
									{
										Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
										Properties: &armcompute.VirtualMachineScaleSetUpdateIPConfigurationProperties{
											Subnet: &armcompute.APIEntityReference{
												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123"),
											},
											Primary: to.Ptr(true),
											PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetUpdatePublicIPAddressConfiguration{
												Name: to.Ptr("a"),
												Properties: &armcompute.VirtualMachineScaleSetUpdatePublicIPAddressConfigurationProperties{
													IdleTimeoutInMinutes: to.Ptr[int32](3),
													DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
														DomainNameLabel: to.Ptr("aaaaaaaaaaaaaaaaaa"),
													},
													DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
												},
											},
											PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
											ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
												},
											},
											ApplicationSecurityGroups: []*armcompute.SubResource{
												{
													ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
												},
											},
											LoadBalancerBackendAddressPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
												},
											},
											LoadBalancerInboundNatPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
												},
											},
										},
									},
								},
								EnableIPForwarding: to.Ptr(true),
								DeleteOption:       to.Ptr(armcompute.DeleteOptionsDelete),
							},
						},
					},
					NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
				},
				SecurityProfile: &armcompute.SecurityProfile{
					EncryptionAtHost: to.Ptr(true),
					UefiSettings: &armcompute.UefiSettings{
						SecureBootEnabled: to.Ptr(true),
						VTpmEnabled:       to.Ptr(true),
					},
					SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
				},
				DiagnosticsProfile: &armcompute.DiagnosticsProfile{
					BootDiagnostics: &armcompute.BootDiagnostics{
						StorageURI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net"),
						Enabled:    to.Ptr(true),
					},
				},
				ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
					Extensions: []*armcompute.VirtualMachineScaleSetExtension{
						{
							Name: to.Ptr("{extension-name}"),
							Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
								AutoUpgradeMinorVersion: to.Ptr(true),
								Publisher:               to.Ptr("{extension-Publisher}"),
								Type:                    to.Ptr("{extension-Type}"),
								TypeHandlerVersion:      to.Ptr("{handler-version}"),
								Settings:                map[string]any{},
								ForceUpdateTag:          to.Ptr("aaaaaaaaa"),
								EnableAutomaticUpgrade:  to.Ptr(true),
								ProtectedSettings:       map[string]any{},
								ProvisionAfterExtensions: []*string{
									to.Ptr("aa"),
								},
								SuppressFailures: to.Ptr(true),
							},
						},
					},
					ExtensionsTimeBudget: to.Ptr("PT1H20M"),
				},
				LicenseType: to.Ptr("aaaaaaaaaaaa"),
				BillingProfile: &armcompute.BillingProfile{
					MaxPrice: to.Ptr[float64](-1),
				},
				ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
					TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
						NotBeforeTimeout: to.Ptr("PT10M"),
						Enable:           to.Ptr(true),
					},
				},
				UserData: to.Ptr("aaaaaaaaaaaaa"),
			},
			Overprovision:                          to.Ptr(true),
			DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(true),
			SinglePlacementGroup:                   to.Ptr(true),
			AdditionalCapabilities: &armcompute.AdditionalCapabilities{
				HibernationEnabled: to.Ptr(true),
				UltraSSDEnabled:    to.Ptr(true),
			},
			ScaleInPolicy: &armcompute.ScaleInPolicy{
				Rules: []*armcompute.VirtualMachineScaleSetScaleInRules{
					to.Ptr(armcompute.VirtualMachineScaleSetScaleInRulesOldestVM),
				},
				ForceDeletion: to.Ptr(true),
			},
			ProximityPlacementGroup: &armcompute.SubResource{
				ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
			},
		},
		Identity: &armcompute.VirtualMachineScaleSetIdentity{
			Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssigned),
			UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
				"key3951": {},
			},
		},
		Zones: []*string{
			to.Ptr("1"),
			to.Ptr("2"),
			to.Ptr("3"),
		},
		Tags: map[string]*string{
			"key246": to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcompute.VirtualMachineScaleSetsClientUpdateResponse{
	// 	VirtualMachineScaleSet: armcompute.VirtualMachineScaleSet{
	// 		SKU: &armcompute.SKU{
	// 			Tier: to.Ptr("Standard"),
	// 			Capacity: to.Ptr[int64](3),
	// 			Name: to.Ptr("Standard_D1_v2"),
	// 		},
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armcompute.VirtualMachineScaleSetProperties{
	// 			Overprovision: to.Ptr(true),
	// 			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
	// 				StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
	// 					ImageReference: &armcompute.ImageReference{
	// 						SKU: to.Ptr("2016-Datacenter"),
	// 						Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 						Version: to.Ptr("latest"),
	// 						Offer: to.Ptr("WindowsServer"),
	// 						ExactVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 						SharedGalleryImageID: to.Ptr("aaaaaa"),
	// 						ID: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 					},
	// 					OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
	// 						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 						ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
	// 							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 							DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 								ID: to.Ptr("aaaaaaaaaaaa"),
	// 							},
	// 						},
	// 						CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 						Name: to.Ptr("aaaaaaaaaaaaaaa"),
	// 						WriteAcceleratorEnabled: to.Ptr(true),
	// 						DiffDiskSettings: &armcompute.DiffDiskSettings{
	// 							Option: to.Ptr(armcompute.DiffDiskOptionsLocal),
	// 							Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
	// 						},
	// 						DiskSizeGB: to.Ptr[int32](6),
	// 						OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 						Image: &armcompute.VirtualHardDisk{
	// 							URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd"),
	// 						},
	// 						VhdContainers: []*string{
	// 							to.Ptr("aa"),
	// 						},
	// 					},
	// 					DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
	// 						{
	// 							Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 							Lun: to.Ptr[int32](26),
	// 							Caching: to.Ptr(armcompute.CachingTypesNone),
	// 							WriteAcceleratorEnabled: to.Ptr(true),
	// 							CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
	// 							DiskSizeGB: to.Ptr[int32](1023),
	// 							ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
	// 								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 								DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 									ID: to.Ptr("aaaaaaaaaaaa"),
	// 								},
	// 							},
	// 							DiskIOPSReadWrite: to.Ptr[int64](28),
	// 							DiskMBpsReadWrite: to.Ptr[int64](15),
	// 						},
	// 					},
	// 				},
	// 				OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
	// 					ComputerNamePrefix: to.Ptr("{vmss-name}"),
	// 					AdminUsername: to.Ptr("{your-username}"),
	// 					CustomData: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 					WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 						ProvisionVMAgent: to.Ptr(true),
	// 						EnableAutomaticUpdates: to.Ptr(true),
	// 						TimeZone: to.Ptr("aaaaaaaaaaaaaaaa"),
	// 						AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
	// 							{
	// 								PassName: to.Ptr("OobeSystem"),
	// 								ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
	// 								SettingName: to.Ptr(armcompute.SettingNamesAutoLogon),
	// 								Content: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 							},
	// 						},
	// 						PatchSettings: &armcompute.PatchSettings{
	// 							PatchMode: to.Ptr(armcompute.WindowsVMGuestPatchModeAutomaticByPlatform),
	// 							EnableHotpatching: to.Ptr(true),
	// 							AssessmentMode: to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
	// 							AutomaticByPlatformSettings: &armcompute.WindowsVMGuestPatchAutomaticByPlatformSettings{
	// 								RebootSetting: to.Ptr(armcompute.WindowsVMGuestPatchAutomaticByPlatformRebootSettingNever),
	// 							},
	// 						},
	// 						WinRM: &armcompute.WinRMConfiguration{
	// 							Listeners: []*armcompute.WinRMListener{
	// 								{
	// 									Protocol: to.Ptr(armcompute.ProtocolTypesHTTP),
	// 									CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 					LinuxConfiguration: &armcompute.LinuxConfiguration{
	// 						DisablePasswordAuthentication: to.Ptr(true),
	// 						SSH: &armcompute.SSHConfiguration{
	// 							PublicKeys: []*armcompute.SSHPublicKey{
	// 								{
	// 									Path: to.Ptr("/home/{your-username}/.ssh/authorized_keys"),
	// 									KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCeClRAk2ipUs/l5voIsDC5q9RI+YSRd1Bvd/O+axgY4WiBzG+4FwJWZm/mLLe5DoOdHQwmU2FrKXZSW4w2sYE70KeWnrFViCOX5MTVvJgPE8ClugNl8RWth/tU849DvM9sT7vFgfVSHcAS2yDRyDlueii+8nF2ym8XWAPltFVCyLHRsyBp5YPqK8JFYIa1eybKsY3hEAxRCA+/7bq8et+Gj3coOsuRmrehav7rE6N12Pb80I6ofa6SM5XNYq4Xk0iYNx7R3kdz0Jj9XgZYWjAHjJmT0gTRoOnt6upOuxK7xI/ykWrllgpXrCPu3Ymz+c+ujaqcxDopnAl2lmf69/J1"),
	// 								},
	// 							},
	// 						},
	// 						ProvisionVMAgent: to.Ptr(true),
	// 						PatchSettings: &armcompute.LinuxPatchSettings{
	// 							PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
	// 							AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
	// 						},
	// 					},
	// 					Secrets: []*armcompute.VaultSecretGroup{
	// 						{
	// 							SourceVault: &armcompute.SubResource{
	// 								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 							},
	// 							VaultCertificates: []*armcompute.VaultCertificate{
	// 								{
	// 									CertificateURL: to.Ptr("aaaaaaa"),
	// 									CertificateStore: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 								},
	// 							},
	// 						},
	// 					},
	// 				},
	// 				NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
	// 					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
	// 						{
	// 							Name: to.Ptr("aaaaaaaa"),
	// 							Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 								Primary: to.Ptr(true),
	// 								EnableIPForwarding: to.Ptr(true),
	// 								IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
	// 									{
	// 										Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 										Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
	// 											Subnet: &armcompute.APIEntityReference{
	// 												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123"),
	// 											},
	// 											Primary: to.Ptr(true),
	// 											PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
	// 												Name: to.Ptr("a"),
	// 												Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
	// 													IdleTimeoutInMinutes: to.Ptr[int32](3),
	// 													DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
	// 														DomainNameLabel: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 													},
	// 													IPTags: []*armcompute.VirtualMachineScaleSetIPTag{
	// 														{
	// 															IPTagType: to.Ptr("aaaaaaa"),
	// 															Tag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 														},
	// 													},
	// 													PublicIPPrefix: &armcompute.SubResource{
	// 														ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 													},
	// 													PublicIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
	// 													DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
	// 												},
	// 												SKU: &armcompute.PublicIPAddressSKU{
	// 													Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
	// 													Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
	// 												},
	// 											},
	// 											PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
	// 											ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
	// 												{
	// 													ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
	// 												},
	// 											},
	// 											ApplicationSecurityGroups: []*armcompute.SubResource{
	// 												{
	// 													ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
	// 												},
	// 											},
	// 											LoadBalancerBackendAddressPools: []*armcompute.SubResource{
	// 												{
	// 													ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
	// 												},
	// 											},
	// 											LoadBalancerInboundNatPools: []*armcompute.SubResource{
	// 												{
	// 													ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
	// 												},
	// 											},
	// 										},
	// 									},
	// 								},
	// 								EnableAcceleratedNetworking: to.Ptr(true),
	// 								EnableFpga: to.Ptr(true),
	// 								NetworkSecurityGroup: &armcompute.SubResource{
	// 									ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
	// 								},
	// 								DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
	// 									DNSServers: []*string{
	// 										to.Ptr("aaaaaaaaaaaa"),
	// 									},
	// 								},
	// 								DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
	// 							},
	// 						},
	// 					},
	// 					HealthProbe: &armcompute.APIEntityReference{
	// 						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123"),
	// 					},
	// 					NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
	// 				},
	// 				SecurityProfile: &armcompute.SecurityProfile{
	// 					UefiSettings: &armcompute.UefiSettings{
	// 						SecureBootEnabled: to.Ptr(true),
	// 						VTpmEnabled: to.Ptr(true),
	// 					},
	// 					EncryptionAtHost: to.Ptr(true),
	// 					SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
	// 				},
	// 				DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 					BootDiagnostics: &armcompute.BootDiagnostics{
	// 						Enabled: to.Ptr(true),
	// 						StorageURI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net"),
	// 					},
	// 				},
	// 				ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
	// 					Extensions: []*armcompute.VirtualMachineScaleSetExtension{
	// 						{
	// 							Name: to.Ptr("{extension-name}"),
	// 							Type: to.Ptr("aaaaa"),
	// 							Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
	// 								ForceUpdateTag: to.Ptr("aaaaaaaaa"),
	// 								Publisher: to.Ptr("{extension-Publisher}"),
	// 								Type: to.Ptr("{extension-Type}"),
	// 								TypeHandlerVersion: to.Ptr("{handler-version}"),
	// 								AutoUpgradeMinorVersion: to.Ptr(true),
	// 								EnableAutomaticUpgrade: to.Ptr(true),
	// 								Settings: map[string]any{
	// 								},
	// 								ProtectedSettings: map[string]any{
	// 								},
	// 								ProvisioningState: to.Ptr("aaaaaaaaaaaaaa"),
	// 								ProvisionAfterExtensions: []*string{
	// 									to.Ptr("aa"),
	// 								},
	// 								SuppressFailures: to.Ptr(true),
	// 							},
	// 							ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
	// 						},
	// 					},
	// 					ExtensionsTimeBudget: to.Ptr("PT1H20M"),
	// 				},
	// 				LicenseType: to.Ptr("aaaaaaaaaaaa"),
	// 				Priority: to.Ptr(armcompute.VirtualMachinePriorityTypesRegular),
	// 				EvictionPolicy: to.Ptr(armcompute.VirtualMachineEvictionPolicyTypesDeallocate),
	// 				BillingProfile: &armcompute.BillingProfile{
	// 					MaxPrice: to.Ptr[float64](-1),
	// 				},
	// 				ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
	// 					TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
	// 						NotBeforeTimeout: to.Ptr("PT10M"),
	// 						Enable: to.Ptr(true),
	// 					},
	// 				},
	// 				UserData: to.Ptr("aaaaaaaaaaaaa"),
	// 				CapacityReservation: &armcompute.CapacityReservationProfile{
	// 					CapacityReservationGroup: &armcompute.SubResource{
	// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 					},
	// 				},
	// 				ApplicationProfile: &armcompute.ApplicationProfile{
	// 					GalleryApplications: []*armcompute.VMGalleryApplication{
	// 						{
	// 							Tags: to.Ptr("aaaaaaaaaaa"),
	// 							Order: to.Ptr[int32](29),
	// 							PackageReferenceID: to.Ptr("aaaaaaaaaa"),
	// 							ConfigurationReference: to.Ptr("aaaaa"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			UpgradePolicy: &armcompute.UpgradePolicy{
	// 				Mode: to.Ptr(armcompute.UpgradeModeManual),
	// 				RollingUpgradePolicy: &armcompute.RollingUpgradePolicy{
	// 					MaxBatchInstancePercent: to.Ptr[int32](49),
	// 					MaxUnhealthyInstancePercent: to.Ptr[int32](81),
	// 					MaxUnhealthyUpgradedInstancePercent: to.Ptr[int32](98),
	// 					PauseTimeBetweenBatches: to.Ptr("aaaaaaaaaaaaaaa"),
	// 					EnableCrossZoneUpgrade: to.Ptr(true),
	// 					PrioritizeUnhealthyInstances: to.Ptr(true),
	// 					MaxSurge: to.Ptr(true),
	// 				},
	// 				AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
	// 					EnableAutomaticOSUpgrade: to.Ptr(true),
	// 					DisableAutomaticRollback: to.Ptr(true),
	// 					OSRollingUpgradeDeferral: to.Ptr(true),
	// 				},
	// 			},
	// 			AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
	// 				Enabled: to.Ptr(true),
	// 				GracePeriod: to.Ptr("PT30M"),
	// 			},
	// 			ProvisioningState: to.Ptr("succeeded"),
	// 			DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(true),
	// 			UniqueID: to.Ptr("aaaaaaaa"),
	// 			SinglePlacementGroup: to.Ptr(true),
	// 			ZoneBalance: to.Ptr(true),
	// 			PlatformFaultDomainCount: to.Ptr[int32](1),
	// 			ProximityPlacementGroup: &armcompute.SubResource{
	// 				ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
	// 			},
	// 			HostGroup: &armcompute.SubResource{
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 			},
	// 			AdditionalCapabilities: &armcompute.AdditionalCapabilities{
	// 				UltraSSDEnabled: to.Ptr(true),
	// 				HibernationEnabled: to.Ptr(true),
	// 			},
	// 			ScaleInPolicy: &armcompute.ScaleInPolicy{
	// 				Rules: []*armcompute.VirtualMachineScaleSetScaleInRules{
	// 					to.Ptr(armcompute.VirtualMachineScaleSetScaleInRulesOldestVM),
	// 				},
	// 				ForceDeletion: to.Ptr(true),
	// 			},
	// 			OrchestrationMode: to.Ptr(armcompute.OrchestrationModeUniform),
	// 			SpotRestorePolicy: &armcompute.SpotRestorePolicy{
	// 				Enabled: to.Ptr(true),
	// 				RestoreTimeout: to.Ptr("aaaaaaaaaa"),
	// 			},
	// 		},
	// 		Plan: &armcompute.Plan{
	// 			Name: to.Ptr("aaaaaaaaaa"),
	// 			Publisher: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
	// 			Product: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 			PromotionCode: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 		},
	// 		Identity: &armcompute.VirtualMachineScaleSetIdentity{
	// 			PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
	// 			TenantID: to.Ptr("aaaaaaaaaaaaaaaa"),
	// 			Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssigned),
	// 			UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
	// 				"key3951": &armcompute.UserAssignedIdentitiesValue{
	// 					PrincipalID: to.Ptr("aaaa"),
	// 					ClientID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 				},
	// 			},
	// 		},
	// 		Zones: []*string{
	// 			to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 		},
	// 		ExtendedLocation: &armcompute.ExtendedLocation{
	// 			Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 			Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 		},
	// 		ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}"),
	// 		Name: to.Ptr("{virtualMachineScaleSetName}"),
	// 		Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
	// 		Tags: map[string]*string{
	// 			"key8425": to.Ptr("aaa"),
	// 		},
	// 	},
	// }
}
