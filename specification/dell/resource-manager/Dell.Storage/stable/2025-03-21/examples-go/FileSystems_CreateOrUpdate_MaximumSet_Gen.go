package armdellstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dell/armdellstorage"
)

// Generated from example definition: 2025-03-21/FileSystems_CreateOrUpdate_MaximumSet_Gen.json
func ExampleFileSystemsClient_BeginCreateOrUpdate_fileSystemsCreateOrUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdellstorage.NewClientFactory("4B6E265D-57CF-4A9D-8B35-3CC68ED9D208", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewFileSystemsClient().BeginCreateOrUpdate(ctx, "rgDell", "abcd", armdellstorage.FileSystemResource{
		Properties: &armdellstorage.FileSystemResourceProperties{
			Marketplace: &armdellstorage.MarketplaceDetails{
				MarketplaceSubscriptionID:     to.Ptr("mvjcxwndudbylynme"),
				PlanID:                        to.Ptr("eekvwfndjoxijeasksnt"),
				OfferID:                       to.Ptr("bcganbkmvznyqfnvhjuag"),
				PublisherID:                   to.Ptr("trdzykoeskmcwpo"),
				MarketplaceSubscriptionStatus: to.Ptr(armdellstorage.MarketplaceSubscriptionStatusPendingFulfillmentStart),
				PrivateOfferID:                to.Ptr("privateOfferId"),
				PlanName:                      to.Ptr("planeName"),
			},
			DelegatedSubnetID:   to.Ptr("rqkpvczbtqcxiaivtbuixblb"),
			DelegatedSubnetCidr: to.Ptr("10.0.0.1/24"),
			User: &armdellstorage.UserDetails{
				Email: to.Ptr("jwogfgznmjabdbcjcljjlkxdpc"),
			},
			SmartConnectFqdn:    to.Ptr("fqdn"),
			OneFsURL:            to.Ptr("oneFsUrl"),
			DellReferenceNumber: to.Ptr("fhewkj"),
			Encryption: &armdellstorage.EncryptionProperties{
				EncryptionType: to.Ptr(armdellstorage.ResourceEncryptionTypeCustomerManagedKeysCMK),
				KeyURL:         to.Ptr("https://contoso.com/keyurl/keyVersion"),
				EncryptionIdentityProperties: &armdellstorage.EncryptionIdentityProperties{
					IdentityType:       to.Ptr(armdellstorage.EncryptionIdentityTypeUserAssigned),
					IdentityResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}"),
				},
			},
		},
		Identity: &armdellstorage.ManagedServiceIdentity{
			Type: to.Ptr(armdellstorage.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armdellstorage.UserAssignedIdentity{
				"key7644": {},
			},
		},
		Tags: map[string]*string{
			"key7594": to.Ptr("sfkwapubiurgedzveido"),
		},
		Location: to.Ptr("cvbmsqftppe"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdellstorage.FileSystemsClientCreateOrUpdateResponse{
	// 	FileSystemResource: &armdellstorage.FileSystemResource{
	// 		Properties: &armdellstorage.FileSystemResourceProperties{
	// 			Capacity: &armdellstorage.Capacity{
	// 				Min: to.Ptr("1"),
	// 				Max: to.Ptr("10"),
	// 				Incremental: to.Ptr("1"),
	// 				Current: to.Ptr("5"),
	// 			},
	// 			Marketplace: &armdellstorage.MarketplaceDetails{
	// 				MarketplaceSubscriptionID: to.Ptr("mvjcxwndudbylynme"),
	// 				PlanID: to.Ptr("eekvwfndjoxijeasksnt"),
	// 				OfferID: to.Ptr("bcganbkmvznyqfnvhjuag"),
	// 				PublisherID: to.Ptr("trdzykoeskmcwpo"),
	// 				MarketplaceSubscriptionStatus: to.Ptr(armdellstorage.MarketplaceSubscriptionStatusPendingFulfillmentStart),
	// 				PrivateOfferID: to.Ptr("privateOfferId"),
	// 				PlanName: to.Ptr("planeName"),
	// 				EndDate: to.Ptr("2023-05-27T17:00:00-07:00"),
	// 				TermUnit: to.Ptr("P1Y"),
	// 			},
	// 			ProvisioningState: to.Ptr(armdellstorage.ProvisioningStateAccepted),
	// 			DelegatedSubnetID: to.Ptr("rqkpvczbtqcxiaivtbquixblb"),
	// 			DelegatedSubnetCidr: to.Ptr("10.0.0.1/24"),
	// 			User: &armdellstorage.UserDetails{
	// 			},
	// 			SmartConnectFqdn: to.Ptr("fqdn"),
	// 			OneFsURL: to.Ptr("oneFsUrl"),
	// 			DellReferenceNumber: to.Ptr("fhewkj"),
	// 			FileSystemID: to.Ptr("filesystemId"),
	// 			Encryption: &armdellstorage.EncryptionProperties{
	// 				EncryptionType: to.Ptr(armdellstorage.ResourceEncryptionTypeCustomerManagedKeysCMK),
	// 				KeyURL: to.Ptr("https://contoso.com/keyurl/keyVersion"),
	// 				EncryptionIdentityProperties: &armdellstorage.EncryptionIdentityProperties{
	// 					IdentityType: to.Ptr(armdellstorage.EncryptionIdentityTypeUserAssigned),
	// 					IdentityResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}"),
	// 				},
	// 			},
	// 		},
	// 		Identity: &armdellstorage.ManagedServiceIdentity{
	// 			PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 			TenantID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 			Type: to.Ptr(armdellstorage.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
	// 			UserAssignedIdentities: map[string]*armdellstorage.UserAssignedIdentity{
	// 				"key7644": &armdellstorage.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 					ClientID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 				},
	// 			},
	// 		},
	// 		Tags: map[string]*string{
	// 			"key7594": to.Ptr("sfkwapubiurgedzveido"),
	// 		},
	// 		Location: to.Ptr("cvbmsqftppe"),
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
	// 		Name: to.Ptr("hyjt"),
	// 		Type: to.Ptr("mzfepehqauauepqojtnobuyiegj"),
	// 		SystemData: &armdellstorage.SystemData{
	// 			CreatedBy: to.Ptr("xfvccbyptfzz"),
	// 			CreatedByType: to.Ptr(armdellstorage.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-28T09:20:01.255Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("uojozqmijcudlqxmvwckofopoori"),
	// 			LastModifiedByType: to.Ptr(armdellstorage.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-28T09:20:01.255Z"); return t}()),
	// 		},
	// 	},
	// }
}
