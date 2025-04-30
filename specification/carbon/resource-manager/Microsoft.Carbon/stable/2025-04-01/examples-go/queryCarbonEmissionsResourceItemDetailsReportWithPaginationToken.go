package armcarbonoptimization_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/carbonoptimization/armcarbonoptimization"
)

// Generated from example definition: 2025-04-01/queryCarbonEmissionsResourceItemDetailsReportWithPaginationToken.json
func ExampleCarbonServiceClient_QueryCarbonEmissionReports_queryCarbonEmissionResourceItemDetailsReportWithPaginationToken() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcarbonoptimization.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCarbonServiceClient().QueryCarbonEmissionReports(ctx, &armcarbonoptimization.ItemDetailsQueryFilter{
		ReportType: to.Ptr(armcarbonoptimization.ReportTypeEnumItemDetailsReport),
		SubscriptionList: []*string{
			to.Ptr("00000000-0000-0000-0000-000000000000"),
			to.Ptr("00000000-0000-0000-0000-000000000001,"),
			to.Ptr("00000000-0000-0000-0000-000000000002"),
			to.Ptr("00000000-0000-0000-0000-000000000003"),
			to.Ptr("00000000-0000-0000-0000-000000000004"),
			to.Ptr("00000000-0000-0000-0000-000000000005"),
			to.Ptr("00000000-0000-0000-0000-000000000006"),
			to.Ptr("00000000-0000-0000-0000-000000000007"),
			to.Ptr("00000000-0000-0000-0000-000000000008"),
		},
		CarbonScopeList: []*armcarbonoptimization.EmissionScopeEnum{
			to.Ptr(armcarbonoptimization.EmissionScopeEnumScope1),
			to.Ptr(armcarbonoptimization.EmissionScopeEnumScope3),
		},
		DateRange: &armcarbonoptimization.DateRange{
			Start: to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-05-01"); return t }()),
			End:   to.Ptr(func() time.Time { t, _ := time.Parse("2006-01-02", "2024-05-01"); return t }()),
		},
		CategoryType:  to.Ptr(armcarbonoptimization.CategoryTypeEnumResource),
		OrderBy:       to.Ptr(armcarbonoptimization.OrderByColumnEnumLatestMonthEmissions),
		SortDirection: to.Ptr(armcarbonoptimization.SortDirectionEnumDesc),
		PageSize:      to.Ptr[int32](100),
		SkipToken:     to.Ptr("dGVzZGZhZGZzZnNkZg=="),
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
	// 			{
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 				Decision: to.Ptr(armcarbonoptimization.AccessDecisionEnumAllowed),
	// 			},
	// 			{
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000002"),
	// 				Decision: to.Ptr(armcarbonoptimization.AccessDecisionEnumAllowed),
	// 			},
	// 			{
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000003"),
	// 				Decision: to.Ptr(armcarbonoptimization.AccessDecisionEnumAllowed),
	// 			},
	// 			{
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000004"),
	// 				Decision: to.Ptr(armcarbonoptimization.AccessDecisionEnumAllowed),
	// 			},
	// 			{
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000005"),
	// 				Decision: to.Ptr(armcarbonoptimization.AccessDecisionEnumAllowed),
	// 			},
	// 			{
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000006"),
	// 				Decision: to.Ptr(armcarbonoptimization.AccessDecisionEnumDenied),
	// 			},
	// 			{
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000007"),
	// 				Decision: to.Ptr(armcarbonoptimization.AccessDecisionEnumDenied),
	// 			},
	// 			{
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000008"),
	// 				Decision: to.Ptr(armcarbonoptimization.AccessDecisionEnumDenied),
	// 				DenialReason: to.Ptr("Carbon Optimization Reader permisison required"),
	// 			},
	// 		},
	// 		Value: []armcarbonoptimization.CarbonEmissionDataClassification{
	// 			&armcarbonoptimization.ResourceCarbonEmissionItemDetailData{
	// 				DataType: to.Ptr(armcarbonoptimization.ResponseDataTypeEnumResourceItemDetailsData),
	// 				LatestMonthEmissions: to.Ptr[float64](0.1),
	// 				PreviousMonthEmissions: to.Ptr[float64](0.05),
	// 				MonthOverMonthEmissionsChangeRatio: to.Ptr[float64](1),
	// 				MonthlyEmissionsChangeValue: to.Ptr[float64](0.05),
	// 				ItemName: to.Ptr("rgName1"),
	// 				ResourceGroup: to.Ptr("rgGroup"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgGroup/providers/microsoft.storage/storageaccounts/rgName1"),
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				CategoryType: to.Ptr(armcarbonoptimization.CategoryTypeEnumResource),
	// 				ResourceType: to.Ptr("microsoft.storage/storageaccounts"),
	// 				Location: to.Ptr("east us"),
	// 			},
	// 			&armcarbonoptimization.ResourceCarbonEmissionItemDetailData{
	// 				DataType: to.Ptr(armcarbonoptimization.ResponseDataTypeEnumResourceItemDetailsData),
	// 				LatestMonthEmissions: to.Ptr[float64](0.1),
	// 				PreviousMonthEmissions: to.Ptr[float64](0.05),
	// 				MonthOverMonthEmissionsChangeRatio: to.Ptr[float64](1),
	// 				MonthlyEmissionsChangeValue: to.Ptr[float64](0.05),
	// 				ItemName: to.Ptr("rgName2"),
	// 				ResourceGroup: to.Ptr("rgGroup"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgGroup/providers/microsoft.storage/storageaccounts/rgName2"),
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				CategoryType: to.Ptr(armcarbonoptimization.CategoryTypeEnumResource),
	// 				ResourceType: to.Ptr("microsoft.storage/storageaccounts"),
	// 				Location: to.Ptr("east us"),
	// 			},
	// 			&armcarbonoptimization.ResourceCarbonEmissionItemDetailData{
	// 				DataType: to.Ptr(armcarbonoptimization.ResponseDataTypeEnumResourceItemDetailsData),
	// 				LatestMonthEmissions: to.Ptr[float64](0.1),
	// 				PreviousMonthEmissions: to.Ptr[float64](0.05),
	// 				MonthOverMonthEmissionsChangeRatio: to.Ptr[float64](1),
	// 				MonthlyEmissionsChangeValue: to.Ptr[float64](0.05),
	// 				ItemName: to.Ptr("rgName3"),
	// 				ResourceGroup: to.Ptr("rgGroup"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgGroup/providers/microsoft.storage/storageaccounts/rgName3"),
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				CategoryType: to.Ptr(armcarbonoptimization.CategoryTypeEnumResource),
	// 				ResourceType: to.Ptr("microsoft.storage/storageaccounts"),
	// 				Location: to.Ptr("east us"),
	// 			},
	// 			&armcarbonoptimization.ResourceCarbonEmissionItemDetailData{
	// 				DataType: to.Ptr(armcarbonoptimization.ResponseDataTypeEnumResourceItemDetailsData),
	// 				LatestMonthEmissions: to.Ptr[float64](0.1),
	// 				PreviousMonthEmissions: to.Ptr[float64](0.05),
	// 				MonthOverMonthEmissionsChangeRatio: to.Ptr[float64](1),
	// 				MonthlyEmissionsChangeValue: to.Ptr[float64](0.05),
	// 				ItemName: to.Ptr("rgName4"),
	// 				ResourceGroup: to.Ptr("rgGroup"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000002/resourceGroups/rgGroup/providers/microsoft.storage/storageaccounts/rgName4"),
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000002"),
	// 				CategoryType: to.Ptr(armcarbonoptimization.CategoryTypeEnumResource),
	// 				ResourceType: to.Ptr("microsoft.storage/storageaccounts"),
	// 				Location: to.Ptr("east us"),
	// 			},
	// 			&armcarbonoptimization.ResourceCarbonEmissionItemDetailData{
	// 				DataType: to.Ptr(armcarbonoptimization.ResponseDataTypeEnumResourceItemDetailsData),
	// 				LatestMonthEmissions: to.Ptr[float64](0.1),
	// 				PreviousMonthEmissions: to.Ptr[float64](0.05),
	// 				MonthOverMonthEmissionsChangeRatio: to.Ptr[float64](1),
	// 				MonthlyEmissionsChangeValue: to.Ptr[float64](0.05),
	// 				ItemName: to.Ptr("rgName5"),
	// 				ResourceGroup: to.Ptr("rgGroup"),
	// 				ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000002/resourceGroups/rgGroup/providers/microsoft.storage/storageaccounts/rgName5"),
	// 				SubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000002"),
	// 				CategoryType: to.Ptr(armcarbonoptimization.CategoryTypeEnumResource),
	// 				ResourceType: to.Ptr("microsoft.storage/storageaccounts"),
	// 				Location: to.Ptr("east us"),
	// 			},
	// 		},
	// 	},
	// }
}
