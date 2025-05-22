package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: 2024-09-01/SapVirtualInstances_CreateInfraOsComputeGalleryImage.json
func ExampleSAPVirtualInstancesClient_BeginCreate_createInfrastructureWithOSConfigurationWithAzureComputeGalleryImage() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("8e17e36c-42e9-4cd5-a078-7b44883414e0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSAPVirtualInstancesClient().BeginCreate(ctx, "test-rg", "X00", armworkloadssapvirtualinstance.SAPVirtualInstance{
		Location: to.Ptr("westcentralus"),
		Properties: &armworkloadssapvirtualinstance.SAPVirtualInstanceProperties{
			Configuration: &armworkloadssapvirtualinstance.DeploymentConfiguration{
				AppLocation:       to.Ptr("eastus"),
				ConfigurationType: to.Ptr(armworkloadssapvirtualinstance.SAPConfigurationTypeDeployment),
				InfrastructureConfiguration: &armworkloadssapvirtualinstance.ThreeTierConfiguration{
					AppResourceGroup: to.Ptr("X00-RG"),
					ApplicationServer: &armworkloadssapvirtualinstance.ApplicationServerConfiguration{
						InstanceCount: to.Ptr[int64](6),
						SubnetID:      to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
						VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
							ImageReference: &armworkloadssapvirtualinstance.ImageReference{
								ID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Compute/galleries/testgallery/images/rhelimagetest/versions/0.0.1"),
							},
							OSProfile: &armworkloadssapvirtualinstance.OSProfile{
								AdminUsername: to.Ptr("{your-username}"),
								OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
									DisablePasswordAuthentication: to.Ptr(true),
									OSType:                        to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
									SSH: &armworkloadssapvirtualinstance.SSHConfiguration{
										PublicKeys: []*armworkloadssapvirtualinstance.SSHPublicKey{
											{
												KeyData: to.Ptr("ssh-rsa public key"),
											},
										},
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
								ID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Compute/galleries/testgallery/images/rhelimagetest/versions/0.0.1"),
							},
							OSProfile: &armworkloadssapvirtualinstance.OSProfile{
								AdminUsername: to.Ptr("{your-username}"),
								OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
									DisablePasswordAuthentication: to.Ptr(true),
									OSType:                        to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
									SSH: &armworkloadssapvirtualinstance.SSHConfiguration{
										PublicKeys: []*armworkloadssapvirtualinstance.SSHPublicKey{
											{
												KeyData: to.Ptr("ssh-rsa public key"),
											},
										},
									},
								},
							},
							VMSize: to.Ptr("Standard_E16ds_v4"),
						},
					},
					DatabaseServer: &armworkloadssapvirtualinstance.DatabaseConfiguration{
						DatabaseType:  to.Ptr(armworkloadssapvirtualinstance.SAPDatabaseTypeHANA),
						InstanceCount: to.Ptr[int64](2),
						SubnetID:      to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/dbsubnet"),
						VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
							ImageReference: &armworkloadssapvirtualinstance.ImageReference{
								ID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Compute/galleries/testgallery/images/rhelimagetest/versions/0.0.1"),
							},
							OSProfile: &armworkloadssapvirtualinstance.OSProfile{
								AdminUsername: to.Ptr("{your-username}"),
								OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
									DisablePasswordAuthentication: to.Ptr(true),
									OSType:                        to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
									SSH: &armworkloadssapvirtualinstance.SSHConfiguration{
										PublicKeys: []*armworkloadssapvirtualinstance.SSHPublicKey{
											{
												KeyData: to.Ptr("ssh-rsa public key"),
											},
										},
									},
								},
							},
							VMSize: to.Ptr("Standard_M32ts"),
						},
					},
					DeploymentType: to.Ptr(armworkloadssapvirtualinstance.SAPDeploymentTypeThreeTier),
					HighAvailabilityConfig: &armworkloadssapvirtualinstance.HighAvailabilityConfiguration{
						HighAvailabilityType: to.Ptr(armworkloadssapvirtualinstance.SAPHighAvailabilityTypeAvailabilityZone),
					},
				},
			},
			Environment: to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeProd),
			SapProduct:  to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
		},
		Tags: map[string]*string{},
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
	// res = armworkloadssapvirtualinstance.SAPVirtualInstancesClientCreateResponse{
	// 	SAPVirtualInstance: &armworkloadssapvirtualinstance.SAPVirtualInstance{
	// 		Name: to.Ptr("X00"),
	// 		Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances"),
	// 		ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00"),
	// 		Location: to.Ptr("westcentralus"),
	// 		Properties: &armworkloadssapvirtualinstance.SAPVirtualInstanceProperties{
	// 			Configuration: &armworkloadssapvirtualinstance.DeploymentConfiguration{
	// 				AppLocation: to.Ptr("eastus"),
	// 				ConfigurationType: to.Ptr(armworkloadssapvirtualinstance.SAPConfigurationTypeDeployment),
	// 				InfrastructureConfiguration: &armworkloadssapvirtualinstance.ThreeTierConfiguration{
	// 					AppResourceGroup: to.Ptr("X00-RG"),
	// 					ApplicationServer: &armworkloadssapvirtualinstance.ApplicationServerConfiguration{
	// 						InstanceCount: to.Ptr[int64](6),
	// 						SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
	// 						VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
	// 							ImageReference: &armworkloadssapvirtualinstance.ImageReference{
	// 								ID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Compute/galleries/testgallery/images/rhelimagetest/versions/0.0.1"),
	// 							},
	// 							OSProfile: &armworkloadssapvirtualinstance.OSProfile{
	// 								AdminUsername: to.Ptr("{your-username}"),
	// 								OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
	// 									DisablePasswordAuthentication: to.Ptr(true),
	// 									OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
	// 									SSH: &armworkloadssapvirtualinstance.SSHConfiguration{
	// 										PublicKeys: []*armworkloadssapvirtualinstance.SSHPublicKey{
	// 											{
	// 												KeyData: to.Ptr("ssh-rsa public key"),
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 							VMSize: to.Ptr("Standard_E32ds_v4"),
	// 						},
	// 					},
	// 					CentralServer: &armworkloadssapvirtualinstance.CentralServerConfiguration{
	// 						InstanceCount: to.Ptr[int64](2),
	// 						SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
	// 						VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
	// 							ImageReference: &armworkloadssapvirtualinstance.ImageReference{
	// 								ID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Compute/galleries/testgallery/images/rhelimagetest/versions/0.0.1"),
	// 							},
	// 							OSProfile: &armworkloadssapvirtualinstance.OSProfile{
	// 								AdminUsername: to.Ptr("{your-username}"),
	// 								OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
	// 									DisablePasswordAuthentication: to.Ptr(true),
	// 									OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
	// 									SSH: &armworkloadssapvirtualinstance.SSHConfiguration{
	// 										PublicKeys: []*armworkloadssapvirtualinstance.SSHPublicKey{
	// 											{
	// 												KeyData: to.Ptr("ssh-rsa public key"),
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 							VMSize: to.Ptr("Standard_E16ds_v4"),
	// 						},
	// 					},
	// 					DatabaseServer: &armworkloadssapvirtualinstance.DatabaseConfiguration{
	// 						DatabaseType: to.Ptr(armworkloadssapvirtualinstance.SAPDatabaseTypeHANA),
	// 						InstanceCount: to.Ptr[int64](2),
	// 						SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/dbsubnet"),
	// 						VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
	// 							ImageReference: &armworkloadssapvirtualinstance.ImageReference{
	// 								ID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Compute/galleries/testgallery/images/rhelimagetest/versions/0.0.1"),
	// 							},
	// 							OSProfile: &armworkloadssapvirtualinstance.OSProfile{
	// 								AdminUsername: to.Ptr("{your-username}"),
	// 								OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
	// 									DisablePasswordAuthentication: to.Ptr(true),
	// 									OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
	// 									SSH: &armworkloadssapvirtualinstance.SSHConfiguration{
	// 										PublicKeys: []*armworkloadssapvirtualinstance.SSHPublicKey{
	// 											{
	// 												KeyData: to.Ptr("ssh-rsa public key"),
	// 											},
	// 										},
	// 									},
	// 								},
	// 							},
	// 							VMSize: to.Ptr("Standard_M32ts"),
	// 						},
	// 					},
	// 					DeploymentType: to.Ptr(armworkloadssapvirtualinstance.SAPDeploymentTypeThreeTier),
	// 					HighAvailabilityConfig: &armworkloadssapvirtualinstance.HighAvailabilityConfiguration{
	// 						HighAvailabilityType: to.Ptr(armworkloadssapvirtualinstance.SAPHighAvailabilityTypeAvailabilityZone),
	// 					},
	// 				},
	// 			},
	// 			Environment: to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeProd),
	// 			Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateUnknown),
	// 			ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
	// 			SapProduct: to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
	// 			State: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStateSoftwareInstallationPending),
	// 			Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatus("Unknown")),
	// 		},
	// 		SystemData: &armworkloadssapvirtualinstance.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 			CreatedBy: to.Ptr("user@xyz.com"),
	// 			CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user@xyz.com"),
	// 			LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 	},
	// }
}
