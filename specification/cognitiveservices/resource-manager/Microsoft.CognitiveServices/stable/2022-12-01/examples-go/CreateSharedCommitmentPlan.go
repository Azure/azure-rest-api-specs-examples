package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-12-01/examples/CreateSharedCommitmentPlan.json
func ExampleCommitmentPlansClient_BeginCreateOrUpdatePlan() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcognitiveservices.NewCommitmentPlansClient("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdatePlan(ctx, "resourceGroupName", "commitmentPlanName", armcognitiveservices.CommitmentPlan{
		Kind:     to.Ptr("SpeechServices"),
		Location: to.Ptr("West US"),
		Properties: &armcognitiveservices.CommitmentPlanProperties{
			AutoRenew: to.Ptr(true),
			Current: &armcognitiveservices.CommitmentPeriod{
				Tier: to.Ptr("T1"),
			},
			HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
			PlanType:     to.Ptr("STT"),
		},
		SKU: &armcognitiveservices.SKU{
			Name: to.Ptr("S0"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CommitmentPlan = armcognitiveservices.CommitmentPlan{
	// 	Name: to.Ptr("commitmentPlanName"),
	// 	Type: to.Ptr("Microsoft.CognitiveServices/commitmentPlans"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanName"),
	// 	Kind: to.Ptr("SpeechServices"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armcognitiveservices.CommitmentPlanProperties{
	// 		AutoRenew: to.Ptr(true),
	// 		Current: &armcognitiveservices.CommitmentPeriod{
	// 			Tier: to.Ptr("T1"),
	// 		},
	// 		HostingModel: to.Ptr(armcognitiveservices.HostingModelWeb),
	// 		PlanType: to.Ptr("STT"),
	// 		ProvisioningState: to.Ptr(armcognitiveservices.CommitmentPlanProvisioningStateSucceeded),
	// 	},
	// 	SKU: &armcognitiveservices.SKU{
	// 		Name: to.Ptr("S0"),
	// 	},
	// }
}
