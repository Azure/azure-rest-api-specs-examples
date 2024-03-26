package armworkloadssapvirtualinstance_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b1e318cbfd2e239db54c80af5e6aea7fdf658851/specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_ListBySubscription.json
func ExampleSAPVirtualInstancesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armworkloadssapvirtualinstance.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSAPVirtualInstancesClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SAPVirtualInstanceList = armworkloadssapvirtualinstance.SAPVirtualInstanceList{
		// 	Value: []*armworkloadssapvirtualinstance.SAPVirtualInstance{
		// 		{
		// 			Name: to.Ptr("X00"),
		// 			Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances"),
		// 			ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X00"),
		// 			SystemData: &armworkloadssapvirtualinstance.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@xyz.com"),
		// 				CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@xyz.com"),
		// 				LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armworkloadssapvirtualinstance.SAPVirtualInstanceProperties{
		// 				Configuration: &armworkloadssapvirtualinstance.DeploymentConfiguration{
		// 					ConfigurationType: to.Ptr(armworkloadssapvirtualinstance.SAPConfigurationTypeDeployment),
		// 					AppLocation: to.Ptr("eastus"),
		// 					InfrastructureConfiguration: &armworkloadssapvirtualinstance.ThreeTierConfiguration{
		// 						AppResourceGroup: to.Ptr("X00-RG"),
		// 						DeploymentType: to.Ptr(armworkloadssapvirtualinstance.SAPDeploymentTypeThreeTier),
		// 						ApplicationServer: &armworkloadssapvirtualinstance.ApplicationServerConfiguration{
		// 							InstanceCount: to.Ptr[int64](10),
		// 							SubnetID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/vnet1/subnets/subnetId3"),
		// 							VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
		// 								ImageReference: &armworkloadssapvirtualinstance.ImageReference{
		// 									Offer: to.Ptr("RHEL-SAP"),
		// 									Publisher: to.Ptr("RedHat"),
		// 									SKU: to.Ptr("84sapha-gen2"),
		// 									Version: to.Ptr("latest"),
		// 								},
		// 								OSProfile: &armworkloadssapvirtualinstance.OSProfile{
		// 									AdminUsername: to.Ptr("{your-username}"),
		// 									OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
		// 										OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
		// 										DisablePasswordAuthentication: to.Ptr(true),
		// 										SSH: &armworkloadssapvirtualinstance.SSHConfiguration{
		// 											PublicKeys: []*armworkloadssapvirtualinstance.SSHPublicKey{
		// 												{
		// 													KeyData: to.Ptr("ssh-rsa public key"),
		// 											}},
		// 										},
		// 									},
		// 								},
		// 								VMSize: to.Ptr("Standard_D8s_v3"),
		// 							},
		// 						},
		// 						CentralServer: &armworkloadssapvirtualinstance.CentralServerConfiguration{
		// 							InstanceCount: to.Ptr[int64](1),
		// 							SubnetID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/vnet1/subnets/subnetId1"),
		// 							VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
		// 								ImageReference: &armworkloadssapvirtualinstance.ImageReference{
		// 									Offer: to.Ptr("RHEL-SAP"),
		// 									Publisher: to.Ptr("RedHat"),
		// 									SKU: to.Ptr("84sapha-gen2"),
		// 									Version: to.Ptr("latest"),
		// 								},
		// 								OSProfile: &armworkloadssapvirtualinstance.OSProfile{
		// 									AdminUsername: to.Ptr("{your-username}"),
		// 									OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
		// 										OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
		// 										DisablePasswordAuthentication: to.Ptr(true),
		// 										SSH: &armworkloadssapvirtualinstance.SSHConfiguration{
		// 											PublicKeys: []*armworkloadssapvirtualinstance.SSHPublicKey{
		// 												{
		// 													KeyData: to.Ptr("ssh-rsa public key"),
		// 											}},
		// 										},
		// 									},
		// 								},
		// 								VMSize: to.Ptr("Standard_D8s_v3"),
		// 							},
		// 						},
		// 						DatabaseServer: &armworkloadssapvirtualinstance.DatabaseConfiguration{
		// 							DatabaseType: to.Ptr(armworkloadssapvirtualinstance.SAPDatabaseTypeHANA),
		// 							InstanceCount: to.Ptr[int64](1),
		// 							SubnetID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/vnet1/subnets/subnetId2"),
		// 							VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
		// 								ImageReference: &armworkloadssapvirtualinstance.ImageReference{
		// 									Offer: to.Ptr("RHEL-SAP"),
		// 									Publisher: to.Ptr("RedHat"),
		// 									SKU: to.Ptr("84sapha-gen2"),
		// 									Version: to.Ptr("latest"),
		// 								},
		// 								OSProfile: &armworkloadssapvirtualinstance.OSProfile{
		// 									AdminUsername: to.Ptr("{your-username}"),
		// 									OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
		// 										OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
		// 										DisablePasswordAuthentication: to.Ptr(true),
		// 										SSH: &armworkloadssapvirtualinstance.SSHConfiguration{
		// 											PublicKeys: []*armworkloadssapvirtualinstance.SSHPublicKey{
		// 												{
		// 													KeyData: to.Ptr("ssh-rsa public key"),
		// 											}},
		// 										},
		// 									},
		// 								},
		// 								VMSize: to.Ptr("Standard_D8s_v3"),
		// 							},
		// 						},
		// 					},
		// 					SoftwareConfiguration: &armworkloadssapvirtualinstance.ServiceInitiatedSoftwareConfiguration{
		// 						SoftwareInstallationType: to.Ptr(armworkloadssapvirtualinstance.SAPSoftwareInstallationTypeServiceInitiated),
		// 						BomURL: to.Ptr("https://myaccount.blob.core.windows.net/mycontainer/myblob"),
		// 						SapBitsStorageAccountID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/ajgupt-test/providers/Microsoft.Storage/storageAccounts/ajguptsavnet"),
		// 						SapFqdn: to.Ptr("sapsystem.contoso.net"),
		// 						SoftwareVersion: to.Ptr("SAP S/4HANA 2020 FPS01"),
		// 					},
		// 				},
		// 				Environment: to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeProd),
		// 				Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
		// 				ManagedResourceGroupConfiguration: &armworkloadssapvirtualinstance.ManagedRGConfiguration{
		// 					Name: to.Ptr("mrg-x00-6d875e77-e412-4d7d-9af4-8895278b4443"),
		// 				},
		// 				ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
		// 				SapProduct: to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
		// 				State: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStateRegistrationComplete),
		// 				Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatusRunning),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("X01"),
		// 			Type: to.Ptr("Microsoft.Workloads/sapVirtualInstances"),
		// 			ID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Workloads/sapVirtualInstances/X01"),
		// 			SystemData: &armworkloadssapvirtualinstance.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
		// 				CreatedBy: to.Ptr("user@xyz.com"),
		// 				CreatedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-19T15:10:46.196Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user@xyz.com"),
		// 				LastModifiedByType: to.Ptr(armworkloadssapvirtualinstance.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armworkloadssapvirtualinstance.SAPVirtualInstanceProperties{
		// 				Configuration: &armworkloadssapvirtualinstance.DeploymentConfiguration{
		// 					ConfigurationType: to.Ptr(armworkloadssapvirtualinstance.SAPConfigurationTypeDeployment),
		// 					AppLocation: to.Ptr("eastus"),
		// 					InfrastructureConfiguration: &armworkloadssapvirtualinstance.ThreeTierConfiguration{
		// 						AppResourceGroup: to.Ptr("X00-RG"),
		// 						DeploymentType: to.Ptr(armworkloadssapvirtualinstance.SAPDeploymentTypeThreeTier),
		// 						ApplicationServer: &armworkloadssapvirtualinstance.ApplicationServerConfiguration{
		// 							InstanceCount: to.Ptr[int64](10),
		// 							SubnetID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/vnet1/subnets/subnetId3"),
		// 							VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
		// 								ImageReference: &armworkloadssapvirtualinstance.ImageReference{
		// 									Offer: to.Ptr("RHEL-SAP"),
		// 									Publisher: to.Ptr("RedHat"),
		// 									SKU: to.Ptr("84sapha-gen2"),
		// 									Version: to.Ptr("latest"),
		// 								},
		// 								OSProfile: &armworkloadssapvirtualinstance.OSProfile{
		// 									AdminUsername: to.Ptr("{your-username}"),
		// 									OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
		// 										OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
		// 										DisablePasswordAuthentication: to.Ptr(true),
		// 										SSH: &armworkloadssapvirtualinstance.SSHConfiguration{
		// 											PublicKeys: []*armworkloadssapvirtualinstance.SSHPublicKey{
		// 												{
		// 													KeyData: to.Ptr("ssh-rsa public key"),
		// 											}},
		// 										},
		// 									},
		// 								},
		// 								VMSize: to.Ptr("Standard_D8s_v3"),
		// 							},
		// 						},
		// 						CentralServer: &armworkloadssapvirtualinstance.CentralServerConfiguration{
		// 							InstanceCount: to.Ptr[int64](1),
		// 							SubnetID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/vnet1/subnets/subnetId1"),
		// 							VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
		// 								ImageReference: &armworkloadssapvirtualinstance.ImageReference{
		// 									Offer: to.Ptr("RHEL-SAP"),
		// 									Publisher: to.Ptr("RedHat"),
		// 									SKU: to.Ptr("84sapha-gen2"),
		// 									Version: to.Ptr("latest"),
		// 								},
		// 								OSProfile: &armworkloadssapvirtualinstance.OSProfile{
		// 									AdminUsername: to.Ptr("{your-username}"),
		// 									OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
		// 										OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
		// 										DisablePasswordAuthentication: to.Ptr(true),
		// 										SSH: &armworkloadssapvirtualinstance.SSHConfiguration{
		// 											PublicKeys: []*armworkloadssapvirtualinstance.SSHPublicKey{
		// 												{
		// 													KeyData: to.Ptr("ssh-rsa public key"),
		// 											}},
		// 										},
		// 									},
		// 								},
		// 								VMSize: to.Ptr("Standard_D8s_v3"),
		// 							},
		// 						},
		// 						DatabaseServer: &armworkloadssapvirtualinstance.DatabaseConfiguration{
		// 							DatabaseType: to.Ptr(armworkloadssapvirtualinstance.SAPDatabaseTypeHANA),
		// 							InstanceCount: to.Ptr[int64](1),
		// 							SubnetID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/vnet1/subnets/subnetId2"),
		// 							VirtualMachineConfiguration: &armworkloadssapvirtualinstance.VirtualMachineConfiguration{
		// 								ImageReference: &armworkloadssapvirtualinstance.ImageReference{
		// 									Offer: to.Ptr("RHEL-SAP"),
		// 									Publisher: to.Ptr("RedHat"),
		// 									SKU: to.Ptr("84sapha-gen2"),
		// 									Version: to.Ptr("latest"),
		// 								},
		// 								OSProfile: &armworkloadssapvirtualinstance.OSProfile{
		// 									AdminUsername: to.Ptr("{your-username}"),
		// 									OSConfiguration: &armworkloadssapvirtualinstance.LinuxConfiguration{
		// 										OSType: to.Ptr(armworkloadssapvirtualinstance.OSTypeLinux),
		// 										DisablePasswordAuthentication: to.Ptr(true),
		// 										SSH: &armworkloadssapvirtualinstance.SSHConfiguration{
		// 											PublicKeys: []*armworkloadssapvirtualinstance.SSHPublicKey{
		// 												{
		// 													KeyData: to.Ptr("ssh-rsa public key"),
		// 											}},
		// 										},
		// 									},
		// 								},
		// 								VMSize: to.Ptr("Standard_D8s_v3"),
		// 							},
		// 						},
		// 					},
		// 					SoftwareConfiguration: &armworkloadssapvirtualinstance.ServiceInitiatedSoftwareConfiguration{
		// 						SoftwareInstallationType: to.Ptr(armworkloadssapvirtualinstance.SAPSoftwareInstallationTypeServiceInitiated),
		// 						BomURL: to.Ptr("https://myaccount.blob.core.windows.net/mycontainer/myblob"),
		// 						SapBitsStorageAccountID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/ajgupt-test/providers/Microsoft.Storage/storageAccounts/ajguptsavnet"),
		// 						SapFqdn: to.Ptr("sapsystem.contoso.net"),
		// 						SoftwareVersion: to.Ptr("SAP S/4HANA 2020 FPS01"),
		// 					},
		// 				},
		// 				Environment: to.Ptr(armworkloadssapvirtualinstance.SAPEnvironmentTypeProd),
		// 				Health: to.Ptr(armworkloadssapvirtualinstance.SAPHealthStateHealthy),
		// 				ProvisioningState: to.Ptr(armworkloadssapvirtualinstance.SapVirtualInstanceProvisioningStateSucceeded),
		// 				SapProduct: to.Ptr(armworkloadssapvirtualinstance.SAPProductTypeS4HANA),
		// 				State: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStateRegistrationComplete),
		// 				Status: to.Ptr(armworkloadssapvirtualinstance.SAPVirtualInstanceStatusRunning),
		// 			},
		// 	}},
		// }
	}
}
