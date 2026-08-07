package armdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/discovery/armdiscovery"
)

// Generated from example definition: 2026-06-01/BookshelfPrivateEndpointConnections_Get_MaximumSet_Gen.json
func ExampleBookshelfPrivateEndpointConnectionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdiscovery.NewClientFactory("A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBookshelfPrivateEndpointConnectionsClient().Get(ctx, "rgdiscovery", "b9893e75cf964912a2", "connection", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdiscovery.BookshelfPrivateEndpointConnectionsClientGetResponse{
	// 	BookshelfPrivateEndpointConnection: armdiscovery.BookshelfPrivateEndpointConnection{
	// 		ID: to.Ptr("/subscriptions/A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/bookshelves/b9893e75cf964912a2/privateEndpointConnections/connection"),
	// 		Name: to.Ptr("connection"),
	// 		Type: to.Ptr("Microsoft.Discovery/bookshelves/privateEndpointConnections"),
	// 		SystemData: &armdiscovery.SystemData{
	// 			CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
	// 			CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
	// 			LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
	// 		},
	// 		Properties: &armdiscovery.PrivateEndpointConnectionProperties{
	// 			GroupIDs: []*string{
	// 				to.Ptr("bpyliugtuio"),
	// 			},
	// 			PrivateEndpoint: &armdiscovery.PrivateEndpoint{
	// 				ID: to.Ptr("bzrnbivshkunzw"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armdiscovery.PrivateLinkServiceConnectionState{
	// 				Status: to.Ptr(armdiscovery.PrivateEndpointServiceConnectionStatusPending),
	// 				Description: to.Ptr("km"),
	// 				ActionsRequired: to.Ptr("xbshniighjomlygqk"),
	// 			},
	// 			ProvisioningState: to.Ptr(armdiscovery.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 		},
	// 	},
	// }
}
