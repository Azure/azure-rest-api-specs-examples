```go
package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPackQueriesSearch.json
func ExampleQueriesClient_NewSearchPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armoperationalinsights.NewQueriesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.NewSearchPager("<resource-group-name>",
		"<query-pack-name>",
		armoperationalinsights.LogAnalyticsQueryPackQuerySearchProperties{
			Related: &armoperationalinsights.LogAnalyticsQueryPackQuerySearchPropertiesRelated{
				Categories: []*string{
					to.Ptr("other"),
					to.Ptr("analytics")},
			},
			Tags: map[string][]*string{
				"my-label": {
					to.Ptr("label1")},
			},
		},
		&armoperationalinsights.QueriesClientSearchOptions{Top: to.Ptr[int64](3),
			IncludeBody: to.Ptr(true),
			SkipToken:   nil,
		})
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Foperationalinsights%2Farmoperationalinsights%2Fv0.5.0/sdk/resourcemanager/operationalinsights/armoperationalinsights/README.md) on how to add the SDK to your project and authenticate.
