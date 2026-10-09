package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan/v2"
)

// Generated from example definition: 2026-05-01-preview/PrivateEndpointConnections_Create_MaximumSet_Gen.json
func ExamplePrivateEndpointConnectionsClient_BeginCreate_privateEndpointConnectionsCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("subscriptionid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreate(ctx, "resourcegroupname", "elasticsanname", "privateendpointconnectionname", armelasticsan.PrivateEndpointConnection{
		Properties: &armelasticsan.PrivateEndpointConnectionProperties{
			GroupIDs: []*string{
				to.Ptr("jdwrzpemdjrpiwzvy"),
			},
			PrivateEndpoint: &armelasticsan.PrivateEndpoint{},
			PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
				Description:     to.Ptr("dxl"),
				ActionsRequired: to.Ptr("jhjdpwvyzipggtn"),
				Status:          to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
			},
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
	// res = armelasticsan.PrivateEndpointConnectionsClientCreateResponse{
	// 	PrivateEndpointConnection: armelasticsan.PrivateEndpointConnection{
	// 		Name: to.Ptr("{privateEndpointConnectionName}"),
	// 		Type: to.Ptr("Microsoft.ElasticSan/elasticSans/privateEndpointConnections"),
	// 		ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/privateEndpointConnections/{privateEndpointConnectionName}"),
	// 		Properties: &armelasticsan.PrivateEndpointConnectionProperties{
	// 			GroupIDs: []*string{
	// 				to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}"),
	// 			},
	// 			PrivateEndpoint: &armelasticsan.PrivateEndpoint{
	// 				ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}"),
	// 			},
	// 			PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
	// 				Description: to.Ptr("Auto-Approved"),
	// 				ActionsRequired: to.Ptr("None"),
	// 				Status: to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
	// 			},
	// 			ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
	// 		},
	// 		SystemData: &armelasticsan.SystemData{
	// 			CreatedAt: to.Ptr(time.Date(2023, time.August, 25, 17, 40, 52, 453000000, time.UTC)),
	// 			CreatedBy: to.Ptr("bgurjvijz"),
	// 			CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(time.Date(2023, time.August, 25, 17, 40, 52, 453000000, time.UTC)),
	// 			LastModifiedBy: to.Ptr("uvexylihjrtinzkeluohusnaxatfqh"),
	// 			LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		},
	// 	},
	// }
}
