package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/907b79c0a6a660826e54dc1f16ea14b831b201d2/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/PrivateEndpointConnections_Get_MinimumSet_Gen.json
func ExamplePrivateEndpointConnectionsClient_Get_privateEndpointConnectionsGetMinimumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionsClient().Get(ctx, "resourcegroupname", "elasticsanname", "privateendpointconnectionname", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnection = armelasticsan.PrivateEndpointConnection{
	// 	Name: to.Ptr("{privateEndpointConnectionName}"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/privateEndpointConnections/{privateEndpointConnectionName}"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T17:40:52.453Z"); return t}()),
	// 		CreatedBy: to.Ptr("bgurjvijz"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T17:40:52.453Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("uvexylihjrtinzkeluohusnaxatfqh"),
	// 		LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 	},
	// 	Properties: &armelasticsan.PrivateEndpointConnectionProperties{
	// 		GroupIDs: []*string{
	// 			to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/volumegroups/{volumeGroupName}")},
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
	// 	}
}
