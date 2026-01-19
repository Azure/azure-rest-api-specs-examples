package armdellstorage_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dell/armdellstorage"
)

// Generated from example definition: 2025-03-21/FileSystems_ListByResourceGroup_MaximumSet_Gen.json
func ExampleFileSystemsClient_NewListByResourceGroupPager_fileSystemsListByResourceGroupMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdellstorage.NewClientFactory("4B6E265D-57CF-4A9D-8B35-3CC68ED9D208", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewFileSystemsClient().NewListByResourceGroupPager("rgDell", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page = armdellstorage.FileSystemsClientListByResourceGroupResponse{
		// 	FileSystemResourceListResult: armdellstorage.FileSystemResourceListResult{
		// 		Value: []*armdellstorage.FileSystemResource{
		// 			{
		// 				Properties: &armdellstorage.FileSystemResourceProperties{
		// 					Capacity: &armdellstorage.Capacity{
		// 						Min: to.Ptr("1"),
		// 						Max: to.Ptr("10"),
		// 						Incremental: to.Ptr("1"),
		// 						Current: to.Ptr("5"),
		// 					},
		// 					Marketplace: &armdellstorage.MarketplaceDetails{
		// 						MarketplaceSubscriptionID: to.Ptr("mvjcxwndudbylynme"),
		// 						PlanID: to.Ptr("eekvwfndjoxijeasksnt"),
		// 						OfferID: to.Ptr("bcganbkmvznyqfnvhjuag"),
		// 						PublisherID: to.Ptr("trdzykoeskmcwpo"),
		// 						MarketplaceSubscriptionStatus: to.Ptr(armdellstorage.MarketplaceSubscriptionStatusPendingFulfillmentStart),
		// 						PrivateOfferID: to.Ptr("privateOfferId"),
		// 						PlanName: to.Ptr("planeName"),
		// 						EndDate: to.Ptr("2023-05-27T17:00:00-07:00"),
		// 					},
		// 					ProvisioningState: to.Ptr(armdellstorage.ProvisioningStateAccepted),
		// 					DelegatedSubnetID: to.Ptr("rqkpvczbtqcxiaivtbquixblb"),
		// 					DelegatedSubnetCidr: to.Ptr("10.0.0.1/24"),
		// 					User: &armdellstorage.UserDetails{
		// 					},
		// 					FileSystemID: to.Ptr("filesystemId"),
		// 					SmartConnectFqdn: to.Ptr("fqdn"),
		// 					DellReferenceNumber: to.Ptr("fhewkj"),
		// 					Encryption: &armdellstorage.EncryptionProperties{
		// 						EncryptionType: to.Ptr(armdellstorage.ResourceEncryptionTypeCustomerManagedKeysCMK),
		// 						KeyURL: to.Ptr("https://contoso.com/keyurl/keyVersion"),
		// 						EncryptionIdentityProperties: &armdellstorage.EncryptionIdentityProperties{
		// 							IdentityType: to.Ptr(armdellstorage.EncryptionIdentityTypeUserAssigned),
		// 							IdentityResourceID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}"),
		// 						},
		// 					},
		// 				},
		// 				Identity: &armdellstorage.ManagedServiceIdentity{
		// 					PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 					TenantID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 					Type: to.Ptr(armdellstorage.ManagedServiceIdentityTypeSystemAssignedUserAssigned),
		// 					UserAssignedIdentities: map[string]*armdellstorage.UserAssignedIdentity{
		// 						"key7644": &armdellstorage.UserAssignedIdentity{
		// 							PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 							ClientID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 						},
		// 					},
		// 				},
		// 				Tags: map[string]*string{
		// 					"key7594": to.Ptr("sfkwapubiurgedzveido"),
		// 				},
		// 				Location: to.Ptr("cvbmsqftppe"),
		// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
		// 				Name: to.Ptr("hyjt"),
		// 				Type: to.Ptr("mzfepehqauauepqojtnobuyiegj"),
		// 				SystemData: &armdellstorage.SystemData{
		// 					CreatedBy: to.Ptr("xfvccbyptfzz"),
		// 					CreatedByType: to.Ptr(armdellstorage.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-28T09:20:01.255Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("uojozqmijcudlqxmvwckofopoori"),
		// 					LastModifiedByType: to.Ptr(armdellstorage.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-08-28T09:20:01.255Z"); return t}()),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
