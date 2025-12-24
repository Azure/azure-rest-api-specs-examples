package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b24c97bfc136b01dd46a1c8ddcecd0bb5a1ab152/specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/SubnetGetWithDelegation.json
func ExampleSubnetsClient_Get_getSubnetWithADelegation() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubnetsClient().Get(ctx, "subnet-test", "vnetname", "subnet1", &armnetwork.SubnetsClientGetOptions{Expand: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Subnet = armnetwork.Subnet{
	// 	ID: to.Ptr("/subscriptions/subId/resourceGroups/subnet-test/providers/Microsoft.Network/virtualNetworks/vnetname/subnets/subnet1"),
	// 	Name: to.Ptr("subnet1"),
	// 	Properties: &armnetwork.SubnetPropertiesFormat{
	// 		AddressPrefix: to.Ptr("10.0.0.0/16"),
	// 		Delegations: []*armnetwork.Delegation{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/subnet1/delegations/myDelegation"),
	// 				Name: to.Ptr("myDelegation"),
	// 				Properties: &armnetwork.ServiceDelegationPropertiesFormat{
	// 					Actions: []*string{
	// 					},
	// 					ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 					ServiceName: to.Ptr("Microsoft.Provider/resourceType"),
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 		Purpose: to.Ptr(""),
	// 	},
	// }
}
