Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvmwarecloudsimple%2Farmvmwarecloudsimple%2Fv0.4.0/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple/README.md) on how to add the SDK to your project and authenticate.

```go
package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/CreateDedicatedCloudNode.json
func ExampleDedicatedCloudNodesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armvmwarecloudsimple.NewDedicatedCloudNodesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<referer>",
		"<dedicated-cloud-node-name>",
		armvmwarecloudsimple.DedicatedCloudNode{
			Location: to.Ptr("<location>"),
			Properties: &armvmwarecloudsimple.DedicatedCloudNodeProperties{
				AvailabilityZoneID: to.Ptr("<availability-zone-id>"),
				NodesCount:         to.Ptr[int32](1),
				PlacementGroupID:   to.Ptr("<placement-group-id>"),
				PurchaseID:         to.Ptr("<purchase-id>"),
				SKUDescription: &armvmwarecloudsimple.SKUDescription{
					Name: to.Ptr("<name>"),
					ID:   to.Ptr("<id>"),
				},
			},
			SKU: &armvmwarecloudsimple.SKU{
				Name: to.Ptr("<name>"),
			},
		},
		&armvmwarecloudsimple.DedicatedCloudNodesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
