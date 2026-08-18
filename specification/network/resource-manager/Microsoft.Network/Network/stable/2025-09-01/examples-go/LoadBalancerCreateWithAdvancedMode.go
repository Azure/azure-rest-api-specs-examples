package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/LoadBalancerCreateWithAdvancedMode.json
func ExampleLoadBalancersClient_BeginCreateOrUpdate_createLoadBalancerWithAdvancedMode() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLoadBalancersClient().BeginCreateOrUpdate(ctx, "rg1", "lb", armnetwork.LoadBalancer{
		Location: to.Ptr("eastus"),
		Properties: &armnetwork.LoadBalancerPropertiesFormat{
			BackendAddressPools: []*armnetwork.BackendAddressPool{
				{
					Name:       to.Ptr("be-lb"),
					Properties: &armnetwork.BackendAddressPoolPropertiesFormat{},
				},
			},
			FrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
				{
					Name: to.Ptr("fe-lb"),
					Properties: &armnetwork.FrontendIPConfigurationPropertiesFormat{
						EnableConnectionTracking: to.Ptr(true),
						Subnet: &armnetwork.Subnet{
							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
						},
					},
				},
			},
			LoadBalancingRules: []*armnetwork.LoadBalancingRule{
				{
					Name: to.Ptr("rulelb"),
					Properties: &armnetwork.LoadBalancingRulePropertiesFormat{
						BackendAddressPool: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"),
						},
						BackendPort:      to.Ptr[int32](4789),
						EnableFloatingIP: to.Ptr(true),
						FrontendIPConfiguration: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
						},
						FrontendPort:     to.Ptr[int32](4789),
						LoadDistribution: to.Ptr(armnetwork.LoadDistributionDefault),
						Probe: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"),
						},
						Protocol: to.Ptr(armnetwork.TransportProtocolUDP),
					},
				},
			},
			Mode: to.Ptr(armnetwork.LoadBalancerModeAdvanced),
			Probes: []*armnetwork.Probe{
				{
					Name: to.Ptr("probe-lb"),
					Properties: &armnetwork.ProbePropertiesFormat{
						IntervalInSeconds: to.Ptr[int32](15),
						NumberOfProbes:    to.Ptr[int32](2),
						Port:              to.Ptr[int32](80),
						ProbeThreshold:    to.Ptr[int32](1),
						RequestPath:       to.Ptr("healthcheck.aspx"),
						Protocol:          to.Ptr(armnetwork.ProbeProtocolHTTP),
					},
				},
			},
			Scope: to.Ptr(armnetwork.LoadBalancerScopePublic),
		},
		SKU: &armnetwork.LoadBalancerSKU{
			Name: to.Ptr(armnetwork.LoadBalancerSKUNameStandard),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.LoadBalancersClientCreateOrUpdateResponse{
	// 	LoadBalancer: armnetwork.LoadBalancer{
	// 		Name: to.Ptr("lb"),
	// 		Type: to.Ptr("Microsoft.Network/loadBalancers"),
	// 		Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb"),
	// 		Location: to.Ptr("eastus"),
	// 		Properties: &armnetwork.LoadBalancerPropertiesFormat{
	// 			BackendAddressPools: []*armnetwork.BackendAddressPool{
	// 				{
	// 					Name: to.Ptr("be-lb"),
	// 					Type: to.Ptr("Microsoft.Network/loadBalancers/backendAddressPools"),
	// 					Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
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
	// 					Type: to.Ptr("Microsoft.Network/loadBalancers/frontendIPConfigurations"),
	// 					Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
	// 					Properties: &armnetwork.FrontendIPConfigurationPropertiesFormat{
	// 						EnableConnectionTracking: to.Ptr(true),
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
	// 			LoadBalancingRules: []*armnetwork.LoadBalancingRule{
	// 				{
	// 					Name: to.Ptr("rulelb"),
	// 					Type: to.Ptr("Microsoft.Network/loadBalancers/loadBalancingRules"),
	// 					Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/loadBalancingRules/rulelb"),
	// 					Properties: &armnetwork.LoadBalancingRulePropertiesFormat{
	// 						BackendAddressPool: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"),
	// 						},
	// 						BackendPort: to.Ptr[int32](4789),
	// 						DisableOutboundSnat: to.Ptr(false),
	// 						EnableFloatingIP: to.Ptr(true),
	// 						FrontendIPConfiguration: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
	// 						},
	// 						FrontendPort: to.Ptr[int32](4789),
	// 						LoadDistribution: to.Ptr(armnetwork.LoadDistributionDefault),
	// 						Probe: &armnetwork.SubResource{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"),
	// 						},
	// 						ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 						Protocol: to.Ptr(armnetwork.TransportProtocolUDP),
	// 					},
	// 				},
	// 			},
	// 			Mode: to.Ptr(armnetwork.LoadBalancerModeAdvanced),
	// 			Probes: []*armnetwork.Probe{
	// 				{
	// 					Name: to.Ptr("probe-lb"),
	// 					Type: to.Ptr("Microsoft.Network/loadBalancers/probes"),
	// 					Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
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
	// 			ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 			Scope: to.Ptr(armnetwork.LoadBalancerScopePublic),
	// 		},
	// 		SKU: &armnetwork.LoadBalancerSKU{
	// 			Name: to.Ptr(armnetwork.LoadBalancerSKUNameStandard),
	// 		},
	// 	},
	// }
}
