package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/LoadBalancerCreateWithInboundNatPool.json
func ExampleLoadBalancersClient_BeginCreateOrUpdate_loadBalancerCreateWithInboundNatPool() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewLoadBalancersClient("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "lb", armnetwork.LoadBalancer{
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
	// TODO: use response item
	_ = res
}
