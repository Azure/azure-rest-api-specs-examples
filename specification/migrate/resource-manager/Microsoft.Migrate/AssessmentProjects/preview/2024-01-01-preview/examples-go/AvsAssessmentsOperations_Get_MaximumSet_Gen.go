package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AvsAssessmentsOperations_Get_MaximumSet_Gen.json
func ExampleAvsAssessmentsOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAvsAssessmentsOperationsClient().Get(ctx, "ayagrawrg", "app18700project", "kuchatur-test", "asm2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AvsAssessment = armmigrationassessment.AvsAssessment{
	// 	Name: to.Ptr("asm2"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/avsAssessments"),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawrg/providers/Microsoft.Migrate/assessmentprojects/app18700project/groups/kuchatur-test/avsAssessments/asm2"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-15T18:24:02.548Z"); return t}()),
	// 		CreatedBy: to.Ptr("sjghfaivpyryuvovuvcqegaorvhp"),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-15T18:24:02.548Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("fjkmzd"),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmigrationassessment.AvsAssessmentProperties{
	// 		ProvisioningState: to.Ptr(armmigrationassessment.ProvisioningStateSucceeded),
	// 		AssessmentErrorSummary: map[string]*int32{
	// 		},
	// 		AssessmentType: to.Ptr(armmigrationassessment.AssessmentTypeAvsAssessment),
	// 		AvsAssessmentScenario: to.Ptr(armmigrationassessment.AvsAssessmentScenarioNewAvsSddc),
	// 		AvsEstimatedExternalStorages: []*armmigrationassessment.AvsEstimatedExternalStorage{
	// 			{
	// 				MonthlyPrice: to.Ptr[float32](250),
	// 				StorageType: to.Ptr(armmigrationassessment.ExternalStorageTypeAnfStandard),
	// 				StorageUtilization: to.Ptr[float32](97),
	// 				TotalStorageInGB: to.Ptr[float32](1024),
	// 			},
	// 			{
	// 				MonthlyPrice: to.Ptr[float32](350),
	// 				StorageType: to.Ptr(armmigrationassessment.ExternalStorageTypeAnfPremium),
	// 				StorageUtilization: to.Ptr[float32](78),
	// 				TotalStorageInGB: to.Ptr[float32](1024),
	// 		}},
	// 		AvsEstimatedNetworks: []*armmigrationassessment.AvsEstimatedNetwork{
	// 			{
	// 				MonthlyPrice: to.Ptr[float32](120),
	// 				NetworkType: to.Ptr(armmigrationassessment.NetworkSKUTypeExpressRouteUltra),
	// 		}},
	// 		AvsEstimatedNodes: []*armmigrationassessment.AvsEstimatedNode{
	// 			{
	// 				CPUUtilization: to.Ptr[float32](46.3),
	// 				FttRaidLevel: to.Ptr(armmigrationassessment.FttAndRaidLevelFtt1Raid1),
	// 				MonthlyPrice: to.Ptr[float32](9088.5),
	// 				NodeNumber: to.Ptr[int32](3),
	// 				NodeType: to.Ptr(armmigrationassessment.AzureAvsNodeTypeAV36),
	// 				PricingModel: to.Ptr(armmigrationassessment.AzureReservedInstanceRI3Year),
	// 				RAMUtilization: to.Ptr[float32](12.73),
	// 				StorageUtilization: to.Ptr[float32](76.99),
	// 				TotalCPU: to.Ptr[float32](108),
	// 				TotalRAM: to.Ptr[float32](1728),
	// 				TotalStorage: to.Ptr[float32](46080),
	// 		}},
	// 		AzureLocation: to.Ptr(armmigrationassessment.AzureLocationEastUs),
	// 		AzureOfferCode: to.Ptr(armmigrationassessment.AzureOfferCodeMsazr0003P),
	// 		ConfidenceRatingInPercentage: to.Ptr[float32](100),
	// 		CostComponents: []*armmigrationassessment.CostComponent{
	// 			{
	// 				Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlyAvsNodeCost),
	// 				Description: to.Ptr("Monthly Cost of AVS Nodes."),
	// 				Value: to.Ptr[float32](9088.5),
	// 			},
	// 			{
	// 				Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlyAvsExternalStorageCost),
	// 				Description: to.Ptr("Monthly AVS External Storage Cost for Assessment."),
	// 				Value: to.Ptr[float32](600),
	// 			},
	// 			{
	// 				Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlyAvsNetworkCost),
	// 				Description: to.Ptr("Monthly AVS Cost of Network."),
	// 				Value: to.Ptr[float32](120),
	// 			},
	// 			{
	// 				Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlyVcfByolCostDifference),
	// 				Description: to.Ptr("Monthly compute cost difference between VCF BYOL and non VCF BYOL SKUs"),
	// 				Value: to.Ptr[float32](7088.5),
	// 		}},
	// 		CPUHeadroom: to.Ptr[float32](12),
	// 		CPUUtilization: to.Ptr[float32](46.3),
	// 		CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T13:35:56.598Z"); return t}()),
	// 		Currency: to.Ptr(armmigrationassessment.AzureCurrencyUSD),
	// 		DedupeCompression: to.Ptr[float32](1.5),
	// 		DiscountPercentage: to.Ptr[float32](0),
	// 		ExternalStorageTypes: []*armmigrationassessment.ExternalStorageType{
	// 			to.Ptr(armmigrationassessment.ExternalStorageTypeAnfStandard),
	// 			to.Ptr(armmigrationassessment.ExternalStorageTypeAnfPremium),
	// 			to.Ptr(armmigrationassessment.ExternalStorageTypeAnfUltra)},
	// 			FailuresToTolerateAndRaidLevel: to.Ptr(armmigrationassessment.FttAndRaidLevelUnknown),
	// 			FailuresToTolerateAndRaidLevelList: []*armmigrationassessment.FttAndRaidLevel{
	// 				to.Ptr(armmigrationassessment.FttAndRaidLevelFtt1Raid1),
	// 				to.Ptr(armmigrationassessment.FttAndRaidLevelFtt1Raid5),
	// 				to.Ptr(armmigrationassessment.FttAndRaidLevelFtt3Raid1)},
	// 				GroupType: to.Ptr(armmigrationassessment.GroupTypeDefault),
	// 				IsStretchClusterEnabled: to.Ptr(false),
	// 				IsVcfByolEnabled: to.Ptr(true),
	// 				LimitingFactor: to.Ptr("Storage"),
	// 				MemOvercommit: to.Ptr[float32](1),
	// 				NodeType: to.Ptr(armmigrationassessment.AzureAvsNodeTypeUnknown),
	// 				NodeTypes: []*armmigrationassessment.AzureAvsNodeType{
	// 					to.Ptr(armmigrationassessment.AzureAvsNodeTypeAV36),
	// 					to.Ptr(armmigrationassessment.AzureAvsNodeTypeAV36P),
	// 					to.Ptr(armmigrationassessment.AzureAvsNodeTypeAV52)},
	// 					NumberOfMachines: to.Ptr[int32](3),
	// 					NumberOfNodes: to.Ptr[int32](3),
	// 					Percentile: to.Ptr(armmigrationassessment.PercentilePercentile95),
	// 					PerfDataEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T13:35:56.567Z"); return t}()),
	// 					PerfDataStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-25T13:35:56.567Z"); return t}()),
	// 					PricesTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-15T16:00:00.000Z"); return t}()),
	// 					RAMUtilization: to.Ptr[float32](12.73),
	// 					ReservedInstance: to.Ptr(armmigrationassessment.AzureReservedInstanceRI3Year),
	// 					ScalingFactor: to.Ptr[float32](1),
	// 					SizingCriterion: to.Ptr(armmigrationassessment.AssessmentSizingCriterionAsOnPremises),
	// 					Stage: to.Ptr(armmigrationassessment.AssessmentStageInProgress),
	// 					Status: to.Ptr(armmigrationassessment.AssessmentStatusCompleted),
	// 					StorageUtilization: to.Ptr[float32](77.43),
	// 					Suitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
	// 					SuitabilityExplanation: to.Ptr(armmigrationassessment.AzureAvsSuitabilityExplanationNotApplicable),
	// 					SuitabilitySummary: map[string]*int32{
	// 						"conditionallySuitable": to.Ptr[int32](3),
	// 					},
	// 					TimeRange: to.Ptr(armmigrationassessment.TimeRangeDay),
	// 					TotalCPUCores: to.Ptr[float32](108),
	// 					TotalMonthlyCost: to.Ptr[float32](9808.5),
	// 					TotalRAMInGB: to.Ptr[float32](1728),
	// 					TotalStorageInGB: to.Ptr[float32](48128),
	// 					UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-26T13:36:01.169Z"); return t}()),
	// 					VcpuOversubscription: to.Ptr[float32](4),
	// 				},
	// 			}
}
