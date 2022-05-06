Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Feventhub%2Farmeventhub%2Fv0.5.0/sdk/resourcemanager/eventhub/armeventhub/README.md) on how to add the SDK to your project and authenticate.

```go
package armeventhub_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/EHNameSpaceCreate.json
func ExampleNamespacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armeventhub.NewNamespacesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<namespace-name>",
		armeventhub.EHNamespace{
			Location: to.Ptr("<location>"),
			Identity: &armeventhub.Identity{
				Type: to.Ptr(armeventhub.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armeventhub.UserAssignedIdentity{
					"/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud1": {},
					"/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2": {},
				},
			},
			Properties: &armeventhub.EHNamespaceProperties{
				ClusterArmID: to.Ptr("<cluster-arm-id>"),
				Encryption: &armeventhub.Encryption{
					KeySource: to.Ptr("<key-source>"),
					KeyVaultProperties: []*armeventhub.KeyVaultProperties{
						{
							Identity: &armeventhub.UserAssignedIdentityProperties{
								UserAssignedIdentity: to.Ptr("<user-assigned-identity>"),
							},
							KeyName:     to.Ptr("<key-name>"),
							KeyVaultURI: to.Ptr("<key-vault-uri>"),
						}},
				},
			},
		},
		&armeventhub.NamespacesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
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
