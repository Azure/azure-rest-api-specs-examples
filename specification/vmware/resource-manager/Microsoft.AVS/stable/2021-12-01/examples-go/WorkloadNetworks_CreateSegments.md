Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Favs%2Farmavs%2Fv0.1.0/sdk/resourcemanager/avs/armavs/README.md) on how to add the SDK to your project and authenticate.

```go
package armavs_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_CreateSegments.json
func ExampleWorkloadNetworksClient_BeginCreateSegments() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateSegments(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<segment-id>",
		armavs.WorkloadNetworkSegment{
			Properties: &armavs.WorkloadNetworkSegmentProperties{
				ConnectedGateway: to.StringPtr("<connected-gateway>"),
				DisplayName:      to.StringPtr("<display-name>"),
				Revision:         to.Int64Ptr(1),
				Subnet: &armavs.WorkloadNetworkSegmentSubnet{
					DhcpRanges: []*string{
						to.StringPtr("40.20.0.0-40.20.0.1")},
					GatewayAddress: to.StringPtr("<gateway-address>"),
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
	log.Printf("WorkloadNetworkSegment.ID: %s\n", *res.ID)
}
```
