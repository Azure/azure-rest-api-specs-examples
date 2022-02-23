Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkeyvault%2Farmkeyvault%2Fv0.3.1/sdk/resourcemanager/keyvault/armkeyvault/README.md) on how to add the SDK to your project and authenticate.

```go
package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// x-ms-original-file: specification/keyvault/resource-manager/Microsoft.KeyVault/preview/2021-11-01-preview/examples/listPrivateLinkResources.json
func ExamplePrivateLinkResourcesClient_ListByVault() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armkeyvault.NewPrivateLinkResourcesClient("<subscription-id>", cred, nil)
	res, err := client.ListByVault(ctx,
		"<resource-group-name>",
		"<vault-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PrivateLinkResourcesClientListByVaultResult)
}
```
