package armportal_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/getDashboard.json
func ExampleDashboardsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armportal.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDashboardsClient().Get(ctx, "testRG", "testDashboard", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Dashboard = armportal.Dashboard{
	// 	Name: to.Ptr("testDashboard"),
	// 	Type: to.Ptr("Microsoft.Portal/dashboards"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/testRG/providers/Microsoft.Portal/dashboards/testDashboard"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armportal.DashboardProperties{
	// 		Lenses: []*armportal.DashboardLens{
	// 			{
	// 				Order: to.Ptr[int32](1),
	// 				Parts: []*armportal.DashboardParts{
	// 					{
	// 						Position: &armportal.DashboardPartsPosition{
	// 							ColSpan: to.Ptr[int32](3),
	// 							RowSpan: to.Ptr[int32](4),
	// 							X: to.Ptr[int32](1),
	// 							Y: to.Ptr[int32](2),
	// 						},
	// 					},
	// 					{
	// 						Position: &armportal.DashboardPartsPosition{
	// 							ColSpan: to.Ptr[int32](6),
	// 							RowSpan: to.Ptr[int32](6),
	// 							X: to.Ptr[int32](5),
	// 							Y: to.Ptr[int32](5),
	// 						},
	// 				}},
	// 			},
	// 			{
	// 				Order: to.Ptr[int32](2),
	// 				Parts: []*armportal.DashboardParts{
	// 				},
	// 		}},
	// 		Metadata: map[string]any{
	// 			"metadata": map[string]any{
	// 				"ColSpan": float64(2),
	// 				"RowSpan": float64(1),
	// 				"X": float64(4),
	// 				"Y": float64(3),
	// 			},
	// 		},
	// 	},
	// 	Tags: map[string]*string{
	// 		"aKey": to.Ptr("aValue"),
	// 		"anotherKey": to.Ptr("anotherValue"),
	// 	},
	// }
}
