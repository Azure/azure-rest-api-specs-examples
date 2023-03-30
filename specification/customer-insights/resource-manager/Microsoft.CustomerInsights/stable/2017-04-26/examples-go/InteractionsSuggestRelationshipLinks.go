package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/InteractionsSuggestRelationshipLinks.json
func ExampleInteractionsClient_SuggestRelationshipLinks() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewInteractionsClient().SuggestRelationshipLinks(ctx, "TestHubRG", "sdkTestHub", "Deposit", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SuggestRelationshipLinksResponse = armcustomerinsights.SuggestRelationshipLinksResponse{
	// 	InteractionName: to.Ptr("Deposit"),
	// 	SuggestedRelationships: []*armcustomerinsights.RelationshipsLookup{
	// 		{
	// 			ProfileName: to.Ptr("Branch"),
	// 			ProfilePropertyReferences: []*armcustomerinsights.ParticipantProfilePropertyReference{
	// 				{
	// 					InteractionPropertyName: to.Ptr("BranchId"),
	// 					ProfilePropertyName: to.Ptr("BranchId"),
	// 			}},
	// 			RelatedProfileName: to.Ptr("Contact"),
	// 			RelatedProfilePropertyReferences: []*armcustomerinsights.ParticipantProfilePropertyReference{
	// 				{
	// 					InteractionPropertyName: to.Ptr("ContactId"),
	// 					ProfilePropertyName: to.Ptr("ContactId"),
	// 			}},
	// 		},
	// 		{
	// 			ProfileName: to.Ptr("Contact"),
	// 			ProfilePropertyReferences: []*armcustomerinsights.ParticipantProfilePropertyReference{
	// 				{
	// 					InteractionPropertyName: to.Ptr("ContactId"),
	// 					ProfilePropertyName: to.Ptr("ContactId"),
	// 			}},
	// 			RelatedProfileName: to.Ptr("Branch"),
	// 			RelatedProfilePropertyReferences: []*armcustomerinsights.ParticipantProfilePropertyReference{
	// 				{
	// 					InteractionPropertyName: to.Ptr("BranchId"),
	// 					ProfilePropertyName: to.Ptr("BranchId"),
	// 			}},
	// 	}},
	// }
}
