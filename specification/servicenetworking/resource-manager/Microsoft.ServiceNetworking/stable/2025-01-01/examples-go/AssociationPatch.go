package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: 2025-01-01/AssociationPatch.json
func ExampleAssociationsInterfaceClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssociationsInterfaceClient().Update(ctx, "rg1", "tc1", "as1", armservicenetworking.AssociationUpdate{
		Properties: &armservicenetworking.AssociationUpdateProperties{
			AssociationType: to.Ptr(armservicenetworking.AssociationTypeSubnets),
			Subnet: &armservicenetworking.AssociationSubnetUpdate{
				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet-tc/subnets/tc-subnet"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armservicenetworking.AssociationsInterfaceClientUpdateResponse{
	// 	Association: &armservicenetworking.Association{
	// 		Name: to.Ptr("as1"),
	// 		Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/associations"),
	// 		ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/tc1/associations/as1"),
	// 		Location: to.Ptr("NorthCentralUS"),
	// 		Properties: &armservicenetworking.AssociationProperties{
	// 			AssociationType: to.Ptr(armservicenetworking.AssociationTypeSubnets),
	// 			Subnet: &armservicenetworking.AssociationSubnet{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet-tc/subnets/tc-subnet"),
	// 			},
	// 		},
	// 	},
	// }
}
