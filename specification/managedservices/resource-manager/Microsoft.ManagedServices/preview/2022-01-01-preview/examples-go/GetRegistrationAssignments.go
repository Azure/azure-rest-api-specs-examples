package armmanagedservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2022-01-01-preview/examples/GetRegistrationAssignments.json
func ExampleRegistrationAssignmentsClient_NewListPager_getRegistrationAssignments() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmanagedservices.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRegistrationAssignmentsClient().NewListPager("subscription/0afefe50-734e-4610-8a82-a144ahf49dea", &armmanagedservices.RegistrationAssignmentsClientListOptions{ExpandRegistrationDefinition: nil,
		Filter: nil,
	})
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
		// page.RegistrationAssignmentList = armmanagedservices.RegistrationAssignmentList{
		// 	Value: []*armmanagedservices.RegistrationAssignment{
		// 		{
		// 			Name: to.Ptr("484a7d5f-9729-4b87-bc9b-26610985a013"),
		// 			Type: to.Ptr("Microsoft.ManagedServices/registrationAssignments"),
		// 			ID: to.Ptr("/subscriptions/0afefe50-734e-4610-8c82-a144aff49dea/providers/Microsoft.ManagedServices/registrationAssignments/484a7d5f-9729-4b87-bc9b-26610985a013"),
		// 			Properties: &armmanagedservices.RegistrationAssignmentProperties{
		// 				ProvisioningState: to.Ptr(armmanagedservices.ProvisioningStateSucceeded),
		// 				RegistrationDefinition: &armmanagedservices.RegistrationAssignmentPropertiesRegistrationDefinition{
		// 					Name: to.Ptr("26c128c2-fefa-4340-9bb1-8e081c90ada2"),
		// 					Type: to.Ptr("Microsoft.ManagedServices/registrationDefinitions"),
		// 					ID: to.Ptr("/subscriptions/0afefe50-734e-4610-8c82-a144aff49dea/providers/Microsoft.ManagedServices/registrationDefinitions/26c128c2-fefa-4340-9bb1-8e081c90ada2"),
		// 					Plan: &armmanagedservices.Plan{
		// 						Name: to.Ptr("addesai-plan"),
		// 						Product: to.Ptr("test"),
		// 						Publisher: to.Ptr("marketplace-test"),
		// 						Version: to.Ptr("1.0.0"),
		// 					},
		// 					Properties: &armmanagedservices.RegistrationAssignmentPropertiesRegistrationDefinitionProperties{
		// 						Description: to.Ptr("Test"),
		// 						Authorizations: []*armmanagedservices.Authorization{
		// 							{
		// 								PrincipalID: to.Ptr("f98d86a2-4cc4-4e9d-ad47-b3e80a1bcdfc"),
		// 								PrincipalIDDisplayName: to.Ptr("Support User"),
		// 								RoleDefinitionID: to.Ptr("acdd72a7-3385-48ef-bd42-f606fba81ae7"),
		// 							},
		// 							{
		// 								DelegatedRoleDefinitionIDs: []*string{
		// 									to.Ptr("b24988ac-6180-42a0-ab88-20f7382dd24c")},
		// 									PrincipalID: to.Ptr("f98d86a2-4cc4-4e9d-ad47-b3e80a1bcdfc"),
		// 									PrincipalIDDisplayName: to.Ptr("User Access Administrator"),
		// 									RoleDefinitionID: to.Ptr("18d7d88d-d35e-4fb5-a5c3-7773c20a72d9"),
		// 							}},
		// 							EligibleAuthorizations: []*armmanagedservices.EligibleAuthorization{
		// 								{
		// 									JustInTimeAccessPolicy: &armmanagedservices.JustInTimeAccessPolicy{
		// 										MaximumActivationDuration: to.Ptr("PT8H"),
		// 										MultiFactorAuthProvider: to.Ptr(armmanagedservices.MultiFactorAuthProviderAzure),
		// 									},
		// 									PrincipalID: to.Ptr("700bddf4-2c3b-4cd1-bb02-6a2c622524f4"),
		// 									RoleDefinitionID: to.Ptr("8e3af657-a8ff-443c-a75c-2fe8c4bcb635"),
		// 							}},
		// 							ManagedByTenantID: to.Ptr("83abe5cd-bcc3-441a-bd86-e6a75360cecc"),
		// 							ManagedByTenantName: to.Ptr("Contoso Corp."),
		// 							ManageeTenantID: to.Ptr("01c0bcd5-4f47-4e4b-b492-418b7e2a8854"),
		// 							ManageeTenantName: to.Ptr("test_test_aad_SbtFhyGiLHPFm"),
		// 							ProvisioningState: to.Ptr(armmanagedservices.ProvisioningStateSucceeded),
		// 							RegistrationDefinitionName: to.Ptr("DefinitionName"),
		// 						},
		// 					},
		// 					RegistrationDefinitionID: to.Ptr("/subscriptions/0afefe50-734e-4610-8c82-a144aff49dea/providers/Microsoft.ManagedServices/registrationDefinitions/26c128c2-fefa-4340-9bb1-8e081c90ada2"),
		// 				},
		// 		}},
		// 	}
	}
}
