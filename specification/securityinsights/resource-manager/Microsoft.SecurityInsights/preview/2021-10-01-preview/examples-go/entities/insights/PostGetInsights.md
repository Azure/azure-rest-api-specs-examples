Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsights%2Farmsecurityinsights%2Fv0.1.1/sdk/resourcemanager/securityinsights/armsecurityinsights/README.md) on how to add the SDK to your project and authenticate.

```go
package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
)

// x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-10-01-preview/examples/entities/insights/PostGetInsights.json
func ExampleEntitiesClient_GetInsights() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armsecurityinsights.NewEntitiesClient("<subscription-id>", cred, nil)
	res, err := client.GetInsights(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<entity-id>",
		armsecurityinsights.EntityGetInsightsParameters{
			AddDefaultExtendedTimeRange: to.BoolPtr(false),
			EndTime:                     to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-01T00:00:00.000Z"); return t }()),
			InsightQueryIDs: []*string{
				to.StringPtr("cae8d0aa-aa45-4d53-8d88-17dd64ffd4e4")},
			StartTime: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T00:00:00.000Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.EntitiesClientGetInsightsResult)
}
```
