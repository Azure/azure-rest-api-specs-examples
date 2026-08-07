package armdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/discovery/armdiscovery"
)

// Generated from example definition: 2026-06-01/Bookshelves_Update_MaximumSet_Gen.json
func ExampleBookshelvesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdiscovery.NewClientFactory("A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBookshelvesClient().BeginUpdate(ctx, "rgdiscovery", "14964dff7a049b02ad", armdiscovery.Bookshelf{
		Properties: &armdiscovery.BookshelfProperties{
			KeyVaultProperties: &armdiscovery.BookshelfKeyVaultProperties{
				KeyName:    to.Ptr("rioczxrgqxcnesqxnxpuc"),
				KeyVersion: to.Ptr("lhpxvapkhljzkdt"),
			},
			PublicNetworkAccess: to.Ptr(armdiscovery.PublicNetworkAccessEnabled),
		},
		Tags: map[string]*string{
			"key5254": to.Ptr("fozqmnqttenfggdjxalzycvqqzfe"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdiscovery.BookshelvesClientUpdateResponse{
	// 	Bookshelf: armdiscovery.Bookshelf{
	// 		ID: to.Ptr("/subscriptions/A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/bookshelves/14964dff7a049b02ad"),
	// 		Name: to.Ptr("14964dff7a049b02ad"),
	// 		Tags: map[string]*string{
	// 			"key3815": to.Ptr("lrfxoxtbcgjpokrmtlzvknkcu"),
	// 		},
	// 		Location: to.Ptr("uksouth"),
	// 		Type: to.Ptr("Microsoft.Discovery/bookshelves"),
	// 		SystemData: &armdiscovery.SystemData{
	// 			CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
	// 			CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
	// 			LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 		},
	// 		Properties: &armdiscovery.BookshelfProperties{
	// 			ProvisioningState: to.Ptr(armdiscovery.ProvisioningStateSucceeded),
	// 			WorkloadIdentities: map[string]*armdiscovery.UserAssignedIdentity{
	// 				"key3650": &armdiscovery.UserAssignedIdentity{
	// 					PrincipalID: to.Ptr("00000011-1111-2222-2222-123456789111"),
	// 					ClientID: to.Ptr("00000011-1111-2222-2222-123456789111"),
	// 				},
	// 			},
	// 			CustomerManagedKeys: to.Ptr(armdiscovery.CustomerManagedKeysEnabled),
	// 			KeyVaultProperties: &armdiscovery.BookshelfKeyVaultProperties{
	// 				KeyVaultURI: to.Ptr("https://microsoft.com/a"),
	// 				KeyName: to.Ptr("rioczxrgqxcnesqxnxpuc"),
	// 				KeyVersion: to.Ptr("lhpxvapkhljzkdt"),
	// 				IdentityClientID: to.Ptr("00000011-1111-2222-2222-123456789111"),
	// 			},
	// 			LogAnalyticsClusterID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.OperationalInsights/clusters/cluster1"),
	// 			PrivateEndpointConnections: []*armdiscovery.PrivateEndpointConnection{
	// 				{
	// 					Properties: &armdiscovery.PrivateEndpointConnectionProperties{
	// 						GroupIDs: []*string{
	// 							to.Ptr("bpyliugtuio"),
	// 						},
	// 						PrivateEndpoint: &armdiscovery.PrivateEndpoint{
	// 							ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/resourceGroups/rgdiscovery/providers/Microsoft.Network/privateEndpoints/privateEndpoint1"),
	// 						},
	// 						PrivateLinkServiceConnectionState: &armdiscovery.PrivateLinkServiceConnectionState{
	// 							Status: to.Ptr(armdiscovery.PrivateEndpointServiceConnectionStatusPending),
	// 							Description: to.Ptr("km"),
	// 							ActionsRequired: to.Ptr("xbshniighjomlygqk"),
	// 						},
	// 						ProvisioningState: to.Ptr(armdiscovery.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 					},
	// 					ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/bookshelves/bookshelves1/privateEndpointConnections/privateEndpointConnection1"),
	// 					Name: to.Ptr("wdmo"),
	// 					Type: to.Ptr("qpcbcqayctbhbbtguk"),
	// 					SystemData: &armdiscovery.SystemData{
	// 						CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
	// 						CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 						LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
	// 						LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 					},
	// 				},
	// 			},
	// 			PublicNetworkAccess: to.Ptr(armdiscovery.PublicNetworkAccessEnabled),
	// 			PrivateEndpointSubnetID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Network/virtualNetworks/virtualnetwork1/subnets/privateEndpointSubnet1"),
	// 			SearchSubnetID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Network/virtualNetworks/virtualnetwork1/subnets/searchSubnet1"),
	// 			ManagedResourceGroup: to.Ptr("tgyjltwleweipcypdzvq"),
	// 			ManagedOnBehalfOfConfiguration: &armdiscovery.WithMoboBrokerResources{
	// 				MoboBrokerResources: []*armdiscovery.MoboBrokerResource{
	// 					{
	// 						ID: to.Ptr("/subscriptions/31735C59-6307-4464-8B80-3675223F23D2/providers/Microsoft.Storage/storageAccounts/storage1"),
	// 					},
	// 				},
	// 			},
	// 			BookshelfURI: to.Ptr("https://microsoft.com/a"),
	// 		},
	// 	},
	// }
}
