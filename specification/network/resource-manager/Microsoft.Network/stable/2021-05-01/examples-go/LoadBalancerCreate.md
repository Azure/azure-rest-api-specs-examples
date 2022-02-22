Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.3.1/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LoadBalancerCreate.json
func ExampleLoadBalancersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewLoadBalancersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<load-balancer-name>",
		armnetwork.LoadBalancer{
			Location: to.StringPtr("<location>"),
			Properties: &armnetwork.LoadBalancerPropertiesFormat{
				BackendAddressPools: []*armnetwork.BackendAddressPool{
					{
						Name:       to.StringPtr("<name>"),
						Properties: &armnetwork.BackendAddressPoolPropertiesFormat{},
					}},
				FrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.FrontendIPConfigurationPropertiesFormat{
							Subnet: &armnetwork.Subnet{
								ID: to.StringPtr("<id>"),
							},
						},
					}},
				InboundNatPools: []*armnetwork.InboundNatPool{},
				InboundNatRules: []*armnetwork.InboundNatRule{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.InboundNatRulePropertiesFormat{
							BackendPort:      to.Int32Ptr(3389),
							EnableFloatingIP: to.BoolPtr(true),
							EnableTCPReset:   to.BoolPtr(false),
							FrontendIPConfiguration: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							FrontendPort:         to.Int32Ptr(3389),
							IdleTimeoutInMinutes: to.Int32Ptr(15),
							Protocol:             armnetwork.TransportProtocol("Tcp").ToPtr(),
						},
					}},
				LoadBalancingRules: []*armnetwork.LoadBalancingRule{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.LoadBalancingRulePropertiesFormat{
							BackendAddressPool: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							BackendPort:      to.Int32Ptr(80),
							EnableFloatingIP: to.BoolPtr(true),
							EnableTCPReset:   to.BoolPtr(false),
							FrontendIPConfiguration: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							FrontendPort:         to.Int32Ptr(80),
							IdleTimeoutInMinutes: to.Int32Ptr(15),
							LoadDistribution:     armnetwork.LoadDistribution("Default").ToPtr(),
							Probe: &armnetwork.SubResource{
								ID: to.StringPtr("<id>"),
							},
							Protocol: armnetwork.TransportProtocol("Tcp").ToPtr(),
						},
					}},
				Probes: []*armnetwork.Probe{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.ProbePropertiesFormat{
							IntervalInSeconds: to.Int32Ptr(15),
							NumberOfProbes:    to.Int32Ptr(2),
							Port:              to.Int32Ptr(80),
							RequestPath:       to.StringPtr("<request-path>"),
							Protocol:          armnetwork.ProbeProtocol("Http").ToPtr(),
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.LoadBalancersClientCreateOrUpdateResult)
}
```
