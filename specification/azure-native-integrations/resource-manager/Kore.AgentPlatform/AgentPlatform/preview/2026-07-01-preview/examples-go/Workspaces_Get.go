package armkoreaiartemis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/koreaiartemis/armkoreaiartemis"
)

// Generated from example definition: 2026-07-01-preview/Workspaces_Get.json
func ExampleWorkspacesClient_Get_workspacesGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkoreaiartemis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().Get(ctx, "contoso-rg", "contoso-org", "contoso-prod", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkoreaiartemis.WorkspacesClientGetResponse{
	// 	WorkspaceResource: armkoreaiartemis.WorkspaceResource{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Kore.AgentPlatform/organizations/contoso-org/workspaces/contoso-prod"),
	// 		Name: to.Ptr("contoso-prod"),
	// 		Type: to.Ptr("Kore.AgentPlatform/organizations/workspaces"),
	// 		Properties: &armkoreaiartemis.WorkspaceProperties{
	// 			PartnerProperties: &armkoreaiartemis.WorkspacePartnerProperties{
	// 				Name: to.Ptr("Contoso Production"),
	// 				OwnerEmail: to.Ptr("owner@contoso.com"),
	// 				PartnerResourceID: to.Ptr("ws_01JK2MXYZ9ABCDEF000002"),
	// 			},
	// 			IsDefault: to.Ptr(false),
	// 			Status: to.Ptr(armkoreaiartemis.WorkspaceLifecycleStatusActive),
	// 			ProvisioningState: to.Ptr(armkoreaiartemis.ResourceProvisioningStateSucceeded),
	// 			MarketplaceDetails: []*armkoreaiartemis.MarketplaceDetails{
	// 				{
	// 					OfferDetails: &armkoreaiartemis.OfferDetails{
	// 						OfferID: to.Ptr("kore-ai-platform"),
	// 						PlanID: to.Ptr("business-monthly"),
	// 						PublisherID: to.Ptr("kore-ai"),
	// 						PlanName: to.Ptr("Business Monthly"),
	// 						TermID: to.Ptr("1m"),
	// 						TermUnit: to.Ptr("P1M"),
	// 					},
	// 					SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Microsoft.SaaS/resources/kore-ai-saas"),
	// 					SubscriptionStatus: to.Ptr(armkoreaiartemis.MarketplaceSubscriptionStatusSubscribed),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
