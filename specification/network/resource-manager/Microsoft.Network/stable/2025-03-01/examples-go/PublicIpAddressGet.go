package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/PublicIpAddressGet.json
func ExamplePublicIPAddressesClient_Get_getPublicIpAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPublicIPAddressesClient().Get(ctx, "rg1", "testDNS-ip", &armnetwork.PublicIPAddressesClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PublicIPAddress = armnetwork.PublicIPAddress{
	// 	Name: to.Ptr("testDNS-ip"),
	// 	Type: to.Ptr("Microsoft.Network/publicIPAddresses"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/testDNS-ip"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armnetwork.PublicIPAddressPropertiesFormat{
	// 		DdosSettings: &armnetwork.DdosSettings{
	// 			DdosProtectionPlan: &armnetwork.SubResource{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/ddosProtectionPlans/test-plan"),
	// 			},
	// 			ProtectionMode: to.Ptr(armnetwork.DdosSettingsProtectionModeEnabled),
	// 		},
	// 		IdleTimeoutInMinutes: to.Ptr[int32](4),
	// 		IPConfiguration: &armnetwork.IPConfiguration{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testDNS649/ipConfigurations/ipconfig1"),
	// 		},
	// 		IPTags: []*armnetwork.IPTag{
	// 			{
	// 				IPTagType: to.Ptr("FirstPartyUsage"),
	// 				Tag: to.Ptr("SQL"),
	// 			},
	// 			{
	// 				IPTagType: to.Ptr("FirstPartyUsage"),
	// 				Tag: to.Ptr("Storage"),
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 		PublicIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 	},
	// }
}
