package armarizeaiobservabilityeval_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/arizeaiobservabilityeval/armarizeaiobservabilityeval"
)

// Generated from example definition: 2024-10-01/Organizations_CreateOrUpdate_MaximumSet_Gen.json
func ExampleOrganizationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armarizeaiobservabilityeval.NewClientFactory("4DEBE8B4-8BA4-42F8-AE50-FBEF318751D1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOrganizationsClient().BeginCreateOrUpdate(ctx, "rgopenapi", "test-organization-1", armarizeaiobservabilityeval.OrganizationResource{
		Properties: &armarizeaiobservabilityeval.OrganizationProperties{
			Marketplace: &armarizeaiobservabilityeval.MarketplaceDetails{
				SubscriptionID:     to.Ptr("meaowktoejxwfqomc"),
				SubscriptionStatus: to.Ptr(armarizeaiobservabilityeval.MarketplaceSubscriptionStatusPendingFulfillmentStart),
				OfferDetails: &armarizeaiobservabilityeval.OfferDetails{
					PublisherID: to.Ptr("flrya"),
					OfferID:     to.Ptr("hwhtxmtmmlwsu"),
					PlanID:      to.Ptr("jozklohkdpng"),
					PlanName:    to.Ptr("clnynwt"),
					TermUnit:    to.Ptr("cbfktammjyqewljjjaokakilog"),
					TermID:      to.Ptr("iugvvvoggusxuz"),
				},
			},
			User: &armarizeaiobservabilityeval.UserDetails{
				FirstName:    to.Ptr("aorfffgdmglvzdvfvdyjohtnblzsfw"),
				LastName:     to.Ptr("tojbqzk"),
				EmailAddress: to.Ptr("btables@arize.com"),
				Upn:          to.Ptr("xzvwwbjpqakqqyfudyp"),
				PhoneNumber:  to.Ptr("akbqdbs"),
			},
			PartnerProperties: &armarizeaiobservabilityeval.PartnerProperties{
				Description: to.Ptr("this is a great description"),
			},
			SingleSignOnProperties: &armarizeaiobservabilityeval.SingleSignOnPropertiesV2{
				Type:            to.Ptr(armarizeaiobservabilityeval.SingleSignOnTypeSaml),
				State:           to.Ptr(armarizeaiobservabilityeval.SingleSignOnStatesInitial),
				EnterpriseAppID: to.Ptr("kqykskeuqffsslmpjryzggphhpeh"),
				URL:             to.Ptr("ihidsswbeahnsjjxxqntz"),
				AADDomains: []*string{
					to.Ptr("tyjdvljasl"),
				},
			},
		},
		Identity: &armarizeaiobservabilityeval.ManagedServiceIdentity{
			Type:                   to.Ptr(armarizeaiobservabilityeval.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armarizeaiobservabilityeval.UserAssignedIdentity{},
		},
		Tags:     map[string]*string{},
		Location: to.Ptr("gigxuxdonjfmqnljxcgctfwqapllu"),
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
	// res = armarizeaiobservabilityeval.OrganizationsClientCreateOrUpdateResponse{
	// 	OrganizationResource: &armarizeaiobservabilityeval.OrganizationResource{
	// 		Properties: &armarizeaiobservabilityeval.OrganizationProperties{
	// 			Marketplace: &armarizeaiobservabilityeval.MarketplaceDetails{
	// 				SubscriptionID: to.Ptr("meaowktoejxwfqomc"),
	// 				SubscriptionStatus: to.Ptr(armarizeaiobservabilityeval.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 				OfferDetails: &armarizeaiobservabilityeval.OfferDetails{
	// 					PublisherID: to.Ptr("flrya"),
	// 					OfferID: to.Ptr("hwhtxmtmmlwsu"),
	// 					PlanID: to.Ptr("jozklohkdpng"),
	// 					PlanName: to.Ptr("clnynwt"),
	// 					TermUnit: to.Ptr("cbfktammjyqewljjjaokakilog"),
	// 					TermID: to.Ptr("iugvvvoggusxuz"),
	// 				},
	// 			},
	// 			User: &armarizeaiobservabilityeval.UserDetails{
	// 				FirstName: to.Ptr("aorfffgdmglvzdvfvdyjohtnblzsfw"),
	// 				LastName: to.Ptr("tojbqzk"),
	// 				EmailAddress: to.Ptr("btables@arize.com"),
	// 				Upn: to.Ptr("xzvwwbjpqakqqyfudyp"),
	// 				PhoneNumber: to.Ptr("akbqdbs"),
	// 			},
	// 			PartnerProperties: &armarizeaiobservabilityeval.PartnerProperties{
	// 				Description: to.Ptr("this is a great description"),
	// 			},
	// 			ProvisioningState: to.Ptr(armarizeaiobservabilityeval.ResourceProvisioningStateSucceeded),
	// 			SingleSignOnProperties: &armarizeaiobservabilityeval.SingleSignOnPropertiesV2{
	// 				Type: to.Ptr(armarizeaiobservabilityeval.SingleSignOnTypeSaml),
	// 				State: to.Ptr(armarizeaiobservabilityeval.SingleSignOnStatesInitial),
	// 				EnterpriseAppID: to.Ptr("kqykskeuqffsslmpjryzggphhpeh"),
	// 				URL: to.Ptr("ihidsswbeahnsjjxxqntz"),
	// 				AADDomains: []*string{
	// 					to.Ptr("tyjdvljasl"),
	// 				},
	// 			},
	// 		},
	// 		Identity: &armarizeaiobservabilityeval.ManagedServiceIdentity{
	// 			Type: to.Ptr(armarizeaiobservabilityeval.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armarizeaiobservabilityeval.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("38fdf748-d66d-4344-ac76-ab1ebd9441fc"),
	// 			TenantID: to.Ptr("9aaa003a-d02e-4a03-a904-c7dd89fc588a"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("gigxuxdonjfmqnljxcgctfwqapllu"),
	// 		ID: to.Ptr("/subscriptions/test-subscription-1/resourceGroups/test-resourcegroup-1/providers/ArizeAi.ObservabilityEval/Organizations/test-organization-1"),
	// 		Name: to.Ptr("btywhrchehozcpizfrv"),
	// 		Type: to.Ptr("ymstff"),
	// 		SystemData: &armarizeaiobservabilityeval.SystemData{
	// 			CreatedBy: to.Ptr("ihyykvakpoaltuuwuwgx"),
	// 			CreatedByType: to.Ptr(armarizeaiobservabilityeval.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-29T03:21:12.627Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("mdkcrvmimq"),
	// 			LastModifiedByType: to.Ptr(armarizeaiobservabilityeval.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-29T03:21:12.627Z"); return t}()),
	// 		},
	// 	},
	// }
}
