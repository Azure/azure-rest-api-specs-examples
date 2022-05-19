Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmanagedservices%2Farmmanagedservices%2Fv0.5.0/sdk/resourcemanager/managedservices/armmanagedservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmanagedservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2022-01-01-preview/examples/PutRegistrationDefinition.json
func ExampleRegistrationDefinitionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmanagedservices.NewRegistrationDefinitionsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"26c128c2-fefa-4340-9bb1-6e081c90ada2",
		"subscription/0afefe50-734e-4610-8a82-a144ahf49dea",
		armmanagedservices.RegistrationDefinition{
			Plan: &armmanagedservices.Plan{
				Name:      to.Ptr("addesai-plan"),
				Product:   to.Ptr("test"),
				Publisher: to.Ptr("marketplace-test"),
				Version:   to.Ptr("1.0.0"),
			},
			Properties: &armmanagedservices.RegistrationDefinitionProperties{
				Description: to.Ptr("Tes1t"),
				Authorizations: []*armmanagedservices.Authorization{
					{
						PrincipalID:            to.Ptr("f98d86a2-4cc4-4e9d-ad47-b3e80a1bcdfc"),
						PrincipalIDDisplayName: to.Ptr("Support User"),
						RoleDefinitionID:       to.Ptr("acdd72a7-3385-48ef-bd42-f606fba81ae7"),
					},
					{
						DelegatedRoleDefinitionIDs: []*string{
							to.Ptr("b24988ac-6180-42a0-ab88-20f7382dd24c")},
						PrincipalID:            to.Ptr("f98d86a2-4cc4-4e9d-ad47-b3e80a1bcdfc"),
						PrincipalIDDisplayName: to.Ptr("User Access Administrator"),
						RoleDefinitionID:       to.Ptr("18d7d88d-d35e-4fb5-a5c3-7773c20a72d9"),
					}},
				EligibleAuthorizations: []*armmanagedservices.EligibleAuthorization{
					{
						JustInTimeAccessPolicy: &armmanagedservices.JustInTimeAccessPolicy{
							ManagedByTenantApprovers: []*armmanagedservices.EligibleApprover{
								{
									PrincipalID:            to.Ptr("d9b22cd6-6407-43cc-8c60-07c56df0b51a"),
									PrincipalIDDisplayName: to.Ptr("Approver Group"),
								}},
							MaximumActivationDuration: to.Ptr("PT8H"),
							MultiFactorAuthProvider:   to.Ptr(armmanagedservices.MultiFactorAuthProviderAzure),
						},
						PrincipalID:            to.Ptr("3e0ed8c6-e902-4fc5-863c-e3ddbb2ae2a2"),
						PrincipalIDDisplayName: to.Ptr("Support User"),
						RoleDefinitionID:       to.Ptr("ae349356-3a1b-4a5e-921d-050484c6347e"),
					}},
				ManagedByTenantID:          to.Ptr("83abe5cd-bcc3-441a-bd86-e6a75360cecc"),
				RegistrationDefinitionName: to.Ptr("DefinitionName"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
```
