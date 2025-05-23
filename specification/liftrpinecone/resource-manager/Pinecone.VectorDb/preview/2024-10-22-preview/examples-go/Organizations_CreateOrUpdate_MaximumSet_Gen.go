package armpineconevectordb_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/pineconevectordb/armpineconevectordb"
)

// Generated from example definition: 2024-10-22-preview/Organizations_CreateOrUpdate_MaximumSet_Gen.json
func ExampleOrganizationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpineconevectordb.NewClientFactory("76a38ef6-c8c1-4f0d-bfe0-00ec782c8077", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOrganizationsClient().BeginCreateOrUpdate(ctx, "rgopenapi", "example-organization-name", armpineconevectordb.OrganizationResource{
		Properties: &armpineconevectordb.OrganizationProperties{
			Marketplace: &armpineconevectordb.MarketplaceDetails{
				SubscriptionID:     to.Ptr("76a38ef6-c8c1-4f0d-bfe0-00ec782c8077"),
				SubscriptionStatus: to.Ptr(armpineconevectordb.MarketplaceSubscriptionStatusPendingFulfillmentStart),
				OfferDetails: &armpineconevectordb.OfferDetails{
					PublisherID: to.Ptr("4d194daf-fa20-46a8-bfb4-5b7d96cae009"),
					OfferID:     to.Ptr("013124d0-bf05-4eab-a6bb-01fa83870642"),
					PlanID:      to.Ptr("62dda065-5acd-4ac5-b418-8610beed92a2"),
					PlanName:    to.Ptr("Freemium"),
					TermUnit:    to.Ptr("der"),
					TermID:      to.Ptr("a2b7ce01-f06d-4874-9f77-6ea4a4875c16"),
				},
			},
			User: &armpineconevectordb.UserDetails{
				FirstName:    to.Ptr("Jimmy"),
				LastName:     to.Ptr("McExample"),
				EmailAddress: to.Ptr("example.user@example.com"),
				Upn:          to.Ptr("example.user@example.com"),
				PhoneNumber:  to.Ptr("555-555-5555"),
			},
			PartnerProperties: &armpineconevectordb.PartnerProperties{
				DisplayName: to.Ptr("My Example Organization"),
			},
			SingleSignOnProperties: &armpineconevectordb.SingleSignOnPropertiesV2{
				Type:            to.Ptr(armpineconevectordb.SingleSignOnTypeSaml),
				State:           to.Ptr(armpineconevectordb.SingleSignOnStatesInitial),
				EnterpriseAppID: to.Ptr("44d3fb26-d8d5-41ff-9b9a-769737f22f13"),
				URL:             to.Ptr("https://login.pinecone.io/?sso=true&connection=dfwgsqzkbrjqrglcsa"),
				AADDomains: []*string{
					to.Ptr("exampledomain"),
				},
			},
		},
		Identity: &armpineconevectordb.ManagedServiceIdentity{
			Type: to.Ptr(armpineconevectordb.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armpineconevectordb.UserAssignedIdentity{
				"ident904655400": {},
			},
		},
		Tags: map[string]*string{
			"my-tag": to.Ptr("tag.value"),
		},
		Location: to.Ptr("us-east"),
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
	// res = armpineconevectordb.OrganizationsClientCreateOrUpdateResponse{
	// 	OrganizationResource: &armpineconevectordb.OrganizationResource{
	// 		Properties: &armpineconevectordb.OrganizationProperties{
	// 			Marketplace: &armpineconevectordb.MarketplaceDetails{
	// 				SubscriptionID: to.Ptr("76a38ef6-c8c1-4f0d-bfe0-00ec782c8077"),
	// 				SubscriptionStatus: to.Ptr(armpineconevectordb.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 				OfferDetails: &armpineconevectordb.OfferDetails{
	// 					PublisherID: to.Ptr("4d194daf-fa20-46a8-bfb4-5b7d96cae009"),
	// 					OfferID: to.Ptr("013124d0-bf05-4eab-a6bb-01fa83870642"),
	// 					PlanID: to.Ptr("62dda065-5acd-4ac5-b418-8610beed92a2"),
	// 					PlanName: to.Ptr("Freemium"),
	// 					TermUnit: to.Ptr("der"),
	// 					TermID: to.Ptr("a2b7ce01-f06d-4874-9f77-6ea4a4875c16"),
	// 				},
	// 			},
	// 			User: &armpineconevectordb.UserDetails{
	// 				FirstName: to.Ptr("Jimmy"),
	// 				LastName: to.Ptr("McExample"),
	// 				EmailAddress: to.Ptr("example.user@example.com"),
	// 				Upn: to.Ptr("example.user@example.com"),
	// 				PhoneNumber: to.Ptr("555-555-5555"),
	// 			},
	// 			ProvisioningState: to.Ptr(armpineconevectordb.ResourceProvisioningStateSucceeded),
	// 			PartnerProperties: &armpineconevectordb.PartnerProperties{
	// 				DisplayName: to.Ptr("My Example Organization"),
	// 			},
	// 			SingleSignOnProperties: &armpineconevectordb.SingleSignOnPropertiesV2{
	// 				Type: to.Ptr(armpineconevectordb.SingleSignOnTypeSaml),
	// 				State: to.Ptr(armpineconevectordb.SingleSignOnStatesInitial),
	// 				EnterpriseAppID: to.Ptr("44d3fb26-d8d5-41ff-9b9a-769737f22f13"),
	// 				URL: to.Ptr("https://login.pinecone.io/?sso=true&connection=dfwgsqzkbrjqrglcsa"),
	// 				AADDomains: []*string{
	// 					to.Ptr("exampledomain"),
	// 				},
	// 			},
	// 		},
	// 		Identity: &armpineconevectordb.ManagedServiceIdentity{
	// 			PrincipalID: to.Ptr("98b0f658-9ee0-4376-9f04-cf211154716c"),
	// 			TenantID: to.Ptr("b118e9e0-4179-41b7-a978-a205cd312e56"),
	// 			Type: to.Ptr(armpineconevectordb.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armpineconevectordb.UserAssignedIdentity{
	// 				"ident904655400": &armpineconevectordb.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("e6beb013-1d4f-4bc7-b6be-33f05f942ae2"),
	// 					ClientID: to.Ptr("7d5610d7-5c8c-4360-98e5-1ae67a4b6ebe"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"my-tag": to.Ptr("tag.value"),
	// 		},
	// 		Location: to.Ptr("us-east"),
	// 		ID: to.Ptr("/subscriptions/76a38ef6-c8c1-4f0d-bfe0-00ec782c8077/resourceGroups/rgopenapi/providers/Pinecone/organizations/example-organization-name"),
	// 		Name: to.Ptr("example-organization-name"),
	// 		Type: to.Ptr("Pinecone.Management/organization"),
	// 		SystemData: &armpineconevectordb.SystemData{
	// 			CreatedBy: to.Ptr("a8006d37-bf85-4ab3-bf9d-2fb4702a1cfb"),
	// 			CreatedByType: to.Ptr(armpineconevectordb.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-22T20:59:36.290Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("a8006d37-bf85-4ab3-bf9d-2fb4702a1cfb"),
	// 			LastModifiedByType: to.Ptr(armpineconevectordb.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-10-22T20:59:36.290Z"); return t}()),
	// 		},
	// 	},
	// }
}
