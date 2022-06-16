package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/ViewCreateOrUpdateByResourceGroup.json
func ExampleViewsClient_CreateOrUpdateByScope() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcostmanagement.NewViewsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdateByScope(ctx,
		"subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG",
		"swaggerExample",
		armcostmanagement.View{
			ETag: to.Ptr("\"1d4ff9fe66f1d10\""),
			Properties: &armcostmanagement.ViewProperties{
				Accumulated: to.Ptr(armcostmanagement.AccumulatedTypeTrue),
				Chart:       to.Ptr(armcostmanagement.ChartTypeTable),
				DisplayName: to.Ptr("swagger Example"),
				Kpis: []*armcostmanagement.KpiProperties{
					{
						Type:    to.Ptr(armcostmanagement.KpiTypeForecast),
						Enabled: to.Ptr(true),
					},
					{
						Type:    to.Ptr(armcostmanagement.KpiTypeBudget),
						Enabled: to.Ptr(true),
						ID:      to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MYDEVTESTRG/providers/Microsoft.Consumption/budgets/swaggerDemo"),
					}},
				Metric: to.Ptr(armcostmanagement.MetricTypeActualCost),
				Pivots: []*armcostmanagement.PivotProperties{
					{
						Name: to.Ptr("ServiceName"),
						Type: to.Ptr(armcostmanagement.PivotTypeDimension),
					},
					{
						Name: to.Ptr("MeterCategory"),
						Type: to.Ptr(armcostmanagement.PivotTypeDimension),
					},
					{
						Name: to.Ptr("swaggerTagKey"),
						Type: to.Ptr(armcostmanagement.PivotTypeTagKey),
					}},
				Query: &armcostmanagement.ReportConfigDefinition{
					Type: to.Ptr(armcostmanagement.ReportTypeUsage),
					DataSet: &armcostmanagement.ReportConfigDataset{
						Aggregation: map[string]*armcostmanagement.ReportConfigAggregation{
							"totalCost": {
								Name:     to.Ptr("PreTaxCost"),
								Function: to.Ptr(armcostmanagement.FunctionTypeSum),
							},
						},
						Granularity: to.Ptr(armcostmanagement.ReportGranularityTypeDaily),
						Grouping:    []*armcostmanagement.ReportConfigGrouping{},
						Sorting: []*armcostmanagement.ReportConfigSorting{
							{
								Name:      to.Ptr("UsageDate"),
								Direction: to.Ptr(armcostmanagement.ReportConfigSortingTypeAscending),
							}},
					},
					Timeframe: to.Ptr(armcostmanagement.ReportTimeframeTypeMonthToDate),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
