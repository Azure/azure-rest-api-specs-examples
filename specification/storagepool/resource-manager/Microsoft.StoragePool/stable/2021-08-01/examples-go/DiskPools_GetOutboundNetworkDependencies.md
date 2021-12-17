Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fstoragepool%2Farmstoragepool%2Fv0.1.0/sdk/resourcemanager/storagepool/armstoragepool/README.md) on how to add the SDK to your project and authenticate.

```go
package armstoragepool_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagepool/armstoragepool"
)

// x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/DiskPools_GetOutboundNetworkDependencies.json
func ExampleDiskPoolsClient_ListOutboundNetworkDependenciesEndpoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armstoragepool.NewDiskPoolsClient("<subscription-id>", cred, nil)
	pager := client.ListOutboundNetworkDependenciesEndpoints("<resource-group-name>",
		"<disk-pool-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
	}
}
```
