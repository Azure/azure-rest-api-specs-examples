package armlambdatesthyperexecute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/lambdatesthyperexecute/armlambdatesthyperexecute"
)

// Generated from example definition: 2024-02-01/Organizations_CreateOrUpdate_MaximumSet_Gen.json
func ExampleOrganizationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlambdatesthyperexecute.NewClientFactory("171E7A75-341B-4472-BC4C-7603C5AB9F32", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOrganizationsClient().BeginCreateOrUpdate(ctx, "rgopenapi", "testorg", armlambdatesthyperexecute.OrganizationResource{
		Properties: &armlambdatesthyperexecute.OrganizationProperties{
			Marketplace: &armlambdatesthyperexecute.MarketplaceDetails{
				SubscriptionID:     to.Ptr("zetdxwryjgcsnosezfpovkpvgvim"),
				SubscriptionStatus: to.Ptr(armlambdatesthyperexecute.MarketplaceSubscriptionStatusPendingFulfillmentStart),
				OfferDetails: &armlambdatesthyperexecute.OfferDetails{
					PublisherID: to.Ptr("ufwwpzit"),
					OfferID:     to.Ptr("fmljkvoivqmfdiwsu"),
					PlanID:      to.Ptr("ssjlabxexw"),
					PlanName:    to.Ptr("mrguqu"),
					TermUnit:    to.Ptr("acvhavsffebfivyaxhxxsaqzt"),
					TermID:      to.Ptr("hxkszxfscsyefeuunyyfskhibr"),
				},
			},
			User: &armlambdatesthyperexecute.UserDetails{
				FirstName:    to.Ptr("ssnzyujsrszbptndzeoqzrmbufrhgp"),
				LastName:     to.Ptr("nsfylyvdyrtfzfeehmji"),
				EmailAddress: to.Ptr("joe@example.com"),
				Upn:          to.Ptr("tfqolz"),
				PhoneNumber:  to.Ptr("jkevskjaaylbwjzofkzmxdysejsoir"),
			},
			PartnerProperties: &armlambdatesthyperexecute.PartnerProperties{
				LicensesSubscribed: to.Ptr[int32](7),
			},
			SingleSignOnProperties: &armlambdatesthyperexecute.SingleSignOnPropertiesV2{
				Type:            to.Ptr(armlambdatesthyperexecute.SingleSignOnTypeSaml),
				State:           to.Ptr(armlambdatesthyperexecute.SingleSignOnStatesInitial),
				EnterpriseAppID: to.Ptr("sonpowym"),
				URL:             to.Ptr("qlshnxrfcdpjcpkxxisrn"),
				AADDomains: []*string{
					to.Ptr("hrgguokssgyrfdzliyjmovtelfu"),
				},
			},
		},
		Identity: &armlambdatesthyperexecute.ManagedServiceIdentity{
			Type:                   to.Ptr(armlambdatesthyperexecute.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armlambdatesthyperexecute.UserAssignedIdentity{},
		},
		Tags:     map[string]*string{},
		Location: to.Ptr("cvymsrlt"),
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
	// res = armlambdatesthyperexecute.OrganizationsClientCreateOrUpdateResponse{
	// 	OrganizationResource: &armlambdatesthyperexecute.OrganizationResource{
	// 		Properties: &armlambdatesthyperexecute.OrganizationProperties{
	// 			Marketplace: &armlambdatesthyperexecute.MarketplaceDetails{
	// 				SubscriptionID: to.Ptr("zetdxwryjgcsnosezfpovkpvgvim"),
	// 				SubscriptionStatus: to.Ptr(armlambdatesthyperexecute.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 				OfferDetails: &armlambdatesthyperexecute.OfferDetails{
	// 					PublisherID: to.Ptr("ufwwpzit"),
	// 					OfferID: to.Ptr("fmljkvoivqmfdiwsu"),
	// 					PlanID: to.Ptr("ssjlabxexw"),
	// 					PlanName: to.Ptr("mrguqu"),
	// 					TermUnit: to.Ptr("acvhavsffebfivyaxhxxsaqzt"),
	// 					TermID: to.Ptr("hxkszxfscsyefeuunyyfskhibr"),
	// 				},
	// 			},
	// 			User: &armlambdatesthyperexecute.UserDetails{
	// 				FirstName: to.Ptr("ssnzyujsrszbptndzeoqzrmbufrhgp"),
	// 				LastName: to.Ptr("nsfylyvdyrtfzfeehmji"),
	// 				EmailAddress: to.Ptr("joe@example.com"),
	// 				Upn: to.Ptr("tfqolz"),
	// 				PhoneNumber: to.Ptr("jkevskjaaylbwjzofkzmxdysejsoir"),
	// 			},
	// 			PartnerProperties: &armlambdatesthyperexecute.PartnerProperties{
	// 				LicensesSubscribed: to.Ptr[int32](7),
	// 			},
	// 			ProvisioningState: to.Ptr(armlambdatesthyperexecute.ResourceProvisioningStateSucceeded),
	// 			SingleSignOnProperties: &armlambdatesthyperexecute.SingleSignOnPropertiesV2{
	// 				Type: to.Ptr(armlambdatesthyperexecute.SingleSignOnTypeSaml),
	// 				State: to.Ptr(armlambdatesthyperexecute.SingleSignOnStatesInitial),
	// 				EnterpriseAppID: to.Ptr("sonpowym"),
	// 				URL: to.Ptr("qlshnxrfcdpjcpkxxisrn"),
	// 				AADDomains: []*string{
	// 					to.Ptr("hrgguokssgyrfdzliyjmovtelfu"),
	// 				},
	// 			},
	// 		},
	// 		Identity: &armlambdatesthyperexecute.ManagedServiceIdentity{
	// 			Type: to.Ptr(armlambdatesthyperexecute.ManagedServiceIdentityTypeNone),
	// 			UserAssignedIdentities: map[string]*armlambdatesthyperexecute.UserAssignedIdentity{
	// 			},
	// 			PrincipalID: to.Ptr("123e4567-e89b-12d3-a456-426614174000"),
	// 			TenantID: to.Ptr("123e4567-e89b-12d3-a456-426614174000"),
	// 		},
	// 		Tags: map[string]*string{
	// 		},
	// 		Location: to.Ptr("cvymsrlt"),
	// 		ID: to.Ptr("/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/myResourceGroup/providers/Lambdatest.Hyoerexecute/organizations/testorg"),
	// 		Name: to.Ptr("rayjyvlsggeccbyzkstzr"),
	// 		Type: to.Ptr("ju"),
	// 		SystemData: &armlambdatesthyperexecute.SystemData{
	// 			CreatedBy: to.Ptr("muialblsfdrvcxxwenlybddkar"),
	// 			CreatedByType: to.Ptr(armlambdatesthyperexecute.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-05T12:10:42.989Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("qnkeaiefwrcpgihdwkesrjw"),
	// 			LastModifiedByType: to.Ptr(armlambdatesthyperexecute.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-09-05T12:10:42.989Z"); return t}()),
	// 		},
	// 	},
	// }
}
