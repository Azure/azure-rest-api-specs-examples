Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fappconfiguration%2Farmappconfiguration%2Fv0.1.0/sdk/resourcemanager/appconfiguration/armappconfiguration/README.md) on how to add the SDK to your project and authenticate.

```go
package armappconfiguration_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
)

// x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-03-01-preview/examples/ConfigurationStoresGet.json
func ExampleConfigurationStoresClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappconfiguration.NewConfigurationStoresClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<config-store-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ConfigurationStore.ID: %s\n", *res.ID)
}
```
