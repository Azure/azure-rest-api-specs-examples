package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v9"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/72410da64f6e945db1e1f1af220e077ba5bdb857/specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/CloudServicePublicIpGet.json
func ExamplePublicIPAddressesClient_GetCloudServicePublicIPAddress() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPublicIPAddressesClient().GetCloudServicePublicIPAddress(ctx, "cs-tester", "cs1", "Test_VM_0", "nic1", "ip1", "pub1", &armnetwork.PublicIPAddressesClientGetCloudServicePublicIPAddressOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PublicIPAddress = armnetwork.PublicIPAddress{
	// 	Name: to.Ptr("pub1"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/cs-tester/providers/Microsoft.Compute/cloudServices/cs1/roleInstances/Test_VM_0/networkInterfaces/nic1/ipConfigurations/ip1/publicIPAddresses/pub1"),
	// 	Properties: &armnetwork.PublicIPAddressPropertiesFormat{
	// 		DNSSettings: &armnetwork.PublicIPAddressDNSSettings{
	// 			DomainNameLabel: to.Ptr("vm1.testvmssacc"),
	// 			Fqdn: to.Ptr("vm1.testvmssacc.southeastasia.cloudapp.azure.com"),
	// 		},
	// 		IdleTimeoutInMinutes: to.Ptr[int32](10),
	// 		IPAddress: to.Ptr("13.67.119.72"),
	// 		IPConfiguration: &armnetwork.IPConfiguration{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/cs-tester/providers/Microsoft.Compute/cloudServices/cs1/roleInstances/Test_VM_0/networkInterfaces/nic1/ipConfigurations/ip1"),
	// 		},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		PublicIPAddressVersion: to.Ptr(armnetwork.IPVersionIPv4),
	// 		PublicIPAllocationMethod: to.Ptr(armnetwork.IPAllocationMethodDynamic),
	// 	},
	// }
}
