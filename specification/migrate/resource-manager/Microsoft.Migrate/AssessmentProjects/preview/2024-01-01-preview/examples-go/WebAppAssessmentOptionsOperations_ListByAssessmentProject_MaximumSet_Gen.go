package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/WebAppAssessmentOptionsOperations_ListByAssessmentProject_MaximumSet_Gen.json
func ExampleWebAppAssessmentOptionsOperationsClient_NewListByAssessmentProjectPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWebAppAssessmentOptionsOperationsClient().NewListByAssessmentProjectPager("rgopenapi", "sumukk-ccy-bcs", nil)
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
		// page.WebAppAssessmentOptionsListResult = armmigrationassessment.WebAppAssessmentOptionsListResult{
		// 	Value: []*armmigrationassessment.WebAppAssessmentOptions{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/webAppAssessmentOptions"),
		// 			ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/sumukk-ccy-bcs/providers/Microsoft.Migrate/assessmentprojects/sumukk-ccy-bcs4557project/webAppAssessmentOptions/default"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-06T05:40:17.507Z"); return t}()),
		// 				CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-11-06T05:40:17.507Z"); return t}()),
		// 				LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmigrationassessment.WebAppAssessmentOptionsProperties{
		// 				ReservedInstanceSupportedCurrencies: []*armmigrationassessment.AzureCurrency{
		// 					to.Ptr(armmigrationassessment.AzureCurrencyUSD)},
		// 					ReservedInstanceSupportedLocations: []*armmigrationassessment.AzureLocation{
		// 						to.Ptr(armmigrationassessment.AzureLocationAustraliaSoutheast)},
		// 						ReservedInstanceSupportedOffers: []*armmigrationassessment.AzureOfferCode{
		// 							to.Ptr(armmigrationassessment.AzureOfferCodeMsazr0003P)},
		// 							ReservedInstanceSupportedWebAppTiers: []*armmigrationassessment.AzureWebAppTier{
		// 								to.Ptr(armmigrationassessment.AzureWebAppTierPremiumV3)},
		// 								SavingsPlanSupportedLocations: []*armmigrationassessment.AzureLocation{
		// 									to.Ptr(armmigrationassessment.AzureLocationUnknown)},
		// 									SavingsPlanSupportedWebAppTiers: []*armmigrationassessment.AzureWebAppTier{
		// 										to.Ptr(armmigrationassessment.AzureWebAppTierPremiumV3)},
		// 										WebAppSKUs: []*armmigrationassessment.WebAppTargetOptions{
		// 											{
		// 												TargetLocations: []*armmigrationassessment.AzureLocation{
		// 													to.Ptr(armmigrationassessment.AzureLocation("PolandCentral"))},
		// 													WebAppTier: to.Ptr(armmigrationassessment.AzureWebAppTierPremiumV3),
		// 											}},
		// 										},
		// 								}},
		// 							}
	}
}
