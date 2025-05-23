package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AssessedSqlInstanceV2Operations_ListBySqlAssessmentV2_MaximumSet_Gen.json
func ExampleAssessedSQLInstanceV2OperationsClient_NewListBySQLAssessmentV2Pager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssessedSQLInstanceV2OperationsClient().NewListBySQLAssessmentV2Pager("rgmigrate", "fci-test6904project", "test_fci_hadr", "test_swagger_1", &armmigrationassessment.AssessedSQLInstanceV2OperationsClientListBySQLAssessmentV2Options{Filter: to.Ptr("(contains(Properties/InstanceName,'MSSQLSERVER'))"),
		PageSize:          to.Ptr[int32](23),
		ContinuationToken: nil,
		TotalRecordCount:  to.Ptr[int32](1),
	})
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
		// page.AssessedSQLInstanceV2ListResult = armmigrationassessment.AssessedSQLInstanceV2ListResult{
		// 	Value: []*armmigrationassessment.AssessedSQLInstanceV2{
		// 		{
		// 			Name: to.Ptr("27b94796-026e-4278-a546-664d46977c0a"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/sqlAssessments/assessedSqlInstances"),
		// 			ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/dhlodhiccy/providers/Microsoft.Migrate/assessmentprojects/dhccy30279project/groups/test_paas_pref_grp_2812/sqlAssessments/test_train_250823/assessedSqlInstances/27b94796-026e-4278-a546-664d46977c0a"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 			},
		// 			Properties: &armmigrationassessment.AssessedSQLInstanceV2Properties{
		// 				AzureSQLDBSuitabilityDetails: &armmigrationassessment.SQLAssessmentV2PaasSuitabilityData{
		// 					CostComponents: []*armmigrationassessment.CostComponent{
		// 						{
		// 							Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlySecurityCost),
		// 							Value: to.Ptr[float32](15),
		// 					}},
		// 					MigrationGuidelines: []*armmigrationassessment.SQLMigrationGuideline{
		// 						{
		// 							GuidelineID: to.Ptr("MigrateDataFromSourceSqlServerToDB"),
		// 							MigrationGuidelineCategory: to.Ptr(armmigrationassessment.SQLMigrationGuidelineCategoryFailoverCluterInstanceGuideLine),
		// 							MigrationGuidelineContext: []*armmigrationassessment.MigrationGuidelineContext{
		// 							},
		// 					}},
		// 					MigrationIssues: []*armmigrationassessment.SQLAssessmentMigrationIssue{
		// 						{
		// 							ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
		// 								{
		// 									ObjectName: to.Ptr("accessibilitytestdiag917"),
		// 									ObjectType: to.Ptr("ServerCredential"),
		// 							}},
		// 							IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryWarning),
		// 							IssueID: to.Ptr("ServerCredentials"),
		// 						},
		// 						{
		// 							ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
		// 							},
		// 							IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryInternal),
		// 							IssueID: to.Ptr("PerformanceDataMissing"),
		// 					}},
		// 					MigrationTargetPlatform: to.Ptr(armmigrationassessment.TargetTypeAzureSQLDatabase),
		// 					MonthlyComputeCost: to.Ptr[float32](1199.994624),
		// 					MonthlyStorageCost: to.Ptr[float32](7.13),
		// 					RecommendationReasonings: []*armmigrationassessment.SQLRecommendationReasoning{
		// 					},
		// 					ReplicaAzureSQLSKU: []*armmigrationassessment.AzureSQLPaasSKUDTO{
		// 					},
		// 					SecuritySuitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
		// 					ShouldProvisionReplicas: to.Ptr(false),
		// 					SKUReplicationMode: to.Ptr(armmigrationassessment.SKUReplicationModeNotApplicable),
		// 					Suitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
		// 				},
		// 				AzureSQLMISuitabilityDetails: &armmigrationassessment.SQLAssessmentV2PaasSuitabilityData{
		// 					AzureSQLSKU: &armmigrationassessment.AzureSQLPaasSKUDTO{
		// 						AzureSQLComputeTier: to.Ptr(armmigrationassessment.ComputeTierProvisioned),
		// 						AzureSQLHardwareGeneration: to.Ptr(armmigrationassessment.HardwareGenerationGen5),
		// 						AzureSQLServiceTier: to.Ptr(armmigrationassessment.AzureSQLServiceTierGeneralPurpose),
		// 						AzureSQLTargetType: to.Ptr(armmigrationassessment.TargetTypeAzureSQLManagedInstance),
		// 						Cores: to.Ptr[int32](4),
		// 						PredictedDataSizeInMB: to.Ptr[float32](6696.25),
		// 						PredictedLogSizeInMB: to.Ptr[float32](14220.125),
		// 						StorageMaxSizeInMB: to.Ptr[float32](65536),
		// 					},
		// 					CostComponents: []*armmigrationassessment.CostComponent{
		// 						{
		// 							Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlySecurityCost),
		// 							Value: to.Ptr[float32](15),
		// 					}},
		// 					MigrationGuidelines: []*armmigrationassessment.SQLMigrationGuideline{
		// 						{
		// 							GuidelineID: to.Ptr("MigrateDataFromSourceSqlServerToMI"),
		// 							MigrationGuidelineCategory: to.Ptr(armmigrationassessment.SQLMigrationGuidelineCategoryFailoverCluterInstanceGuideLine),
		// 							MigrationGuidelineContext: []*armmigrationassessment.MigrationGuidelineContext{
		// 							},
		// 					}},
		// 					MigrationIssues: []*armmigrationassessment.SQLAssessmentMigrationIssue{
		// 						{
		// 							ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
		// 							},
		// 							IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryInternal),
		// 							IssueID: to.Ptr("PerformanceDataMissing"),
		// 					}},
		// 					MigrationTargetPlatform: to.Ptr(armmigrationassessment.TargetTypeAzureSQLManagedInstance),
		// 					MonthlyComputeCost: to.Ptr[float32](199.999104),
		// 					MonthlyStorageCost: to.Ptr[float32](3.68),
		// 					RecommendationReasonings: []*armmigrationassessment.SQLRecommendationReasoning{
		// 						{
		// 							ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 								{
		// 									ContextKey: to.Ptr("CoreUtilizationPercentage"),
		// 									ContextValue: to.Ptr("0"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("AllocatedCores"),
		// 									ContextValue: to.Ptr("4"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("TargetCoresRecommended"),
		// 									ContextValue: to.Ptr("4"),
		// 							}},
		// 							ReasoningCategory: to.Ptr("CpuType"),
		// 							ReasoningID: to.Ptr("MeetsInstanceCpuScalingFactorRequirement"),
		// 						},
		// 						{
		// 							ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 								{
		// 									ContextKey: to.Ptr("TotalStorageConsumedByDataFilesInGB"),
		// 									ContextValue: to.Ptr("6.539"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("TotalStorageConsumedByLogFilesInGB"),
		// 									ContextValue: to.Ptr("13.887"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("TotalRecommendedStorageInGB"),
		// 									ContextValue: to.Ptr("64"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("IncrementsInStorageInGB"),
		// 									ContextValue: to.Ptr("32"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("MaxAvailableStorageInGB"),
		// 									ContextValue: to.Ptr("2048"),
		// 							}},
		// 							ReasoningCategory: to.Ptr("StorageType"),
		// 							ReasoningID: to.Ptr("MeetsInstanceStorageRequirement"),
		// 						},
		// 						{
		// 							ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 								{
		// 									ContextKey: to.Ptr("InstanceLevelThroughputRequiredInMBPerSec"),
		// 									ContextValue: to.Ptr("0"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("InstanceLevelThroughputAvailableInMBPerSec"),
		// 									ContextValue: to.Ptr("10"),
		// 							}},
		// 							ReasoningCategory: to.Ptr("ThroughputType"),
		// 							ReasoningID: to.Ptr("MeetsInstanceIOThroughputRequirement"),
		// 						},
		// 						{
		// 							ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 								{
		// 									ContextKey: to.Ptr("InstanceLevelIOPSRequiredInIOPerSec"),
		// 									ContextValue: to.Ptr("0"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("InstanceLevelIOPSAvailableInIOPerSec"),
		// 									ContextValue: to.Ptr("6000"),
		// 							}},
		// 							ReasoningCategory: to.Ptr("IOPSType"),
		// 							ReasoningID: to.Ptr("MeetsInstanceIOPSDataLogRequirement"),
		// 						},
		// 						{
		// 							ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 								{
		// 									ContextKey: to.Ptr("ComputeTier"),
		// 									ContextValue: to.Ptr("Provisioned"),
		// 							}},
		// 							ReasoningCategory: to.Ptr("ComputeTier"),
		// 							ReasoningID: to.Ptr("MeetsProvisionedComputeTierRequirement"),
		// 						},
		// 						{
		// 							ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 								{
		// 									ContextKey: to.Ptr("ServiceTier"),
		// 									ContextValue: to.Ptr("GeneralPurpose"),
		// 							}},
		// 							ReasoningCategory: to.Ptr("ServiceTier"),
		// 							ReasoningID: to.Ptr("MeetsServiceTierRequirement"),
		// 					}},
		// 					ReplicaAzureSQLSKU: []*armmigrationassessment.AzureSQLPaasSKUDTO{
		// 					},
		// 					SecuritySuitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
		// 					SharedResources: &armmigrationassessment.SharedResourcesDTO{
		// 						NumberOfMounts: to.Ptr[int32](0),
		// 						SharedDataDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
		// 						},
		// 						SharedLogDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
		// 						},
		// 						SharedTempDbDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
		// 						},
		// 					},
		// 					ShouldProvisionReplicas: to.Ptr(false),
		// 					SKUReplicationMode: to.Ptr(armmigrationassessment.SKUReplicationModeNotApplicable),
		// 					Suitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
		// 				},
		// 				AzureSQLVMSuitabilityDetails: &armmigrationassessment.SQLAssessmentV2IaasSuitabilityData{
		// 					AzureSQLSKU: &armmigrationassessment.AzureSQLIaasSKUDTO{
		// 						AzureSQLTargetType: to.Ptr(armmigrationassessment.TargetTypeAzureSQLVirtualMachine),
		// 						DataDiskSizes: []*armmigrationassessment.AzureManagedDiskSKUDTO{
		// 							{
		// 								DiskRedundancy: to.Ptr(armmigrationassessment.AzureManagedDiskSKUDTODiskRedundancyLRS),
		// 								DiskSize: to.Ptr(armmigrationassessment.AzureDiskSizePremiumP2),
		// 								DiskType: to.Ptr(armmigrationassessment.AzureManagedDiskSKUDTODiskTypePremium),
		// 								StorageCost: to.Ptr[float32](1.2),
		// 						}},
		// 						LogDiskSizes: []*armmigrationassessment.AzureManagedDiskSKUDTO{
		// 							{
		// 								DiskRedundancy: to.Ptr(armmigrationassessment.AzureManagedDiskSKUDTODiskRedundancyLRS),
		// 								DiskSize: to.Ptr(armmigrationassessment.AzureDiskSizePremiumP3),
		// 								DiskType: to.Ptr(armmigrationassessment.AzureManagedDiskSKUDTODiskTypePremium),
		// 								StorageCost: to.Ptr[float32](2.4),
		// 						}},
		// 						VirtualMachineSize: &armmigrationassessment.AzureVirtualMachineSKUDTO{
		// 							AvailableCores: to.Ptr[int32](2),
		// 							AzureSKUName: to.Ptr(armmigrationassessment.AzureVMSizeStandardD2AsV4),
		// 							AzureVMFamily: to.Ptr(armmigrationassessment.AzureVMFamilyDasv4Series),
		// 							Cores: to.Ptr[int32](2),
		// 							MaxNetworkInterfaces: to.Ptr[int32](0),
		// 						},
		// 					},
		// 					CostComponents: []*armmigrationassessment.CostComponent{
		// 						{
		// 							Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlySecurityCost),
		// 							Value: to.Ptr[float32](15),
		// 					}},
		// 					MigrationGuidelines: []*armmigrationassessment.SQLMigrationGuideline{
		// 						{
		// 							GuidelineID: to.Ptr("MigrateDataFromSourceSqlServerToVM"),
		// 							MigrationGuidelineCategory: to.Ptr(armmigrationassessment.SQLMigrationGuidelineCategoryFailoverCluterInstanceGuideLine),
		// 							MigrationGuidelineContext: []*armmigrationassessment.MigrationGuidelineContext{
		// 							},
		// 					}},
		// 					MigrationIssues: []*armmigrationassessment.SQLAssessmentMigrationIssue{
		// 						{
		// 							ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
		// 							},
		// 							IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryInternal),
		// 							IssueID: to.Ptr("PerformanceDataMissing"),
		// 					}},
		// 					MigrationTargetPlatform: to.Ptr(armmigrationassessment.TargetTypeAzureSQLVirtualMachine),
		// 					MonthlyComputeCost: to.Ptr[float32](26.361408),
		// 					MonthlyStorageCost: to.Ptr[float32](3.6),
		// 					RecommendationReasonings: []*armmigrationassessment.SQLRecommendationReasoning{
		// 						{
		// 							ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 								{
		// 									ContextKey: to.Ptr("CoreUtilizationPercentage"),
		// 									ContextValue: to.Ptr("0"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("AllocatedCores"),
		// 									ContextValue: to.Ptr("4"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("TargetCoresRecommended"),
		// 									ContextValue: to.Ptr("2"),
		// 							}},
		// 							ReasoningCategory: to.Ptr("CpuType"),
		// 							ReasoningID: to.Ptr("MeetsVirtualMachineCpuScalingFactorRequirement"),
		// 						},
		// 						{
		// 							ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 								{
		// 									ContextKey: to.Ptr("BufferPoolSizeInGB"),
		// 									ContextValue: to.Ptr("0"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("RAMRecommendedInGB"),
		// 									ContextValue: to.Ptr("8"),
		// 							}},
		// 							ReasoningCategory: to.Ptr("MemoryType"),
		// 							ReasoningID: to.Ptr("MeetsVirtualMachineMemoryRequirement"),
		// 						},
		// 						{
		// 							ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 								{
		// 									ContextKey: to.Ptr("VirtualMachineIOPSRequiredInIOPerSec"),
		// 									ContextValue: to.Ptr("0"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("VirtualMachineUncachedDiskIOPSAvailableInIOPerSec"),
		// 									ContextValue: to.Ptr("3200"),
		// 							}},
		// 							ReasoningCategory: to.Ptr("IOPSType"),
		// 							ReasoningID: to.Ptr("MeetsVirtualMachineIOPSRequirement"),
		// 						},
		// 						{
		// 							ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 								{
		// 									ContextKey: to.Ptr("VirtualMachineIOThroughputRequiredInInMBPerSec"),
		// 									ContextValue: to.Ptr("0"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("VirtualMachineIOThroughputUncachedAvailableInIOPerSec"),
		// 									ContextValue: to.Ptr("48"),
		// 							}},
		// 							ReasoningCategory: to.Ptr("ThroughputType"),
		// 							ReasoningID: to.Ptr("MeetsVirtualMachineIOThroughputRequirement"),
		// 						},
		// 						{
		// 							ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 								{
		// 									ContextKey: to.Ptr("CommaSeperatedManagedDisks"),
		// 									ContextValue: to.Ptr("1 PremiumSSD_P2,1 PremiumSSD_P3"),
		// 							}},
		// 							ReasoningCategory: to.Ptr("StorageType"),
		// 							ReasoningID: to.Ptr("MeetsVirtualMachineStorageRequirement"),
		// 						},
		// 						{
		// 							ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 								{
		// 									ContextKey: to.Ptr("VirtualMachineSize"),
		// 									ContextValue: to.Ptr("standardDASv4Family"),
		// 							}},
		// 							ReasoningCategory: to.Ptr("ServiceTier"),
		// 							ReasoningID: to.Ptr("MeetsVirtualMachineFamilyRequirement"),
		// 					}},
		// 					ReplicaAzureSQLSKU: []*armmigrationassessment.AzureSQLIaasSKUDTO{
		// 					},
		// 					SecuritySuitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
		// 					SharedResources: &armmigrationassessment.SharedResourcesDTO{
		// 						NumberOfMounts: to.Ptr[int32](0),
		// 						SharedDataDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
		// 						},
		// 						SharedLogDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
		// 						},
		// 						SharedTempDbDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
		// 						},
		// 					},
		// 					ShouldProvisionReplicas: to.Ptr(false),
		// 					SKUReplicationMode: to.Ptr(armmigrationassessment.SKUReplicationModeNotApplicable),
		// 					Suitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
		// 				},
		// 				ConfidenceRatingInPercentage: to.Ptr[float32](0),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T09:20:24.554Z"); return t}()),
		// 				DatabaseSummary: &armmigrationassessment.AssessedSQLInstanceDatabaseSummary{
		// 					LargestDatabaseSizeInMB: to.Ptr[float32](20815),
		// 					NumberOfUserDatabases: to.Ptr[int32](12),
		// 					TotalDatabaseSizeInMB: to.Ptr[float32](20916.375),
		// 					TotalDiscoveredUserDatabases: to.Ptr[int32](12),
		// 				},
		// 				HasScanOccurred: to.Ptr(true),
		// 				InstanceName: to.Ptr("MSSQLSERVER"),
		// 				IsClustered: to.Ptr(false),
		// 				IsHighAvailabilityEnabled: to.Ptr(false),
		// 				LogicalDisks: []*armmigrationassessment.AssessedSQLInstanceDiskDetails{
		// 					{
		// 						DiskID: to.Ptr("d:"),
		// 						DiskSizeInMB: to.Ptr[float32](20916.375),
		// 						MegabytesPerSecondOfRead: to.Ptr[float32](0),
		// 						MegabytesPerSecondOfWrite: to.Ptr[float32](0),
		// 						NumberOfReadOperationsPerSecond: to.Ptr[float32](0),
		// 						NumberOfWriteOperationsPerSecond: to.Ptr[float32](0),
		// 				}},
		// 				MachineArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/dhlodhiccy/providers/Microsoft.Migrate/assessmentprojects/dhccy30279project/machines/daf51245-560b-4caa-9c18-5409ab93d319"),
		// 				MachineName: to.Ptr("SQLTestDBVM46"),
		// 				MemoryInUseInMB: to.Ptr[float32](0),
		// 				NumberOfCoresAllocated: to.Ptr[int32](4),
		// 				PercentageCoresUtilization: to.Ptr[float32](0),
		// 				ProductSupportStatus: &armmigrationassessment.ProductSupportStatus{
		// 					CurrentEsuYear: to.Ptr("Unknown"),
		// 					EsuStatus: to.Ptr("Unknown"),
		// 					Eta: to.Ptr[int32](34),
		// 					ExtendedSecurityUpdateYear1EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 					ExtendedSecurityUpdateYear2EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 					ExtendedSecurityUpdateYear3EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 					ExtendedSupportEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-07-14T00:00:00.000Z"); return t}()),
		// 					MainstreamEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-13T00:00:00.000Z"); return t}()),
		// 					ServicePackStatus: to.Ptr("Unknown"),
		// 					SupportStatus: to.Ptr("Extended"),
		// 				},
		// 				RecommendedAzureSQLTargetType: to.Ptr(armmigrationassessment.TargetTypeAzureSQLVirtualMachine),
		// 				RecommendedSuitability: to.Ptr(armmigrationassessment.RecommendedSuitabilitySuitableForSQLVM),
		// 				RecommendedTargetReasonings: []*armmigrationassessment.SQLRecommendationReasoning{
		// 					{
		// 						ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 							{
		// 								ContextKey: to.Ptr("MigrationIssuesCount"),
		// 								ContextValue: to.Ptr("0"),
		// 							},
		// 							{
		// 								ContextKey: to.Ptr("MigrationWarningsCount"),
		// 								ContextValue: to.Ptr("0"),
		// 							},
		// 							{
		// 								ContextKey: to.Ptr("RecommendedCost"),
		// 								ContextValue: to.Ptr("29.961408"),
		// 						}},
		// 						ReasoningCategory: to.Ptr("RecommendedTargetReasoning"),
		// 						ReasoningID: to.Ptr("RecommendedForSQLVM"),
		// 					},
		// 					{
		// 						ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 							{
		// 								ContextKey: to.Ptr("RecommendedCost"),
		// 								ContextValue: to.Ptr("29.961408"),
		// 						}},
		// 						ReasoningCategory: to.Ptr("RecommendedTargetReasoning"),
		// 						ReasoningID: to.Ptr("RecommendedCost"),
		// 					},
		// 					{
		// 						ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 							{
		// 								ContextKey: to.Ptr("SQLMICost"),
		// 								ContextValue: to.Ptr("203.679104"),
		// 						}},
		// 						ReasoningCategory: to.Ptr("RecommendedTargetReasoning"),
		// 						ReasoningID: to.Ptr("SQLMICost"),
		// 					},
		// 					{
		// 						ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 							{
		// 								ContextKey: to.Ptr("SQLDBCost"),
		// 								ContextValue: to.Ptr("1207.124624"),
		// 						}},
		// 						ReasoningCategory: to.Ptr("RecommendedTargetReasoning"),
		// 						ReasoningID: to.Ptr("SQLDBCost"),
		// 				}},
		// 				SizingCriterion: to.Ptr(armmigrationassessment.AssessmentSizingCriterionPerformanceBased),
		// 				SQLEdition: to.Ptr("Enterprise Edition: Core-based Licensing (64-bit)"),
		// 				SQLInstanceSdsArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourcegroups/dhlodhiccy/providers/microsoft.offazure/mastersites/dhccy37819mastersite/sqlsites/dhccy37819sqlsites/sqlservers/a5a6441d-9f63-463e-a720-73fecd90a2e4"),
		// 				SQLVersion: to.Ptr("SQL Server 2016"),
		// 				StorageTypeBasedDetails: []*armmigrationassessment.AssessedSQLInstanceStorageDetails{
		// 					{
		// 						DiskSizeInMB: to.Ptr[float32](6696.25),
		// 						MegabytesPerSecondOfRead: to.Ptr[float32](0),
		// 						MegabytesPerSecondOfWrite: to.Ptr[float32](0),
		// 						NumberOfReadOperationsPerSecond: to.Ptr[float32](0),
		// 						NumberOfWriteOperationsPerSecond: to.Ptr[float32](0),
		// 						StorageType: to.Ptr("Rows"),
		// 					},
		// 					{
		// 						DiskSizeInMB: to.Ptr[float32](14220.125),
		// 						MegabytesPerSecondOfRead: to.Ptr[float32](0),
		// 						MegabytesPerSecondOfWrite: to.Ptr[float32](0),
		// 						NumberOfReadOperationsPerSecond: to.Ptr[float32](0),
		// 						NumberOfWriteOperationsPerSecond: to.Ptr[float32](0),
		// 						StorageType: to.Ptr("Log"),
		// 				}},
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T09:20:25.219Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
