package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/WebAppAssessmentV2SummaryOperations_ListByWebAppAssessmentV2_MaximumSet_Gen.json
func ExampleWebAppAssessmentV2SummaryOperationsClient_NewListByWebAppAssessmentV2Pager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebAppAssessmentV2SummaryOperationsClient().NewListByWebAppAssessmentV2Pager("rgopenapi", "sumukk-ccy-bcs4557project", "anraghun-selfhost-v2", "anraghun-selfhost-v2", nil)
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
		// page.WebAppAssessmentV2SummaryListResult = armmigrationassessment.WebAppAssessmentV2SummaryListResult{
		// 	Value: []*armmigrationassessment.WebAppAssessmentV2Summary{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentProjects/groups/webAppAssessments"),
		// 			ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sumukk-ccy-bcs/providers/Microsoft.Migrate/assessmentProjects/sumukk-ccy-bcs4557project/groups/anraghun-selfhost-v2/webAppAssessments/anraghun-selfhost-v2/summaries/default"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:41:40.597Z"); return t}()),
		// 				CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-03T05:41:40.597Z"); return t}()),
		// 				LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmigrationassessment.WebAppAssessmentV2SummaryProperties{
		// 				AssessmentSummary: map[string]*armmigrationassessment.DiscoveredEntitiesSummary{
		// 					"discoveredEntitiesSummary": &armmigrationassessment.DiscoveredEntitiesSummary{
		// 						NumberOfMachines: to.Ptr[int32](24),
		// 						NumberOfServers: to.Ptr[int32](15),
		// 						NumberOfWebApps: to.Ptr[int32](6),
		// 						WebAppSummary: map[string]*int32{
		// 							"iis": to.Ptr[int32](30),
		// 						},
		// 						WebServerSummary: map[string]*int32{
		// 							"iis": to.Ptr[int32](13),
		// 						},
		// 					},
		// 				},
		// 				TargetSpecificSummary: map[string]*armmigrationassessment.TargetSpecificSummary{
		// 					"azureAppService": &armmigrationassessment.TargetSpecificSummary{
		// 						ReadinessSummary: map[string]*int32{
		// 							"conditionallySuitable": to.Ptr[int32](13),
		// 						},
		// 						RecommendationResultSKUDetails: map[string]*armmigrationassessment.WebAppSKUDetails{
		// 							"isolated": &armmigrationassessment.WebAppSKUDetails{
		// 								AppServicePlanCount: to.Ptr[int32](2),
		// 								MonthlySecurityCost: to.Ptr[float32](0),
		// 								SKUName: to.Ptr("Isolated"),
		// 								SKUSize: to.Ptr("I1"),
		// 								TotalMonthlyCost: to.Ptr[float32](584),
		// 							},
		// 						},
		// 					},
		// 				},
		// 			},
		// 	}},
		// }
	}
}
