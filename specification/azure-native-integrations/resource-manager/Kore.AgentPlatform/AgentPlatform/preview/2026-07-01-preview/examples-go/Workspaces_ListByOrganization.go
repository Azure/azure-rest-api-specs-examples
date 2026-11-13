package armkoreaiartemis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/koreaiartemis/armkoreaiartemis"
)

// Generated from example definition: 2026-07-01-preview/Workspaces_ListByOrganization.json
func ExampleWorkspacesClient_NewListByOrganizationPager_workspacesListByOrganization() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkoreaiartemis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacesClient().NewListByOrganizationPager("contoso-rg", "contoso-org", nil)
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
		// page = armkoreaiartemis.WorkspacesClientListByOrganizationResponse{
		// 	WorkspaceResourceListResult: armkoreaiartemis.WorkspaceResourceListResult{
		// 		Value: []*armkoreaiartemis.WorkspaceResource{
		// 			{
		// 				Type: to.Ptr("Kore.AgentPlatform/organizations/workspaces"),
		// 				Name: to.Ptr("my first workspace"),
		// 				Properties: &armkoreaiartemis.WorkspaceProperties{
		// 					MarketplaceDetails: []*armkoreaiartemis.MarketplaceDetails{
		// 						{
		// 							SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Microsoft.SaaS/resources/kore-ai-saas"),
		// 							SubscriptionStatus: to.Ptr(armkoreaiartemis.MarketplaceSubscriptionStatusSubscribed),
		// 							OfferDetails: &armkoreaiartemis.OfferDetails{
		// 								PublisherID: to.Ptr("kore-ai"),
		// 								OfferID: to.Ptr("kore-ai-platform"),
		// 								PlanID: to.Ptr("business-monthly"),
		// 								PlanName: to.Ptr("Business Monthly"),
		// 								TermUnit: to.Ptr("P1M"),
		// 								TermID: to.Ptr("1m"),
		// 							},
		// 						},
		// 					},
		// 					IsDefault: to.Ptr(true),
		// 					PartnerProperties: &armkoreaiartemis.WorkspacePartnerProperties{
		// 						OwnerEmail: to.Ptr("admin@contoso.com"),
		// 						Name: to.Ptr("my first workspace"),
		// 					},
		// 					Status: to.Ptr(armkoreaiartemis.WorkspaceLifecycleStatusActive),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Kore.AgentPlatform/organizations/contoso-org/workspaces/my-first-workspace"),
		// 			},
		// 			{
		// 				Type: to.Ptr("Kore.AgentPlatform/organizations/workspaces"),
		// 				Name: to.Ptr("Contoso Production"),
		// 				Properties: &armkoreaiartemis.WorkspaceProperties{
		// 					MarketplaceDetails: []*armkoreaiartemis.MarketplaceDetails{
		// 						{
		// 							SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Microsoft.SaaS/resources/kore-ai-saas"),
		// 							SubscriptionStatus: to.Ptr(armkoreaiartemis.MarketplaceSubscriptionStatusSubscribed),
		// 							OfferDetails: &armkoreaiartemis.OfferDetails{
		// 								PublisherID: to.Ptr("kore-ai"),
		// 								OfferID: to.Ptr("kore-ai-platform"),
		// 								PlanID: to.Ptr("business-monthly"),
		// 								PlanName: to.Ptr("Business Monthly"),
		// 								TermUnit: to.Ptr("P1M"),
		// 								TermID: to.Ptr("1m"),
		// 							},
		// 						},
		// 					},
		// 					IsDefault: to.Ptr(false),
		// 					PartnerProperties: &armkoreaiartemis.WorkspacePartnerProperties{
		// 						OwnerEmail: to.Ptr("owner@contoso.com"),
		// 						Name: to.Ptr("Contoso Production"),
		// 					},
		// 					Status: to.Ptr(armkoreaiartemis.WorkspaceLifecycleStatusActive),
		// 				},
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Kore.AgentPlatform/organizations/contoso-org/workspaces/Contoso-Production"),
		// 			},
		// 		},
		// 	},
		// }
	}
}
