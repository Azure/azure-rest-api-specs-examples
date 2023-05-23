package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/5d2adf9b7fda669b4a2538c65e937ee74fe3f966/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-09-04/examples/CloudService_Update_ToIncludeTags.json
func ExampleCloudServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudServicesClient().BeginUpdate(ctx, "ConstosoRG", "{cs-name}", armcompute.CloudServiceUpdate{
		Tags: map[string]*string{
			"Documentation": to.Ptr("RestAPI"),
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
	// res.CloudService = armcompute.CloudService{
	// 	Name: to.Ptr("{cs-name}"),
	// 	Type: to.Ptr("Microsoft.Compute/cloudServices"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Compute/cloudServices/{cs-name}"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Properties: &armcompute.CloudServiceProperties{
	// 		Configuration: to.Ptr("{ServiceConfiguration}"),
	// 		NetworkProfile: &armcompute.CloudServiceNetworkProfile{
	// 			LoadBalancerConfigurations: []*armcompute.LoadBalancerConfiguration{
	// 				{
	// 					Name: to.Ptr("contosolb"),
	// 					Properties: &armcompute.LoadBalancerConfigurationProperties{
	// 						FrontendIPConfigurations: []*armcompute.LoadBalancerFrontendIPConfiguration{
	// 							{
	// 								Name: to.Ptr("contosofe"),
	// 								Properties: &armcompute.LoadBalancerFrontendIPConfigurationProperties{
	// 									PublicIPAddress: &armcompute.SubResource{
	// 										ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Network/publicIPAddresses/contosopublicip"),
	// 									},
	// 								},
	// 						}},
	// 					},
	// 			}},
	// 		},
	// 		OSProfile: &armcompute.CloudServiceOsProfile{
	// 			Secrets: []*armcompute.CloudServiceVaultSecretGroup{
	// 			},
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RoleProfile: &armcompute.CloudServiceRoleProfile{
	// 			Roles: []*armcompute.CloudServiceRoleProfileProperties{
	// 				{
	// 					Name: to.Ptr("ContosoFrontend"),
	// 					SKU: &armcompute.CloudServiceRoleSKU{
	// 						Name: to.Ptr("Standard_D1_v2"),
	// 						Capacity: to.Ptr[int64](2),
	// 						Tier: to.Ptr("Standard"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("ContosoBackend"),
	// 					SKU: &armcompute.CloudServiceRoleSKU{
	// 						Name: to.Ptr("Standard_D1_v2"),
	// 						Capacity: to.Ptr[int64](2),
	// 						Tier: to.Ptr("Standard"),
	// 					},
	// 			}},
	// 		},
	// 		UniqueID: to.Ptr("4ccb4323-4740-4545-bb81-780b27375947"),
	// 		UpgradeMode: to.Ptr(armcompute.CloudServiceUpgradeModeAuto),
	// 	},
	// 	Tags: map[string]*string{
	// 		"Documentation": to.Ptr("RestAPI"),
	// 	},
	// }
}
