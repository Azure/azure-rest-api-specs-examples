package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/ListSharedCommitmentPlansBySubscription.json
func ExampleCommitmentPlansClient_NewListPlansBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCommitmentPlansClient().NewListPlansBySubscriptionPager(nil)
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
		// page.CommitmentPlanListResult = armcognitiveservices.CommitmentPlanListResult{
		// 	Value: []*armcognitiveservices.CommitmentPlan{
		// 		{
		// 			Name: to.Ptr("commitmentPlanName"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/commitmentPlans"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanName"),
		// 			Kind: to.Ptr("SpeechServices"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armcognitiveservices.CommitmentPlanProperties{
		// 				AutoRenew: to.Ptr(true),
		// 				Current: &armcognitiveservices.CommitmentPeriod{
		// 					Tier: to.Ptr("T1"),
		// 				},
		// 				HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
		// 				PlanType: to.Ptr("STT"),
		// 				ProvisioningState: to.Ptr(armcognitiveservices.CommitmentPlanProvisioningStateSucceeded),
		// 			},
		// 			SKU: &armcognitiveservices.SKU{
		// 				Name: to.Ptr("S0"),
		// 			},
		// 	}},
		// }
	}
}
