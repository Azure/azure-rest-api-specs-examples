package armastro_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/astro/armastro"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/liftrastronomer/resource-manager/Astronomer.Astro/stable/2023-08-01/examples/Organizations_Update_MaximumSet_Gen.json
func ExampleOrganizationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armastro.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOrganizationsClient().BeginUpdate(ctx, "rgastronomer", "6.", armastro.OrganizationResourceUpdate{
		Identity: &armastro.ManagedServiceIdentity{
			Type:                   to.Ptr(armastro.ManagedServiceIdentityTypeNone),
			UserAssignedIdentities: map[string]*armastro.UserAssignedIdentity{},
		},
		Properties: &armastro.OrganizationResourceUpdateProperties{
			PartnerOrganizationProperties: &armastro.LiftrBaseDataPartnerOrganizationPropertiesUpdate{
				OrganizationID:   to.Ptr("lrtmbkvyvvoszhjevohkmyjhfyty"),
				OrganizationName: to.Ptr("U2P_"),
				SingleSignOnProperties: &armastro.LiftrBaseSingleSignOnProperties{
					AADDomains: []*string{
						to.Ptr("kfbleh")},
					EnterpriseAppID:   to.Ptr("mklfypyujwumgwdzae"),
					ProvisioningState: to.Ptr(armastro.ResourceProvisioningStateSucceeded),
					SingleSignOnState: to.Ptr(armastro.SingleSignOnStatesInitial),
					SingleSignOnURL:   to.Ptr("ymmtzkyghvinvhgnqlzwrr"),
				},
				WorkspaceID:   to.Ptr("xsepuskdhejaadusyxq"),
				WorkspaceName: to.Ptr("L.-y_--:"),
			},
			User: &armastro.LiftrBaseUserDetailsUpdate{
				EmailAddress: to.Ptr(".K_@e7N-g1.xjqnbPs"),
				FirstName:    to.Ptr("qeuofehzypzljgcuysugefbgxde"),
				LastName:     to.Ptr("g"),
				PhoneNumber:  to.Ptr("aqpyxznvqpgkzohevynofrjdfgoo"),
				Upn:          to.Ptr("uwtprzdfpsqmktx"),
			},
		},
		Tags: map[string]*string{
			"key1474": to.Ptr("bqqyipxnbbxryhznyaosmtpo"),
		},
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
	// res.OrganizationResource = armastro.OrganizationResource{
	// 	Name: to.Ptr("ycyrfvupthkudm"),
	// 	Type: to.Ptr("ldwwclcpqssjomo"),
	// 	ID: to.Ptr("bhslekyvgkfomahtvjiin"),
	// 	SystemData: &armastro.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-25T11:16:12.868Z"); return t}()),
	// 		CreatedBy: to.Ptr("zw"),
	// 		CreatedByType: to.Ptr(armastro.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-25T11:16:12.868Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("isirkhwcppaqoqzoebybzikzbzkjzf"),
	// 		LastModifiedByType: to.Ptr(armastro.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("mhqthlsatwvqkl"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armastro.ManagedServiceIdentity{
	// 		Type: to.Ptr(armastro.ManagedServiceIdentityTypeNone),
	// 		PrincipalID: to.Ptr("b5684bd7-7958-4c0e-9795-d686c31746d2"),
	// 		TenantID: to.Ptr("b5684bd7-7958-4c0e-9795-d686c31746d2"),
	// 		UserAssignedIdentities: map[string]*armastro.UserAssignedIdentity{
	// 		},
	// 	},
	// 	Properties: &armastro.LiftrBaseDataOrganizationProperties{
	// 		Marketplace: &armastro.LiftrBaseMarketplaceDetails{
	// 			OfferDetails: &armastro.LiftrBaseOfferDetails{
	// 				OfferID: to.Ptr("krzkefmpxztqyusidzgpchfaswuyce"),
	// 				PlanID: to.Ptr("kndxzygsanuiqzwbfbbvoipv"),
	// 				PlanName: to.Ptr("pwqjwlq"),
	// 				PublisherID: to.Ptr("gfsqxygpnerxmvols"),
	// 				TermID: to.Ptr("pwds"),
	// 				TermUnit: to.Ptr("xyygyzcazkuelz"),
	// 			},
	// 			SubscriptionID: to.Ptr("ntthclydlpqmasr"),
	// 			SubscriptionStatus: to.Ptr(armastro.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 		},
	// 		PartnerOrganizationProperties: &armastro.LiftrBaseDataPartnerOrganizationProperties{
	// 			OrganizationID: to.Ptr("lrtmbkvyvvoszhjevohkmyjhfyty"),
	// 			OrganizationName: to.Ptr("U2P_"),
	// 			SingleSignOnProperties: &armastro.LiftrBaseSingleSignOnProperties{
	// 				AADDomains: []*string{
	// 					to.Ptr("kfbleh")},
	// 					EnterpriseAppID: to.Ptr("mklfypyujwumgwdzae"),
	// 					ProvisioningState: to.Ptr(armastro.ResourceProvisioningStateSucceeded),
	// 					SingleSignOnState: to.Ptr(armastro.SingleSignOnStatesInitial),
	// 					SingleSignOnURL: to.Ptr("ymmtzkyghvinvhgnqlzwrr"),
	// 				},
	// 				WorkspaceID: to.Ptr("xsepuskdhejaadusyxq"),
	// 				WorkspaceName: to.Ptr("L.-y_--:"),
	// 			},
	// 			ProvisioningState: to.Ptr(armastro.ResourceProvisioningStateSucceeded),
	// 			User: &armastro.LiftrBaseUserDetails{
	// 				EmailAddress: to.Ptr(".K_@e7N-g1.xjqnbPs"),
	// 				FirstName: to.Ptr("qeuofehzypzljgcuysugefbgxde"),
	// 				LastName: to.Ptr("g"),
	// 				PhoneNumber: to.Ptr("aqpyxznvqpgkzohevynofrjdfgoo"),
	// 				Upn: to.Ptr("uwtprzdfpsqmktx"),
	// 			},
	// 		},
	// 	}
}
