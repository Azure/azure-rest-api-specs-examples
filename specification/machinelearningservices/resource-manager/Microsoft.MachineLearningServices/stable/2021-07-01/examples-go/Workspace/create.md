Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fmachinelearningservices%2Farmmachinelearningservices%2Fv0.4.0/sdk/resourcemanager/machinelearningservices/armmachinelearningservices/README.md) on how to add the SDK to your project and authenticate.

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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2021-07-01/examples/Workspace/create.json
func ExampleWorkspacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmachinelearningservices.NewWorkspacesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		armmachinelearningservices.Workspace{
			Identity: &armmachinelearningservices.Identity{
				Type: to.Ptr(armmachinelearningservices.ResourceIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armmachinelearningservices.UserAssignedIdentity{
					"/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/workspace-1234/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testuai": {},
				},
			},
			Location: to.Ptr("<location>"),
			Properties: &armmachinelearningservices.WorkspaceProperties{
				Description:         to.Ptr("<description>"),
				ApplicationInsights: to.Ptr("<application-insights>"),
				ContainerRegistry:   to.Ptr("<container-registry>"),
				Encryption: &armmachinelearningservices.EncryptionProperty{
					Identity: &armmachinelearningservices.IdentityForCmk{
						UserAssignedIdentity: to.Ptr("<user-assigned-identity>"),
					},
					KeyVaultProperties: &armmachinelearningservices.KeyVaultProperties{
						IdentityClientID: to.Ptr("<identity-client-id>"),
						KeyIdentifier:    to.Ptr("<key-identifier>"),
						KeyVaultArmID:    to.Ptr("<key-vault-arm-id>"),
					},
					Status: to.Ptr(armmachinelearningservices.EncryptionStatusEnabled),
				},
				FriendlyName: to.Ptr("<friendly-name>"),
				HbiWorkspace: to.Ptr(false),
				KeyVault:     to.Ptr("<key-vault>"),
				SharedPrivateLinkResources: []*armmachinelearningservices.SharedPrivateLinkResource{
					{
						Name: to.Ptr("<name>"),
						Properties: &armmachinelearningservices.SharedPrivateLinkResourceProperty{
							GroupID:               to.Ptr("<group-id>"),
							PrivateLinkResourceID: to.Ptr("<private-link-resource-id>"),
							RequestMessage:        to.Ptr("<request-message>"),
							Status:                to.Ptr(armmachinelearningservices.PrivateEndpointServiceConnectionStatusApproved),
						},
					}},
				StorageAccount: to.Ptr("<storage-account>"),
			},
		},
		&armmachinelearningservices.WorkspacesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
