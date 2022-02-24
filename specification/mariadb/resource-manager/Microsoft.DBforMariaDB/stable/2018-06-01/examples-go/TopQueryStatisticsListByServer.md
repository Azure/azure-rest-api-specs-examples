Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmariadb%2Farmmariadb%2Fv0.2.1/sdk/resourcemanager/mariadb/armmariadb/README.md) on how to add the SDK to your project and authenticate.

```go
package armmariadb_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mariadb/armmariadb"
)

// x-ms-original-file: specification/mariadb/resource-manager/Microsoft.DBforMariaDB/stable/2018-06-01/examples/TopQueryStatisticsListByServer.json
func ExampleTopQueryStatisticsClient_ListByServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmariadb.NewTopQueryStatisticsClient("<subscription-id>", cred, nil)
	pager := client.ListByServer("<resource-group-name>",
		"<server-name>",
		armmariadb.TopQueryStatisticsInput{
			Properties: &armmariadb.TopQueryStatisticsInputProperties{
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
