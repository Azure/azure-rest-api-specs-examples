package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e7bf3adfa2d5e5cdbb804eec35279501794f461c/specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2021-11-01/examples/media-service-operation-result-by-id.json
func ExampleOperationResultsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmediaservices.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOperationResultsClient().Get(ctx, "westus", "6FBA62C4-99B5-4FF8-9826-FC4744A8864F", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MediaService = armmediaservices.MediaService{
	// 	Name: to.Ptr("contosomovies"),
	// 	Type: to.Ptr("Microsoft.Media/mediaServices"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaServices/contosomovies"),
	// 	Location: to.Ptr("South Central US"),
	// 	Tags: map[string]*string{
	// 		"key1": to.Ptr("value1"),
	// 		"key2": to.Ptr("value2"),
	// 	},
	// 	Identity: &armmediaservices.MediaServiceIdentity{
	// 		Type: to.Ptr("UserAssigned"),
	// 		UserAssignedIdentities: map[string]*armmediaservices.UserAssignedManagedIdentity{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": &armmediaservices.UserAssignedManagedIdentity{
	// 				ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": &armmediaservices.UserAssignedManagedIdentity{
	// 				ClientID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 				PrincipalID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armmediaservices.MediaServiceProperties{
	// 		Encryption: &armmediaservices.AccountEncryption{
	// 			Type: to.Ptr(armmediaservices.AccountEncryptionKeyTypeCustomerKey),
	// 			Identity: &armmediaservices.ResourceIdentity{
	// 				UseSystemAssignedIdentity: to.Ptr(false),
	// 				UserAssignedIdentity: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1"),
	// 			},
	// 			KeyVaultProperties: &armmediaservices.KeyVaultProperties{
	// 				CurrentKeyIdentifier: to.Ptr("https://keyvault.vault.azure.net/keys/key1/ver1"),
	// 				KeyIdentifier: to.Ptr("https://keyvault.vault.azure.net/keys/key1"),
	// 			},
	// 		},
	// 		KeyDelivery: &armmediaservices.KeyDelivery{
	// 			AccessControl: &armmediaservices.AccessControl{
	// 				DefaultAction: to.Ptr(armmediaservices.DefaultActionAllow),
	// 			},
	// 		},
	// 		PrivateEndpointConnections: []*armmediaservices.PrivateEndpointConnection{
	// 			{
	// 				Name: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 				Type: to.Ptr("Microsoft.Media/mediaservices/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosotv/privateEndpointConnections/00000000-0000-0000-0000-000000000001"),
	// 				Properties: &armmediaservices.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armmediaservices.PrivateEndpoint{
	// 						ID: to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/reosuceGroup1/providers/Microsoft.Network/privateEndpoints/pe1"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armmediaservices.PrivateLinkServiceConnectionState{
	// 						Description: to.Ptr("test description"),
	// 						Status: to.Ptr(armmediaservices.PrivateEndpointServiceConnectionStatusApproved),
	// 					},
	// 					ProvisioningState: to.Ptr(armmediaservices.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 				},
	// 			},
	// 			{
	// 				Name: to.Ptr("00000000-0000-0000-0000-000000000002"),
	// 				Type: to.Ptr("Microsoft.Media/mediaservices/privateEndpointConnections"),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Media/mediaservices/contosotv/privateEndpointConnections/00000000-0000-0000-0000-000000000002"),
	// 				Properties: &armmediaservices.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armmediaservices.PrivateEndpoint{
	// 						ID: to.Ptr("/subscriptions/22222222-2222-2222-2222-222222222222/resourceGroups/reosuceGroup2/providers/Microsoft.Network/privateEndpoints/pe2"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armmediaservices.PrivateLinkServiceConnectionState{
	// 						Description: to.Ptr("test description"),
	// 						Status: to.Ptr(armmediaservices.PrivateEndpointServiceConnectionStatusPending),
	// 					},
	// 					ProvisioningState: to.Ptr(armmediaservices.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armmediaservices.ProvisioningStateSucceeded),
	// 		PublicNetworkAccess: to.Ptr(armmediaservices.PublicNetworkAccessEnabled),
	// 		StorageAccounts: []*armmediaservices.StorageAccount{
	// 			{
	// 				Type: to.Ptr(armmediaservices.StorageAccountTypePrimary),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Storage/storageAccounts/contososportsstore"),
	// 				Identity: &armmediaservices.ResourceIdentity{
	// 					UseSystemAssignedIdentity: to.Ptr(false),
	// 					UserAssignedIdentity: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1"),
	// 				},
	// 		}},
	// 		StorageAuthentication: to.Ptr(armmediaservices.StorageAuthenticationManagedIdentity),
	// 	},
	// }
}
