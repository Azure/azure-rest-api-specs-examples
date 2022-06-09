```go
package armcompute_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Update_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.VirtualMachineScaleSetUpdate{
			Tags: map[string]*string{
				"key246": to.Ptr("aaaaaaaaaaaaaaaaaaaaaaaa"),
			},
			Identity: &armcompute.VirtualMachineScaleSetIdentity{
				Type: to.Ptr(armcompute.ResourceIdentityTypeSystemAssigned),
				UserAssignedIdentities: map[string]*armcompute.VirtualMachineScaleSetIdentityUserAssignedIdentitiesValue{
					"key3951": {},
				},
			},
			Plan: &armcompute.Plan{
				Name:          to.Ptr("<name>"),
				Product:       to.Ptr("<product>"),
				PromotionCode: to.Ptr("<promotion-code>"),
				Publisher:     to.Ptr("<publisher>"),
			},
			Properties: &armcompute.VirtualMachineScaleSetUpdateProperties{
				AdditionalCapabilities: &armcompute.AdditionalCapabilities{
					HibernationEnabled: to.Ptr(true),
					UltraSSDEnabled:    to.Ptr(true),
				},
				AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
					Enabled:     to.Ptr(true),
					GracePeriod: to.Ptr("<grace-period>"),
				},
				DoNotRunExtensionsOnOverprovisionedVMs: to.Ptr(true),
				Overprovision:                          to.Ptr(true),
				ProximityPlacementGroup: &armcompute.SubResource{
					ID: to.Ptr("<id>"),
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
					},
					Mode: to.Ptr(armcompute.UpgradeModeManual),
					RollingUpgradePolicy: &armcompute.RollingUpgradePolicy{
						EnableCrossZoneUpgrade:              to.Ptr(true),
						MaxBatchInstancePercent:             to.Ptr[int32](49),
						MaxUnhealthyInstancePercent:         to.Ptr[int32](81),
						MaxUnhealthyUpgradedInstancePercent: to.Ptr[int32](98),
						PauseTimeBetweenBatches:             to.Ptr("<pause-time-between-batches>"),
						PrioritizeUnhealthyInstances:        to.Ptr(true),
					},
				},
				VirtualMachineProfile: &armcompute.VirtualMachineScaleSetUpdateVMProfile{
					BillingProfile: &armcompute.BillingProfile{
						MaxPrice: to.Ptr[float64](-1),
					},
					DiagnosticsProfile: &armcompute.DiagnosticsProfile{
						BootDiagnostics: &armcompute.BootDiagnostics{
							Enabled:    to.Ptr(true),
							StorageURI: to.Ptr("<storage-uri>"),
						},
					},
					ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
						ExtensionsTimeBudget: to.Ptr("<extensions-time-budget>"),
						Extensions: []*armcompute.VirtualMachineScaleSetExtension{
							{
								Name: to.Ptr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
									Type:                    to.Ptr("<type>"),
									AutoUpgradeMinorVersion: to.Ptr(true),
									EnableAutomaticUpgrade:  to.Ptr(true),
									ForceUpdateTag:          to.Ptr("<force-update-tag>"),
									ProtectedSettings:       map[string]interface{}{},
									ProvisionAfterExtensions: []*string{
										to.Ptr("aa")},
									Publisher:          to.Ptr("<publisher>"),
									Settings:           map[string]interface{}{},
									SuppressFailures:   to.Ptr(true),
									TypeHandlerVersion: to.Ptr("<type-handler-version>"),
								},
							}},
					},
					LicenseType: to.Ptr("<license-type>"),
					NetworkProfile: &armcompute.VirtualMachineScaleSetUpdateNetworkProfile{
						HealthProbe: &armcompute.APIEntityReference{
							ID: to.Ptr("<id>"),
						},
						NetworkAPIVersion: to.Ptr(armcompute.NetworkAPIVersionTwoThousandTwenty1101),
						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetUpdateNetworkConfiguration{
							{
								ID:   to.Ptr("<id>"),
								Name: to.Ptr("<name>"),
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
											ID:   to.Ptr("<id>"),
											Name: to.Ptr("<name>"),
											Properties: &armcompute.VirtualMachineScaleSetUpdateIPConfigurationProperties{
												ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
													{
														ID: to.Ptr("<id>"),
													}},
												ApplicationSecurityGroups: []*armcompute.SubResource{
													{
														ID: to.Ptr("<id>"),
													}},
												LoadBalancerBackendAddressPools: []*armcompute.SubResource{
													{
														ID: to.Ptr("<id>"),
													}},
												LoadBalancerInboundNatPools: []*armcompute.SubResource{
													{
														ID: to.Ptr("<id>"),
													}},
												Primary:                 to.Ptr(true),
												PrivateIPAddressVersion: to.Ptr(armcompute.IPVersionIPv4),
												PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetUpdatePublicIPAddressConfiguration{
													Name: to.Ptr("<name>"),
													Properties: &armcompute.VirtualMachineScaleSetUpdatePublicIPAddressConfigurationProperties{
														DeleteOption: to.Ptr(armcompute.DeleteOptionsDelete),
														DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
															DomainNameLabel: to.Ptr("<domain-name-label>"),
														},
														IdleTimeoutInMinutes: to.Ptr[int32](3),
													},
												},
												Subnet: &armcompute.APIEntityReference{
													ID: to.Ptr("<id>"),
												},
											},
										}},
									NetworkSecurityGroup: &armcompute.SubResource{
										ID: to.Ptr("<id>"),
									},
									Primary: to.Ptr(true),
								},
							}},
					},
					OSProfile: &armcompute.VirtualMachineScaleSetUpdateOSProfile{
						CustomData: to.Ptr("<custom-data>"),
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
										Path:    to.Ptr("<path>"),
										KeyData: to.Ptr("<key-data>"),
									}},
							},
						},
						Secrets: []*armcompute.VaultSecretGroup{
							{
								SourceVault: &armcompute.SubResource{
									ID: to.Ptr("<id>"),
								},
								VaultCertificates: []*armcompute.VaultCertificate{
									{
										CertificateStore: to.Ptr("<certificate-store>"),
										CertificateURL:   to.Ptr("<certificate-url>"),
									}},
							}},
						WindowsConfiguration: &armcompute.WindowsConfiguration{
							AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
								{
									ComponentName: to.Ptr("<component-name>"),
									Content:       to.Ptr("<content>"),
									PassName:      to.Ptr("<pass-name>"),
									SettingName:   to.Ptr(armcompute.SettingNamesAutoLogon),
								}},
							EnableAutomaticUpdates: to.Ptr(true),
							PatchSettings: &armcompute.PatchSettings{
								AssessmentMode:    to.Ptr(armcompute.WindowsPatchAssessmentModeImageDefault),
								EnableHotpatching: to.Ptr(true),
								PatchMode:         to.Ptr(armcompute.WindowsVMGuestPatchModeAutomaticByOS),
							},
							ProvisionVMAgent: to.Ptr(true),
							TimeZone:         to.Ptr("<time-zone>"),
							WinRM: &armcompute.WinRMConfiguration{
								Listeners: []*armcompute.WinRMListener{
									{
										CertificateURL: to.Ptr("<certificate-url>"),
										Protocol:       to.Ptr(armcompute.ProtocolTypesHTTP),
									}},
							},
						},
					},
					ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
						TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
							Enable:           to.Ptr(true),
							NotBeforeTimeout: to.Ptr("<not-before-timeout>"),
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
								Name:              to.Ptr("<name>"),
								Caching:           to.Ptr(armcompute.CachingTypesNone),
								CreateOption:      to.Ptr(armcompute.DiskCreateOptionTypesEmpty),
								DiskIOPSReadWrite: to.Ptr[int64](28),
								DiskMBpsReadWrite: to.Ptr[int64](15),
								DiskSizeGB:        to.Ptr[int32](1023),
								Lun:               to.Ptr[int32](26),
								ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
									DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
										ID: to.Ptr("<id>"),
									},
									StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
								},
								WriteAcceleratorEnabled: to.Ptr(true),
							}},
						ImageReference: &armcompute.ImageReference{
							ID:                   to.Ptr("<id>"),
							Offer:                to.Ptr("<offer>"),
							Publisher:            to.Ptr("<publisher>"),
							SharedGalleryImageID: to.Ptr("<shared-gallery-image-id>"),
							SKU:                  to.Ptr("<sku>"),
							Version:              to.Ptr("<version>"),
						},
						OSDisk: &armcompute.VirtualMachineScaleSetUpdateOSDisk{
							Caching:    to.Ptr(armcompute.CachingTypesReadWrite),
							DiskSizeGB: to.Ptr[int32](6),
							Image: &armcompute.VirtualHardDisk{
								URI: to.Ptr("<uri>"),
							},
							ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
								DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
									ID: to.Ptr("<id>"),
								},
								StorageAccountType: to.Ptr(armcompute.StorageAccountTypesStandardLRS),
							},
							VhdContainers: []*string{
								to.Ptr("aa")},
							WriteAcceleratorEnabled: to.Ptr(true),
						},
					},
					UserData: to.Ptr("<user-data>"),
				},
			},
			SKU: &armcompute.SKU{
				Name:     to.Ptr("<name>"),
				Capacity: to.Ptr[int64](7),
				Tier:     to.Ptr("<tier>"),
			},
		},
		&armcompute.VirtualMachineScaleSetsClientBeginUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.7.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.
