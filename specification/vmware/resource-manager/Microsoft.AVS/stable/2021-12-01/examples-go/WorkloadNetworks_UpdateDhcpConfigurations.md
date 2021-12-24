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

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_UpdateDhcpConfigurations.json
func ExampleWorkloadNetworksClient_BeginUpdateDhcp() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdateDhcp(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<dhcp-id>",
		armavs.WorkloadNetworkDhcp{
			Properties: &armavs.WorkloadNetworkDhcpServer{
				WorkloadNetworkDhcpEntity: armavs.WorkloadNetworkDhcpEntity{
					DhcpType: armavs.DhcpTypeEnumSERVER.ToPtr(),
					Revision: to.Int64Ptr(1),
				},
				LeaseTime:     to.Int64Ptr(86400),
				ServerAddress: to.StringPtr("<server-address>"),
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
	log.Printf("WorkloadNetworkDhcp.ID: %s\n", *res.ID)
}
```
