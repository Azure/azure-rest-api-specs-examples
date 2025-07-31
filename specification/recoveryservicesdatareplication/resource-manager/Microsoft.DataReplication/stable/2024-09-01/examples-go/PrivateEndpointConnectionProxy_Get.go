package armrecoveryservicesdatareplication_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservicesdatareplication/armrecoveryservicesdatareplication"
)

// Generated from example definition: 2024-09-01/PrivateEndpointConnectionProxy_Get.json
func ExamplePrivateEndpointConnectionProxiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrecoveryservicesdatareplication.NewClientFactory("930CEC23-4430-4513-B855-DBA237E2F3BF", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionProxiesClient().Get(ctx, "rgswagger_2024-09-01", "4", "d", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armrecoveryservicesdatareplication.PrivateEndpointConnectionProxiesClientGetResponse{
	// 	PrivateEndpointConnectionProxy: &armrecoveryservicesdatareplication.PrivateEndpointConnectionProxy{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataReplication/replicationVaults/vault1/privateEndpointConnectionProxies/proxy1"),
	// 		Name: to.Ptr("wrbeymbilwm"),
	// 		Type: to.Ptr("xpfgxxykisvjcifsnlvtgjakwifak"),
	// 		Etag: to.Ptr("hruibxrezstxroxrxweh"),
	// 		Properties: &armrecoveryservicesdatareplication.PrivateEndpointConnectionProxyProperties{
	// 			RemotePrivateEndpoint: &armrecoveryservicesdatareplication.RemotePrivateEndpoint{
	// 				ID: to.Ptr("yipalno"),
	// 				PrivateLinkServiceConnections: []*armrecoveryservicesdatareplication.PrivateLinkServiceConnection{
	// 					{
	// 						Name: to.Ptr("jqwntlzfsksl"),
	// 						GroupIDs: []*string{
	// 							to.Ptr("hvejynjktikteipnioyeja"),
	// 						},
	// 						RequestMessage: to.Ptr("bukgzpkvcvfbmcdmpcbiigbvugicqa"),
	// 					},
	// 				},
	// 				ManualPrivateLinkServiceConnections: []*armrecoveryservicesdatareplication.PrivateLinkServiceConnection{
	// 					{
	// 						Name: to.Ptr("jqwntlzfsksl"),
	// 						GroupIDs: []*string{
	// 							to.Ptr("hvejynjktikteipnioyeja"),
	// 						},
	// 						RequestMessage: to.Ptr("bukgzpkvcvfbmcdmpcbiigbvugicqa"),
	// 					},
	// 				},
	// 				PrivateLinkServiceProxies: []*armrecoveryservicesdatareplication.PrivateLinkServiceProxy{
	// 					{
	// 						ID: to.Ptr("nzqxevuyqeedrqnkbnlcyrrrbzxvl"),
	// 						RemotePrivateLinkServiceConnectionState: &armrecoveryservicesdatareplication.PrivateLinkServiceConnectionState{
	// 							Status: to.Ptr(armrecoveryservicesdatareplication.PrivateEndpointConnectionStatusApproved),
	// 							Description: to.Ptr("y"),
	// 							ActionsRequired: to.Ptr("afwbq"),
	// 						},
	// 						RemotePrivateEndpointConnection: &armrecoveryservicesdatareplication.RemotePrivateEndpointConnection{
	// 							ID: to.Ptr("ocunsgawjsqohkrcyxiv"),
	// 						},
	// 						GroupConnectivityInformation: []*armrecoveryservicesdatareplication.GroupConnectivityInformation{
	// 							{
	// 								GroupID: to.Ptr("per"),
	// 								MemberName: to.Ptr("ybptuypgdqoxkuwqx"),
	// 								CustomerVisibleFqdns: []*string{
	// 									to.Ptr("vedcg"),
	// 								},
	// 								InternalFqdn: to.Ptr("maqavwhxwzzhbzjbryyquvitmup"),
	// 								RedirectMapID: to.Ptr("pezncxcq"),
	// 								PrivateLinkServiceArmRegion: to.Ptr("rerkqqxinteevmlbrdkktaqhcch"),
	// 							},
	// 						},
	// 					},
	// 				},
	// 				ConnectionDetails: []*armrecoveryservicesdatareplication.ConnectionDetails{
	// 					{
	// 						ID: to.Ptr("lenqkogzkes"),
	// 						PrivateIPAddress: to.Ptr("cyiacdzzyqmxjpijjbwgasegehtqe"),
	// 						LinkIdentifier: to.Ptr("ravfufhkdowufd"),
	// 						GroupID: to.Ptr("pjrlygpadir"),
	// 						MemberName: to.Ptr("ybuysjrlfupewxe"),
	// 					},
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
