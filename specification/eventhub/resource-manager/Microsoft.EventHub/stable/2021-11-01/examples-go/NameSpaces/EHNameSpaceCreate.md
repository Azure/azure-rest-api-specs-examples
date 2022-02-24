Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Feventhub%2Farmeventhub%2Fv0.3.1/sdk/resourcemanager/eventhub/armeventhub/README.md) on how to add the SDK to your project and authenticate.

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

// x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/EHNameSpaceCreate.json
func ExampleNamespacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armeventhub.NewNamespacesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<namespace-name>",
		armeventhub.EHNamespace{
			Location: to.StringPtr("<location>"),
			Identity: &armeventhub.Identity{
				Type: armeventhub.ManagedServiceIdentityTypeSystemAssignedUserAssigned.ToPtr(),
				UserAssignedIdentities: map[string]*armeventhub.UserAssignedIdentity{
					"/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud1": {},
					"/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2": {},
				},
			},
			Properties: &armeventhub.EHNamespaceProperties{
				ClusterArmID: to.StringPtr("<cluster-arm-id>"),
				Encryption: &armeventhub.Encryption{
					KeySource: to.StringPtr("<key-source>"),
					KeyVaultProperties: []*armeventhub.KeyVaultProperties{
						{
							Identity: &armeventhub.UserAssignedIdentityProperties{
								UserAssignedIdentity: to.StringPtr("<user-assigned-identity>"),
							},
							KeyName:     to.StringPtr("<key-name>"),
							KeyVaultURI: to.StringPtr("<key-vault-uri>"),
						}},
				},
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
	log.Printf("Response result: %#v\n", res.NamespacesClientCreateOrUpdateResult)
}
```
