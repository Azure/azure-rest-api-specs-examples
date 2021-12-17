Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmachinelearningservices%2Farmmachinelearningservices%2Fv0.1.0/sdk/resourcemanager/machinelearningservices/armmachinelearningservices/README.md) on how to add the SDK to your project and authenticate.

```go
package armmachinelearningservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearningservices/armmachinelearningservices"
)

// x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/create.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armmachinelearningservices.Workspace{
			Identity: &armmachinelearningservices.Identity{
				Type: armmachinelearningservices.ResourceIdentityTypeSystemAssignedUserAssigned.ToPtr(),
				UserAssignedIdentities: map[string]*armmachinelearningservices.UserAssignedIdentity{
					"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai": {},
				},
			},
			Location: to.StringPtr("<location>"),
			Properties: &armmachinelearningservices.WorkspaceProperties{
				Description:         to.StringPtr("<description>"),
				ApplicationInsights: to.StringPtr("<application-insights>"),
				ContainerRegistry:   to.StringPtr("<container-registry>"),
				Encryption: &armmachinelearningservices.EncryptionProperty{
					Identity: &armmachinelearningservices.IdentityForCmk{
						UserAssignedIdentity: to.StringPtr("<user-assigned-identity>"),
					},
					KeyVaultProperties: &armmachinelearningservices.KeyVaultProperties{
						IdentityClientID: to.StringPtr("<identity-client-id>"),
						KeyIdentifier:    to.StringPtr("<key-identifier>"),
						KeyVaultArmID:    to.StringPtr("<key-vault-arm-id>"),
					},
					Status: armmachinelearningservices.EncryptionStatusEnabled.ToPtr(),
				},
				FriendlyName: to.StringPtr("<friendly-name>"),
				HbiWorkspace: to.BoolPtr(false),
				KeyVault:     to.StringPtr("<key-vault>"),
				SharedPrivateLinkResources: []*armmachinelearningservices.SharedPrivateLinkResource{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armmachinelearningservices.SharedPrivateLinkResourceProperty{
							GroupID:               to.StringPtr("<group-id>"),
							PrivateLinkResourceID: to.StringPtr("<private-link-resource-id>"),
							RequestMessage:        to.StringPtr("<request-message>"),
							Status:                armmachinelearningservices.PrivateEndpointServiceConnectionStatusApproved.ToPtr(),
						},
					}},
				StorageAccount: to.StringPtr("<storage-account>"),
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
	log.Printf("Workspace.ID: %s\n", *res.ID)
}
```
