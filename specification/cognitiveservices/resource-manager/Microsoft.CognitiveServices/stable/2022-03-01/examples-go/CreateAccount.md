Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcognitiveservices%2Farmcognitiveservices%2Fv0.6.0/sdk/resourcemanager/cognitiveservices/armcognitiveservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armcognitiveservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/CreateAccount.json
func ExampleAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewAccountsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armcognitiveservices.Account{
			Identity: &armcognitiveservices.Identity{
				Type: to.Ptr(armcognitiveservices.ResourceIdentityTypeSystemAssigned),
			},
			Kind:     to.Ptr("<kind>"),
			Location: to.Ptr("<location>"),
			Properties: &armcognitiveservices.AccountProperties{
				Encryption: &armcognitiveservices.Encryption{
					KeySource: to.Ptr(armcognitiveservices.KeySourceMicrosoftKeyVault),
					KeyVaultProperties: &armcognitiveservices.KeyVaultProperties{
						KeyName:     to.Ptr("<key-name>"),
						KeyVaultURI: to.Ptr("<key-vault-uri>"),
						KeyVersion:  to.Ptr("<key-version>"),
					},
				},
				UserOwnedStorage: []*armcognitiveservices.UserOwnedStorage{
					{
						ResourceID: to.Ptr("<resource-id>"),
					}},
			},
			SKU: &armcognitiveservices.SKU{
				Name: to.Ptr("<name>"),
			},
		},
		&armcognitiveservices.AccountsClientBeginCreateOptions{ResumeToken: ""})
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
