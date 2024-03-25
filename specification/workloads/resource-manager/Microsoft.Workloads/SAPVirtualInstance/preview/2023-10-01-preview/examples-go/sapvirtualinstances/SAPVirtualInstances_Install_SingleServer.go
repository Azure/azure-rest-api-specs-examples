package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Install_SingleServer.json
func ExampleSAPVirtualInstancesClient_BeginCreate_installSapSoftwareOnSingleServerSystem() {
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
		Location: to.Ptr("eastus2"),
		Tags:     map[string]*string{},
		Properties: &armworkloadssapvirtualinstance.SAPVirtualInstanceProperties{
			Configuration: &armworkloadssapvirtualinstance.DeploymentWithOSConfiguration{
				ConfigurationType: to.Ptr(armworkloadssapvirtualinstance.SAPConfigurationTypeDeploymentWithOSConfig),
				AppLocation:       to.Ptr("eastus"),
				InfrastructureConfiguration: &armworkloadssapvirtualinstance.SingleServerConfiguration{
					AppResourceGroup: to.Ptr("test-rg"),
					DeploymentType:   to.Ptr(armworkloadssapvirtualinstance.SAPDeploymentTypeSingleServer),
					SubnetID:         to.Ptr("/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/testsubnet"),
					VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
						ImageReference: &armworkloadssapvirtualinstance.ImageReference{
							Offer:     to.Ptr("SLES-SAP"),
							Publisher: to.Ptr("SUSE"),
							SKU:       to.Ptr("12-sp4-gen2"),
							Version:   to.Ptr("2022.02.01"),
						},
						OSProfile: &armworkloadssapvirtualinstance.OSProfile{
							AdminUsername: to.Ptr("azureappadmin"),
							OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
								OSType:                        to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
								DisablePasswordAuthentication: to.Ptr(true),
								SSHKeyPair: &armworkloadssapvirtualinstance.SSHKeyPair{
									PrivateKey: to.Ptr("{{privateKey}}"),
									PublicKey:  to.Ptr("{{sshkey}}"),
								},
							},
						},
						VMSize: to.Ptr("Standard_E32ds_v4"),
					},
				},
				OSSapConfiguration: &armworkloadssapvirtualinstance.OsSapConfiguration{
					SapFqdn: to.Ptr("sap.bpaas.com"),
				},
				SoftwareConfiguration: &armworkloadssapvirtualinstance.SAPInstallWithoutOSConfigSoftwareConfiguration{
					SoftwareInstallationType: to.Ptr(armworkloadssapvirtualinstance.SAPSoftwareInstallationTypeSAPInstallWithoutOSConfig),
					BomURL:                   to.Ptr("https://teststorageaccount.blob.core.windows.net/sapbits/sapfiles/boms/S41909SPS03_v0011ms/S41909SPS03_v0011ms.yaml"),
					SapBitsStorageAccountID:  to.Ptr("/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/teststorageaccount"),
					SoftwareVersion:          to.Ptr("SAP S/4HANA 1909 SPS 03"),
				},
			},
			Environment: to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeNonProd),
			SapProduct:  to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
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
	// 	Type: to.Ptr("microsoft.workloads/sapvirtualinstances"),
	// 	ID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/aa-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00"),
	// 	SystemData: &armworkloadssapvirtualinstance.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-18T18:25:55.240Z"); return t}()),
	// 		CreatedBy: to.Ptr("f1bc9460-9d57-4c16-99a8-5a19378a3a93"),
	// 		CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-18T18:50:59.194Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("f1bc9460-9d57-4c16-99a8-5a19378a3a93"),
	// 		LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("eastus2"),
	// 	Tags: map[string]*string{
	// 		"created by": to.Ptr("azureuser"),
	// 	},
	// 	Properties: &armworkloadssapvirtualinstance.SAPVirtualInstanceProperties{
	// 		Configuration: &armworkloadssapvirtualinstance.DeploymentWithOSConfiguration{
	// 			ConfigurationType: to.Ptr(armworkloadssapvirtualinstance.SAPConfigurationTypeDeploymentWithOSConfig),
	// 			AppLocation: to.Ptr("eastus"),
	// 			InfrastructureConfiguration: &armworkloadssapvirtualinstance.ThreeTierConfiguration{
	// 				AppResourceGroup: to.Ptr("aa-rg"),
	// 				DeploymentType: to.Ptr(armworkloadssapvirtualinstance.SAPDeploymentTypeThreeTier),
	// 				ApplicationServer: &armworkloadssapvirtualinstance.ApplicationServerConfiguration{
	// 					InstanceCount: to.Ptr[int64](2),
	// 					SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app"),
	// 					VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
	// 						ImageReference: &armworkloadssapvirtualinstance.ImageReference{
	// 							Offer: to.Ptr("RHEL-SAP-HA"),
	// 							Publisher: to.Ptr("RedHat"),
	// 							SKU: to.Ptr("84sapha-gen2"),
	// 							Version: to.Ptr("latest"),
	// 						},
	// 						OSProfile: &armworkloadssapvirtualinstance.OSProfile{
	// 							AdminUsername: to.Ptr("azureuser"),
	// 							OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
	// 								OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								SSHKeyPair: &armworkloadssapvirtualinstance.SSHKeyPair{
	// 									PublicKey: to.Ptr("ssh-rsa AAA generated-by-azure"),
	// 								},
	// 							},
	// 						},
	// 						VMSize: to.Ptr("Standard_E4ds_v4"),
	// 					},
	// 				},
	// 				CentralServer: &armworkloadssapvirtualinstance.CentralServerConfiguration{
	// 					InstanceCount: to.Ptr[int64](1),
	// 					SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app"),
	// 					VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
	// 						ImageReference: &armworkloadssapvirtualinstance.ImageReference{
	// 							Offer: to.Ptr("RHEL-SAP-HA"),
	// 							Publisher: to.Ptr("RedHat"),
	// 							SKU: to.Ptr("84sapha-gen2"),
	// 							Version: to.Ptr("latest"),
	// 						},
	// 						OSProfile: &armworkloadssapvirtualinstance.OSProfile{
	// 							AdminUsername: to.Ptr("azureuser"),
	// 							OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
	// 								OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								SSHKeyPair: &armworkloadssapvirtualinstance.SSHKeyPair{
	// 									PublicKey: to.Ptr("ssh-rsa AAA generated-by-azure"),
	// 								},
	// 							},
	// 						},
	// 						VMSize: to.Ptr("Standard_E4ds_v4"),
	// 					},
	// 				},
	// 				DatabaseServer: &armworkloadssapvirtualinstance.DatabaseConfiguration{
	// 					InstanceCount: to.Ptr[int64](1),
	// 					SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app"),
	// 					VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
	// 						ImageReference: &armworkloadssapvirtualinstance.ImageReference{
	// 							Offer: to.Ptr("RHEL-SAP-HA"),
	// 							Publisher: to.Ptr("RedHat"),
	// 							SKU: to.Ptr("84sapha-gen2"),
	// 							Version: to.Ptr("latest"),
	// 						},
	// 						OSProfile: &armworkloadssapvirtualinstance.OSProfile{
	// 							AdminUsername: to.Ptr("azureuser"),
	// 							OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
	// 								OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								SSHKeyPair: &armworkloadssapvirtualinstance.SSHKeyPair{
	// 									PublicKey: to.Ptr("ssh-rsa AAA generated-by-azure"),
	// 								},
	// 							},
	// 						},
	// 						VMSize: to.Ptr("Standard_M32ts"),
	// 					},
	// 				},
	// 				NetworkConfiguration: &armworkloadssapvirtualinstance.NetworkConfiguration{
	// 					IsSecondaryIPEnabled: to.Ptr(true),
	// 				},
	// 			},
	// 			OSSapConfiguration: &armworkloadssapvirtualinstance.OsSapConfiguration{
	// 				SapFqdn: to.Ptr("sap.bpaas.com"),
	// 			},
	// 			SoftwareConfiguration: &armworkloadssapvirtualinstance.SAPInstallWithoutOSConfigSoftwareConfiguration{
	// 				SoftwareInstallationType: to.Ptr(armworkloadssapvirtualinstance.SAPSoftwareInstallationTypeSAPInstallWithoutOSConfig),
	// 				BomURL: to.Ptr("https://teststorageaccount.blob.core.windows.net/sapbits/sapfiles/boms/S41909SPS03_v0011ms/S41909SPS03_v0011ms.yaml"),
	// 				SapBitsStorageAccountID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/teststorageaccount"),
	// 				SoftwareVersion: to.Ptr("SAP S/4HANA 1909 SPS 03"),
	// 			},
	// 		},
	// 		Environment: to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeProd),
	// 		ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
	// 		SapProduct: to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
	// 		State: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStateSoftwareInstallationInProgress),
	// 	},
	// }
}
