Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fpolicyinsights%2Farmpolicyinsights%2Fv0.1.0/sdk/resourcemanager/policyinsights/armpolicyinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armpolicyinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/policyinsights/armpolicyinsights"
)

// x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_FilterAndAggregateOnly.json
func ExamplePolicyEventsClient_ListQueryResultsForSubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armpolicyinsights.NewPolicyEventsClient(cred, nil)
	pager := client.ListQueryResultsForSubscription(armpolicyinsights.Enum1("default"),
		"<subscription-id>",
		&armpolicyinsights.QueryOptions{Top: nil,
			Filter:    to.StringPtr("<filter>"),
			OrderBy:   nil,
			Select:    nil,
			From:      to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-05T18:00:00Z"); return t }()),
			To:        nil,
			Apply:     to.StringPtr("<apply>"),
			SkipToken: nil,
			Expand:    nil,
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
