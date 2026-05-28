package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/RecoverableDatabaseGetWithExpandEqualsKeys.json
func ExampleRecoverableDatabasesClient_Get_getsARecoverableDatabaseWithExpandEqualsKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRecoverableDatabasesClient().Get(ctx, "recoverabledatabasetest-6852", "recoverabledatabasetest-2080", "recoverabledatabasetest-9187", &armsql.RecoverableDatabasesClientGetOptions{
		Expand: to.Ptr("keys")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.RecoverableDatabasesClientGetResponse{
	// 	RecoverableDatabase: armsql.RecoverableDatabase{
	// 		Name: to.Ptr("recoverabledatabasetest-9187"),
	// 		Type: to.Ptr("Microsoft.Sql/servers/recoverableDatabases"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/recoverabledatabasetest-6852/providers/Microsoft.Sql/servers/recoverabledatabasetest-2080/recoverableDatabases/recoverabledatabasetest-9187"),
	// 		Properties: &armsql.RecoverableDatabaseProperties{
	// 			Edition: to.Ptr("Basic"),
	// 			Keys: map[string]*armsql.DatabaseKey{
	// 				"https://your-key-vault-name.vault.azure.net/yourKey/yourKeyVersion": &armsql.DatabaseKey{
	// 				},
	// 				"https://your-key-vault-name.vault.azure.net/yourKey2/yourKey2Version": &armsql.DatabaseKey{
	// 				},
	// 			},
	// 			LastAvailableBackupDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-26T01:06:29.78Z"); return t}()),
	// 			ServiceLevelObjective: to.Ptr("Basic"),
	// 		},
	// 	},
	// }
}
