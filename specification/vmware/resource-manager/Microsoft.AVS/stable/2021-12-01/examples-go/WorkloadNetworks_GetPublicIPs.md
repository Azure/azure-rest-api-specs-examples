Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Favs%2Farmavs%2Fv0.1.0/sdk/resourcemanager/avs/armavs/README.md) on how to add the SDK to your project and authenticate.

```go
package armavs_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs"
)

// x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/WorkloadNetworks_GetPublicIPs.json
func ExampleWorkloadNetworksClient_GetPublicIP() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armavs.NewWorkloadNetworksClient("<subscription-id>", cred, nil)
	res, err := client.GetPublicIP(ctx,
		"<resource-group-name>",
		"<private-cloud-name>",
		"<public-ipid>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("WorkloadNetworkPublicIP.ID: %s\n", *res.ID)
}
```
