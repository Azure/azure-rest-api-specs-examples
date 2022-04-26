Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.6.0/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AdaptiveNetworkHardenings/EnforceAdaptiveNetworkHardeningRules_example.json
func ExampleAdaptiveNetworkHardeningsClient_BeginEnforce() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armsecurity.NewAdaptiveNetworkHardeningsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginEnforce(ctx,
		"<resource-group-name>",
		"<resource-namespace>",
		"<resource-type>",
		"<resource-name>",
		"<adaptive-network-hardening-resource-name>",
		armsecurity.AdaptiveNetworkHardeningEnforceRequest{
			NetworkSecurityGroups: []*string{
				to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/nsg1"),
				to.Ptr("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/rg2/providers/Microsoft.Network/networkSecurityGroups/nsg2")},
			Rules: []*armsecurity.Rule{
				{
					Name:            to.Ptr("<name>"),
					DestinationPort: to.Ptr[int32](3389),
					Direction:       to.Ptr(armsecurity.DirectionInbound),
					IPAddresses: []*string{
						to.Ptr("100.10.1.1"),
						to.Ptr("200.20.2.2"),
						to.Ptr("81.199.3.0/24")},
					Protocols: []*armsecurity.TransportProtocol{
						to.Ptr(armsecurity.TransportProtocolTCP)},
				},
				{
					Name:            to.Ptr("<name>"),
					DestinationPort: to.Ptr[int32](22),
					Direction:       to.Ptr(armsecurity.DirectionInbound),
					IPAddresses:     []*string{},
					Protocols: []*armsecurity.TransportProtocol{
						to.Ptr(armsecurity.TransportProtocolTCP)},
				}},
		},
		&armsecurity.AdaptiveNetworkHardeningsClientBeginEnforceOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```
