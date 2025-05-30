package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/SqlAssessmentV2Operations_Create_MaximumSet_Gen.json
func ExampleSQLAssessmentV2OperationsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSQLAssessmentV2OperationsClient().BeginCreate(ctx, "rgmigrate", "fci-test6904project", "test_fci_hadr", "test_swagger_1", armmigrationassessment.SQLAssessmentV2{
		Properties: &armmigrationassessment.SQLAssessmentV2Properties{
			AsyncCommitModeIntent: to.Ptr(armmigrationassessment.AsyncCommitModeIntentDisasterRecovery),
			AzureLocation:         to.Ptr("SoutheastAsia"),
			AzureOfferCode:        to.Ptr(armmigrationassessment.AzureOfferCodeMsazr0003P),
			AzureOfferCodeForVM:   to.Ptr(armmigrationassessment.AzureOfferCodeMsazr0003P),
			AzureSQLDatabaseSettings: &armmigrationassessment.SQLDbSettings{
				AzureSQLComputeTier:   to.Ptr(armmigrationassessment.ComputeTierAutomatic),
				AzureSQLDataBaseType:  to.Ptr(armmigrationassessment.AzureSQLDataBaseTypeSingleDatabase),
				AzureSQLPurchaseModel: to.Ptr(armmigrationassessment.AzureSQLPurchaseModelVCore),
				AzureSQLServiceTier:   to.Ptr(armmigrationassessment.AzureSQLServiceTierAutomatic),
			},
			AzureSQLManagedInstanceSettings: &armmigrationassessment.SQLMiSettings{
				AzureSQLInstanceType: to.Ptr(armmigrationassessment.AzureSQLInstanceTypeSingleInstance),
				AzureSQLServiceTier:  to.Ptr(armmigrationassessment.AzureSQLServiceTierAutomatic),
			},
			AzureSQLVMSettings: &armmigrationassessment.SQLVMSettings{
				InstanceSeries: []*armmigrationassessment.AzureVMFamily{
					to.Ptr(armmigrationassessment.AzureVMFamilyEadsv5Series)},
			},
			Currency:                 to.Ptr(armmigrationassessment.AzureCurrencyUSD),
			DisasterRecoveryLocation: to.Ptr(armmigrationassessment.AzureLocationEastAsia),
			DiscountPercentage:       to.Ptr[float32](0),
			EnableHadrAssessment:     to.Ptr(true),
			EntityUptime: &armmigrationassessment.EntityUptime{
				DaysPerMonth: to.Ptr[int32](30),
				HoursPerDay:  to.Ptr[int32](24),
			},
			EnvironmentType:       to.Ptr(armmigrationassessment.EnvironmentTypeProduction),
			MultiSubnetIntent:     to.Ptr(armmigrationassessment.MultiSubnetIntentDisasterRecovery),
			OptimizationLogic:     to.Ptr(armmigrationassessment.OptimizationLogicMinimizeCost),
			OSLicense:             to.Ptr(armmigrationassessment.OsLicenseUnknown),
			Percentile:            to.Ptr(armmigrationassessment.PercentilePercentile95),
			ReservedInstance:      to.Ptr(armmigrationassessment.AzureReservedInstanceNone),
			ReservedInstanceForVM: to.Ptr(armmigrationassessment.AzureReservedInstanceNone),
			ScalingFactor:         to.Ptr[float32](1),
			SizingCriterion:       to.Ptr(armmigrationassessment.AssessmentSizingCriterionPerformanceBased),
			SQLServerLicense:      to.Ptr(armmigrationassessment.SQLServerLicenseUnknown),
			TimeRange:             to.Ptr(armmigrationassessment.TimeRangeDay),
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
	// res.SQLAssessmentV2 = armmigrationassessment.SQLAssessmentV2{
	// 	Name: to.Ptr("test_swagger_1"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/sqlAssessments"),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.Migrate/assessmentprojects/fci-ankit-test6904project/groups/test_fci_hadr/sqlAssessments/test_swagger_1"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 	},
	// 	Properties: &armmigrationassessment.SQLAssessmentV2Properties{
	// 		AssessmentType: to.Ptr(armmigrationassessment.AssessmentTypeSQLAssessment),
	// 		AsyncCommitModeIntent: to.Ptr(armmigrationassessment.AsyncCommitModeIntentHighAvailability),
	// 		AzureLocation: to.Ptr("SoutheastAsia"),
	// 		AzureOfferCode: to.Ptr(armmigrationassessment.AzureOfferCodeMsazr0003P),
	// 		AzureOfferCodeForVM: to.Ptr(armmigrationassessment.AzureOfferCodeMsazr0003P),
	// 		AzureSecurityOfferingType: to.Ptr(armmigrationassessment.AzureSecurityOfferingTypeNO),
	// 		AzureSQLDatabaseSettings: &armmigrationassessment.SQLDbSettings{
	// 			AzureSQLComputeTier: to.Ptr(armmigrationassessment.ComputeTierProvisioned),
	// 			AzureSQLDataBaseType: to.Ptr(armmigrationassessment.AzureSQLDataBaseTypeSingleDatabase),
	// 			AzureSQLPurchaseModel: to.Ptr(armmigrationassessment.AzureSQLPurchaseModelVCore),
	// 			AzureSQLServiceTier: to.Ptr(armmigrationassessment.AzureSQLServiceTierAutomatic),
	// 		},
	// 		AzureSQLManagedInstanceSettings: &armmigrationassessment.SQLMiSettings{
	// 			AzureSQLInstanceType: to.Ptr(armmigrationassessment.AzureSQLInstanceTypeSingleInstance),
	// 			AzureSQLServiceTier: to.Ptr(armmigrationassessment.AzureSQLServiceTierAutomatic),
	// 		},
	// 		AzureSQLVMSettings: &armmigrationassessment.SQLVMSettings{
	// 			InstanceSeries: []*armmigrationassessment.AzureVMFamily{
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyDadsv5Series),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyDasv4Series),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyDdsv4Series),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyDdsv5Series),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyEadsv5Series),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyEasv4Series),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyEbdsv5Series),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyEbsv5Series),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyEdsv4Series),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyEdsv5Series),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyMSeries),
	// 				to.Ptr(armmigrationassessment.AzureVMFamilyMdsv2Series)},
	// 			},
	// 			ConfidenceRatingInPercentage: to.Ptr[float32](0),
	// 			CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-08T05:21:45.980Z"); return t}()),
	// 			Currency: to.Ptr(armmigrationassessment.AzureCurrencyUSD),
	// 			DisasterRecoveryLocation: to.Ptr(armmigrationassessment.AzureLocationEastAsia),
	// 			DiscountPercentage: to.Ptr[float32](0),
	// 			EnableHadrAssessment: to.Ptr(true),
	// 			EntityUptime: &armmigrationassessment.EntityUptime{
	// 				DaysPerMonth: to.Ptr[int32](31),
	// 				HoursPerDay: to.Ptr[int32](24),
	// 			},
	// 			EnvironmentType: to.Ptr(armmigrationassessment.EnvironmentTypeProduction),
	// 			GroupType: to.Ptr(armmigrationassessment.GroupTypeDefault),
	// 			IsInternetAccessAvailable: to.Ptr(false),
	// 			MultiSubnetIntent: to.Ptr(armmigrationassessment.MultiSubnetIntentHighAvailability),
	// 			OptimizationLogic: to.Ptr(armmigrationassessment.OptimizationLogicMinimizeCost),
	// 			OSLicense: to.Ptr(armmigrationassessment.OsLicenseYes),
	// 			Percentile: to.Ptr(armmigrationassessment.PercentilePercentile95),
	// 			PerfDataEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-08T05:21:45.964Z"); return t}()),
	// 			PerfDataStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-07T05:21:45.964Z"); return t}()),
	// 			PricesTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-15T16:00:00.000Z"); return t}()),
	// 			ReservedInstance: to.Ptr(armmigrationassessment.AzureReservedInstanceRI3Year),
	// 			ReservedInstanceForVM: to.Ptr(armmigrationassessment.AzureReservedInstanceRI3Year),
	// 			ScalingFactor: to.Ptr[float32](1),
	// 			SchemaVersion: to.Ptr("2.0"),
	// 			SizingCriterion: to.Ptr(armmigrationassessment.AssessmentSizingCriterionPerformanceBased),
	// 			SQLServerLicense: to.Ptr(armmigrationassessment.SQLServerLicenseYes),
	// 			Stage: to.Ptr(armmigrationassessment.AssessmentStageInProgress),
	// 			Status: to.Ptr(armmigrationassessment.AssessmentStatusCompleted),
	// 			TimeRange: to.Ptr(armmigrationassessment.TimeRangeDay),
	// 			UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-08T05:23:45.268Z"); return t}()),
	// 		},
	// 	}
}
