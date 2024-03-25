package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Create_CustomFullResourceNames_HA_AvSet.json
func ExampleSAPVirtualInstancesClient_BeginCreate_createInfrastructureWithOsConfigurationWithCustomResourceNamesForHaSystemWithAvailabilitySet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPVirtualInstancesClient().BeginCreate(ctx, "test-rg", "X00", armworkloadssapvirtualinstance.SAPVirtualInstance{
		Location: to.Ptr("eastus"),
		Tags:     map[string]*string{},
		Properties: &armworkloadssapvirtualinstance.SAPVirtualInstanceProperties{
			Configuration: &armworkloadssapvirtualinstance.DeploymentWithOSConfiguration{
				ConfigurationType: to.Ptr(armworkloadssapvirtualinstance.SAPConfigurationTypeDeploymentWithOSConfig),
				AppLocation:       to.Ptr("eastus"),
				InfrastructureConfiguration: &armworkloadssapvirtualinstance.ThreeTierConfiguration{
					AppResourceGroup: to.Ptr("X00-RG"),
					DeploymentType:   to.Ptr(armworkloadssapvirtualinstance.SAPDeploymentTypeThreeTier),
					ApplicationServer: &armworkloadssapvirtualinstance.ApplicationServerConfiguration{
						InstanceCount: to.Ptr[int64](6),
						SubnetID:      to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
						VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
							ImageReference: &armworkloadssapvirtualinstance.ImageReference{
								Offer:     to.Ptr("RHEL-SAP"),
								Publisher: to.Ptr("RedHat"),
								SKU:       to.Ptr("84sapha-gen2"),
								Version:   to.Ptr("latest"),
							},
							OSProfile: &armworkloadssapvirtualinstance.OSProfile{
								AdminUsername: to.Ptr("{your-username}"),
								OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
									OSType:                        to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
									DisablePasswordAuthentication: to.Ptr(true),
									SSHKeyPair: &armworkloadssapvirtualinstance.SSHKeyPair{
										PrivateKey: to.Ptr("xyz"),
										PublicKey:  to.Ptr("abc"),
									},
								},
							},
							VMSize: to.Ptr("Standard_E32ds_v4"),
						},
					},
					CentralServer: &armworkloadssapvirtualinstance.CentralServerConfiguration{
						InstanceCount: to.Ptr[int64](2),
						SubnetID:      to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
						VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
							ImageReference: &armworkloadssapvirtualinstance.ImageReference{
								Offer:     to.Ptr("RHEL-SAP"),
								Publisher: to.Ptr("RedHat"),
								SKU:       to.Ptr("84sapha-gen2"),
								Version:   to.Ptr("latest"),
							},
							OSProfile: &armworkloadssapvirtualinstance.OSProfile{
								AdminUsername: to.Ptr("{your-username}"),
								OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
									OSType:                        to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
									DisablePasswordAuthentication: to.Ptr(true),
									SSHKeyPair: &armworkloadssapvirtualinstance.SSHKeyPair{
										PrivateKey: to.Ptr("xyz"),
										PublicKey:  to.Ptr("abc"),
									},
								},
							},
							VMSize: to.Ptr("Standard_E16ds_v4"),
						},
					},
					CustomResourceNames: &armworkloadssapvirtualinstance.ThreeTierFullResourceNames{
						NamingPatternType: to.Ptr(armworkloadssapvirtualinstance.NamingPatternTypeFullResourceName),
						ApplicationServer: &armworkloadssapvirtualinstance.ApplicationServerFullResourceNames{
							AvailabilitySetName: to.Ptr("appAvSet"),
							VirtualMachines: []*armworkloadssapvirtualinstance.VirtualMachineResourceNames{
								{
									DataDiskNames: map[string][]*string{
										"default": {
											to.Ptr("app0disk0")},
									},
									HostName: to.Ptr("apphostName0"),
									NetworkInterfaces: []*armworkloadssapvirtualinstance.NetworkInterfaceResourceNames{
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
									NetworkInterfaces: []*armworkloadssapvirtualinstance.NetworkInterfaceResourceNames{
										{
											NetworkInterfaceName: to.Ptr("appnic1"),
										}},
									OSDiskName: to.Ptr("app1osdisk"),
									VMName:     to.Ptr("appvm1"),
								}},
						},
						CentralServer: &armworkloadssapvirtualinstance.CentralServerFullResourceNames{
							AvailabilitySetName: to.Ptr("csAvSet"),
							LoadBalancer: &armworkloadssapvirtualinstance.LoadBalancerResourceNames{
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
							VirtualMachines: []*armworkloadssapvirtualinstance.VirtualMachineResourceNames{
								{
									HostName: to.Ptr("ascshostName"),
									NetworkInterfaces: []*armworkloadssapvirtualinstance.NetworkInterfaceResourceNames{
										{
											NetworkInterfaceName: to.Ptr("ascsnic"),
										}},
									OSDiskName: to.Ptr("ascsosdisk"),
									VMName:     to.Ptr("ascsvm"),
								},
								{
									HostName: to.Ptr("ershostName"),
									NetworkInterfaces: []*armworkloadssapvirtualinstance.NetworkInterfaceResourceNames{
										{
											NetworkInterfaceName: to.Ptr("ersnic"),
										}},
									OSDiskName: to.Ptr("ersosdisk"),
									VMName:     to.Ptr("ersvm"),
								}},
						},
						DatabaseServer: &armworkloadssapvirtualinstance.DatabaseServerFullResourceNames{
							AvailabilitySetName: to.Ptr("dbAvSet"),
							LoadBalancer: &armworkloadssapvirtualinstance.LoadBalancerResourceNames{
								BackendPoolNames: []*string{
									to.Ptr("dbBackendPool")},
								FrontendIPConfigurationNames: []*string{
									to.Ptr("dbip")},
								HealthProbeNames: []*string{
									to.Ptr("dbHealthProbe")},
								LoadBalancerName: to.Ptr("dblb"),
							},
							VirtualMachines: []*armworkloadssapvirtualinstance.VirtualMachineResourceNames{
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
									NetworkInterfaces: []*armworkloadssapvirtualinstance.NetworkInterfaceResourceNames{
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
									NetworkInterfaces: []*armworkloadssapvirtualinstance.NetworkInterfaceResourceNames{
										{
											NetworkInterfaceName: to.Ptr("dbsrnic"),
										}},
									OSDiskName: to.Ptr("dbsrosdisk"),
									VMName:     to.Ptr("dbvmsr"),
								}},
						},
						SharedStorage: &armworkloadssapvirtualinstance.SharedStorageResourceNames{
							SharedStorageAccountName:                to.Ptr("storageacc"),
							SharedStorageAccountPrivateEndPointName: to.Ptr("peForxNFS"),
						},
					},
					DatabaseServer: &armworkloadssapvirtualinstance.DatabaseConfiguration{
						DatabaseType:  to.Ptr(armworkloadssapvirtualinstance.SAPDatabaseTypeHANA),
						InstanceCount: to.Ptr[int64](2),
						SubnetID:      to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/dbsubnet"),
						VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
							ImageReference: &armworkloadssapvirtualinstance.ImageReference{
								Offer:     to.Ptr("RHEL-SAP"),
								Publisher: to.Ptr("RedHat"),
								SKU:       to.Ptr("84sapha-gen2"),
								Version:   to.Ptr("latest"),
							},
							OSProfile: &armworkloadssapvirtualinstance.OSProfile{
								AdminUsername: to.Ptr("{your-username}"),
								OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
									OSType:                        to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
									DisablePasswordAuthentication: to.Ptr(true),
									SSHKeyPair: &armworkloadssapvirtualinstance.SSHKeyPair{
										PrivateKey: to.Ptr("xyz"),
										PublicKey:  to.Ptr("abc"),
									},
								},
							},
							VMSize: to.Ptr("Standard_M32ts"),
						},
					},
					HighAvailabilityConfig: &armworkloadssapvirtualinstance.HighAvailabilityConfiguration{
						HighAvailabilityType: to.Ptr(armworkloadssapvirtualinstance.SAPHighAvailabilityTypeAvailabilitySet),
					},
				},
				OSSapConfiguration: &armworkloadssapvirtualinstance.OsSapConfiguration{
					SapFqdn: to.Ptr("xyz.test.com"),
				},
			},
			Environment: to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeProd),
			ManagedResourceGroupConfiguration: &armworkloadssapvirtualinstance.ManagedRGConfiguration{
				Name: to.Ptr("mrg-X00-8e17e36c-42e9-4cd5-a078-7b44883414e0"),
			},
			SapProduct: to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
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
	// res.SAPVirtualInstance = armworkloadssapvirtualinstance.SAPVirtualInstance{
	// 	Name: to.Ptr("X00"),
	// 	Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances"),
	// 	ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00"),
	// 	SystemData: &armworkloadssapvirtualinstance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		CreatedBy: to.Ptr("user@xyz.com"),
	// 		CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user@xyz.com"),
	// 		LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armworkloadssapvirtualinstance.SAPVirtualInstanceProperties{
	// 		Configuration: &armworkloadssapvirtualinstance.DeploymentWithOSConfiguration{
	// 			ConfigurationType: to.Ptr(armworkloadssapvirtualinstance.SAPConfigurationTypeDeploymentWithOSConfig),
	// 			AppLocation: to.Ptr("eastus"),
	// 			InfrastructureConfiguration: &armworkloadssapvirtualinstance.ThreeTierConfiguration{
	// 				AppResourceGroup: to.Ptr("X00-RG"),
	// 				DeploymentType: to.Ptr(armworkloadssapvirtualinstance.SAPDeploymentTypeThreeTier),
	// 				ApplicationServer: &armworkloadssapvirtualinstance.ApplicationServerConfiguration{
	// 					InstanceCount: to.Ptr[int64](6),
	// 					SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
	// 					VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
	// 						ImageReference: &armworkloadssapvirtualinstance.ImageReference{
	// 							Offer: to.Ptr("RHEL-SAP"),
	// 							Publisher: to.Ptr("RedHat"),
	// 							SKU: to.Ptr("84sapha-gen2"),
	// 							Version: to.Ptr("latest"),
	// 						},
	// 						OSProfile: &armworkloadssapvirtualinstance.OSProfile{
	// 							AdminUsername: to.Ptr("{your-username}"),
	// 							OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
	// 								OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								SSHKeyPair: &armworkloadssapvirtualinstance.SSHKeyPair{
	// 									PublicKey: to.Ptr("abc"),
	// 								},
	// 							},
	// 						},
	// 						VMSize: to.Ptr("Standard_E32ds_v4"),
	// 					},
	// 				},
	// 				CentralServer: &armworkloadssapvirtualinstance.CentralServerConfiguration{
	// 					InstanceCount: to.Ptr[int64](2),
	// 					SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
	// 					VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
	// 						ImageReference: &armworkloadssapvirtualinstance.ImageReference{
	// 							Offer: to.Ptr("RHEL-SAP"),
	// 							Publisher: to.Ptr("RedHat"),
	// 							SKU: to.Ptr("84sapha-gen2"),
	// 							Version: to.Ptr("latest"),
	// 						},
	// 						OSProfile: &armworkloadssapvirtualinstance.OSProfile{
	// 							AdminUsername: to.Ptr("{your-username}"),
	// 							OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
	// 								OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								SSHKeyPair: &armworkloadssapvirtualinstance.SSHKeyPair{
	// 									PublicKey: to.Ptr("abc"),
	// 								},
	// 							},
	// 						},
	// 						VMSize: to.Ptr("Standard_E16ds_v4"),
	// 					},
	// 				},
	// 				CustomResourceNames: &armworkloadssapvirtualinstance.ThreeTierFullResourceNames{
	// 					NamingPatternType: to.Ptr(armworkloadssapvirtualinstance.NamingPatternTypeFullResourceName),
	// 					ApplicationServer: &armworkloadssapvirtualinstance.ApplicationServerFullResourceNames{
	// 						AvailabilitySetName: to.Ptr("appAvSet"),
	// 						VirtualMachines: []*armworkloadssapvirtualinstance.VirtualMachineResourceNames{
	// 							{
	// 								DataDiskNames: map[string][]*string{
	// 									"default": []*string{
	// 										to.Ptr("app0disk0")},
	// 									},
	// 									HostName: to.Ptr("apphostName0"),
	// 									NetworkInterfaces: []*armworkloadssapvirtualinstance.NetworkInterfaceResourceNames{
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
	// 										NetworkInterfaces: []*armworkloadssapvirtualinstance.NetworkInterfaceResourceNames{
	// 											{
	// 												NetworkInterfaceName: to.Ptr("appnic1"),
	// 										}},
	// 										OSDiskName: to.Ptr("app1osdisk"),
	// 										VMName: to.Ptr("appvm1"),
	// 								}},
	// 							},
	// 							CentralServer: &armworkloadssapvirtualinstance.CentralServerFullResourceNames{
	// 								AvailabilitySetName: to.Ptr("csAvSet"),
	// 								LoadBalancer: &armworkloadssapvirtualinstance.LoadBalancerResourceNames{
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
	// 											VirtualMachines: []*armworkloadssapvirtualinstance.VirtualMachineResourceNames{
	// 												{
	// 													HostName: to.Ptr("ascshostName"),
	// 													NetworkInterfaces: []*armworkloadssapvirtualinstance.NetworkInterfaceResourceNames{
	// 														{
	// 															NetworkInterfaceName: to.Ptr("ascsnic"),
	// 													}},
	// 													OSDiskName: to.Ptr("ascsosdisk"),
	// 													VMName: to.Ptr("ascsvm"),
	// 												},
	// 												{
	// 													HostName: to.Ptr("ershostName"),
	// 													NetworkInterfaces: []*armworkloadssapvirtualinstance.NetworkInterfaceResourceNames{
	// 														{
	// 															NetworkInterfaceName: to.Ptr("ersnic"),
	// 													}},
	// 													OSDiskName: to.Ptr("ersosdisk"),
	// 													VMName: to.Ptr("ersvm"),
	// 											}},
	// 										},
	// 										DatabaseServer: &armworkloadssapvirtualinstance.DatabaseServerFullResourceNames{
	// 											AvailabilitySetName: to.Ptr("dbAvSet"),
	// 											LoadBalancer: &armworkloadssapvirtualinstance.LoadBalancerResourceNames{
	// 												BackendPoolNames: []*string{
	// 													to.Ptr("dbBackendPool")},
	// 													FrontendIPConfigurationNames: []*string{
	// 														to.Ptr("dbip")},
	// 														HealthProbeNames: []*string{
	// 															to.Ptr("dbHealthProbe")},
	// 															LoadBalancerName: to.Ptr("dblb"),
	// 														},
	// 														VirtualMachines: []*armworkloadssapvirtualinstance.VirtualMachineResourceNames{
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
	// 																				NetworkInterfaces: []*armworkloadssapvirtualinstance.NetworkInterfaceResourceNames{
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
	// 																								NetworkInterfaces: []*armworkloadssapvirtualinstance.NetworkInterfaceResourceNames{
	// 																									{
	// 																										NetworkInterfaceName: to.Ptr("dbsrnic"),
	// 																								}},
	// 																								OSDiskName: to.Ptr("dbsrosdisk"),
	// 																								VMName: to.Ptr("dbvmsr"),
	// 																						}},
	// 																					},
	// 																					SharedStorage: &armworkloadssapvirtualinstance.SharedStorageResourceNames{
	// 																						SharedStorageAccountName: to.Ptr("storageacc"),
	// 																						SharedStorageAccountPrivateEndPointName: to.Ptr("peForxNFS"),
	// 																					},
	// 																				},
	// 																				DatabaseServer: &armworkloadssapvirtualinstance.DatabaseConfiguration{
	// 																					DatabaseType: to.Ptr(armworkloadssapvirtualinstance.SAPDatabaseTypeHANA),
	// 																					InstanceCount: to.Ptr[int64](2),
	// 																					SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/dbsubnet"),
	// 																					VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
	// 																						ImageReference: &armworkloadssapvirtualinstance.ImageReference{
	// 																							Offer: to.Ptr("RHEL-SAP"),
	// 																							Publisher: to.Ptr("RedHat"),
	// 																							SKU: to.Ptr("84sapha-gen2"),
	// 																							Version: to.Ptr("latest"),
	// 																						},
	// 																						OSProfile: &armworkloadssapvirtualinstance.OSProfile{
	// 																							AdminUsername: to.Ptr("{your-username}"),
	// 																							OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
	// 																								OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
	// 																								DisablePasswordAuthentication: to.Ptr(true),
	// 																								SSHKeyPair: &armworkloadssapvirtualinstance.SSHKeyPair{
	// 																									PublicKey: to.Ptr("abc"),
	// 																								},
	// 																							},
	// 																						},
	// 																						VMSize: to.Ptr("Standard_M32ts"),
	// 																					},
	// 																				},
	// 																				HighAvailabilityConfig: &armworkloadssapvirtualinstance.HighAvailabilityConfiguration{
	// 																					HighAvailabilityType: to.Ptr(armworkloadssapvirtualinstance.SAPHighAvailabilityTypeAvailabilitySet),
	// 																				},
	// 																			},
	// 																			OSSapConfiguration: &armworkloadssapvirtualinstance.OsSapConfiguration{
	// 																				SapFqdn: to.Ptr("xyz.test.com"),
	// 																			},
	// 																		},
	// 																		Environment: to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeProd),
	// 																		Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateUnknown),
	// 																		ManagedResourceGroupConfiguration: &armworkloadssapvirtualinstance.ManagedRGConfiguration{
	// 																			Name: to.Ptr("mrg-X00-8e17e36c-42e9-4cd5-a078-7b44883414e0"),
	// 																		},
	// 																		ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
	// 																		SapProduct: to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
	// 																		State: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStateSoftwareInstallationPending),
	// 																		Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatus("Unknown")),
	// 																	},
	// 																}
}
