package armmysql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/WaitStatisticsListByServer.json
func ExampleWaitStatisticsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmysql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWaitStatisticsClient().NewListByServerPager("testResourceGroupName", "testServerName", armmysql.WaitStatisticsInput{
		Properties: &armmysql.WaitStatisticsInputProperties{
			AggregationWindow:    to.Ptr("PT15M"),
			ObservationEndTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-07T20:00:00.000Z"); return t }()),
			ObservationStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-01T20:00:00.000Z"); return t }()),
		},
	}, nil)
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
		// page.WaitStatisticsResultList = armmysql.WaitStatisticsResultList{
		// 	Value: []*armmysql.WaitStatistic{
		// 		{
		// 			Name: to.Ptr("636927606000000000-636927615000000000-send-wait/io/socket/sql/client_connection-2--0"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/waitStatistics"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroupName/providers/Microsoft.DBforMySQL/servers/testServerName/waitStatistics/636927606000000000-636927615000000000-send-wait/io/socket/sql/client_connection-2--0"),
		// 			Properties: &armmysql.WaitStatisticProperties{
		// 				Count: to.Ptr[int64](2),
		// 				DatabaseName: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testResourceGroupName/providers/Microsoft.DBforMySQL/servers/testServerName/databases/mysql"),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-06T17:45:00.000Z"); return t}()),
		// 				EventName: to.Ptr("wait/io/socket/sql/client_connection"),
		// 				EventTypeName: to.Ptr("send"),
		// 				QueryID: to.Ptr[int64](2),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-06T17:30:00.000Z"); return t}()),
		// 				TotalTimeInMs: to.Ptr[float64](12.345),
		// 				UserID: to.Ptr[int64](0),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("636927606000000000-636927615000000000-lock-wait/synch/mutex/mysys/THR_LOCK::mutex-2--0"),
		// 			Type: to.Ptr("Microsoft.DBforMySQL/servers/waitStatistics"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/hyshim-test/providers/Microsoft.DBforMySQL/servers/hyshim-wait-stats-fix/waitStatistics/636927606000000000-636927615000000000-lock-wait/synch/mutex/mysys/THR_LOCK::mutex-2--0"),
		// 			Properties: &armmysql.WaitStatisticProperties{
		// 				Count: to.Ptr[int64](4),
		// 				DatabaseName: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/hyshim-test/providers/Microsoft.DBforMySQL/servers/hyshim-wait-stats-fix/databases/"),
		// 				EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-06T17:45:00.000Z"); return t}()),
		// 				EventName: to.Ptr("wait/synch/mutex/mysys/THR_LOCK::mutex"),
		// 				EventTypeName: to.Ptr("lock"),
		// 				QueryID: to.Ptr[int64](2),
		// 				StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-06T17:30:00.000Z"); return t}()),
		// 				TotalTimeInMs: to.Ptr[float64](56.789),
		// 				UserID: to.Ptr[int64](0),
		// 			},
		// 	}},
		// }
	}
}
