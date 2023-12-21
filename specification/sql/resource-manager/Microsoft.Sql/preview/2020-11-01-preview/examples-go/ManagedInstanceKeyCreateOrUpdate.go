package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedInstanceKeyCreateOrUpdate.json
func ExampleManagedInstanceKeysClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedInstanceKeysClient().BeginCreateOrUpdate(ctx, "sqlcrudtest-7398", "sqlcrudtest-4645", "someVault_someKey_01234567890123456789012345678901", armsql.ManagedInstanceKey{
		Properties: &armsql.ManagedInstanceKeyProperties{
			ServerKeyType: to.Ptr(armsql.ServerKeyTypeAzureKeyVault),
			URI:           to.Ptr("https://someVault.vault.azure.net/keys/someKey/01234567890123456789012345678901"),
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
	// res.ManagedInstanceKey = armsql.ManagedInstanceKey{
	// 	Name: to.Ptr("sqlcrudtest-4645"),
	// 	Type: to.Ptr("Microsoft.Sql/managedInstances/keys"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/managedInstances/sqlcrudtest-4645/keys/someVault_someKey_01234567890123456789012345678901"),
	// 	Kind: to.Ptr("azurekeyvault"),
	// 	Properties: &armsql.ManagedInstanceKeyProperties{
	// 		AutoRotationEnabled: to.Ptr(false),
	// 		CreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-01T00:00:00.000Z"); return t}()),
	// 		Thumbprint: to.Ptr("00112233445566778899AABBCCDDEEFFAABBCCDD"),
	// 	},
	// }
}
