package armmanagedservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2022-01-01-preview/examples/GetRegistrationDefinition.json
func ExampleRegistrationDefinitionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedservices.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRegistrationDefinitionsClient().Get(ctx, "subscription/0afefe50-734e-4610-8a82-a144ahf49dea", "26c128c2-fefa-4340-9bb1-6e081c90ada2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RegistrationDefinition = armmanagedservices.RegistrationDefinition{
	// 	Name: to.Ptr("26c128c2-fefa-4340-9bb1-6e081c90ada2"),
	// 	Type: to.Ptr("Microsoft.ManagedServices/registrationDefinitions"),
	// 	ID: to.Ptr("/subscriptions/0afefe50-734e-4610-8a82-a144ahf49dea/providers/Microsoft.ManagedServices/registrationDefinitions/26c128c2-fefa-4340-9bb1-6e081c90ada2"),
	// 	Plan: &armmanagedservices.Plan{
	// 		Name: to.Ptr("addesai-plan"),
	// 		Product: to.Ptr("test"),
	// 		Publisher: to.Ptr("marketplace-test"),
	// 		Version: to.Ptr("1.0.0"),
	// 	},
	// 	Properties: &armmanagedservices.RegistrationDefinitionProperties{
	// 		Description: to.Ptr("Test"),
	// 		Authorizations: []*armmanagedservices.Authorization{
	// 			{
	// 				PrincipalID: to.Ptr("f98g86a2-4cc4-4e6d-ad47-b3e80a1bcdfc"),
	// 				PrincipalIDDisplayName: to.Ptr("Support User"),
	// 				RoleDefinitionID: to.Ptr("acdd72a7-3385-48ef-bd42-f606fba81ae7"),
	// 			},
	// 			{
	// 				DelegatedRoleDefinitionIDs: []*string{
	// 					to.Ptr("b24988ac-6180-42a0-ab88-20f7382dd24c")},
	// 					PrincipalID: to.Ptr("f98d86a2-4cc4-4e9d-ad47-b3e80a1bcdfc"),
	// 					PrincipalIDDisplayName: to.Ptr("User Access Administrator"),
	// 					RoleDefinitionID: to.Ptr("18d7d88d-d35e-4fb5-a5c3-7773c20a72d9"),
	// 			}},
	// 			EligibleAuthorizations: []*armmanagedservices.EligibleAuthorization{
	// 				{
	// 					JustInTimeAccessPolicy: &armmanagedservices.JustInTimeAccessPolicy{
	// 						ManagedByTenantApprovers: []*armmanagedservices.EligibleApprover{
	// 							{
	// 								PrincipalID: to.Ptr("d9b22cd6-6407-43cc-8c60-07c56df0b51a"),
	// 								PrincipalIDDisplayName: to.Ptr("Approver Group"),
	// 						}},
	// 						MaximumActivationDuration: to.Ptr("PT8H"),
	// 						MultiFactorAuthProvider: to.Ptr(armmanagedservices.MultiFactorAuthProviderAzure),
	// 					},
	// 					PrincipalID: to.Ptr("3e0ed8c6-e902-4fc5-863c-e3ddbb2ae2a2"),
	// 					PrincipalIDDisplayName: to.Ptr("Support User"),
	// 					RoleDefinitionID: to.Ptr("ae349356-3a1b-4a5e-921d-050484c6347e"),
	// 			}},
	// 			ManagedByTenantID: to.Ptr("83ace5cd-bcc3-441a-hd86-e6a75360cecc"),
	// 			ManagedByTenantName: to.Ptr("Test Tenant"),
	// 			ManageeTenantID: to.Ptr("0e06d6a3-55ae-40a3-ac29-350808980808"),
	// 			ManageeTenantName: to.Ptr("Test customer"),
	// 			ProvisioningState: to.Ptr(armmanagedservices.ProvisioningStateSucceeded),
	// 			RegistrationDefinitionName: to.Ptr("DefinitionName"),
	// 		},
	// 	}
}
