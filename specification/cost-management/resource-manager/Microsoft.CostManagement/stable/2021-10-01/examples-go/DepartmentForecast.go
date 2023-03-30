package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/DepartmentForecast.json
func ExampleForecastClient_Usage_departmentForecast() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewForecastClient().Usage(ctx, "providers/Microsoft.Billing/billingAccounts/12345:6789/departments/123", armcostmanagement.ForecastDefinition{
		Type: to.Ptr(armcostmanagement.ForecastTypeUsage),
		Dataset: &armcostmanagement.ForecastDataset{
			Filter: &armcostmanagement.QueryFilter{
				And: []*armcostmanagement.QueryFilter{
					{
						Or: []*armcostmanagement.QueryFilter{
							{
								Dimensions: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("ResourceLocation"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("East US"),
										to.Ptr("West Europe")},
								},
							},
							{
								Tags: &armcostmanagement.QueryComparisonExpression{
									Name:     to.Ptr("Environment"),
									Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
									Values: []*string{
										to.Ptr("UAT"),
										to.Ptr("Prod")},
								},
							}},
					},
					{
						Dimensions: &armcostmanagement.QueryComparisonExpression{
							Name:     to.Ptr("ResourceGroup"),
							Operator: to.Ptr(armcostmanagement.QueryOperatorTypeIn),
							Values: []*string{
								to.Ptr("API")},
						},
					}},
			},
			Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
		},
		IncludeActualCost:       to.Ptr(false),
		IncludeFreshPartialCost: to.Ptr(false),
		Timeframe:               to.Ptr(armcostmanagement.ForecastTimeframeTypeMonthToDate),
	}, &armcostmanagement.ForecastClientUsageOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryResult = armcostmanagement.QueryResult{
	// 	Name: to.Ptr("ad67fd91-c131-4bda-9ba9-7187ecb1cebd"),
	// 	Type: to.Ptr("microsoft.costmanagement/Query"),
	// 	ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/12345:6789/departments/123/providers/Microsoft.CostManagement/query/ad67fd91-c131-4bda-9ba9-7187ecb1cebd"),
	// 	Properties: &armcostmanagement.QueryProperties{
	// 		Columns: []*armcostmanagement.QueryColumn{
	// 			{
	// 				Name: to.Ptr("PreTaxCost"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("UsageDate"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("CostStatus"),
	// 				Type: to.Ptr("String"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Currency"),
	// 				Type: to.Ptr("String"),
	// 		}},
	// 		Rows: [][]any{
	// 			[]any{
	// 				float64(2.10333307059661),
	// 				float64(20180331),
	// 				"Forecast",
	// 				"USD"},
	// 				[]any{
	// 					float64(218.68795741935486),
	// 					float64(20180331),
	// 					"Forecast",
	// 					"USD"},
	// 					[]any{
	// 						float64(0.14384913581657052),
	// 						float64(20180401),
	// 						"Forecast",
	// 						"USD"},
	// 						[]any{
	// 							float64(0.009865586851323632),
	// 							float64(20180429),
	// 							"Forecast",
	// 							"USD"}},
	// 						},
	// 					}
}
