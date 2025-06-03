package armdellstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dell/armdellstorage"
)

// Generated from example definition: 2025-03-21-preview/FileSystems_Update_MaximumSet_Gen.json
func ExampleFileSystemsClient_Update_fileSystemsUpdateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdellstorage.NewClientFactory("4B6E265D-57CF-4A9D-8B35-3CC68ED9D208", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewFileSystemsClient().Update(ctx, "rgDell", "abcd", armdellstorage.FileSystemResourceUpdate{
		Identity: &armdellstorage.ManagedServiceIdentityUpdate{
			Type: to.Ptr(armdellstorage.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
			UserAssignedIdentities: map[string]*armdellstorage.UserAssignedIdentity{
				"key7645": {},
			},
		},
		Tags: map[string]*string{
			"key6099": to.Ptr("ursbxlphfcguvntuevleacwq"),
		},
		Properties: &armdellstorage.FileSystemResourceUpdateProperties{
			DelegatedSubnetID: to.Ptr("bfpuabdz"),
			Capacity: &armdellstorage.Capacity{
				Current: to.Ptr("5"),
			},
			Encryption: &armdellstorage.EncryptionUpdateProperties{
				EncryptionType: to.Ptr(armdellstorage.ResourceEncryptionTypeCustomerManagedKeysCMK),
				KeyURL:         to.Ptr("https://contoso.com/keyurl/keyVersion"),
				EncryptionIdentityProperties: &armdellstorage.EncryptionIdentityUpdateProperties{
					IdentityType:       to.Ptr(armdellstorage.EncryptionIdentityTypeUserAssigned),
					IdentityResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}"),
				},
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdellstorage.FileSystemsClientUpdateResponse{
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
	// 			},
	// 			ProvisioningState: to.Ptr(armdellstorage.ProvisioningStateAccepted),
	// 			DelegatedSubnetID: to.Ptr("rqkpvczbtqcxiaivtbquixblb"),
	// 			DelegatedSubnetCidr: to.Ptr("10.0.0.1/24"),
	// 			User: &armdellstorage.UserDetails{
	// 			},
	// 			FileSystemID: to.Ptr("filesystemId"),
	// 			SmartConnectFqdn: to.Ptr("fqdn"),
	// 			DellReferenceNumber: to.Ptr("fhewkj"),
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
	// 			PrincipalID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 			TenantID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 			Type: to.Ptr(armdellstorage.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
	// 			UserAssignedIdentities: map[string]*armdellstorage.UserAssignedIdentity{
	// 				"key7645": &armdellstorage.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("22222222-2222-2222-2222-222222222222"),
	// 					ClientID: to.Ptr("22222222-2222-2222-2222-222222222222"),
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
