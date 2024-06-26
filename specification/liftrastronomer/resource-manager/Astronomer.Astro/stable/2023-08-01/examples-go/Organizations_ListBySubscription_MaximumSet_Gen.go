package armastro_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/astro/armastro"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/liftrastronomer/resource-manager/Astronomer.Astro/stable/2023-08-01/examples/Organizations_ListBySubscription_MaximumSet_Gen.json
func ExampleOrganizationsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armastro.NewClientFactory("<subscription-id>", cred, nil)
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
		// page.OrganizationResourceListResult = armastro.OrganizationResourceListResult{
		// 	Value: []*armastro.OrganizationResource{
		// 		{
		// 			Name: to.Ptr("ycyrfvupthkudm"),
		// 			Type: to.Ptr("ldwwclcpqssjomo"),
		// 			ID: to.Ptr("bhslekyvgkfomahtvjiin"),
		// 			SystemData: &armastro.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-25T11:16:12.868Z"); return t}()),
		// 				CreatedBy: to.Ptr("zw"),
		// 				CreatedByType: to.Ptr(armastro.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-25T11:16:12.868Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("isirkhwcppaqoqzoebybzikzbzkjzf"),
		// 				LastModifiedByType: to.Ptr(armastro.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("mhqthlsatwvqkl"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armastro.ManagedServiceIdentity{
		// 				Type: to.Ptr(armastro.ManagedServiceIdentityTypeNone),
		// 				PrincipalID: to.Ptr("b5684bd7-7958-4c0e-9795-d686c31746d2"),
		// 				TenantID: to.Ptr("b5684bd7-7958-4c0e-9795-d686c31746d2"),
		// 				UserAssignedIdentities: map[string]*armastro.UserAssignedIdentity{
		// 				},
		// 			},
		// 			Properties: &armastro.LiftrBaseDataOrganizationProperties{
		// 				Marketplace: &armastro.LiftrBaseMarketplaceDetails{
		// 					OfferDetails: &armastro.LiftrBaseOfferDetails{
		// 						OfferID: to.Ptr("krzkefmpxztqyusidzgpchfaswuyce"),
		// 						PlanID: to.Ptr("kndxzygsanuiqzwbfbbvoipv"),
		// 						PlanName: to.Ptr("pwqjwlq"),
		// 						PublisherID: to.Ptr("gfsqxygpnerxmvols"),
		// 						TermID: to.Ptr("pwds"),
		// 						TermUnit: to.Ptr("xyygyzcazkuelz"),
		// 					},
		// 					SubscriptionID: to.Ptr("ntthclydlpqmasr"),
		// 					SubscriptionStatus: to.Ptr(armastro.MarketplaceSubscriptionStatusPendingFulfillmentStart),
		// 				},
		// 				PartnerOrganizationProperties: &armastro.LiftrBaseDataPartnerOrganizationProperties{
		// 					OrganizationID: to.Ptr("lskgzdmziusgrsucv"),
		// 					OrganizationName: to.Ptr("3-"),
		// 					SingleSignOnProperties: &armastro.LiftrBaseSingleSignOnProperties{
		// 						AADDomains: []*string{
		// 							to.Ptr("kfbleh")},
		// 							EnterpriseAppID: to.Ptr("mklfypyujwumgwdzae"),
		// 							ProvisioningState: to.Ptr(armastro.ResourceProvisioningStateSucceeded),
		// 							SingleSignOnState: to.Ptr(armastro.SingleSignOnStatesInitial),
		// 							SingleSignOnURL: to.Ptr("ymmtzkyghvinvhgnqlzwrr"),
		// 						},
		// 						WorkspaceID: to.Ptr("vcrupxwpaba"),
		// 						WorkspaceName: to.Ptr("9.:06"),
		// 					},
		// 					ProvisioningState: to.Ptr(armastro.ResourceProvisioningStateSucceeded),
		// 					User: &armastro.LiftrBaseUserDetails{
		// 						EmailAddress: to.Ptr(".K_@e7N-g1.xjqnbPs"),
		// 						FirstName: to.Ptr("nfh"),
		// 						LastName: to.Ptr("lazfbstcccykibvcrxpmglqam"),
		// 						PhoneNumber: to.Ptr("inxkscllh"),
		// 						Upn: to.Ptr("xtutvycpxjrtoftx"),
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}
