package armkoreaiartemis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/koreaiartemis/armkoreaiartemis"
)

// Generated from example definition: 2026-07-01-preview/Organizations_Get.json
func ExampleOrganizationsClient_Get_organizationsGet() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkoreaiartemis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationsClient().Get(ctx, "contoso-rg", "contoso-org", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkoreaiartemis.OrganizationsClientGetResponse{
	// 	OrganizationResource: armkoreaiartemis.OrganizationResource{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Kore.AgentPlatform/organizations/contoso-org"),
	// 		Name: to.Ptr("contoso-org"),
	// 		Type: to.Ptr("Kore.AgentPlatform/organizations"),
	// 		Location: to.Ptr("eastus2euap"),
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 		},
	// 		Properties: &armkoreaiartemis.OrganizationProperties{
	// 			PartnerProperties: &armkoreaiartemis.OrganizationPartnerProperties{
	// 				OrganizationName: to.Ptr("Contoso AI Platform"),
	// 				BillingEmail: to.Ptr("billing@contoso.com"),
	// 				OwnerEmail: to.Ptr("admin@contoso.com"),
	// 				DefaultWorkspace: &armkoreaiartemis.DefaultWorkspaceProperties{
	// 					Name: to.Ptr("my first workspace"),
	// 					MarketplaceDetails: []*armkoreaiartemis.MarketplaceDetails{
	// 						{
	// 							OfferDetails: &armkoreaiartemis.OfferDetails{
	// 								OfferID: to.Ptr("kore-ai-platform"),
	// 								PlanID: to.Ptr("business-monthly"),
	// 								PublisherID: to.Ptr("kore-ai"),
	// 								PlanName: to.Ptr("Business Monthly"),
	// 								TermID: to.Ptr("1m"),
	// 								TermUnit: to.Ptr("P1M"),
	// 							},
	// 							SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Microsoft.SaaS/resources/kore-ai-saas"),
	// 							SubscriptionStatus: to.Ptr(armkoreaiartemis.MarketplaceSubscriptionStatusSubscribed),
	// 						},
	// 					},
	// 				},
	// 				PartnerResourceID: to.Ptr("org_01JK2MXYZ9ABCDEF123456"),
	// 			},
	// 			User: &armkoreaiartemis.UserDetails{
	// 				FirstName: to.Ptr("Jane"),
	// 				LastName: to.Ptr("Doe"),
	// 				EmailAddress: to.Ptr("jane.doe@contoso.com"),
	// 			},
	// 			SingleSignOn: &armkoreaiartemis.SingleSignOnProperties{
	// 				Type: to.Ptr(armkoreaiartemis.SingleSignOnTypeSaml),
	// 				URL: to.Ptr("https://app.kore.ai/sso/login"),
	// 			},
	// 			ProvisioningState: to.Ptr(armkoreaiartemis.ResourceProvisioningStateSucceeded),
	// 		},
	// 		SystemData: &armkoreaiartemis.SystemData{
	// 			CreatedBy: to.Ptr("user@contoso.com"),
	// 			CreatedByType: to.Ptr(armkoreaiartemis.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.May, 18, 8, 0, 0, 0, time.UTC)),
	// 			LastModifiedBy: to.Ptr("user@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armkoreaiartemis.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.May, 18, 8, 0, 0, 0, time.UTC)),
	// 		},
	// 	},
	// }
}
