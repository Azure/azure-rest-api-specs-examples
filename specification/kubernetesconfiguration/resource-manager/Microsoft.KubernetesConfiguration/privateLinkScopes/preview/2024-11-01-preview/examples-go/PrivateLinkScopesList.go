package armprivatelinkscopes_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armprivatelinkscopes"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4a2bb0762eaad11e725516708483598e0c12cabb/specification/kubernetesconfiguration/resource-manager/Microsoft.KubernetesConfiguration/privateLinkScopes/preview/2024-11-01-preview/examples/PrivateLinkScopesList.json
func ExampleClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armprivatelinkscopes.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewClient().NewListPager(nil)
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
		// page.KubernetesConfigurationPrivateLinkScopeListResult = armprivatelinkscopes.KubernetesConfigurationPrivateLinkScopeListResult{
		// 	Value: []*armprivatelinkscopes.KubernetesConfigurationPrivateLinkScope{
		// 		{
		// 			Name: to.Ptr("my-privatelinkscope"),
		// 			Type: to.Ptr("Microsoft.KubernetesConfiguration/privateLinkScopes"),
		// 			ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/Microsoft.KubernetesConfiguration/privateLinkScopes/my-privatelinkscope"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armprivatelinkscopes.KubernetesConfigurationPrivateLinkScopeProperties{
		// 				ClusterResourceID: to.Ptr("/subscriptions/e9c17b5c-b7ef-4c29-aae7-9338ed5dcb43/resourceGroups/my-resource-group/providers/Microsoft.Kubernetes/ConnectedClusters/my-clusterName"),
		// 				PrivateLinkScopeID: to.Ptr("e5dc51d3-92ed-4d7e-947a-775ea79b4919"),
		// 				ProvisioningState: to.Ptr(armprivatelinkscopes.ProvisioningStateSucceeded),
		// 				PublicNetworkAccess: to.Ptr(armprivatelinkscopes.PublicNetworkAccessTypeDisabled),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("my-other-privatelinkscope"),
		// 			Type: to.Ptr("Microsoft.KubernetesConfiguration/privateLinkScopes"),
		// 			ID: to.Ptr("/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4919/resourceGroups/my-resource-group/providers/Microsoft.KubernetesConfiguration/privateLinkScopes/my-other-privatelinkscope"),
		// 			Location: to.Ptr("westus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Properties: &armprivatelinkscopes.KubernetesConfigurationPrivateLinkScopeProperties{
		// 				ClusterResourceID: to.Ptr("/subscriptions/e9c17b5c-b7ef-4c29-aae7-9338ed5dcb43/resourceGroups/my-resource-group/providers/Microsoft.Kubernetes/ConnectedClusters/my-clusterName"),
		// 				PrivateEndpointConnections: []*armprivatelinkscopes.PrivateEndpointConnection{
		// 					{
		// 						Name: to.Ptr("private-endpoint-connection-name"),
		// 						Type: to.Ptr("Microsoft.KubernetesConfiguration/privateLinkScopes/privateEndpointConnections"),
		// 						ID: to.Ptr("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.KubernetesConfiguration/privateLinkScopes/myPrivateLinkScope/privateEndpointConnections/private-endpoint-connection-name"),
		// 						Properties: &armprivatelinkscopes.PrivateEndpointConnectionProperties{
		// 							PrivateEndpoint: &armprivatelinkscopes.PrivateEndpoint{
		// 								ID: to.Ptr("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-Network/providers/Microsoft.Network/privateEndpoints/private-endpoint-name"),
		// 							},
		// 							PrivateLinkServiceConnectionState: &armprivatelinkscopes.PrivateLinkServiceConnectionState{
		// 								Description: to.Ptr("Auto-approved"),
		// 								ActionsRequired: to.Ptr("None"),
		// 								Status: to.Ptr(armprivatelinkscopes.PrivateEndpointServiceConnectionStatusApproved),
		// 							},
		// 							ProvisioningState: to.Ptr(armprivatelinkscopes.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 						},
		// 				}},
		// 				PrivateLinkScopeID: to.Ptr("f5dc51d3-92ed-4d7e-947a-775ea79b4919"),
		// 				ProvisioningState: to.Ptr(armprivatelinkscopes.ProvisioningStateSucceeded),
		// 				PublicNetworkAccess: to.Ptr(armprivatelinkscopes.PublicNetworkAccessTypeDisabled),
		// 			},
		// 	}},
		// }
	}
}
