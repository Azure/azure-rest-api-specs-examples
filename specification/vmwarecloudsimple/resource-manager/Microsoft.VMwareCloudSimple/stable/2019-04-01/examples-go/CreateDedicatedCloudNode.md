```go
package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/CreateDedicatedCloudNode.json
func ExampleDedicatedCloudNodesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armvmwarecloudsimple.NewDedicatedCloudNodesClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"myResourceGroup",
		"https://management.azure.com/",
		"myNode",
		armvmwarecloudsimple.DedicatedCloudNode{
			Location: to.Ptr("westus"),
			Properties: &armvmwarecloudsimple.DedicatedCloudNodeProperties{
				AvailabilityZoneID: to.Ptr("az1"),
				NodesCount:         to.Ptr[int32](1),
				PlacementGroupID:   to.Ptr("n1"),
				PurchaseID:         to.Ptr("56acbd46-3d36-4bbf-9b08-57c30fdf6932"),
				SKUDescription: &armvmwarecloudsimple.SKUDescription{
					Name: to.Ptr("CS28-Node"),
					ID:   to.Ptr("general"),
				},
			},
			SKU: &armvmwarecloudsimple.SKU{
				Name: to.Ptr("VMware_CloudSimple_CS28"),
			},
		},
		nil)
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fvmwarecloudsimple%2Farmvmwarecloudsimple%2Fv1.0.0/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple/README.md) on how to add the SDK to your project and authenticate.
