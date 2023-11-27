package armmysql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/WaitStatisticsGet.json
func ExampleWaitStatisticsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWaitStatisticsClient().Get(ctx, "testResourceGroupName", "testServerName", "636927606000000000-636927615000000000-send-wait/io/socket/sql/client_connection-2--0", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WaitStatistic = armmysql.WaitStatistic{
	// 	Name: to.Ptr("636927606000000000-636927615000000000-send-wait/io/socket/sql/client_connection-2--0"),
	// 	Type: to.Ptr("Microsoft.DBforMySQL/servers/waitStatistics"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroupName/providers/Microsoft.DBforMySQL/servers/testServerName/waitStatistics/636927606000000000-636927615000000000-send-wait/io/socket/sql/client_connection-2--0"),
	// 	Properties: &armmysql.WaitStatisticProperties{
	// 		Count: to.Ptr[int64](3),
	// 		DatabaseName: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroupName/providers/Microsoft.DBforMySQL/servers/testServerName/databases/mysql"),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-06T17:45:00.000Z"); return t}()),
	// 		EventName: to.Ptr("wait/io/socket/sql/client_connection"),
	// 		EventTypeName: to.Ptr("send"),
	// 		QueryID: to.Ptr[int64](2),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-06T17:30:00.000Z"); return t}()),
	// 		TotalTimeInMs: to.Ptr[float64](12.345),
	// 		UserID: to.Ptr[int64](0),
	// 	},
	// }
}
