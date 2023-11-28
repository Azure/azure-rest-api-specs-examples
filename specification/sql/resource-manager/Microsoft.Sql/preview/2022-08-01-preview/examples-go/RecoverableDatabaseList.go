package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c78b5d8bd3aff2d82a5f034d9164b1a9ac030e09/specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/RecoverableDatabaseList.json
func ExampleRecoverableDatabasesClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRecoverableDatabasesClient().NewListByServerPager("recoverabledatabasetest-1234", "recoverabledatabasetest-7177", nil)
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
		// page.RecoverableDatabaseListResult = armsql.RecoverableDatabaseListResult{
		// 	Value: []*armsql.RecoverableDatabase{
		// 		{
		// 			Name: to.Ptr("recoverabledatabasetest-1235"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/recoverableDatabases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/recoverabledatabasetest-1234/providers/Microsoft.Sql/servers/recoverabledatabasetest-7177/recoverableDatabases/recoverabledatabasetest-1235"),
		// 			Properties: &armsql.RecoverableDatabaseProperties{
		// 				Edition: to.Ptr("Standard"),
		// 				LastAvailableBackupDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-26T01:06:29.780Z"); return t}()),
		// 				ServiceLevelObjective: to.Ptr("S0"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("recoverabledatabasetest-9231"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/recoverableDatabases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/recoverabledatabasetest-1234/providers/Microsoft.Sql/servers/recoverabledatabasetest-7177/recoverableDatabases/recoverabledatabasetest-9231"),
		// 			Properties: &armsql.RecoverableDatabaseProperties{
		// 				Edition: to.Ptr("Premium"),
		// 				LastAvailableBackupDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-26T03:20:31.780Z"); return t}()),
		// 				ServiceLevelObjective: to.Ptr("P1"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("recoverabledatabasetest-0342"),
		// 			Type: to.Ptr("Microsoft.Sql/servers/recoverabledatabases"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/recoverabledatabasetest-1234/providers/Microsoft.Sql/servers/recoverabledatabasetest-7177/recoverabledatabases/recoverabledatabasetest-0342"),
		// 			Properties: &armsql.RecoverableDatabaseProperties{
		// 				Edition: to.Ptr("Basic"),
		// 				LastAvailableBackupDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-24T02:06:18.780Z"); return t}()),
		// 				ServiceLevelObjective: to.Ptr("Basic"),
		// 			},
		// 	}},
		// }
	}
}
