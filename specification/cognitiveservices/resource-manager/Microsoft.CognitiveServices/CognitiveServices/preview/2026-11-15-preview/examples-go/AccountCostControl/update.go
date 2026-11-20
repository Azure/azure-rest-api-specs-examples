package armcognitiveservices_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-11-15-preview/AccountCostControl/update.json
func ExampleAccountCostControlsClient_Update_mergeNamedBudgetSettingsAndRestoreOwnerDefaultGrouping() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAccountCostControlsClient().Update(ctx, "foundry-resource-group", "foundry-account", "production-agents", armcognitiveservices.CostControlPatch{
		Properties: &armcognitiveservices.CostControlPatchProperties{
			Grouping: []*armcognitiveservices.CostControlGrouping{},
			Amount:   to.Ptr[float64](1500),
			TimePeriod: &armcognitiveservices.CostControlTimePeriodPatch{
				EndDate: to.Ptr(time.Date(2028, time.November, 15, 0, 0, 0, 0, time.UTC)),
			},
			Notifications: map[string]*armcognitiveservices.CostControlNotificationPatch{
				"usageAlert": {
					Enabled: to.Ptr(false),
				},
			},
			ProviderProperties: &armcognitiveservices.CostControlProviderPropertiesPatch{
				Actions: map[string]*armcognitiveservices.CostControlActionPatch{
					"stopRequests": {
						Threshold: to.Ptr[float64](90),
					},
				},
			},
		},
	}, &armcognitiveservices.AccountCostControlsClientUpdateOptions{
		IfMatch: to.Ptr("\"00000000-0000-0000-0000-000000000001\"")})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.AccountCostControlsClientUpdateResponse{
	// 	Etag: to.Ptr(azcore.ETag("\"00000000-0000-0000-0000-000000000002\"")),
	// 	AccountCostControl: armcognitiveservices.AccountCostControl{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/foundry-resource-group/providers/Microsoft.CognitiveServices/accounts/foundry-account/costControls/production-agents"),
	// 		Name: to.Ptr("production-agents"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/costControls"),
	// 		Etag: to.Ptr(azcore.ETag("\"00000000-0000-0000-0000-000000000002\"")),
	// 		Properties: &armcognitiveservices.CostControlProperties{
	// 			Provider: to.Ptr("Microsoft.CognitiveServices"),
	// 			Category: to.Ptr(armcognitiveservices.CostControlCategoryCost),
	// 			Amount: to.Ptr[float64](1500),
	// 			TimeGrain: to.Ptr(armcognitiveservices.CostControlTimeGrainMonthly),
	// 			TimePeriod: &armcognitiveservices.CostControlTimePeriod{
	// 				StartDate: to.Ptr(time.Date(2026, time.November, 15, 0, 0, 0, 0, time.UTC)),
	// 				EndDate: to.Ptr(time.Date(2028, time.November, 15, 0, 0, 0, 0, time.UTC)),
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
	// 					Enabled: to.Ptr(false),
	// 					Operator: to.Ptr("GreaterThanOrEqualTo"),
	// 				},
	// 			},
	// 			ProviderProperties: &armcognitiveservices.CostControlProviderProperties{
	// 				Actions: map[string]*armcognitiveservices.CostControlAction{
	// 					"stopRequests": &armcognitiveservices.CostControlAction{
	// 						Type: to.Ptr("Block"),
	// 						ThresholdType: to.Ptr(armcognitiveservices.CostControlThresholdTypePercentage),
	// 						Threshold: to.Ptr[float64](90),
	// 						Operator: to.Ptr("GreaterThanOrEqualTo"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
