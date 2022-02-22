Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurity%2Farmsecurity%2Fv0.3.1/sdk/resourcemanager/security/armsecurity/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurity_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
)

// x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-05-01-preview/examples/SoftwareInventories/GetSoftware_example.json
func ExampleSoftwareInventoriesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurity.NewSoftwareInventoriesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<resource-namespace>",
		"<resource-type>",
		"<resource-name>",
		"<software-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.SoftwareInventoriesClientGetResult)
}
```
