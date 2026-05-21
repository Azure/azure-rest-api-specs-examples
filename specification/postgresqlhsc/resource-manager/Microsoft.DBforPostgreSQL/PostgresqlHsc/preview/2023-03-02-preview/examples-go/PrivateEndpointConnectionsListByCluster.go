package armpostgresqlhsc_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc"
)

// Generated from example definition: 2023-03-02-preview/PrivateEndpointConnectionsListByCluster.json
func ExamplePrivateEndpointConnectionsClient_NewListByClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armpostgresqlhsc.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListByClusterPager("TestResourceGroup", "testcluster", nil)
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
		// page = armpostgresqlhsc.PrivateEndpointConnectionsClientListByClusterResponse{
		// 	PrivateEndpointConnectionListResult: armpostgresqlhsc.PrivateEndpointConnectionListResult{
		// 		Value: []*armpostgresqlhsc.PrivateEndpointConnection{
		// 			{
		// 				Name: to.Ptr("private-endpoint-connection-name"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/privateEndpointConnections/private-endpoint-connection-name"),
		// 				Properties: &armpostgresqlhsc.PrivateEndpointConnectionProperties{
		// 					GroupIDs: []*string{
		// 						to.Ptr("coordinator"),
		// 					},
		// 					PrivateEndpoint: &armpostgresqlhsc.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armpostgresqlhsc.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Auto-approved"),
		// 						ActionsRequired: to.Ptr("None"),
		// 						Status: to.Ptr(armpostgresqlhsc.PrivateEndpointServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armpostgresqlhsc.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 				SystemData: &armpostgresqlhsc.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-01T17:18:19.1234567Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-02T17:18:19.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("private-endpoint-connection-name-2"),
		// 				Type: to.Ptr("Microsoft.DBforPostgreSQL/serverGroupsv2/privateEndpointConnections"),
		// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestResourceGroup/providers/Microsoft.DBforPostgreSQL/serverGroupsv2/testcluster/privateEndpointConnections/private-endpoint-connection-name-2"),
		// 				Properties: &armpostgresqlhsc.PrivateEndpointConnectionProperties{
		// 					GroupIDs: []*string{
		// 						to.Ptr("worker-0"),
		// 					},
		// 					PrivateEndpoint: &armpostgresqlhsc.PrivateEndpoint{
		// 						ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name-2"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armpostgresqlhsc.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("Auto-approved"),
		// 						ActionsRequired: to.Ptr("None"),
		// 						Status: to.Ptr(armpostgresqlhsc.PrivateEndpointServiceConnectionStatusApproved),
		// 					},
		// 					ProvisioningState: to.Ptr(armpostgresqlhsc.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 				},
		// 				SystemData: &armpostgresqlhsc.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-01T17:18:19.1234567Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-01-02T17:18:19.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armpostgresqlhsc.CreatedByTypeUser),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
