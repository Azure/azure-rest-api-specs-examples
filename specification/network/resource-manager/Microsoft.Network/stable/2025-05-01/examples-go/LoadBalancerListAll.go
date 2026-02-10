package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/LoadBalancerListAll.json
func ExampleLoadBalancersClient_NewListAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLoadBalancersClient().NewListAllPager(nil)
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
		// page.LoadBalancerListResult = armnetwork.LoadBalancerListResult{
		// 	Value: []*armnetwork.LoadBalancer{
		// 		{
		// 			Name: to.Ptr("lb"),
		// 			Type: to.Ptr("Microsoft.Network/loadBalancers"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.LoadBalancerPropertiesFormat{
		// 				BackendAddressPools: []*armnetwork.BackendAddressPool{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/belb"),
		// 						Name: to.Ptr("belb"),
		// 						Type: to.Ptr("Microsoft.Network/loadBalancers/backendAddressPools"),
		// 						Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 						Properties: &armnetwork.BackendAddressPoolPropertiesFormat{
		// 							LoadBalancingRules: []*armnetwork.SubResource{
		// 								{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
		// 							}},
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 						},
		// 				}},
		// 				FrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/felb"),
		// 						Name: to.Ptr("felb"),
		// 						Type: to.Ptr("Microsoft.Network/loadBalancers/frontendIPConfigurations"),
		// 						Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 						Properties: &armnetwork.FrontendIPConfigurationPropertiesFormat{
		// 							InboundNatRules: []*armnetwork.SubResource{
		// 								{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/inboundNatRules/inrlb"),
		// 							}},
		// 							LoadBalancingRules: []*armnetwork.SubResource{
		// 								{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
		// 							}},
		// 							PrivateIPAddress: to.Ptr("10.0.1.4"),
		// 							PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							Subnet: &armnetwork.Subnet{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
		// 							},
		// 						},
		// 				}},
		// 				InboundNatPools: []*armnetwork.InboundNatPool{
		// 				},
		// 				InboundNatRules: []*armnetwork.InboundNatRule{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/inboundNatRules/inrlb"),
		// 						Name: to.Ptr("inrlb"),
		// 						Type: to.Ptr("Microsoft.Network/loadBalancers/inboundNatRules"),
		// 						Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 						Properties: &armnetwork.InboundNatRulePropertiesFormat{
		// 							BackendPort: to.Ptr[int32](3389),
		// 							EnableFloatingIP: to.Ptr(true),
		// 							EnableTCPReset: to.Ptr(true),
		// 							FrontendIPConfiguration: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/felb"),
		// 							},
		// 							FrontendPort: to.Ptr[int32](3389),
		// 							IdleTimeoutInMinutes: to.Ptr[int32](15),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
		// 						},
		// 				}},
		// 				LoadBalancingRules: []*armnetwork.LoadBalancingRule{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
		// 						Name: to.Ptr("rulelb"),
		// 						Type: to.Ptr("Microsoft.Network/loadBalancers/loadBalancingRules"),
		// 						Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 						Properties: &armnetwork.LoadBalancingRulePropertiesFormat{
		// 							BackendAddressPool: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/belb"),
		// 							},
		// 							BackendPort: to.Ptr[int32](80),
		// 							EnableFloatingIP: to.Ptr(true),
		// 							EnableTCPReset: to.Ptr(true),
		// 							FrontendIPConfiguration: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/felb"),
		// 							},
		// 							FrontendPort: to.Ptr[int32](80),
		// 							IdleTimeoutInMinutes: to.Ptr[int32](15),
		// 							LoadDistribution: to.Ptr(armnetwork.LoadDistributionDefault),
		// 							Probe: &armnetwork.SubResource{
		// 								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/prlb"),
		// 							},
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
		// 						},
		// 				}},
		// 				Probes: []*armnetwork.Probe{
		// 					{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/prlb"),
		// 						Name: to.Ptr("prlb"),
		// 						Type: to.Ptr("Microsoft.Network/loadBalancers/probes"),
		// 						Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 						Properties: &armnetwork.ProbePropertiesFormat{
		// 							IntervalInSeconds: to.Ptr[int32](15),
		// 							LoadBalancingRules: []*armnetwork.SubResource{
		// 								{
		// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
		// 							}},
		// 							NumberOfProbes: to.Ptr[int32](2),
		// 							Port: to.Ptr[int32](80),
		// 							ProbeThreshold: to.Ptr[int32](1),
		// 							ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 							RequestPath: to.Ptr("healthcheck.aspx"),
		// 							Protocol: to.Ptr(armnetwork.ProbeProtocolHTTP),
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 			},
		// 			SKU: &armnetwork.LoadBalancerSKU{
		// 				Name: to.Ptr(armnetwork.LoadBalancerSKUNameBasic),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("lb3"),
		// 			Type: to.Ptr("Microsoft.Network/loadBalancers"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/loadBalancers/lb3"),
		// 			Location: to.Ptr("westus"),
		// 			Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
		// 			Properties: &armnetwork.LoadBalancerPropertiesFormat{
		// 				BackendAddressPools: []*armnetwork.BackendAddressPool{
		// 				},
		// 				FrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
		// 				},
		// 				InboundNatPools: []*armnetwork.InboundNatPool{
		// 				},
		// 				InboundNatRules: []*armnetwork.InboundNatRule{
		// 				},
		// 				LoadBalancingRules: []*armnetwork.LoadBalancingRule{
		// 				},
		// 				Probes: []*armnetwork.Probe{
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
		// 			},
		// 			SKU: &armnetwork.LoadBalancerSKU{
		// 				Name: to.Ptr(armnetwork.LoadBalancerSKUNameBasic),
		// 			},
		// 	}},
		// }
	}
}
