Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmysql%2Farmmysql%2Fv0.3.1/sdk/resourcemanager/mysql/armmysql/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/TopQueryStatisticsListByServer.json
func ExampleTopQueryStatisticsClient_ListByServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmysql.NewTopQueryStatisticsClient("<subscription-id>", cred, nil)
	pager := client.ListByServer("<resource-group-name>",
		"<server-name>",
		armmysql.TopQueryStatisticsInput{
			Properties: &armmysql.TopQueryStatisticsInputProperties{
				AggregationFunction:  to.StringPtr("<aggregation-function>"),
				AggregationWindow:    to.StringPtr("<aggregation-window>"),
				NumberOfTopQueries:   to.Int32Ptr(5),
				ObservationEndTime:   to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-07T20:00:00.000Z"); return t }()),
				ObservationStartTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-01T20:00:00.000Z"); return t }()),
				ObservedMetric:       to.StringPtr("<observed-metric>"),
			},
		},
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}
```
