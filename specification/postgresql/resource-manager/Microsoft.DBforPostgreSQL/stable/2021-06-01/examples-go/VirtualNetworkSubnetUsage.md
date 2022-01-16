Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpostgresql%2Farmpostgresqlflexibleservers%2Fv0.3.0/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/README.md) on how to add the SDK to your project and authenticate.

```go
package armpostgresqlflexibleservers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers"
)

// x-ms-original-file: specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/examples/VirtualNetworkSubnetUsage.json
func ExampleVirtualNetworkSubnetUsageClient_Execute() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpostgresqlflexibleservers.NewVirtualNetworkSubnetUsageClient("<subscription-id>", cred, nil)
	res, err := client.Execute(ctx,
		"<location-name>",
		armpostgresqlflexibleservers.VirtualNetworkSubnetUsageParameter{
			VirtualNetworkArmResourceID: to.StringPtr("<virtual-network-arm-resource-id>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VirtualNetworkSubnetUsageClientExecuteResult)
}
```
