package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/BusinessCaseOperations_Get_MaximumSet_Gen.json
func ExampleBusinessCaseOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBusinessCaseOperationsClient().Get(ctx, "rgopenapi", "multipleto8617project", "sample-business-case", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BusinessCase = armmigrationassessment.BusinessCase{
	// 	Name: to.Ptr("sample-business-case"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentProjects/businessCases"),
	// 	ID: to.Ptr("/subscriptions/5E8485F1-A440-404E-9F91-AFAC5A9993B2/resourceGroups/rgv2024-01-01-preview/providers/Microsoft.Migrate/assessmentProjects/multipleto8617project/businessCases/sample-business-case"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T11:57:41.238Z"); return t}()),
	// 		CreatedBy: to.Ptr("t72jdt@company.com"),
	// 		CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T11:57:41.238Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("t72jdt@company.com"),
	// 		LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
	// 	},
	// 	Properties: &armmigrationassessment.BusinessCaseProperties{
	// 		ProvisioningState: to.Ptr(armmigrationassessment.ProvisioningState2Succeeded),
	// 		ReportStatusDetails: []*armmigrationassessment.ReportDetails{
	// 			{
	// 				ReportStatus: to.Ptr(armmigrationassessment.ReportStatusCompleted),
	// 				ReportType: to.Ptr(armmigrationassessment.ReportTypeExcel),
	// 		}},
	// 		Settings: &armmigrationassessment.Settings{
	// 			AzureArcSettings: &armmigrationassessment.AzureArcSettings{
	// 				AzureArcState: to.Ptr(armmigrationassessment.AzureArcStateEnabled),
	// 				LaborCostPercentage: to.Ptr[float32](70),
	// 				ManagementSettings: &armmigrationassessment.AzureArcManagementSettings{
	// 					MonitoringSettings: &armmigrationassessment.AzureArcMonitoringSettings{
	// 						AlertRulesCount: to.Ptr[int32](10),
	// 						LogsVolumeInGB: to.Ptr[float32](0.5),
	// 					},
	// 				},
	// 			},
	// 			AzureSettings: &armmigrationassessment.AzureSettings{
	// 				AvsLaborCostPercentage: to.Ptr[float32](69),
	// 				BusinessCaseType: to.Ptr(armmigrationassessment.MigrationStrategyOptimizeForCost),
	// 				ComfortFactor: to.Ptr[float32](8),
	// 				Currency: to.Ptr(armmigrationassessment.BusinessCaseCurrencyUSD),
	// 				DiscountPercentage: to.Ptr[float32](52),
	// 				IaasLaborCostPercentage: to.Ptr[float32](73),
	// 				InfrastructureGrowthRate: to.Ptr[float32](44),
	// 				NetworkCostPercentage: to.Ptr[float32](71),
	// 				PaasLaborCostPercentage: to.Ptr[float32](50),
	// 				PerYearMigrationCompletionPercentage: map[string]*float32{
	// 					"Year0": to.Ptr[float32](20),
	// 					"Year1": to.Ptr[float32](30),
	// 					"Year2": to.Ptr[float32](60),
	// 					"Year3": to.Ptr[float32](90),
	// 				},
	// 				PerformanceDataEndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T11:57:41.237Z"); return t}()),
	// 				PerformanceDataStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-10-05T11:57:41.237Z"); return t}()),
	// 				PerformanceUtilizationPercentile: to.Ptr[float32](64),
	// 				SavingsOption: to.Ptr(armmigrationassessment.SavingsOptionRI3Year),
	// 				TargetLocation: to.Ptr("WestUs2"),
	// 				Wacc: to.Ptr[float32](5),
	// 				WorkloadDiscoverySource: to.Ptr(armmigrationassessment.DiscoverySourceAppliance),
	// 			},
	// 			OnPremiseSettings: &armmigrationassessment.OnPremiseSettings{
	// 				ComputeSettings: &armmigrationassessment.ComputeSettings{
	// 					HyperthreadCoreToMemoryRatio: to.Ptr[float32](5),
	// 					Price: to.Ptr[float32](14),
	// 					RhelLinuxServerLicensing: &armmigrationassessment.LinuxServerLicensingSettings{
	// 						LicenseCost: to.Ptr[float32](9),
	// 					},
	// 					SQLServerLicensing: []*armmigrationassessment.SQLServerLicensingSettings{
	// 						{
	// 							LicenseCost: to.Ptr[float32](14),
	// 							SoftwareAssuranceCost: to.Ptr[float32](13),
	// 							Version: to.Ptr(armmigrationassessment.SQLServerLicenseTypeEnterprise),
	// 					}},
	// 					SuseLinuxServerLicensing: &armmigrationassessment.LinuxServerLicensingSettings{
	// 						LicenseCost: to.Ptr[float32](9),
	// 					},
	// 					VirtualizationSoftwareSettings: &armmigrationassessment.VirtualizationSoftwareSettings{
	// 						VMwareCloudFoundationLicenseCost: to.Ptr[float32](7),
	// 					},
	// 					WindowsServerLicensing: &armmigrationassessment.WindowsServerLicensingSettings{
	// 						LicenseCost: to.Ptr[float32](13),
	// 						LicensesPerCore: to.Ptr[int32](8),
	// 						SoftwareAssuranceCost: to.Ptr[float32](14),
	// 					},
	// 				},
	// 				FacilitySettings: &armmigrationassessment.FacilitySettings{
	// 					FacilitiesCostPerKwh: to.Ptr[float32](28),
	// 				},
	// 				LaborSettings: &armmigrationassessment.LaborSettings{
	// 					HourlyAdminCost: to.Ptr[float32](15),
	// 					PhysicalServersPerAdmin: to.Ptr[int32](30),
	// 					VirtualMachinesPerAdmin: to.Ptr[int32](1),
	// 				},
	// 				ManagementSettings: &armmigrationassessment.ManagementSettings{
	// 					HypervVirtualizationManagementSettings: &armmigrationassessment.HypervVirtualizationManagementSettings{
	// 						LicenseAndSupportList: []*armmigrationassessment.HypervLicense{
	// 							{
	// 								LicenseCost: to.Ptr[float32](12),
	// 								LicenseType: to.Ptr(armmigrationassessment.HyperVLicenseTypeStandard),
	// 						}},
	// 						NumberOfPhysicalCoresPerLicense: to.Ptr[int32](21),
	// 						SoftwareAssuranceCost: to.Ptr[float32](1),
	// 					},
	// 					OtherManagementCostsSettings: &armmigrationassessment.OtherManagementCostsSettings{
	// 						DataProtectionCostPerServerPerYear: to.Ptr[float32](28),
	// 						MonitoringCostPerServerPerYear: to.Ptr[float32](15),
	// 						PatchingCostPerServerPerYear: to.Ptr[float32](15),
	// 					},
	// 					ThirdPartyManagementSettings: &armmigrationassessment.ThirdPartyManagementSettings{
	// 						LicenseCost: to.Ptr[float32](6),
	// 						SupportCost: to.Ptr[float32](2),
	// 					},
	// 				},
	// 				NetworkSettings: &armmigrationassessment.NetworkSettings{
	// 					HardwareSoftwareCostPercentage: to.Ptr[float32](79),
	// 					MaintenanceCostPercentage: to.Ptr[float32](82),
	// 				},
	// 				SecuritySettings: &armmigrationassessment.SecuritySettings{
	// 					ServerSecurityCostPerServerPerYear: to.Ptr[float32](2),
	// 					SQLServerSecurityCostPerServerPerYear: to.Ptr[float32](23),
	// 				},
	// 				StorageSettings: &armmigrationassessment.StorageSettings{
	// 					CostPerGbPerMonth: to.Ptr[float32](15),
	// 					MaintainanceCostPercentageToAcquisitionCost: to.Ptr[float32](91),
	// 				},
	// 			},
	// 		},
	// 		State: to.Ptr(armmigrationassessment.BusinessCaseStateCompleted),
	// 	},
	// }
}
