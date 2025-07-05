package armmongodbatlas_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongodbatlas/armmongodbatlas"
)

// Generated from example definition: 2025-06-01/Organizations_CreateOrUpdate_MaximumSet_Gen.json
func ExampleOrganizationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongodbatlas.NewClientFactory("AD4FE133-6EF1-4ED8-82DB-5C1CBA58597E", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOrganizationsClient().BeginCreateOrUpdate(ctx, "rgopenapi", "U.1-:7", armmongodbatlas.OrganizationResource{
		Properties: &armmongodbatlas.OrganizationProperties{
			Marketplace: &armmongodbatlas.MarketplaceDetails{
				SubscriptionID:     to.Ptr("o"),
				SubscriptionStatus: to.Ptr(armmongodbatlas.MarketplaceSubscriptionStatusPendingFulfillmentStart),
				OfferDetails: &armmongodbatlas.OfferDetails{
					PublisherID: to.Ptr("rxglearenxsgpwzlsxmiicynks"),
					OfferID:     to.Ptr("ohnquleylybvjrtnpjupvwlk"),
					PlanID:      to.Ptr("obhxnhvrtbcnoovgofbs"),
					PlanName:    to.Ptr("lkwdzpfhvjezjusrqzyftcikxdt"),
					TermUnit:    to.Ptr("omkxrnburbnruglwqgjlahvjmbfcse"),
					TermID:      to.Ptr("bqmmltwmtpdcdeszbka"),
				},
			},
			User: &armmongodbatlas.UserDetails{
				FirstName:    to.Ptr("aslybvdwwddqxwazxvxhjrs"),
				LastName:     to.Ptr("cnuitqoqpcyvmuqowgnxpwxjcveyr"),
				EmailAddress: to.Ptr(".K_@e7N-g1.xjqnbPs"),
				Upn:          to.Ptr("howdzmfy"),
				PhoneNumber:  to.Ptr("ilypntsrbmbbbexbasuu"),
				CompanyName:  to.Ptr("oxdcwwl"),
			},
			PartnerProperties: &armmongodbatlas.PartnerProperties{
				OrganizationID:   to.Ptr("lyombjlhvwxithkiy"),
				RedirectURL:      to.Ptr("cbxwtehraetlluocdihfgchvjzockn"),
				OrganizationName: to.Ptr("U.1-:7"),
			},
		},
		Identity: &armmongodbatlas.ManagedServiceIdentity{
			Type:                   to.Ptr(armmongodbatlas.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armmongodbatlas.UserAssignedIdentity{},
		},
		Tags:     map[string]*string{},
		Location: to.Ptr("wobqn"),
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
	// res = armmongodbatlas.OrganizationsClientCreateOrUpdateResponse{
	// 	OrganizationResource: &armmongodbatlas.OrganizationResource{
	// 		Properties: &armmongodbatlas.OrganizationProperties{
	// 			Marketplace: &armmongodbatlas.MarketplaceDetails{
	// 				SubscriptionID: to.Ptr("o"),
	// 				SubscriptionStatus: to.Ptr(armmongodbatlas.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 				OfferDetails: &armmongodbatlas.OfferDetails{
	// 					PublisherID: to.Ptr("rxglearenxsgpwzlsxmiicynks"),
	// 					OfferID: to.Ptr("ohnquleylybvjrtnpjupvwlk"),
	// 					PlanID: to.Ptr("obhxnhvrtbcnoovgofbs"),
	// 					PlanName: to.Ptr("lkwdzpfhvjezjusrqzyftcikxdt"),
	// 					TermUnit: to.Ptr("omkxrnburbnruglwqgjlahvjmbfcse"),
	// 					TermID: to.Ptr("bqmmltwmtpdcdeszbka"),
	// 				},
	// 			},
	// 			User: &armmongodbatlas.UserDetails{
	// 				FirstName: to.Ptr("aslybvdwwddqxwazxvxhjrs"),
	// 				LastName: to.Ptr("cnuitqoqpcyvmuqowgnxpwxjcveyr"),
	// 				EmailAddress: to.Ptr(".K_@e7N-g1.xjqnbPs"),
	// 				Upn: to.Ptr("howdzmfy"),
	// 				PhoneNumber: to.Ptr("ilypntsrbmbbbexbasuu"),
	// 				CompanyName: to.Ptr("oxdcwwl"),
	// 			},
	// 			ProvisioningState: to.Ptr(armmongodbatlas.ResourceProvisioningStateSucceeded),
	// 			PartnerProperties: &armmongodbatlas.PartnerProperties{
	// 				OrganizationID: to.Ptr("lpmdubvuizsthw"),
	// 				RedirectURL: to.Ptr("cbxwtehraetlluocdihfgchvjzockn"),
	// 				OrganizationName: to.Ptr("U.1-:7"),
	// 			},
	// 		},
	// 		Identity: &armmongodbatlas.ManagedServiceIdentity{
	// 			PrincipalID: to.Ptr("123e4567-e89b-12d3-a456-426614174000"),
	// 			TenantID: to.Ptr("123e4567-e89b-12d3-a456-426614174000"),
	// 			Type: to.Ptr(armmongodbatlas.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armmongodbatlas.UserAssignedIdentity{
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key8183": to.Ptr("ibdiopeuhthfvnjjbtkwipooa"),
	// 		},
	// 		Location: to.Ptr("wobqn"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/MongoDB.Atlas/organizations/testorg"),
	// 		Name: to.Ptr("qcdj"),
	// 		Type: to.Ptr("pfmxfnhzmwmnmbdbluynkxomgwtqz"),
	// 		SystemData: &armmongodbatlas.SystemData{
	// 			CreatedBy: to.Ptr("rwy"),
	// 			CreatedByType: to.Ptr(armmongodbatlas.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-04T04:51:37.109Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ceflhxjkjslg"),
	// 			LastModifiedByType: to.Ptr(armmongodbatlas.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-12-04T04:51:37.109Z"); return t}()),
	// 		},
	// 	},
	// }
}
