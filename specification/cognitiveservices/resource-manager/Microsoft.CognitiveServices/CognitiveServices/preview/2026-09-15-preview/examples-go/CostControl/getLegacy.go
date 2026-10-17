package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-09-15-preview/CostControl/getLegacy.json
func ExampleCostControlsClient_Get_readALegacyCostControlWithoutAuthoringItsLegacyValues() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCostControlsClient().Get(ctx, "foundry-resource-group", "foundry-account", "legacy-budget", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.CostControlsClientGetResponse{
	// 	Etag: to.Ptr("\"00000000-0000-0000-0000-000000000001\""),
	// 	CostControl: armcognitiveservices.CostControl{
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/foundry-resource-group/providers/Microsoft.CognitiveServices/accounts/foundry-account/costControls/legacy-budget"),
	// 		Name: to.Ptr("legacy-budget"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/accounts/costControls"),
	// 		Etag: to.Ptr("\"00000000-0000-0000-0000-000000000001\""),
	// 		Properties: &armcognitiveservices.CostControlProperties{
	// 			DisplayName: to.Ptr("Legacy cost center budget"),
	// 			Rules: []*armcognitiveservices.CostControlRule{
	// 				{
	// 					Name: to.Ptr("legacy-custom-budget"),
	// 					CounterKey: &armcognitiveservices.CostControlDimension{
	// 						Type: to.Ptr(armcognitiveservices.CostControlDimensionTypeCustom),
	// 						Attribute: to.Ptr("legacy cost center value"),
	// 					},
	// 					Unit: to.Ptr(armcognitiveservices.CostControlUnitUsd),
	// 					Amount: to.Ptr[float64](1000),
	// 					Period: to.Ptr(armcognitiveservices.CostControlPeriodYear),
	// 					Recurring: to.Ptr(true),
	// 					Thresholds: []*armcognitiveservices.CostControlThreshold{
	// 						{
	// 							Type: to.Ptr(armcognitiveservices.CostControlThresholdTypePercentage),
	// 							Value: to.Ptr[float64](80),
	// 							Action: to.Ptr(armcognitiveservices.CostControlThresholdActionAudit),
	// 						},
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
