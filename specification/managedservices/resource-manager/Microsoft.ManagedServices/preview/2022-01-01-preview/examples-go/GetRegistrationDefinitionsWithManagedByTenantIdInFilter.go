package armmanagedservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2022-01-01-preview/examples/GetRegistrationDefinitionsWithManagedByTenantIdInFilter.json
func ExampleRegistrationDefinitionsClient_NewListPager_getRegistrationDefinitionsWithManagedByTenantIdInFilter() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedservices.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRegistrationDefinitionsClient().NewListPager("subscription/0afefe50-734e-4610-8a82-a144ahf49dea", &armmanagedservices.RegistrationDefinitionsClientListOptions{Filter: to.Ptr("$filter=managedByTenantId in (83ace5cd-bcc3-441a-hd86-e6a75360cecc, de83f4a9-a76a-4025-a91a-91171923eac7)")})
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
		// page.RegistrationDefinitionList = armmanagedservices.RegistrationDefinitionList{
		// 	Value: []*armmanagedservices.RegistrationDefinition{
		// 		{
		// 			Name: to.Ptr("26c128c2-fefa-4340-9bb1-6e081c90ada2"),
		// 			Type: to.Ptr("Microsoft.ManagedServices/registrationDefinitions"),
		// 			ID: to.Ptr("/subscriptions/0afefe50-734e-4610-8a82-a144ahf49dea/providers/Microsoft.ManagedServices/registrationDefinitions/26c128c2-fefa-4340-9bb1-6e081c90ada2"),
		// 			Plan: &armmanagedservices.Plan{
		// 				Name: to.Ptr("addesai-plan"),
		// 				Product: to.Ptr("test"),
		// 				Publisher: to.Ptr("marketplace-test"),
		// 				Version: to.Ptr("1.0.0"),
		// 			},
		// 			Properties: &armmanagedservices.RegistrationDefinitionProperties{
		// 				Description: to.Ptr("Test"),
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
		// 						{
		// 							JustInTimeAccessPolicy: &armmanagedservices.JustInTimeAccessPolicy{
		// 								ManagedByTenantApprovers: []*armmanagedservices.EligibleApprover{
		// 									{
		// 										PrincipalID: to.Ptr("d9b22cd6-6407-43cc-8c60-07c56df0b51a"),
		// 										PrincipalIDDisplayName: to.Ptr("Approver Group"),
		// 								}},
		// 								MaximumActivationDuration: to.Ptr("PT8H"),
		// 								MultiFactorAuthProvider: to.Ptr(armmanagedservices.MultiFactorAuthProviderAzure),
		// 							},
		// 							PrincipalID: to.Ptr("3e0ed8c6-e902-4fc5-863c-e3ddbb2ae2a2"),
		// 							PrincipalIDDisplayName: to.Ptr("Support User"),
		// 							RoleDefinitionID: to.Ptr("ae349356-3a1b-4a5e-921d-050484c6347e"),
		// 					}},
		// 					ManagedByTenantID: to.Ptr("83ace5cd-bcc3-441a-hd86-e6a75360cecc"),
		// 					ManagedByTenantName: to.Ptr("Contoso Corp."),
		// 					ManageeTenantID: to.Ptr("01c0bcd5-4f47-4e4b-b492-418b7e2a8854"),
		// 					ManageeTenantName: to.Ptr("test_test_aad_SbtFhyGiLHPFm"),
		// 					ProvisioningState: to.Ptr(armmanagedservices.ProvisioningStateSucceeded),
		// 					RegistrationDefinitionName: to.Ptr("DefinitionName"),
		// 				},
		// 				SystemData: &armmanagedservices.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-22T19:56:05.070Z"); return t}()),
		// 					CreatedBy: to.Ptr("testuser@outlook.com"),
		// 					CreatedByType: to.Ptr(armmanagedservices.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-22T19:56:05.070Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("testuser@outlook.com"),
		// 					LastModifiedByType: to.Ptr(armmanagedservices.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("2a60751f-03d6-45a1-8797-24678246d54d"),
		// 				Type: to.Ptr("Microsoft.ManagedServices/registrationDefinitions"),
		// 				ID: to.Ptr("/subscriptions/0afefe50-734e-4610-8a82-a144ahf49dea/providers/Microsoft.ManagedServices/registrationDefinitions/2a60751f-03d6-45a1-8797-24678246d54d"),
		// 				Properties: &armmanagedservices.RegistrationDefinitionProperties{
		// 					Description: to.Ptr("Test 2"),
		// 					Authorizations: []*armmanagedservices.Authorization{
		// 						{
		// 							PrincipalID: to.Ptr("a2e38003-c234-42dc-a769-3ca55be53098"),
		// 							PrincipalIDDisplayName: to.Ptr("Support User"),
		// 							RoleDefinitionID: to.Ptr("acdd72a7-3385-48ef-bd42-f606fba81ae7"),
		// 					}},
		// 					ManagedByTenantID: to.Ptr("de83f4a9-a76a-4025-a91a-91171923eac7"),
		// 					ManagedByTenantName: to.Ptr("Test Corp."),
		// 					ManageeTenantID: to.Ptr("01c0bcd5-4f47-4e4b-b492-418b7e2a8854"),
		// 					ManageeTenantName: to.Ptr("test_test_aad_SbtFhyGiLHPFm"),
		// 					ProvisioningState: to.Ptr(armmanagedservices.ProvisioningStateSucceeded),
		// 					RegistrationDefinitionName: to.Ptr("DefinitionName"),
		// 				},
		// 				SystemData: &armmanagedservices.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-21T08:56:05.070Z"); return t}()),
		// 					CreatedBy: to.Ptr("testuser@msp.com"),
		// 					CreatedByType: to.Ptr(armmanagedservices.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-11-22T08:56:05.070Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("testuser@msp.com"),
		// 					LastModifiedByType: to.Ptr(armmanagedservices.CreatedByTypeUser),
		// 				},
		// 		}},
		// 	}
	}
}
