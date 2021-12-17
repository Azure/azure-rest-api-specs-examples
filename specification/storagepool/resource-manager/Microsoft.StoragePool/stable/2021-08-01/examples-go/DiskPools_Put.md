Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragepool%2Farmstoragepool%2Fv0.1.0/sdk/resourcemanager/storagepool/armstoragepool/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_Put.json
func ExampleDiskPoolsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstoragepool.NewDiskPoolsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<disk-pool-name>",
		armstoragepool.DiskPoolCreate{
			Location: to.StringPtr("<location>"),
			Properties: &armstoragepool.DiskPoolCreateProperties{
				AvailabilityZones: []*string{
					to.StringPtr("1")},
				Disks: []*armstoragepool.Disk{
					{
						ID: to.StringPtr("<id>"),
					},
					{
						ID: to.StringPtr("<id>"),
					}},
				SubnetID: to.StringPtr("<subnet-id>"),
			},
			SKU: &armstoragepool.SKU{
				Name: to.StringPtr("<name>"),
				Tier: to.StringPtr("<tier>"),
			},
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
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
	log.Printf("DiskPool.ID: %s\n", *res.ID)
}
```
