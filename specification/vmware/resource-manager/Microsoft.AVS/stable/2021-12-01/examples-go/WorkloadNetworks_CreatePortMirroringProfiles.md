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

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_CreatePortMirroringProfiles.json
func ExampleWorkloadNetworksClient_BeginCreatePortMirroring() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreatePortMirroring(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<port-mirroring-id>",
		armavs.WorkloadNetworkPortMirroring{
			Properties: &armavs.WorkloadNetworkPortMirroringProperties{
				Destination: to.StringPtr("<destination>"),
				Direction:   armavs.PortMirroringDirectionEnumBIDIRECTIONAL.ToPtr(),
				DisplayName: to.StringPtr("<display-name>"),
				Revision:    to.Int64Ptr(1),
				Source:      to.StringPtr("<source>"),
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
	log.Printf("WorkloadNetworkPortMirroring.ID: %s\n", *res.ID)
}
```
