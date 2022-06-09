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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/PUTVault.json
func ExampleVaultsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armrecoveryservices.NewVaultsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<vault-name>",
		armrecoveryservices.Vault{
			Location: to.Ptr("<location>"),
			Identity: &armrecoveryservices.IdentityData{
				Type: to.Ptr(armrecoveryservices.ResourceIdentityTypeSystemAssigned),
			},
			Properties: &armrecoveryservices.VaultProperties{},
			SKU: &armrecoveryservices.SKU{
				Name: to.Ptr(armrecoveryservices.SKUNameStandard),
			},
		},
		&armrecoveryservices.VaultsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Frecoveryservices%2Farmrecoveryservices%2Fv0.5.0/sdk/resourcemanager/recoveryservices/armrecoveryservices/README.md) on how to add the SDK to your project and authenticate.
