Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fhybridcompute%2Farmhybridcompute%2Fv0.4.0/sdk/resourcemanager/hybridcompute/armhybridcompute/README.md) on how to add the SDK to your project and authenticate.

```go
package armhybridcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2021-12-10-preview/examples/Machines_CreateOrUpdate.json
func ExampleMachinesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armhybridcompute.NewMachinesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<machine-name>",
		armhybridcompute.Machine{
			Location: to.Ptr("<location>"),
			Identity: &armhybridcompute.Identity{
				Type: to.Ptr("<type>"),
			},
			Properties: &armhybridcompute.MachineProperties{
				ClientPublicKey: to.Ptr("<client-public-key>"),
				LocationData: &armhybridcompute.LocationData{
					Name: to.Ptr("<name>"),
				},
				ParentClusterResourceID:    to.Ptr("<parent-cluster-resource-id>"),
				PrivateLinkScopeResourceID: to.Ptr("<private-link-scope-resource-id>"),
				VMID:                       to.Ptr("<vmid>"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
