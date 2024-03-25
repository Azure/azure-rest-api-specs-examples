package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Create_DiskDetails_SingleServer.json
func ExampleSAPVirtualInstancesClient_BeginCreate_createInfrastructureWithDiskAndOsConfigurationsForSingleServerSystemRecommended() {
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
				InfrastructureConfiguration: &armworkloadssapvirtualinstance.SingleServerConfiguration{
					AppResourceGroup: to.Ptr("X00-RG"),
					DeploymentType:   to.Ptr(armworkloadssapvirtualinstance.SAPDeploymentTypeSingleServer),
					DatabaseType:     to.Ptr(armworkloadssapvirtualinstance.SAPDatabaseTypeHANA),
					DbDiskConfiguration: &armworkloadssapvirtualinstance.DiskConfiguration{
						DiskVolumeConfigurations: map[string]*armworkloadssapvirtualinstance.DiskVolumeConfiguration{
							"backup": {
								Count:  to.Ptr[int64](2),
								SizeGB: to.Ptr[int64](256),
								SKU: &armworkloadssapvirtualinstance.DiskSKU{
									Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
								},
							},
							"hana/data": {
								Count:  to.Ptr[int64](4),
								SizeGB: to.Ptr[int64](128),
								SKU: &armworkloadssapvirtualinstance.DiskSKU{
									Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
								},
							},
							"hana/log": {
								Count:  to.Ptr[int64](3),
								SizeGB: to.Ptr[int64](128),
								SKU: &armworkloadssapvirtualinstance.DiskSKU{
									Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
								},
							},
							"hana/shared": {
								Count:  to.Ptr[int64](1),
								SizeGB: to.Ptr[int64](256),
								SKU: &armworkloadssapvirtualinstance.DiskSKU{
									Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
								},
							},
							"os": {
								Count:  to.Ptr[int64](1),
								SizeGB: to.Ptr[int64](64),
								SKU: &armworkloadssapvirtualinstance.DiskSKU{
									Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
								},
							},
							"usr/sap": {
								Count:  to.Ptr[int64](1),
								SizeGB: to.Ptr[int64](128),
								SKU: &armworkloadssapvirtualinstance.DiskSKU{
									Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
								},
							},
						},
					},
					NetworkConfiguration: &armworkloadssapvirtualinstance.NetworkConfiguration{
						IsSecondaryIPEnabled: to.Ptr(true),
					},
					SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/dindurkhya-e2etesting/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
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
				OSSapConfiguration: &armworkloadssapvirtualinstance.OsSapConfiguration{
					SapFqdn: to.Ptr("xyz.test.com"),
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
	// 			InfrastructureConfiguration: &armworkloadssapvirtualinstance.SingleServerConfiguration{
	// 				AppResourceGroup: to.Ptr("X00-RG"),
	// 				DeploymentType: to.Ptr(armworkloadssapvirtualinstance.SAPDeploymentTypeSingleServer),
	// 				DatabaseType: to.Ptr(armworkloadssapvirtualinstance.SAPDatabaseTypeHANA),
	// 				DbDiskConfiguration: &armworkloadssapvirtualinstance.DiskConfiguration{
	// 					DiskVolumeConfigurations: map[string]*armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 						"backup": &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 							Count: to.Ptr[int64](2),
	// 							SizeGB: to.Ptr[int64](256),
	// 							SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 								Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 							},
	// 						},
	// 						"hana/data": &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 							Count: to.Ptr[int64](4),
	// 							SizeGB: to.Ptr[int64](128),
	// 							SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 								Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 							},
	// 						},
	// 						"hana/log": &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 							Count: to.Ptr[int64](3),
	// 							SizeGB: to.Ptr[int64](128),
	// 							SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 								Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 							},
	// 						},
	// 						"hana/shared": &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 							Count: to.Ptr[int64](1),
	// 							SizeGB: to.Ptr[int64](256),
	// 							SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 								Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 							},
	// 						},
	// 						"os": &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 							Count: to.Ptr[int64](1),
	// 							SizeGB: to.Ptr[int64](64),
	// 							SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 								Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNameStandardSSDLRS),
	// 							},
	// 						},
	// 						"usr/sap": &armworkloadssapvirtualinstance.DiskVolumeConfiguration{
	// 							Count: to.Ptr[int64](1),
	// 							SizeGB: to.Ptr[int64](128),
	// 							SKU: &armworkloadssapvirtualinstance.DiskSKU{
	// 								Name: to.Ptr(armworkloadssapvirtualinstance.DiskSKUNamePremiumLRS),
	// 							},
	// 						},
	// 					},
	// 				},
	// 				NetworkConfiguration: &armworkloadssapvirtualinstance.NetworkConfiguration{
	// 					IsSecondaryIPEnabled: to.Ptr(true),
	// 				},
	// 				SubnetID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/dindurkhya-e2etesting/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet"),
	// 				VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
	// 					ImageReference: &armworkloadssapvirtualinstance.ImageReference{
	// 						Offer: to.Ptr("RHEL-SAP"),
	// 						Publisher: to.Ptr("RedHat"),
	// 						SKU: to.Ptr("84sapha-gen2"),
	// 						Version: to.Ptr("latest"),
	// 					},
	// 					OSProfile: &armworkloadssapvirtualinstance.OSProfile{
	// 						AdminUsername: to.Ptr("{your-username}"),
	// 						OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
	// 							OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
	// 							DisablePasswordAuthentication: to.Ptr(true),
	// 							SSHKeyPair: &armworkloadssapvirtualinstance.SSHKeyPair{
	// 								PublicKey: to.Ptr("abc"),
	// 							},
	// 						},
	// 					},
	// 					VMSize: to.Ptr("Standard_E32ds_v4"),
	// 				},
	// 			},
	// 			OSSapConfiguration: &armworkloadssapvirtualinstance.OsSapConfiguration{
	// 				SapFqdn: to.Ptr("xyz.test.com"),
	// 			},
	// 		},
	// 		Environment: to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeNonProd),
	// 		Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateUnknown),
	// 		ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
	// 		SapProduct: to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
	// 		State: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStateSoftwareInstallationPending),
	// 		Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatus("Unknown")),
	// 	},
	// }
}
