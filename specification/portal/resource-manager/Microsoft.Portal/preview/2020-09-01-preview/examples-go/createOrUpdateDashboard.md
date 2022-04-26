Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fportal%2Farmportal%2Fv0.4.0/sdk/resourcemanager/portal/armportal/README.md) on how to add the SDK to your project and authenticate.

```go
package armportal_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/createOrUpdateDashboard.json
func ExampleDashboardsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armportal.NewDashboardsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<dashboard-name>",
		armportal.Dashboard{
			Location: to.Ptr("<location>"),
			Properties: &armportal.DashboardProperties{
				Lenses: []*armportal.DashboardLens{
					{
						Order: to.Ptr[int32](1),
						Parts: []*armportal.DashboardParts{
							{
								Position: &armportal.DashboardPartsPosition{
									ColSpan: to.Ptr[int32](3),
									RowSpan: to.Ptr[int32](4),
									X:       to.Ptr[int32](1),
									Y:       to.Ptr[int32](2),
								},
							},
							{
								Position: &armportal.DashboardPartsPosition{
									ColSpan: to.Ptr[int32](6),
									RowSpan: to.Ptr[int32](6),
									X:       to.Ptr[int32](5),
									Y:       to.Ptr[int32](5),
								},
							}},
					},
					{
						Order: to.Ptr[int32](2),
						Parts: []*armportal.DashboardParts{},
					}},
				Metadata: map[string]interface{}{
					"metadata": map[string]interface{}{
						"ColSpan": float64(2),
						"RowSpan": float64(1),
						"X":       float64(4),
						"Y":       float64(3),
					},
				},
			},
			Tags: map[string]*string{
				"aKey":       to.Ptr("aValue"),
				"anotherKey": to.Ptr("anotherValue"),
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
