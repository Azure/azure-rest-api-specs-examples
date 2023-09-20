package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80c21c17b4a7aa57f637ee594f7cfd653255a7e0/specification/network/resource-manager/Microsoft.Network/stable/2023-05-01/examples/ExpressRoutePortsLocationGet.json
func ExampleExpressRoutePortsLocationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewExpressRoutePortsLocationsClient().Get(ctx, "locationName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ExpressRoutePortsLocation = armnetwork.ExpressRoutePortsLocation{
	// 	Name: to.Ptr("locationName"),
	// 	Type: to.Ptr("Microsoft.Network/expressRoutePortsLocations"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Network/expressRoutePortsLocations/locationName"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.ExpressRoutePortsLocationPropertiesFormat{
	// 		Address: to.Ptr("123 Main Street, City, State, Zip"),
	// 		AvailableBandwidths: []*armnetwork.ExpressRoutePortsLocationBandwidths{
	// 			{
	// 				OfferName: to.Ptr("100 Gbps"),
	// 				ValueInGbps: to.Ptr[int32](100),
	// 		}},
	// 		Contact: to.Ptr("email@address.com"),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}
