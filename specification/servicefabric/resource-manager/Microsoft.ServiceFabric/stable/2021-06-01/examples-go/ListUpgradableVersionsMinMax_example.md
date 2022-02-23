Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fservicefabric%2Farmservicefabric%2Fv0.3.1/sdk/resourcemanager/servicefabric/armservicefabric/README.md) on how to add the SDK to your project and authenticate.

```go
package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ListUpgradableVersionsMinMax_example.json
func ExampleClustersClient_ListUpgradableVersions() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armservicefabric.NewClustersClient("<subscription-id>", cred, nil)
	res, err := client.ListUpgradableVersions(ctx,
		"<resource-group-name>",
		"<cluster-name>",
		&armservicefabric.ClustersClientListUpgradableVersionsOptions{VersionsDescription: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClustersClientListUpgradableVersionsResult)
}
```
