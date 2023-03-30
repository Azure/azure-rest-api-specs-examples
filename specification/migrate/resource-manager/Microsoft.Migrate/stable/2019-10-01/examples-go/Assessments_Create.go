package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/Assessments_Create.json
func ExampleAssessmentsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessmentsClient().Create(ctx, "abgoyal-westEurope", "abgoyalWEselfhostb72bproject", "Group2", "assessment_5_14_2019_16_48_47", &armmigrate.AssessmentsClientCreateOptions{Assessment: &armmigrate.Assessment{
		ETag: to.Ptr("\"1e000c2c-0000-0d00-0000-5cdaa4190000\""),
		Properties: &armmigrate.AssessmentProperties{
			AzureDiskType:          to.Ptr(armmigrate.AzureDiskTypeStandardOrPremium),
			AzureHybridUseBenefit:  to.Ptr(armmigrate.AzureHybridUseBenefitYes),
			AzureLocation:          to.Ptr(armmigrate.AzureLocationNorthEurope),
			AzureOfferCode:         to.Ptr(armmigrate.AzureOfferCodeMSAZR0003P),
			AzurePricingTier:       to.Ptr(armmigrate.AzurePricingTierStandard),
			AzureStorageRedundancy: to.Ptr(armmigrate.AzureStorageRedundancyLocallyRedundant),
			AzureVMFamilies: []*armmigrate.AzureVMFamily{
				to.Ptr(armmigrate.AzureVMFamilyDv2Series),
				to.Ptr(armmigrate.AzureVMFamilyFSeries),
				to.Ptr(armmigrate.AzureVMFamilyDv3Series),
				to.Ptr(armmigrate.AzureVMFamilyDSSeries),
				to.Ptr(armmigrate.AzureVMFamilyDSv2Series),
				to.Ptr(armmigrate.AzureVMFamilyFsSeries),
				to.Ptr(armmigrate.AzureVMFamilyDsv3Series),
				to.Ptr(armmigrate.AzureVMFamilyEv3Series),
				to.Ptr(armmigrate.AzureVMFamilyEsv3Series),
				to.Ptr(armmigrate.AzureVMFamilyDSeries),
				to.Ptr(armmigrate.AzureVMFamilyMSeries),
				to.Ptr(armmigrate.AzureVMFamilyFsv2Series),
				to.Ptr(armmigrate.AzureVMFamilyHSeries)},
			Currency:           to.Ptr(armmigrate.CurrencyUSD),
			DiscountPercentage: to.Ptr[float64](100),
			Percentile:         to.Ptr(armmigrate.PercentilePercentile95),
			ReservedInstance:   to.Ptr(armmigrate.ReservedInstanceRI3Year),
			ScalingFactor:      to.Ptr[float64](1),
			SizingCriterion:    to.Ptr(armmigrate.AssessmentSizingCriterionPerformanceBased),
			Stage:              to.Ptr(armmigrate.AssessmentStageInProgress),
			TimeRange:          to.Ptr(armmigrate.TimeRangeDay),
			VMUptime: &armmigrate.VMUptime{
				DaysPerMonth: to.Ptr[int32](31),
				HoursPerDay:  to.Ptr[int32](24),
			},
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Assessment = armmigrate.Assessment{
	// 	Name: to.Ptr("assessment_5_14_2019_16_48_47"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/assessments"),
	// 	ETag: to.Ptr("\"1e000c2c-0000-0d00-0000-5cdaa4190000\""),
	// 	ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westeurope/providers/Microsoft.Migrate/assessmentprojects/abgoyalWEselfhostb72bproject/groups/Group2/assessments/assessment_5_14_2019_16_48_47"),
	// 	Properties: &armmigrate.AssessmentProperties{
	// 		AzureDiskType: to.Ptr(armmigrate.AzureDiskTypeStandardOrPremium),
	// 		AzureHybridUseBenefit: to.Ptr(armmigrate.AzureHybridUseBenefitYes),
	// 		AzureLocation: to.Ptr(armmigrate.AzureLocationNorthEurope),
	// 		AzureOfferCode: to.Ptr(armmigrate.AzureOfferCodeMSAZR0003P),
	// 		AzurePricingTier: to.Ptr(armmigrate.AzurePricingTierStandard),
	// 		AzureStorageRedundancy: to.Ptr(armmigrate.AzureStorageRedundancyLocallyRedundant),
	// 		AzureVMFamilies: []*armmigrate.AzureVMFamily{
	// 			to.Ptr(armmigrate.AzureVMFamilyDv2Series),
	// 			to.Ptr(armmigrate.AzureVMFamilyFSeries),
	// 			to.Ptr(armmigrate.AzureVMFamilyDv3Series),
	// 			to.Ptr(armmigrate.AzureVMFamilyDSSeries),
	// 			to.Ptr(armmigrate.AzureVMFamilyDSv2Series),
	// 			to.Ptr(armmigrate.AzureVMFamilyFsSeries),
	// 			to.Ptr(armmigrate.AzureVMFamilyDsv3Series),
	// 			to.Ptr(armmigrate.AzureVMFamilyEv3Series),
	// 			to.Ptr(armmigrate.AzureVMFamilyEsv3Series),
	// 			to.Ptr(armmigrate.AzureVMFamilyDSeries),
	// 			to.Ptr(armmigrate.AzureVMFamilyMSeries),
	// 			to.Ptr(armmigrate.AzureVMFamilyFsv2Series),
	// 			to.Ptr(armmigrate.AzureVMFamilyHSeries)},
	// 			ConfidenceRatingInPercentage: to.Ptr[float64](0),
	// 			CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-14T11:18:47.7893715Z"); return t}()),
	// 			Currency: to.Ptr(armmigrate.CurrencyUSD),
	// 			DiscountPercentage: to.Ptr[float64](0),
	// 			MonthlyBandwidthCost: to.Ptr[float64](0),
	// 			MonthlyComputeCost: to.Ptr[float64](607.443264),
	// 			MonthlyPremiumStorageCost: to.Ptr[float64](0),
	// 			MonthlyStandardSSDStorageCost: to.Ptr[float64](0),
	// 			MonthlyStorageCost: to.Ptr[float64](111.36),
	// 			NumberOfMachines: to.Ptr[int32](5),
	// 			Percentile: to.Ptr(armmigrate.PercentilePercentile95),
	// 			PerfDataEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-14T11:18:47.7893715Z"); return t}()),
	// 			PerfDataStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-13T11:18:47.7893715Z"); return t}()),
	// 			PricesTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-03-26T11:07:37.139768Z"); return t}()),
	// 			ReservedInstance: to.Ptr(armmigrate.ReservedInstanceRI3Year),
	// 			ScalingFactor: to.Ptr[float64](1),
	// 			SizingCriterion: to.Ptr(armmigrate.AssessmentSizingCriterionPerformanceBased),
	// 			Stage: to.Ptr(armmigrate.AssessmentStageInProgress),
	// 			Status: to.Ptr(armmigrate.AssessmentStatusCompleted),
	// 			TimeRange: to.Ptr(armmigrate.TimeRangeDay),
	// 			UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-14T11:18:47.7893715Z"); return t}()),
	// 			VMUptime: &armmigrate.VMUptime{
	// 				DaysPerMonth: to.Ptr[int32](31),
	// 				HoursPerDay: to.Ptr[int32](24),
	// 			},
	// 		},
	// 	}
}
