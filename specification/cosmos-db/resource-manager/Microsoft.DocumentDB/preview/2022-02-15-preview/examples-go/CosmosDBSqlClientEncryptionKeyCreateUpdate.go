package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-02-15-preview/examples/CosmosDBSqlClientEncryptionKeyCreateUpdate.json
func ExampleSQLResourcesClient_BeginCreateUpdateClientEncryptionKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcosmos.NewSQLResourcesClient("subId", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateUpdateClientEncryptionKey(ctx,
		"rgName",
		"accountName",
		"databaseName",
		"cekName",
		armcosmos.ClientEncryptionKeyCreateUpdateParameters{
			Properties: &armcosmos.ClientEncryptionKeyCreateUpdateProperties{
				Resource: &armcosmos.ClientEncryptionKeyResource{
					EncryptionAlgorithm: to.Ptr("AEAD_AES_256_CBC_HMAC_SHA256"),
					ID:                  to.Ptr("cekName"),
					KeyWrapMetadata: &armcosmos.KeyWrapMetadata{
						Name:      to.Ptr("customerManagedKey"),
						Type:      to.Ptr("AzureKeyVault"),
						Algorithm: to.Ptr("RSA-OAEP"),
						Value:     to.Ptr("AzureKeyVault Key URL"),
					},
					WrappedDataEncryptionKey: []byte("This is actually an array of bytes. This request/response is being presented as a string for readability in the example"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
