package armmanagedservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2022-01-01-preview/examples/GetMarketplaceRegistrationDefinitionsAtTenantScope.json
func ExampleMarketplaceRegistrationDefinitionsWithoutScopeClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedservices.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMarketplaceRegistrationDefinitionsWithoutScopeClient().NewListPager(&armmanagedservices.MarketplaceRegistrationDefinitionsWithoutScopeClientListOptions{Filter: to.Ptr("planIdentifier eq 'publisher.offerIdentifier.planName.version'")})
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
		// page.MarketplaceRegistrationDefinitionList = armmanagedservices.MarketplaceRegistrationDefinitionList{
		// 	Value: []*armmanagedservices.MarketplaceRegistrationDefinition{
		// 		{
		// 			Name: to.Ptr("marketplace-test.test.test-plan.1.0.0"),
		// 			Type: to.Ptr("Microsoft.ManagedServices/marketplaceRegistrationDefinitions"),
		// 			ID: to.Ptr("/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/marketplace-test.test.test-plan.1.0.0"),
		// 			Plan: &armmanagedservices.Plan{
		// 				Name: to.Ptr("test-plan"),
		// 				Product: to.Ptr("test"),
		// 				Publisher: to.Ptr("marketplace-test"),
		// 				Version: to.Ptr("1.0.0"),
		// 			},
		// 			Properties: &armmanagedservices.MarketplaceRegistrationDefinitionProperties{
		// 				Authorizations: []*armmanagedservices.Authorization{
		// 					{
		// 						PrincipalID: to.Ptr("f98g86a2-4cc4-4e6d-ad47-b3e80a1bcdfc"),
		// 						PrincipalIDDisplayName: to.Ptr("Support User"),
		// 						RoleDefinitionID: to.Ptr("acdd72a7-3385-48ef-bd42-f606fba81ae7"),
		// 					},
		// 					{
		// 						DelegatedRoleDefinitionIDs: []*string{
		// 							to.Ptr("b24988ac-6180-42a0-ab88-20f7382dd24c")},
		// 							PrincipalID: to.Ptr("f98d86a2-4cc4-4e9d-ad47-b3e80a1bcdfc"),
		// 							PrincipalIDDisplayName: to.Ptr("User Access Administrator"),
		// 							RoleDefinitionID: to.Ptr("18d7d88d-d35e-4fb5-a5c3-7773c20a72d9"),
		// 					}},
		// 					EligibleAuthorizations: []*armmanagedservices.EligibleAuthorization{
		// 					},
		// 					ManagedByTenantID: to.Ptr("83ace5cd-bcc3-441a-hd86-e6a75360cecc"),
		// 					OfferDisplayName: to.Ptr("Marketplace Test Offer"),
		// 					PlanDisplayName: to.Ptr("Test Plan"),
		// 					PublisherDisplayName: to.Ptr("Marketplace Test Publisher"),
		// 				},
		// 		}},
		// 	}
	}
}
