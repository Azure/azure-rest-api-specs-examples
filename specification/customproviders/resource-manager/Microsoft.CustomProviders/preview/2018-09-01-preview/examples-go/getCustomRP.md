Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcustomproviders%2Farmcustomproviders%2Fv0.1.0/sdk/resourcemanager/customproviders/armcustomproviders/README.md) on how to add the SDK to your project and authenticate.

```go
package armcustomproviders_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customproviders/armcustomproviders"
)

// x-ms-original-file: specification/customproviders/resource-manager/Microsoft.CustomProviders/preview/2018-09-01-preview/examples/getCustomRP.json
func ExampleCustomResourceProviderClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcustomproviders.NewCustomResourceProviderClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<resource-provider-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CustomRPManifest.ID: %s\n", *res.ID)
}
```
