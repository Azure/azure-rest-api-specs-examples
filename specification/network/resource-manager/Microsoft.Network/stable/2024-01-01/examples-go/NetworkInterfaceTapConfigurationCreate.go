package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/network/resource-manager/Microsoft.Network/stable/2024-01-01/examples/NetworkInterfaceTapConfigurationCreate.json
func ExampleInterfaceTapConfigurationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewInterfaceTapConfigurationsClient().BeginCreateOrUpdate(ctx, "testrg", "mynic", "tapconfiguration1", armnetwork.InterfaceTapConfiguration{
		Properties: &armnetwork.InterfaceTapConfigurationPropertiesFormat{
			VirtualNetworkTap: &armnetwork.VirtualNetworkTap{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworkTaps/testvtap"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.InterfaceTapConfiguration = armnetwork.InterfaceTapConfiguration{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/networkInterfaces/mynic/tapConfigurations/tapConfiguration1"),
	// 	Name: to.Ptr("tapConfiguration1"),
	// 	Type: to.Ptr("Microsoft.Network/networkInterfaces/tapConfigurations"),
	// 	Etag: to.Ptr("etag"),
	// 	Properties: &armnetwork.InterfaceTapConfigurationPropertiesFormat{
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		VirtualNetworkTap: &armnetwork.VirtualNetworkTap{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/testrg/providers/Microsoft.Network/virtualNetworkTaps/testvtap"),
	// 		},
	// 	},
	// }
}
