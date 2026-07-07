package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices/v4"
)

// Generated from example definition: 2026-05-15-preview/ListCommitmentPlans.json
func ExampleCommitmentPlansClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("00000000-1111-2222-3333-444444444444", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCommitmentPlansClient().NewListPager("resourceGroupName", "accountName", nil)
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
		// page = armcognitiveservices.CommitmentPlansClientListResponse{
		// 	CommitmentPlanListResult: armcognitiveservices.CommitmentPlanListResult{
		// 		Value: []*armcognitiveservices.CommitmentPlan{
		// 			{
		// 				Name: to.Ptr("commitmentPlanName"),
		// 				Type: to.Ptr("Microsoft.CognitiveServices/accounts/commitmentPlans"),
		// 				ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/commitmentPlans/commitmentPlanName"),
		// 				Properties: &armcognitiveservices.CommitmentPlanProperties{
		// 					AutoRenew: to.Ptr(true),
		// 					Current: &armcognitiveservices.CommitmentPeriod{
		// 						Tier: to.Ptr("T1"),
		// 					},
		// 					HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
		// 					PlanType: to.Ptr("Speech2Text"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
