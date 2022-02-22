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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateLinkServiceCreate.json
func ExamplePrivateLinkServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewPrivateLinkServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		armnetwork.PrivateLinkService{
			Location: to.StringPtr("<location>"),
			Properties: &armnetwork.PrivateLinkServiceProperties{
				AutoApproval: &armnetwork.PrivateLinkServicePropertiesAutoApproval{
					Subscriptions: []*string{
						to.StringPtr("subscription1"),
						to.StringPtr("subscription2")},
				},
				Fqdns: []*string{
					to.StringPtr("fqdn1"),
					to.StringPtr("fqdn2"),
					to.StringPtr("fqdn3")},
				IPConfigurations: []*armnetwork.PrivateLinkServiceIPConfiguration{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
							PrivateIPAddress:          to.StringPtr("<private-ipaddress>"),
							PrivateIPAddressVersion:   armnetwork.IPVersion("IPv4").ToPtr(),
							PrivateIPAllocationMethod: armnetwork.IPAllocationMethod("Static").ToPtr(),
							Subnet: &armnetwork.Subnet{
								ID: to.StringPtr("<id>"),
							},
						},
					}},
				LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
					{
						ID: to.StringPtr("<id>"),
					}},
				Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
					Subscriptions: []*string{
						to.StringPtr("subscription1"),
						to.StringPtr("subscription2"),
						to.StringPtr("subscription3")},
				},
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
	log.Printf("Response result: %#v\n", res.PrivateLinkServicesClientCreateOrUpdateResult)
}
```
