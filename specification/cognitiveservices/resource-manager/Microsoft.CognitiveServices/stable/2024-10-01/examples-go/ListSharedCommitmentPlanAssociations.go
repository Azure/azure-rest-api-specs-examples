package armcognitiveservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/069a65e8a6d1a6c0c58d9a9d97610b7103b6e8a5/specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/ListSharedCommitmentPlanAssociations.json
func ExampleCommitmentPlansClient_NewListAssociationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcognitiveservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCommitmentPlansClient().NewListAssociationsPager("resourceGroupName", "commitmentPlanName", nil)
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
		// page.CommitmentPlanAccountAssociationListResult = armcognitiveservices.CommitmentPlanAccountAssociationListResult{
		// 	Value: []*armcognitiveservices.CommitmentPlanAccountAssociation{
		// 		{
		// 			Name: to.Ptr("accountAssociationName"),
		// 			Type: to.Ptr("Microsoft.CognitiveServices/commitmentPlans/accountAssociations"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/commitmentPlans/commitmentPlanName/accountAssociations/accountAssociationName"),
		// 			Properties: &armcognitiveservices.CommitmentPlanAccountAssociationProperties{
		// 				AccountID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName"),
		// 			},
		// 	}},
		// }
	}
}
