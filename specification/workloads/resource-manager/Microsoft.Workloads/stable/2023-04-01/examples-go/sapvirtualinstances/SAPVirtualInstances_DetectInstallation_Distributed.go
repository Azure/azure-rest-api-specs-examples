package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/1e7b408f3323e7f5424745718fe62c7a043a2337/specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPVirtualInstances_DetectInstallation_Distributed.json
func ExampleSAPVirtualInstancesClient_BeginCreate_detectSapSoftwareInstallationOnADistributedSystem() {
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
		Location: to.Ptr("eastus2"),
		Tags: map[string]*string{
			"created by": to.Ptr("azureuser"),
		},
		Properties: &armworkloads.SAPVirtualInstanceProperties{
			Configuration: &armworkloads.DeploymentWithOSConfiguration{
				ConfigurationType: to.Ptr(armworkloads.SAPConfigurationTypeDeploymentWithOSConfig),
				AppLocation:       to.Ptr("eastus"),
				InfrastructureConfiguration: &armworkloads.ThreeTierConfiguration{
					AppResourceGroup: to.Ptr("{{resourcegrp}}"),
					DeploymentType:   to.Ptr(armworkloads.SAPDeploymentTypeThreeTier),
					ApplicationServer: &armworkloads.ApplicationServerConfiguration{
						InstanceCount: to.Ptr[int64](2),
						SubnetID:      to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app"),
						VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
							ImageReference: &armworkloads.ImageReference{
								Offer:     to.Ptr("RHEL-SAP-HA"),
								Publisher: to.Ptr("RedHat"),
								SKU:       to.Ptr("84sapha-gen2"),
								Version:   to.Ptr("latest"),
							},
							OSProfile: &armworkloads.OSProfile{
								AdminUsername: to.Ptr("azureuser"),
								OSConfiguration: &armworkloads.LinuxConfiguration{
									OSType:                        to.Ptr(armworkloads.OSTypeLinux),
									DisablePasswordAuthentication: to.Ptr(true),
									SSHKeyPair: &armworkloads.SSHKeyPair{
										PrivateKey: to.Ptr("{{privateKey}}"),
										PublicKey:  to.Ptr("{{sshkey}}"),
									},
								},
							},
							VMSize: to.Ptr("Standard_E4ds_v4"),
						},
					},
					CentralServer: &armworkloads.CentralServerConfiguration{
						InstanceCount: to.Ptr[int64](1),
						SubnetID:      to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app"),
						VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
							ImageReference: &armworkloads.ImageReference{
								Offer:     to.Ptr("RHEL-SAP-HA"),
								Publisher: to.Ptr("RedHat"),
								SKU:       to.Ptr("84sapha-gen2"),
								Version:   to.Ptr("latest"),
							},
							OSProfile: &armworkloads.OSProfile{
								AdminUsername: to.Ptr("azureuser"),
								OSConfiguration: &armworkloads.LinuxConfiguration{
									OSType:                        to.Ptr(armworkloads.OSTypeLinux),
									DisablePasswordAuthentication: to.Ptr(true),
									SSHKeyPair: &armworkloads.SSHKeyPair{
										PrivateKey: to.Ptr("{{privateKey}}"),
										PublicKey:  to.Ptr("{{sshkey}}"),
									},
								},
							},
							VMSize: to.Ptr("Standard_E4ds_v4"),
						},
					},
					DatabaseServer: &armworkloads.DatabaseConfiguration{
						InstanceCount: to.Ptr[int64](1),
						SubnetID:      to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app"),
						VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
							ImageReference: &armworkloads.ImageReference{
								Offer:     to.Ptr("RHEL-SAP-HA"),
								Publisher: to.Ptr("RedHat"),
								SKU:       to.Ptr("84sapha-gen2"),
								Version:   to.Ptr("latest"),
							},
							OSProfile: &armworkloads.OSProfile{
								AdminUsername: to.Ptr("azureuser"),
								OSConfiguration: &armworkloads.LinuxConfiguration{
									OSType:                        to.Ptr(armworkloads.OSTypeLinux),
									DisablePasswordAuthentication: to.Ptr(true),
									SSHKeyPair: &armworkloads.SSHKeyPair{
										PrivateKey: to.Ptr("{{privateKey}}"),
										PublicKey:  to.Ptr("{{sshkey}}"),
									},
								},
							},
							VMSize: to.Ptr("Standard_M32ts"),
						},
					},
					NetworkConfiguration: &armworkloads.NetworkConfiguration{
						IsSecondaryIPEnabled: to.Ptr(true),
					},
				},
				OSSapConfiguration: &armworkloads.OsSapConfiguration{
					SapFqdn: to.Ptr("sap.bpaas.com"),
				},
				SoftwareConfiguration: &armworkloads.ExternalInstallationSoftwareConfiguration{
					SoftwareInstallationType: to.Ptr(armworkloads.SAPSoftwareInstallationTypeExternal),
					CentralServerVMID:        to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/sapq20scsvm0"),
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
	// 	Type: to.Ptr("microsoft.workloads/sapvirtualinstances"),
	// 	ID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00"),
	// 	SystemData: &armworkloads.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-18T18:25:55.240Z"); return t}()),
	// 		CreatedBy: to.Ptr("f1bc9460-9d57-4c16-99a8-5a19378a3a93"),
	// 		CreatedByType: to.Ptr(armworkloads.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-18T18:50:59.194Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("f1bc9460-9d57-4c16-99a8-5a19378a3a93"),
	// 		LastModifiedByType: to.Ptr(armworkloads.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("eastus2"),
	// 	Tags: map[string]*string{
	// 		"created by": to.Ptr("azureuser"),
	// 	},
	// 	Properties: &armworkloads.SAPVirtualInstanceProperties{
	// 		Configuration: &armworkloads.DeploymentWithOSConfiguration{
	// 			ConfigurationType: to.Ptr(armworkloads.SAPConfigurationTypeDeploymentWithOSConfig),
	// 			AppLocation: to.Ptr("eastus"),
	// 			InfrastructureConfiguration: &armworkloads.ThreeTierConfiguration{
	// 				AppResourceGroup: to.Ptr("aa-rg"),
	// 				DeploymentType: to.Ptr(armworkloads.SAPDeploymentTypeThreeTier),
	// 				ApplicationServer: &armworkloads.ApplicationServerConfiguration{
	// 					InstanceCount: to.Ptr[int64](2),
	// 					SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app"),
	// 					VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
	// 						ImageReference: &armworkloads.ImageReference{
	// 							Offer: to.Ptr("RHEL-SAP-HA"),
	// 							Publisher: to.Ptr("RedHat"),
	// 							SKU: to.Ptr("84sapha-gen2"),
	// 							Version: to.Ptr("latest"),
	// 						},
	// 						OSProfile: &armworkloads.OSProfile{
	// 							AdminUsername: to.Ptr("azureuser"),
	// 							OSConfiguration: &armworkloads.LinuxConfiguration{
	// 								OSType: to.Ptr(armworkloads.OSTypeLinux),
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								SSHKeyPair: &armworkloads.SSHKeyPair{
	// 									PublicKey: to.Ptr("ssh-rsa AAA generated-by-azure"),
	// 								},
	// 							},
	// 						},
	// 						VMSize: to.Ptr("Standard_E4ds_v4"),
	// 					},
	// 				},
	// 				CentralServer: &armworkloads.CentralServerConfiguration{
	// 					InstanceCount: to.Ptr[int64](1),
	// 					SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app"),
	// 					VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
	// 						ImageReference: &armworkloads.ImageReference{
	// 							Offer: to.Ptr("RHEL-SAP-HA"),
	// 							Publisher: to.Ptr("RedHat"),
	// 							SKU: to.Ptr("84sapha-gen2"),
	// 							Version: to.Ptr("latest"),
	// 						},
	// 						OSProfile: &armworkloads.OSProfile{
	// 							AdminUsername: to.Ptr("azureuser"),
	// 							OSConfiguration: &armworkloads.LinuxConfiguration{
	// 								OSType: to.Ptr(armworkloads.OSTypeLinux),
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								SSHKeyPair: &armworkloads.SSHKeyPair{
	// 									PublicKey: to.Ptr("ssh-rsa AAA generated-by-azure"),
	// 								},
	// 							},
	// 						},
	// 						VMSize: to.Ptr("Standard_E4ds_v4"),
	// 					},
	// 				},
	// 				DatabaseServer: &armworkloads.DatabaseConfiguration{
	// 					InstanceCount: to.Ptr[int64](1),
	// 					SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app"),
	// 					VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
	// 						ImageReference: &armworkloads.ImageReference{
	// 							Offer: to.Ptr("RHEL-SAP-HA"),
	// 							Publisher: to.Ptr("RedHat"),
	// 							SKU: to.Ptr("84sapha-gen2"),
	// 							Version: to.Ptr("latest"),
	// 						},
	// 						OSProfile: &armworkloads.OSProfile{
	// 							AdminUsername: to.Ptr("azureuser"),
	// 							OSConfiguration: &armworkloads.LinuxConfiguration{
	// 								OSType: to.Ptr(armworkloads.OSTypeLinux),
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								SSHKeyPair: &armworkloads.SSHKeyPair{
	// 									PublicKey: to.Ptr("ssh-rsa AAA generated-by-azure"),
	// 								},
	// 							},
	// 						},
	// 						VMSize: to.Ptr("Standard_M32ts"),
	// 					},
	// 				},
	// 				NetworkConfiguration: &armworkloads.NetworkConfiguration{
	// 					IsSecondaryIPEnabled: to.Ptr(true),
	// 				},
	// 			},
	// 			OSSapConfiguration: &armworkloads.OsSapConfiguration{
	// 				SapFqdn: to.Ptr("sap.bpaas.com"),
	// 			},
	// 			SoftwareConfiguration: &armworkloads.ExternalInstallationSoftwareConfiguration{
	// 				SoftwareInstallationType: to.Ptr(armworkloads.SAPSoftwareInstallationTypeExternal),
	// 				CentralServerVMID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/sapq20scsvm0"),
	// 			},
	// 		},
	// 		Environment: to.Ptr(armworkloads.SAPEnvironmentTypeProd),
	// 		ProvisioningState: to.Ptr(armworkloads.SapVirtualInstanceProvisioningStateSucceeded),
	// 		SapProduct: to.Ptr(armworkloads.SAPProductTypeS4HANA),
	// 		State: to.Ptr(armworkloads.SAPVirtualInstanceStateRegistrationComplete),
	// 	},
	// }
}
