package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPVirtualInstances_Create_CustomFullResourceNames_HA_AvSet.json
func ExampleSAPVirtualInstancesClient_BeginCreate_createInfrastructureWithOsConfigurationWithCustomResourceNamesForHaSystemWithAvailabilitySet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloads.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPVirtualInstancesClient().BeginCreate(ctx, "test-rg", "X00", armworkloads.SAPVirtualInstance{
		Location: to.Ptr("westcentralus"),
		Tags:     map[string]*string{},
		Properties: &armworkloads.SAPVirtualInstanceProperties{
			Configuration: &armworkloads.DeploymentWithOSConfiguration{
				ConfigurationType: to.Ptr(armworkloads.SAPConfigurationTypeDeploymentWithOSConfig),
				AppLocation:       to.Ptr("eastus"),
				InfrastructureConfiguration: &armworkloads.ThreeTierConfiguration{
					AppResourceGroup: to.Ptr("X00-RG"),
					DeploymentType:   to.Ptr(armworkloads.SAPDeploymentTypeThreeTier),
					ApplicationServer: &armworkloads.ApplicationServerConfiguration{
						InstanceCount: to.Ptr[int64](6),
						SubnetID:      to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
						VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
							ImageReference: &armworkloads.ImageReference{
								Offer:     to.Ptr("RHEL-SAP"),
								Publisher: to.Ptr("RedHat"),
								SKU:       to.Ptr("84sapha-gen2"),
								Version:   to.Ptr("latest"),
							},
							OSProfile: &armworkloads.OSProfile{
								AdminUsername: to.Ptr("{your-username}"),
								OSConfiguration: &armworkloads.LinuxConfiguration{
									OSType:                        to.Ptr(armworkloads.OSTypeLinux),
									DisablePasswordAuthentication: to.Ptr(true),
									SSHKeyPair: &armworkloads.SSHKeyPair{
										PrivateKey: to.Ptr("xyz"),
										PublicKey:  to.Ptr("abc"),
									},
								},
							},
							VMSize: to.Ptr("Standard_E32ds_v4"),
						},
					},
					CentralServer: &armworkloads.CentralServerConfiguration{
						InstanceCount: to.Ptr[int64](2),
						SubnetID:      to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
						VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
							ImageReference: &armworkloads.ImageReference{
								Offer:     to.Ptr("RHEL-SAP"),
								Publisher: to.Ptr("RedHat"),
								SKU:       to.Ptr("84sapha-gen2"),
								Version:   to.Ptr("latest"),
							},
							OSProfile: &armworkloads.OSProfile{
								AdminUsername: to.Ptr("{your-username}"),
								OSConfiguration: &armworkloads.LinuxConfiguration{
									OSType:                        to.Ptr(armworkloads.OSTypeLinux),
									DisablePasswordAuthentication: to.Ptr(true),
									SSHKeyPair: &armworkloads.SSHKeyPair{
										PrivateKey: to.Ptr("xyz"),
										PublicKey:  to.Ptr("abc"),
									},
								},
							},
							VMSize: to.Ptr("Standard_E16ds_v4"),
						},
					},
					CustomResourceNames: &armworkloads.ThreeTierFullResourceNames{
						NamingPatternType: to.Ptr(armworkloads.NamingPatternTypeFullResourceName),
						ApplicationServer: &armworkloads.ApplicationServerFullResourceNames{
							AvailabilitySetName: to.Ptr("appAvSet"),
							VirtualMachines: []*armworkloads.VirtualMachineResourceNames{
								{
									DataDiskNames: map[string][]*string{
										"default": {
											to.Ptr("app0disk0")},
									},
									HostName: to.Ptr("apphostName0"),
									NetworkInterfaces: []*armworkloads.NetworkInterfaceResourceNames{
										{
											NetworkInterfaceName: to.Ptr("appnic0"),
										}},
									OSDiskName: to.Ptr("app0osdisk"),
									VMName:     to.Ptr("appvm0"),
								},
								{
									DataDiskNames: map[string][]*string{
										"default": {
											to.Ptr("app1disk0")},
									},
									HostName: to.Ptr("apphostName1"),
									NetworkInterfaces: []*armworkloads.NetworkInterfaceResourceNames{
										{
											NetworkInterfaceName: to.Ptr("appnic1"),
										}},
									OSDiskName: to.Ptr("app1osdisk"),
									VMName:     to.Ptr("appvm1"),
								}},
						},
						CentralServer: &armworkloads.CentralServerFullResourceNames{
							AvailabilitySetName: to.Ptr("csAvSet"),
							LoadBalancer: &armworkloads.LoadBalancerResourceNames{
								BackendPoolNames: []*string{
									to.Ptr("ascsBackendPool")},
								FrontendIPConfigurationNames: []*string{
									to.Ptr("ascsip0"),
									to.Ptr("ersip0")},
								HealthProbeNames: []*string{
									to.Ptr("ascsHealthProbe"),
									to.Ptr("ersHealthProbe")},
								LoadBalancerName: to.Ptr("ascslb"),
							},
							VirtualMachines: []*armworkloads.VirtualMachineResourceNames{
								{
									HostName: to.Ptr("ascshostName"),
									NetworkInterfaces: []*armworkloads.NetworkInterfaceResourceNames{
										{
											NetworkInterfaceName: to.Ptr("ascsnic"),
										}},
									OSDiskName: to.Ptr("ascsosdisk"),
									VMName:     to.Ptr("ascsvm"),
								},
								{
									HostName: to.Ptr("ershostName"),
									NetworkInterfaces: []*armworkloads.NetworkInterfaceResourceNames{
										{
											NetworkInterfaceName: to.Ptr("ersnic"),
										}},
									OSDiskName: to.Ptr("ersosdisk"),
									VMName:     to.Ptr("ersvm"),
								}},
						},
						DatabaseServer: &armworkloads.DatabaseServerFullResourceNames{
							AvailabilitySetName: to.Ptr("dbAvSet"),
							LoadBalancer: &armworkloads.LoadBalancerResourceNames{
								BackendPoolNames: []*string{
									to.Ptr("dbBackendPool")},
								FrontendIPConfigurationNames: []*string{
									to.Ptr("dbip")},
								HealthProbeNames: []*string{
									to.Ptr("dbHealthProbe")},
								LoadBalancerName: to.Ptr("dblb"),
							},
							VirtualMachines: []*armworkloads.VirtualMachineResourceNames{
								{
									DataDiskNames: map[string][]*string{
										"hanaData": {
											to.Ptr("hanadatapr0"),
											to.Ptr("hanadatapr1")},
										"hanaLog": {
											to.Ptr("hanalogpr0"),
											to.Ptr("hanalogpr1"),
											to.Ptr("hanalogpr2")},
										"hanaShared": {
											to.Ptr("hanasharedpr0"),
											to.Ptr("hanasharedpr1")},
										"usrSap": {
											to.Ptr("usrsappr0")},
									},
									HostName: to.Ptr("dbprhostName"),
									NetworkInterfaces: []*armworkloads.NetworkInterfaceResourceNames{
										{
											NetworkInterfaceName: to.Ptr("dbprnic"),
										}},
									OSDiskName: to.Ptr("dbprosdisk"),
									VMName:     to.Ptr("dbvmpr"),
								},
								{
									DataDiskNames: map[string][]*string{
										"hanaData": {
											to.Ptr("hanadatasr0"),
											to.Ptr("hanadatasr1")},
										"hanaLog": {
											to.Ptr("hanalogsr0"),
											to.Ptr("hanalogsr1"),
											to.Ptr("hanalogsr2")},
										"hanaShared": {
											to.Ptr("hanasharedsr0"),
											to.Ptr("hanasharedsr1")},
										"usrSap": {
											to.Ptr("usrsapsr0")},
									},
									HostName: to.Ptr("dbsrhostName"),
									NetworkInterfaces: []*armworkloads.NetworkInterfaceResourceNames{
										{
											NetworkInterfaceName: to.Ptr("dbsrnic"),
										}},
									OSDiskName: to.Ptr("dbsrosdisk"),
									VMName:     to.Ptr("dbvmsr"),
								}},
						},
						SharedStorage: &armworkloads.SharedStorageResourceNames{
							SharedStorageAccountName:                to.Ptr("storageacc"),
							SharedStorageAccountPrivateEndPointName: to.Ptr("peForxNFS"),
						},
					},
					DatabaseServer: &armworkloads.DatabaseConfiguration{
						DatabaseType:  to.Ptr(armworkloads.SAPDatabaseTypeHANA),
						InstanceCount: to.Ptr[int64](2),
						SubnetID:      to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/dbsubnet"),
						VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
							ImageReference: &armworkloads.ImageReference{
								Offer:     to.Ptr("RHEL-SAP"),
								Publisher: to.Ptr("RedHat"),
								SKU:       to.Ptr("84sapha-gen2"),
								Version:   to.Ptr("latest"),
							},
							OSProfile: &armworkloads.OSProfile{
								AdminUsername: to.Ptr("{your-username}"),
								OSConfiguration: &armworkloads.LinuxConfiguration{
									OSType:                        to.Ptr(armworkloads.OSTypeLinux),
									DisablePasswordAuthentication: to.Ptr(true),
									SSHKeyPair: &armworkloads.SSHKeyPair{
										PrivateKey: to.Ptr("xyz"),
										PublicKey:  to.Ptr("abc"),
									},
								},
							},
							VMSize: to.Ptr("Standard_M32ts"),
						},
					},
					HighAvailabilityConfig: &armworkloads.HighAvailabilityConfiguration{
						HighAvailabilityType: to.Ptr(armworkloads.SAPHighAvailabilityTypeAvailabilitySet),
					},
				},
				OSSapConfiguration: &armworkloads.OsSapConfiguration{
					SapFqdn: to.Ptr("xyz.test.com"),
				},
			},
			Environment: to.Ptr(armworkloads.SAPEnvironmentTypeProd),
			SapProduct:  to.Ptr(armworkloads.SAPProductTypeS4HANA),
		},
	}, nil)
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
	// res.SAPVirtualInstance = armworkloads.SAPVirtualInstance{
	// 	Name: to.Ptr("X00"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00"),
	// 	SystemData: &armworkloads.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@xyz.com"),
	// 		CreatedByType: to.Ptr(armworkloads.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@xyz.com"),
	// 		LastModifiedByType: to.Ptr(armworkloads.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("westcentralus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armworkloads.SAPVirtualInstanceProperties{
	// 		Configuration: &armworkloads.DeploymentWithOSConfiguration{
	// 			ConfigurationType: to.Ptr(armworkloads.SAPConfigurationTypeDeploymentWithOSConfig),
	// 			AppLocation: to.Ptr("eastus"),
	// 			InfrastructureConfiguration: &armworkloads.ThreeTierConfiguration{
	// 				AppResourceGroup: to.Ptr("X00-RG"),
	// 				DeploymentType: to.Ptr(armworkloads.SAPDeploymentTypeThreeTier),
	// 				ApplicationServer: &armworkloads.ApplicationServerConfiguration{
	// 					InstanceCount: to.Ptr[int64](6),
	// 					SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
	// 					VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
	// 						ImageReference: &armworkloads.ImageReference{
	// 							Offer: to.Ptr("RHEL-SAP"),
	// 							Publisher: to.Ptr("RedHat"),
	// 							SKU: to.Ptr("84sapha-gen2"),
	// 							Version: to.Ptr("latest"),
	// 						},
	// 						OSProfile: &armworkloads.OSProfile{
	// 							AdminUsername: to.Ptr("{your-username}"),
	// 							OSConfiguration: &armworkloads.LinuxConfiguration{
	// 								OSType: to.Ptr(armworkloads.OSTypeLinux),
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								SSHKeyPair: &armworkloads.SSHKeyPair{
	// 									PublicKey: to.Ptr("abc"),
	// 								},
	// 							},
	// 						},
	// 						VMSize: to.Ptr("Standard_E32ds_v4"),
	// 					},
	// 				},
	// 				CentralServer: &armworkloads.CentralServerConfiguration{
	// 					InstanceCount: to.Ptr[int64](2),
	// 					SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
	// 					VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
	// 						ImageReference: &armworkloads.ImageReference{
	// 							Offer: to.Ptr("RHEL-SAP"),
	// 							Publisher: to.Ptr("RedHat"),
	// 							SKU: to.Ptr("84sapha-gen2"),
	// 							Version: to.Ptr("latest"),
	// 						},
	// 						OSProfile: &armworkloads.OSProfile{
	// 							AdminUsername: to.Ptr("{your-username}"),
	// 							OSConfiguration: &armworkloads.LinuxConfiguration{
	// 								OSType: to.Ptr(armworkloads.OSTypeLinux),
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								SSHKeyPair: &armworkloads.SSHKeyPair{
	// 									PublicKey: to.Ptr("abc"),
	// 								},
	// 							},
	// 						},
	// 						VMSize: to.Ptr("Standard_E16ds_v4"),
	// 					},
	// 				},
	// 				CustomResourceNames: &armworkloads.ThreeTierFullResourceNames{
	// 					NamingPatternType: to.Ptr(armworkloads.NamingPatternTypeFullResourceName),
	// 					ApplicationServer: &armworkloads.ApplicationServerFullResourceNames{
	// 						AvailabilitySetName: to.Ptr("appAvSet"),
	// 						VirtualMachines: []*armworkloads.VirtualMachineResourceNames{
	// 							{
	// 								DataDiskNames: map[string][]*string{
	// 									"default": []*string{
	// 										to.Ptr("app0disk0")},
	// 									},
	// 									HostName: to.Ptr("apphostName0"),
	// 									NetworkInterfaces: []*armworkloads.NetworkInterfaceResourceNames{
	// 										{
	// 											NetworkInterfaceName: to.Ptr("appnic0"),
	// 									}},
	// 									OSDiskName: to.Ptr("app0osdisk"),
	// 									VMName: to.Ptr("appvm0"),
	// 								},
	// 								{
	// 									DataDiskNames: map[string][]*string{
	// 										"default": []*string{
	// 											to.Ptr("app1disk0")},
	// 										},
	// 										HostName: to.Ptr("apphostName1"),
	// 										NetworkInterfaces: []*armworkloads.NetworkInterfaceResourceNames{
	// 											{
	// 												NetworkInterfaceName: to.Ptr("appnic1"),
	// 										}},
	// 										OSDiskName: to.Ptr("app1osdisk"),
	// 										VMName: to.Ptr("appvm1"),
	// 								}},
	// 							},
	// 							CentralServer: &armworkloads.CentralServerFullResourceNames{
	// 								AvailabilitySetName: to.Ptr("csAvSet"),
	// 								LoadBalancer: &armworkloads.LoadBalancerResourceNames{
	// 									BackendPoolNames: []*string{
	// 										to.Ptr("ascsBackendPool")},
	// 										FrontendIPConfigurationNames: []*string{
	// 											to.Ptr("ascsip0"),
	// 											to.Ptr("ersip0")},
	// 											HealthProbeNames: []*string{
	// 												to.Ptr("ascsHealthProbe"),
	// 												to.Ptr("ersHealthProbe")},
	// 												LoadBalancerName: to.Ptr("ascslb"),
	// 											},
	// 											VirtualMachines: []*armworkloads.VirtualMachineResourceNames{
	// 												{
	// 													HostName: to.Ptr("ascshostName"),
	// 													NetworkInterfaces: []*armworkloads.NetworkInterfaceResourceNames{
	// 														{
	// 															NetworkInterfaceName: to.Ptr("ascsnic"),
	// 													}},
	// 													OSDiskName: to.Ptr("ascsosdisk"),
	// 													VMName: to.Ptr("ascsvm"),
	// 												},
	// 												{
	// 													HostName: to.Ptr("ershostName"),
	// 													NetworkInterfaces: []*armworkloads.NetworkInterfaceResourceNames{
	// 														{
	// 															NetworkInterfaceName: to.Ptr("ersnic"),
	// 													}},
	// 													OSDiskName: to.Ptr("ersosdisk"),
	// 													VMName: to.Ptr("ersvm"),
	// 											}},
	// 										},
	// 										DatabaseServer: &armworkloads.DatabaseServerFullResourceNames{
	// 											AvailabilitySetName: to.Ptr("dbAvSet"),
	// 											LoadBalancer: &armworkloads.LoadBalancerResourceNames{
	// 												BackendPoolNames: []*string{
	// 													to.Ptr("dbBackendPool")},
	// 													FrontendIPConfigurationNames: []*string{
	// 														to.Ptr("dbip")},
	// 														HealthProbeNames: []*string{
	// 															to.Ptr("dbHealthProbe")},
	// 															LoadBalancerName: to.Ptr("dblb"),
	// 														},
	// 														VirtualMachines: []*armworkloads.VirtualMachineResourceNames{
	// 															{
	// 																DataDiskNames: map[string][]*string{
	// 																	"hanaData": []*string{
	// 																		to.Ptr("hanadatapr0"),
	// 																		to.Ptr("hanadatapr1")},
	// 																		"hanaLog": []*string{
	// 																			to.Ptr("hanalogpr0"),
	// 																			to.Ptr("hanalogpr1"),
	// 																			to.Ptr("hanalogpr2")},
	// 																			"hanaShared": []*string{
	// 																				to.Ptr("hanasharedpr0"),
	// 																				to.Ptr("hanasharedpr1")},
	// 																				"usrSap": []*string{
	// 																					to.Ptr("usrsappr0")},
	// 																				},
	// 																				HostName: to.Ptr("dbprhostName"),
	// 																				NetworkInterfaces: []*armworkloads.NetworkInterfaceResourceNames{
	// 																					{
	// 																						NetworkInterfaceName: to.Ptr("dbprnic"),
	// 																				}},
	// 																				OSDiskName: to.Ptr("dbprosdisk"),
	// 																				VMName: to.Ptr("dbvmpr"),
	// 																			},
	// 																			{
	// 																				DataDiskNames: map[string][]*string{
	// 																					"hanaData": []*string{
	// 																						to.Ptr("hanadatasr0"),
	// 																						to.Ptr("hanadatasr1")},
	// 																						"hanaLog": []*string{
	// 																							to.Ptr("hanalogsr0"),
	// 																							to.Ptr("hanalogsr1"),
	// 																							to.Ptr("hanalogsr2")},
	// 																							"hanaShared": []*string{
	// 																								to.Ptr("hanasharedsr0"),
	// 																								to.Ptr("hanasharedsr1")},
	// 																								"usrSap": []*string{
	// 																									to.Ptr("usrsapsr0")},
	// 																								},
	// 																								HostName: to.Ptr("dbsrhostName"),
	// 																								NetworkInterfaces: []*armworkloads.NetworkInterfaceResourceNames{
	// 																									{
	// 																										NetworkInterfaceName: to.Ptr("dbsrnic"),
	// 																								}},
	// 																								OSDiskName: to.Ptr("dbsrosdisk"),
	// 																								VMName: to.Ptr("dbvmsr"),
	// 																						}},
	// 																					},
	// 																					SharedStorage: &armworkloads.SharedStorageResourceNames{
	// 																						SharedStorageAccountName: to.Ptr("storageacc"),
	// 																						SharedStorageAccountPrivateEndPointName: to.Ptr("peForxNFS"),
	// 																					},
	// 																				},
	// 																				DatabaseServer: &armworkloads.DatabaseConfiguration{
	// 																					DatabaseType: to.Ptr(armworkloads.SAPDatabaseTypeHANA),
	// 																					InstanceCount: to.Ptr[int64](2),
	// 																					SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/dbsubnet"),
	// 																					VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
	// 																						ImageReference: &armworkloads.ImageReference{
	// 																							Offer: to.Ptr("RHEL-SAP"),
	// 																							Publisher: to.Ptr("RedHat"),
	// 																							SKU: to.Ptr("84sapha-gen2"),
	// 																							Version: to.Ptr("latest"),
	// 																						},
	// 																						OSProfile: &armworkloads.OSProfile{
	// 																							AdminUsername: to.Ptr("{your-username}"),
	// 																							OSConfiguration: &armworkloads.LinuxConfiguration{
	// 																								OSType: to.Ptr(armworkloads.OSTypeLinux),
	// 																								DisablePasswordAuthentication: to.Ptr(true),
	// 																								SSHKeyPair: &armworkloads.SSHKeyPair{
	// 																									PublicKey: to.Ptr("abc"),
	// 																								},
	// 																							},
	// 																						},
	// 																						VMSize: to.Ptr("Standard_M32ts"),
	// 																					},
	// 																				},
	// 																				HighAvailabilityConfig: &armworkloads.HighAvailabilityConfiguration{
	// 																					HighAvailabilityType: to.Ptr(armworkloads.SAPHighAvailabilityTypeAvailabilitySet),
	// 																				},
	// 																			},
	// 																			OSSapConfiguration: &armworkloads.OsSapConfiguration{
	// 																				SapFqdn: to.Ptr("xyz.test.com"),
	// 																			},
	// 																		},
	// 																		Environment: to.Ptr(armworkloads.SAPEnvironmentTypeProd),
	// 																		Health: to.Ptr(armworkloads.SAPHealthStateUnknown),
	// 																		ProvisioningState: to.Ptr(armworkloads.SapVirtualInstanceProvisioningStateSucceeded),
	// 																		SapProduct: to.Ptr(armworkloads.SAPProductTypeS4HANA),
	// 																		State: to.Ptr(armworkloads.SAPVirtualInstanceStateSoftwareInstallationPending),
	// 																		Status: to.Ptr(armworkloads.SAPVirtualInstanceStatus("Unknown")),
	// 																	},
	// 																}
}
