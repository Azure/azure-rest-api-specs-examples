package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/PredictionsGetTrainingResults.json
func ExamplePredictionsClient_GetTrainingResults() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPredictionsClient().GetTrainingResults(ctx, "TestHubRG", "sdkTestHub", "sdktest", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PredictionTrainingResults = armcustomerinsights.PredictionTrainingResults{
	// 	CanonicalProfiles: []*armcustomerinsights.CanonicalProfileDefinition{
	// 		{
	// 			CanonicalProfileID: to.Ptr[int32](100),
	// 			Properties: []*armcustomerinsights.CanonicalProfileDefinitionPropertiesItem{
	// 				{
	// 					Type: to.Ptr(armcustomerinsights.CanonicalPropertyValueTypeNumeric),
	// 					ProfileName: to.Ptr("Customers"),
	// 					ProfilePropertyName: to.Ptr("MobilePhone"),
	// 					Rank: to.Ptr[int32](1),
	// 					Value: to.Ptr("139162668.0000 - 34325747410.0000"),
	// 				},
	// 				{
	// 					Type: to.Ptr(armcustomerinsights.CanonicalPropertyValueTypeNumeric),
	// 					ProfileName: to.Ptr("Customers"),
	// 					ProfilePropertyName: to.Ptr("CustomerId"),
	// 					Rank: to.Ptr[int32](2),
	// 					Value: to.Ptr("592266139745.0000 - 592266139864.0000"),
	// 			}},
	// 	}},
	// 	PredictionDistribution: &armcustomerinsights.PredictionDistributionDefinition{
	// 		Distributions: []*armcustomerinsights.PredictionDistributionDefinitionDistributionsItem{
	// 			{
	// 				Negatives: to.Ptr[int64](0),
	// 				NegativesAboveThreshold: to.Ptr[int64](0),
	// 				Positives: to.Ptr[int64](0),
	// 				PositivesAboveThreshold: to.Ptr[int64](0),
	// 				ScoreThreshold: to.Ptr[int32](0),
	// 			},
	// 			{
	// 				Negatives: to.Ptr[int64](0),
	// 				NegativesAboveThreshold: to.Ptr[int64](0),
	// 				Positives: to.Ptr[int64](0),
	// 				PositivesAboveThreshold: to.Ptr[int64](0),
	// 				ScoreThreshold: to.Ptr[int32](1),
	// 		}},
	// 		TotalNegatives: to.Ptr[int64](0),
	// 		TotalPositives: to.Ptr[int64](0),
	// 	},
	// 	PrimaryProfileInstanceCount: to.Ptr[int64](0),
	// 	ScoreName: to.Ptr("sdktest"),
	// 	TenantID: to.Ptr("sdkTestHub"),
	// }
}
