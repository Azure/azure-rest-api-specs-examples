package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_Validate.json
func ExamplePrivateEndpointConnectionProxiesClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPrivateEndpointConnectionProxiesClient().Validate(ctx, "test-rg", "contoso", "peexample01", armdeviceupdate.PrivateEndpointConnectionProxy{
		RemotePrivateEndpoint: &armdeviceupdate.RemotePrivateEndpoint{
			ID:                      to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}"),
			ImmutableResourceID:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
			ImmutableSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
			Location:                to.Ptr("westus2"),
			ManualPrivateLinkServiceConnections: []*armdeviceupdate.PrivateLinkServiceConnection{
				{
					Name: to.Ptr("{privateEndpointConnectionProxyId}"),
					GroupIDs: []*string{
						to.Ptr("DeviceUpdate")},
					RequestMessage: to.Ptr("Please approve my connection, thanks."),
				}},
			PrivateLinkServiceProxies: []*armdeviceupdate.PrivateLinkServiceProxy{
				{
					GroupConnectivityInformation: []*armdeviceupdate.GroupConnectivityInformation{},
					ID:                           to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}/privateLinkServiceProxies/{privateEndpointConnectionProxyId}"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
