Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Favs%2Farmavs%2Fv0.4.0/sdk/resourcemanager/avs/armavs/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_CreatePortMirroringProfiles.json
func ExampleWorkloadNetworksClient_BeginCreatePortMirroring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreatePortMirroring(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<port-mirroring-id>",
		armavs.WorkloadNetworkPortMirroring{
			Properties: &armavs.WorkloadNetworkPortMirroringProperties{
				Destination: to.Ptr("<destination>"),
				Direction:   to.Ptr(armavs.PortMirroringDirectionEnumBIDIRECTIONAL),
				DisplayName: to.Ptr("<display-name>"),
				Revision:    to.Ptr[int64](1),
				Source:      to.Ptr("<source>"),
			},
		},
		&armavs.WorkloadNetworksClientBeginCreatePortMirroringOptions{ResumeToken: ""})
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
