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

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_CreateVMGroups.json
func ExampleWorkloadNetworksClient_BeginCreateVMGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateVMGroup(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<vm-group-id>",
		armavs.WorkloadNetworkVMGroup{
			Properties: &armavs.WorkloadNetworkVMGroupProperties{
				DisplayName: to.StringPtr("<display-name>"),
				Members: []*string{
					to.StringPtr("564d43da-fefc-2a3b-1d92-42855622fa50")},
				Revision: to.Int64Ptr(1),
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
	log.Printf("WorkloadNetworkVMGroup.ID: %s\n", *res.ID)
}
```
