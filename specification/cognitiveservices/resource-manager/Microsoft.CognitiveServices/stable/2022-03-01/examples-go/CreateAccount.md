Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcognitiveservices%2Farmcognitiveservices%2Fv0.4.0/sdk/resourcemanager/cognitiveservices/armcognitiveservices/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/CreateAccount.json
func ExampleAccountsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcognitiveservices.NewAccountsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreate(ctx,
		"<resource-group-name>",
		"<account-name>",
		armcognitiveservices.Account{
			Identity: &armcognitiveservices.Identity{
				Type: armcognitiveservices.ResourceIdentityTypeSystemAssigned.ToPtr(),
			},
			Kind:     to.StringPtr("<kind>"),
			Location: to.StringPtr("<location>"),
			Properties: &armcognitiveservices.AccountProperties{
				Encryption: &armcognitiveservices.Encryption{
					KeySource: armcognitiveservices.KeySource("Microsoft.KeyVault").ToPtr(),
					KeyVaultProperties: &armcognitiveservices.KeyVaultProperties{
						KeyName:     to.StringPtr("<key-name>"),
						KeyVaultURI: to.StringPtr("<key-vault-uri>"),
						KeyVersion:  to.StringPtr("<key-version>"),
					},
				},
				UserOwnedStorage: []*armcognitiveservices.UserOwnedStorage{
					{
						ResourceID: to.StringPtr("<resource-id>"),
					}},
			},
			SKU: &armcognitiveservices.SKU{
				Name: to.StringPtr("<name>"),
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
	log.Printf("Response result: %#v\n", res.AccountsClientCreateResult)
}
```
