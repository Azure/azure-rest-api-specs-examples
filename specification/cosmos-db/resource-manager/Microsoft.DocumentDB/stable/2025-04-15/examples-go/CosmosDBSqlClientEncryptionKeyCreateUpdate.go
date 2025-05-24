package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/011ecc5633300a5eefe43dde748f269d39e96458/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2025-04-15/examples/CosmosDBSqlClientEncryptionKeyCreateUpdate.json
func ExampleSQLResourcesClient_BeginCreateUpdateClientEncryptionKey() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLResourcesClient().BeginCreateUpdateClientEncryptionKey(ctx, "rgName", "accountName", "databaseName", "cekName", armcosmos.ClientEncryptionKeyCreateUpdateParameters{
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
				WrappedDataEncryptionKey: []byte("U3dhZ2dlciByb2Nrcw=="),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClientEncryptionKeyGetResults = armcosmos.ClientEncryptionKeyGetResults{
	// 	Name: to.Ptr("cekName"),
	// 	Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/clientEncryptionKey"),
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.DocumentDB/databaseAccounts/accountName/sqlDatabases/databaseName/clientEncryptionKeys/cekName"),
	// 	Properties: &armcosmos.ClientEncryptionKeyGetProperties{
	// 		Resource: &armcosmos.ClientEncryptionKeyGetPropertiesResource{
	// 			EncryptionAlgorithm: to.Ptr("AEAD_AES_256_CBC_HMAC_SHA256"),
	// 			ID: to.Ptr("cekName"),
	// 			KeyWrapMetadata: &armcosmos.KeyWrapMetadata{
	// 				Name: to.Ptr("customerManagedKey"),
	// 				Type: to.Ptr("AzureKeyVault"),
	// 				Algorithm: to.Ptr("RSA-OAEP"),
	// 				Value: to.Ptr("AzureKeyVault Key URL"),
	// 			},
	// 			WrappedDataEncryptionKey: []byte("U3dhZ2dlciByb2Nrcw=="),
	// 			Etag: to.Ptr("00000000-0000-0000-7a1f-bc0828e801d7"),
	// 			Rid: to.Ptr("tNc4AAAAAAAQkjzWAgAAAA=="),
	// 			Ts: to.Ptr[float32](1626425552),
	// 		},
	// 	},
	// }
}
