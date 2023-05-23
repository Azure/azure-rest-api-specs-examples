package armcostmanagement_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/SubscriptionForecast.json
func ExampleForecastClient_Usage_subscriptionForecast() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewForecastClient().Usage(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", armcostmanagement.ForecastDefinition{
		Type: to.Ptr(armcostmanagement.ForecastTypeUsage),
		Dataset: &armcostmanagement.ForecastDataset{
			Aggregation: map[string]*armcostmanagement.ForecastAggregation{
				"totalCost": {
					Name:     to.Ptr(armcostmanagement.FunctionNameCost),
					Function: to.Ptr(armcostmanagement.FunctionTypeSum),
				},
			},
			Filter: &armcostmanagement.ForecastFilter{
				And: []*armcostmanagement.ForecastFilter{
					{
						Or: []*armcostmanagement.ForecastFilter{
							{
								Dimensions: &armcostmanagement.ForecastComparisonExpression{
									Name:     to.Ptr("ResourceLocation"),
									Operator: to.Ptr(armcostmanagement.ForecastOperatorTypeIn),
									Values: []*string{
										to.Ptr("East US"),
										to.Ptr("West Europe")},
								},
							},
							{
								Tags: &armcostmanagement.ForecastComparisonExpression{
									Name:     to.Ptr("Environment"),
									Operator: to.Ptr(armcostmanagement.ForecastOperatorTypeIn),
									Values: []*string{
										to.Ptr("UAT"),
										to.Ptr("Prod")},
								},
							}},
					},
					{
						Dimensions: &armcostmanagement.ForecastComparisonExpression{
							Name:     to.Ptr("ResourceGroup"),
							Operator: to.Ptr(armcostmanagement.ForecastOperatorTypeIn),
							Values: []*string{
								to.Ptr("API")},
						},
					}},
			},
			Granularity: to.Ptr(armcostmanagement.GranularityTypeDaily),
		},
		IncludeActualCost:       to.Ptr(false),
		IncludeFreshPartialCost: to.Ptr(false),
		TimePeriod: &armcostmanagement.ForecastTimePeriod{
			From: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-01T00:00:00+00:00"); return t }()),
			To:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-08-31T23:59:59+00:00"); return t }()),
		},
		Timeframe: to.Ptr(armcostmanagement.ForecastTimeframeCustom),
	}, &armcostmanagement.ForecastClientUsageOptions{Filter: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ForecastResult = armcostmanagement.ForecastResult{
	// 	Name: to.Ptr("55312978-ba1b-415c-9304-cfd9c43c0481"),
	// 	Type: to.Ptr("Microsoft.CostManagement/query"),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/query/00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armcostmanagement.ForecastProperties{
	// 		Columns: []*armcostmanagement.ForecastColumn{
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
