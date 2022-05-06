Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcosmos%2Farmcosmos%2Fv0.5.0/sdk/resourcemanager/cosmos/armcosmos/README.md) on how to add the SDK to your project and authenticate.

```go
package armcosmos_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/CosmosDBSqlClientEncryptionKeyCreateUpdate.json
func ExampleSQLResourcesClient_BeginCreateUpdateClientEncryptionKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcosmos.NewSQLResourcesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateUpdateClientEncryptionKey(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<database-name>",
		"<client-encryption-key-name>",
		armcosmos.ClientEncryptionKeyCreateUpdateParameters{
			Properties: &armcosmos.ClientEncryptionKeyCreateUpdateProperties{
				Resource: &armcosmos.ClientEncryptionKeyResource{
					EncryptionAlgorithm: to.Ptr("<encryption-algorithm>"),
					ID:                  to.Ptr("<id>"),
					KeyWrapMetadata: &armcosmos.KeyWrapMetadata{
						Name:      to.Ptr("<name>"),
						Type:      to.Ptr("<type>"),
						Algorithm: to.Ptr("<algorithm>"),
						Value:     to.Ptr("<value>"),
					},
					WrappedDataEncryptionKey: []byte("This is actually an array of bytes. This request/response is being presented as a string for readability in the example"),
				},
			},
		},
		&armcosmos.SQLResourcesClientBeginCreateUpdateClientEncryptionKeyOptions{ResumeToken: ""})
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
