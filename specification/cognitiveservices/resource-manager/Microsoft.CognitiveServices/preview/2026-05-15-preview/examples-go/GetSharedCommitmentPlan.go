package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-05-15-preview/GetSharedCommitmentPlan.json
func ExampleCommitmentPlansClient_GetPlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCommitmentPlansClient().GetPlan(ctx, "resourceGroupName", "commitmentPlanName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armcognitiveservices.CommitmentPlansClientGetPlanResponse{
	// 	CommitmentPlan: armcognitiveservices.CommitmentPlan{
	// 		Name: to.Ptr("commitmentPlanName"),
	// 		Type: to.Ptr("Microsoft.CognitiveServices/commitmentPlans"),
	// 		ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanName"),
	// 		Kind: to.Ptr("SpeechServices"),
	// 		Location: to.Ptr("West US"),
	// 		Properties: &armcognitiveservices.CommitmentPlanProperties{
	// 			AutoRenew: to.Ptr(true),
	// 			Current: &armcognitiveservices.CommitmentPeriod{
	// 				Tier: to.Ptr("T1"),
	// 			},
	// 			HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
	// 			PlanType: to.Ptr("STT"),
	// 			ProvisioningState: to.Ptr(armcognitiveservices.CommitmentPlanProvisioningStateSucceeded),
	// 		},
	// 		SKU: &armcognitiveservices.SKU{
	// 			Name: to.Ptr("S0"),
	// 		},
	// 	},
	// }
}
