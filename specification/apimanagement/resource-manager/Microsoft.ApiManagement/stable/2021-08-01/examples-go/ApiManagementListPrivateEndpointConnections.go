package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListPrivateEndpointConnections.json
func ExamplePrivateEndpointConnectionClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionClient().NewListByServicePager("rg1", "apimService1", nil)
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
		// page.PrivateEndpointConnectionListResult = armapimanagement.PrivateEndpointConnectionListResult{
		// 	Value: []*armapimanagement.PrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("privateEndpointProxyName"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/privateEndpointConnections/connectionName"),
		// 			Properties: &armapimanagement.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armapimanagement.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/privateEndpointName"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armapimanagement.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Please approve my request, thanks"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armapimanagement.PrivateEndpointServiceConnectionStatusPending),
		// 				},
		// 				ProvisioningState: to.Ptr(armapimanagement.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("privateEndpointProxyName2"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/privateEndpointConnections/privateEndpointProxyName2"),
		// 			Properties: &armapimanagement.PrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armapimanagement.PrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/privateEndpoints/privateEndpointName2"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armapimanagement.PrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("Please approve my request, thanks"),
		// 					ActionsRequired: to.Ptr("None"),
		// 					Status: to.Ptr(armapimanagement.PrivateEndpointServiceConnectionStatusPending),
		// 				},
		// 				ProvisioningState: to.Ptr(armapimanagement.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
