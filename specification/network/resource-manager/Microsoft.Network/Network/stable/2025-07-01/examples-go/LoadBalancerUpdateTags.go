package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/LoadBalancerUpdateTags.json
func ExampleLoadBalancersClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLoadBalancersClient().UpdateTags(ctx, "rg1", "lb", armnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.LoadBalancersClientUpdateTagsResponse{
	// 	LoadBalancer: armnetwork.LoadBalancer{
	// 		Name: to.Ptr("lb"),
	// 		Type: to.Ptr("Microsoft.Network/loadBalancers"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb"),
	// 		Location: to.Ptr("westus"),
	// 		Properties: &armnetwork.LoadBalancerPropertiesFormat{
	// 			BackendAddressPools: []*armnetwork.BackendAddressPool{
	// 				{
	// 					Name: to.Ptr("be-lb"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"),
	// 					Properties: &armnetwork.BackendAddressPoolPropertiesFormat{
	// 						LoadBalancingRules: []*armnetwork.SubResource{
	// 							{
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
	// 							},
	// 						},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					},
	// 				},
	// 			},
	// 			FrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
	// 				{
	// 					Name: to.Ptr("fe-lb"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
	// 					Properties: &armnetwork.FrontendIPConfigurationPropertiesFormat{
	// 						InboundNatRules: []*armnetwork.SubResource{
	// 							{
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/inboundNatRules/in-nat-rule"),
	// 							},
	// 						},
	// 						LoadBalancingRules: []*armnetwork.SubResource{
	// 							{
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
	// 							},
	// 						},
	// 						PrivateIPAddress: to.Ptr("10.0.1.4"),
	// 						PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						Subnet: &armnetwork.Subnet{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
	// 						},
	// 					},
	// 				},
	// 			},
	// 			InboundNatPools: []*armnetwork.InboundNatPool{
	// 			},
	// 			InboundNatRules: []*armnetwork.InboundNatRule{
	// 				{
	// 					Name: to.Ptr("in-nat-rule"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/inboundNatRules/in-nat-rule"),
	// 					Properties: &armnetwork.InboundNatRulePropertiesFormat{
	// 						BackendPort: to.Ptr[int32](3389),
	// 						EnableFloatingIP: to.Ptr(true),
	// 						FrontendIPConfiguration: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
	// 						},
	// 						FrontendPort: to.Ptr[int32](3389),
	// 						IdleTimeoutInMinutes: to.Ptr[int32](15),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
	// 					},
	// 				},
	// 			},
	// 			LoadBalancingRules: []*armnetwork.LoadBalancingRule{
	// 				{
	// 					Name: to.Ptr("rulelb"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
	// 					Properties: &armnetwork.LoadBalancingRulePropertiesFormat{
	// 						BackendAddressPool: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"),
	// 						},
	// 						BackendPort: to.Ptr[int32](80),
	// 						DisableOutboundSnat: to.Ptr(false),
	// 						EnableFloatingIP: to.Ptr(true),
	// 						FrontendIPConfiguration: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
	// 						},
	// 						FrontendPort: to.Ptr[int32](80),
	// 						IdleTimeoutInMinutes: to.Ptr[int32](15),
	// 						LoadDistribution: to.Ptr(armnetwork.LoadDistributionDefault),
	// 						Probe: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"),
	// 						},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
	// 					},
	// 				},
	// 			},
	// 			OutboundRules: []*armnetwork.OutboundRule{
	// 			},
	// 			Probes: []*armnetwork.Probe{
	// 				{
	// 					Name: to.Ptr("probe-lb"),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"),
	// 					Properties: &armnetwork.ProbePropertiesFormat{
	// 						IntervalInSeconds: to.Ptr[int32](15),
	// 						LoadBalancingRules: []*armnetwork.SubResource{
	// 							{
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
	// 							},
	// 						},
	// 						NumberOfProbes: to.Ptr[int32](2),
	// 						Port: to.Ptr[int32](80),
	// 						ProbeThreshold: to.Ptr[int32](1),
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						RequestPath: to.Ptr("healthcheck.aspx"),
	// 						Protocol: to.Ptr(armnetwork.ProbeProtocolHTTP),
	// 					},
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		},
	// 		Tags: map[string]*string{
	// 			"tag1": to.Ptr("value1"),
	// 			"tag2": to.Ptr("value2"),
	// 		},
	// 	},
	// }
}
