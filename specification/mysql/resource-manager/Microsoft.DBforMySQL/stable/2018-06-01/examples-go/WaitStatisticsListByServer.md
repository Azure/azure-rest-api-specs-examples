Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmysql%2Farmmysql%2Fv0.5.0/sdk/resourcemanager/mysql/armmysql/README.md) on how to add the SDK to your project and authenticate.

```go
package armmysql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/WaitStatisticsListByServer.json
func ExampleWaitStatisticsClient_NewListByServerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmysql.NewWaitStatisticsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewListByServerPager("<resource-group-name>",
		"<server-name>",
		armmysql.WaitStatisticsInput{
			Properties: &armmysql.WaitStatisticsInputProperties{
				AggregationWindow:    to.Ptr("<aggregation-window>"),
				ObservationEndTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-07T20:00:00.000Z"); return t }()),
				ObservationStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-01T20:00:00.000Z"); return t }()),
			},
		},
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
