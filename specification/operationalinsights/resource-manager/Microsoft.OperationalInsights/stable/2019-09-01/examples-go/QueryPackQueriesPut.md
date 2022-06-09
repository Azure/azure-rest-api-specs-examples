```go
package armoperationalinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPackQueriesPut.json
func ExampleQueriesClient_Put() {
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
	res, err := client.Put(ctx,
		"<resource-group-name>",
		"<query-pack-name>",
		"<id>",
		armoperationalinsights.LogAnalyticsQueryPackQuery{
			Properties: &armoperationalinsights.LogAnalyticsQueryPackQueryProperties{
				Description: to.Ptr("<description>"),
				Body:        to.Ptr("<body>"),
				DisplayName: to.Ptr("<display-name>"),
				Related: &armoperationalinsights.LogAnalyticsQueryPackQueryPropertiesRelated{
					Categories: []*string{
						to.Ptr("analytics")},
				},
				Tags: map[string][]*string{
					"my-label": {
						to.Ptr("label1")},
					"my-other-label": {
						to.Ptr("label2")},
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Foperationalinsights%2Farmoperationalinsights%2Fv0.5.0/sdk/resourcemanager/operationalinsights/armoperationalinsights/README.md) on how to add the SDK to your project and authenticate.
