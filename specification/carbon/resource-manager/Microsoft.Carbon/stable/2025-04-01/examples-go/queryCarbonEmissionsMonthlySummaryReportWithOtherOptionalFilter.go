package armcarbonoptimization_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/carbonoptimization/armcarbonoptimization"
)

// Generated from example definition: 2025-04-01/queryCarbonEmissionsMonthlySummaryReportWithOtherOptionalFilter.json
func ExampleCarbonServiceClient_QueryCarbonEmissionReports_queryCarbonEmissionMonthlySummaryReportWithOptionalFilterLocationListResourceTypeListResourceGroupUrlList() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcarbonoptimization.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCarbonServiceClient().QueryCarbonEmissionReports(ctx, &armcarbonoptimization.MonthlySummaryReportQueryFilter{
		ReportType: to.Ptr(armcarbonoptimization.ReportTypeEnumMonthlySummaryReport),
		SubscriptionList: []*string{
			to.Ptr("00000000-0000-0000-0000-000000000000"),
		},
		CarbonScopeList: []*armcarbonoptimization.EmissionScopeEnum{
			to.Ptr(armcarbonoptimization.EmissionScopeEnumScope1),
			to.Ptr(armcarbonoptimization.EmissionScopeEnumScope3),
		},
		DateRange: &armcarbonoptimization.DateRange{
			Start: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-03-01"); return t }()),
			End:   to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-05-01"); return t }()),
		},
		LocationList: []*string{
			to.Ptr("east us"),
			to.Ptr("west us"),
		},
		ResourceTypeList: []*string{
			to.Ptr("microsoft.storage/storageaccounts"),
			to.Ptr("microsoft.databricks/workspaces"),
		},
		ResourceGroupURLList: []*string{
			to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg-name"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcarbonoptimization.CarbonServiceClientQueryCarbonEmissionReportsResponse{
	// 	CarbonEmissionDataListResult: &armcarbonoptimization.CarbonEmissionDataListResult{
	// 		SubscriptionAccessDecisionList: []*armcarbonoptimization.SubscriptionAccessDecision{
	// 			{
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				Decision: to.Ptr(armcarbonoptimization.AccessDecisionEnumAllowed),
	// 			},
	// 		},
	// 		Value: []armcarbonoptimization.CarbonEmissionDataClassification{
	// 			&armcarbonoptimization.CarbonEmissionMonthlySummaryData{
	// 				DataType: to.Ptr(armcarbonoptimization.ResponseDataTypeEnumMonthlySummaryData),
	// 				LatestMonthEmissions: to.Ptr[float64](0.1),
	// 				PreviousMonthEmissions: to.Ptr[float64](0.05),
	// 				MonthOverMonthEmissionsChangeRatio: to.Ptr[float64](1),
	// 				MonthlyEmissionsChangeValue: to.Ptr[float64](0.05),
	// 				Date: to.Ptr("2024-05-01"),
	// 				CarbonIntensity: to.Ptr[float64](22),
	// 			},
	// 			&armcarbonoptimization.CarbonEmissionMonthlySummaryData{
	// 				DataType: to.Ptr(armcarbonoptimization.ResponseDataTypeEnumMonthlySummaryData),
	// 				LatestMonthEmissions: to.Ptr[float64](0.1),
	// 				PreviousMonthEmissions: to.Ptr[float64](0.05),
	// 				MonthOverMonthEmissionsChangeRatio: to.Ptr[float64](1),
	// 				MonthlyEmissionsChangeValue: to.Ptr[float64](0.05),
	// 				Date: to.Ptr("2024-04-01"),
	// 				CarbonIntensity: to.Ptr[float64](22),
	// 			},
	// 			&armcarbonoptimization.CarbonEmissionMonthlySummaryData{
	// 				DataType: to.Ptr(armcarbonoptimization.ResponseDataTypeEnumMonthlySummaryData),
	// 				LatestMonthEmissions: to.Ptr[float64](0.1),
	// 				PreviousMonthEmissions: to.Ptr[float64](0.05),
	// 				MonthOverMonthEmissionsChangeRatio: to.Ptr[float64](1),
	// 				MonthlyEmissionsChangeValue: to.Ptr[float64](0.05),
	// 				Date: to.Ptr("2024-03-01"),
	// 				CarbonIntensity: to.Ptr[float64](22),
	// 			},
	// 		},
	// 	},
	// }
}
