package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/EHNameSpaceCreate.json
func ExampleNamespacesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armeventhub.NewNamespacesClient("SampleSubscription", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"ResurceGroupSample",
		"NamespaceSample",
		armeventhub.EHNamespace{
			Location: to.Ptr("East US"),
			Identity: &armeventhub.Identity{
				Type: to.Ptr(armeventhub.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
				UserAssignedIdentities: map[string]*armeventhub.UserAssignedIdentity{
					"/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud1": {},
					"/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2": {},
				},
			},
			Properties: &armeventhub.EHNamespaceProperties{
				ClusterArmID: to.Ptr("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.EventHub/clusters/enc-test"),
				Encryption: &armeventhub.Encryption{
					KeySource: to.Ptr("Microsoft.KeyVault"),
					KeyVaultProperties: []*armeventhub.KeyVaultProperties{
						{
							Identity: &armeventhub.UserAssignedIdentityProperties{
								UserAssignedIdentity: to.Ptr("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud1"),
							},
							KeyName:     to.Ptr("Samplekey"),
							KeyVaultURI: to.Ptr("https://aprao-keyvault-user.vault-int.azure-int.net/"),
						}},
				},
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
