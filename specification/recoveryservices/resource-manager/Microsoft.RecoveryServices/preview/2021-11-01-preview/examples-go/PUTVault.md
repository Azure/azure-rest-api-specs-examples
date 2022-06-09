```go
package armrecoveryservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservices"
)

// x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/preview/2021-11-01-preview/examples/PUTVault.json
func ExampleVaultsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armrecoveryservices.NewVaultsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vault-name>",
		armrecoveryservices.Vault{
			Location: to.StringPtr("<location>"),
			Identity: &armrecoveryservices.IdentityData{
				Type: armrecoveryservices.ResourceIdentityType("SystemAssigned").ToPtr(),
			},
			Properties: &armrecoveryservices.VaultProperties{},
			SKU: &armrecoveryservices.SKU{
				Name: armrecoveryservices.SKUName("Standard").ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.VaultsClientCreateOrUpdateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservices%2Fv0.3.1/sdk/resourcemanager/recoveryservices/armrecoveryservices/README.md) on how to add the SDK to your project and authenticate.
