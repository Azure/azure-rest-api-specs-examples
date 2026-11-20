package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-11-15-preview/AccountCostControl/get.json
func ExampleAccountCostControlsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountCostControlsClient().Get(ctx, "foundry-resource-group", "foundry-account", "production-agents", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.AccountCostControlsClientGetResponse{
	// 	Etag: to.Ptr(azcore.ETag("\"00000000-0000-0000-0000-000000000002\"")),
	// 	AccountCostControl: armcognitiveservices.AccountCostControl{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/foundry-resource-group/providers/Microsoft.CognitiveServices/accounts/foundry-account/costControls/production-agents"),
	// 		Name: to.Ptr("production-agents"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/costControls"),
	// 		Etag: to.Ptr(azcore.ETag("\"00000000-0000-0000-0000-000000000002\"")),
	// 		Properties: &armcognitiveservices.CostControlProperties{
	// 			Provider: to.Ptr("Microsoft.CognitiveServices"),
	// 			Grouping: []*armcognitiveservices.CostControlGrouping{
	// 				{
	// 					Name: to.Ptr(armcognitiveservices.CostControlGroupingNameAgent),
	// 					Type: to.Ptr("AdditionalInfoKey"),
	// 				},
	// 			},
	// 			Category: to.Ptr(armcognitiveservices.CostControlCategoryCost),
	// 			Amount: to.Ptr[float64](1500),
	// 			TimeGrain: to.Ptr(armcognitiveservices.CostControlTimeGrainMonthly),
	// 			TimePeriod: &armcognitiveservices.CostControlTimePeriod{
	// 				StartDate: to.Ptr(time.Date(2026, time.November, 15, 0, 0, 0, 0, time.UTC)),
	// 				EndDate: to.Ptr(time.Date(2027, time.November, 15, 0, 0, 0, 0, time.UTC)),
	// 			},
	// 			Filter: &armcognitiveservices.CostControlFilter{
	// 				And: []*armcognitiveservices.CostControlFilter{
	// 					{
	// 						AdditionalInfo: &armcognitiveservices.CostControlAdditionalInfoFilter{
	// 							Name: to.Ptr(armcognitiveservices.CostControlFilterNameProject),
	// 							Operator: to.Ptr("In"),
	// 							Values: []*string{
	// 								to.Ptr("production"),
	// 							},
	// 						},
	// 					},
	// 					{
	// 						AdditionalInfo: &armcognitiveservices.CostControlAdditionalInfoFilter{
	// 							Name: to.Ptr(armcognitiveservices.CostControlFilterNameAgent),
	// 							Operator: to.Ptr("In"),
	// 							Values: []*string{
	// 								to.Ptr("customer-support"),
	// 								to.Ptr("order-assistant"),
	// 							},
	// 						},
	// 					},
	// 					{
	// 						AdditionalInfo: &armcognitiveservices.CostControlAdditionalInfoFilter{
	// 							Name: to.Ptr(armcognitiveservices.CostControlFilterNameDeployment),
	// 							Operator: to.Ptr("In"),
	// 							Values: []*string{
	// 								to.Ptr("customer-support-model"),
	// 							},
	// 						},
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
