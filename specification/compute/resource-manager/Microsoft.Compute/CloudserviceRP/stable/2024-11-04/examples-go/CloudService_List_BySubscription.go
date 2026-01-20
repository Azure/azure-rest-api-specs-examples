package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudService_List_BySubscription.json
func ExampleCloudServicesClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCloudServicesClient().NewListAllPager(nil)
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
		// page.CloudServiceListResult = armcompute.CloudServiceListResult{
		// 	Value: []*armcompute.CloudService{
		// 		{
		// 			Name: to.Ptr("{cs-name}"),
		// 			Type: to.Ptr("Microsoft.Compute/cloudServices"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Properties: &armcompute.CloudServiceProperties{
		// 				Configuration: to.Ptr("{ServiceConfiguration}"),
		// 				ExtensionProfile: &armcompute.CloudServiceExtensionProfile{
		// 					Extensions: []*armcompute.Extension{
		// 						{
		// 							Name: to.Ptr("RDPExtension"),
		// 							Properties: &armcompute.CloudServiceExtensionProperties{
		// 								Type: to.Ptr("RDP"),
		// 								AutoUpgradeMinorVersion: to.Ptr(false),
		// 								ProvisioningState: to.Ptr("Succeeded"),
		// 								Publisher: to.Ptr("Microsoft.Windows.Azure.Extensions"),
		// 								RolesAppliedTo: []*string{
		// 									to.Ptr("*")},
		// 									Settings: "<PublicConfig><UserName>userazure</UserName><Expiration>01/12/2022 16:29:02</Expiration></PublicConfig>",
		// 									TypeHandlerVersion: to.Ptr("1.2"),
		// 								},
		// 						}},
		// 					},
		// 					NetworkProfile: &armcompute.CloudServiceNetworkProfile{
		// 						LoadBalancerConfigurations: []*armcompute.LoadBalancerConfiguration{
		// 							{
		// 								Name: to.Ptr("contosolb"),
		// 								Properties: &armcompute.LoadBalancerConfigurationProperties{
		// 									FrontendIPConfigurations: []*armcompute.LoadBalancerFrontendIPConfiguration{
		// 										{
		// 											Name: to.Ptr("contosofe"),
		// 											Properties: &armcompute.LoadBalancerFrontendIPConfigurationProperties{
		// 												PublicIPAddress: &armcompute.SubResource{
		// 													ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Network/publicIPAddresses/contosopublicip"),
		// 												},
		// 											},
		// 									}},
		// 								},
		// 						}},
		// 					},
		// 					OSProfile: &armcompute.CloudServiceOsProfile{
		// 						Secrets: []*armcompute.CloudServiceVaultSecretGroup{
		// 						},
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					RoleProfile: &armcompute.CloudServiceRoleProfile{
		// 						Roles: []*armcompute.CloudServiceRoleProfileProperties{
		// 							{
		// 								Name: to.Ptr("ContosoFrontend"),
		// 								SKU: &armcompute.CloudServiceRoleSKU{
		// 									Name: to.Ptr("Standard_D1_v2"),
		// 									Capacity: to.Ptr[int64](2),
		// 									Tier: to.Ptr("Standard"),
		// 								},
		// 							},
		// 							{
		// 								Name: to.Ptr("ContosoBackend"),
		// 								SKU: &armcompute.CloudServiceRoleSKU{
		// 									Name: to.Ptr("Standard_D1_v2"),
		// 									Capacity: to.Ptr[int64](2),
		// 									Tier: to.Ptr("Standard"),
		// 								},
		// 						}},
		// 					},
		// 					UniqueID: to.Ptr("4ccb4323-4740-4545-bb81-780b27375947"),
		// 					UpgradeMode: to.Ptr(armcompute.CloudServiceUpgradeModeAuto),
		// 				},
		// 		}},
		// 	}
	}
}
