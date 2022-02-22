Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv0.3.1/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

```go
package armsql_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListJobExecutionsByAgentWithFilter.json
func ExampleJobExecutionsClient_ListByAgent() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsql.NewJobExecutionsClient("<subscription-id>", cred, nil)
	pager := client.ListByAgent("<resource-group-name>",
		"<server-name>",
		"<job-agent-name>",
		&armsql.JobExecutionsClientListByAgentOptions{CreateTimeMin: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-21T19:00:00Z"); return t }()),
			CreateTimeMax: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-21T19:05:00Z"); return t }()),
			EndTimeMin:    to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-21T19:20:00Z"); return t }()),
			EndTimeMax:    to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-21T19:25:00Z"); return t }()),
			IsActive:      to.BoolPtr(false),
			Skip:          nil,
			Top:           nil,
		})
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
