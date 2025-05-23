package armneonpostgres_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/neonpostgres/armneonpostgres"
)

// Generated from example definition: 2024-08-01-preview/Organizations_Get_MaximumSet_Gen.json
func ExampleOrganizationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armneonpostgres.NewClientFactory("1178323D-8270-4757-B639-D528B6266487", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationsClient().Get(ctx, "rgneon", "5", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armneonpostgres.OrganizationsClientGetResponse{
	// 	OrganizationResource: &armneonpostgres.OrganizationResource{
	// 		Properties: &armneonpostgres.OrganizationProperties{
	// 			MarketplaceDetails: &armneonpostgres.MarketplaceDetails{
	// 				SubscriptionID: to.Ptr("yxmkfivp"),
	// 				SubscriptionStatus: to.Ptr(armneonpostgres.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 				OfferDetails: &armneonpostgres.OfferDetails{
	// 					PublisherID: to.Ptr("hporaxnopmolttlnkbarw"),
	// 					OfferID: to.Ptr("bunyeeupoedueofwrzej"),
	// 					PlanID: to.Ptr("nlbfiwtslenfwek"),
	// 					PlanName: to.Ptr("ljbmgpkfqklaufacbpml"),
	// 					TermUnit: to.Ptr("qbcq"),
	// 					TermID: to.Ptr("aedlchikwqckuploswthvshe"),
	// 				},
	// 			},
	// 			UserDetails: &armneonpostgres.UserDetails{
	// 				FirstName: to.Ptr("buwwe"),
	// 				LastName: to.Ptr("escynjpynkoox"),
	// 				EmailAddress: to.Ptr("3i_%@w8-y.H-p.tvj.dG"),
	// 				Upn: to.Ptr("fwedjamgwwrotcjaucuzdwycfjdqn"),
	// 				PhoneNumber: to.Ptr("dlrqoowumy"),
	// 			},
	// 			CompanyDetails: &armneonpostgres.CompanyDetails{
	// 				CompanyName: to.Ptr("uxn"),
	// 				Country: to.Ptr("lpajqzptqchuko"),
	// 				OfficeAddress: to.Ptr("chpkrlpmfslmawgunjxdllzcrctykq"),
	// 				BusinessPhone: to.Ptr("hbeb"),
	// 				Domain: to.Ptr("krjldeakhwiepvs"),
	// 				NumberOfEmployees: to.Ptr[int64](23),
	// 			},
	// 			ProvisioningState: to.Ptr(armneonpostgres.ResourceProvisioningStateSucceeded),
	// 			PartnerOrganizationProperties: &armneonpostgres.PartnerOrganizationProperties{
	// 				OrganizationID: to.Ptr("nrhvoqzulowcunhmvwfgjcaibvwcl"),
	// 				OrganizationName: to.Ptr("2__.-"),
	// 				SingleSignOnProperties: &armneonpostgres.SingleSignOnProperties{
	// 					SingleSignOnState: to.Ptr(armneonpostgres.SingleSignOnStatesInitial),
	// 					EnterpriseAppID: to.Ptr("fpibacregjfncfdsojs"),
	// 					SingleSignOnURL: to.Ptr("tmojh"),
	// 					AADDomains: []*string{
	// 						to.Ptr("kndszgrwzbvvlssvkej"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key2099": to.Ptr("omjjymaqtrqzksxszhzgyl"),
	// 		},
	// 		Location: to.Ptr("upxxgikyqrbnv"),
	// 		ID: to.Ptr("/subscriptions/1178323D-8270-4757-B639-D528B6266487/resourceGroups/rgneon/providers/Microsoft.Neon/organizations/2__.-"),
	// 		Name: to.Ptr("grhdqtqnkqmu"),
	// 		Type: to.Ptr("gapeymltyvlqlvpgdgfxidkkd"),
	// 		SystemData: &armneonpostgres.SystemData{
	// 			CreatedBy: to.Ptr("qfhekdgpvdtqcohjhvlyhzd"),
	// 			CreatedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-30T15:12:24.902Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("dqsjroejrtucfjyqcoonpdopfaa"),
	// 			LastModifiedByType: to.Ptr(armneonpostgres.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-07-30T15:12:24.902Z"); return t}()),
	// 		},
	// 	},
	// }
}
