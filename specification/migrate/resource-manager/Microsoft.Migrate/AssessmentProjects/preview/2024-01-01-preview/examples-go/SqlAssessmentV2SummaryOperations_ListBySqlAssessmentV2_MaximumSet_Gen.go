package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/SqlAssessmentV2SummaryOperations_ListBySqlAssessmentV2_MaximumSet_Gen.json
func ExampleSQLAssessmentV2SummaryOperationsClient_NewListBySQLAssessmentV2Pager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSQLAssessmentV2SummaryOperationsClient().NewListBySQLAssessmentV2Pager("rgmigrate", "fci-test6904project", "test_fci_hadr", "test_swagger_1", nil)
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
		// page.SQLAssessmentV2SummaryListResult = armmigrationassessment.SQLAssessmentV2SummaryListResult{
		// 	Value: []*armmigrationassessment.SQLAssessmentV2Summary{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/groups/sqlAssessments/summaries"),
		// 			ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/bansalankit-rg/providers/Microsoft.Migrate/assessmentprojects/fci-ankit-test6904project/groups/test_fci_hadr/sqlAssessments/test_swagger_1/summaries/default"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0-12-31T15:54:17.000Z"); return t}()),
		// 			},
		// 			Properties: &armmigrationassessment.SQLAssessmentV2SummaryProperties{
		// 				AssessmentSummary: map[string]*armmigrationassessment.SQLAssessmentV2SummaryData{
		// 					"azureSqlDatabase": &armmigrationassessment.SQLAssessmentV2SummaryData{
		// 						ConfidenceScore: to.Ptr[float32](0),
		// 						MonthlyComputeCost: to.Ptr[float32](175079.0739839982),
		// 						MonthlyLicenseCost: to.Ptr[float32](214401.07883999476),
		// 						MonthlySecurityCost: to.Ptr[float32](0),
		// 						MonthlyStorageCost: to.Ptr[float32](1054.8825000000395),
		// 						SuitabilitySummary: map[string]*int32{
		// 							"readinessUnknown": to.Ptr[int32](1468),
		// 							"suitableForSqlDB": to.Ptr[int32](16),
		// 						},
		// 					},
		// 					"azureSqlManagedInstance": &armmigrationassessment.SQLAssessmentV2SummaryData{
		// 						ConfidenceScore: to.Ptr[float32](0),
		// 						MonthlyComputeCost: to.Ptr[float32](4581.992448),
		// 						MonthlyLicenseCost: to.Ptr[float32](6568.510799999999),
		// 						MonthlySecurityCost: to.Ptr[float32](0),
		// 						MonthlyStorageCost: to.Ptr[float32](0),
		// 						SuitabilitySummary: map[string]*int32{
		// 							"notSuitable": to.Ptr[int32](7),
		// 							"readinessUnknown": to.Ptr[int32](15),
		// 							"suitableForSqlMI": to.Ptr[int32](1),
		// 						},
		// 					},
		// 					"azureSqlVirtualMachine": &armmigrationassessment.SQLAssessmentV2SummaryData{
		// 						ConfidenceScore: to.Ptr[float32](0),
		// 						MonthlyComputeCost: to.Ptr[float32](889.4966399999996),
		// 						MonthlyLicenseCost: to.Ptr[float32](25281.36),
		// 						MonthlySecurityCost: to.Ptr[float32](0),
		// 						MonthlyStorageCost: to.Ptr[float32](73.94399999999999),
		// 						SuitabilitySummary: map[string]*int32{
		// 							"readinessUnknown": to.Ptr[int32](2),
		// 							"suitableForSqlVM": to.Ptr[int32](21),
		// 						},
		// 					},
		// 					"azureVirtualMachine": &armmigrationassessment.SQLAssessmentV2SummaryData{
		// 						ConfidenceScore: to.Ptr[float32](0),
		// 						MonthlyComputeCost: to.Ptr[float32](625.9420799999998),
		// 						MonthlyLicenseCost: to.Ptr[float32](1300.5120000000002),
		// 						MonthlySecurityCost: to.Ptr[float32](0),
		// 						MonthlyStorageCost: to.Ptr[float32](374.48999999999995),
		// 						SuitabilitySummary: map[string]*int32{
		// 							"suitableForVM": to.Ptr[int32](19),
		// 						},
		// 					},
		// 					"recommended": &armmigrationassessment.SQLAssessmentV2SummaryData{
		// 						ConfidenceScore: to.Ptr[float32](0),
		// 						MonthlyComputeCost: to.Ptr[float32](823.6079999999997),
		// 						MonthlyLicenseCost: to.Ptr[float32](22912.896),
		// 						MonthlySecurityCost: to.Ptr[float32](0),
		// 						MonthlyStorageCost: to.Ptr[float32](100.88400000000001),
		// 						SuitabilitySummary: map[string]*int32{
		// 							"suitableForSqlVM": to.Ptr[int32](19),
		// 							"suitableForVM": to.Ptr[int32](4),
		// 						},
		// 					},
		// 				},
		// 				DatabaseDistributionBySizingCriterion: map[string]*int32{
		// 					"asOnPremises": to.Ptr[int32](0),
		// 					"perfBasedFallbackToAsOnPremises": to.Ptr[int32](1484),
		// 					"performanceBased": to.Ptr[int32](0),
		// 				},
		// 				DistributionByServicePackInsight: map[string]*int32{
		// 					"unknown": to.Ptr[int32](23),
		// 				},
		// 				DistributionBySQLEdition: map[string]*int32{
		// 					"enterprise Edition (64-bit)": to.Ptr[int32](1),
		// 					"enterprise Edition: Core-based Licensing": to.Ptr[int32](1),
		// 					"enterprise Edition: Core-based Licensing (64-bit)": to.Ptr[int32](15),
		// 					"enterprise Evaluation Edition (64-bit)": to.Ptr[int32](1),
		// 					"express Edition (64-bit)": to.Ptr[int32](2),
		// 					"standard Edition (64-bit)": to.Ptr[int32](2),
		// 					"web Edition": to.Ptr[int32](1),
		// 				},
		// 				DistributionBySQLVersion: map[string]*int32{
		// 					"sql Server 2012": to.Ptr[int32](9),
		// 					"sql Server 2014": to.Ptr[int32](4),
		// 					"sql Server 2016": to.Ptr[int32](5),
		// 					"sql Server 2017": to.Ptr[int32](2),
		// 					"sql Server 2019": to.Ptr[int32](3),
		// 				},
		// 				DistributionBySupportStatus: map[string]*int32{
		// 					"extended": to.Ptr[int32](11),
		// 					"mainstream": to.Ptr[int32](3),
		// 					"outOfSupport": to.Ptr[int32](9),
		// 				},
		// 				InstanceDistributionBySizingCriterion: map[string]*int32{
		// 					"asOnPremises": to.Ptr[int32](0),
		// 					"perfBasedFallbackToAsOnPremises": to.Ptr[int32](23),
		// 					"performanceBased": to.Ptr[int32](0),
		// 				},
		// 				NumberOfFciInstances: to.Ptr[int32](4),
		// 				NumberOfMachines: to.Ptr[int32](19),
		// 				NumberOfSQLAvailabilityGroups: to.Ptr[int32](5),
		// 				NumberOfSQLDatabases: to.Ptr[int32](1484),
		// 				NumberOfSQLInstances: to.Ptr[int32](23),
		// 			},
		// 	}},
		// }
	}
}
