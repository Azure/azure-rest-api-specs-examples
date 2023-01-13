package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2021-11-01/examples/async-accounts-create.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewClient("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "contoso", "contososports", armmediaservices.MediaService{
		Location: to.Ptr("South Central US"),
		Tags: map[string]*string{
			"key1": to.Ptr("value1"),
			"key2": to.Ptr("value2"),
		},
		Identity: &armmediaservices.MediaServiceIdentity{
			Type: to.Ptr("UserAssigned"),
			UserAssignedIdentities: map[string]*armmediaservices.UserAssignedManagedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1": {},
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id2": {},
			},
		},
		Properties: &armmediaservices.MediaServiceProperties{
			Encryption: &armmediaservices.AccountEncryption{
				Type: to.Ptr(armmediaservices.AccountEncryptionKeyTypeCustomerKey),
				Identity: &armmediaservices.ResourceIdentity{
					UseSystemAssignedIdentity: to.Ptr(false),
					UserAssignedIdentity:      to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1"),
				},
			},
			KeyDelivery: &armmediaservices.KeyDelivery{
				AccessControl: &armmediaservices.AccessControl{
					DefaultAction: to.Ptr(armmediaservices.DefaultActionAllow),
				},
			},
			PublicNetworkAccess: to.Ptr(armmediaservices.PublicNetworkAccessEnabled),
			StorageAccounts: []*armmediaservices.StorageAccount{
				{
					Type: to.Ptr(armmediaservices.StorageAccountTypePrimary),
					ID:   to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Storage/storageAccounts/contososportsstore"),
					Identity: &armmediaservices.ResourceIdentity{
						UseSystemAssignedIdentity: to.Ptr(false),
						UserAssignedIdentity:      to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1"),
					},
				}},
			StorageAuthentication: to.Ptr(armmediaservices.StorageAuthenticationManagedIdentity),
		},
	}, nil)
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
