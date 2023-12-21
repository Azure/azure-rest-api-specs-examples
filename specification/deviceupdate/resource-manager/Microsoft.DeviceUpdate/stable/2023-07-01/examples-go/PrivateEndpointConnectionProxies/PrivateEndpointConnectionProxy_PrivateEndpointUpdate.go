package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b8c74fd80b415fa1ebb6fa787d454694c39e0fd5/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_PrivateEndpointUpdate.json
func ExamplePrivateEndpointConnectionProxiesClient_UpdatePrivateEndpointProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPrivateEndpointConnectionProxiesClient().UpdatePrivateEndpointProperties(ctx, "test-rg", "contoso", "peexample01", armdeviceupdate.PrivateEndpointUpdate{
		ID:                      to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
		ImmutableResourceID:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
		ImmutableSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		Location:                to.Ptr("westus2"),
		VnetTrafficTag:          to.Ptr("12345678"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
