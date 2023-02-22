package armworkloads_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloads/armworkloads"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Get.json
func ExampleSAPVirtualInstancesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armworkloads.NewSAPVirtualInstancesClient("8e17e36c-42e9-4cd5-a078-7b44883414e0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx, "test-rg", "X00", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
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
	// 		Configuration: &armworkloads.DeploymentConfiguration{
	// 			ConfigurationType: to.Ptr(armworkloads.SAPConfigurationTypeDeployment),
	// 			AppLocation: to.Ptr("eastus"),
	// 			InfrastructureConfiguration: &armworkloads.ThreeTierConfiguration{
	// 				AppResourceGroup: to.Ptr("X00-RG"),
	// 				DeploymentType: to.Ptr(armworkloads.SAPDeploymentTypeThreeTier),
	// 				ApplicationServer: &armworkloads.ApplicationServerConfiguration{
	// 					InstanceCount: to.Ptr[int64](10),
	// 					SubnetID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/vnet1/subnets/subnetId3"),
	// 					VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
	// 						ImageReference: &armworkloads.ImageReference{
	// 							Offer: to.Ptr("RHEL-SAP"),
	// 							Publisher: to.Ptr("RedHat"),
	// 							SKU: to.Ptr("7.4"),
	// 							Version: to.Ptr("7.4.2019062505"),
	// 						},
	// 						OSProfile: &armworkloads.OSProfile{
	// 							AdminUsername: to.Ptr("{your-username}"),
	// 							OSConfiguration: &armworkloads.LinuxConfiguration{
	// 								OSType: to.Ptr(armworkloads.OSTypeLinux),
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								SSH: &armworkloads.SSHConfiguration{
	// 									PublicKeys: []*armworkloads.SSHPublicKey{
	// 										{
	// 											KeyData: to.Ptr("ssh-rsa public key"),
	// 									}},
	// 								},
	// 							},
	// 						},
	// 						VMSize: to.Ptr("Standard_D8s_v3"),
	// 					},
	// 				},
	// 				CentralServer: &armworkloads.CentralServerConfiguration{
	// 					InstanceCount: to.Ptr[int64](1),
	// 					SubnetID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/vnet1/subnets/subnetId1"),
	// 					VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
	// 						ImageReference: &armworkloads.ImageReference{
	// 							Offer: to.Ptr("RHEL-SAP"),
	// 							Publisher: to.Ptr("RedHat"),
	// 							SKU: to.Ptr("7.4"),
	// 							Version: to.Ptr("7.4.2019062505"),
	// 						},
	// 						OSProfile: &armworkloads.OSProfile{
	// 							AdminUsername: to.Ptr("{your-username}"),
	// 							OSConfiguration: &armworkloads.LinuxConfiguration{
	// 								OSType: to.Ptr(armworkloads.OSTypeLinux),
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								SSH: &armworkloads.SSHConfiguration{
	// 									PublicKeys: []*armworkloads.SSHPublicKey{
	// 										{
	// 											KeyData: to.Ptr("ssh-rsa public key"),
	// 									}},
	// 								},
	// 							},
	// 						},
	// 						VMSize: to.Ptr("Standard_D8s_v3"),
	// 					},
	// 				},
	// 				DatabaseServer: &armworkloads.DatabaseConfiguration{
	// 					DatabaseType: to.Ptr(armworkloads.SAPDatabaseTypeHANA),
	// 					InstanceCount: to.Ptr[int64](1),
	// 					SubnetID: to.Ptr("/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/vnet1/subnets/subnetId2"),
	// 					VirtualMachineConfiguration: &armworkloads.VirtualMachineConfiguration{
	// 						ImageReference: &armworkloads.ImageReference{
	// 							Offer: to.Ptr("RHEL-SAP"),
	// 							Publisher: to.Ptr("RedHat"),
	// 							SKU: to.Ptr("7.4"),
	// 							Version: to.Ptr("7.4.2019062505"),
	// 						},
	// 						OSProfile: &armworkloads.OSProfile{
	// 							AdminUsername: to.Ptr("{your-username}"),
	// 							OSConfiguration: &armworkloads.LinuxConfiguration{
	// 								OSType: to.Ptr(armworkloads.OSTypeLinux),
	// 								DisablePasswordAuthentication: to.Ptr(true),
	// 								SSH: &armworkloads.SSHConfiguration{
	// 									PublicKeys: []*armworkloads.SSHPublicKey{
	// 										{
	// 											KeyData: to.Ptr("ssh-rsa public key"),
	// 									}},
	// 								},
	// 							},
	// 						},
	// 						VMSize: to.Ptr("Standard_D8s_v3"),
	// 					},
	// 				},
	// 			},
	// 			SoftwareConfiguration: &armworkloads.ServiceInitiatedSoftwareConfiguration{
	// 				SoftwareInstallationType: to.Ptr(armworkloads.SAPSoftwareInstallationTypeServiceInitiated),
	// 				BomURL: to.Ptr("https://myaccount.blob.core.windows.net/mycontainer/myblob"),
	// 				SapBitsStorageAccountID: to.Ptr("/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/ajgupt-test/providers/Microsoft.Storage/storageAccounts/ajguptsavnet"),
	// 				SapFqdn: to.Ptr("sapsystem.contoso.net"),
	// 				SoftwareVersion: to.Ptr("SAP S/4HANA 2020 FPS01"),
	// 			},
	// 		},
	// 		Environment: to.Ptr(armworkloads.SAPEnvironmentTypeProd),
	// 		Health: to.Ptr(armworkloads.SAPHealthStateHealthy),
	// 		ManagedResourceGroupConfiguration: &armworkloads.ManagedRGConfiguration{
	// 			Name: to.Ptr("mrg-x00-6d875e77-e412-4d7d-9af4-8895278b4443"),
	// 		},
	// 		ProvisioningState: to.Ptr(armworkloads.SapVirtualInstanceProvisioningStateSucceeded),
	// 		SapProduct: to.Ptr(armworkloads.SAPProductTypeS4HANA),
	// 		State: to.Ptr(armworkloads.SAPVirtualInstanceStateRegistrationComplete),
	// 		Status: to.Ptr(armworkloads.SAPVirtualInstanceStatusRunning),
	// 	},
	// }
}
