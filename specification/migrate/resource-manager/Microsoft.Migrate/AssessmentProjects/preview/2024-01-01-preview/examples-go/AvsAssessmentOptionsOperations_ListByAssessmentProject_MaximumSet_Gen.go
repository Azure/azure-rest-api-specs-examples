package armmigrationassessment_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/dd668996bc8ba729784c02c7a99b6b0042612bb3/specification/migrate/resource-manager/Microsoft.Migrate/AssessmentProjects/preview/2024-01-01-preview/examples/AvsAssessmentOptionsOperations_ListByAssessmentProject_MaximumSet_Gen.json
func ExampleAvsAssessmentOptionsOperationsClient_NewListByAssessmentProjectPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrationassessment.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAvsAssessmentOptionsOperationsClient().NewListByAssessmentProjectPager("ayagrawrg", "app18700project", nil)
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
		// page.AvsAssessmentOptionsListResult = armmigrationassessment.AvsAssessmentOptionsListResult{
		// 	Value: []*armmigrationassessment.AvsAssessmentOptions{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/{assessmentOptionsName}"),
		// 			ID: to.Ptr("/subscriptions/4bd2aa0f-2bd2-4d67-91a8-5a4533d58600/resourceGroups/ayagrawrg/providers/Microsoft.Migrate/assessmentprojects/app18700project/avsAssessmentOptions/default"),
		// 			SystemData: &armmigrationassessment.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:22:18.589Z"); return t}()),
		// 				CreatedBy: to.Ptr("asd"),
		// 				CreatedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-09-15T07:22:18.589Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("tnmop"),
		// 				LastModifiedByType: to.Ptr(armmigrationassessment.CreatedByTypeUser),
		// 			},
		// 			Properties: &armmigrationassessment.AvsAssessmentOptionsProperties{
		// 				AvsExternalStorageTypes: []*armmigrationassessment.AvsExternalStorageOptions{
		// 					{
		// 						StorageType: to.Ptr(armmigrationassessment.ExternalStorageTypeAnfStandard),
		// 						TargetLocations: []*armmigrationassessment.AzureLocation{
		// 							to.Ptr(armmigrationassessment.AzureLocationEastAsia),
		// 							to.Ptr(armmigrationassessment.AzureLocationSoutheastAsia),
		// 							to.Ptr(armmigrationassessment.AzureLocationAustraliaEast),
		// 							to.Ptr(armmigrationassessment.AzureLocationAustraliaSoutheast),
		// 							to.Ptr(armmigrationassessment.AzureLocationBrazilSouth),
		// 							to.Ptr(armmigrationassessment.AzureLocationCanadaCentral),
		// 							to.Ptr(armmigrationassessment.AzureLocationCanadaEast),
		// 							to.Ptr(armmigrationassessment.AzureLocationWestEurope),
		// 							to.Ptr(armmigrationassessment.AzureLocationNorthEurope),
		// 							to.Ptr(armmigrationassessment.AzureLocationJapanEast),
		// 							to.Ptr(armmigrationassessment.AzureLocationJapanWest),
		// 							to.Ptr(armmigrationassessment.AzureLocationUkWest),
		// 							to.Ptr(armmigrationassessment.AzureLocationUkSouth),
		// 							to.Ptr(armmigrationassessment.AzureLocationNorthCentralUs),
		// 							to.Ptr(armmigrationassessment.AzureLocationEastUs),
		// 							to.Ptr(armmigrationassessment.AzureLocationSouthCentralUs),
		// 							to.Ptr(armmigrationassessment.AzureLocationCentralUs),
		// 							to.Ptr(armmigrationassessment.AzureLocationEastUs2),
		// 							to.Ptr(armmigrationassessment.AzureLocationWestUs),
		// 							to.Ptr(armmigrationassessment.AzureLocationFranceCentral),
		// 							to.Ptr(armmigrationassessment.AzureLocationSouthAfricaNorth),
		// 							to.Ptr(armmigrationassessment.AzureLocationGermanyWestCentral),
		// 							to.Ptr(armmigrationassessment.AzureLocationSwedenCentral)},
		// 						},
		// 						{
		// 							StorageType: to.Ptr(armmigrationassessment.ExternalStorageTypeAnfPremium),
		// 							TargetLocations: []*armmigrationassessment.AzureLocation{
		// 								to.Ptr(armmigrationassessment.AzureLocationWestUs),
		// 								to.Ptr(armmigrationassessment.AzureLocationFranceCentral),
		// 								to.Ptr(armmigrationassessment.AzureLocationSouthAfricaNorth),
		// 								to.Ptr(armmigrationassessment.AzureLocationGermanyWestCentral),
		// 								to.Ptr(armmigrationassessment.AzureLocationSwedenCentral)},
		// 						}},
		// 						AvsNodes: []*armmigrationassessment.AvsSKUOptions{
		// 							{
		// 								NodeType: to.Ptr(armmigrationassessment.AzureAvsNodeTypeAV36),
		// 								TargetLocations: []*armmigrationassessment.AzureLocation{
		// 									to.Ptr(armmigrationassessment.AzureLocationEastAsia),
		// 									to.Ptr(armmigrationassessment.AzureLocationSoutheastAsia),
		// 									to.Ptr(armmigrationassessment.AzureLocationAustraliaEast),
		// 									to.Ptr(armmigrationassessment.AzureLocationAustraliaSoutheast),
		// 									to.Ptr(armmigrationassessment.AzureLocationBrazilSouth),
		// 									to.Ptr(armmigrationassessment.AzureLocationCanadaCentral),
		// 									to.Ptr(armmigrationassessment.AzureLocationCanadaEast),
		// 									to.Ptr(armmigrationassessment.AzureLocationWestEurope),
		// 									to.Ptr(armmigrationassessment.AzureLocationNorthEurope),
		// 									to.Ptr(armmigrationassessment.AzureLocationJapanEast),
		// 									to.Ptr(armmigrationassessment.AzureLocationJapanWest),
		// 									to.Ptr(armmigrationassessment.AzureLocationUkWest),
		// 									to.Ptr(armmigrationassessment.AzureLocationUkSouth),
		// 									to.Ptr(armmigrationassessment.AzureLocationNorthCentralUs),
		// 									to.Ptr(armmigrationassessment.AzureLocationEastUs),
		// 									to.Ptr(armmigrationassessment.AzureLocationSouthCentralUs),
		// 									to.Ptr(armmigrationassessment.AzureLocationCentralUs),
		// 									to.Ptr(armmigrationassessment.AzureLocationEastUs2),
		// 									to.Ptr(armmigrationassessment.AzureLocationWestUs),
		// 									to.Ptr(armmigrationassessment.AzureLocationFranceCentral),
		// 									to.Ptr(armmigrationassessment.AzureLocationSouthAfricaNorth),
		// 									to.Ptr(armmigrationassessment.AzureLocationGermanyWestCentral),
		// 									to.Ptr(armmigrationassessment.AzureLocationSwedenCentral)},
		// 								},
		// 								{
		// 									NodeType: to.Ptr(armmigrationassessment.AzureAvsNodeTypeAV36P),
		// 									TargetLocations: []*armmigrationassessment.AzureLocation{
		// 										to.Ptr(armmigrationassessment.AzureLocationAustraliaEast),
		// 										to.Ptr(armmigrationassessment.AzureLocationCanadaCentral),
		// 										to.Ptr(armmigrationassessment.AzureLocationUkSouth),
		// 										to.Ptr(armmigrationassessment.AzureLocationNorthCentralUs),
		// 										to.Ptr(armmigrationassessment.AzureLocationEastUs),
		// 										to.Ptr(armmigrationassessment.AzureLocationWestUs2),
		// 										to.Ptr(armmigrationassessment.AzureLocationSouthCentralUs),
		// 										to.Ptr(armmigrationassessment.AzureLocationEastUs2),
		// 										to.Ptr(armmigrationassessment.AzureLocationWestUs),
		// 										to.Ptr(armmigrationassessment.AzureLocationQatarCentral),
		// 										to.Ptr(armmigrationassessment.AzureLocationWestEurope)},
		// 									},
		// 									{
		// 										NodeType: to.Ptr(armmigrationassessment.AzureAvsNodeTypeAV52),
		// 										TargetLocations: []*armmigrationassessment.AzureLocation{
		// 											to.Ptr(armmigrationassessment.AzureLocationWestEurope),
		// 											to.Ptr(armmigrationassessment.AzureLocationUkSouth),
		// 											to.Ptr(armmigrationassessment.AzureLocationEastUs2)},
		// 									}},
		// 									FailuresToTolerateAndRaidLevelValues: []*armmigrationassessment.FttAndRaidLevel{
		// 										to.Ptr(armmigrationassessment.FttAndRaidLevelFtt1Raid1),
		// 										to.Ptr(armmigrationassessment.FttAndRaidLevelFtt1Raid5),
		// 										to.Ptr(armmigrationassessment.FttAndRaidLevelFtt2Raid1),
		// 										to.Ptr(armmigrationassessment.FttAndRaidLevelFtt2Raid6),
		// 										to.Ptr(armmigrationassessment.FttAndRaidLevelFtt3Raid1)},
		// 										ReservedInstanceAvsNodes: []*armmigrationassessment.AzureAvsNodeType{
		// 											to.Ptr(armmigrationassessment.AzureAvsNodeTypeAV36),
		// 											to.Ptr(armmigrationassessment.AzureAvsNodeTypeAV36P),
		// 											to.Ptr(armmigrationassessment.AzureAvsNodeTypeAV52)},
		// 											ReservedInstanceSupportedCurrencies: []*armmigrationassessment.AzureCurrency{
		// 												to.Ptr(armmigrationassessment.AzureCurrencyUSD),
		// 												to.Ptr(armmigrationassessment.AzureCurrencyDKK),
		// 												to.Ptr(armmigrationassessment.AzureCurrencyCAD),
		// 												to.Ptr(armmigrationassessment.AzureCurrencyJPY),
		// 												to.Ptr(armmigrationassessment.AzureCurrencyKRW),
		// 												to.Ptr(armmigrationassessment.AzureCurrencyNZD),
		// 												to.Ptr(armmigrationassessment.AzureCurrencyNOK),
		// 												to.Ptr(armmigrationassessment.AzureCurrencyRUB),
		// 												to.Ptr(armmigrationassessment.AzureCurrencySEK),
		// 												to.Ptr(armmigrationassessment.AzureCurrencyGBP),
		// 												to.Ptr(armmigrationassessment.AzureCurrencyINR),
		// 												to.Ptr(armmigrationassessment.AzureCurrencyBRL),
		// 												to.Ptr(armmigrationassessment.AzureCurrencyTWD),
		// 												to.Ptr(armmigrationassessment.AzureCurrencyEUR),
		// 												to.Ptr(armmigrationassessment.AzureCurrencyCHF),
		// 												to.Ptr(armmigrationassessment.AzureCurrencyAUD)},
		// 												ReservedInstanceSupportedLocations: []*armmigrationassessment.AzureLocation{
		// 													to.Ptr(armmigrationassessment.AzureLocationEastAsia),
		// 													to.Ptr(armmigrationassessment.AzureLocationSoutheastAsia),
		// 													to.Ptr(armmigrationassessment.AzureLocationAustraliaEast),
		// 													to.Ptr(armmigrationassessment.AzureLocationAustraliaSoutheast),
		// 													to.Ptr(armmigrationassessment.AzureLocationBrazilSouth),
		// 													to.Ptr(armmigrationassessment.AzureLocationCanadaCentral),
		// 													to.Ptr(armmigrationassessment.AzureLocationCanadaEast),
		// 													to.Ptr(armmigrationassessment.AzureLocationWestEurope),
		// 													to.Ptr(armmigrationassessment.AzureLocationNorthEurope),
		// 													to.Ptr(armmigrationassessment.AzureLocationJapanEast),
		// 													to.Ptr(armmigrationassessment.AzureLocationJapanWest),
		// 													to.Ptr(armmigrationassessment.AzureLocationUkWest),
		// 													to.Ptr(armmigrationassessment.AzureLocationUkSouth),
		// 													to.Ptr(armmigrationassessment.AzureLocationNorthCentralUs),
		// 													to.Ptr(armmigrationassessment.AzureLocationEastUs),
		// 													to.Ptr(armmigrationassessment.AzureLocationSouthCentralUs),
		// 													to.Ptr(armmigrationassessment.AzureLocationCentralUs),
		// 													to.Ptr(armmigrationassessment.AzureLocationEastUs2),
		// 													to.Ptr(armmigrationassessment.AzureLocationWestUs),
		// 													to.Ptr(armmigrationassessment.AzureLocationFranceCentral),
		// 													to.Ptr(armmigrationassessment.AzureLocationSouthAfricaNorth),
		// 													to.Ptr(armmigrationassessment.AzureLocationGermanyWestCentral),
		// 													to.Ptr(armmigrationassessment.AzureLocationSwedenCentral),
		// 													to.Ptr(armmigrationassessment.AzureLocationWestUs2),
		// 													to.Ptr(armmigrationassessment.AzureLocationQatarCentral)},
		// 													ReservedInstanceSupportedOffers: []*armmigrationassessment.AzureOfferCode{
		// 														to.Ptr(armmigrationassessment.AzureOfferCodeMsazr0003P)},
		// 													},
		// 											}},
		// 										}
	}
}
