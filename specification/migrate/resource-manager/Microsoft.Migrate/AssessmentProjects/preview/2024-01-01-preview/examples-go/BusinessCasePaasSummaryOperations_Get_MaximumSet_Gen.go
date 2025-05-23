package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/BusinessCasePaasSummaryOperations_Get_MaximumSet_Gen.json
func ExampleBusinessCasePaasSummaryOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBusinessCasePaasSummaryOperationsClient().Get(ctx, "rgopenapi", "multipleto8617project", "sample-business-case", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PaasSummary = armmigrationassessment.PaasSummary{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentProjects/businessCases/paasSummaries"),
	// 	ID: to.Ptr("/subscriptions/ADC896AD-6A38-454E-9A62-AFC618F5F4BC/resourceGroups/rgopenapi/providers/Microsoft.Migrate/assessmentProjects/multipleto8617project/businessCases/sample-business-case/paasSummaries/default"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-08T07:09:55.036Z"); return t}()),
	// 		CreatedBy: to.Ptr("t72jdt@company.com"),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-08T07:09:55.036Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("t72jdt@company.com"),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmigrationassessment.PaasSummaryProperties{
	// 		Azure: &armmigrationassessment.AzurePaasSummary{
	// 			AzureAppServiceContainerSummary: &armmigrationassessment.AzureAppServiceContainerSummary{
	// 				DistributionByApp: []*armmigrationassessment.WebAppDistribution{
	// 					{
	// 						NumberOfWebApps: to.Ptr[float32](14),
	// 						WebAppType: to.Ptr(armmigrationassessment.WebAppTypesUnknown),
	// 				}},
	// 				DistributionBySKU: []*armmigrationassessment.WebAppTargetSKU{
	// 					{
	// 						Cost: to.Ptr[float32](14),
	// 						Count: to.Ptr[int32](30),
	// 						SKUName: to.Ptr("Premium_V3"),
	// 				}},
	// 				EstimatedCostByOffer: []*armmigrationassessment.EstimatedCostByOffer{
	// 					{
	// 						Cost: to.Ptr[float32](17),
	// 						OfferName: to.Ptr("3 Year ASP"),
	// 				}},
	// 			},
	// 			AzureAppServiceSummary: &armmigrationassessment.AzureAppServiceSummary{
	// 				DistributionByApp: []*armmigrationassessment.WebAppDistribution{
	// 					{
	// 						NumberOfWebApps: to.Ptr[float32](14),
	// 						WebAppType: to.Ptr(armmigrationassessment.WebAppTypesUnknown),
	// 				}},
	// 				DistributionBySKU: []*armmigrationassessment.AppServiceSKU{
	// 					{
	// 						Cost: to.Ptr[float32](6),
	// 						Count: to.Ptr[int32](27),
	// 						SKUName: to.Ptr("Premium_V3"),
	// 				}},
	// 				EstimatedCostByOffer: []*armmigrationassessment.EstimatedCostByOffer{
	// 					{
	// 						Cost: to.Ptr[float32](17),
	// 						OfferName: to.Ptr("3 Year ASP"),
	// 				}},
	// 			},
	// 			AzureKubernetesServiceSummary: &armmigrationassessment.AzureKubernetesServiceSummary{
	// 				DistributionByApp: []*armmigrationassessment.WebAppDistribution{
	// 					{
	// 						NumberOfWebApps: to.Ptr[float32](14),
	// 						WebAppType: to.Ptr(armmigrationassessment.WebAppTypesUnknown),
	// 				}},
	// 				DistributionBySKU: []*armmigrationassessment.WebAppTargetSKU{
	// 					{
	// 						Cost: to.Ptr[float32](14),
	// 						Count: to.Ptr[int32](30),
	// 						SKUName: to.Ptr("Premium_V3"),
	// 				}},
	// 				EstimatedCostByOffer: []*armmigrationassessment.EstimatedCostByOffer{
	// 					{
	// 						Cost: to.Ptr[float32](17),
	// 						OfferName: to.Ptr("3 Year ASP"),
	// 				}},
	// 			},
	// 			AzureSQLSummary: &armmigrationassessment.AzureSQLSummary{
	// 				DistributionByServiceTier: []*armmigrationassessment.SQLServiceTier{
	// 					{
	// 						ServiceTierName: to.Ptr("GeneralPurpose"),
	// 						ServiceTierNumber: to.Ptr[float32](1),
	// 				}},
	// 				DistributionByServiceTierForSQLDb: []*armmigrationassessment.SQLServiceTier{
	// 					{
	// 						ServiceTierName: to.Ptr("GeneralPurpose"),
	// 						ServiceTierNumber: to.Ptr[float32](1),
	// 				}},
	// 				DistributionByServiceTierForSQLMi: []*armmigrationassessment.SQLServiceTier{
	// 					{
	// 						ServiceTierName: to.Ptr("GeneralPurpose"),
	// 						ServiceTierNumber: to.Ptr[float32](1),
	// 				}},
	// 				EstimatedCostByOffer: []*armmigrationassessment.EstimatedCostByOffer{
	// 					{
	// 						Cost: to.Ptr[float32](17),
	// 						OfferName: to.Ptr("3 Year RI"),
	// 				}},
	// 			},
	// 			Cores: to.Ptr[int32](3),
	// 			EstimatedCostByTarget: &armmigrationassessment.EstimatedCostByTarget{
	// 				AppServiceContainerCost: to.Ptr[float32](23),
	// 				AppServiceCost: to.Ptr[float32](5),
	// 				AzureKubernetesServiceCost: to.Ptr[float32](22),
	// 				AzureSQLDbCost: to.Ptr[float32](6),
	// 				AzureSQLMiCost: to.Ptr[float32](27),
	// 			},
	// 			EstimatedSQLServerLicensingCost: to.Ptr[float32](18),
	// 			Memory: to.Ptr[float32](29),
	// 			Storage: to.Ptr[float32](8),
	// 			SuitableSQLEntities: to.Ptr[int32](8),
	// 			SuitableWebApps: to.Ptr[int32](8),
	// 			TotalAzurePaasCost: &armmigrationassessment.CostDetails{
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
	// 			TotalSQLEntities: to.Ptr[int32](1),
	// 			TotalWebApps: to.Ptr[int32](22),
	// 		},
	// 		OnPremises: &armmigrationassessment.OnPremisesPaasSummary{
	// 			CPUUtilization: to.Ptr[float32](18),
	// 			MemoryUtilization: to.Ptr[float32](14),
	// 			OnPremisesPaasCostForDecommisioned: &armmigrationassessment.CostDetails{
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
	// 			OnPremisesPaasLicensingCost: &armmigrationassessment.OnPremisesPaasLicensingCost{
	// 				DecomissionServerCost: to.Ptr[float32](1),
	// 				TotalCost: to.Ptr[float32](29),
	// 			},
	// 			OnPremisesSQLSummary: &armmigrationassessment.OnPremisesSQLSummary{
	// 				DistributionBySQLEdition: []*armmigrationassessment.DistributionByType{
	// 					{
	// 						Type: to.Ptr("Enterprise Edition (64-bit)"),
	// 						Count: to.Ptr[int32](27),
	// 				}},
	// 				DistributionBySQLVersion: []*armmigrationassessment.SQLVersionDetails{
	// 					{
	// 						NumberOfInstances: to.Ptr[int32](20),
	// 						SQLVersion: to.Ptr("SQL Server 2019"),
	// 				}},
	// 				SQLDatabases: to.Ptr[int32](21),
	// 				SQLInstances: to.Ptr[int32](26),
	// 				SQLOnPremisesUtilizationData: &armmigrationassessment.UtilizationData{
	// 					NumberOfActiveEntities: to.Ptr[int32](10),
	// 					NumberOfDecommisionEntities: to.Ptr[int32](10),
	// 					NumberOfInactiveEntities: to.Ptr[int32](18),
	// 					NumberOfUnknownEntities: to.Ptr[int32](13),
	// 				},
	// 			},
	// 			OnPremisesWebAppSummary: &armmigrationassessment.OnPremisesWebAppSummary{
	// 				NumberOfWebApplications: to.Ptr[int32](27),
	// 				NumberOfWebAppsPerType: map[string]*int32{
	// 					"key2586": to.Ptr[int32](28),
	// 				},
	// 				NumberOfWebServers: to.Ptr[int32](1),
	// 			},
	// 			OSServicePackInsight: &armmigrationassessment.ServicePackInsight{
	// 				Patched: to.Ptr[int32](14),
	// 				UnknownServicePack: to.Ptr[int32](3),
	// 				Unpatched: to.Ptr[int32](5),
	// 			},
	// 			OSSupportStatusDistribution: &armmigrationassessment.SupportStatusDistribution{
	// 				Extended: to.Ptr[int32](26),
	// 				MainStream: to.Ptr[int32](20),
	// 				OutOfSupport: to.Ptr[int32](29),
	// 				UnknownSupport: to.Ptr[int32](2),
	// 			},
	// 			SQLServicePackInsight: &armmigrationassessment.ServicePackInsight{
	// 				Patched: to.Ptr[int32](14),
	// 				UnknownServicePack: to.Ptr[int32](3),
	// 				Unpatched: to.Ptr[int32](5),
	// 			},
	// 			SQLSupportStatusDistribution: &armmigrationassessment.SupportStatusDistribution{
	// 				Extended: to.Ptr[int32](26),
	// 				MainStream: to.Ptr[int32](20),
	// 				OutOfSupport: to.Ptr[int32](29),
	// 				UnknownSupport: to.Ptr[int32](2),
	// 			},
	// 			TotalOnPremisesPaasCost: &armmigrationassessment.CostDetails{
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
	// 			TotalServers: to.Ptr[int32](29),
	// 		},
	// 	},
	// }
}
