package armkoreaiartemis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/koreaiartemis/armkoreaiartemis"
)

// Generated from example definition: 2026-07-01-preview/Organizations_Create_MaximumSet_Gen.json
func ExampleOrganizationsClient_BeginCreate_organizationsCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkoreaiartemis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOrganizationsClient().BeginCreate(ctx, "contoso-rg", "contoso-org", armkoreaiartemis.OrganizationResource{
		Location: to.Ptr("eastus2euap"),
		Tags: map[string]*string{
			"environment": to.Ptr("production"),
		},
		Properties: &armkoreaiartemis.OrganizationProperties{
			PartnerProperties: &armkoreaiartemis.OrganizationPartnerProperties{
				OrganizationName: to.Ptr("Contoso AI Platform"),
				BillingEmail:     to.Ptr("billing@contoso.com"),
				OwnerEmail:       to.Ptr("admin@contoso.com"),
				DefaultWorkspace: &armkoreaiartemis.DefaultWorkspaceProperties{
					Name:    to.Ptr("my first workspace"),
					Product: to.Ptr(armkoreaiartemis.ProductAiforwork),
					MarketplaceDetails: []*armkoreaiartemis.MarketplaceDetails{
						{
							OfferDetails: &armkoreaiartemis.OfferDetails{
								OfferID:     to.Ptr("kore-ai-platform"),
								PlanID:      to.Ptr("business-monthly"),
								PublisherID: to.Ptr("kore-ai"),
								PlanName:    to.Ptr("Business Monthly"),
								TermID:      to.Ptr("1m"),
								TermUnit:    to.Ptr("P1M"),
							},
							SubscriptionID:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Microsoft.SaaS/resources/kore-ai-saas"),
							SubscriptionStatus: to.Ptr(armkoreaiartemis.MarketplaceSubscriptionStatusSubscribed),
						},
					},
				},
			},
			User: &armkoreaiartemis.UserDetails{
				FirstName:      to.Ptr("Jane"),
				LastName:       to.Ptr("Doe"),
				EmailAddress:   to.Ptr("jane.doe@contoso.com"),
				Upn:            to.Ptr("jane.doe@contoso.onmicrosoft.com"),
				PhoneNumber:    to.Ptr("+1-425-555-0100"),
				AlternateEmail: to.Ptr("jane.alternate@contoso.com"),
				Address:        to.Ptr("1 Microsoft Way, Redmond, WA 98052"),
			},
			CompanyDetails: &armkoreaiartemis.CompanyDetails{
				CompanyName:       to.Ptr("Contoso Ltd."),
				Country:           to.Ptr("US"),
				Domain:            to.Ptr("contoso.com"),
				BusinessType:      to.Ptr("Technology"),
				NumberOfEmployees: to.Ptr[int32](500),
				OfficeAddress:     to.Ptr("1 Microsoft Way, Redmond, WA 98052"),
			},
			SingleSignOn: &armkoreaiartemis.SingleSignOnProperties{
				Type:            to.Ptr(armkoreaiartemis.SingleSignOnTypeSaml),
				State:           to.Ptr(armkoreaiartemis.SingleSignOnStatesEnable),
				EnterpriseAppID: to.Ptr("00000000-0000-0000-0000-000000000000"),
				URL:             to.Ptr("https://sso.kore.ai/contoso"),
				AADDomains: []*string{
					to.Ptr("contoso.com"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkoreaiartemis.OrganizationsClientCreateResponse{
	// 	OrganizationResource: armkoreaiartemis.OrganizationResource{
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
	// 					Product: to.Ptr(armkoreaiartemis.ProductAiforwork),
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
	// 				Upn: to.Ptr("jane.doe@contoso.onmicrosoft.com"),
	// 				PhoneNumber: to.Ptr("+1-425-555-0100"),
	// 				AlternateEmail: to.Ptr("jane.alternate@contoso.com"),
	// 				Address: to.Ptr("1 Microsoft Way, Redmond, WA 98052"),
	// 			},
	// 			CompanyDetails: &armkoreaiartemis.CompanyDetails{
	// 				CompanyName: to.Ptr("Contoso Ltd."),
	// 				Country: to.Ptr("US"),
	// 				Domain: to.Ptr("contoso.com"),
	// 				BusinessType: to.Ptr("Technology"),
	// 				NumberOfEmployees: to.Ptr[int32](500),
	// 				OfficeAddress: to.Ptr("1 Microsoft Way, Redmond, WA 98052"),
	// 			},
	// 			SingleSignOn: &armkoreaiartemis.SingleSignOnProperties{
	// 				Type: to.Ptr(armkoreaiartemis.SingleSignOnTypeSaml),
	// 				State: to.Ptr(armkoreaiartemis.SingleSignOnStatesEnable),
	// 				EnterpriseAppID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				URL: to.Ptr("https://sso.kore.ai/contoso"),
	// 				AADDomains: []*string{
	// 					to.Ptr("contoso.com"),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armkoreaiartemis.ResourceProvisioningStateSucceeded),
	// 		},
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Kore.AgentPlatform/organizations/contoso-org"),
	// 		Name: to.Ptr("contoso-org"),
	// 		Type: to.Ptr("Kore.AgentPlatform/organizations"),
	// 		SystemData: &armkoreaiartemis.SystemData{
	// 			CreatedBy: to.Ptr("jane.doe@contoso.com"),
	// 			CreatedByType: to.Ptr(armkoreaiartemis.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(time.Date(2026, time.May, 25, 5, 38, 33, 975000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("jane.doe@contoso.com"),
	// 			LastModifiedByType: to.Ptr(armkoreaiartemis.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2026, time.May, 25, 5, 38, 33, 975000000, time.UTC)),
	// 		},
	// 	},
	// }
}
