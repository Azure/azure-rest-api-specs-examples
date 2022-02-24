Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmaps%2Farmmaps%2Fv0.2.1/sdk/resourcemanager/maps/armmaps/README.md) on how to add the SDK to your project and authenticate.

```go
package armmaps_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/maps/armmaps"
)

// x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/preview/2021-12-01-preview/examples/RegenerateKey.json
func ExampleAccountsClient_RegenerateKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmaps.NewAccountsClient("<subscription-id>", cred, nil)
	res, err := client.RegenerateKeys(ctx,
		"<resource-group-name>",
		"<account-name>",
		armmaps.KeySpecification{
			KeyType: armmaps.KeyType("primary").ToPtr(),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AccountsClientRegenerateKeysResult)
}
```
