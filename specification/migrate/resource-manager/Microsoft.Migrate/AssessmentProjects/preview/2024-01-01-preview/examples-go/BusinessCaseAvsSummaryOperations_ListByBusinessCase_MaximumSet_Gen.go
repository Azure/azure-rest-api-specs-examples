package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/BusinessCaseAvsSummaryOperations_ListByBusinessCase_MaximumSet_Gen.json
func ExampleBusinessCaseAvsSummaryOperationsClient_NewListByBusinessCasePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBusinessCaseAvsSummaryOperationsClient().NewListByBusinessCasePager("rgopenapi", "multipleto8617project", "sample-business-case", nil)
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
		// page.AvsSummaryListResult = armmigrationassessment.AvsSummaryListResult{
		// 	Value: []*armmigrationassessment.AvsSummary{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentProjects/businessCases/avsSummaries"),
		// 			ID: to.Ptr("/subscriptions/ADC896AD-6A38-454E-9A62-AFC618F5F4BC/resourceGroups/rgopenapi/providers/Microsoft.Migrate/assessmentProjects/multipleto8617project/businessCases/sample-business-case/avsSummaries/default"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-08T07:09:55.036Z"); return t}()),
		// 				CreatedBy: to.Ptr("t72jdt@company.com"),
		// 				CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-08T07:09:55.036Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("t72jdt@company.com"),
		// 				LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmigrationassessment.AvsSummaryProperties{
		// 				AzureAvsSummary: &armmigrationassessment.AzureAvsSummary{
		// 					AvsCostDetails: &armmigrationassessment.CostDetails{
		// 						AhubSavings: to.Ptr[float32](1),
		// 						ComputeCost: to.Ptr[float32](19),
		// 						EsuSavings: to.Ptr[float32](28),
		// 						FacilitiesCost: to.Ptr[float32](19),
		// 						ItLaborCost: to.Ptr[float32](29),
		// 						LinuxAhubSavings: to.Ptr[float32](18),
		// 						ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
		// 							ManagementCost: to.Ptr[float32](22),
		// 							ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
		// 								{
		// 									Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
		// 									Value: to.Ptr[float32](16),
		// 							}},
		// 						},
		// 						NetworkCost: to.Ptr[float32](28),
		// 						SecurityCost: to.Ptr[float32](4),
		// 						StorageCost: to.Ptr[float32](27),
		// 					},
		// 					AvsNodeSummary: &armmigrationassessment.AvsNodeSummary{
		// 						AvsNodeCostDetails: &armmigrationassessment.CostDetails{
		// 							AhubSavings: to.Ptr[float32](1),
		// 							ComputeCost: to.Ptr[float32](19),
		// 							EsuSavings: to.Ptr[float32](28),
		// 							FacilitiesCost: to.Ptr[float32](19),
		// 							ItLaborCost: to.Ptr[float32](29),
		// 							LinuxAhubSavings: to.Ptr[float32](18),
		// 							ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
		// 								ManagementCost: to.Ptr[float32](22),
		// 								ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
		// 									{
		// 										Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
		// 										Value: to.Ptr[float32](16),
		// 								}},
		// 							},
		// 							NetworkCost: to.Ptr[float32](28),
		// 							SecurityCost: to.Ptr[float32](4),
		// 							StorageCost: to.Ptr[float32](27),
		// 						},
		// 						Cores: to.Ptr[int32](17),
		// 						EstimatedCost: to.Ptr[float32](29),
		// 						EstimatedCostByRecommendedOffer: []*armmigrationassessment.EstimatedCostByOffer{
		// 							{
		// 								Cost: to.Ptr[float32](17),
		// 								OfferName: to.Ptr("3 Year RI"),
		// 						}},
		// 						EstimatedCostWithVcfByol: to.Ptr[float32](24),
		// 						EstimatedExternalStorage: []*armmigrationassessment.EstimatedExternalStorage{
		// 							{
		// 								StorageType: to.Ptr(armmigrationassessment.ExternalStorageTypeAnfStandard),
		// 								StorageUtilization: to.Ptr[float32](97),
		// 								TotalStorageCost: to.Ptr[float32](250),
		// 								TotalStorageInGB: to.Ptr[float32](1024),
		// 							},
		// 							{
		// 								StorageType: to.Ptr(armmigrationassessment.ExternalStorageTypeAnfPremium),
		// 								StorageUtilization: to.Ptr[float32](99),
		// 								TotalStorageCost: to.Ptr[float32](350),
		// 								TotalStorageInGB: to.Ptr[float32](1024),
		// 						}},
		// 						EstimatedNetwork: []*armmigrationassessment.EstimatedNetwork{
		// 							{
		// 								Cost: to.Ptr[float32](200),
		// 								SKUType: to.Ptr(armmigrationassessment.NetworkSKUTypeExpressRouteUltra),
		// 						}},
		// 						MemoryGb: to.Ptr[float32](19),
		// 						OSLicensingCost: to.Ptr[float32](18),
		// 						RecommendedAvsNodeTypeSummary: []*armmigrationassessment.RecommendedAvsNodeTypeSummary{
		// 							{
		// 								AvsNodeType: to.Ptr("AV36"),
		// 								FailuresToTolerateAndRaidLevel: to.Ptr(armmigrationassessment.FttAndRaidLevelFtt1Raid1),
		// 								NumberOfNodes: to.Ptr[int32](16),
		// 						}},
		// 						ServersSuitableForMigration: to.Ptr[int32](28),
		// 						TotalNumberOfServers: to.Ptr[int32](4),
		// 					},
		// 					YearOnYearEstimates: &armmigrationassessment.YearOnYearEstimates{
		// 						AzureCapexCost: map[string]*float32{
		// 							"Year0": to.Ptr[float32](1049.81370090828),
		// 							"Year1": to.Ptr[float32](881.843508762955),
		// 							"Year2": to.Ptr[float32](578.709802625689),
		// 							"Year3": to.Ptr[float32](0),
		// 						},
		// 						AzureCost: map[string]*float32{
		// 							"Year0": to.Ptr[float32](351803.869241272),
		// 							"Year1": to.Ptr[float32](374738.191162668),
		// 							"Year2": to.Ptr[float32](391989.235419251),
		// 							"Year3": to.Ptr[float32](396114.705),
		// 						},
		// 						AzureOpexCost: map[string]*float32{
		// 							"Year0": to.Ptr[float32](350754.055540363),
		// 							"Year1": to.Ptr[float32](373856.347653905),
		// 							"Year2": to.Ptr[float32](391410.525616625),
		// 							"Year3": to.Ptr[float32](396114.705),
		// 						},
		// 						OnPremisesCapexCost: map[string]*float32{
		// 							"Year0": to.Ptr[float32](1049.81370090828),
		// 							"Year1": to.Ptr[float32](1102.30438595369),
		// 							"Year2": to.Ptr[float32](1157.41960525138),
		// 							"Year3": to.Ptr[float32](1215.29058551395),
		// 						},
		// 						OnPremisesCost: map[string]*float32{
		// 							"Year0": to.Ptr[float32](351803.869241272),
		// 							"Year1": to.Ptr[float32](369394.062703335),
		// 							"Year2": to.Ptr[float32](387863.765838502),
		// 							"Year3": to.Ptr[float32](408410.454130427),
		// 						},
		// 						OnPremisesOpexCost: map[string]*float32{
		// 							"Year0": to.Ptr[float32](350754.055540363),
		// 							"Year1": to.Ptr[float32](368291.758317382),
		// 							"Year2": to.Ptr[float32](386706.346233251),
		// 							"Year3": to.Ptr[float32](407195.163544913),
		// 						},
		// 						PaybackPeriod: to.Ptr[int32](0),
		// 						Savings: map[string]*float32{
		// 							"Year0": to.Ptr[float32](0),
		// 							"Year1": to.Ptr[float32](-5344.12845933292),
		// 							"Year2": to.Ptr[float32](-4125.46958074899),
		// 							"Year3": to.Ptr[float32](12295.7491304271),
		// 						},
		// 					},
		// 				},
		// 				OnPremisesAvsSummary: &armmigrationassessment.OnPremisesIaasSummary{
		// 					Cores: to.Ptr[int32](26),
		// 					CPUUtilization: to.Ptr[float32](24),
		// 					DistributionByOperatingSystem: []*armmigrationassessment.DistributionByOperatingSystem{
		// 						{
		// 							Count: to.Ptr[int32](4),
		// 							OSClassificationType: to.Ptr("WindowsOSInSupport"),
		// 					}},
		// 					DistributionByOperatingSystemVersion: []*armmigrationassessment.DistributionByType{
		// 						{
		// 							Type: to.Ptr("Microsoft Windows Server 2016 Datacenter"),
		// 							Count: to.Ptr[int32](27),
		// 					}},
		// 					DistributionBySQLEdition: []*armmigrationassessment.DistributionByType{
		// 						{
		// 							Type: to.Ptr("Enterprise Evaluation Edition"),
		// 							Count: to.Ptr[int32](27),
		// 					}},
		// 					DistributionBySQLVersion: []*armmigrationassessment.DistributionByType{
		// 						{
		// 							Type: to.Ptr("SQL Server 2012"),
		// 							Count: to.Ptr[int32](27),
		// 					}},
		// 					DistributionByVirtualization: []*armmigrationassessment.DistributionByVirtualization{
		// 						{
		// 							Count: to.Ptr[int32](30),
		// 							VirtualizationType: to.Ptr("Vmware"),
		// 					}},
		// 					IaasOsSupportStatusDistribution: &armmigrationassessment.SupportStatusDistribution{
		// 						Extended: to.Ptr[int32](26),
		// 						MainStream: to.Ptr[int32](20),
		// 						OutOfSupport: to.Ptr[int32](29),
		// 						UnknownSupport: to.Ptr[int32](2),
		// 					},
		// 					IaasSQLSupportStatusDistribution: &armmigrationassessment.SupportStatusDistribution{
		// 						Extended: to.Ptr[int32](26),
		// 						MainStream: to.Ptr[int32](20),
		// 						OutOfSupport: to.Ptr[int32](29),
		// 						UnknownSupport: to.Ptr[int32](2),
		// 					},
		// 					MemoryGb: to.Ptr[float32](4),
		// 					MemoryUtilization: to.Ptr[float32](10),
		// 					OnPremisesIaasCostDetails: &armmigrationassessment.CostDetails{
		// 						AhubSavings: to.Ptr[float32](1),
		// 						ComputeCost: to.Ptr[float32](19),
		// 						EsuSavings: to.Ptr[float32](28),
		// 						FacilitiesCost: to.Ptr[float32](19),
		// 						ItLaborCost: to.Ptr[float32](29),
		// 						LinuxAhubSavings: to.Ptr[float32](18),
		// 						ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
		// 							ManagementCost: to.Ptr[float32](22),
		// 							ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
		// 								{
		// 									Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
		// 									Value: to.Ptr[float32](16),
		// 							}},
		// 						},
		// 						NetworkCost: to.Ptr[float32](28),
		// 						SecurityCost: to.Ptr[float32](4),
		// 						StorageCost: to.Ptr[float32](27),
		// 					},
		// 					OnPremisesIaasCostForDecommissioned: &armmigrationassessment.CostDetails{
		// 						AhubSavings: to.Ptr[float32](1),
		// 						ComputeCost: to.Ptr[float32](19),
		// 						EsuSavings: to.Ptr[float32](28),
		// 						FacilitiesCost: to.Ptr[float32](19),
		// 						ItLaborCost: to.Ptr[float32](29),
		// 						LinuxAhubSavings: to.Ptr[float32](18),
		// 						ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
		// 							ManagementCost: to.Ptr[float32](22),
		// 							ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
		// 								{
		// 									Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
		// 									Value: to.Ptr[float32](16),
		// 							}},
		// 						},
		// 						NetworkCost: to.Ptr[float32](28),
		// 						SecurityCost: to.Ptr[float32](4),
		// 						StorageCost: to.Ptr[float32](27),
		// 					},
		// 					OnPremisesUtilizationData: &armmigrationassessment.UtilizationData{
		// 						NumberOfActiveEntities: to.Ptr[int32](10),
		// 						NumberOfDecommisionEntities: to.Ptr[int32](10),
		// 						NumberOfInactiveEntities: to.Ptr[int32](18),
		// 						NumberOfUnknownEntities: to.Ptr[int32](13),
		// 					},
		// 					OSLicensingDetails: []*armmigrationassessment.IaasOsLicensingDetails{
		// 						{
		// 							DecomissionCost: to.Ptr[float32](29),
		// 							OSType: to.Ptr("Windows"),
		// 							TotalCost: to.Ptr[float32](11),
		// 					}},
		// 					OSServicePackInsight: &armmigrationassessment.ServicePackInsight{
		// 						Patched: to.Ptr[int32](14),
		// 						UnknownServicePack: to.Ptr[int32](3),
		// 						Unpatched: to.Ptr[int32](5),
		// 					},
		// 					Servers: to.Ptr[int32](29),
		// 					SQLServicePackInsight: &armmigrationassessment.ServicePackInsight{
		// 						Patched: to.Ptr[int32](14),
		// 						UnknownServicePack: to.Ptr[int32](3),
		// 						Unpatched: to.Ptr[int32](5),
		// 					},
		// 					StorageUtilization: to.Ptr[float32](26),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
