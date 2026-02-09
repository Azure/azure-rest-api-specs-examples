package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/ServiceGatewayGetAddressLocationsResponse.json
func ExampleServiceGatewaysClient_NewGetAddressLocationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewServiceGatewaysClient().NewGetAddressLocationsPager("rg1", "sg", nil)
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
		// page.GetServiceGatewayAddressLocationsResult = armnetwork.GetServiceGatewayAddressLocationsResult{
		// 	Value: []*armnetwork.ServiceGatewayAddressLocationResponse{
		// 		{
		// 			AddressLocation: to.Ptr("192.0.0.1"),
		// 			Addresses: []*armnetwork.ServiceGatewayAddress{
		// 				{
		// 					Address: to.Ptr("10.0.0.4"),
		// 					Services: []*string{
		// 						to.Ptr("Service1")},
		// 				}},
		// 			},
		// 			{
		// 				AddressLocation: to.Ptr("192.0.0.2"),
		// 				Addresses: []*armnetwork.ServiceGatewayAddress{
		// 					{
		// 						Address: to.Ptr("10.0.0.5"),
		// 						Services: []*string{
		// 							to.Ptr("Service2")},
		// 					}},
		// 			}},
		// 		}
	}
}
