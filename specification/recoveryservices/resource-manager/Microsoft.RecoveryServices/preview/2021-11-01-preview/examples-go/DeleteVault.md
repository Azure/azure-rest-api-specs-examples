Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservices%2Fv0.3.1/sdk/resourcemanager/recoveryservices/armrecoveryservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armrecoveryservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices"
)

// x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/preview/2021-11-01-preview/examples/DeleteVault.json
func ExampleVaultsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservices.NewVaultsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<vault-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}
```
