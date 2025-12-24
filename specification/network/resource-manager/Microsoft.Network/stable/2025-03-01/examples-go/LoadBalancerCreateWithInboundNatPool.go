package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/LoadBalancerCreateWithInboundNatPool.json
func ExampleLoadBalancersClient_BeginCreateOrUpdate_createLoadBalancerWithInboundNatPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewLoadBalancersClient().BeginCreateOrUpdate(ctx, "rg1", "lb", armnetwork.LoadBalancer{
		Location: to.Ptr("eastus"),
		Properties: &armnetwork.LoadBalancerPropertiesFormat{
			BackendAddressPools: []*armnetwork.BackendAddressPool{},
			FrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
				{
					ID:   to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/test"),
					Name: to.Ptr("test"),
					Properties: &armnetwork.FrontendIPConfigurationPropertiesFormat{
						PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
						Subnet: &armnetwork.Subnet{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/lbvnet/subnets/lbsubnet"),
						},
					},
					Zones: []*string{},
				}},
			InboundNatPools: []*armnetwork.InboundNatPool{
				{
					ID:   to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/inboundNatPools/test"),
					Name: to.Ptr("test"),
					Properties: &armnetwork.InboundNatPoolPropertiesFormat{
						BackendPort:      to.Ptr[int32](8888),
						EnableFloatingIP: to.Ptr(true),
						EnableTCPReset:   to.Ptr(true),
						FrontendIPConfiguration: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/test"),
						},
						FrontendPortRangeEnd:   to.Ptr[int32](8085),
						FrontendPortRangeStart: to.Ptr[int32](8080),
						IdleTimeoutInMinutes:   to.Ptr[int32](10),
						Protocol:               to.Ptr(armnetwork.TransportProtocolTCP),
					},
				}},
			InboundNatRules:    []*armnetwork.InboundNatRule{},
			LoadBalancingRules: []*armnetwork.LoadBalancingRule{},
			OutboundRules:      []*armnetwork.OutboundRule{},
			Probes:             []*armnetwork.Probe{},
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
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LoadBalancer = armnetwork.LoadBalancer{
	// 	Name: to.Ptr("lb"),
	// 	Type: to.Ptr("Microsoft.Network/loadBalancers"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb"),
	// 	Location: to.Ptr("eastus"),
	// 	Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 	Properties: &armnetwork.LoadBalancerPropertiesFormat{
	// 		BackendAddressPools: []*armnetwork.BackendAddressPool{
	// 		},
	// 		FrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/test"),
	// 				Name: to.Ptr("test"),
	// 				Type: to.Ptr("Microsoft.Network/loadBalancers/frontendIPConfigurations"),
	// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 				Properties: &armnetwork.FrontendIPConfigurationPropertiesFormat{
	// 					InboundNatPools: []*armnetwork.SubResource{
	// 						{
	// 							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/inboundNatPools/test"),
	// 					}},
	// 					PrivateIPAddress: to.Ptr("10.0.1.4"),
	// 					PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					Subnet: &armnetwork.Subnet{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/lbvnet/subnets/lbsubnet"),
	// 					},
	// 				},
	// 		}},
	// 		InboundNatPools: []*armnetwork.InboundNatPool{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/inboundNatPools/test"),
	// 				Name: to.Ptr("test"),
	// 				Type: to.Ptr("Microsoft.Network/loadBalancers/inboundNatPools"),
	// 				Etag: to.Ptr("W/\"00000000-0000-0000-0000-00000000\""),
	// 				Properties: &armnetwork.InboundNatPoolPropertiesFormat{
	// 					BackendPort: to.Ptr[int32](8888),
	// 					EnableFloatingIP: to.Ptr(true),
	// 					EnableTCPReset: to.Ptr(true),
	// 					FrontendIPConfiguration: &armnetwork.SubResource{
	// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/test"),
	// 					},
	// 					FrontendPortRangeEnd: to.Ptr[int32](8085),
	// 					FrontendPortRangeStart: to.Ptr[int32](8080),
	// 					IdleTimeoutInMinutes: to.Ptr[int32](10),
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
	// 				},
	// 		}},
	// 		InboundNatRules: []*armnetwork.InboundNatRule{
	// 		},
	// 		LoadBalancingRules: []*armnetwork.LoadBalancingRule{
	// 		},
	// 		OutboundRules: []*armnetwork.OutboundRule{
	// 		},
	// 		Probes: []*armnetwork.Probe{
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		ResourceGUID: to.Ptr("00000000-0000-0000-0000-00000000"),
	// 	},
	// 	SKU: &armnetwork.LoadBalancerSKU{
	// 		Name: to.Ptr(armnetwork.LoadBalancerSKUNameStandard),
	// 	},
	// }
}
