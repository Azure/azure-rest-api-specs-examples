package armsql_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a3913f4b26467aed413cdc907116e99894f08994/specification/sql/resource-manager/Microsoft.Sql/preview/2023-05-01-preview/examples/ServerGet.json
func ExampleServersClient_Get_getServer() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsql.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServersClient().Get(ctx, "sqlcrudtest-7398", "sqlcrudtest-4645", &armsql.ServersClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Server = armsql.Server{
	// 	Name: to.Ptr("sqlcrudtest-4645"),
	// 	Type: to.Ptr("Microsoft.Sql/servers"),
	// 	ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645"),
	// 	Location: to.Ptr("japaneast"),
	// 	Tags: map[string]*string{
	// 		"tagKey1": to.Ptr("TagValue1"),
	// 	},
	// 	Kind: to.Ptr("v12.0"),
	// 	Properties: &armsql.ServerProperties{
	// 		AdministratorLogin: to.Ptr("dummylogin"),
	// 		FullyQualifiedDomainName: to.Ptr("sqlcrudtest-4645.database.windows.net"),
	// 		IsIPv6Enabled: to.Ptr(armsql.ServerNetworkAccessFlagEnabled),
	// 		PrivateEndpointConnections: []*armsql.ServerPrivateEndpointConnection{
	// 			{
	// 				ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/sqlcrudtest-7398/providers/Microsoft.Sql/servers/sqlcrudtest-4645/privateEndpointConnections/private-endpoint-name-00000000-1111-2222-3333-444444444444"),
	// 				Properties: &armsql.PrivateEndpointConnectionProperties{
	// 					PrivateEndpoint: &armsql.PrivateEndpointProperty{
	// 						ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
	// 					},
	// 					PrivateLinkServiceConnectionState: &armsql.PrivateLinkServiceConnectionStateProperty{
	// 						Description: to.Ptr("Auto-approved"),
	// 						ActionsRequired: to.Ptr(armsql.PrivateLinkServiceConnectionStateActionsRequireNone),
	// 						Status: to.Ptr(armsql.PrivateLinkServiceConnectionStateStatusApproved),
	// 					},
	// 					ProvisioningState: to.Ptr(armsql.PrivateEndpointProvisioningState("Succeeded")),
	// 				},
	// 		}},
	// 		PublicNetworkAccess: to.Ptr(armsql.ServerPublicNetworkAccessFlagEnabled),
	// 		RestrictOutboundNetworkAccess: to.Ptr(armsql.ServerNetworkAccessFlagEnabled),
	// 		State: to.Ptr("Ready"),
	// 		Version: to.Ptr("12.0"),
	// 		WorkspaceFeature: to.Ptr(armsql.ServerWorkspaceFeatureConnected),
	// 	},
	// }
}
