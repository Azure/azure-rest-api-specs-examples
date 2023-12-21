package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/01e99457ccf5613a95d5b2960d31a12f84018863/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/DataWarehouseRestorePointsGet.json
func ExampleRestorePointsClient_Get_getsADatawarehouseDatabaseRestorePoint() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRestorePointsClient().Get(ctx, "Default-SQL-SouthEastAsia", "testserver", "testDatabase", "131546477590000000", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RestorePoint = armsql.RestorePoint{
	// 	Name: to.Ptr("131546477590000000"),
	// 	Type: to.Ptr("Microsoft.Sql/servers/databases/restorePoints"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Sql/servers/testserver/databases/testDatabase/restorePoints/131546477590000000"),
	// 	Location: to.Ptr("japaneast"),
	// 	Properties: &armsql.RestorePointProperties{
	// 		RestorePointCreationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-10T08:00:00.000Z"); return t}()),
	// 		RestorePointLabel: to.Ptr("mylabel"),
	// 		RestorePointType: to.Ptr(armsql.RestorePointTypeDISCRETE),
	// 	},
	// }
}
