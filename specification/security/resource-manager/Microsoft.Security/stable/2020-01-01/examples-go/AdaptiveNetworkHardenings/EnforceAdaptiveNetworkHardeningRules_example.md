Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.4.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AdaptiveNetworkHardenings/EnforceAdaptiveNetworkHardeningRules_example.json
func ExampleAdaptiveNetworkHardeningsClient_BeginEnforce() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewAdaptiveNetworkHardeningsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginEnforce(ctx,
		"<resource-group-name>",
		"<resource-namespace>",
		"<resource-type>",
		"<resource-name>",
		"<adaptive-network-hardening-resource-name>",
		armsecurity.AdaptiveNetworkHardeningEnforceRequest{
			NetworkSecurityGroups: []*string{
				to.StringPtr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1"),
				to.StringPtr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg2/providers/Microsoft.Network/networkSecurityGroups/nsg2")},
			Rules: []*armsecurity.Rule{
				{
					Name:            to.StringPtr("<name>"),
					DestinationPort: to.Int32Ptr(3389),
					Direction:       armsecurity.Direction("Inbound").ToPtr(),
					IPAddresses: []*string{
						to.StringPtr("100.10.1.1"),
						to.StringPtr("200.20.2.2"),
						to.StringPtr("81.199.3.0/24")},
					Protocols: []*armsecurity.TransportProtocol{
						armsecurity.TransportProtocol("TCP").ToPtr()},
				},
				{
					Name:            to.StringPtr("<name>"),
					DestinationPort: to.Int32Ptr(22),
					Direction:       armsecurity.Direction("Inbound").ToPtr(),
					IPAddresses:     []*string{},
					Protocols: []*armsecurity.TransportProtocol{
						armsecurity.TransportProtocol("TCP").ToPtr()},
				}},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
```
