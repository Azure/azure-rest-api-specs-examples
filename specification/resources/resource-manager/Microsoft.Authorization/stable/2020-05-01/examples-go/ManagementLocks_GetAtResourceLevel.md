Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fresources%2Farmlocks%2Fv0.1.1/sdk/resourcemanager/resources/armlocks/README.md) on how to add the SDK to your project and authenticate.

```go
package armlocks_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlocks"
)

// x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-05-01/examples/ManagementLocks_GetAtResourceLevel.json
func ExampleManagementLocksClient_GetAtResourceLevel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armlocks.NewManagementLocksClient("<subscription-id>", cred, nil)
	res, err := client.GetAtResourceLevel(ctx,
		"<resource-group-name>",
		"<resource-provider-namespace>",
		"<parent-resource-path>",
		"<resource-type>",
		"<resource-name>",
		"<lock-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ManagementLockObject.ID: %s\n", *res.ID)
}
```
