package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ccea12be6cd5e64743499d46f7a9c8b60df52db1/specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2024-11-04/examples/CloudService_Create_WithMultiRole_WithZones.json
func ExampleCloudServicesClient_BeginCreateOrUpdate_createNewCloudServiceWithMultipleRolesInASpecificAvailabilityZone() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcompute.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCloudServicesClient().BeginCreateOrUpdate(ctx, "ConstosoRG", "{cs-name}", armcompute.CloudService{
		Location: to.Ptr("westus"),
		Properties: &armcompute.CloudServiceProperties{
			Configuration: to.Ptr("{ServiceConfiguration}"),
			NetworkProfile: &armcompute.CloudServiceNetworkProfile{
				LoadBalancerConfigurations: []*armcompute.LoadBalancerConfiguration{
					{
						Name: to.Ptr("contosolb"),
						Properties: &armcompute.LoadBalancerConfigurationProperties{
							FrontendIPConfigurations: []*armcompute.LoadBalancerFrontendIPConfiguration{
								{
									Name: to.Ptr("contosofe"),
									Properties: &armcompute.LoadBalancerFrontendIPConfigurationProperties{
										PublicIPAddress: &armcompute.SubResource{
											ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/ConstosoRG/providers/Microsoft.Network/publicIPAddresses/contosopublicip"),
										},
									},
								}},
						},
					}},
			},
			PackageURL: to.Ptr("{PackageUrl}"),
			RoleProfile: &armcompute.CloudServiceRoleProfile{
				Roles: []*armcompute.CloudServiceRoleProfileProperties{
					{
						Name: to.Ptr("ContosoFrontend"),
						SKU: &armcompute.CloudServiceRoleSKU{
							Name:     to.Ptr("Standard_D1_v2"),
							Capacity: to.Ptr[int64](1),
							Tier:     to.Ptr("Standard"),
						},
					},
					{
						Name: to.Ptr("ContosoBackend"),
						SKU: &armcompute.CloudServiceRoleSKU{
							Name:     to.Ptr("Standard_D1_v2"),
							Capacity: to.Ptr[int64](1),
							Tier:     to.Ptr("Standard"),
						},
					}},
			},
			UpgradeMode: to.Ptr(armcompute.CloudServiceUpgradeModeAuto),
		},
		Zones: []*string{
			to.Ptr("1")},
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
	// 	Location: to.Ptr("westus"),
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
	// 		PackageURL: to.Ptr("{PackageUrl}"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		RoleProfile: &armcompute.CloudServiceRoleProfile{
	// 			Roles: []*armcompute.CloudServiceRoleProfileProperties{
	// 				{
	// 					Name: to.Ptr("ContosoFrontend"),
	// 					SKU: &armcompute.CloudServiceRoleSKU{
	// 						Name: to.Ptr("Standard_D1_v2"),
	// 						Capacity: to.Ptr[int64](1),
	// 						Tier: to.Ptr("Standard"),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("ContosoBackend"),
	// 					SKU: &armcompute.CloudServiceRoleSKU{
	// 						Name: to.Ptr("Standard_D1_v2"),
	// 						Capacity: to.Ptr[int64](1),
	// 						Tier: to.Ptr("Standard"),
	// 					},
	// 			}},
	// 		},
	// 		UniqueID: to.Ptr("7f3edf91-cb34-4a3e-971a-177dc3dd43cb"),
	// 		UpgradeMode: to.Ptr(armcompute.CloudServiceUpgradeModeAuto),
	// 	},
	// 	SystemData: &armcompute.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}
