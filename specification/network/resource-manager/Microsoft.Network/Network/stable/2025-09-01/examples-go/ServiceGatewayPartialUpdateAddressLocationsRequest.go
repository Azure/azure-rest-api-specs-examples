package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v11"
)

// Generated from example definition: 2025-09-01/ServiceGatewayPartialUpdateAddressLocationsRequest.json
func ExampleServiceGatewaysClient_UpdateAddressLocations_partialUpdateCreateUpdateOrDeleteAddressLocationsInTheServiceGateway() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewServiceGatewaysClient().UpdateAddressLocations(ctx, "rg1", "sg", armnetwork.ServiceGatewayUpdateAddressLocationsRequest{
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armnetwork.ServiceGatewaysClientUpdateAddressLocationsResponse{
	// 	ServiceGatewayActionOkResponseBody: armnetwork.ServiceGatewayActionOkResponseBody{
	// 	},
	// }
}
