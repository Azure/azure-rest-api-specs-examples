Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmanagedservices%2Farmmanagedservices%2Fv0.4.0/sdk/resourcemanager/managedservices/armmanagedservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmanagedservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2022-01-01-preview/examples/PutRegistrationDefinition.json
func ExampleRegistrationDefinitionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmanagedservices.NewRegistrationDefinitionsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<registration-definition-id>",
		"<scope>",
		armmanagedservices.RegistrationDefinition{
			Plan: &armmanagedservices.Plan{
				Name:      to.Ptr("<name>"),
				Product:   to.Ptr("<product>"),
				Publisher: to.Ptr("<publisher>"),
				Version:   to.Ptr("<version>"),
			},
			Properties: &armmanagedservices.RegistrationDefinitionProperties{
				Description: to.Ptr("<description>"),
				Authorizations: []*armmanagedservices.Authorization{
					{
						PrincipalID:            to.Ptr("<principal-id>"),
						PrincipalIDDisplayName: to.Ptr("<principal-iddisplay-name>"),
						RoleDefinitionID:       to.Ptr("<role-definition-id>"),
					},
					{
						DelegatedRoleDefinitionIDs: []*string{
							to.Ptr("b24988ac-6180-42a0-ab88-20f7382dd24c")},
						PrincipalID:            to.Ptr("<principal-id>"),
						PrincipalIDDisplayName: to.Ptr("<principal-iddisplay-name>"),
						RoleDefinitionID:       to.Ptr("<role-definition-id>"),
					}},
				EligibleAuthorizations: []*armmanagedservices.EligibleAuthorization{
					{
						JustInTimeAccessPolicy: &armmanagedservices.JustInTimeAccessPolicy{
							ManagedByTenantApprovers: []*armmanagedservices.EligibleApprover{
								{
									PrincipalID:            to.Ptr("<principal-id>"),
									PrincipalIDDisplayName: to.Ptr("<principal-iddisplay-name>"),
								}},
							MaximumActivationDuration: to.Ptr("<maximum-activation-duration>"),
							MultiFactorAuthProvider:   to.Ptr(armmanagedservices.MultiFactorAuthProviderAzure),
						},
						PrincipalID:            to.Ptr("<principal-id>"),
						PrincipalIDDisplayName: to.Ptr("<principal-iddisplay-name>"),
						RoleDefinitionID:       to.Ptr("<role-definition-id>"),
					}},
				ManagedByTenantID:          to.Ptr("<managed-by-tenant-id>"),
				RegistrationDefinitionName: to.Ptr("<registration-definition-name>"),
			},
		},
		&armmanagedservices.RegistrationDefinitionsClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
```
