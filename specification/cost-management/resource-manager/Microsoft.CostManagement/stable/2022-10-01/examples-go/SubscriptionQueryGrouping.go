package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/SubscriptionQueryGrouping.json
func ExampleQueryClient_Usage_subscriptionQueryGroupingLegacy() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQueryClient().Usage(ctx, "subscriptions/00000000-0000-0000-0000-000000000000", armcostmanagement.QueryDefinition{
		Type: to.Ptr(armcostmanagement.ExportTypeUsage),
		Dataset: &armcostmanagement.QueryDataset{
			Aggregation: map[string]*armcostmanagement.QueryAggregation{
				"totalCost": {
					Name:     to.Ptr("PreTaxCost"),
					Function: to.Ptr(armcostmanagement.FunctionTypeSum),
				},
			},
			Granularity: to.Ptr(armcostmanagement.GranularityType("None")),
			Grouping: []*armcostmanagement.QueryGrouping{
				{
					Name: to.Ptr("ResourceGroup"),
					Type: to.Ptr(armcostmanagement.QueryColumnTypeDimension),
				}},
		},
		Timeframe: to.Ptr(armcostmanagement.TimeframeTypeTheLastMonth),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryResult = armcostmanagement.QueryResult{
	// 	Name: to.Ptr("55312978-ba1b-415c-9304-cfd9c43c0481"),
	// 	Type: to.Ptr("microsoft.costmanagement/Query"),
	// 	ID: to.Ptr("subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.CostManagement/Query/00000000-0000-0000-0000-000000000000"),
	// 	Properties: &armcostmanagement.QueryProperties{
	// 		Columns: []*armcostmanagement.QueryColumn{
	// 			{
	// 				Name: to.Ptr("PreTaxCost"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("ResourceGroup"),
	// 				Type: to.Ptr("String"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Currency"),
	// 				Type: to.Ptr("String"),
	// 		}},
	// 		Rows: [][]any{
	// 			[]any{
	// 				float64(0.009865586851323632),
	// 				"Ict_StratAndPlan_GoldSprova_Prod_0",
	// 				"USD"},
	// 				[]any{
	// 					float64(218.68795741935486),
	// 					"Ict_StratAndPlan_GoldSprova_Prod_1",
	// 					"USD"},
	// 					[]any{
	// 						float64(2.10333307059661),
	// 						"ScreenSharingTest-peer1",
	// 						"USD"},
	// 						[]any{
	// 							float64(0.14384913581657052),
	// 							"Ssbciotelement01",
	// 							"USD"}},
	// 						},
	// 					}
}
