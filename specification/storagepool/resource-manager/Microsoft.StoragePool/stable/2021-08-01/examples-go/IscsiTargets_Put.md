Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragepool%2Farmstoragepool%2Fv0.2.1/sdk/resourcemanager/storagepool/armstoragepool/README.md) on how to add the SDK to your project and authenticate.

```go
package armstoragepool_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool"
)

// x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/IscsiTargets_Put.json
func ExampleIscsiTargetsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstoragepool.NewIscsiTargetsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<disk-pool-name>",
		"<iscsi-target-name>",
		armstoragepool.IscsiTargetCreate{
			Properties: &armstoragepool.IscsiTargetCreateProperties{
				ACLMode: armstoragepool.IscsiTargetACLMode("Dynamic").ToPtr(),
				Luns: []*armstoragepool.IscsiLun{
					{
						Name:                       to.StringPtr("<name>"),
						ManagedDiskAzureResourceID: to.StringPtr("<managed-disk-azure-resource-id>"),
					}},
				TargetIqn: to.StringPtr("<target-iqn>"),
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
	log.Printf("Response result: %#v\n", res.IscsiTargetsClientCreateOrUpdateResult)
}
```
