package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceKeyList.json
func ExampleManagedInstanceKeysClient_NewListByInstancePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedInstanceKeysClient().NewListByInstancePager("sqlcrudtest-7398", "sqlcrudtest-4645", &armsql.ManagedInstanceKeysClientListByInstanceOptions{Filter: nil})
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
		// page.ManagedInstanceKeyListResult = armsql.ManagedInstanceKeyListResult{
		// 	Value: []*armsql.ManagedInstanceKey{
		// 		{
		// 			Name: to.Ptr("someVault_someKey_01234567890123456789012345678901"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/keys"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/managedInstances/sqlcrudtest-4645/keys/someVault_someKey_01234567890123456789012345678901"),
		// 			Kind: to.Ptr("azurekeyvault"),
		// 			Properties: &armsql.ManagedInstanceKeyProperties{
		// 				AutoRotationEnabled: to.Ptr(false),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-15T00:00:00.000Z"); return t}()),
		// 				Thumbprint: to.Ptr("00112233445566778899AABBCCDDEEFFAABBCCDD"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myVault_myKey_11111111111111111111111111111111"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/keys"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/managedInstances/sqlcrudtest-4645/keys/myVault_myKey_11111111111111111111111111111111"),
		// 			Kind: to.Ptr("azurekeyvault"),
		// 			Properties: &armsql.ManagedInstanceKeyProperties{
		// 				AutoRotationEnabled: to.Ptr(false),
		// 				CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-15T00:00:00.000Z"); return t}()),
		// 				Thumbprint: to.Ptr("AAAAAAAAAAAAAAABBBBBBBBBBBBBBBBBBBBBBBBB"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("ServiceManaged"),
		// 			Type: to.Ptr("Microsoft.Sql/managedInstances/keys"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/managedInstances/sqlcrudtest-4645/keys/ServiceManaged"),
		// 			Kind: to.Ptr("servicemanaged"),
		// 			Properties: &armsql.ManagedInstanceKeyProperties{
		// 			},
		// 	}},
		// }
	}
}
