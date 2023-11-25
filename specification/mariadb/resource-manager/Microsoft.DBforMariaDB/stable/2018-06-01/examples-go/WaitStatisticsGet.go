package armmariadb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/WaitStatisticsGet.json
func ExampleWaitStatisticsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmariadb.NewClientFactory("<subscription-id>", cred, nil)
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
	// res.WaitStatistic = armmariadb.WaitStatistic{
	// 	Name: to.Ptr("636927606000000000-636927615000000000-send-wait/io/socket/sql/client_connection-2--0"),
	// 	Type: to.Ptr("Microsoft.DBforMariaDB/servers/waitStatistics"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroupName/providers/Microsoft.DBforMariaDB/servers/testServerName/waitStatistics/636927606000000000-636927615000000000-send-wait/io/socket/sql/client_connection-2--0"),
	// 	Properties: &armmariadb.WaitStatisticProperties{
	// 		Count: to.Ptr[int64](2),
	// 		DatabaseName: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroupName/providers/Microsoft.DBforMariaDB/servers/testServerName/databases/mariadb"),
	// 		EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-06T17:45:00.000Z"); return t}()),
	// 		EventName: to.Ptr("wait/io/socket/sql/client_connection"),
	// 		EventTypeName: to.Ptr("send"),
	// 		QueryID: to.Ptr[int64](2),
	// 		StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-06T17:30:00.000Z"); return t}()),
	// 		TotalTimeInMs: to.Ptr[float64](12.346),
	// 		UserID: to.Ptr[int64](0),
	// 	},
	// }
}
