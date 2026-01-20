package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Update_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginUpdate_virtualMachineScaleSetUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVirtualMachineScaleSetsClient().BeginUpdate(ctx, "rgcompute", "aaaaaaaaaaaaa", armcompute.VirtualMachineScaleSetUpdate{
		Tags: map[string]*string{
			"key246": to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
		},
		Identity: &armcompute.VirtualMachineScaleSetIdentity{
			Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssigned),
			UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
				"key3951": {},
			},
		},
		Plan: &armcompute.Plan{
			Name:          to.Ptr("windows2016"),
			Product:       to.Ptr("windows-data-science-vm"),
			PromotionCode: to.Ptr("aaaaaaaaaa"),
			Publisher:     to.Ptr("microsoft-ads"),
		},
		Properties: &armcompute.VirtualMachineScaleSetUpdateProperties{
			AdditionalCapabilities: &armcompute.AdditionalCapabilities{
				HibernationEnabled: to.Ptr(true),
				UltraSSDEnabled:    to.Ptr(true),
			},
			AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
				Enabled:     to.Ptr(true),
				GracePeriod: to.Ptr("PT30M"),
			},
			DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(true),
			Overprovision:                          to.Ptr(true),
			ProximityPlacementGroup: &armcompute.SubResource{
				ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
			},
			ScaleInPolicy: &armcompute.ScaleInPolicy{
				ForceDeletion: to.Ptr(true),
				Rules: []*armcompute.VirtualMachineScaleSetScaleInRules{
					to.Ptr(armcompute.VirtualMachineScaleSetScaleInRulesOldestVM)},
			},
			SinglePlacementGroup: to.Ptr(true),
			UpgradePolicy: &armcompute.UpgradePolicy{
				AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
					DisableAutomaticRollback: to.Ptr(true),
					EnableAutomaticOSUpgrade: to.Ptr(true),
					OSRollingUpgradeDeferral: to.Ptr(true),
				},
				Mode: to.Ptr(armcompute.UpgradeModeManual),
				RollingUpgradePolicy: &armcompute.RollingUpgradePolicy{
					EnableCrossZoneUpgrade:                to.Ptr(true),
					MaxBatchInstancePercent:               to.Ptr[int32](49),
					MaxSurge:                              to.Ptr(true),
					MaxUnhealthyInstancePercent:           to.Ptr[int32](81),
					MaxUnhealthyUpgradedInstancePercent:   to.Ptr[int32](98),
					PauseTimeBetweenBatches:               to.Ptr("aaaaaaaaaaaaaaa"),
					PrioritizeUnhealthyInstances:          to.Ptr(true),
					RollbackFailedInstancesOnPolicyBreach: to.Ptr(true),
				},
			},
			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetUpdateVMProfile{
				BillingProfile: &armcompute.BillingProfile{
					MaxPrice: to.Ptr[float64](-1),
				},
				DiagnosticsProfile: &armcompute.DiagnosticsProfile{
					BootDiagnostics: &armcompute.BootDiagnostics{
						Enabled:    to.Ptr(true),
						StorageURI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net"),
					},
				},
				ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
					ExtensionsTimeBudget: to.Ptr("PT1H20M"),
					Extensions: []*armcompute.VirtualMachineScaleSetExtension{
						{
							Name: to.Ptr("{extension-name}"),
							Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
								Type:                    to.Ptr("{extension-Type}"),
								AutoUpgradeMinorVersion: to.Ptr(true),
								EnableAutomaticUpgrade:  to.Ptr(true),
								ForceUpdateTag:          to.Ptr("aaaaaaaaa"),
								ProtectedSettings:       map[string]any{},
								ProvisionAfterExtensions: []*string{
									to.Ptr("aa")},
								Publisher:          to.Ptr("{extension-Publisher}"),
								Settings:           map[string]any{},
								SuppressFailures:   to.Ptr(true),
								TypeHandlerVersion: to.Ptr("{handler-version}"),
							},
						}},
				},
				LicenseType: to.Ptr("aaaaaaaaaaaa"),
				NetworkProfile: &armcompute.VirtualMachineScaleSetUpdateNetworkProfile{
					HealthProbe: &armcompute.APIEntityReference{
						ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123"),
					},
					NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
					NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetUpdateNetworkConfiguration{
						{
							Name: to.Ptr("aaaaaaaa"),
							Properties: &armcompute.VirtualMachineScaleSetUpdateNetworkConfigurationProperties{
								DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
								DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
									DNSServers: []*string{},
								},
								EnableAcceleratedNetworking: to.Ptr(true),
								EnableFpga:                  to.Ptr(true),
								EnableIPForwarding:          to.Ptr(true),
								IPConfigurations: []*armcompute.VirtualMachineScaleSetUpdateIPConfiguration{
									{
										Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
										Properties: &armcompute.VirtualMachineScaleSetUpdateIPConfigurationProperties{
											ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
												}},
											ApplicationSecurityGroups: []*armcompute.SubResource{
												{
													ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
												}},
											LoadBalancerBackendAddressPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
												}},
											LoadBalancerInboundNatPools: []*armcompute.SubResource{
												{
													ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
												}},
											Primary:                 to.Ptr(true),
											PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
											PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetUpdatePublicIPAddressConfiguration{
												Name: to.Ptr("a"),
												Properties: &armcompute.VirtualMachineScaleSetUpdatePublicIPAddressConfigurationProperties{
													DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
													DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
														DomainNameLabel: to.Ptr("aaaaaaaaaaaaaaaaaa"),
													},
													IdleTimeoutInMinutes: to.Ptr[int32](3),
												},
											},
											Subnet: &armcompute.APIEntityReference{
												ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123"),
											},
										},
									}},
								NetworkSecurityGroup: &armcompute.SubResource{
									ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
								},
								Primary: to.Ptr(true),
							},
						}},
				},
				OSProfile: &armcompute.VirtualMachineScaleSetUpdateOSProfile{
					CustomData: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
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
									Path:    to.Ptr("/home/{your-username}/.ssh/authorized_keys"),
									KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCeClRAk2ipUs/l5voIsDC5q9RI+YSRd1Bvd/O+axgY4WiBzG+4FwJWZm/mLLe5DoOdHQwmU2FrKXZSW4w2sYE70KeWnrFViCOX5MTVvJgPE8ClugNl8RWth/tU849DvM9sT7vFgfVSHcAS2yDRyDlueii+8nF2ym8XWAPltFVCyLHRsyBp5YPqK8JFYIa1eybKsY3hEAxRCA+/7bq8et+Gj3coOsuRmrehav7rE6N12Pb80I6ofa6SM5XNYq4Xk0iYNx7R3kdz0Jj9XgZYWjAHjJmT0gTRoOnt6upOuxK7xI/ykWrllgpXrCPu3Ymz+c+ujaqcxDopnAl2lmf69/J1"),
								}},
						},
					},
					Secrets: []*armcompute.VaultSecretGroup{
						{
							SourceVault: &armcompute.SubResource{
								ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
							},
							VaultCertificates: []*armcompute.VaultCertificate{
								{
									CertificateStore: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
									CertificateURL:   to.Ptr("aaaaaaa"),
								}},
						}},
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
							AssessmentMode: to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
							AutomaticByPlatformSettings: &armcompute.WindowsVMGuestPatchAutomaticByPlatformSettings{
								RebootSetting: to.Ptr(armcompute.WindowsVMGuestPatchAutomaticByPlatformRebootSettingNever),
							},
							EnableHotpatching: to.Ptr(true),
							PatchMode:         to.Ptr(armcompute.WindowsVMGuestPatchModeAutomaticByPlatform),
						},
						ProvisionVMAgent: to.Ptr(true),
						TimeZone:         to.Ptr("aaaaaaaaaaaaaaaa"),
						WinRM: &armcompute.WinRMConfiguration{
							Listeners: []*armcompute.WinRMListener{
								{
									CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
									Protocol:       to.Ptr(armcompute.ProtocolTypesHTTP),
								}},
						},
					},
				},
				ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
					TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
						Enable:           to.Ptr(true),
						NotBeforeTimeout: to.Ptr("PT10M"),
					},
				},
				SecurityProfile: &armcompute.SecurityProfile{
					EncryptionAtHost: to.Ptr(true),
					SecurityType:     to.Ptr(armcompute.SecurityTypesTrustedLaunch),
					UefiSettings: &armcompute.UefiSettings{
						SecureBootEnabled: to.Ptr(true),
						VTpmEnabled:       to.Ptr(true),
					},
				},
				StorageProfile: &armcompute.VirtualMachineScaleSetUpdateStorageProfile{
					DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
						{
							Name:              to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
							Caching:           to.Ptr(armcompute.CachingTypesNone),
							CreateOption:      to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
							DiskIOPSReadWrite: to.Ptr[int64](28),
							DiskMBpsReadWrite: to.Ptr[int64](15),
							DiskSizeGB:        to.Ptr[int32](1023),
							Lun:               to.Ptr[int32](26),
							ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
								DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
									ID: to.Ptr("aaaaaaaaaaaa"),
								},
								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
							},
							WriteAcceleratorEnabled: to.Ptr(true),
						}},
					ImageReference: &armcompute.ImageReference{
						ID:                   to.Ptr("aaaaaaaaaaaaaaaaaaa"),
						Offer:                to.Ptr("WindowsServer"),
						Publisher:            to.Ptr("MicrosoftWindowsServer"),
						SharedGalleryImageID: to.Ptr("aaaaaa"),
						SKU:                  to.Ptr("2016-Datacenter"),
						Version:              to.Ptr("latest"),
					},
					OSDisk: &armcompute.VirtualMachineScaleSetUpdateOSDisk{
						Caching: to.Ptr(armcompute.CachingTypesReadWrite),
						DiffDiskSettings: &armcompute.DiffDiskSettings{
							Option:    to.Ptr(armcompute.DiffDiskOptionsLocal),
							Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
						},
						DiskSizeGB: to.Ptr[int32](6),
						Image: &armcompute.VirtualHardDisk{
							URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd"),
						},
						ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
							DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
								ID: to.Ptr("aaaaaaaaaaaa"),
							},
							StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
						},
						VhdContainers: []*string{
							to.Ptr("aa")},
						WriteAcceleratorEnabled: to.Ptr(true),
					},
				},
				UserData: to.Ptr("aaaaaaaaaaaaa"),
			},
		},
		SKU: &armcompute.SKU{
			Name:     to.Ptr("DSv3-Type1"),
			Capacity: to.Ptr[int64](7),
			Tier:     to.Ptr("aaa"),
		},
		Zones: []*string{
			to.Ptr("1"),
			to.Ptr("2"),
			to.Ptr("3")},
	}, &armcompute.VirtualMachineScaleSetsClientBeginUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualMachineScaleSet = armcompute.VirtualMachineScaleSet{
	// 	Name: to.Ptr("{virtualMachineScaleSetName}"),
	// 	Type: to.Ptr("Microsoft.Compute/virtualMachineScaleSets"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}"),
	// 	Location: to.Ptr("westus"),
	// 	Tags: map[string]*string{
	// 		"key8425": to.Ptr("aaa"),
	// 	},
	// 	ExtendedLocation: &armcompute.ExtendedLocation{
	// 		Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaa"),
	// 		Type: to.Ptr(armcompute.ExtendedLocationTypesEdgeZone),
	// 	},
	// 	Identity: &armcompute.VirtualMachineScaleSetIdentity{
	// 		Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("aaaaaaaaaaaaaaa"),
	// 		TenantID: to.Ptr("aaaaaaaaaaaaaaaa"),
	// 		UserAssignedIdentities: map[string]*armcompute.UserAssignedIdentitiesValue{
	// 			"key3951": &armcompute.UserAssignedIdentitiesValue{
	// 				ClientID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 				PrincipalID: to.Ptr("aaaa"),
	// 			},
	// 		},
	// 	},
	// 	Plan: &armcompute.Plan{
	// 		Name: to.Ptr("aaaaaaaaaa"),
	// 		Product: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 		PromotionCode: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 		Publisher: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
	// 	},
	// 	Properties: &armcompute.VirtualMachineScaleSetProperties{
	// 		AdditionalCapabilities: &armcompute.AdditionalCapabilities{
	// 			HibernationEnabled: to.Ptr(true),
	// 			UltraSSDEnabled: to.Ptr(true),
	// 		},
	// 		AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
	// 			Enabled: to.Ptr(true),
	// 			GracePeriod: to.Ptr("PT30M"),
	// 		},
	// 		DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(true),
	// 		HostGroup: &armcompute.SubResource{
	// 			ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 		},
	// 		OrchestrationMode: to.Ptr(armcompute.OrchestrationModeUniform),
	// 		Overprovision: to.Ptr(true),
	// 		PlatformFaultDomainCount: to.Ptr[int32](1),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ProximityPlacementGroup: &armcompute.SubResource{
	// 			ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
	// 		},
	// 		ScaleInPolicy: &armcompute.ScaleInPolicy{
	// 			ForceDeletion: to.Ptr(true),
	// 			Rules: []*armcompute.VirtualMachineScaleSetScaleInRules{
	// 				to.Ptr(armcompute.VirtualMachineScaleSetScaleInRulesOldestVM)},
	// 			},
	// 			SinglePlacementGroup: to.Ptr(true),
	// 			SpotRestorePolicy: &armcompute.SpotRestorePolicy{
	// 				Enabled: to.Ptr(true),
	// 				RestoreTimeout: to.Ptr("aaaaaaaaaa"),
	// 			},
	// 			UniqueID: to.Ptr("aaaaaaaa"),
	// 			UpgradePolicy: &armcompute.UpgradePolicy{
	// 				AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
	// 					DisableAutomaticRollback: to.Ptr(true),
	// 					EnableAutomaticOSUpgrade: to.Ptr(true),
	// 					OSRollingUpgradeDeferral: to.Ptr(true),
	// 				},
	// 				Mode: to.Ptr(armcompute.UpgradeModeManual),
	// 				RollingUpgradePolicy: &armcompute.RollingUpgradePolicy{
	// 					EnableCrossZoneUpgrade: to.Ptr(true),
	// 					MaxBatchInstancePercent: to.Ptr[int32](49),
	// 					MaxSurge: to.Ptr(true),
	// 					MaxUnhealthyInstancePercent: to.Ptr[int32](81),
	// 					MaxUnhealthyUpgradedInstancePercent: to.Ptr[int32](98),
	// 					PauseTimeBetweenBatches: to.Ptr("aaaaaaaaaaaaaaa"),
	// 					PrioritizeUnhealthyInstances: to.Ptr(true),
	// 				},
	// 			},
	// 			VirtualMachineProfile: &armcompute.VirtualMachineScaleSetVMProfile{
	// 				ApplicationProfile: &armcompute.ApplicationProfile{
	// 					GalleryApplications: []*armcompute.VMGalleryApplication{
	// 						{
	// 							ConfigurationReference: to.Ptr("aaaaa"),
	// 							Order: to.Ptr[int32](29),
	// 							PackageReferenceID: to.Ptr("aaaaaaaaaa"),
	// 							Tags: to.Ptr("aaaaaaaaaaa"),
	// 					}},
	// 				},
	// 				BillingProfile: &armcompute.BillingProfile{
	// 					MaxPrice: to.Ptr[float64](-1),
	// 				},
	// 				CapacityReservation: &armcompute.CapacityReservationProfile{
	// 					CapacityReservationGroup: &armcompute.SubResource{
	// 						ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 					},
	// 				},
	// 				DiagnosticsProfile: &armcompute.DiagnosticsProfile{
	// 					BootDiagnostics: &armcompute.BootDiagnostics{
	// 						Enabled: to.Ptr(true),
	// 						StorageURI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net"),
	// 					},
	// 				},
	// 				EvictionPolicy: to.Ptr(armcompute.VirtualMachineEvictionPolicyTypesDeallocate),
	// 				ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
	// 					ExtensionsTimeBudget: to.Ptr("PT1H20M"),
	// 					Extensions: []*armcompute.VirtualMachineScaleSetExtension{
	// 						{
	// 							ID: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
	// 							Name: to.Ptr("{extension-name}"),
	// 							Type: to.Ptr("aaaaa"),
	// 							Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
	// 								Type: to.Ptr("{extension-Type}"),
	// 								AutoUpgradeMinorVersion: to.Ptr(true),
	// 								EnableAutomaticUpgrade: to.Ptr(true),
	// 								ForceUpdateTag: to.Ptr("aaaaaaaaa"),
	// 								ProtectedSettings: map[string]any{
	// 								},
	// 								ProvisionAfterExtensions: []*string{
	// 									to.Ptr("aa")},
	// 									ProvisioningState: to.Ptr("Succeeded"),
	// 									Publisher: to.Ptr("{extension-Publisher}"),
	// 									Settings: map[string]any{
	// 									},
	// 									SuppressFailures: to.Ptr(true),
	// 									TypeHandlerVersion: to.Ptr("{handler-version}"),
	// 								},
	// 						}},
	// 					},
	// 					LicenseType: to.Ptr("aaaaaaaaaaaa"),
	// 					NetworkProfile: &armcompute.VirtualMachineScaleSetNetworkProfile{
	// 						HealthProbe: &armcompute.APIEntityReference{
	// 							ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123"),
	// 						},
	// 						NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
	// 						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetNetworkConfiguration{
	// 							{
	// 								Name: to.Ptr("aaaaaaaa"),
	// 								Properties: &armcompute.VirtualMachineScaleSetNetworkConfigurationProperties{
	// 									DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
	// 									DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
	// 										DNSServers: []*string{
	// 											to.Ptr("aaaaaaaaaaaa")},
	// 										},
	// 										EnableAcceleratedNetworking: to.Ptr(true),
	// 										EnableFpga: to.Ptr(true),
	// 										EnableIPForwarding: to.Ptr(true),
	// 										IPConfigurations: []*armcompute.VirtualMachineScaleSetIPConfiguration{
	// 											{
	// 												Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 												Properties: &armcompute.VirtualMachineScaleSetIPConfigurationProperties{
	// 													ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
	// 														{
	// 															ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
	// 													}},
	// 													ApplicationSecurityGroups: []*armcompute.SubResource{
	// 														{
	// 															ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
	// 													}},
	// 													LoadBalancerBackendAddressPools: []*armcompute.SubResource{
	// 														{
	// 															ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
	// 													}},
	// 													LoadBalancerInboundNatPools: []*armcompute.SubResource{
	// 														{
	// 															ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
	// 													}},
	// 													Primary: to.Ptr(true),
	// 													PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
	// 													PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetPublicIPAddressConfiguration{
	// 														Name: to.Ptr("a"),
	// 														Properties: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationProperties{
	// 															DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
	// 															DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
	// 																DomainNameLabel: to.Ptr("aaaaaaaaaaaaaaaaaa"),
	// 															},
	// 															IdleTimeoutInMinutes: to.Ptr[int32](3),
	// 															IPTags: []*armcompute.VirtualMachineScaleSetIPTag{
	// 																{
	// 																	IPTagType: to.Ptr("aaaaaaa"),
	// 																	Tag: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 															}},
	// 															PublicIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
	// 															PublicIPPrefix: &armcompute.SubResource{
	// 																ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 															},
	// 														},
	// 														SKU: &armcompute.PublicIPAddressSKU{
	// 															Name: to.Ptr(armcompute.PublicIPAddressSKUNameBasic),
	// 															Tier: to.Ptr(armcompute.PublicIPAddressSKUTierRegional),
	// 														},
	// 													},
	// 													Subnet: &armcompute.APIEntityReference{
	// 														ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/disk123"),
	// 													},
	// 												},
	// 										}},
	// 										NetworkSecurityGroup: &armcompute.SubResource{
	// 											ID: to.Ptr("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"),
	// 										},
	// 										Primary: to.Ptr(true),
	// 									},
	// 							}},
	// 						},
	// 						OSProfile: &armcompute.VirtualMachineScaleSetOSProfile{
	// 							AdminUsername: to.Ptr("{your-username}"),
	// 							ComputerNamePrefix: to.Ptr("{vmss-name}"),
	// 							CustomData: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 							LinuxConfiguration: &armcompute.LinuxConfiguration{
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								PatchSettings: &armcompute.LinuxPatchSettings{
	// 									AssessmentMode: to.Ptr(armcompute.LinuxPatchAssessmentModeImageDefault),
	// 									PatchMode: to.Ptr(armcompute.LinuxVMGuestPatchModeImageDefault),
	// 								},
	// 								ProvisionVMAgent: to.Ptr(true),
	// 								SSH: &armcompute.SSHConfiguration{
	// 									PublicKeys: []*armcompute.SSHPublicKey{
	// 										{
	// 											Path: to.Ptr("/home/{your-username}/.ssh/authorized_keys"),
	// 											KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCeClRAk2ipUs/l5voIsDC5q9RI+YSRd1Bvd/O+axgY4WiBzG+4FwJWZm/mLLe5DoOdHQwmU2FrKXZSW4w2sYE70KeWnrFViCOX5MTVvJgPE8ClugNl8RWth/tU849DvM9sT7vFgfVSHcAS2yDRyDlueii+8nF2ym8XWAPltFVCyLHRsyBp5YPqK8JFYIa1eybKsY3hEAxRCA+/7bq8et+Gj3coOsuRmrehav7rE6N12Pb80I6ofa6SM5XNYq4Xk0iYNx7R3kdz0Jj9XgZYWjAHjJmT0gTRoOnt6upOuxK7xI/ykWrllgpXrCPu3Ymz+c+ujaqcxDopnAl2lmf69/J1"),
	// 									}},
	// 								},
	// 							},
	// 							Secrets: []*armcompute.VaultSecretGroup{
	// 								{
	// 									SourceVault: &armcompute.SubResource{
	// 										ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/availabilitySets/{availabilitySetName}"),
	// 									},
	// 									VaultCertificates: []*armcompute.VaultCertificate{
	// 										{
	// 											CertificateStore: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 											CertificateURL: to.Ptr("aaaaaaa"),
	// 									}},
	// 							}},
	// 							WindowsConfiguration: &armcompute.WindowsConfiguration{
	// 								AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
	// 									{
	// 										ComponentName: to.Ptr("Microsoft-Windows-Shell-Setup"),
	// 										Content: to.Ptr("aaaaaaaaaaaaaaaaaaaa"),
	// 										PassName: to.Ptr("OobeSystem"),
	// 										SettingName: to.Ptr(armcompute.SettingNamesAutoLogon),
	// 								}},
	// 								EnableAutomaticUpdates: to.Ptr(true),
	// 								PatchSettings: &armcompute.PatchSettings{
	// 									AssessmentMode: to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
	// 									AutomaticByPlatformSettings: &armcompute.WindowsVMGuestPatchAutomaticByPlatformSettings{
	// 										RebootSetting: to.Ptr(armcompute.WindowsVMGuestPatchAutomaticByPlatformRebootSettingNever),
	// 									},
	// 									EnableHotpatching: to.Ptr(true),
	// 									PatchMode: to.Ptr(armcompute.WindowsVMGuestPatchModeAutomaticByPlatform),
	// 								},
	// 								ProvisionVMAgent: to.Ptr(true),
	// 								TimeZone: to.Ptr("aaaaaaaaaaaaaaaa"),
	// 								WinRM: &armcompute.WinRMConfiguration{
	// 									Listeners: []*armcompute.WinRMListener{
	// 										{
	// 											CertificateURL: to.Ptr("aaaaaaaaaaaaaaaaaaaaaa"),
	// 											Protocol: to.Ptr(armcompute.ProtocolTypesHTTP),
	// 									}},
	// 								},
	// 							},
	// 						},
	// 						Priority: to.Ptr(armcompute.VirtualMachinePriorityTypesRegular),
	// 						ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
	// 							TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
	// 								Enable: to.Ptr(true),
	// 								NotBeforeTimeout: to.Ptr("PT10M"),
	// 							},
	// 						},
	// 						SecurityProfile: &armcompute.SecurityProfile{
	// 							EncryptionAtHost: to.Ptr(true),
	// 							SecurityType: to.Ptr(armcompute.SecurityTypesTrustedLaunch),
	// 							UefiSettings: &armcompute.UefiSettings{
	// 								SecureBootEnabled: to.Ptr(true),
	// 								VTpmEnabled: to.Ptr(true),
	// 							},
	// 						},
	// 						StorageProfile: &armcompute.VirtualMachineScaleSetStorageProfile{
	// 							DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
	// 								{
	// 									Name: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 									Caching: to.Ptr(armcompute.CachingTypesNone),
	// 									CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
	// 									DiskIOPSReadWrite: to.Ptr[int64](28),
	// 									DiskMBpsReadWrite: to.Ptr[int64](15),
	// 									DiskSizeGB: to.Ptr[int32](1023),
	// 									Lun: to.Ptr[int32](26),
	// 									ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
	// 										DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 											ID: to.Ptr("aaaaaaaaaaaa"),
	// 										},
	// 										StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 									},
	// 									WriteAcceleratorEnabled: to.Ptr(true),
	// 							}},
	// 							ImageReference: &armcompute.ImageReference{
	// 								ID: to.Ptr("aaaaaaaaaaaaaaaaaaa"),
	// 								ExactVersion: to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"),
	// 								Offer: to.Ptr("WindowsServer"),
	// 								Publisher: to.Ptr("MicrosoftWindowsServer"),
	// 								SharedGalleryImageID: to.Ptr("aaaaaa"),
	// 								SKU: to.Ptr("2016-Datacenter"),
	// 								Version: to.Ptr("latest"),
	// 							},
	// 							OSDisk: &armcompute.VirtualMachineScaleSetOSDisk{
	// 								Name: to.Ptr("aaaaaaaaaaaaaaa"),
	// 								Caching: to.Ptr(armcompute.CachingTypesReadWrite),
	// 								CreateOption: to.Ptr(armcompute.DiskCreateOptionTypesFromImage),
	// 								DiffDiskSettings: &armcompute.DiffDiskSettings{
	// 									Option: to.Ptr(armcompute.DiffDiskOptionsLocal),
	// 									Placement: to.Ptr(armcompute.DiffDiskPlacementCacheDisk),
	// 								},
	// 								DiskSizeGB: to.Ptr[int32](6),
	// 								Image: &armcompute.VirtualHardDisk{
	// 									URI: to.Ptr("http://{existing-storage-account-name}.blob.core.windows.net/{existing-container-name}/myDisk.vhd"),
	// 								},
	// 								ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
	// 									DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
	// 										ID: to.Ptr("aaaaaaaaaaaa"),
	// 									},
	// 									StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
	// 								},
	// 								OSType: to.Ptr(armcompute.OperatingSystemTypesWindows),
	// 								VhdContainers: []*string{
	// 									to.Ptr("aa")},
	// 									WriteAcceleratorEnabled: to.Ptr(true),
	// 								},
	// 							},
	// 							UserData: to.Ptr("aaaaaaaaaaaaa"),
	// 						},
	// 						ZoneBalance: to.Ptr(true),
	// 					},
	// 					SKU: &armcompute.SKU{
	// 						Name: to.Ptr("Standard_D1_v2"),
	// 						Capacity: to.Ptr[int64](3),
	// 						Tier: to.Ptr("Standard"),
	// 					},
	// 					Zones: []*string{
	// 						to.Ptr("aaaaaaaaaaaaaaaaaaaa")},
	// 					}
}
