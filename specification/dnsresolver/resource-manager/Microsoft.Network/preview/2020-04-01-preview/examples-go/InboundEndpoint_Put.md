Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fdnsresolver%2Farmdnsresolver%2Fv0.4.0/sdk/resourcemanager/dnsresolver/armdnsresolver/README.md) on how to add the SDK to your project and authenticate.

```go
package armdnsresolver_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/dnsresolver/resource-manager/Microsoft.Network/preview/2020-04-01-preview/examples/InboundEndpoint_Put.json
func ExampleInboundEndpointsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armdnsresolver.NewInboundEndpointsClient("abdd4249-9f34-4cc6-8e42-c2e32110603e", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"sampleResourceGroup",
		"sampleDnsResolver",
		"sampleInboundEndpoint",
		armdnsresolver.InboundEndpoint{
			Location: to.Ptr("westus2"),
			Tags: map[string]*string{
				"key1": to.Ptr("value1"),
			},
			Properties: &armdnsresolver.InboundEndpointProperties{
				IPConfigurations: []*armdnsresolver.IPConfiguration{
					{
						PrivateIPAddress:          to.Ptr("255.255.255.255"),
						PrivateIPAllocationMethod: to.Ptr(armdnsresolver.IPAllocationMethodStatic),
						Subnet: &armdnsresolver.SubResource{
							ID: to.Ptr("/subscriptions/0403cfa9-9659-4f33-9f30-1f191c51d111/resourceGroups/sampleVnetResourceGroupName/providers/Microsoft.Network/virtualNetworks/sampleVirtualNetwork/subnets/sampleSubnet"),
						},
					}},
			},
		},
		&armdnsresolver.InboundEndpointsClientBeginCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
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
