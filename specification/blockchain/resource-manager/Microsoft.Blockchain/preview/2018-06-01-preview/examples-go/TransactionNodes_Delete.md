Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fblockchain%2Farmblockchain%2Fv0.4.0/sdk/resourcemanager/blockchain/armblockchain/README.md) on how to add the SDK to your project and authenticate.

```go
package armblockchain_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/blockchain/armblockchain"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/blockchain/resource-manager/Microsoft.Blockchain/preview/2018-06-01-preview/examples/TransactionNodes_Delete.json
func ExampleTransactionNodesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armblockchain.NewTransactionNodesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDelete(ctx,
		"<blockchain-member-name>",
		"<transaction-node-name>",
		"<resource-group-name>",
		&armblockchain.TransactionNodesClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
```
