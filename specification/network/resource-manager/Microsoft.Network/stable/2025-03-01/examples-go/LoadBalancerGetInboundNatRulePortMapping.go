package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/LoadBalancerGetInboundNatRulePortMapping.json
func ExampleLoadBalancersClient_Get_getLoadBalancerWithInboundNatRulePortMapping() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLoadBalancersClient().Get(ctx, "rg1", "lb", &armnetwork.LoadBalancersClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LoadBalancer = armnetwork.LoadBalancer{
	// 	Name: to.Ptr("lb"),
	// 	Type: to.Ptr("Microsoft.Network/loadBalancers"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb"),
	// 	Location: to.Ptr("westus"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 	Properties: &armnetwork.LoadBalancerPropertiesFormat{
	// 		BackendAddressPools: []*armnetwork.BackendAddressPool{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"),
	// 				Name: to.Ptr("be-lb"),
	// 				Type: to.Ptr("Microsoft.Network/loadBalancers/backendAddressPools"),
	// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 				Properties: &armnetwork.BackendAddressPoolPropertiesFormat{
	// 					InboundNatRules: []*armnetwork.SubResource{
	// 						{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/inboundNatRules/in-nat-rule"),
	// 					}},
	// 					LoadBalancerBackendAddresses: []*armnetwork.LoadBalancerBackendAddress{
	// 						{
	// 							Name: to.Ptr("8bec96ef-0a6a-45b8-9860-4ffde7bf3572"),
	// 							Properties: &armnetwork.LoadBalancerBackendAddressPropertiesFormat{
	// 								InboundNatRulesPortMapping: []*armnetwork.NatRulePortMapping{
	// 									{
	// 										BackendPort: to.Ptr[int32](3389),
	// 										FrontendPort: to.Ptr[int32](3389),
	// 										InboundNatRuleName: to.Ptr("natRule"),
	// 								}},
	// 								IPAddress: to.Ptr("10.0.0.4"),
	// 								Subnet: &armnetwork.SubResource{
	// 									ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
	// 								},
	// 							},
	// 					}},
	// 					LoadBalancingRules: []*armnetwork.SubResource{
	// 						{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
	// 					}},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 				},
	// 		}},
	// 		FrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
	// 				Name: to.Ptr("fe-lb"),
	// 				Type: to.Ptr("Microsoft.Network/loadBalancers/frontendIPConfigurations"),
	// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 				Properties: &armnetwork.FrontendIPConfigurationPropertiesFormat{
	// 					LoadBalancingRules: []*armnetwork.SubResource{
	// 						{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
	// 					}},
	// 					PrivateIPAddress: to.Ptr("10.0.1.4"),
	// 					PrivateIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 					PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					Subnet: &armnetwork.Subnet{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
	// 					},
	// 				},
	// 		}},
	// 		InboundNatPools: []*armnetwork.InboundNatPool{
	// 		},
	// 		InboundNatRules: []*armnetwork.InboundNatRule{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/inboundNatRules/in-nat-rule"),
	// 				Name: to.Ptr("in-nat-rule"),
	// 				Type: to.Ptr("Microsoft.Network/loadBalancers/inboundNatRules"),
	// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 				Properties: &armnetwork.InboundNatRulePropertiesFormat{
	// 					BackendAddressPool: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"),
	// 					},
	// 					BackendPort: to.Ptr[int32](3389),
	// 					EnableFloatingIP: to.Ptr(true),
	// 					EnableTCPReset: to.Ptr(true),
	// 					FrontendIPConfiguration: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
	// 					},
	// 					FrontendPort: to.Ptr[int32](0),
	// 					FrontendPortRangeEnd: to.Ptr[int32](4000),
	// 					FrontendPortRangeStart: to.Ptr[int32](3389),
	// 					IdleTimeoutInMinutes: to.Ptr[int32](15),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
	// 				},
	// 		}},
	// 		LoadBalancingRules: []*armnetwork.LoadBalancingRule{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
	// 				Name: to.Ptr("rulelb"),
	// 				Type: to.Ptr("Microsoft.Network/loadBalancers/loadBalancingRules"),
	// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 				Properties: &armnetwork.LoadBalancingRulePropertiesFormat{
	// 					BackendAddressPool: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"),
	// 					},
	// 					BackendPort: to.Ptr[int32](80),
	// 					DisableOutboundSnat: to.Ptr(false),
	// 					EnableFloatingIP: to.Ptr(true),
	// 					EnableTCPReset: to.Ptr(true),
	// 					FrontendIPConfiguration: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
	// 					},
	// 					FrontendPort: to.Ptr[int32](80),
	// 					IdleTimeoutInMinutes: to.Ptr[int32](15),
	// 					LoadDistribution: to.Ptr(armnetwork.LoadDistributionDefault),
	// 					Probe: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"),
	// 					},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
	// 				},
	// 		}},
	// 		Probes: []*armnetwork.Probe{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"),
	// 				Name: to.Ptr("probe-lb"),
	// 				Type: to.Ptr("Microsoft.Network/loadBalancers/probes"),
	// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 				Properties: &armnetwork.ProbePropertiesFormat{
	// 					IntervalInSeconds: to.Ptr[int32](15),
	// 					LoadBalancingRules: []*armnetwork.SubResource{
	// 						{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
	// 					}},
	// 					NumberOfProbes: to.Ptr[int32](2),
	// 					Port: to.Ptr[int32](80),
	// 					ProbeThreshold: to.Ptr[int32](1),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					RequestPath: to.Ptr("healthcheck.aspx"),
	// 					Protocol: to.Ptr(armnetwork.ProbeProtocolHTTP),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 	},
	// 	SKU: &armnetwork.LoadBalancerSKU{
	// 		Name: to.Ptr(armnetwork.LoadBalancerSKUNameStandard),
	// 	},
	// }
}
