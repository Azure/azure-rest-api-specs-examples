package armmysql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2020-01-01/examples/ServerKeyList.json
func ExampleServerKeysClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.ServerKeyListResult = armmysql.ServerKeyListResult{
		// 	Value: []*armmysql.ServerKey{
		// 		{
		// 			Name: to.Ptr("someVault_someKey_01234567890123456789012345678901"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/keys"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBforMySQL/servers/testserver/keys/someVault_someKey_01234567890123456789012345678901"),
		// 			Kind: to.Ptr("azurekeyvault"),
		// 			Properties: &armmysql.ServerKeyProperties{
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-12-01T00:00:00.000Z"); return t}()),
		// 				ServerKeyType: to.Ptr(armmysql.ServerKeyTypeAzureKeyVault),
		// 				URI: to.Ptr("https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901"),
		// 			},
		// 	}},
		// }
	}
}
