Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcompute%2Farmcompute%2Fv0.5.0/sdk/resourcemanager/compute/armcompute/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSets_Update_MaximumSet_Gen.json
func ExampleVirtualMachineScaleSetsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcompute.NewVirtualMachineScaleSetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<vm-scale-set-name>",
		armcompute.VirtualMachineScaleSetUpdate{
			Tags: map[string]*string{
				"key246": to.StringPtr("aaaaaaaaaaaaaaaaaaaaaaaa"),
			},
			Identity: &armcompute.VirtualMachineScaleSetIdentity{
				Type: armcompute.ResourceIdentityTypeSystemAssigned.ToPtr(),
				UserAssignedIdentities: map[string]*armcompute.VirtualMachineScaleSetIdentityUserAssignedIdentitiesValue{
					"key3951": {},
				},
			},
			Plan: &armcompute.Plan{
				Name:          to.StringPtr("<name>"),
				Product:       to.StringPtr("<product>"),
				PromotionCode: to.StringPtr("<promotion-code>"),
				Publisher:     to.StringPtr("<publisher>"),
			},
			Properties: &armcompute.VirtualMachineScaleSetUpdateProperties{
				AdditionalCapabilities: &armcompute.AdditionalCapabilities{
					HibernationEnabled: to.BoolPtr(true),
					UltraSSDEnabled:    to.BoolPtr(true),
				},
				AutomaticRepairsPolicy: &armcompute.AutomaticRepairsPolicy{
					Enabled:     to.BoolPtr(true),
					GracePeriod: to.StringPtr("<grace-period>"),
				},
				DoNotRunExtensionsOnOverprovisionedVMs: to.BoolPtr(true),
				Overprovision:                          to.BoolPtr(true),
				ProximityPlacementGroup: &armcompute.SubResource{
					ID: to.StringPtr("<id>"),
				},
				ScaleInPolicy: &armcompute.ScaleInPolicy{
					ForceDeletion: to.BoolPtr(true),
					Rules: []*armcompute.VirtualMachineScaleSetScaleInRules{
						armcompute.VirtualMachineScaleSetScaleInRules("OldestVM").ToPtr()},
				},
				SinglePlacementGroup: to.BoolPtr(true),
				UpgradePolicy: &armcompute.UpgradePolicy{
					AutomaticOSUpgradePolicy: &armcompute.AutomaticOSUpgradePolicy{
						DisableAutomaticRollback: to.BoolPtr(true),
						EnableAutomaticOSUpgrade: to.BoolPtr(true),
					},
					Mode: armcompute.UpgradeModeManual.ToPtr(),
					RollingUpgradePolicy: &armcompute.RollingUpgradePolicy{
						EnableCrossZoneUpgrade:              to.BoolPtr(true),
						MaxBatchInstancePercent:             to.Int32Ptr(49),
						MaxUnhealthyInstancePercent:         to.Int32Ptr(81),
						MaxUnhealthyUpgradedInstancePercent: to.Int32Ptr(98),
						PauseTimeBetweenBatches:             to.StringPtr("<pause-time-between-batches>"),
						PrioritizeUnhealthyInstances:        to.BoolPtr(true),
					},
				},
				VirtualMachineProfile: &armcompute.VirtualMachineScaleSetUpdateVMProfile{
					BillingProfile: &armcompute.BillingProfile{
						MaxPrice: to.Float64Ptr(-1),
					},
					DiagnosticsProfile: &armcompute.DiagnosticsProfile{
						BootDiagnostics: &armcompute.BootDiagnostics{
							Enabled:    to.BoolPtr(true),
							StorageURI: to.StringPtr("<storage-uri>"),
						},
					},
					ExtensionProfile: &armcompute.VirtualMachineScaleSetExtensionProfile{
						ExtensionsTimeBudget: to.StringPtr("<extensions-time-budget>"),
						Extensions: []*armcompute.VirtualMachineScaleSetExtension{
							{
								Name: to.StringPtr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetExtensionProperties{
									Type:                    to.StringPtr("<type>"),
									AutoUpgradeMinorVersion: to.BoolPtr(true),
									EnableAutomaticUpgrade:  to.BoolPtr(true),
									ForceUpdateTag:          to.StringPtr("<force-update-tag>"),
									ProtectedSettings:       map[string]interface{}{},
									ProvisionAfterExtensions: []*string{
										to.StringPtr("aa")},
									Publisher:          to.StringPtr("<publisher>"),
									Settings:           map[string]interface{}{},
									SuppressFailures:   to.BoolPtr(true),
									TypeHandlerVersion: to.StringPtr("<type-handler-version>"),
								},
							}},
					},
					LicenseType: to.StringPtr("<license-type>"),
					NetworkProfile: &armcompute.VirtualMachineScaleSetUpdateNetworkProfile{
						HealthProbe: &armcompute.APIEntityReference{
							ID: to.StringPtr("<id>"),
						},
						NetworkAPIVersion: armcompute.NetworkAPIVersion("2020-11-01").ToPtr(),
						NetworkInterfaceConfigurations: []*armcompute.VirtualMachineScaleSetUpdateNetworkConfiguration{
							{
								ID:   to.StringPtr("<id>"),
								Name: to.StringPtr("<name>"),
								Properties: &armcompute.VirtualMachineScaleSetUpdateNetworkConfigurationProperties{
									DeleteOption: armcompute.DeleteOptions("Delete").ToPtr(),
									DNSSettings: &armcompute.VirtualMachineScaleSetNetworkConfigurationDNSSettings{
										DNSServers: []*string{},
									},
									EnableAcceleratedNetworking: to.BoolPtr(true),
									EnableFpga:                  to.BoolPtr(true),
									EnableIPForwarding:          to.BoolPtr(true),
									IPConfigurations: []*armcompute.VirtualMachineScaleSetUpdateIPConfiguration{
										{
											ID:   to.StringPtr("<id>"),
											Name: to.StringPtr("<name>"),
											Properties: &armcompute.VirtualMachineScaleSetUpdateIPConfigurationProperties{
												ApplicationGatewayBackendAddressPools: []*armcompute.SubResource{
													{
														ID: to.StringPtr("<id>"),
													}},
												ApplicationSecurityGroups: []*armcompute.SubResource{
													{
														ID: to.StringPtr("<id>"),
													}},
												LoadBalancerBackendAddressPools: []*armcompute.SubResource{
													{
														ID: to.StringPtr("<id>"),
													}},
												LoadBalancerInboundNatPools: []*armcompute.SubResource{
													{
														ID: to.StringPtr("<id>"),
													}},
												Primary:                 to.BoolPtr(true),
												PrivateIPAddressVersion: armcompute.IPVersion("IPv4").ToPtr(),
												PublicIPAddressConfiguration: &armcompute.VirtualMachineScaleSetUpdatePublicIPAddressConfiguration{
													Name: to.StringPtr("<name>"),
													Properties: &armcompute.VirtualMachineScaleSetUpdatePublicIPAddressConfigurationProperties{
														DeleteOption: armcompute.DeleteOptions("Delete").ToPtr(),
														DNSSettings: &armcompute.VirtualMachineScaleSetPublicIPAddressConfigurationDNSSettings{
															DomainNameLabel: to.StringPtr("<domain-name-label>"),
														},
														IdleTimeoutInMinutes: to.Int32Ptr(3),
													},
												},
												Subnet: &armcompute.APIEntityReference{
													ID: to.StringPtr("<id>"),
												},
											},
										}},
									NetworkSecurityGroup: &armcompute.SubResource{
										ID: to.StringPtr("<id>"),
									},
									Primary: to.BoolPtr(true),
								},
							}},
					},
					OSProfile: &armcompute.VirtualMachineScaleSetUpdateOSProfile{
						CustomData: to.StringPtr("<custom-data>"),
						LinuxConfiguration: &armcompute.LinuxConfiguration{
							DisablePasswordAuthentication: to.BoolPtr(true),
							PatchSettings: &armcompute.LinuxPatchSettings{
								AssessmentMode: armcompute.LinuxPatchAssessmentMode("ImageDefault").ToPtr(),
								PatchMode:      armcompute.LinuxVMGuestPatchMode("ImageDefault").ToPtr(),
							},
							ProvisionVMAgent: to.BoolPtr(true),
							SSH: &armcompute.SSHConfiguration{
								PublicKeys: []*armcompute.SSHPublicKey{
									{
										Path:    to.StringPtr("<path>"),
										KeyData: to.StringPtr("<key-data>"),
									}},
							},
						},
						Secrets: []*armcompute.VaultSecretGroup{
							{
								SourceVault: &armcompute.SubResource{
									ID: to.StringPtr("<id>"),
								},
								VaultCertificates: []*armcompute.VaultCertificate{
									{
										CertificateStore: to.StringPtr("<certificate-store>"),
										CertificateURL:   to.StringPtr("<certificate-url>"),
									}},
							}},
						WindowsConfiguration: &armcompute.WindowsConfiguration{
							AdditionalUnattendContent: []*armcompute.AdditionalUnattendContent{
								{
									ComponentName: to.StringPtr("<component-name>"),
									Content:       to.StringPtr("<content>"),
									PassName:      to.StringPtr("<pass-name>"),
									SettingName:   armcompute.SettingNamesAutoLogon.ToPtr(),
								}},
							EnableAutomaticUpdates: to.BoolPtr(true),
							PatchSettings: &armcompute.PatchSettings{
								AssessmentMode:    armcompute.WindowsPatchAssessmentMode("ImageDefault").ToPtr(),
								EnableHotpatching: to.BoolPtr(true),
								PatchMode:         armcompute.WindowsVMGuestPatchMode("AutomaticByOS").ToPtr(),
							},
							ProvisionVMAgent: to.BoolPtr(true),
							TimeZone:         to.StringPtr("<time-zone>"),
							WinRM: &armcompute.WinRMConfiguration{
								Listeners: []*armcompute.WinRMListener{
									{
										CertificateURL: to.StringPtr("<certificate-url>"),
										Protocol:       armcompute.ProtocolTypesHTTP.ToPtr(),
									}},
							},
						},
					},
					ScheduledEventsProfile: &armcompute.ScheduledEventsProfile{
						TerminateNotificationProfile: &armcompute.TerminateNotificationProfile{
							Enable:           to.BoolPtr(true),
							NotBeforeTimeout: to.StringPtr("<not-before-timeout>"),
						},
					},
					SecurityProfile: &armcompute.SecurityProfile{
						EncryptionAtHost: to.BoolPtr(true),
						SecurityType:     armcompute.SecurityTypes("TrustedLaunch").ToPtr(),
						UefiSettings: &armcompute.UefiSettings{
							SecureBootEnabled: to.BoolPtr(true),
							VTpmEnabled:       to.BoolPtr(true),
						},
					},
					StorageProfile: &armcompute.VirtualMachineScaleSetUpdateStorageProfile{
						DataDisks: []*armcompute.VirtualMachineScaleSetDataDisk{
							{
								Name:              to.StringPtr("<name>"),
								Caching:           armcompute.CachingTypesNone.ToPtr(),
								CreateOption:      armcompute.DiskCreateOptionTypes("Empty").ToPtr(),
								DiskIOPSReadWrite: to.Int64Ptr(28),
								DiskMBpsReadWrite: to.Int64Ptr(15),
								DiskSizeGB:        to.Int32Ptr(1023),
								Lun:               to.Int32Ptr(26),
								ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
									DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
										ID: to.StringPtr("<id>"),
									},
									StorageAccountType: armcompute.StorageAccountTypes("Standard_LRS").ToPtr(),
								},
								WriteAcceleratorEnabled: to.BoolPtr(true),
							}},
						ImageReference: &armcompute.ImageReference{
							ID:                   to.StringPtr("<id>"),
							Offer:                to.StringPtr("<offer>"),
							Publisher:            to.StringPtr("<publisher>"),
							SharedGalleryImageID: to.StringPtr("<shared-gallery-image-id>"),
							SKU:                  to.StringPtr("<sku>"),
							Version:              to.StringPtr("<version>"),
						},
						OSDisk: &armcompute.VirtualMachineScaleSetUpdateOSDisk{
							Caching:    armcompute.CachingTypesReadWrite.ToPtr(),
							DiskSizeGB: to.Int32Ptr(6),
							Image: &armcompute.VirtualHardDisk{
								URI: to.StringPtr("<uri>"),
							},
							ManagedDisk: &armcompute.VirtualMachineScaleSetManagedDiskParameters{
								DiskEncryptionSet: &armcompute.DiskEncryptionSetParameters{
									ID: to.StringPtr("<id>"),
								},
								StorageAccountType: armcompute.StorageAccountTypes("Standard_LRS").ToPtr(),
							},
							VhdContainers: []*string{
								to.StringPtr("aa")},
							WriteAcceleratorEnabled: to.BoolPtr(true),
						},
					},
					UserData: to.StringPtr("<user-data>"),
				},
			},
			SKU: &armcompute.SKU{
				Name:     to.StringPtr("<name>"),
				Capacity: to.Int64Ptr(7),
				Tier:     to.StringPtr("<tier>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualMachineScaleSetsClientUpdateResult)
}
```
