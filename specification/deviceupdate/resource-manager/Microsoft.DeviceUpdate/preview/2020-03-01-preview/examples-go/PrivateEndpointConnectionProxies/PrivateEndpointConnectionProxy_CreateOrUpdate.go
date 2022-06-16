package armdeviceupdate_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2020-03-01-preview/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_CreateOrUpdate.json
func ExamplePrivateEndpointConnectionProxiesClient_BeginCreateOrUpdate() {
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
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<private-endpoint-connection-proxy-id>",
		armdeviceupdate.PrivateEndpointConnectionProxy{
			RemotePrivateEndpoint: &armdeviceupdate.RemotePrivateEndpoint{
				ID:                      to.Ptr("<id>"),
				ImmutableResourceID:     to.Ptr("<immutable-resource-id>"),
				ImmutableSubscriptionID: to.Ptr("<immutable-subscription-id>"),
				Location:                to.Ptr("<location>"),
				ManualPrivateLinkServiceConnections: []*armdeviceupdate.PrivateLinkServiceConnection{
					{
						Name: to.Ptr("<name>"),
						GroupIDs: []*string{
							to.Ptr("DeviceUpdate")},
						RequestMessage: to.Ptr("<request-message>"),
					}},
				PrivateLinkServiceProxies: []*armdeviceupdate.PrivateLinkServiceProxy{
					{
						GroupConnectivityInformation: []*armdeviceupdate.GroupConnectivityInformation{},
						ID:                           to.Ptr("<id>"),
					}},
			},
		},
		&armdeviceupdate.PrivateEndpointConnectionProxiesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}
