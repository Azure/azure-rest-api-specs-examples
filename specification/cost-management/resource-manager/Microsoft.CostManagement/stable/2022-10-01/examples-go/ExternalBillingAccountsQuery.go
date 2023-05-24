package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExternalBillingAccountsQuery.json
func ExampleQueryClient_UsageByExternalCloudProviderType_externalBillingAccountQueryList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQueryClient().UsageByExternalCloudProviderType(ctx, armcostmanagement.ExternalCloudProviderTypeExternalBillingAccounts, "100", armcostmanagement.QueryDefinition{
		Type: to.Ptr(armcostmanagement.ExportTypeUsage),
		Dataset: &armcostmanagement.QueryDataset{
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
		Timeframe: to.Ptr(armcostmanagement.TimeframeTypeMonthToDate),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.QueryResult = armcostmanagement.QueryResult{
	// 	Name: to.Ptr("6dc7b06a-d90a-4df5-b655-ce6cf1c0814d"),
	// 	Type: to.Ptr("Microsoft.CostManagement/query"),
	// 	ID: to.Ptr("providers/Microsoft.CostManagement/externalBillingAccounts/100/query/6dc7b06a-d90a-4df5-b655-ce6cf1c0814d"),
	// 	Properties: &armcostmanagement.QueryProperties{
	// 		Columns: []*armcostmanagement.QueryColumn{
	// 			{
	// 				Name: to.Ptr("PreTaxCost"),
	// 				Type: to.Ptr("Number"),
	// 			},
	// 			{
	// 				Name: to.Ptr("ServiceName"),
	// 				Type: to.Ptr("String"),
	// 			},
	// 			{
	// 				Name: to.Ptr("Currency"),
	// 				Type: to.Ptr("String"),
	// 		}},
	// 		Rows: [][]any{
	// 			[]any{
	// 				float64(0),
	// 				"abc db",
	// 				"USD"},
	// 				[]any{
	// 					float64(30.2572751438),
	// 					"abc compute cloud",
	// 					"USD"},
	// 					[]any{
	// 						float64(0.07675760200000002),
	// 						"abc file system",
	// 						"USD"},
	// 						[]any{
	// 							float64(50.43096419040001),
	// 							"abc elasticache",
	// 							"USD"}},
	// 						},
	// 					}
}
