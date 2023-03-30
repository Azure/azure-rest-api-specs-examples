package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/PredictionsListByHub.json
func ExamplePredictionsClient_NewListByHubPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPredictionsClient().NewListByHubPager("TestHubRG", "sdkTestHub", nil)
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
		// page.PredictionListResult = armcustomerinsights.PredictionListResult{
		// 	Value: []*armcustomerinsights.PredictionResourceFormat{
		// 		{
		// 			Name: to.Ptr("sdkTestHub/sdktest"),
		// 			Type: to.Ptr("Microsoft.CustomerInsights/hubs/predictions"),
		// 			ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/azSdkTestHub/predictions/sdktest"),
		// 			Properties: &armcustomerinsights.Prediction{
		// 				Description: map[string]*string{
		// 					"en-us": to.Ptr("sdktest"),
		// 				},
		// 				AutoAnalyze: to.Ptr(true),
		// 				DisplayName: map[string]*string{
		// 					"en-us": to.Ptr("sdktest"),
		// 				},
		// 				Grades: []*armcustomerinsights.PredictionGradesItem{
		// 				},
		// 				InvolvedInteractionTypes: []*string{
		// 				},
		// 				InvolvedKpiTypes: []*string{
		// 				},
		// 				InvolvedRelationships: []*string{
		// 				},
		// 				Mappings: &armcustomerinsights.PredictionMappings{
		// 					Grade: to.Ptr("sdktest_Grade"),
		// 					Reason: to.Ptr("sdktest_Reason"),
		// 					Score: to.Ptr("sdktest_Score"),
		// 				},
		// 				NegativeOutcomeExpression: to.Ptr("Customers.FirstName = 'Mike'"),
		// 				PositiveOutcomeExpression: to.Ptr("Customers.FirstName = 'David'"),
		// 				PredictionName: to.Ptr("sdktest"),
		// 				PrimaryProfileType: to.Ptr("Customers"),
		// 				ProvisioningState: to.Ptr(armcustomerinsights.ProvisioningStatesSucceeded),
		// 				ScopeExpression: to.Ptr("*"),
		// 				ScoreLabel: to.Ptr("score label"),
		// 				SystemGeneratedEntities: &armcustomerinsights.PredictionSystemGeneratedEntities{
		// 					GeneratedInteractionTypes: []*string{
		// 						to.Ptr("_predictions_sdktest")},
		// 						GeneratedKpis: map[string]*string{
		// 							"leadingProfiles": to.Ptr("_predictions_sdktest_LeadingProfiles"),
		// 							"negativeOutcomeByGrade": to.Ptr("_predictions_sdktest_NegativeOutcomeByGrade"),
		// 							"positiveOutcomeByGrade": to.Ptr("_predictions_sdktest_PositiveOutcomeByGrade"),
		// 							"predictionPerformance": to.Ptr("_predictions_sdktest_PredictionPerformance"),
		// 						},
		// 						GeneratedLinks: []*string{
		// 							to.Ptr("_predictions_link_sdktest")},
		// 						},
		// 						TenantID: to.Ptr("predtest620"),
		// 					},
		// 			}},
		// 		}
	}
}
