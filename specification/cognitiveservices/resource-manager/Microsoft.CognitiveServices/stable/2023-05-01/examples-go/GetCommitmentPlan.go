package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/GetCommitmentPlan.json
func ExampleCommitmentPlansClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCommitmentPlansClient().Get(ctx, "resourceGroupName", "accountName", "commitmentPlanName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CommitmentPlan = armcognitiveservices.CommitmentPlan{
	// 	Name: to.Ptr("commitmentPlanName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/accounts/commitmentPlans"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/commitmentPlans/commitmentPlanName"),
	// 	Properties: &armcognitiveservices.CommitmentPlanProperties{
	// 		AutoRenew: to.Ptr(true),
	// 		Current: &armcognitiveservices.CommitmentPeriod{
	// 			Tier: to.Ptr("T1"),
	// 		},
	// 		HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
	// 		PlanType: to.Ptr("Speech2Text"),
	// 	},
	// }
}
