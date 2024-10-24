package armmongocluster_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster"
)

// Generated from example definition: D:/ws/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster/TempTypeSpecFiles/DocumentDB.MongoCluster.Management/examples/2024-07-01/MongoClusters_PrivateEndpointConnectionPut.json
func ExamplePrivateEndpointConnectionsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmongocluster.NewClientFactory("ffffffff-ffff-ffff-ffff-ffffffffffff", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreate(ctx, "TestGroup", "myMongoCluster", "pecTest", armmongocluster.PrivateEndpointConnectionResource{
		Properties: &armmongocluster.PrivateEndpointConnectionProperties{
			PrivateLinkServiceConnectionState: &armmongocluster.PrivateLinkServiceConnectionState{
				Status:      to.Ptr(armmongocluster.PrivateEndpointServiceConnectionStatusApproved),
				Description: to.Ptr("Auto-Approved"),
			},
		},
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
	// res = armmongocluster.PrivateEndpointConnectionsClientCreateResponse{
	// 	PrivateEndpointConnectionResource: &armmongocluster.PrivateEndpointConnectionResource{
	// 		Name: to.Ptr("pecTest.5d393f64-ef64-46d0-9959-308321c44ac0"),
	// 		ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.DocumentDB/mongoClusters/myMongoCluster/privateEndpointConnections/pecTest.5d393f64-ef64-46d0-9959-308321c44ac0"),
	// 		Type: to.Ptr("Microsoft.DocumentDB/mongoClusters/privateEndpointConnections"),
	// 		SystemData: &armmongocluster.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-09T05:51:31.1386869Z"); return t}()),
	// 			CreatedBy: to.Ptr("2df9eb86-36b5-49dc-86ae-9a63135bfa8c"),
	// 			CreatedByType: to.Ptr(armmongocluster.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2024-02-09T05:51:31.1386869Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("2ff9eb86-36b5-49dc-86ae-9a63135bfa8c"),
	// 			LastModifiedByType: to.Ptr(armmongocluster.CreatedByTypeApplication),
	// 		},
	// 		Properties: &armmongocluster.PrivateEndpointConnectionProperties{
	// 			ProvisioningState: to.Ptr(armmongocluster.PrivateEndpointConnectionProvisioningStateSucceeded),
	// 			PrivateEndpoint: &armmongocluster.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/TestGroup/providers/Microsoft.Network/privateEndpoints/pecTest"),
	// 			},
	// 			GroupIDs: []*string{
	// 				to.Ptr("MongoCluster"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armmongocluster.PrivateLinkServiceConnectionState{
	// 				Status: to.Ptr(armmongocluster.PrivateEndpointServiceConnectionStatusApproved),
	// 				Description: to.Ptr("Approved by admin"),
	// 				ActionsRequired: to.Ptr("None"),
	// 			},
	// 		},
	// 	},
	// }
}
