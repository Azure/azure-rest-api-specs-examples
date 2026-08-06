package armdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/discovery/armdiscovery"
)

// Generated from example definition: 2026-06-01/Bookshelves_ListBySubscription_MaximumSet_Gen.json
func ExampleBookshelvesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdiscovery.NewClientFactory("A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBookshelvesClient().NewListBySubscriptionPager(nil)
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
		// page = armdiscovery.BookshelvesClientListBySubscriptionResponse{
		// 	BookshelfListResult: armdiscovery.BookshelfListResult{
		// 		Value: []*armdiscovery.Bookshelf{
		// 			{
		// 				ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/Bookshelves/mulgmkcahzfyhpehrptfkswy"),
		// 				Name: to.Ptr("mulgmkcahzfyhpehrptfkswy"),
		// 				Tags: map[string]*string{
		// 					"key3815": to.Ptr("lrfxoxtbcgjpokrmtlzvknkcu"),
		// 				},
		// 				Location: to.Ptr("uksouth"),
		// 				Type: to.Ptr("Microsoft.Discovery/Bookshelves"),
		// 				SystemData: &armdiscovery.SystemData{
		// 					CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
		// 					CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
		// 					LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
		// 				},
		// 				Properties: &armdiscovery.BookshelfProperties{
		// 					ProvisioningState: to.Ptr(armdiscovery.ProvisioningStateSucceeded),
		// 					WorkloadIdentities: map[string]*armdiscovery.UserAssignedIdentity{
		// 						"key3650": &armdiscovery.UserAssignedIdentity{
		// 							PrincipalID: to.Ptr("00000011-1111-2222-2222-123456789111"),
		// 							ClientID: to.Ptr("00000011-1111-2222-2222-123456789111"),
		// 						},
		// 					},
		// 					CustomerManagedKeys: to.Ptr(armdiscovery.CustomerManagedKeysEnabled),
		// 					KeyVaultProperties: &armdiscovery.BookshelfKeyVaultProperties{
		// 						KeyVaultURI: to.Ptr("https://microsoft.com/a"),
		// 						KeyName: to.Ptr("picc"),
		// 						KeyVersion: to.Ptr("bnzaxtmzrsjovfifuizqsfsphspdyc"),
		// 						IdentityClientID: to.Ptr("00000011-1111-2222-2222-123456789111"),
		// 					},
		// 					LogAnalyticsClusterID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.OperationalInsights/clusters/cluster1"),
		// 					PrivateEndpointConnections: []*armdiscovery.PrivateEndpointConnection{
		// 						{
		// 							Properties: &armdiscovery.PrivateEndpointConnectionProperties{
		// 								GroupIDs: []*string{
		// 									to.Ptr("bpyliugtuio"),
		// 								},
		// 								PrivateEndpoint: &armdiscovery.PrivateEndpoint{
		// 									ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/resourceGroups/rgdiscovery/providers/Microsoft.Network/privateEndpoints/privateEndpoint1"),
		// 								},
		// 								PrivateLinkServiceConnectionState: &armdiscovery.PrivateLinkServiceConnectionState{
		// 									Status: to.Ptr(armdiscovery.PrivateEndpointServiceConnectionStatusPending),
		// 									Description: to.Ptr("km"),
		// 									ActionsRequired: to.Ptr("xbshniighjomlygqk"),
		// 								},
		// 								ProvisioningState: to.Ptr(armdiscovery.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 							},
		// 							ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/bookshelves/bookshelves1/privateEndpointConnections/privateEndpointConnection1"),
		// 							Name: to.Ptr("wdmo"),
		// 							Type: to.Ptr("qpcbcqayctbhbbtguk"),
		// 							SystemData: &armdiscovery.SystemData{
		// 								CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
		// 								CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
		// 								CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
		// 								LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
		// 								LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
		// 								LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
		// 							},
		// 						},
		// 					},
		// 					PublicNetworkAccess: to.Ptr(armdiscovery.PublicNetworkAccessEnabled),
		// 					PrivateEndpointSubnetID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Network/virtualNetworks/virtualnetwork1/subnets/privateEndpointSubnet1"),
		// 					SearchSubnetID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Network/virtualNetworks/virtualnetwork1/subnets/searchSubnet1"),
		// 					ManagedResourceGroup: to.Ptr("tgyjltwleweipcypdzvq"),
		// 					ManagedOnBehalfOfConfiguration: &armdiscovery.WithMoboBrokerResources{
		// 						MoboBrokerResources: []*armdiscovery.MoboBrokerResource{
		// 							{
		// 								ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Storage/storageAccounts/storage1"),
		// 							},
		// 						},
		// 					},
		// 					BookshelfURI: to.Ptr("https://microsoft.com/a"),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
