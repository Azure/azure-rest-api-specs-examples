package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ExpressRouteGatewayGetResiliencyInformation.json
func ExampleExpressRouteGatewaysClient_BeginGetResiliencyInformation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewExpressRouteGatewaysClient().BeginGetResiliencyInformation(ctx, "rg1", "ergw1", &armnetwork.ExpressRouteGatewaysClientBeginGetResiliencyInformationOptions{
		AttemptRefresh: to.Ptr(false)})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ExpressRouteGatewaysClientGetResiliencyInformationResponse{
	// 	GatewayResiliencyInformation: armnetwork.GatewayResiliencyInformation{
	// 		OverallScore: to.Ptr("85"),
	// 		ScoreChange: to.Ptr("5"),
	// 		MinScoreFromRecommendations: to.Ptr("2"),
	// 		MaxScoreFromRecommendations: to.Ptr("10"),
	// 		LastComputedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-01T00:00:00Z"); return t}()),
	// 		NextEligibleComputeTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2025-01-01T01:00:00Z"); return t}()),
	// 		Components: []*armnetwork.ResiliencyRecommendationComponents{
	// 			{
	// 				Name: to.Ptr("Disaster Recovery"),
	// 				CurrentScore: to.Ptr("-"),
	// 				MaxScore: to.Ptr("-"),
	// 				Recommendations: []*armnetwork.GatewayResiliencyRecommendation{
	// 					{
	// 						RecommendationTitle: to.Ptr("Site Failover Not Yet Validated"),
	// 						RecommendationID: to.Ptr("6"),
	// 						Severity: to.Ptr("High"),
	// 						RecommendationText: to.Ptr("Use Resiliency Validation Tests at least once every 3 months."),
	// 						CallToActionText: to.Ptr("Run Resiliency Validation Tests."),
	// 						CallToActionLink: to.Ptr("https://learn.microsoft.com/en-us/azure/expressroute/resiliency-insights"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
