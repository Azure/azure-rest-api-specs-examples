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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateLinkServiceCreate.json
func ExamplePrivateLinkServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armnetwork.NewPrivateLinkServicesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<service-name>",
		armnetwork.PrivateLinkService{
			Location: to.Ptr("<location>"),
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
						Name: to.Ptr("<name>"),
						Properties: &armnetwork.PrivateLinkServiceIPConfigurationProperties{
							PrivateIPAddress:          to.Ptr("<private-ipaddress>"),
							PrivateIPAddressVersion:   to.Ptr(armnetwork.IPVersionIPv4),
							PrivateIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
							Subnet: &armnetwork.Subnet{
								ID: to.Ptr("<id>"),
							},
						},
					}},
				LoadBalancerFrontendIPConfigurations: []*armnetwork.FrontendIPConfiguration{
					{
						ID: to.Ptr("<id>"),
					}},
				Visibility: &armnetwork.PrivateLinkServicePropertiesVisibility{
					Subscriptions: []*string{
						to.Ptr("subscription1"),
						to.Ptr("subscription2"),
						to.Ptr("subscription3")},
				},
			},
		},
		&armnetwork.PrivateLinkServicesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fnetwork%2Farmnetwork%2Fv0.5.0/sdk/resourcemanager/network/armnetwork/README.md) on how to add the SDK to your project and authenticate.
