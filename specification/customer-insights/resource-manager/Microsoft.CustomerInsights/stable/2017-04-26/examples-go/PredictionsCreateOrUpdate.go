package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/PredictionsCreateOrUpdate.json
func ExamplePredictionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcustomerinsights.NewPredictionsClient("c909e979-ef71-4def-a970-bc7c154db8c5", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"TestHubRG",
		"sdkTestHub",
		"sdktest",
		armcustomerinsights.PredictionResourceFormat{
			Properties: &armcustomerinsights.Prediction{
				Description: map[string]*string{
					"en-us": to.Ptr("sdktest"),
				},
				AutoAnalyze: to.Ptr(true),
				DisplayName: map[string]*string{
					"en-us": to.Ptr("sdktest"),
				},
				Grades:                   []*armcustomerinsights.PredictionGradesItem{},
				InvolvedInteractionTypes: []*string{},
				InvolvedKpiTypes:         []*string{},
				InvolvedRelationships:    []*string{},
				Mappings: &armcustomerinsights.PredictionMappings{
					Grade:  to.Ptr("sdktest_Grade"),
					Reason: to.Ptr("sdktest_Reason"),
					Score:  to.Ptr("sdktest_Score"),
				},
				NegativeOutcomeExpression: to.Ptr("Customers.FirstName = 'Mike'"),
				PositiveOutcomeExpression: to.Ptr("Customers.FirstName = 'David'"),
				PredictionName:            to.Ptr("sdktest"),
				PrimaryProfileType:        to.Ptr("Customers"),
				ScopeExpression:           to.Ptr("*"),
				ScoreLabel:                to.Ptr("score label"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
