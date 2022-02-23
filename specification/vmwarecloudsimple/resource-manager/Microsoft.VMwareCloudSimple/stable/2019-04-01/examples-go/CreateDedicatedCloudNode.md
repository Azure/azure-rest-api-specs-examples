Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvmwarecloudsimple%2Farmvmwarecloudsimple%2Fv0.2.1/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/CreateDedicatedCloudNode.json
func ExampleDedicatedCloudNodesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armvmwarecloudsimple.NewDedicatedCloudNodesClient("<subscription-id>",
		"<referer>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<dedicated-cloud-node-name>",
		armvmwarecloudsimple.DedicatedCloudNode{
			Location: to.StringPtr("<location>"),
			Properties: &armvmwarecloudsimple.DedicatedCloudNodeProperties{
				AvailabilityZoneID: to.StringPtr("<availability-zone-id>"),
				NodesCount:         to.Int32Ptr(1),
				PlacementGroupID:   to.StringPtr("<placement-group-id>"),
				PurchaseID:         to.StringPtr("<purchase-id>"),
				SKUDescription: &armvmwarecloudsimple.SKUDescription{
					Name: to.StringPtr("<name>"),
					ID:   to.StringPtr("<id>"),
				},
			},
			SKU: &armvmwarecloudsimple.SKU{
				Name: to.StringPtr("<name>"),
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
	log.Printf("Response result: %#v\n", res.DedicatedCloudNodesClientCreateOrUpdateResult)
}
```
