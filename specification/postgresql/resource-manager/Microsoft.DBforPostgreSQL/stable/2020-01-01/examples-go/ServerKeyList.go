package armpostgresql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2020-01-01/examples/ServerKeyList.json
func ExampleServerKeysClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServerKeysClient().NewListPager("testrg", "testserver", nil)
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
		// page.ServerKeyListResult = armpostgresql.ServerKeyListResult{
		// 	Value: []*armpostgresql.ServerKey{
		// 		{
		// 			Name: to.Ptr("someVault_someKey_01234567890123456789012345678901"),
		// 			Type: to.Ptr("Microsoft.DBforPostgreSQL/servers/keys"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforPostgreSQL/servers/testserver/keys/someVault_someKey_01234567890123456789012345678901"),
		// 			Kind: to.Ptr("azurekeyvault"),
		// 			Properties: &armpostgresql.ServerKeyProperties{
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-01T00:00:00.000Z"); return t}()),
		// 				ServerKeyType: to.Ptr(armpostgresql.ServerKeyTypeAzureKeyVault),
		// 				URI: to.Ptr("https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901"),
		// 			},
		// 	}},
		// }
	}
}
