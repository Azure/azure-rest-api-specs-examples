package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: 2026-08-01-preview/DeletedServerGet.json
func ExampleDeletedServersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDeletedServersClient().Get(ctx, "japaneast", "sqlcrudtest-d-1414", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armsql.DeletedServersClientGetResponse{
	// 	DeletedServer: armsql.DeletedServer{
	// 		Name: to.Ptr("sqlcrudtest-d-1414"),
	// 		Type: to.Ptr("Microsoft.Sql/deletedServers"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/locations/japaneast/deletedServers/sqlcrudtest-d-1414"),
	// 		Properties: &armsql.DeletedServerProperties{
	// 			DeletionTime: to.Ptr(time.Date(2017, time.June, 15, 11, 20, 0, 345000000, time.UTC)),
	// 			FullyQualifiedDomainName: to.Ptr("sqlcrudtest-d-1414.database.windows.net"),
	// 			OriginalID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default/providers/Microsoft.Sql/servers/sqlcrudtest-d-1414"),
	// 			Version: to.Ptr("12.0"),
	// 			OriginalResourceGroup: to.Ptr("Default"),
	// 			ScheduledPurgeTime: to.Ptr(time.Date(2017, time.June, 22, 11, 20, 0, 345000000, time.UTC)),
	// 		},
	// 	},
	// }
}
