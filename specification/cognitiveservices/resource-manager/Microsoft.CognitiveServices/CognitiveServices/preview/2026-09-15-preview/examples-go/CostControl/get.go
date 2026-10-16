package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-09-15-preview/CostControl/get.json
func ExampleCostControlsClient_Get_getACostControlWithAllSettings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCostControlsClient().Get(ctx, "foundry-resource-group", "foundry-account", "production-agents", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.CostControlsClientGetResponse{
	// 	Etag: to.Ptr("\"00000000-0000-0000-0000-000000000001\""),
	// 	CostControl: armcognitiveservices.CostControl{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/foundry-resource-group/providers/Microsoft.CognitiveServices/accounts/foundry-account/costControls/production-agents"),
	// 		Name: to.Ptr("production-agents"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/costControls"),
	// 		Etag: to.Ptr("\"00000000-0000-0000-0000-000000000001\""),
	// 		Properties: &armcognitiveservices.CostControlProperties{
	// 			DisplayName: to.Ptr("Production agent monthly budget"),
	// 			Rules: []*armcognitiveservices.CostControlRule{
	// 				{
	// 					Name: to.Ptr("monthly-agent-budget"),
	// 					Unit: to.Ptr(armcognitiveservices.CostControlUnitUsd),
	// 					Amount: to.Ptr[float64](1000),
	// 					Period: to.Ptr(armcognitiveservices.CostControlPeriodMonth),
	// 					Recurring: to.Ptr(true),
	// 					Match: &armcognitiveservices.CostControlMatch{
	// 						IdentityObjectIDs: []*string{
	// 							to.Ptr("11111111-2222-3333-4444-555555555555"),
	// 						},
	// 						SessionIDs: []*string{
	// 							to.Ptr("production-session"),
	// 						},
	// 						ProjectIDs: []*string{
	// 							to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/foundry-resource-group/providers/Microsoft.CognitiveServices/accounts/foundry-account/projects/production"),
	// 						},
	// 						AgentResourceIDs: []*string{
	// 							to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/accounts/foundry-account/project/production/agent/customer-support-agent"),
	// 							to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/accounts/foundry-account/project/production/agent/sales-assistant-agent"),
	// 						},
	// 					},
	// 					Thresholds: []*armcognitiveservices.CostControlThreshold{
	// 						{
	// 							Type: to.Ptr(armcognitiveservices.CostControlThresholdTypePercentage),
	// 							Value: to.Ptr[float64](80),
	// 							Action: to.Ptr(armcognitiveservices.CostControlThresholdActionAlert),
	// 						},
	// 						{
	// 							Type: to.Ptr(armcognitiveservices.CostControlThresholdTypeAbsolute),
	// 							Value: to.Ptr[float64](1000),
	// 							Action: to.Ptr(armcognitiveservices.CostControlThresholdActionBlock),
	// 						},
	// 					},
	// 					CounterKey: &armcognitiveservices.CostControlDimension{
	// 						Type: to.Ptr(armcognitiveservices.CostControlDimensionTypeAgent),
	// 					},
	// 				},
	// 				{
	// 					Name: to.Ptr("daily-account-budget"),
	// 					Unit: to.Ptr(armcognitiveservices.CostControlUnitUsd),
	// 					Amount: to.Ptr[float64](100),
	// 					Period: to.Ptr(armcognitiveservices.CostControlPeriodDay),
	// 					Recurring: to.Ptr(true),
	// 					Thresholds: []*armcognitiveservices.CostControlThreshold{
	// 						{
	// 							Type: to.Ptr(armcognitiveservices.CostControlThresholdTypePercentage),
	// 							Value: to.Ptr[float64](90),
	// 							Action: to.Ptr(armcognitiveservices.CostControlThresholdActionAlert),
	// 						},
	// 					},
	// 					CounterKey: &armcognitiveservices.CostControlDimension{
	// 						Type: to.Ptr(armcognitiveservices.CostControlDimensionTypeAccount),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
