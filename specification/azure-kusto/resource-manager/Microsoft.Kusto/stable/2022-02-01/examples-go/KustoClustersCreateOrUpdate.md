Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkusto%2Farmkusto%2Fv0.4.0/sdk/resourcemanager/kusto/armkusto/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoClustersCreateOrUpdate.json
func ExampleClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armkusto.NewClustersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		armkusto.Cluster{
			Location: to.Ptr("<location>"),
			Identity: &armkusto.Identity{
				Type: to.Ptr(armkusto.IdentityTypeSystemAssigned),
			},
			Properties: &armkusto.ClusterProperties{
				AllowedIPRangeList: []*string{
					to.Ptr("0.0.0.0/0")},
				EnableAutoStop:         to.Ptr(true),
				EnableDoubleEncryption: to.Ptr(false),
				EnablePurge:            to.Ptr(true),
				EnableStreamingIngest:  to.Ptr(true),
				PublicIPType:           to.Ptr(armkusto.PublicIPTypeDualStack),
				PublicNetworkAccess:    to.Ptr(armkusto.PublicNetworkAccessEnabled),
			},
			SKU: &armkusto.AzureSKU{
				Name:     to.Ptr(armkusto.AzureSKUNameStandardL8S),
				Capacity: to.Ptr[int32](2),
				Tier:     to.Ptr(armkusto.AzureSKUTierStandard),
			},
		},
		&armkusto.ClustersClientBeginCreateOrUpdateOptions{IfMatch: nil,
			IfNoneMatch: nil,
			ResumeToken: "",
		})
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
