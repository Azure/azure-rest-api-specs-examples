package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: 2025-07-15-preview/Topics_CreateOrUpdate.json
func ExampleTopicsClient_BeginCreateOrUpdate_topicsCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("8f6b6269-84f2-4d09-9e31-1127efcd1e40", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewTopicsClient().BeginCreateOrUpdate(ctx, "examplerg", "exampletopic1", armeventgrid.Topic{
		Identity: &armeventgrid.IdentityInfo{
			Type: to.Ptr(armeventgrid.IdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armeventgrid.UserIdentityProperties{
				"/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/azureeventgridrunnerrgcentraluseuap/providers/Microsoft.ManagedIdentity/userAssignedIdentities/user-assigned-id": {},
			},
		},
		Location: to.Ptr("westus2"),
		Properties: &armeventgrid.TopicProperties{
			Encryption: &armeventgrid.KeyEncryption{
				CustomerManagedKeyEncryption: []*armeventgrid.CustomerManagedKeyEncryption{
					{
						KeyEncryptionKeyIdentity: &armeventgrid.KeyEncryptionKeyIdentity{
							Type:                           to.Ptr(armeventgrid.KeyEncryptionIdentityTypeUserAssigned),
							UserAssignedIdentityResourceID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/azureeventgridrunnerrgcentraluseuap/providers/Microsoft.ManagedIdentity/userAssignedIdentities/user-assigned-id"),
						},
						KeyEncryptionKeyURL: to.Ptr("https://ege2ekeyvault.vault.azure.net/keys/ValidKey1"),
					},
				},
			},
			PlatformCapabilities: &armeventgrid.PlatformCapabilities{
				ConfidentialCompute: &armeventgrid.ConfidentialCompute{
					Mode: to.Ptr(armeventgrid.ConfidentialComputeModeEnabled),
				},
			},
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
