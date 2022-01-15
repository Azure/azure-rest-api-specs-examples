Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcostmanagement%2Farmcostmanagement%2Fv0.2.0/sdk/resourcemanager/costmanagement/armcostmanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/PrivateViewCreateOrUpdate.json
func ExampleViewsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcostmanagement.NewViewsClient(cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<view-name>",
		armcostmanagement.View{
			ETag: to.StringPtr("<etag>"),
			Properties: &armcostmanagement.ViewProperties{
				Accumulated: armcostmanagement.AccumulatedType("true").ToPtr(),
				Chart:       armcostmanagement.ChartType("Table").ToPtr(),
				DisplayName: to.StringPtr("<display-name>"),
				Kpis: []*armcostmanagement.KpiProperties{
					{
						Type:    armcostmanagement.KpiType("Forecast").ToPtr(),
						Enabled: to.BoolPtr(true),
					},
					{
						Type:    armcostmanagement.KpiType("Budget").ToPtr(),
						Enabled: to.BoolPtr(true),
						ID:      to.StringPtr("<id>"),
					}},
				Metric: armcostmanagement.MetricType("ActualCost").ToPtr(),
				Pivots: []*armcostmanagement.PivotProperties{
					{
						Name: to.StringPtr("<name>"),
						Type: armcostmanagement.PivotType("Dimension").ToPtr(),
					},
					{
						Name: to.StringPtr("<name>"),
						Type: armcostmanagement.PivotType("Dimension").ToPtr(),
					},
					{
						Name: to.StringPtr("<name>"),
						Type: armcostmanagement.PivotType("TagKey").ToPtr(),
					}},
				Query: &armcostmanagement.ReportConfigDefinition{
					Type: armcostmanagement.ReportType("Usage").ToPtr(),
					DataSet: &armcostmanagement.ReportConfigDataset{
						Aggregation: map[string]*armcostmanagement.ReportConfigAggregation{
							"totalCost": {
								Name:     to.StringPtr("<name>"),
								Function: armcostmanagement.FunctionType("Sum").ToPtr(),
							},
						},
						Granularity: armcostmanagement.ReportGranularityType("Daily").ToPtr(),
						Grouping:    []*armcostmanagement.ReportConfigGrouping{},
						Sorting: []*armcostmanagement.ReportConfigSorting{
							{
								Name:      to.StringPtr("<name>"),
								Direction: armcostmanagement.ReportConfigSortingDirection("Ascending").ToPtr(),
							}},
					},
					Timeframe: armcostmanagement.ReportTimeframeType("MonthToDate").ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ViewsClientCreateOrUpdateResult)
}
```
