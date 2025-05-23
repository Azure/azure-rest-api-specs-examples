package armmigrationassessment_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AssessmentsOperations_Create_MaximumSet_Gen.json
func ExampleAssessmentsOperationsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAssessmentsOperationsClient().BeginCreate(ctx, "ayagrawrg", "app18700project", "kuchatur-test", "asm1", armmigrationassessment.Assessment{
		Properties: &armmigrationassessment.MachineAssessmentProperties{
			ProvisioningState: to.Ptr(armmigrationassessment.ProvisioningStateSucceeded),
			AssessmentType:    to.Ptr(armmigrationassessment.AssessmentTypeUnknown),
			AzureDiskTypes: []*armmigrationassessment.AzureDiskType{
				to.Ptr(armmigrationassessment.AzureDiskTypePremium),
				to.Ptr(armmigrationassessment.AzureDiskTypePremiumV2),
				to.Ptr(armmigrationassessment.AzureDiskTypeStandardSSD)},
			AzureHybridUseBenefit:  to.Ptr(armmigrationassessment.AzureHybridUseBenefitUnknown),
			AzureLocation:          to.Ptr("njxbwdtsxzhichsnk"),
			AzureOfferCode:         to.Ptr(armmigrationassessment.AzureOfferCodeUnknown),
			AzurePricingTier:       to.Ptr(armmigrationassessment.AzurePricingTierStandard),
			AzureStorageRedundancy: to.Ptr(armmigrationassessment.AzureStorageRedundancyUnknown),
			AzureVMFamilies: []*armmigrationassessment.AzureVMFamily{
				to.Ptr(armmigrationassessment.AzureVMFamilyDSeries),
				to.Ptr(armmigrationassessment.AzureVMFamilyLsv2Series),
				to.Ptr(armmigrationassessment.AzureVMFamilyMSeries),
				to.Ptr(armmigrationassessment.AzureVMFamilyMdsv2Series),
				to.Ptr(armmigrationassessment.AzureVMFamilyMsv2Series),
				to.Ptr(armmigrationassessment.AzureVMFamilyMv2Series)},
			Currency:                   to.Ptr(armmigrationassessment.AzureCurrencyUnknown),
			DiscountPercentage:         to.Ptr[float32](6),
			EaSubscriptionID:           to.Ptr("kwsu"),
			GroupType:                  to.Ptr(armmigrationassessment.GroupTypeDefault),
			LinuxAzureHybridUseBenefit: to.Ptr(armmigrationassessment.AzureHybridUseBenefitUnknown),
			Percentile:                 to.Ptr(armmigrationassessment.PercentilePercentile50),
			PerfDataEndTime:            to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T09:36:48.491Z"); return t }()),
			PerfDataStartTime:          to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T09:36:48.491Z"); return t }()),
			ReservedInstance:           to.Ptr(armmigrationassessment.AzureReservedInstanceNone),
			ScalingFactor:              to.Ptr[float32](24),
			SizingCriterion:            to.Ptr(armmigrationassessment.AssessmentSizingCriterionPerformanceBased),
			Stage:                      to.Ptr(armmigrationassessment.AssessmentStageInProgress),
			Status:                     to.Ptr(armmigrationassessment.AssessmentStatusCreated),
			TimeRange:                  to.Ptr(armmigrationassessment.TimeRangeDay),
			VMUptime: &armmigrationassessment.VMUptime{
				DaysPerMonth: to.Ptr[int32](13),
				HoursPerDay:  to.Ptr[int32](26),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Assessment = armmigrationassessment.Assessment{
	// 	Name: to.Ptr("asm1"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/assessments"),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawrg/providers/Microsoft.Migrate/assessmentprojects/app18700project/groups/kuchatur-test/assessments/asm1"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T09:36:29.583Z"); return t}()),
	// 		CreatedBy: to.Ptr("sakanwar"),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T09:36:29.583Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("sakanwar"),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmigrationassessment.MachineAssessmentProperties{
	// 		ProvisioningState: to.Ptr(armmigrationassessment.ProvisioningStateSucceeded),
	// 		AssessmentErrorSummary: map[string]*int32{
	// 		},
	// 		AssessmentType: to.Ptr(armmigrationassessment.AssessmentTypeMachineAssessment),
	// 		AzureDiskTypes: []*armmigrationassessment.AzureDiskType{
	// 			to.Ptr(armmigrationassessment.AzureDiskTypePremium),
	// 			to.Ptr(armmigrationassessment.AzureDiskTypePremiumV2),
	// 			to.Ptr(armmigrationassessment.AzureDiskTypeStandardSSD)},
	// 			AzureHybridUseBenefit: to.Ptr(armmigrationassessment.AzureHybridUseBenefitYes),
	// 			AzureLocation: to.Ptr("SoutheastAsia"),
	// 			AzureOfferCode: to.Ptr(armmigrationassessment.AzureOfferCodeMsazr0003P),
	// 			AzurePricingTier: to.Ptr(armmigrationassessment.AzurePricingTierStandard),
	// 			AzureStorageRedundancy: to.Ptr(armmigrationassessment.AzureStorageRedundancyLocallyRedundant),
	// 			AzureVMFamilies: []*armmigrationassessment.AzureVMFamily{
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyDSeries),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyLsv2Series),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyMSeries),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyMdsv2Series),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyMsv2Series),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyMv2Series)},
	// 				ConfidenceRatingInPercentage: to.Ptr[float32](28),
	// 				CostComponents: []*armmigrationassessment.CostComponent{
	// 					{
	// 						Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlyAzureHybridCostSavings),
	// 						Value: to.Ptr[float32](547.584),
	// 					},
	// 					{
	// 						Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlySecurityCost),
	// 						Value: to.Ptr[float32](44.64),
	// 					},
	// 					{
	// 						Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlyPremiumV2StorageCost),
	// 						Value: to.Ptr[float32](25.141248),
	// 				}},
	// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T09:14:05.732Z"); return t}()),
	// 				Currency: to.Ptr(armmigrationassessment.AzureCurrencyUSD),
	// 				DiscountPercentage: to.Ptr[float32](6),
	// 				DistributionByOsName: map[string]*int32{
	// 					"microsoft Windows Server 2016 or later (64-bit)": to.Ptr[int32](3),
	// 				},
	// 				DistributionByServicePackInsight: map[string]*int32{
	// 				},
	// 				DistributionBySupportStatus: map[string]*int32{
	// 					"extended": to.Ptr[int32](3),
	// 				},
	// 				GroupType: to.Ptr(armmigrationassessment.GroupTypeDefault),
	// 				LinuxAzureHybridUseBenefit: to.Ptr(armmigrationassessment.AzureHybridUseBenefitYes),
	// 				MonthlyBandwidthCost: to.Ptr[float32](21),
	// 				MonthlyComputeCost: to.Ptr[float32](10),
	// 				MonthlyPremiumStorageCost: to.Ptr[float32](13),
	// 				MonthlyStandardSsdStorageCost: to.Ptr[float32](3),
	// 				MonthlyStorageCost: to.Ptr[float32](20),
	// 				MonthlyUltraStorageCost: to.Ptr[float32](21),
	// 				NumberOfMachines: to.Ptr[int32](3),
	// 				Percentile: to.Ptr(armmigrationassessment.PercentilePercentile95),
	// 				PerfDataEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T09:36:48.491Z"); return t}()),
	// 				PerfDataStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T09:36:48.491Z"); return t}()),
	// 				PricesTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-15T16:00:00.000Z"); return t}()),
	// 				ReservedInstance: to.Ptr(armmigrationassessment.AzureReservedInstanceRI3Year),
	// 				ScalingFactor: to.Ptr[float32](24),
	// 				SizingCriterion: to.Ptr(armmigrationassessment.AssessmentSizingCriterionPerformanceBased),
	// 				Stage: to.Ptr(armmigrationassessment.AssessmentStageInProgress),
	// 				Status: to.Ptr(armmigrationassessment.AssessmentStatusCreated),
	// 				SuitabilitySummary: map[string]*int32{
	// 					"suitable": to.Ptr[int32](3),
	// 				},
	// 				TimeRange: to.Ptr(armmigrationassessment.TimeRangeDay),
	// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T09:14:05.732Z"); return t}()),
	// 				VMUptime: &armmigrationassessment.VMUptime{
	// 					DaysPerMonth: to.Ptr[int32](13),
	// 					HoursPerDay: to.Ptr[int32](20),
	// 				},
	// 			},
	// 		}
}
