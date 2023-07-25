package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c53808ba54beef57059371708f1fa6949a11a280/specification/elasticsan/resource-manager/Microsoft.ElasticSan/preview/2022-12-01-preview/examples/PrivateEndpointConnections_Create_MaximumSet_Gen.json
func ExamplePrivateEndpointConnectionsClient_BeginCreate_privateEndpointConnectionsCreateMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionsClient().BeginCreate(ctx, "resourcegroupname", "elasticsanname", "privateendpointconnectionname", armelasticsan.PrivateEndpointConnection{
		Properties: &armelasticsan.PrivateEndpointConnectionProperties{
			GroupIDs: []*string{
				to.Ptr("sytxzqlcoapcaywthgwvwcw")},
			PrivateEndpoint: &armelasticsan.PrivateEndpoint{},
			PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
				Description:     to.Ptr("Auto-Approved"),
				ActionsRequired: to.Ptr("None"),
				Status:          to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
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
	// res.PrivateEndpointConnection = armelasticsan.PrivateEndpointConnection{
	// 	Name: to.Ptr("{privateEndpointConnectionName}"),
	// 	Type: to.Ptr("Microsoft.ElasticSan/elasticSans/privateEndpointConnections"),
	// 	ID: to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ElasticSan/elasticSans/{elasticSanName}/privateEndpointConnections/{privateEndpointConnectionName}"),
	// 	SystemData: &armelasticsan.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T09:59:45.919Z"); return t}()),
	// 		CreatedBy: to.Ptr("otfifnrahdshqombvtg"),
	// 		CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-07-03T09:59:45.919Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("jnaxavnlhrboshtidtib"),
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
