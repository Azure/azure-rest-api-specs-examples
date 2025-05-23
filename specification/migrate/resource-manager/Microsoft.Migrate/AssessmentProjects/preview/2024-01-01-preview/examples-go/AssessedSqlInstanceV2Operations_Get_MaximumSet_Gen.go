package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AssessedSqlInstanceV2Operations_Get_MaximumSet_Gen.json
func ExampleAssessedSQLInstanceV2OperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessedSQLInstanceV2OperationsClient().Get(ctx, "rgmigrate", "fci-test6904project", "test_fci_hadr", "test_swagger_1", "3c6574cf-b4e1-4fdc-93db-6bbcc570dda2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessedSQLInstanceV2 = armmigrationassessment.AssessedSQLInstanceV2{
	// 	Name: to.Ptr("3c6574cf-b4e1-4fdc-93db-6bbcc570dda2"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/sqlAssessments/assessedSqlInstances"),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.Migrate/assessmentprojects/fci-ankit-test6904project/groups/test_fci_hadr/sqlAssessments/test_swagger_1/assessedSqlInstances/3c6574cf-b4e1-4fdc-93db-6bbcc570dda2"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 	},
	// 	Properties: &armmigrationassessment.AssessedSQLInstanceV2Properties{
	// 		AvailabilityReplicaSummary: &armmigrationassessment.SQLAvailabilityReplicaSummary{
	// 			NumberOfAsynchronousNonReadReplicas: to.Ptr[int32](0),
	// 			NumberOfAsynchronousReadReplicas: to.Ptr[int32](0),
	// 			NumberOfPrimaryReplicas: to.Ptr[int32](6),
	// 			NumberOfSynchronousNonReadReplicas: to.Ptr[int32](0),
	// 			NumberOfSynchronousReadReplicas: to.Ptr[int32](0),
	// 		},
	// 		AzureSQLDBSuitabilityDetails: &armmigrationassessment.SQLAssessmentV2PaasSuitabilityData{
	// 			CostComponents: []*armmigrationassessment.CostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlySecurityCost),
	// 					Value: to.Ptr[float32](0),
	// 			}},
	// 			MigrationGuidelines: []*armmigrationassessment.SQLMigrationGuideline{
	// 			},
	// 			MigrationIssues: []*armmigrationassessment.SQLAssessmentMigrationIssue{
	// 				{
	// 					ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
	// 					},
	// 					IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryInternal),
	// 					IssueID: to.Ptr("SuitabilityReportMissing"),
	// 				},
	// 				{
	// 					ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
	// 					},
	// 					IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryInternal),
	// 					IssueID: to.Ptr("PerformanceDataMissing"),
	// 				},
	// 				{
	// 					ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
	// 						{
	// 							ObjectName: to.Ptr("MSSQLSERVER"),
	// 							ObjectType: to.Ptr("SqlInstance"),
	// 					}},
	// 					IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryWarning),
	// 					IssueID: to.Ptr("PerfBasedFallbackToAsOnPremises"),
	// 			}},
	// 			MigrationTargetPlatform: to.Ptr(armmigrationassessment.TargetTypeAzureSQLDatabase),
	// 			MonthlyComputeCost: to.Ptr[float32](833.165424),
	// 			MonthlyStorageCost: to.Ptr[float32](24.69),
	// 			RecommendationReasonings: []*armmigrationassessment.SQLRecommendationReasoning{
	// 			},
	// 			ReplicaAzureSQLSKU: []*armmigrationassessment.AzureSQLPaasSKUDTO{
	// 			},
	// 			SecuritySuitability: to.Ptr(armmigrationassessment.CloudSuitabilityUnknown),
	// 			ShouldProvisionReplicas: to.Ptr(false),
	// 			SKUReplicationMode: to.Ptr(armmigrationassessment.SKUReplicationModeNotApplicable),
	// 			Suitability: to.Ptr(armmigrationassessment.CloudSuitabilityReadinessUnknown),
	// 		},
	// 		AzureSQLMISuitabilityDetails: &armmigrationassessment.SQLAssessmentV2PaasSuitabilityData{
	// 			AzureSQLSKU: &armmigrationassessment.AzureSQLPaasSKUDTO{
	// 				AzureSQLComputeTier: to.Ptr(armmigrationassessment.ComputeTierProvisioned),
	// 				AzureSQLHardwareGeneration: to.Ptr(armmigrationassessment.HardwareGenerationGen5),
	// 				AzureSQLServiceTier: to.Ptr(armmigrationassessment.AzureSQLServiceTierBusinessCritical),
	// 				AzureSQLTargetType: to.Ptr(armmigrationassessment.TargetTypeAzureSQLManagedInstance),
	// 				Cores: to.Ptr[int32](4),
	// 				PredictedDataSizeInMB: to.Ptr[float32](112),
	// 				PredictedLogSizeInMB: to.Ptr[float32](112),
	// 				StorageMaxSizeInMB: to.Ptr[float32](32768),
	// 			},
	// 			CostComponents: []*armmigrationassessment.CostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlySecurityCost),
	// 					Value: to.Ptr[float32](0),
	// 			}},
	// 			MigrationGuidelines: []*armmigrationassessment.SQLMigrationGuideline{
	// 			},
	// 			MigrationIssues: []*armmigrationassessment.SQLAssessmentMigrationIssue{
	// 				{
	// 					ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
	// 					},
	// 					IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryInternal),
	// 					IssueID: to.Ptr("SuitabilityReportMissing"),
	// 				},
	// 				{
	// 					ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
	// 					},
	// 					IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryInternal),
	// 					IssueID: to.Ptr("PerformanceDataMissing"),
	// 				},
	// 				{
	// 					ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
	// 						{
	// 							ObjectName: to.Ptr("MSSQLSERVER"),
	// 							ObjectType: to.Ptr("SqlInstance"),
	// 					}},
	// 					IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryWarning),
	// 					IssueID: to.Ptr("PerfBasedFallbackToAsOnPremises"),
	// 			}},
	// 			MigrationTargetPlatform: to.Ptr(armmigrationassessment.TargetTypeAzureSQLManagedInstance),
	// 			MonthlyComputeCost: to.Ptr[float32](475.999296),
	// 			MonthlyStorageCost: to.Ptr[float32](0),
	// 			RecommendationReasonings: []*armmigrationassessment.SQLRecommendationReasoning{
	// 				{
	// 					ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
	// 						{
	// 							ContextKey: to.Ptr("AllocatedCores"),
	// 							ContextValue: to.Ptr("2"),
	// 						},
	// 						{
	// 							ContextKey: to.Ptr("TargetCoresRecommended"),
	// 							ContextValue: to.Ptr("4"),
	// 					}},
	// 					ReasoningCategory: to.Ptr("CpuType"),
	// 					ReasoningID: to.Ptr("InstanceCpuScalingFactorRequirementUnavailable"),
	// 				},
	// 				{
	// 					ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
	// 						{
	// 							ContextKey: to.Ptr("TotalStorageConsumedByDataFilesInGB"),
	// 							ContextValue: to.Ptr("0.109"),
	// 						},
	// 						{
	// 							ContextKey: to.Ptr("TotalStorageConsumedByLogFilesInGB"),
	// 							ContextValue: to.Ptr("0.109"),
	// 						},
	// 						{
	// 							ContextKey: to.Ptr("TotalRecommendedStorageInGB"),
	// 							ContextValue: to.Ptr("32"),
	// 						},
	// 						{
	// 							ContextKey: to.Ptr("IncrementsInStorageInGB"),
	// 							ContextValue: to.Ptr("32"),
	// 						},
	// 						{
	// 							ContextKey: to.Ptr("MaxAvailableStorageInGB"),
	// 							ContextValue: to.Ptr("1024"),
	// 					}},
	// 					ReasoningCategory: to.Ptr("StorageType"),
	// 					ReasoningID: to.Ptr("MeetsInstanceStorageRequirement"),
	// 				},
	// 				{
	// 					ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
	// 						{
	// 							ContextKey: to.Ptr("AlwaysOnDatabaseServiceTier"),
	// 							ContextValue: to.Ptr("BusinessCritical"),
	// 					}},
	// 					ReasoningCategory: to.Ptr("AlwaysOnAG"),
	// 					ReasoningID: to.Ptr("MeetsAlwaysOnInstanceReplicaRequirementWithoutSecondaries"),
	// 				},
	// 				{
	// 					ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
	// 						{
	// 							ContextKey: to.Ptr("ComputeTier"),
	// 							ContextValue: to.Ptr("Provisioned"),
	// 					}},
	// 					ReasoningCategory: to.Ptr("ComputeTier"),
	// 					ReasoningID: to.Ptr("MeetsProvisionedComputeTierRequirement"),
	// 				},
	// 				{
	// 					ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
	// 						{
	// 							ContextKey: to.Ptr("ServiceTier"),
	// 							ContextValue: to.Ptr("BusinessCritical"),
	// 					}},
	// 					ReasoningCategory: to.Ptr("ServiceTier"),
	// 					ReasoningID: to.Ptr("MeetsServiceTierRequirement"),
	// 			}},
	// 			ReplicaAzureSQLSKU: []*armmigrationassessment.AzureSQLPaasSKUDTO{
	// 			},
	// 			SecuritySuitability: to.Ptr(armmigrationassessment.CloudSuitabilityUnknown),
	// 			SharedResources: &armmigrationassessment.SharedResourcesDTO{
	// 				NumberOfMounts: to.Ptr[int32](0),
	// 				SharedDataDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 				},
	// 				SharedLogDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 				},
	// 				SharedTempDbDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 				},
	// 			},
	// 			ShouldProvisionReplicas: to.Ptr(false),
	// 			SKUReplicationMode: to.Ptr(armmigrationassessment.SKUReplicationModeNotApplicable),
	// 			Suitability: to.Ptr(armmigrationassessment.CloudSuitabilityReadinessUnknown),
	// 		},
	// 		AzureSQLVMSuitabilityDetails: &armmigrationassessment.SQLAssessmentV2IaasSuitabilityData{
	// 			AzureSQLSKU: &armmigrationassessment.AzureSQLIaasSKUDTO{
	// 				AzureSQLTargetType: to.Ptr(armmigrationassessment.TargetTypeAzureSQLVirtualMachine),
	// 				DataDiskSizes: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 					{
	// 						DiskRedundancy: to.Ptr(armmigrationassessment.AzureManagedDiskSKUDTODiskRedundancyLRS),
	// 						DiskSize: to.Ptr(armmigrationassessment.AzureDiskSizePremiumP2),
	// 						DiskType: to.Ptr(armmigrationassessment.AzureManagedDiskSKUDTODiskTypePremium),
	// 						RecommendedIops: to.Ptr[float32](120),
	// 						RecommendedSizeInGib: to.Ptr[float32](8),
	// 						RecommendedThroughputInMbps: to.Ptr[float32](25),
	// 						StorageCost: to.Ptr[float32](1.56),
	// 				}},
	// 				LogDiskSizes: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 					{
	// 						DiskRedundancy: to.Ptr(armmigrationassessment.AzureManagedDiskSKUDTODiskRedundancyLRS),
	// 						DiskSize: to.Ptr(armmigrationassessment.AzureDiskSizePremiumP2),
	// 						DiskType: to.Ptr(armmigrationassessment.AzureManagedDiskSKUDTODiskTypePremium),
	// 						RecommendedIops: to.Ptr[float32](120),
	// 						RecommendedSizeInGib: to.Ptr[float32](8),
	// 						RecommendedThroughputInMbps: to.Ptr[float32](25),
	// 						StorageCost: to.Ptr[float32](1.56),
	// 				}},
	// 				VirtualMachineSize: &armmigrationassessment.AzureVirtualMachineSKUDTO{
	// 					AvailableCores: to.Ptr[int32](2),
	// 					AzureSKUName: to.Ptr(armmigrationassessment.AzureVMSizeStandardD2AsV4),
	// 					AzureVMFamily: to.Ptr(armmigrationassessment.AzureVMFamilyDasv4Series),
	// 					Cores: to.Ptr[int32](2),
	// 					MaxNetworkInterfaces: to.Ptr[int32](0),
	// 				},
	// 			},
	// 			CostComponents: []*armmigrationassessment.CostComponent{
	// 				{
	// 					Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlySecurityCost),
	// 					Value: to.Ptr[float32](0),
	// 			}},
	// 			MigrationGuidelines: []*armmigrationassessment.SQLMigrationGuideline{
	// 				{
	// 					GuidelineID: to.Ptr("MigrateDataFromSourceSqlServerToVM"),
	// 					MigrationGuidelineCategory: to.Ptr(armmigrationassessment.SQLMigrationGuidelineCategoryFailoverCluterInstanceGuideLine),
	// 					MigrationGuidelineContext: []*armmigrationassessment.MigrationGuidelineContext{
	// 					},
	// 				},
	// 				{
	// 					GuidelineID: to.Ptr("MigrateNonReplicatedDatabasesAndAutoSeedSecondaryIfNeeded"),
	// 					MigrationGuidelineCategory: to.Ptr(armmigrationassessment.SQLMigrationGuidelineCategoryAvailabilityGroupGuideline),
	// 					MigrationGuidelineContext: []*armmigrationassessment.MigrationGuidelineContext{
	// 					},
	// 				},
	// 				{
	// 					GuidelineID: to.Ptr("MigratePrimaryReplicaAndAllOtherReplicasToIaaS"),
	// 					MigrationGuidelineCategory: to.Ptr(armmigrationassessment.SQLMigrationGuidelineCategoryAvailabilityGroupGuideline),
	// 					MigrationGuidelineContext: []*armmigrationassessment.MigrationGuidelineContext{
	// 					},
	// 			}},
	// 			MigrationIssues: []*armmigrationassessment.SQLAssessmentMigrationIssue{
	// 				{
	// 					ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
	// 					},
	// 					IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryInternal),
	// 					IssueID: to.Ptr("PerformanceDataMissing"),
	// 				},
	// 				{
	// 					ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
	// 						{
	// 							ObjectName: to.Ptr("MSSQLSERVER"),
	// 							ObjectType: to.Ptr("SqlInstance"),
	// 					}},
	// 					IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryWarning),
	// 					IssueID: to.Ptr("PerfBasedFallbackToAsOnPremises"),
	// 			}},
	// 			MigrationTargetPlatform: to.Ptr(armmigrationassessment.TargetTypeAzureSQLVirtualMachine),
	// 			MonthlyComputeCost: to.Ptr[float32](32.94432),
	// 			MonthlyStorageCost: to.Ptr[float32](3.12),
	// 			RecommendationReasonings: []*armmigrationassessment.SQLRecommendationReasoning{
	// 				{
	// 					ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
	// 						{
	// 							ContextKey: to.Ptr("AllocatedCores"),
	// 							ContextValue: to.Ptr("2"),
	// 						},
	// 						{
	// 							ContextKey: to.Ptr("TargetCoresRecommended"),
	// 							ContextValue: to.Ptr("2"),
	// 					}},
	// 					ReasoningCategory: to.Ptr("CpuType"),
	// 					ReasoningID: to.Ptr("VirtualMachineCpuScalingFactorRequirementUnavailable"),
	// 				},
	// 				{
	// 					ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
	// 						{
	// 							ContextKey: to.Ptr("BufferPoolSizeInGB"),
	// 							ContextValue: to.Ptr("4"),
	// 						},
	// 						{
	// 							ContextKey: to.Ptr("RAMRecommendedInGB"),
	// 							ContextValue: to.Ptr("8"),
	// 					}},
	// 					ReasoningCategory: to.Ptr("MemoryType"),
	// 					ReasoningID: to.Ptr("MeetsVirtualMachineMemoryRequirement"),
	// 				},
	// 				{
	// 					ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
	// 						{
	// 							ContextKey: to.Ptr("CommaSeperatedManagedDisks"),
	// 							ContextValue: to.Ptr("2 PremiumSSD(P2)"),
	// 					}},
	// 					ReasoningCategory: to.Ptr("StorageType"),
	// 					ReasoningID: to.Ptr("MeetsVirtualMachineStorageRequirement"),
	// 				},
	// 				{
	// 					ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
	// 						{
	// 							ContextKey: to.Ptr("VirtualMachineSize"),
	// 							ContextValue: to.Ptr("standardDASv4Family"),
	// 					}},
	// 					ReasoningCategory: to.Ptr("ServiceTier"),
	// 					ReasoningID: to.Ptr("MeetsVirtualMachineFamilyRequirement"),
	// 			}},
	// 			ReplicaAzureSQLSKU: []*armmigrationassessment.AzureSQLIaasSKUDTO{
	// 			},
	// 			SecuritySuitability: to.Ptr(armmigrationassessment.CloudSuitabilityUnknown),
	// 			SharedResources: &armmigrationassessment.SharedResourcesDTO{
	// 				NumberOfMounts: to.Ptr[int32](0),
	// 				SharedDataDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 				},
	// 				SharedLogDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 				},
	// 				SharedTempDbDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 				},
	// 			},
	// 			ShouldProvisionReplicas: to.Ptr(false),
	// 			SKUReplicationMode: to.Ptr(armmigrationassessment.SKUReplicationModeNotApplicable),
	// 			Suitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
	// 		},
	// 		ConfidenceRatingInPercentage: to.Ptr[float32](0),
	// 		CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-08T05:22:34.935Z"); return t}()),
	// 		DatabaseSummary: &armmigrationassessment.AssessedSQLInstanceDatabaseSummary{
	// 			LargestDatabaseSizeInMB: to.Ptr[float32](144),
	// 			NumberOfUserDatabases: to.Ptr[int32](6),
	// 			TotalDatabaseSizeInMB: to.Ptr[float32](224),
	// 			TotalDiscoveredUserDatabases: to.Ptr[int32](7),
	// 		},
	// 		HasScanOccurred: to.Ptr(true),
	// 		InstanceName: to.Ptr("MSSQLSERVER"),
	// 		IsClustered: to.Ptr(false),
	// 		IsHighAvailabilityEnabled: to.Ptr(true),
	// 		LogicalDisks: []*armmigrationassessment.AssessedSQLInstanceDiskDetails{
	// 			{
	// 				DiskID: to.Ptr("s:"),
	// 				DiskSizeInMB: to.Ptr[float32](104),
	// 				MegabytesPerSecondOfRead: to.Ptr[float32](0),
	// 				MegabytesPerSecondOfWrite: to.Ptr[float32](0),
	// 				NumberOfReadOperationsPerSecond: to.Ptr[float32](0),
	// 				NumberOfWriteOperationsPerSecond: to.Ptr[float32](0),
	// 			},
	// 			{
	// 				DiskID: to.Ptr("l:"),
	// 				DiskSizeInMB: to.Ptr[float32](104),
	// 				MegabytesPerSecondOfRead: to.Ptr[float32](0),
	// 				MegabytesPerSecondOfWrite: to.Ptr[float32](0),
	// 				NumberOfReadOperationsPerSecond: to.Ptr[float32](0),
	// 				NumberOfWriteOperationsPerSecond: to.Ptr[float32](0),
	// 			},
	// 			{
	// 				DiskID: to.Ptr("c:"),
	// 				DiskSizeInMB: to.Ptr[float32](16),
	// 				MegabytesPerSecondOfRead: to.Ptr[float32](0),
	// 				MegabytesPerSecondOfWrite: to.Ptr[float32](0),
	// 				NumberOfReadOperationsPerSecond: to.Ptr[float32](0),
	// 				NumberOfWriteOperationsPerSecond: to.Ptr[float32](0),
	// 		}},
	// 		MachineArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.Migrate/assessmentprojects/fci-ankit-test6904project/machines/184b8433-6991-45db-9f57-3dd3f6397a60"),
	// 		MachineName: to.Ptr("SQLAGVM02"),
	// 		MemoryInUseInMB: to.Ptr[float32](4096),
	// 		NumberOfCoresAllocated: to.Ptr[int32](2),
	// 		PercentageCoresUtilization: to.Ptr[float32](100),
	// 		ProductSupportStatus: &armmigrationassessment.ProductSupportStatus{
	// 			CurrentEsuYear: to.Ptr("Unknown"),
	// 			EsuStatus: to.Ptr("Unknown"),
	// 			Eta: to.Ptr[int32](34),
	// 			ExtendedSecurityUpdateYear1EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 			ExtendedSecurityUpdateYear2EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 			ExtendedSecurityUpdateYear3EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 			ExtendedSupportEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-07-14T00:00:00.000Z"); return t}()),
	// 			MainstreamEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-13T00:00:00.000Z"); return t}()),
	// 			ServicePackStatus: to.Ptr("Unknown"),
	// 			SupportStatus: to.Ptr("Extended"),
	// 		},
	// 		RecommendedAzureSQLTargetType: to.Ptr(armmigrationassessment.TargetTypeAzureSQLVirtualMachine),
	// 		RecommendedSuitability: to.Ptr(armmigrationassessment.RecommendedSuitabilitySuitableForSQLVM),
	// 		RecommendedTargetReasonings: []*armmigrationassessment.SQLRecommendationReasoning{
	// 			{
	// 				ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
	// 					{
	// 						ContextKey: to.Ptr("MigrationIssuesCount"),
	// 						ContextValue: to.Ptr("0"),
	// 					},
	// 					{
	// 						ContextKey: to.Ptr("MigrationWarningsCount"),
	// 						ContextValue: to.Ptr("1"),
	// 					},
	// 					{
	// 						ContextKey: to.Ptr("RecommendedCost"),
	// 						ContextValue: to.Ptr("36.06432"),
	// 				}},
	// 				ReasoningCategory: to.Ptr("RecommendedTargetReasoning"),
	// 				ReasoningID: to.Ptr("RecommendedWithWarningsForSQLVM"),
	// 			},
	// 			{
	// 				ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
	// 					{
	// 						ContextKey: to.Ptr("RecommendedCost"),
	// 						ContextValue: to.Ptr("36.06432"),
	// 				}},
	// 				ReasoningCategory: to.Ptr("RecommendedTargetReasoning"),
	// 				ReasoningID: to.Ptr("RecommendedCost"),
	// 		}},
	// 		SizingCriterion: to.Ptr(armmigrationassessment.AssessmentSizingCriterionAsOnPremises),
	// 		SQLEdition: to.Ptr("Enterprise Edition: Core-based Licensing (64-bit)"),
	// 		SQLInstanceSdsArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourcegroups/bansalankit-rg/providers/microsoft.offazure/mastersites/fci-ankit-test6065mastersite/sqlsites/fci-ankit-test6065sqlsites/sqlservers/13c17b85-8842-4d50-a640-df0e45e56670"),
	// 		SQLVersion: to.Ptr("SQL Server 2016"),
	// 		StorageTypeBasedDetails: []*armmigrationassessment.AssessedSQLInstanceStorageDetails{
	// 			{
	// 				DiskSizeInMB: to.Ptr[float32](112),
	// 				MegabytesPerSecondOfRead: to.Ptr[float32](0),
	// 				MegabytesPerSecondOfWrite: to.Ptr[float32](0),
	// 				NumberOfReadOperationsPerSecond: to.Ptr[float32](0),
	// 				NumberOfWriteOperationsPerSecond: to.Ptr[float32](0),
	// 				StorageType: to.Ptr("Rows"),
	// 			},
	// 			{
	// 				DiskSizeInMB: to.Ptr[float32](112),
	// 				MegabytesPerSecondOfRead: to.Ptr[float32](0),
	// 				MegabytesPerSecondOfWrite: to.Ptr[float32](0),
	// 				NumberOfReadOperationsPerSecond: to.Ptr[float32](0),
	// 				NumberOfWriteOperationsPerSecond: to.Ptr[float32](0),
	// 				StorageType: to.Ptr("Log"),
	// 		}},
	// 		UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-08T05:22:35.966Z"); return t}()),
	// 	},
	// }
}
