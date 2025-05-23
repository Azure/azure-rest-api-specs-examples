package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/BusinessCaseIaasSummaryOperations_Get_MaximumSet_Gen.json
func ExampleBusinessCaseIaasSummaryOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBusinessCaseIaasSummaryOperationsClient().Get(ctx, "rgopenapi", "multipleto8617project", "sample-business-case", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IaasSummary = armmigrationassessment.IaasSummary{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentProjects/businessCases/iaasSummaries"),
	// 	ID: to.Ptr("/subscriptions/ADC896AD-6A38-454E-9A62-AFC618F5F4BC/resourceGroups/rgopenapi/providers/Microsoft.Migrate/assessmentProjects/multipleto8617project/businessCases/sample-business-case/iaasSummaries/default"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-08T07:09:55.036Z"); return t}()),
	// 		CreatedBy: to.Ptr("t72jdt@company.com"),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-08T07:09:55.036Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("t72jdt@company.com"),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmigrationassessment.IaasSummaryProperties{
	// 		AzureIaasSummary: &armmigrationassessment.AzureIaasSummary{
	// 			AzureIaasCostDetails: &armmigrationassessment.CostDetails{
	// 				AhubSavings: to.Ptr[float32](1),
	// 				ComputeCost: to.Ptr[float32](19),
	// 				EsuSavings: to.Ptr[float32](28),
	// 				FacilitiesCost: to.Ptr[float32](19),
	// 				ItLaborCost: to.Ptr[float32](29),
	// 				LinuxAhubSavings: to.Ptr[float32](18),
	// 				ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 					ManagementCost: to.Ptr[float32](22),
	// 					ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 						{
	// 							Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 							Value: to.Ptr[float32](16),
	// 					}},
	// 				},
	// 				NetworkCost: to.Ptr[float32](28),
	// 				SecurityCost: to.Ptr[float32](4),
	// 				StorageCost: to.Ptr[float32](27),
	// 			},
	// 			AzureIaasSQLSummary: &armmigrationassessment.AzureIaasSQLSummary{
	// 				AzureSQLCostDetails: &armmigrationassessment.CostDetails{
	// 					AhubSavings: to.Ptr[float32](1),
	// 					ComputeCost: to.Ptr[float32](19),
	// 					EsuSavings: to.Ptr[float32](28),
	// 					FacilitiesCost: to.Ptr[float32](19),
	// 					ItLaborCost: to.Ptr[float32](29),
	// 					LinuxAhubSavings: to.Ptr[float32](18),
	// 					ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 						ManagementCost: to.Ptr[float32](22),
	// 						ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 							{
	// 								Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 								Value: to.Ptr[float32](16),
	// 						}},
	// 					},
	// 					NetworkCost: to.Ptr[float32](28),
	// 					SecurityCost: to.Ptr[float32](4),
	// 					StorageCost: to.Ptr[float32](27),
	// 				},
	// 				CostByStorageType: []*armmigrationassessment.CostByStorageType{
	// 					{
	// 						Cost: to.Ptr[float32](12),
	// 						StorageType: to.Ptr("StandardSsd"),
	// 				}},
	// 				EstimatedCost: to.Ptr[float32](23),
	// 				EstimatedCostByRecommendedOffer: []*armmigrationassessment.EstimatedCostByOffer{
	// 					{
	// 						Cost: to.Ptr[float32](17),
	// 						OfferName: to.Ptr("3 Year RI"),
	// 				}},
	// 				OSLicensingCost: to.Ptr[float32](8),
	// 				RecommendedVMFamilySummary: []*armmigrationassessment.RecommendedVMFamilySummary{
	// 					{
	// 						AzureVMFamily: to.Ptr("Standard_F4"),
	// 						NumberOfMachines: to.Ptr[int32](26),
	// 				}},
	// 				SuitableSQLEntities: to.Ptr[int32](21),
	// 				TotalSQLEntities: to.Ptr[int32](7),
	// 			},
	// 			AzureIaasVMSummary: &armmigrationassessment.AzureIaasVMSummary{
	// 				AzureVMCostDetails: &armmigrationassessment.CostDetails{
	// 					AhubSavings: to.Ptr[float32](1),
	// 					ComputeCost: to.Ptr[float32](19),
	// 					EsuSavings: to.Ptr[float32](28),
	// 					FacilitiesCost: to.Ptr[float32](19),
	// 					ItLaborCost: to.Ptr[float32](29),
	// 					LinuxAhubSavings: to.Ptr[float32](18),
	// 					ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 						ManagementCost: to.Ptr[float32](22),
	// 						ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 							{
	// 								Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 								Value: to.Ptr[float32](16),
	// 						}},
	// 					},
	// 					NetworkCost: to.Ptr[float32](28),
	// 					SecurityCost: to.Ptr[float32](4),
	// 					StorageCost: to.Ptr[float32](27),
	// 				},
	// 				Cores: to.Ptr[int32](6),
	// 				CostByStorageType: []*armmigrationassessment.CostByStorageType{
	// 					{
	// 						Cost: to.Ptr[float32](12),
	// 						StorageType: to.Ptr("StandardSsd"),
	// 				}},
	// 				EstimatedCost: to.Ptr[float32](19),
	// 				EstimatedCostByRecommendedOffer: []*armmigrationassessment.EstimatedCostByOffer{
	// 					{
	// 						Cost: to.Ptr[float32](17),
	// 						OfferName: to.Ptr("3 Year RI"),
	// 				}},
	// 				MemoryGb: to.Ptr[float32](26),
	// 				OSLicensingCost: to.Ptr[float32](12),
	// 				RecommendedVMFamilySummary: []*armmigrationassessment.RecommendedVMFamilySummary{
	// 					{
	// 						AzureVMFamily: to.Ptr("Standard_F4"),
	// 						NumberOfMachines: to.Ptr[int32](26),
	// 				}},
	// 				ServersSuitableForMigration: to.Ptr[int32](15),
	// 				TotalNumberOfServers: to.Ptr[int32](27),
	// 			},
	// 			YearOnYearEstimates: &armmigrationassessment.YearOnYearEstimates{
	// 				AzureCapexCost: map[string]*float32{
	// 					"Year0": to.Ptr[float32](1049.81370090828),
	// 					"Year1": to.Ptr[float32](881.843508762955),
	// 					"Year2": to.Ptr[float32](578.709802625689),
	// 					"Year3": to.Ptr[float32](0),
	// 				},
	// 				AzureCost: map[string]*float32{
	// 					"Year0": to.Ptr[float32](351803.869241272),
	// 					"Year1": to.Ptr[float32](374738.191162668),
	// 					"Year2": to.Ptr[float32](391989.235419251),
	// 					"Year3": to.Ptr[float32](396114.705),
	// 				},
	// 				AzureOpexCost: map[string]*float32{
	// 					"Year0": to.Ptr[float32](350754.055540363),
	// 					"Year1": to.Ptr[float32](373856.347653905),
	// 					"Year2": to.Ptr[float32](391410.525616625),
	// 					"Year3": to.Ptr[float32](396114.705),
	// 				},
	// 				OnPremisesCapexCost: map[string]*float32{
	// 					"Year0": to.Ptr[float32](1049.81370090828),
	// 					"Year1": to.Ptr[float32](1102.30438595369),
	// 					"Year2": to.Ptr[float32](1157.41960525138),
	// 					"Year3": to.Ptr[float32](1215.29058551395),
	// 				},
	// 				OnPremisesCost: map[string]*float32{
	// 					"Year0": to.Ptr[float32](351803.869241272),
	// 					"Year1": to.Ptr[float32](369394.062703335),
	// 					"Year2": to.Ptr[float32](387863.765838502),
	// 					"Year3": to.Ptr[float32](408410.454130427),
	// 				},
	// 				OnPremisesOpexCost: map[string]*float32{
	// 					"Year0": to.Ptr[float32](350754.055540363),
	// 					"Year1": to.Ptr[float32](368291.758317382),
	// 					"Year2": to.Ptr[float32](386706.346233251),
	// 					"Year3": to.Ptr[float32](407195.163544913),
	// 				},
	// 				PaybackPeriod: to.Ptr[int32](0),
	// 				Savings: map[string]*float32{
	// 					"Year0": to.Ptr[float32](0),
	// 					"Year1": to.Ptr[float32](-5344.12845933292),
	// 					"Year2": to.Ptr[float32](-4125.46958074899),
	// 					"Year3": to.Ptr[float32](12295.7491304271),
	// 				},
	// 			},
	// 		},
	// 		OnPremisesIaasSummary: &armmigrationassessment.OnPremisesIaasSummary{
	// 			Cores: to.Ptr[int32](26),
	// 			CPUUtilization: to.Ptr[float32](24),
	// 			DistributionByOperatingSystem: []*armmigrationassessment.DistributionByOperatingSystem{
	// 				{
	// 					Count: to.Ptr[int32](4),
	// 					OSClassificationType: to.Ptr("LinuxOSInSupport"),
	// 			}},
	// 			DistributionByOperatingSystemVersion: []*armmigrationassessment.DistributionByType{
	// 				{
	// 					Type: to.Ptr("Ubuntu"),
	// 					Count: to.Ptr[int32](27),
	// 			}},
	// 			DistributionBySQLEdition: []*armmigrationassessment.DistributionByType{
	// 				{
	// 					Type: to.Ptr("Enterprise Evaluation Edition"),
	// 					Count: to.Ptr[int32](27),
	// 			}},
	// 			DistributionBySQLVersion: []*armmigrationassessment.DistributionByType{
	// 				{
	// 					Type: to.Ptr("SQL Server 2012"),
	// 					Count: to.Ptr[int32](27),
	// 			}},
	// 			DistributionByVirtualization: []*armmigrationassessment.DistributionByVirtualization{
	// 				{
	// 					Count: to.Ptr[int32](30),
	// 					VirtualizationType: to.Ptr("Vmware"),
	// 			}},
	// 			IaasOsSupportStatusDistribution: &armmigrationassessment.SupportStatusDistribution{
	// 				Extended: to.Ptr[int32](26),
	// 				MainStream: to.Ptr[int32](20),
	// 				OutOfSupport: to.Ptr[int32](29),
	// 				UnknownSupport: to.Ptr[int32](2),
	// 			},
	// 			IaasSQLSupportStatusDistribution: &armmigrationassessment.SupportStatusDistribution{
	// 				Extended: to.Ptr[int32](26),
	// 				MainStream: to.Ptr[int32](20),
	// 				OutOfSupport: to.Ptr[int32](29),
	// 				UnknownSupport: to.Ptr[int32](2),
	// 			},
	// 			MemoryGb: to.Ptr[float32](4),
	// 			MemoryUtilization: to.Ptr[float32](10),
	// 			OnPremisesIaasCostDetails: &armmigrationassessment.CostDetails{
	// 				AhubSavings: to.Ptr[float32](1),
	// 				ComputeCost: to.Ptr[float32](19),
	// 				EsuSavings: to.Ptr[float32](28),
	// 				FacilitiesCost: to.Ptr[float32](19),
	// 				ItLaborCost: to.Ptr[float32](29),
	// 				LinuxAhubSavings: to.Ptr[float32](18),
	// 				ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 					ManagementCost: to.Ptr[float32](22),
	// 					ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 						{
	// 							Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 							Value: to.Ptr[float32](16),
	// 					}},
	// 				},
	// 				NetworkCost: to.Ptr[float32](28),
	// 				SecurityCost: to.Ptr[float32](4),
	// 				StorageCost: to.Ptr[float32](27),
	// 			},
	// 			OnPremisesIaasCostForDecommissioned: &armmigrationassessment.CostDetails{
	// 				AhubSavings: to.Ptr[float32](1),
	// 				ComputeCost: to.Ptr[float32](19),
	// 				EsuSavings: to.Ptr[float32](28),
	// 				FacilitiesCost: to.Ptr[float32](19),
	// 				ItLaborCost: to.Ptr[float32](29),
	// 				LinuxAhubSavings: to.Ptr[float32](18),
	// 				ManagementCostDetails: &armmigrationassessment.ManagementCostDetails{
	// 					ManagementCost: to.Ptr[float32](22),
	// 					ManagementCostComponents: []*armmigrationassessment.ManagementCostComponent{
	// 						{
	// 							Name: to.Ptr(armmigrationassessment.ManagementCostComponentNameUnknown),
	// 							Value: to.Ptr[float32](16),
	// 					}},
	// 				},
	// 				NetworkCost: to.Ptr[float32](28),
	// 				SecurityCost: to.Ptr[float32](4),
	// 				StorageCost: to.Ptr[float32](27),
	// 			},
	// 			OnPremisesUtilizationData: &armmigrationassessment.UtilizationData{
	// 				NumberOfActiveEntities: to.Ptr[int32](10),
	// 				NumberOfDecommisionEntities: to.Ptr[int32](10),
	// 				NumberOfInactiveEntities: to.Ptr[int32](18),
	// 				NumberOfUnknownEntities: to.Ptr[int32](13),
	// 			},
	// 			OSLicensingDetails: []*armmigrationassessment.IaasOsLicensingDetails{
	// 				{
	// 					DecomissionCost: to.Ptr[float32](29),
	// 					OSType: to.Ptr("Windows"),
	// 					TotalCost: to.Ptr[float32](11),
	// 			}},
	// 			OSServicePackInsight: &armmigrationassessment.ServicePackInsight{
	// 				Patched: to.Ptr[int32](14),
	// 				UnknownServicePack: to.Ptr[int32](3),
	// 				Unpatched: to.Ptr[int32](5),
	// 			},
	// 			Servers: to.Ptr[int32](29),
	// 			SQLServicePackInsight: &armmigrationassessment.ServicePackInsight{
	// 				Patched: to.Ptr[int32](14),
	// 				UnknownServicePack: to.Ptr[int32](3),
	// 				Unpatched: to.Ptr[int32](5),
	// 			},
	// 			StorageUtilization: to.Ptr[float32](26),
	// 		},
	// 	},
	// }
}
