package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2025-02-01-preview/DatabaseRestorePointsListByDatabase.json
func ExampleRestorePointsClient_NewListByDatabasePager_listDatabaseRestorePoints() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRestorePointsClient().NewListByDatabasePager("sqlcrudtest-6730", "sqlcrudtest-9007", "3481", nil)
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
		// page = armsql.RestorePointsClientListByDatabaseResponse{
		// 	RestorePointListResult: armsql.RestorePointListResult{
		// 		Value: []*armsql.RestorePoint{
		// 			{
		// 				Name: to.Ptr("ContinuousRestorePoint"),
		// 				Type: to.Ptr("Microsoft.Sql/servers/databases/restorePoints"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-6730/providers/Microsoft.Sql/servers/sqlcrudtest-9007/databases/3481/restorepoints/ContinuousRestorePoint"),
		// 				Location: to.Ptr("West US"),
		// 				Properties: &armsql.RestorePointProperties{
		// 					EarliestRestoreDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-12T00:00:00Z"); return t}()),
		// 					RestorePointType: to.Ptr(armsql.RestorePointTypeCONTINUOUS),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
