Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkusto%2Farmkusto%2Fv0.1.0/sdk/resourcemanager/kusto/armkusto/README.md) on how to add the SDK to your project and authenticate.

```go
package armkusto_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kusto/armkusto"
)

// x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClustersCreateOrUpdate.json
func ExampleClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkusto.NewClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armkusto.Cluster{
			TrackedResource: armkusto.TrackedResource{
				Location: to.StringPtr("<location>"),
			},
			Identity: &armkusto.Identity{
				Type: armkusto.IdentityTypeSystemAssigned.ToPtr(),
			},
			Properties: &armkusto.ClusterProperties{
				AllowedIPRangeList: []*string{
					to.StringPtr("0.0.0.0/0")},
				EnableAutoStop:         to.BoolPtr(true),
				EnableDoubleEncryption: to.BoolPtr(false),
				EnablePurge:            to.BoolPtr(true),
				EnableStreamingIngest:  to.BoolPtr(true),
				PublicNetworkAccess:    armkusto.PublicNetworkAccessEnabled.ToPtr(),
			},
			SKU: &armkusto.AzureSKU{
				Name:     armkusto.AzureSKUNameStandardL8S.ToPtr(),
				Capacity: to.Int32Ptr(2),
				Tier:     armkusto.AzureSKUTierStandard.ToPtr(),
			},
		},
		&armkusto.ClustersBeginCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Cluster.ID: %s\n", *res.ID)
}
```
