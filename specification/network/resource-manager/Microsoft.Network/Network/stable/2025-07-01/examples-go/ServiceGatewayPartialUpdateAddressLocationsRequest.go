package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v10"
)

// Generated from example definition: 2025-07-01/ServiceGatewayPartialUpdateAddressLocationsRequest.json
func ExampleServiceGatewaysClient_BeginUpdateAddressLocations_partialUpdateCreateUpdateOrDeleteAddressLocationsInTheServiceGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
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
							to.Ptr("Service1"),
						},
					},
				},
			},
			{
				AddressLocation:     to.Ptr("192.0.0.2"),
				AddressUpdateAction: to.Ptr(armnetwork.AddressUpdateActionPartialUpdate),
				Addresses: []*armnetwork.ServiceGatewayAddress{
					{
						Address: to.Ptr("10.0.0.5"),
						Services: []*string{
							to.Ptr("Service2"),
						},
					},
					{
						Address: to.Ptr("10.0.0.6"),
					},
				},
			},
			{
				AddressLocation: to.Ptr("192.0.0.3"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
}
