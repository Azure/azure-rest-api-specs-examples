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
	}
	ctx := context.Background()
	client, err := armportal.NewDashboardsClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"testRG",
		"testDashboard",
		armportal.Dashboard{
			Location: to.Ptr("eastus"),
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
	}
	// TODO: use response item
	_ = res
}
