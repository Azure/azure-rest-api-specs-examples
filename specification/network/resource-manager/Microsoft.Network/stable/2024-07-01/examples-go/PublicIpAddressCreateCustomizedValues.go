package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c58fa033619b12c7cfa8a0ec5a9bf03bb18869ab/specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/PublicIpAddressCreateCustomizedValues.json
func ExamplePublicIPAddressesClient_BeginCreateOrUpdate_createPublicIpAddressAllocationMethod() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPublicIPAddressesClient().BeginCreateOrUpdate(ctx, "rg1", "test-ip", armnetwork.PublicIPAddress{
		Location: to.Ptr("eastus"),
		Properties: &armnetwork.PublicIPAddressPropertiesFormat{
			IdleTimeoutInMinutes:     to.Ptr[int32](10),
			PublicIPAddressVersion:   to.Ptr(armnetwork.IPVersionIPv4),
			PublicIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
		},
		SKU: &armnetwork.PublicIPAddressSKU{
			Name: to.Ptr(armnetwork.PublicIPAddressSKUNameStandard),
			Tier: to.Ptr(armnetwork.PublicIPAddressSKUTierGlobal),
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
	// res.PublicIPAddress = armnetwork.PublicIPAddress{
	// 	Name: to.Ptr("testDNS-ip"),
	// 	Type: to.Ptr("Microsoft.Network/publicIPAddresses"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/test-ip"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armnetwork.PublicIPAddressPropertiesFormat{
	// 		DdosSettings: &armnetwork.DdosSettings{
	// 			ProtectionMode: to.Ptr(armnetwork.DdosSettingsProtectionModeVirtualNetworkInherited),
	// 		},
	// 		IdleTimeoutInMinutes: to.Ptr[int32](10),
	// 		IPConfiguration: &armnetwork.IPConfiguration{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testDNS649/ipConfigurations/ipconfig1"),
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 		PublicIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodStatic),
	// 	},
	// 	SKU: &armnetwork.PublicIPAddressSKU{
	// 		Name: to.Ptr(armnetwork.PublicIPAddressSKUNameStandard),
	// 		Tier: to.Ptr(armnetwork.PublicIPAddressSKUTierGlobal),
	// 	},
	// 	Zones: []*string{
	// 		to.Ptr("1")},
	// 	}
}
