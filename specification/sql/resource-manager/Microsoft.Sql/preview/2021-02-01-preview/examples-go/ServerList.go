package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/08894fa8d66cb44dc62a73f7a09530f905985fa3/specification/sql/resource-manager/Microsoft.Sql/preview/2021-02-01-preview/examples/ServerList.json
func ExampleServersClient_NewListPager_listServers() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServersClient().NewListPager(&armsql.ServersClientListOptions{Expand: nil})
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
		// page.ServerListResult = armsql.ServerListResult{
		// 	Value: []*armsql.Server{
		// 		{
		// 			Name: to.Ptr("sqlcrudtest-4645"),
		// 			Type: to.Ptr("Microsoft.Sql/servers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645"),
		// 			Location: to.Ptr("japaneast"),
		// 			Kind: to.Ptr("v12.0"),
		// 			Properties: &armsql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("dummylogin"),
		// 				FullyQualifiedDomainName: to.Ptr("sqlcrudtest-4645.database.windows.net"),
		// 				PrivateEndpointConnections: []*armsql.ServerPrivateEndpointConnection{
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645/privateEndpointConnections/private-endpoint-name-00000000-1111-2222-3333-444444444444"),
		// 						Properties: &armsql.PrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armsql.PrivateEndpointProperty{
		// 								ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armsql.PrivateLinkServiceConnectionStateProperty{
		// 								Description: to.Ptr("Auto-approved"),
		// 								ActionsRequired: to.Ptr(armsql.PrivateLinkServiceConnectionStateActionsRequireNone),
		// 								Status: to.Ptr(armsql.PrivateLinkServiceConnectionStateStatusApproved),
		// 							},
		// 							ProvisioningState: to.Ptr(armsql.PrivateEndpointProvisioningState("Succeeded")),
		// 						},
		// 				}},
		// 				PublicNetworkAccess: to.Ptr(armsql.ServerNetworkAccessFlagEnabled),
		// 				RestrictOutboundNetworkAccess: to.Ptr(armsql.ServerNetworkAccessFlagEnabled),
		// 				State: to.Ptr("Ready"),
		// 				Version: to.Ptr("12.0"),
		// 				WorkspaceFeature: to.Ptr(armsql.ServerWorkspaceFeatureConnected),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sqlcrudtest-6661"),
		// 			Type: to.Ptr("Microsoft.Sql/servers"),
		// 			ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-6661"),
		// 			Location: to.Ptr("japaneast"),
		// 			Kind: to.Ptr("v12.0"),
		// 			Properties: &armsql.ServerProperties{
		// 				AdministratorLogin: to.Ptr("dummylogin"),
		// 				FullyQualifiedDomainName: to.Ptr("sqlcrudtest-6661.database.windows.net"),
		// 				PrivateEndpointConnections: []*armsql.ServerPrivateEndpointConnection{
		// 					{
		// 						ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645/privateEndpointConnections/private-endpoint-name-00000000-1111-2222-3333-444444444444"),
		// 						Properties: &armsql.PrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armsql.PrivateEndpointProperty{
		// 								ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armsql.PrivateLinkServiceConnectionStateProperty{
		// 								Description: to.Ptr("Auto-approved"),
		// 								ActionsRequired: to.Ptr(armsql.PrivateLinkServiceConnectionStateActionsRequireNone),
		// 								Status: to.Ptr(armsql.PrivateLinkServiceConnectionStateStatusApproved),
		// 							},
		// 							ProvisioningState: to.Ptr(armsql.PrivateEndpointProvisioningState("Succeeded")),
		// 						},
		// 				}},
		// 				PublicNetworkAccess: to.Ptr(armsql.ServerNetworkAccessFlagEnabled),
		// 				RestrictOutboundNetworkAccess: to.Ptr(armsql.ServerNetworkAccessFlagEnabled),
		// 				State: to.Ptr("Ready"),
		// 				Version: to.Ptr("12.0"),
		// 				WorkspaceFeature: to.Ptr(armsql.ServerWorkspaceFeatureConnected),
		// 			},
		// 	}},
		// }
	}
}
