package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/CloudServicePublicIpListAll.json
func ExamplePublicIPAddressesClient_NewListCloudServicePublicIPAddressesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPublicIPAddressesClient().NewListCloudServicePublicIPAddressesPager("cs-tester", "cs1", nil)
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
		// page.PublicIPAddressListResult = armnetwork.PublicIPAddressListResult{
		// 	Value: []*armnetwork.PublicIPAddress{
		// 		{
		// 			Name: to.Ptr("pub1"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/cs-tester/providers/Microsoft.Compute/cloudServices/cs1/roleInstances/Test_VM_0/networkInterfaces/nic1/ipConfigurations/ip1/publicIPAddresses/pub1"),
		// 			Properties: &armnetwork.PublicIPAddressPropertiesFormat{
		// 				DNSSettings: &armnetwork.PublicIPAddressDNSSettings{
		// 					DomainNameLabel: to.Ptr("vm1.testvmssacc"),
		// 					Fqdn: to.Ptr("vm1.testvmssacc.southeastasia.cloudapp.azure.com"),
		// 				},
		// 				IdleTimeoutInMinutes: to.Ptr[int32](10),
		// 				IPAddress: to.Ptr("13.67.119.72"),
		// 				IPConfiguration: &armnetwork.IPConfiguration{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/cs-tester/providers/Microsoft.Compute/cloudServices/cs1/roleInstances/Test_VM_0/networkInterfaces/nic1/ipConfigurations/ip1"),
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 				PublicIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("pub1"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/cs-tester/providers/Microsoft.Compute/cloudServices/cs1/roleInstances/Test_VM_3/networkInterfaces/nic1/ipConfigurations/ip1/publicIPAddresses/pub1"),
		// 			Properties: &armnetwork.PublicIPAddressPropertiesFormat{
		// 				DNSSettings: &armnetwork.PublicIPAddressDNSSettings{
		// 					DomainNameLabel: to.Ptr("vm3.testvmssacc"),
		// 					Fqdn: to.Ptr("vm3.testvmssacc.southeastasia.cloudapp.azure.com"),
		// 				},
		// 				IdleTimeoutInMinutes: to.Ptr[int32](10),
		// 				IPAddress: to.Ptr("13.67.118.216"),
		// 				IPConfiguration: &armnetwork.IPConfiguration{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/cs-tester/providers/Microsoft.Compute/cloudServices/cs1/roleInstances/Test_VM_3/networkInterfaces/nic1/ipConfigurations/ip1"),
		// 				},
		// 				ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
		// 				PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
		// 				PublicIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
		// 			},
		// 	}},
		// }
	}
}
