package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-04-15/examples/CosmosDBSqlClientEncryptionKeysList.json
func ExampleSQLResourcesClient_NewListClientEncryptionKeysPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLResourcesClient().NewListClientEncryptionKeysPager("rgName", "accountName", "databaseName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.ClientEncryptionKeysListResult = armcosmos.ClientEncryptionKeysListResult{
		// 	Value: []*armcosmos.ClientEncryptionKeyGetResults{
		// 		{
		// 			Name: to.Ptr("cekName1"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/clientEncryptionKey"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.DocumentDB/databaseAccounts/accountName/sqlDatabases/databaseName/clientEncryptionKeys/cekName1"),
		// 			Properties: &armcosmos.ClientEncryptionKeyGetProperties{
		// 				Resource: &armcosmos.ClientEncryptionKeyGetPropertiesResource{
		// 					EncryptionAlgorithm: to.Ptr("AEAD_AES_256_CBC_HMAC_SHA256"),
		// 					ID: to.Ptr("cekName1"),
		// 					KeyWrapMetadata: &armcosmos.KeyWrapMetadata{
		// 						Name: to.Ptr("customerManagedKey1"),
		// 						Type: to.Ptr("AzureKeyVault"),
		// 						Algorithm: to.Ptr("RSA-OAEP"),
		// 						Value: to.Ptr("AzureKeyVault Key URL for customerManagedKey1"),
		// 					},
		// 					WrappedDataEncryptionKey: []byte("U3dhZ2dlciByb2Nrcw=="),
		// 					Etag: to.Ptr("00000000-0000-0000-7a1f-bc0828e801d7"),
		// 					Rid: to.Ptr("nAMyAAAAAADPw1kKAgAAAA=="),
		// 					Ts: to.Ptr[float32](1626425552),
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("cekName2"),
		// 			Type: to.Ptr("Microsoft.DocumentDB/databaseAccounts/sqlDatabases/clientEncryptionKey"),
		// 			ID: to.Ptr("/subscriptions/subId/resourceGroups/rgName/providers/Microsoft.DocumentDB/databaseAccounts/accountName/sqlDatabases/databaseName/clientEncryptionKeys/cekName2"),
		// 			Properties: &armcosmos.ClientEncryptionKeyGetProperties{
		// 				Resource: &armcosmos.ClientEncryptionKeyGetPropertiesResource{
		// 					EncryptionAlgorithm: to.Ptr("AEAD_AES_256_CBC_HMAC_SHA256"),
		// 					ID: to.Ptr("cekName2"),
		// 					KeyWrapMetadata: &armcosmos.KeyWrapMetadata{
		// 						Name: to.Ptr("customerManagedKey2"),
		// 						Type: to.Ptr("AzureKeyVault"),
		// 						Algorithm: to.Ptr("RSA-OAEP"),
		// 						Value: to.Ptr("AzureKeyVault Key URL for customerManagedKey2"),
		// 					},
		// 					WrappedDataEncryptionKey: []byte("U3dhZ2dlciByb2Nrcw=="),
		// 					Etag: to.Ptr("00000000-0000-0000-7a21-7788a38c01d7"),
		// 					Rid: to.Ptr("nAMyAAAAAAAWWfxHAgAAAA=="),
		// 					Ts: to.Ptr[float32](1626425631),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
