package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AssessedSqlDatabaseV2Operations_ListBySqlAssessmentV2_MaximumSet_Gen.json
func ExampleAssessedSQLDatabaseV2OperationsClient_NewListBySQLAssessmentV2Pager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssessedSQLDatabaseV2OperationsClient().NewListBySQLAssessmentV2Pager("rgmigrate", "fci-test6904project", "test_fci_hadr", "test_swagger_1", &armmigrationassessment.AssessedSQLDatabaseV2OperationsClientListBySQLAssessmentV2Options{Filter: to.Ptr("(contains(Properties/DatabaseName,'adv130'))"),
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
		// page.AssessedSQLDatabaseV2ListResult = armmigrationassessment.AssessedSQLDatabaseV2ListResult{
		// 	Value: []*armmigrationassessment.AssessedSQLDatabaseV2{
		// 		{
		// 			Name: to.Ptr("858eb860-9e07-417c-91b6-bca1bffb3bf5"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/sqlAssessments/assessedSqlDatabases"),
		// 			ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.Migrate/assessmentprojects/fci-ankit-test6904project/groups/test_fci_hadr/sqlAssessments/test_swagger_1/assessedSqlDatabases/858eb860-9e07-417c-91b6-bca1bffb3bf5"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 			},
		// 			Properties: &armmigrationassessment.AssessedSQLDatabaseV2Properties{
		// 				AssessedSQLInstanceArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.Migrate/assessmentprojects/fci-ankit-test6904project/groups/test_fci_hadr/sqlAssessments/test_swagger_1/assessedSqlInstances/b11d90cc-b528-49e4-aac2-1c9a53209f8e"),
		// 				AzureSQLDBSuitabilityDetails: &armmigrationassessment.SQLAssessmentV2PaasSuitabilityData{
		// 					AzureSQLSKU: &armmigrationassessment.AzureSQLPaasSKUDTO{
		// 						AzureSQLComputeTier: to.Ptr(armmigrationassessment.ComputeTierProvisioned),
		// 						AzureSQLHardwareGeneration: to.Ptr(armmigrationassessment.HardwareGenerationGen5),
		// 						AzureSQLServiceTier: to.Ptr(armmigrationassessment.AzureSQLServiceTierGeneralPurpose),
		// 						AzureSQLTargetType: to.Ptr(armmigrationassessment.TargetTypeAzureSQLDatabase),
		// 						Cores: to.Ptr[int32](2),
		// 						PredictedDataSizeInMB: to.Ptr[float32](10),
		// 						PredictedLogSizeInMB: to.Ptr[float32](5),
		// 						StorageMaxSizeInMB: to.Ptr[float32](1024),
		// 					},
		// 					CostComponents: []*armmigrationassessment.CostComponent{
		// 						{
		// 							Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlySecurityCost),
		// 							Value: to.Ptr[float32](0),
		// 					}},
		// 					MigrationGuidelines: []*armmigrationassessment.SQLMigrationGuideline{
		// 					},
		// 					MigrationIssues: []*armmigrationassessment.SQLAssessmentMigrationIssue{
		// 						{
		// 							ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
		// 							},
		// 							IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryInternal),
		// 							IssueID: to.Ptr("SuitabilityReportMissing"),
		// 						},
		// 						{
		// 							ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
		// 							},
		// 							IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryInternal),
		// 							IssueID: to.Ptr("PerformanceDataMissing"),
		// 						},
		// 						{
		// 							ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
		// 								{
		// 									ObjectName: to.Ptr("Finance_User10"),
		// 									ObjectType: to.Ptr("SqlDatabase"),
		// 							}},
		// 							IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryWarning),
		// 							IssueID: to.Ptr("PerfBasedFallbackToAsOnPremises"),
		// 					}},
		// 					MigrationTargetPlatform: to.Ptr(armmigrationassessment.TargetTypeAzureSQLDatabase),
		// 					MonthlyComputeCost: to.Ptr[float32](118.999824),
		// 					MonthlyStorageCost: to.Ptr[float32](0.69),
		// 					RecommendationReasonings: []*armmigrationassessment.SQLRecommendationReasoning{
		// 						{
		// 							ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 								{
		// 									ContextKey: to.Ptr("AllocatedCores"),
		// 									ContextValue: to.Ptr("2"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("TargetCoresRecommended"),
		// 									ContextValue: to.Ptr("2"),
		// 							}},
		// 							ReasoningCategory: to.Ptr("CpuType"),
		// 							ReasoningID: to.Ptr("DatabaseCpuScalingFactorRequirementUnavailable"),
		// 						},
		// 						{
		// 							ContextParameters: []*armmigrationassessment.SQLRecommendationReasoningContext{
		// 								{
		// 									ContextKey: to.Ptr("TotalStorageConsumedByDataFilesInGB"),
		// 									ContextValue: to.Ptr("0.01"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("TotalStorageConsumedByLogFilesInGB"),
		// 									ContextValue: to.Ptr("0.005"),
		// 								},
		// 								{
		// 									ContextKey: to.Ptr("TotalRecommendedStorageInGB"),
		// 									ContextValue: to.Ptr("1024"),
		// 							}},
		// 							ReasoningCategory: to.Ptr("StorageType"),
		// 							ReasoningID: to.Ptr("MeetsDatabaseStorageRequirement"),
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
		// 					SecuritySuitability: to.Ptr(armmigrationassessment.CloudSuitabilityUnknown),
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
		// 					Suitability: to.Ptr(armmigrationassessment.CloudSuitabilityReadinessUnknown),
		// 				},
		// 				AzureSQLMISuitabilityDetails: &armmigrationassessment.SQLAssessmentV2PaasSuitabilityData{
		// 					CostComponents: []*armmigrationassessment.CostComponent{
		// 						{
		// 							Name: to.Ptr(armmigrationassessment.CostComponentNameMonthlySecurityCost),
		// 							Value: to.Ptr[float32](0),
		// 					}},
		// 					MigrationGuidelines: []*armmigrationassessment.SQLMigrationGuideline{
		// 					},
		// 					MigrationIssues: []*armmigrationassessment.SQLAssessmentMigrationIssue{
		// 						{
		// 							ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
		// 							},
		// 							IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryInternal),
		// 							IssueID: to.Ptr("SuitabilityReportMissing"),
		// 						},
		// 						{
		// 							ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
		// 							},
		// 							IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryInternal),
		// 							IssueID: to.Ptr("PerformanceDataMissing"),
		// 						},
		// 						{
		// 							ImpactedObjects: []*armmigrationassessment.ImpactedAssessmentObject{
		// 								{
		// 									ObjectName: to.Ptr("Finance_User10"),
		// 									ObjectType: to.Ptr("SqlDatabase"),
		// 							}},
		// 							IssueCategory: to.Ptr(armmigrationassessment.SQLAssessmentMigrationIssueCategoryIssue),
		// 							IssueID: to.Ptr("SkuNotFound"),
		// 					}},
		// 					MigrationTargetPlatform: to.Ptr(armmigrationassessment.TargetTypeAzureSQLManagedInstance),
		// 					MonthlyComputeCost: to.Ptr[float32](0),
		// 					MonthlyStorageCost: to.Ptr[float32](0),
		// 					RecommendationReasonings: []*armmigrationassessment.SQLRecommendationReasoning{
		// 					},
		// 					ReplicaAzureSQLSKU: []*armmigrationassessment.AzureSQLPaasSKUDTO{
		// 					},
		// 					SecuritySuitability: to.Ptr(armmigrationassessment.CloudSuitabilityUnknown),
		// 					ShouldProvisionReplicas: to.Ptr(false),
		// 					SKUReplicationMode: to.Ptr(armmigrationassessment.SKUReplicationModeNotApplicable),
		// 					Suitability: to.Ptr(armmigrationassessment.CloudSuitabilityNotSuitable),
		// 				},
		// 				BufferCacheSizeInMB: to.Ptr[float32](18.87557603686636),
		// 				CompatibilityLevel: to.Ptr(armmigrationassessment.CompatibilityLevelCompatLevel120),
		// 				ConfidenceRatingInPercentage: to.Ptr[float32](0),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-08T05:22:34.857Z"); return t}()),
		// 				DatabaseName: to.Ptr("Finance_User10"),
		// 				DatabaseSizeInMB: to.Ptr[float32](15),
		// 				InstanceName: to.Ptr("MSSQLSERVER"),
		// 				IsDatabaseHighlyAvailable: to.Ptr(false),
		// 				MachineArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.Migrate/assessmentprojects/fci-ankit-test6904project/machines/cc64c9dc-b38e-435d-85ad-d509df5d92c6"),
		// 				MachineName: to.Ptr("SQLHAVM17"),
		// 				MegabytesPerSecondOfRead: to.Ptr[float32](0),
		// 				MegabytesPerSecondOfWrite: to.Ptr[float32](0),
		// 				NumberOfReadOperationsPerSecond: to.Ptr[float32](0),
		// 				NumberOfWriteOperationsPerSecond: to.Ptr[float32](0),
		// 				PercentageCoresUtilization: to.Ptr[float32](100),
		// 				ProductSupportStatus: &armmigrationassessment.ProductSupportStatus{
		// 					CurrentEsuYear: to.Ptr("Unknown"),
		// 					EsuStatus: to.Ptr("Unknown"),
		// 					Eta: to.Ptr[int32](9),
		// 					ExtendedSecurityUpdateYear1EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 					ExtendedSecurityUpdateYear2EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 					ExtendedSecurityUpdateYear3EndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 					ExtendedSupportEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-09T00:00:00.000Z"); return t}()),
		// 					MainstreamEndDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-07-09T00:00:00.000Z"); return t}()),
		// 					ServicePackStatus: to.Ptr("Unknown"),
		// 					SupportStatus: to.Ptr("Extended"),
		// 				},
		// 				RecommendedAzureSQLTargetType: to.Ptr(armmigrationassessment.TargetTypeAzureSQLVirtualMachine),
		// 				RecommendedSuitability: to.Ptr(armmigrationassessment.RecommendedSuitabilitySuitableForSQLVM),
		// 				SizingCriterion: to.Ptr(armmigrationassessment.AssessmentSizingCriterionAsOnPremises),
		// 				SQLDatabaseSdsArmID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourcegroups/bansalankit-rg/providers/microsoft.offazure/mastersites/fci-ankit-test6065mastersite/sqlsites/fci-ankit-test6065sqlsites/sqldatabases/ea92ba27-3656-5370-8bcc-e5eed9d7ba5e"),
		// 				UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-08T05:22:36.466Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
