```go
package armblockchain_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// x-ms-original-file: specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/TransactionNodes_ListRegenerateApiKeys.json
func ExampleTransactionNodesClient_ListRegenerateAPIKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armblockchain.NewTransactionNodesClient("<subscription-id>", cred, nil)
	_, err = client.ListRegenerateAPIKeys(ctx,
		"<blockchain-member-name>",
		"<transaction-node-name>",
		"<resource-group-name>",
		&armblockchain.TransactionNodesListRegenerateAPIKeysOptions{APIKey: &armblockchain.APIKey{
			KeyName: to.StringPtr("<key-name>"),
		},
		})
	if err != nil {
		log.Fatal(err)
	}
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fblockchain%2Farmblockchain%2Fv0.1.0/sdk/resourcemanager/blockchain/armblockchain/README.md) on how to add the SDK to your project and authenticate.
