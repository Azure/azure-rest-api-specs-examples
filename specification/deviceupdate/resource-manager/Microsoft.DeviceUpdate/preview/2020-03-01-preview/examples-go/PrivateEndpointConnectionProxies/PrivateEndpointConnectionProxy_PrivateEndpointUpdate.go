package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2020-03-01-preview/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_PrivateEndpointUpdate.json
func ExamplePrivateEndpointConnectionProxiesClient_UpdatePrivateEndpointProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armdeviceupdate.NewPrivateEndpointConnectionProxiesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	_, err = client.UpdatePrivateEndpointProperties(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<private-endpoint-connection-proxy-id>",
		armdeviceupdate.PrivateEndpointUpdate{
			ID:                      to.Ptr("<id>"),
			ImmutableResourceID:     to.Ptr("<immutable-resource-id>"),
			ImmutableSubscriptionID: to.Ptr("<immutable-subscription-id>"),
			Location:                to.Ptr("<location>"),
			VnetTrafficTag:          to.Ptr("<vnet-traffic-tag>"),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
}
