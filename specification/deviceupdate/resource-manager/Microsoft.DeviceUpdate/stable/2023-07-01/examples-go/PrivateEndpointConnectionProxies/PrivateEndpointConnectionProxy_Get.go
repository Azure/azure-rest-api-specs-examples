package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_Get.json
func ExamplePrivateEndpointConnectionProxiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionProxiesClient().Get(ctx, "test-rg", "contoso", "peexample01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnectionProxy = armdeviceupdate.PrivateEndpointConnectionProxy{
	// 	RemotePrivateEndpoint: &armdeviceupdate.RemotePrivateEndpoint{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
	// 		ImmutableResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
	// 		ImmutableSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		Location: to.Ptr("westus2"),
	// 		ManualPrivateLinkServiceConnections: []*armdeviceupdate.PrivateLinkServiceConnection{
	// 			{
	// 				Name: to.Ptr("{plsConnectionName}"),
	// 				GroupIDs: []*string{
	// 					to.Ptr("DeviceUpdate")},
	// 					RequestMessage: to.Ptr("Please approve my connection, thanks."),
	// 			}},
	// 			PrivateLinkServiceProxies: []*armdeviceupdate.PrivateLinkServiceProxy{
	// 				{
	// 					GroupConnectivityInformation: []*armdeviceupdate.GroupConnectivityInformation{
	// 						{
	// 							GroupID: to.Ptr("DeviceUpdate"),
	// 							MemberName: to.Ptr("adu"),
	// 					}},
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}/privateLinkServiceProxies/{privateEndpointConnectionProxyId}"),
	// 			}},
	// 		},
	// 		Name: to.Ptr("peexample01"),
	// 		Type: to.Ptr("Microsoft.DeviceUpdate/accounts/privateEndpointConnectionProxies"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DeviceUpdate/accounts/contoso/privateEndpointConnectionProxies/peexample01"),
	// 	}
}
