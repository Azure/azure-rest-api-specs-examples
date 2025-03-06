package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AksAssessmentOperations_ListByAssessmentProject_MaximumSet_Gen.json
func ExampleAksAssessmentOperationsClient_NewListByAssessmentProjectPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAksAssessmentOperationsClient().NewListByAssessmentProjectPager("rgaksswagger", "testproject", &armmigrationassessment.AksAssessmentOperationsClientListByAssessmentProjectOptions{ContinuationToken: to.Ptr("polt"),
		Top:              to.Ptr[int32](5),
		Filter:           to.Ptr("azekdtdhupdngbqxzdppicwf"),
		TotalRecordCount: to.Ptr[int32](18),
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.AKSAssessmentListResult = armmigrationassessment.AKSAssessmentListResult{
		// 	Value: []*armmigrationassessment.AKSAssessment{
		// 		{
		// 			Name: to.Ptr("testaksassessment"),
		// 			Type: to.Ptr("AKSAssessment"),
		// 			ID: to.Ptr("/subscriptions/D6F60DF4-CE70-4E39-8217-B8FBE7CA85AA/resourceGroups/rgaksswagger/providers/Microsoft.Migrate/assessmentProjects/testproject/aksAssessments/testaksassessment"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-07T06:51:24.108Z"); return t}()),
		// 				CreatedBy: to.Ptr("User"),
		// 				CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-07T06:51:24.108Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("User"),
		// 				LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 			},
		// 			ETag: to.Ptr("00000000-0000-0000-a616-12d4724c01d9"),
		// 			Properties: &armmigrationassessment.AKSAssessmentProperties{
		// 				ProvisioningState: to.Ptr(armmigrationassessment.ProvisioningStatusSucceeded),
		// 				Scope: &armmigrationassessment.AssessmentScopeParameters{
		// 					ServerGroupID: to.Ptr("/subscriptions/D6F60DF4-CE70-4E39-8217-B8FBE7CA85AA/resourceGroups/rgaksswagger/providers/Microsoft.Migrate/assessmentProjects/testproject/groups/testgrp"),
		// 				},
		// 				Settings: &armmigrationassessment.AKSAssessmentSettings{
		// 					AzureLocation: to.Ptr("Unknown"),
		// 					Currency: to.Ptr(armmigrationassessment.AzureCurrencyUnknown),
		// 					DiscountPercentage: to.Ptr[float32](15),
		// 					EnvironmentType: to.Ptr(armmigrationassessment.AzureEnvironmentTypeUnknown),
		// 					LicensingProgram: to.Ptr(armmigrationassessment.LicensingProgramDefault),
		// 					PerformanceData: &armmigrationassessment.PerfDataSettings{
		// 						Percentile: to.Ptr(armmigrationassessment.PercentilePercentile50),
		// 						PerfDataEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-07T06:51:24.320Z"); return t}()),
		// 						PerfDataStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-07T06:51:24.320Z"); return t}()),
		// 						TimeRange: to.Ptr(armmigrationassessment.TimeRangeDay),
		// 					},
		// 					ScalingFactor: to.Ptr[float32](3),
		// 					SizingCriteria: to.Ptr(armmigrationassessment.AssessmentSizingCriterionPerformanceBased),
		// 					Category: to.Ptr(armmigrationassessment.AzureVMCategoryAll),
		// 					Consolidation: to.Ptr(armmigrationassessment.ConsolidationTypeFull),
		// 					PricingTier: to.Ptr(armmigrationassessment.PricingTierStandard),
		// 					SavingsOptions: to.Ptr(armmigrationassessment.SavingsOptionsNone),
		// 				},
		// 				Details: &armmigrationassessment.AKSAssessmentDetails{
		// 					ConfidenceRatingInPercentage: to.Ptr[float32](1),
		// 					CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-07T06:51:24.320Z"); return t}()),
		// 					PricesTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-07T06:51:24.320Z"); return t}()),
		// 					Status: to.Ptr(armmigrationassessment.AssessmentStatusCreated),
		// 					UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-07T06:51:24.320Z"); return t}()),
		// 					MachineCount: to.Ptr[int32](15),
		// 					TotalMonthlyCost: to.Ptr[float32](1),
		// 					WebAppCount: to.Ptr[int32](25),
		// 					WebServerCount: to.Ptr[int32](29),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
