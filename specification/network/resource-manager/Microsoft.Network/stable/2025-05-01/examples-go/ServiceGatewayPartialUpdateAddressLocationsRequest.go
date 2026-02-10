package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ServiceGatewayPartialUpdateAddressLocationsRequest.json
func ExampleServiceGatewaysClient_BeginUpdateAddressLocations_partialUpdateCreateUpdateOrDeleteAddressLocationsInTheServiceGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewServiceGatewaysClient().BeginUpdateAddressLocations(ctx, "rg1", "sg", armnetwork.ServiceGatewayUpdateAddressLocationsRequest{
		Action: to.Ptr(armnetwork.UpdateActionPartialUpdate),
		AddressLocations: []*armnetwork.ServiceGatewayAddressLocation{
			{
				AddressLocation:     to.Ptr("192.0.0.1"),
				AddressUpdateAction: to.Ptr(armnetwork.AddressUpdateActionFullUpdate),
				Addresses: []*armnetwork.ServiceGatewayAddress{
					{
						Address: to.Ptr("10.0.0.4"),
						Services: []*string{
							to.Ptr("Service1")},
					}},
			},
			{
				AddressLocation:     to.Ptr("192.0.0.2"),
				AddressUpdateAction: to.Ptr(armnetwork.AddressUpdateActionPartialUpdate),
				Addresses: []*armnetwork.ServiceGatewayAddress{
					{
						Address: to.Ptr("10.0.0.5"),
						Services: []*string{
							to.Ptr("Service2")},
					},
					{
						Address: to.Ptr("10.0.0.6"),
					}},
			},
			{
				AddressLocation: to.Ptr("192.0.0.3"),
			}},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
