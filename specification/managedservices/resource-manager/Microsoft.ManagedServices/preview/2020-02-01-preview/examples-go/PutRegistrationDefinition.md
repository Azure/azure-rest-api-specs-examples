Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmanagedservices%2Farmmanagedservices%2Fv0.2.1/sdk/resourcemanager/managedservices/armmanagedservices/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/managedservices/resource-manager/Microsoft.ManagedServices/preview/2020-02-01-preview/examples/PutRegistrationDefinition.json
func ExampleRegistrationDefinitionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmanagedservices.NewRegistrationDefinitionsClient(cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<registration-definition-id>",
		"<scope>",
		armmanagedservices.RegistrationDefinition{
			Plan: &armmanagedservices.Plan{
				Name:      to.StringPtr("<name>"),
				Product:   to.StringPtr("<product>"),
				Publisher: to.StringPtr("<publisher>"),
				Version:   to.StringPtr("<version>"),
			},
			Properties: &armmanagedservices.RegistrationDefinitionProperties{
				Description: to.StringPtr("<description>"),
				Authorizations: []*armmanagedservices.Authorization{
					{
						PrincipalID:            to.StringPtr("<principal-id>"),
						PrincipalIDDisplayName: to.StringPtr("<principal-iddisplay-name>"),
						RoleDefinitionID:       to.StringPtr("<role-definition-id>"),
					},
					{
						DelegatedRoleDefinitionIDs: []*string{
							to.StringPtr("b24988ac-6180-42a0-ab88-20f7382dd24c")},
						PrincipalID:            to.StringPtr("<principal-id>"),
						PrincipalIDDisplayName: to.StringPtr("<principal-iddisplay-name>"),
						RoleDefinitionID:       to.StringPtr("<role-definition-id>"),
					}},
				EligibleAuthorizations: []*armmanagedservices.EligibleAuthorization{
					{
						JustInTimeAccessPolicy: &armmanagedservices.JustInTimeAccessPolicy{
							ManagedByTenantApprovers: []*armmanagedservices.EligibleApprover{
								{
									PrincipalID:            to.StringPtr("<principal-id>"),
									PrincipalIDDisplayName: to.StringPtr("<principal-iddisplay-name>"),
								}},
							MaximumActivationDuration: to.StringPtr("<maximum-activation-duration>"),
							MultiFactorAuthProvider:   armmanagedservices.MultiFactorAuthProvider("Azure").ToPtr(),
						},
						PrincipalID:            to.StringPtr("<principal-id>"),
						PrincipalIDDisplayName: to.StringPtr("<principal-iddisplay-name>"),
						RoleDefinitionID:       to.StringPtr("<role-definition-id>"),
					}},
				ManagedByTenantID:          to.StringPtr("<managed-by-tenant-id>"),
				RegistrationDefinitionName: to.StringPtr("<registration-definition-name>"),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.RegistrationDefinitionsClientCreateOrUpdateResult)
}
```
