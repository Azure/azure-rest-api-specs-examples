Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsql%2Farmsql%2Fv1.0.0/sdk/resourcemanager/sql/armsql/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ListJobExecutionsByAgentWithFilter.json
func ExampleJobExecutionsClient_NewListByAgentPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewJobExecutionsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByAgentPager("group1",
		"server1",
		"agent1",
		&armsql.JobExecutionsClientListByAgentOptions{CreateTimeMin: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-21T19:00:00Z"); return t }()),
			CreateTimeMax: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-21T19:05:00Z"); return t }()),
			EndTimeMin:    to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-21T19:20:00Z"); return t }()),
			EndTimeMax:    to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-21T19:25:00Z"); return t }()),
			IsActive:      to.Ptr(false),
			Skip:          nil,
			Top:           nil,
		})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
```
