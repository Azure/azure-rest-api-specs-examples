Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.5.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LoadBalancerCreate.json
func ExampleLoadBalancersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewLoadBalancersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<load-balancer-name>",
		armnetwork.LoadBalancer{
			Location: to.Ptr("<location>"),
			Properties: &armnetwork.LoadBalancerPropertiesFormat{
				BackendAddressPools: []*armnetwork.BackendAddressPool{
					{
						Name:       to.Ptr("<name>"),
						Properties: &armnetwork.BackendAddressPoolPropertiesFormat{},
					}},
				FrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.FrontendIPConfigurationPropertiesFormat{
							Subnet: &armnetwork.Subnet{
								ID: to.Ptr("<id>"),
							},
						},
					}},
				InboundNatPools: []*armnetwork.InboundNatPool{},
				InboundNatRules: []*armnetwork.InboundNatRule{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.InboundNatRulePropertiesFormat{
							BackendPort:      to.Ptr[int32](3389),
							EnableFloatingIP: to.Ptr(true),
							EnableTCPReset:   to.Ptr(false),
							FrontendIPConfiguration: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							FrontendPort:         to.Ptr[int32](3389),
							IdleTimeoutInMinutes: to.Ptr[int32](15),
							Protocol:             to.Ptr(armnetwork.TransportProtocolTCP),
						},
					}},
				LoadBalancingRules: []*armnetwork.LoadBalancingRule{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.LoadBalancingRulePropertiesFormat{
							BackendAddressPool: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							BackendPort:      to.Ptr[int32](80),
							EnableFloatingIP: to.Ptr(true),
							EnableTCPReset:   to.Ptr(false),
							FrontendIPConfiguration: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							FrontendPort:         to.Ptr[int32](80),
							IdleTimeoutInMinutes: to.Ptr[int32](15),
							LoadDistribution:     to.Ptr(armnetwork.LoadDistributionDefault),
							Probe: &armnetwork.SubResource{
								ID: to.Ptr("<id>"),
							},
							Protocol: to.Ptr(armnetwork.TransportProtocolTCP),
						},
					}},
				Probes: []*armnetwork.Probe{
					{
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.ProbePropertiesFormat{
							IntervalInSeconds: to.Ptr[int32](15),
							NumberOfProbes:    to.Ptr[int32](2),
							Port:              to.Ptr[int32](80),
							RequestPath:       to.Ptr("<request-path>"),
							Protocol:          to.Ptr(armnetwork.ProbeProtocolHTTP),
						},
					}},
			},
		},
		&armnetwork.LoadBalancersClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
