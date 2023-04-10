package armkeyvault_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/0cc5e2efd6ffccf30e80d1e150b488dd87198b94/specification/keyvault/resource-manager/Microsoft.KeyVault/stable/2021-10-01/examples/ManagedHsm_ListPrivateEndpointConnectionsByResource.json
func ExampleMHSMPrivateEndpointConnectionsClient_NewListByResourcePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkeyvault.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMHSMPrivateEndpointConnectionsClient().NewListByResourcePager("sample-group", "sample-mhsm", nil)
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
		// page.MHSMPrivateEndpointConnectionsListResult = armkeyvault.MHSMPrivateEndpointConnectionsListResult{
		// 	Value: []*armkeyvault.MHSMPrivateEndpointConnection{
		// 		{
		// 			Name: to.Ptr("sample-pec1"),
		// 			Type: to.Ptr("Microsoft.KeyVault/managedhsms/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedhsms/sample-mhsm/privateEndpointConnections/sample-pec1"),
		// 			Etag: to.Ptr(""),
		// 			Properties: &armkeyvault.MHSMPrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armkeyvault.MHSMPrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-1234-000000000000/resourceGroups/sample-group/providers/Microsoft.Network/privateEndpoints/sample-pe1"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armkeyvault.MHSMPrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("This was automatically approved by user1234@contoso.com"),
		// 					ActionsRequired: to.Ptr(armkeyvault.ActionsRequiredNone),
		// 					Status: to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armkeyvault.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("sample-pec2"),
		// 			Type: to.Ptr("Microsoft.KeyVault/managedhsms/privateEndpointConnections"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-group/providers/Microsoft.KeyVault/managedhsms/sample-mhsm/privateEndpointConnections/sample-pec2"),
		// 			Etag: to.Ptr(""),
		// 			Properties: &armkeyvault.MHSMPrivateEndpointConnectionProperties{
		// 				PrivateEndpoint: &armkeyvault.MHSMPrivateEndpoint{
		// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-1234-000000000000/resourceGroups/sample-group/providers/Microsoft.Network/privateEndpoints/sample-pe2"),
		// 				},
		// 				PrivateLinkServiceConnectionState: &armkeyvault.MHSMPrivateLinkServiceConnectionState{
		// 					Description: to.Ptr("This was automatically approved by user1234@contoso.com"),
		// 					ActionsRequired: to.Ptr(armkeyvault.ActionsRequiredNone),
		// 					Status: to.Ptr(armkeyvault.PrivateEndpointServiceConnectionStatusApproved),
		// 				},
		// 				ProvisioningState: to.Ptr(armkeyvault.PrivateEndpointConnectionProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
