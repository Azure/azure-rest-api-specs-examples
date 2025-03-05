package armmigrationassessment_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/WebAppAssessmentV2Operations_Create_MaximumSet_Gen.json
func ExampleWebAppAssessmentV2OperationsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewWebAppAssessmentV2OperationsClient().BeginCreate(ctx, "rgopenapi", "sumukk-ccy-bcs4557project", "anraghun-selfhost-v2", "anraghun-selfhost-v2", armmigrationassessment.WebAppAssessmentV2{
		Properties: &armmigrationassessment.WebAppAssessmentV2Properties{
			AppSvcContainerSettings: &armmigrationassessment.AppSvcContainerSettings{
				IsolationRequired: to.Ptr(true),
			},
			AppSvcNativeSettings: &armmigrationassessment.AppSvcNativeSettings{
				IsolationRequired: to.Ptr(true),
			},
			AssessmentType:               to.Ptr(armmigrationassessment.AssessmentTypeWebAppAssessment),
			AzureLocation:                to.Ptr("UkWest"),
			AzureOfferCode:               to.Ptr(armmigrationassessment.AzureOfferCodeMsazr0003P),
			AzureSecurityOfferingType:    to.Ptr(armmigrationassessment.AzureSecurityOfferingTypeNO),
			ConfidenceRatingInPercentage: to.Ptr[float32](13),
			Currency:                     to.Ptr(armmigrationassessment.AzureCurrencyUSD),
			DiscountPercentage:           to.Ptr[float32](13),
			DiscoveredEntityLightSummary: &armmigrationassessment.DiscoveredEntityLightSummary{
				NumberOfMachines: to.Ptr[int32](27),
				NumberOfServers:  to.Ptr[int32](5),
				NumberOfWebApps:  to.Ptr[int32](23),
			},
			EaSubscriptionID: to.Ptr(""),
			EntityUptime: &armmigrationassessment.EntityUptime{
				DaysPerMonth: to.Ptr[int32](18),
				HoursPerDay:  to.Ptr[int32](13),
			},
			EnvironmentType:   to.Ptr(armmigrationassessment.EnvironmentTypeProduction),
			GroupType:         to.Ptr(armmigrationassessment.GroupTypeDefault),
			Percentile:        to.Ptr(armmigrationassessment.PercentilePercentile50),
			PerfDataEndTime:   to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:42:45.496Z"); return t }()),
			PerfDataStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:42:45.496Z"); return t }()),
			ReservedInstance:  to.Ptr(armmigrationassessment.AzureReservedInstanceNone),
			ScalingFactor:     to.Ptr[float32](17),
			SizingCriterion:   to.Ptr(armmigrationassessment.AssessmentSizingCriterionPerformanceBased),
			Stage:             to.Ptr(armmigrationassessment.AssessmentStageInProgress),
			Status:            to.Ptr(armmigrationassessment.AssessmentStatusCreated),
			TimeRange:         to.Ptr(armmigrationassessment.TimeRangeDay),
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
	// res.WebAppAssessmentV2 = armmigrationassessment.WebAppAssessmentV2{
	// 	Name: to.Ptr("anraghun-v2-test"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentProjects/groups/webAppAssessments"),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sumukk-ccy-bcs/providers/Microsoft.Migrate/assessmentProjects/sumukk-ccy-bcs4557project/groups/anraghun-selfhost-v2/webAppAssessments/anraghun-v2-test"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:41:40.597Z"); return t}()),
	// 		CreatedBy: to.Ptr("eatnpvdxza"),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:41:40.597Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("rnfkzoaidy"),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmigrationassessment.WebAppAssessmentV2Properties{
	// 		ProvisioningState: to.Ptr(armmigrationassessment.ProvisioningState2Succeeded),
	// 		AppSvcContainerSettings: &armmigrationassessment.AppSvcContainerSettings{
	// 			IsolationRequired: to.Ptr(true),
	// 		},
	// 		AppSvcNativeSettings: &armmigrationassessment.AppSvcNativeSettings{
	// 			IsolationRequired: to.Ptr(true),
	// 		},
	// 		AssessmentType: to.Ptr(armmigrationassessment.AssessmentTypeWebAppAssessment),
	// 		AzureLocation: to.Ptr("UkWest"),
	// 		AzureOfferCode: to.Ptr(armmigrationassessment.AzureOfferCodeMsazr0003P),
	// 		AzureSecurityOfferingType: to.Ptr(armmigrationassessment.AzureSecurityOfferingTypeNO),
	// 		ConfidenceRatingInPercentage: to.Ptr[float32](13),
	// 		CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:42:45.496Z"); return t}()),
	// 		Currency: to.Ptr(armmigrationassessment.AzureCurrencyUSD),
	// 		DiscountPercentage: to.Ptr[float32](13),
	// 		DiscoveredEntityLightSummary: &armmigrationassessment.DiscoveredEntityLightSummary{
	// 			NumberOfMachines: to.Ptr[int32](27),
	// 			NumberOfServers: to.Ptr[int32](5),
	// 			NumberOfWebApps: to.Ptr[int32](23),
	// 		},
	// 		EaSubscriptionID: to.Ptr(""),
	// 		EntityUptime: &armmigrationassessment.EntityUptime{
	// 			DaysPerMonth: to.Ptr[int32](18),
	// 			HoursPerDay: to.Ptr[int32](13),
	// 		},
	// 		EnvironmentType: to.Ptr(armmigrationassessment.EnvironmentTypeProduction),
	// 		GroupType: to.Ptr(armmigrationassessment.GroupTypeDefault),
	// 		Percentile: to.Ptr(armmigrationassessment.PercentilePercentile50),
	// 		PerfDataEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:42:45.496Z"); return t}()),
	// 		PerfDataStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:42:45.496Z"); return t}()),
	// 		PricesTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:42:45.496Z"); return t}()),
	// 		ReservedInstance: to.Ptr(armmigrationassessment.AzureReservedInstanceNone),
	// 		ScalingFactor: to.Ptr[float32](17),
	// 		SchemaVersion: to.Ptr("2.0"),
	// 		SizingCriterion: to.Ptr(armmigrationassessment.AssessmentSizingCriterionPerformanceBased),
	// 		Stage: to.Ptr(armmigrationassessment.AssessmentStageInProgress),
	// 		Status: to.Ptr(armmigrationassessment.AssessmentStatusCreated),
	// 		TimeRange: to.Ptr(armmigrationassessment.TimeRangeDay),
	// 		UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:42:45.496Z"); return t}()),
	// 	},
	// }
}
