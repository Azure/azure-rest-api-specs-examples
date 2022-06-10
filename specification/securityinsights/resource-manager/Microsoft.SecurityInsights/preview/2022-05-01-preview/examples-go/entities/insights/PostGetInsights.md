```go
package armsecurityinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-05-01-preview/examples/entities/insights/PostGetInsights.json
func ExampleEntitiesClient_GetInsights() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armsecurityinsights.NewEntitiesClient("d0cfe6b2-9ac0-4464-9919-dccaee2e48c0", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetInsights(ctx,
		"myRg",
		"myWorkspace",
		"e1d3d618-e11f-478b-98e3-bb381539a8e1",
		armsecurityinsights.EntityGetInsightsParameters{
			AddDefaultExtendedTimeRange: to.Ptr(false),
			EndTime:                     to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-10-01T00:00:00.000Z"); return t }()),
			InsightQueryIDs: []*string{
				to.Ptr("cae8d0aa-aa45-4d53-8d88-17dd64ffd4e4")},
			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-01T00:00:00.000Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fsecurityinsights%2Farmsecurityinsights%2Fv2.0.0-beta.1/sdk/resourcemanager/securityinsights/armsecurityinsights/README.md) on how to add the SDK to your project and authenticate.
