Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcostmanagement%2Farmcostmanagement%2Fv0.1.0/sdk/resourcemanager/costmanagement/armcostmanagement/README.md) on how to add the SDK to your project and authenticate.

```go
package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ViewCreateOrUpdateByResourceGroup.json
func ExampleViewsClient_CreateOrUpdateByScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcostmanagement.NewViewsClient(cred, nil)
	res, err := client.CreateOrUpdateByScope(ctx,
		"<scope>",
		"<view-name>",
		armcostmanagement.View{
			ProxyResource: armcostmanagement.ProxyResource{
				ETag: to.StringPtr("<etag>"),
			},
			Properties: &armcostmanagement.ViewProperties{
				Accumulated: armcostmanagement.AccumulatedTypeTrue.ToPtr(),
				Chart:       armcostmanagement.ChartTypeTable.ToPtr(),
				DisplayName: to.StringPtr("<display-name>"),
				Kpis: []*armcostmanagement.KpiProperties{
					{
						Type:    armcostmanagement.KpiTypeForecast.ToPtr(),
						Enabled: to.BoolPtr(true),
						ID:      to.StringPtr("<id>"),
					},
					{
						Type:    armcostmanagement.KpiTypeBudget.ToPtr(),
						Enabled: to.BoolPtr(true),
						ID:      to.StringPtr("<id>"),
					}},
				Metric: armcostmanagement.MetricTypeActualCost.ToPtr(),
				Pivots: []*armcostmanagement.PivotProperties{
					{
						Name: to.StringPtr("<name>"),
						Type: armcostmanagement.PivotTypeDimension.ToPtr(),
					},
					{
						Name: to.StringPtr("<name>"),
						Type: armcostmanagement.PivotTypeDimension.ToPtr(),
					},
					{
						Name: to.StringPtr("<name>"),
						Type: armcostmanagement.PivotTypeTagKey.ToPtr(),
					}},
				Query: &armcostmanagement.ReportConfigDefinition{
					Type: armcostmanagement.ReportTypeUsage.ToPtr(),
					DataSet: &armcostmanagement.ReportConfigDataset{
						Aggregation: map[string]*armcostmanagement.ReportConfigAggregation{
							"totalCost": {
								Name:     to.StringPtr("<name>"),
								Function: armcostmanagement.FunctionTypeSum.ToPtr(),
							},
						},
						Granularity: armcostmanagement.ReportGranularityTypeDaily.ToPtr(),
						Grouping:    []*armcostmanagement.ReportConfigGrouping{},
						Sorting: []*armcostmanagement.ReportConfigSorting{
							{
								Name:      to.StringPtr("<name>"),
								Direction: armcostmanagement.ReportConfigSortingDirectionAscending.ToPtr(),
							}},
					},
					Timeframe: armcostmanagement.ReportTimeframeTypeMonthToDate.ToPtr(),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("View.ID: %s\n", *res.ID)
}
```
