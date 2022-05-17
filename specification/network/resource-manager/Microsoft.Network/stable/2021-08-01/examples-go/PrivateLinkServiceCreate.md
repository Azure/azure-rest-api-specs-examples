Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv1.0.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.

```go
package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/PrivateLinkServiceCreate.json
func ExamplePrivateLinkServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armnetwork.NewPrivateLinkServicesClient("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"rg1",
		"testPls",
		armnetwork.PrivateLinkService{
			Location: to.Ptr("eastus"),
			Properties: &armnetwork.PrivateLinkServiceProperties{
				AutoApproval: &armnetwork.PrivateLinkServicePropertiesAutoApproval{
					Subscriptions: []*string{
						to.Ptr("subscription1"),
						to.Ptr("subscription2")},
				},
				Fqdns: []*string{
					to.Ptr("fqdn1"),
					to.Ptr("fqdn2"),
					to.Ptr("fqdn3")},
				IPConfigurations: []*armnetwork.PrivateLinkServiceIPConfiguration{
					{
						Name: to.Ptr("fe-lb"),
						Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
							PrivateIPAddress:          to.Ptr("10.0.1.4"),
							PrivateIPAddressVersion:   to.Ptr(armnetwork.IPVersionIPv4),
							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
							Subnet: &armnetwork.Subnet{
								ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnetlb/subnets/subnetlb"),
							},
						},
					}},
				LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
					{
						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/loadBalancers/lb/frontendIPConfigurations/fe-lb"),
					}},
				Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
					Subscriptions: []*string{
						to.Ptr("subscription1"),
						to.Ptr("subscription2"),
						to.Ptr("subscription3")},
				},
			},
		},
		nil)
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
```
