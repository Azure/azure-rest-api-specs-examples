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

// x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/LocalNetworkGatewayCreate.json
func ExampleLocalNetworkGatewaysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armnetwork.NewLocalNetworkGatewaysClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<local-network-gateway-name>",
		armnetwork.LocalNetworkGateway{
			Location: to.StringPtr("<location>"),
			Properties: &armnetwork.LocalNetworkGatewayPropertiesFormat{
				Fqdn:             to.StringPtr("<fqdn>"),
				GatewayIPAddress: to.StringPtr("<gateway-ipaddress>"),
				LocalNetworkAddressSpace: &armnetwork.AddressSpace{
					AddressPrefixes: []*string{
						to.StringPtr("10.1.0.0/16")},
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
	log.Printf("Response result: %#v\n", res.LocalNetworkGatewaysClientCreateOrUpdateResult)
}
```
