package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AssessedSqlRecommendedEntityOperations_Get_MaximumSet_Gen.json
func ExampleAssessedSQLRecommendedEntityOperationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssessedSQLRecommendedEntityOperationsClient().Get(ctx, "rgmigrate", "fci-test6904project", "test_fci_hadr", "test_swagger_1", "cc64c9dc-b38e-435d-85ad-d509df5d92c6", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AssessedSQLRecommendedEntity = armmigrationassessment.AssessedSQLRecommendedEntity{
	// 	Name: to.Ptr("b11d90cc-b528-49e4-aac2-1c9a53209f8e"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/sqlAssessments/recommendedAssessedEntities"),
	// 	ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.Migrate/assessmentprojects/fci-ankit-test6904project/groups/test_fci_hadr/sqlAssessments/test_swagger_1/recommendedAssessedEntities/b11d90cc-b528-49e4-aac2-1c9a53209f8e"),
	// 	SystemData: &armmigrationassessment.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 	},
	// 	Properties: &armmigrationassessment.AssessedSQLRecommendedEntityProperties{
	// 		AssessedSQLEntityArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.Migrate/assessmentprojects/fci-ankit-test6904project/groups/test_fci_hadr/sqlAssessments/test_swagger_1/assessedSqlInstances/b11d90cc-b528-49e4-aac2-1c9a53209f8e"),
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
	// 			MonthlyComputeCost: to.Ptr[float32](25822.9618079999),
	// 			MonthlyStorageCost: to.Ptr[float32](149.73),
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
	// 					IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryIssue),
	// 					IssueID: to.Ptr("SkuNotFound"),
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
	// 			MonthlyComputeCost: to.Ptr[float32](0),
	// 			MonthlyStorageCost: to.Ptr[float32](0),
	// 			RecommendationReasonings: []*armmigrationassessment.SQLRecommendationReasoning{
	// 			},
	// 			ReplicaAzureSQLSKU: []*armmigrationassessment.AzureSQLPaasSKUDTO{
	// 			},
	// 			SecuritySuitability: to.Ptr(armmigrationassessment.CloudSuitabilityUnknown),
	// 			ShouldProvisionReplicas: to.Ptr(false),
	// 			SKUReplicationMode: to.Ptr(armmigrationassessment.SKUReplicationModeNotApplicable),
	// 			Suitability: to.Ptr(armmigrationassessment.CloudSuitabilityNotSuitable),
	// 		},
	// 		AzureSQLVMSuitabilityDetails: &armmigrationassessment.SQLAssessmentV2IaasSuitabilityData{
	// 			AzureSQLSKU: &armmigrationassessment.AzureSQLIaasSKUDTO{
	// 				AzureSQLTargetType: to.Ptr(armmigrationassessment.TargetTypeAzureSQLVirtualMachine),
	// 				DataDiskSizes: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 				},
	// 				LogDiskSizes: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 				},
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
	// 					GuidelineID: to.Ptr("ProvisionInstanceAsFciOnSqlVM"),
	// 					MigrationGuidelineCategory: to.Ptr(armmigrationassessment.SQLMigrationGuidelineCategoryFailoverCluterInstanceGuideLine),
	// 					MigrationGuidelineContext: []*armmigrationassessment.MigrationGuidelineContext{
	// 					},
	// 				},
	// 				{
	// 					GuidelineID: to.Ptr("ProvisionDisksBasedOnLayoutInSku"),
	// 					MigrationGuidelineCategory: to.Ptr(armmigrationassessment.SQLMigrationGuidelineCategoryFailoverCluterInstanceGuideLine),
	// 					MigrationGuidelineContext: []*armmigrationassessment.MigrationGuidelineContext{
	// 					},
	// 				},
	// 				{
	// 					GuidelineID: to.Ptr("ProvisionSharedDiskForQuorum"),
	// 					MigrationGuidelineCategory: to.Ptr(armmigrationassessment.SQLMigrationGuidelineCategoryFailoverCluterInstanceGuideLine),
	// 					MigrationGuidelineContext: []*armmigrationassessment.MigrationGuidelineContext{
	// 					},
	// 				},
	// 				{
	// 					GuidelineID: to.Ptr("MigrateNonReplicatedDatabasesAndAutoSeedSecondaryIfNeeded"),
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
	// 			MonthlyComputeCost: to.Ptr[float32](65.88864),
	// 			MonthlyStorageCost: to.Ptr[float32](3.276),
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
	// 					},
	// 					ReasoningCategory: to.Ptr("AlwaysOnFCI"),
	// 					ReasoningID: to.Ptr("AlwaysOnVirtualMachineFCIDiskWitness"),
	// 				},
	// 				{
	// 					ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
	// 					},
	// 					ReasoningCategory: to.Ptr("AlwaysOnFCI"),
	// 					ReasoningID: to.Ptr("AlwaysOnVirtualMachineFCISharedZRSDisks"),
	// 				},
	// 				{
	// 					ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
	// 						{
	// 							ContextKey: to.Ptr("CommaSeperatedManagedDisks"),
	// 							ContextValue: to.Ptr(""),
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
	// 				{
	// 					AzureSQLTargetType: to.Ptr(armmigrationassessment.TargetTypeAzureSQLVirtualMachine),
	// 					DataDiskSizes: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 					},
	// 					LogDiskSizes: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 					},
	// 					VirtualMachineSize: &armmigrationassessment.AzureVirtualMachineSKUDTO{
	// 						AvailableCores: to.Ptr[int32](2),
	// 						AzureSKUName: to.Ptr(armmigrationassessment.AzureVMSizeStandardD2AsV4),
	// 						AzureVMFamily: to.Ptr(armmigrationassessment.AzureVMFamilyDasv4Series),
	// 						Cores: to.Ptr[int32](2),
	// 						MaxNetworkInterfaces: to.Ptr[int32](0),
	// 					},
	// 			}},
	// 			SecuritySuitability: to.Ptr(armmigrationassessment.CloudSuitabilityUnknown),
	// 			SharedResources: &armmigrationassessment.SharedResourcesDTO{
	// 				NumberOfMounts: to.Ptr[int32](2),
	// 				QuorumWitness: &armmigrationassessment.AzureQuorumWitnessDTO{
	// 					QuorumWitnessType: to.Ptr(armmigrationassessment.AzureQuorumWitnessDTOQuorumWitnessTypeDisk),
	// 				},
	// 				SharedDataDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 					{
	// 						DiskRedundancy: to.Ptr(armmigrationassessment.AzureManagedDiskSKUDTODiskRedundancyLRS),
	// 						DiskSize: to.Ptr(armmigrationassessment.AzureDiskSizePremiumP2),
	// 						DiskType: to.Ptr(armmigrationassessment.AzureManagedDiskSKUDTODiskTypePremium),
	// 						RecommendedIops: to.Ptr[float32](120),
	// 						RecommendedSizeInGib: to.Ptr[float32](8),
	// 						RecommendedThroughputInMbps: to.Ptr[float32](25),
	// 						StorageCost: to.Ptr[float32](1.6380000000000001),
	// 				}},
	// 				SharedLogDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 					{
	// 						DiskRedundancy: to.Ptr(armmigrationassessment.AzureManagedDiskSKUDTODiskRedundancyLRS),
	// 						DiskSize: to.Ptr(armmigrationassessment.AzureDiskSizePremiumP2),
	// 						DiskType: to.Ptr(armmigrationassessment.AzureManagedDiskSKUDTODiskTypePremium),
	// 						RecommendedIops: to.Ptr[float32](120),
	// 						RecommendedSizeInGib: to.Ptr[float32](8),
	// 						RecommendedThroughputInMbps: to.Ptr[float32](25),
	// 						StorageCost: to.Ptr[float32](1.6380000000000001),
	// 				}},
	// 				SharedTempDbDisks: []*armmigrationassessment.AzureManagedDiskSKUDTO{
	// 				},
	// 			},
	// 			ShouldProvisionReplicas: to.Ptr(true),
	// 			SKUReplicationMode: to.Ptr(armmigrationassessment.SKUReplicationModeNotApplicable),
	// 			Suitability: to.Ptr(armmigrationassessment.CloudSuitabilitySuitable),
	// 		},
	// 		DbCount: to.Ptr[int32](217),
	// 		DiscoveredDBCount: to.Ptr[int32](217),
	// 		HasScanOccurred: to.Ptr(true),
	// 		InstanceName: to.Ptr("MSSQLSERVER"),
	// 		IsClustered: to.Ptr(true),
	// 		IsHighAvailabilityEnabled: to.Ptr(true),
	// 		MachineName: to.Ptr("FC7-SQL2K14.FPL.COM"),
	// 		ProductSupportStatus: &armmigrationassessment.ProductSupportStatus{
	// 			CurrentEsuYear: to.Ptr("Unknown"),
	// 			EsuStatus: to.Ptr("Unknown"),
	// 			Eta: to.Ptr[int32](9),
	// 			ExtendedSecurityUpdateYear1EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 			ExtendedSecurityUpdateYear2EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 			ExtendedSecurityUpdateYear3EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
	// 			ExtendedSupportEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-09T00:00:00.000Z"); return t}()),
	// 			MainstreamEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-09T00:00:00.000Z"); return t}()),
	// 			ServicePackStatus: to.Ptr("Unknown"),
	// 			SupportStatus: to.Ptr("Extended"),
	// 		},
	// 		RecommendedAzureSQLTargetType: to.Ptr(armmigrationassessment.TargetTypeAzureSQLVirtualMachine),
	// 		RecommendedSuitability: to.Ptr(armmigrationassessment.RecommendedSuitabilitySuitableForSQLVM),
	// 		SizingCriterion: to.Ptr(armmigrationassessment.AssessmentSizingCriterionAsOnPremises),
	// 		SQLEdition: to.Ptr("Enterprise Edition: Core-based Licensing (64-bit)"),
	// 		SQLVersion: to.Ptr("SQL Server 2014"),
	// 	},
	// }
}
