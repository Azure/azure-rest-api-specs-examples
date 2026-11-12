package armkoreaiartemis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/koreaiartemis/armkoreaiartemis"
)

// Generated from example definition: 2026-07-01-preview/Organizations_ListBySubscription_MaximumSet_Gen.json
func ExampleOrganizationsClient_NewListBySubscriptionPager_organizationsListBySubscriptionMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkoreaiartemis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOrganizationsClient().NewListBySubscriptionPager(nil)
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
		// page = armkoreaiartemis.OrganizationsClientListBySubscriptionResponse{
		// 	OrganizationResourceListResult: armkoreaiartemis.OrganizationResourceListResult{
		// 		Value: []*armkoreaiartemis.OrganizationResource{
		// 			{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Kore.AgentPlatform/organizations/contoso-org"),
		// 				Name: to.Ptr("contoso-org"),
		// 				Type: to.Ptr("Kore.AgentPlatform/organizations"),
		// 				Location: to.Ptr("eastus2euap"),
		// 				Properties: &armkoreaiartemis.OrganizationProperties{
		// 					PartnerProperties: &armkoreaiartemis.OrganizationPartnerProperties{
		// 						OrganizationName: to.Ptr("Contoso AI Platform"),
		// 						BillingEmail: to.Ptr("billing@contoso.com"),
		// 						OwnerEmail: to.Ptr("admin@contoso.com"),
		// 						DefaultWorkspace: &armkoreaiartemis.DefaultWorkspaceProperties{
		// 							Name: to.Ptr("my first workspace"),
		// 							MarketplaceDetails: []*armkoreaiartemis.MarketplaceDetails{
		// 								{
		// 									OfferDetails: &armkoreaiartemis.OfferDetails{
		// 										OfferID: to.Ptr("kore-ai-platform"),
		// 										PlanID: to.Ptr("business-monthly"),
		// 										PublisherID: to.Ptr("kore-ai"),
		// 										PlanName: to.Ptr("Business Monthly"),
		// 										TermID: to.Ptr("1m"),
		// 										TermUnit: to.Ptr("P1M"),
		// 									},
		// 									SubscriptionID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Microsoft.SaaS/resources/kore-ai-saas"),
		// 									SubscriptionStatus: to.Ptr(armkoreaiartemis.MarketplaceSubscriptionStatusSubscribed),
		// 								},
		// 							},
		// 						},
		// 						PartnerResourceID: to.Ptr("org_01JK2MXYZ9ABCDEF123456"),
		// 					},
		// 					ProvisioningState: to.Ptr(armkoreaiartemis.ResourceProvisioningStateSucceeded),
		// 					User: &armkoreaiartemis.UserDetails{
		// 						FirstName: to.Ptr("Jane"),
		// 						LastName: to.Ptr("Doe"),
		// 						EmailAddress: to.Ptr("jane.alternate@contoso.com"),
		// 						AlternateEmail: to.Ptr("jane.alternate@contoso.com"),
		// 						Upn: to.Ptr("jane.doe@contoso.onmicrosoft.com"),
		// 						Address: to.Ptr("1 Microsoft Way, Redmond, WA 98052"),
		// 						PhoneNumber: to.Ptr("+1-425-555-0100"),
		// 					},
		// 					CompanyDetails: &armkoreaiartemis.CompanyDetails{
		// 						CompanyName: to.Ptr("Contoso Ltd."),
		// 						OfficeAddress: to.Ptr("1 Microsoft Way, Redmond, WA 98052"),
		// 						Country: to.Ptr("US"),
		// 						Domain: to.Ptr("contoso.com"),
		// 						BusinessType: to.Ptr("Technology"),
		// 						NumberOfEmployees: to.Ptr[int32](28),
		// 					},
		// 					SingleSignOn: &armkoreaiartemis.SingleSignOnProperties{
		// 						Type: to.Ptr(armkoreaiartemis.SingleSignOnTypeSaml),
		// 						State: to.Ptr(armkoreaiartemis.SingleSignOnStatesEnable),
		// 						EnterpriseAppID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 						URL: to.Ptr("https://sso.kore.ai/contoso"),
		// 						AADDomains: []*string{
		// 							to.Ptr("contoso.com"),
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"environment": to.Ptr("production"),
		// 				},
		// 				SystemData: &armkoreaiartemis.SystemData{
		// 					CreatedBy: to.Ptr("jane.doe@contoso.com"),
		// 					CreatedByType: to.Ptr(armkoreaiartemis.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(time.Date(2026, time.May, 25, 5, 38, 33, 975000000, time.UTC)),
		// 					LastModifiedBy: to.Ptr("jane.doe@contoso.com"),
		// 					LastModifiedByType: to.Ptr(armkoreaiartemis.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(time.Date(2026, time.May, 25, 5, 38, 33, 975000000, time.UTC)),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/ahcxjr"),
		// 	},
		// }
	}
}
