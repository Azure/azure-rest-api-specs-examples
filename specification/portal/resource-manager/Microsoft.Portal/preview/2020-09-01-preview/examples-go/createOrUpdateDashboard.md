Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fportal%2Farmportal%2Fv0.2.0/sdk/resourcemanager/portal/armportal/README.md) on how to add the SDK to your project and authenticate.

```go
package armportal_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"
)

// x-ms-original-file: specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/createOrUpdateDashboard.json
func ExampleDashboardsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armportal.NewDashboardsClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<dashboard-name>",
		armportal.Dashboard{
			Location: to.StringPtr("<location>"),
			Properties: &armportal.DashboardProperties{
				Lenses: []*armportal.DashboardLens{
					{
						Order: to.Int32Ptr(1),
						Parts: []*armportal.DashboardParts{
							{
								Position: &armportal.DashboardPartsPosition{
									ColSpan: to.Int32Ptr(3),
									RowSpan: to.Int32Ptr(4),
									X:       to.Int32Ptr(1),
									Y:       to.Int32Ptr(2),
								},
							},
							{
								Position: &armportal.DashboardPartsPosition{
									ColSpan: to.Int32Ptr(6),
									RowSpan: to.Int32Ptr(6),
									X:       to.Int32Ptr(5),
									Y:       to.Int32Ptr(5),
								},
							}},
					},
					{
						Order: to.Int32Ptr(2),
						Parts: []*armportal.DashboardParts{},
					}},
				Metadata: map[string]map[string]interface{}{
					"metadata": {
						"ColSpan": float64(2),
						"RowSpan": float64(1),
						"X":       float64(4),
						"Y":       float64(3),
					},
				},
			},
			Tags: map[string]*string{
				"aKey":       to.StringPtr("aValue"),
				"anotherKey": to.StringPtr("anotherValue"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.DashboardsClientCreateOrUpdateResult)
}
```
