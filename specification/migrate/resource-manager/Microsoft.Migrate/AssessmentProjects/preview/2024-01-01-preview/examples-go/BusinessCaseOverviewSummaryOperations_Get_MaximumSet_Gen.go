package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/BusinessCaseOverviewSummaryOperations_Get_MaximumSet_Gen.json
func ExampleBusinessCaseOverviewSummaryOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBusinessCaseOverviewSummaryOperationsClient().Get(ctx, "rgopenapi", "multipleto8617project", "sample-business-case", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OverviewSummary = armmigrationassessment.OverviewSummary{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentProjects/businessCases/overviewSummaries"),
	// 	ID: to.Ptr("/subscriptions/ADC896AD-6A38-454E-9A62-AFC618F5F4BC/resourceGroups/rgopenapi/providers/Microsoft.Migrate/assessmentProjects/multipleto8617project/businessCases/sample-business-case/overviewSummaries/default"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-08T07:09:55.036Z"); return t}()),
	// 		CreatedBy: to.Ptr("t72jdt@company.com"),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-08T07:09:55.036Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("t72jdt@company.com"),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmigrationassessment.OverviewSummaryProperties{
	// 		AzureArcEnabledOnPremisesCost: to.Ptr[float32](6),
	// 		AzureArcServicesCost: to.Ptr[float32](27),
	// 		EsuSavingsFor4Years: to.Ptr[float32](20),
	// 		FutureAzureArcEnabledOnPremisesCost: to.Ptr[float32](13),
	// 		FutureAzureArcServicesCost: to.Ptr[float32](7),
	// 		FutureAzureIaasCost: to.Ptr[float32](23),
	// 		FutureAzurePaasCost: to.Ptr[float32](28),
	// 		FutureCostIncludingAzureArc: to.Ptr[float32](27),
	// 		FutureEsuSavingsFor4YearsIncludingAzureArc: to.Ptr[float32](7),
	// 		FutureManagementCostSavingsIncludingAzureArc: to.Ptr[float32](24),
	// 		FutureSecurityCostSavingsIncludingAzureArc: to.Ptr[float32](18),
	// 		IaasOsDistribution: &armmigrationassessment.IaasOsDistribution{
	// 			Linux: to.Ptr[int32](22),
	// 			Other: to.Ptr[int32](14),
	// 			Windows: to.Ptr[int32](5),
	// 		},
	// 		LinuxAhubSavings: to.Ptr[float32](26),
	// 		ManagementCostSavings: to.Ptr[float32](19),
	// 		OSSupportStatusDistribution: &armmigrationassessment.SupportStatusDistribution{
	// 			Extended: to.Ptr[int32](26),
	// 			MainStream: to.Ptr[int32](20),
	// 			OutOfSupport: to.Ptr[int32](29),
	// 			UnknownSupport: to.Ptr[int32](2),
	// 		},
	// 		PaasDistribution: &armmigrationassessment.PaasDistribution{
	// 			IisWebServers: to.Ptr[int32](9),
	// 			SQLServers: to.Ptr[int32](28),
	// 			WebAppDistribution: map[string]*int32{
	// 				"key683": to.Ptr[int32](28),
	// 			},
	// 		},
	// 		SecurityCostSavings: to.Ptr[float32](8),
	// 		ServersDiscovered: &armmigrationassessment.ServersDiscovered{
	// 			HyperV: to.Ptr[int32](28),
	// 			NotApplicable: to.Ptr[int32](13),
	// 			Physical: to.Ptr[int32](16),
	// 			VMWare: to.Ptr[int32](6),
	// 		},
	// 		SQLAhubSavings: to.Ptr[float32](29),
	// 		SQLSupportStatusDistribution: &armmigrationassessment.SupportStatusDistribution{
	// 			Extended: to.Ptr[int32](26),
	// 			MainStream: to.Ptr[int32](20),
	// 			OutOfSupport: to.Ptr[int32](29),
	// 			UnknownSupport: to.Ptr[int32](2),
	// 		},
	// 		TotalAzureAvsCost: to.Ptr[float32](21),
	// 		TotalAzureCost: to.Ptr[float32](20),
	// 		TotalAzureIaasCost: to.Ptr[float32](29),
	// 		TotalAzurePaasCost: to.Ptr[float32](18),
	// 		TotalOnPremisesCost: to.Ptr[float32](20),
	// 		UtilizationData: &armmigrationassessment.UtilizationData{
	// 			NumberOfActiveEntities: to.Ptr[int32](10),
	// 			NumberOfDecommisionEntities: to.Ptr[int32](10),
	// 			NumberOfInactiveEntities: to.Ptr[int32](18),
	// 			NumberOfUnknownEntities: to.Ptr[int32](13),
	// 		},
	// 		WindowsAhubSavings: to.Ptr[float32](26),
	// 		YearOnYearEstimates: &armmigrationassessment.YearOnYearEstimates{
	// 			AzureArcEnabledOnPremisesCost: []*armmigrationassessment.YearOnYearCost{
	// 				{
	// 					Cost: to.Ptr[float32](8646893.36496789),
	// 					Year: to.Ptr(armmigrationassessment.YearYear0),
	// 				},
	// 				{
	// 					Cost: to.Ptr[float32](6452444.63321628),
	// 					Year: to.Ptr(armmigrationassessment.YearYear1),
	// 				},
	// 				{
	// 					Cost: to.Ptr[float32](7361476.3148771),
	// 					Year: to.Ptr(armmigrationassessment.YearYear2),
	// 				},
	// 				{
	// 					Cost: to.Ptr[float32](7893482.58062095),
	// 					Year: to.Ptr(armmigrationassessment.YearYear3),
	// 			}},
	// 			AzureCapexCost: map[string]*float32{
	// 				"Year0": to.Ptr[float32](1049.81370090828),
	// 				"Year1": to.Ptr[float32](881.843508762955),
	// 				"Year2": to.Ptr[float32](578.709802625689),
	// 				"Year3": to.Ptr[float32](0),
	// 			},
	// 			AzureCost: map[string]*float32{
	// 				"Year0": to.Ptr[float32](351803.869241272),
	// 				"Year1": to.Ptr[float32](374738.191162668),
	// 				"Year2": to.Ptr[float32](391989.235419251),
	// 				"Year3": to.Ptr[float32](396114.705),
	// 			},
	// 			AzureOpexCost: map[string]*float32{
	// 				"Year0": to.Ptr[float32](350754.055540363),
	// 				"Year1": to.Ptr[float32](373856.347653905),
	// 				"Year2": to.Ptr[float32](391410.525616625),
	// 				"Year3": to.Ptr[float32](396114.705),
	// 			},
	// 			FutureAzureArcEnabledOnPremisesEsuCost: []*armmigrationassessment.YearOnYearCost{
	// 				{
	// 					Cost: to.Ptr[float32](2777682.6),
	// 					Year: to.Ptr(armmigrationassessment.YearYear0),
	// 				},
	// 				{
	// 					Cost: to.Ptr[float32](217405.65),
	// 					Year: to.Ptr(armmigrationassessment.YearYear1),
	// 				},
	// 				{
	// 					Cost: to.Ptr[float32](334683.25),
	// 					Year: to.Ptr(armmigrationassessment.YearYear2),
	// 				},
	// 				{
	// 					Cost: to.Ptr[float32](0),
	// 					Year: to.Ptr(armmigrationassessment.YearYear3),
	// 			}},
	// 			FutureCost: []*armmigrationassessment.YearOnYearCost{
	// 				{
	// 					Cost: to.Ptr[float32](8619047.96496789),
	// 					Year: to.Ptr(armmigrationassessment.YearYear0),
	// 				},
	// 				{
	// 					Cost: to.Ptr[float32](5847511.33627395),
	// 					Year: to.Ptr(armmigrationassessment.YearYear1),
	// 				},
	// 				{
	// 					Cost: to.Ptr[float32](5363132.85669087),
	// 					Year: to.Ptr(armmigrationassessment.YearYear2),
	// 				},
	// 				{
	// 					Cost: to.Ptr[float32](3616793.89850464),
	// 					Year: to.Ptr(armmigrationassessment.YearYear3),
	// 			}},
	// 			FutureOnPremisesEsuCost: []*armmigrationassessment.YearOnYearCost{
	// 				{
	// 					Cost: to.Ptr[float32](3742242.49),
	// 					Year: to.Ptr(armmigrationassessment.YearYear0),
	// 				},
	// 				{
	// 					Cost: to.Ptr[float32](413811.16),
	// 					Year: to.Ptr(armmigrationassessment.YearYear1),
	// 				},
	// 				{
	// 					Cost: to.Ptr[float32](250961.91),
	// 					Year: to.Ptr(armmigrationassessment.YearYear2),
	// 				},
	// 				{
	// 					Cost: to.Ptr[float32](0),
	// 					Year: to.Ptr(armmigrationassessment.YearYear3),
	// 			}},
	// 			OnPremisesCapexCost: map[string]*float32{
	// 				"Year0": to.Ptr[float32](1049.81370090828),
	// 				"Year1": to.Ptr[float32](1102.30438595369),
	// 				"Year2": to.Ptr[float32](1157.41960525138),
	// 				"Year3": to.Ptr[float32](1215.29058551395),
	// 			},
	// 			OnPremisesCost: map[string]*float32{
	// 				"Year0": to.Ptr[float32](351803.869241272),
	// 				"Year1": to.Ptr[float32](369394.062703335),
	// 				"Year2": to.Ptr[float32](387863.765838502),
	// 				"Year3": to.Ptr[float32](408410.454130427),
	// 			},
	// 			OnPremisesOpexCost: map[string]*float32{
	// 				"Year0": to.Ptr[float32](350754.055540363),
	// 				"Year1": to.Ptr[float32](368291.758317382),
	// 				"Year2": to.Ptr[float32](386706.346233251),
	// 				"Year3": to.Ptr[float32](407195.163544913),
	// 			},
	// 			PaybackPeriod: to.Ptr[int32](0),
	// 			Savings: map[string]*float32{
	// 				"Year0": to.Ptr[float32](0),
	// 				"Year1": to.Ptr[float32](-5344.12845933292),
	// 				"Year2": to.Ptr[float32](-4125.46958074899),
	// 				"Year3": to.Ptr[float32](12295.7491304271),
	// 			},
	// 		},
	// 	},
	// }
}
