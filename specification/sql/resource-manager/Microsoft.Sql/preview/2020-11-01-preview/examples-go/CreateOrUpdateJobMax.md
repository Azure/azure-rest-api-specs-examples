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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/CreateOrUpdateJobMax.json
func ExampleJobsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsql.NewJobsClient("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"group1",
		"server1",
		"agent1",
		"job1",
		armsql.Job{
			Properties: &armsql.JobProperties{
				Description: to.Ptr("my favourite job"),
				Schedule: &armsql.JobSchedule{
					Type:      to.Ptr(armsql.JobScheduleTypeRecurring),
					Enabled:   to.Ptr(true),
					EndTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-24T23:59:59Z"); return t }()),
					Interval:  to.Ptr("PT5M"),
					StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2015-09-24T18:30:01Z"); return t }()),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
