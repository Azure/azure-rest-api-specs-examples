package armcustomerinsights_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RelationshipLinksListByHub.json
func ExampleRelationshipLinksClient_NewListByHubPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcustomerinsights.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRelationshipLinksClient().NewListByHubPager("TestHubRG", "sdkTestHub", nil)
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
		// page.RelationshipLinkListResult = armcustomerinsights.RelationshipLinkListResult{
		// 	Value: []*armcustomerinsights.RelationshipLinkResourceFormat{
		// 		{
		// 			Name: to.Ptr("sdkTestHub/Somelink"),
		// 			Type: to.Ptr("Microsoft.CustomerInsights/hubs/relationshipLinks"),
		// 			ID: to.Ptr("/subscriptions/c909e979-ef71-4def-a970-bc7c154db8c5/resourceGroups/TestHubRG/providers/Microsoft.CustomerInsights/hubs/sdkTestHub/relationshipLinks/Somelink"),
		// 			Properties: &armcustomerinsights.RelationshipLinkDefinition{
		// 				Description: map[string]*string{
		// 					"en-us": to.Ptr("Link Description"),
		// 				},
		// 				DisplayName: map[string]*string{
		// 					"en-us": to.Ptr("Link DisplayName"),
		// 				},
		// 				InteractionType: to.Ptr("testInteraction4332"),
		// 				LinkName: to.Ptr("Somelink"),
		// 				ProfilePropertyReferences: []*armcustomerinsights.ParticipantProfilePropertyReference{
		// 					{
		// 						InteractionPropertyName: to.Ptr("profile1"),
		// 						ProfilePropertyName: to.Ptr("ProfileId"),
		// 				}},
		// 				ProvisioningState: to.Ptr(armcustomerinsights.ProvisioningStatesSucceeded),
		// 				RelatedProfilePropertyReferences: []*armcustomerinsights.ParticipantProfilePropertyReference{
		// 					{
		// 						InteractionPropertyName: to.Ptr("profile1"),
		// 						ProfilePropertyName: to.Ptr("ProfileId"),
		// 				}},
		// 				RelationshipName: to.Ptr("testProfile2326994"),
		// 				TenantID: to.Ptr("sdktesthub"),
		// 			},
		// 	}},
		// }
	}
}
