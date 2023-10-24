package armelasticsan_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/3066a973f4baf2e2bf072a013b585a820bb10146/specification/elasticsan/resource-manager/Microsoft.ElasticSan/stable/2023-01-01/examples/PrivateEndpointConnections_List_MaximumSet_Gen.json
func ExamplePrivateEndpointConnectionsClient_NewListPager_privateEndpointConnectionsListMaximumSetGen() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armelasticsan.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionsClient().NewListPager("resourcegroupname", "elasticsanname", nil)
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
		// page.PrivateEndpointConnectionListResult = armelasticsan.PrivateEndpointConnectionListResult{
		// 	Value: []*armelasticsan.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("vyzqckpcwufpvalbspekxikt"),
		// 			Type: to.Ptr("ldolsnjwzutewucdfessitnxqb"),
		// 			ID: to.Ptr("ynin"),
		// 			SystemData: &armelasticsan.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T17:40:52.453Z"); return t}()),
		// 				CreatedBy: to.Ptr("bgurjvijz"),
		// 				CreatedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2023-08-25T17:40:52.453Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("uvexylihjrtinzkeluohusnaxatfqh"),
		// 				LastModifiedByType: to.Ptr(armelasticsan.CreatedByTypeUser),
		// 			},
		// 			Properties: &armelasticsan.PrivateEndpointConnectionProperties{
		// 				GroupIDs: []*string{
		// 					to.Ptr("jdwrzpemdjrpiwzvy")},
		// 					PrivateEndpoint: &armelasticsan.PrivateEndpoint{
		// 						ID: to.Ptr("gktekmqchmjqxhfvywq"),
		// 					},
		// 					PrivateLinkServiceConnectionState: &armelasticsan.PrivateLinkServiceConnectionState{
		// 						Description: to.Ptr("dxl"),
		// 						ActionsRequired: to.Ptr("jhjdpwvyzipggtn"),
		// 						Status: to.Ptr(armelasticsan.PrivateEndpointServiceConnectionStatusPending),
		// 					},
		// 					ProvisioningState: to.Ptr(armelasticsan.ProvisioningStatesSucceeded),
		// 				},
		// 		}},
		// 	}
	}
}
