package armcostmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2021-10-01/examples/MCABillingProfileQueryGrouping.json
func ExampleQueryClient_Usage_billingProfileQueryGroupingModern() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcostmanagement.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewQueryClient().Usage(ctx, "providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579", armcostmanagement.QueryDefinition{
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
	// 	Name: to.Ptr("ad67fd91-c131-4bda-9ba9-7187ecb1cebd"),
	// 	Type: to.Ptr("microsoft.costmanagement/Query"),
	// 	ID: to.Ptr("providers/Microsoft.Billing/billingAccounts/12345:6789/billingProfiles/13579/providers/Microsoft.CostManagement/Query/ad67fd91-c131-4bda-9ba9-7187ecb1cebd"),
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
	// 				float64(19.545363672276512),
	// 				"JapanUnifia-Trial",
	// 				"USD"},
	// 				[]any{
	// 					float64(173.41979241290323),
	// 					"RVIIOT-TRIAL",
	// 					"USD"},
	// 					[]any{
	// 						float64(20.35941656262545),
	// 						"VSTSHOL-1595322048000",
	// 						"USD"}},
	// 					},
	// 				}
}
