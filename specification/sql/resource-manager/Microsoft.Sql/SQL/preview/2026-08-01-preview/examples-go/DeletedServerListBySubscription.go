package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2026-08-01-preview/DeletedServerListBySubscription.json
func ExampleDeletedServersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDeletedServersClient().NewListPager(nil)
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
		// page = armsql.DeletedServersClientListResponse{
		// 	DeletedServerListResult: armsql.DeletedServerListResult{
		// 		Value: []*armsql.DeletedServer{
		// 			{
		// 				Name: to.Ptr("sqlcrudtest-d-1414"),
		// 				Type: to.Ptr("Microsoft.Sql/deletedServers"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/deletedServers/sqlcrudtest-d-1414"),
		// 				Properties: &armsql.DeletedServerProperties{
		// 					DeletionTime: to.Ptr(time.Date(2017, time.June, 15, 20, 20, 0, 345000000, time.UTC)),
		// 					FullyQualifiedDomainName: to.Ptr("sqlcrudtest-d-1414.database.windows.net"),
		// 					OriginalID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sqlcrudtest-d-1414"),
		// 					Version: to.Ptr("12.0"),
		// 					OriginalResourceGroup: to.Ptr("Default"),
		// 					ScheduledPurgeTime: to.Ptr(time.Date(2017, time.June, 20, 10, 10, 0, 678000000, time.UTC)),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("sqlcrudtest-d-2424"),
		// 				Type: to.Ptr("Microsoft.Sql/deletedServers"),
		// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/deletedServers/sqlcrudtest-d-2424"),
		// 				Properties: &armsql.DeletedServerProperties{
		// 					DeletionTime: to.Ptr(time.Date(2017, time.June, 13, 10, 10, 0, 678000000, time.UTC)),
		// 					FullyQualifiedDomainName: to.Ptr("sqlcrudtest-d-2424.database.windows.net"),
		// 					OriginalID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sqlcrudtest-d-2424"),
		// 					Version: to.Ptr("12.0"),
		// 					OriginalResourceGroup: to.Ptr("Default"),
		// 					ScheduledPurgeTime: to.Ptr(time.Date(2017, time.June, 20, 10, 10, 0, 678000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
