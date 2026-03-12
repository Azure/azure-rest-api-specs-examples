package armeventgrid_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid/v2"
)

// Generated from example definition: 2025-07-15-preview/SystemTopics_Get.json
func ExampleSystemTopicsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventgrid.NewClientFactory("8f6b6269-84f2-4d09-9e31-1127efcd1e40", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSystemTopicsClient().Get(ctx, "examplerg", "exampleSystemTopic2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armeventgrid.SystemTopicsClientGetResponse{
	// 	SystemTopic: &armeventgrid.SystemTopic{
	// 		Name: to.Ptr("exampleSystemTopic2"),
	// 		Type: to.Ptr("Microsoft.EventGrid/systemTopics"),
	// 		ID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/systemTopics/exampleSystemTopic2"),
	// 		Identity: &armeventgrid.IdentityInfo{
	// 			Type: to.Ptr(armeventgrid.IdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armeventgrid.UserIdentityProperties{
	// 				"/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/azureeventgridrunnerrgcentraluseuap/providers/Microsoft.ManagedIdentity/userAssignedIdentities/user-assigned-id": &armeventgrid.UserIdentityProperties{
	// 					ClientID: to.Ptr("11af827c-d951-4766-9c41-15565fa13157"),
	// 					PrincipalID: to.Ptr("a5435982-c82d-453b-b244-3115b7c8789c"),
	// 				},
	// 			},
	// 		},
	// 		Location: to.Ptr("centraluseuap"),
	// 		Properties: &armeventgrid.SystemTopicProperties{
	// 			Encryption: &armeventgrid.KeyEncryption{
	// 				CustomerManagedKeyEncryption: []*armeventgrid.CustomerManagedKeyEncryption{
	// 					{
	// 						KeyEncryptionKeyIdentity: &armeventgrid.KeyEncryptionKeyIdentity{
	// 							Type: to.Ptr(armeventgrid.KeyEncryptionIdentityTypeUserAssigned),
	// 							UserAssignedIdentityResourceID: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/azureeventgridrunnerrgcentraluseuap/providers/Microsoft.ManagedIdentity/userAssignedIdentities/user-assigned-id"),
	// 						},
	// 						KeyEncryptionKeyStatus: to.Ptr(armeventgrid.KeyEncryptionKeyStatusActive),
	// 						KeyEncryptionKeyStatusFriendlyDescription: to.Ptr("Customer Managed Key (CMK) Encryption is 'Active' and running normally."),
	// 						KeyEncryptionKeyURL: to.Ptr("https://ege2ekeyvault.vault.azure.net/keys/ValidKey1"),
	// 					},
	// 				},
	// 			},
	// 			MetricResourceID: to.Ptr("183c0fb1-17ff-47b6-ac77-5a47420ab01e"),
	// 			PlatformCapabilities: &armeventgrid.PlatformCapabilities{
	// 				ConfidentialCompute: &armeventgrid.ConfidentialCompute{
	// 					Mode: to.Ptr(armeventgrid.ConfidentialComputeModeEnabled),
	// 				},
	// 			},
	// 			ProvisioningState: to.Ptr(armeventgrid.ResourceProvisioningStateSucceeded),
	// 			Source: to.Ptr("/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/azureeventgridrunnerrgcentraluseuap/providers/microsoft.storage/storageaccounts/pubstgrunnerb71cd29e"),
	// 			TopicType: to.Ptr("microsoft.storage.storageaccounts"),
	// 		},
	// 	},
	// }
}
