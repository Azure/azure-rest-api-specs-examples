Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fkeyvault%2Farmkeyvault%2Fv1.0.0/sdk/resourcemanager/keyvault/armkeyvault/README.md) on how to add the SDK to your project and authenticate.

```go
package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/deleteVault.json
func ExampleVaultsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armkeyvault.NewVaultsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = client.Delete(ctx,
		"sample-resource-group",
		"sample-vault",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
```
