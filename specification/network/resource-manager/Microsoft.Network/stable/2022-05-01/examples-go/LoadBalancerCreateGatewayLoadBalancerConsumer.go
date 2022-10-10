package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/LoadBalancerCreateGatewayLoadBalancerConsumer.json
func ExampleLoadBalancersClient_BeginCreateOrUpdate_createLoadBalancerWithGatewayLoadBalancerConsumerConfigured() {
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
			BackendAddressPools: []*armnetwork.BackendAddressPool{
				{
					Name:       to.Ptr("be-lb"),
					Properties: &armnetwork.BackendAddressPoolPropertiesFormat{},
				}},
			FrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
				{
					Name: to.Ptr("fe-lb"),
					Properties: &armnetwork.FrontendIPConfigurationPropertiesFormat{
						GatewayLoadBalancer: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb-provider"),
						},
						Subnet: &armnetwork.Subnet{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
						},
					},
				}},
			InboundNatPools: []*armnetwork.InboundNatPool{},
			InboundNatRules: []*armnetwork.InboundNatRule{
				{
					Name: to.Ptr("in-nat-rule"),
					Properties: &armnetwork.InboundNatRulePropertiesFormat{
						BackendPort:      to.Ptr[int32](3389),
						EnableFloatingIP: to.Ptr(true),
						FrontendIPConfiguration: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
						},
						FrontendPort:         to.Ptr[int32](3389),
						IdleTimeoutInMinutes: to.Ptr[int32](15),
						Protocol:             to.Ptr(armnetwork.TransportProtocolTCP),
					},
				}},
			LoadBalancingRules: []*armnetwork.LoadBalancingRule{
				{
					Name: to.Ptr("rulelb"),
					Properties: &armnetwork.LoadBalancingRulePropertiesFormat{
						BackendAddressPool: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/backendAddressPools/be-lb"),
						},
						BackendPort:      to.Ptr[int32](80),
						EnableFloatingIP: to.Ptr(true),
						FrontendIPConfiguration: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
						},
						FrontendPort:         to.Ptr[int32](80),
						IdleTimeoutInMinutes: to.Ptr[int32](15),
						LoadDistribution:     to.Ptr(armnetwork.LoadDistributionDefault),
						Probe: &armnetwork.SubResource{
							ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/probes/probe-lb"),
						},
						Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
					},
				}},
			OutboundRules: []*armnetwork.OutboundRule{},
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
				}},
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
