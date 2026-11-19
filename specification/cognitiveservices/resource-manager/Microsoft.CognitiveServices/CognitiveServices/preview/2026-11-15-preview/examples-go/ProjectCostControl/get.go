package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-11-15-preview/ProjectCostControl/get.json
func ExampleProjectCostControlsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewProjectCostControlsClient().Get(ctx, "foundry-resource-group", "foundry-account", "production", "project-token-budget", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.ProjectCostControlsClientGetResponse{
	// 	Etag: to.Ptr(azcore.ETag("\"00000000-0000-0000-0000-000000000002\"")),
	// 	ProjectCostControl: armcognitiveservices.ProjectCostControl{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/foundry-resource-group/providers/Microsoft.CognitiveServices/accounts/foundry-account/projects/production/costControls/project-token-budget"),
	// 		Name: to.Ptr("project-token-budget"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/projects/costControls"),
	// 		Etag: to.Ptr(azcore.ETag("\"00000000-0000-0000-0000-000000000002\"")),
	// 		Properties: &armcognitiveservices.CostControlProperties{
	// 			Provider: to.Ptr("Microsoft.CognitiveServices"),
	// 			Category: to.Ptr(armcognitiveservices.CostControlCategoryTokens),
	// 			Amount: to.Ptr[float64](2000000),
	// 			TimeGrain: to.Ptr(armcognitiveservices.CostControlTimeGrainDaily),
	// 			TimePeriod: &armcognitiveservices.CostControlTimePeriod{
	// 				StartDate: to.Ptr(time.Date(2026, time.November, 15, 0, 0, 0, 0, time.UTC)),
	// 				EndDate: to.Ptr(time.Date(2027, time.November, 15, 0, 0, 0, 0, time.UTC)),
	// 			},
	// 			Filter: &armcognitiveservices.CostControlFilter{
	// 				AdditionalInfo: &armcognitiveservices.CostControlAdditionalInfoFilter{
	// 					Name: to.Ptr(armcognitiveservices.CostControlFilterNameTokenType),
	// 					Operator: to.Ptr("In"),
	// 					Values: []*string{
	// 						to.Ptr("input"),
	// 						to.Ptr("output"),
	// 					},
	// 				},
	// 			},
	// 			Notifications: map[string]*armcognitiveservices.CostControlNotification{
	// 				"usageAlert": &armcognitiveservices.CostControlNotification{
	// 					EventGridConnectionID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/foundry-resource-group/providers/Microsoft.CognitiveServices/accounts/foundry-account/connections/budget-alerts"),
	// 					ThresholdType: to.Ptr(armcognitiveservices.CostControlThresholdTypePercentage),
	// 					Threshold: to.Ptr[float64](80),
	// 					Enabled: to.Ptr(true),
	// 					Operator: to.Ptr("GreaterThanOrEqualTo"),
	// 				},
	// 			},
	// 			ProviderProperties: &armcognitiveservices.CostControlProviderProperties{
	// 				Actions: map[string]*armcognitiveservices.CostControlAction{
	// 					"stopRequests": &armcognitiveservices.CostControlAction{
	// 						Type: to.Ptr("Block"),
	// 						ThresholdType: to.Ptr(armcognitiveservices.CostControlThresholdTypePercentage),
	// 						Threshold: to.Ptr[float64](100),
	// 						Operator: to.Ptr("GreaterThanOrEqualTo"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
