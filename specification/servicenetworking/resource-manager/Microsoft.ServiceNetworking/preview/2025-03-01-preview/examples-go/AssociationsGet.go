package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: 2025-03-01-preview/AssociationsGet.json
func ExampleAssociationsInterfaceClient_NewListByTrafficControllerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("subid", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssociationsInterfaceClient().NewListByTrafficControllerPager("rg1", "tc1", nil)
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
		// page = armservicenetworking.AssociationsInterfaceClientListByTrafficControllerResponse{
		// 	AssociationListResult: armservicenetworking.AssociationListResult{
		// 		Value: []*armservicenetworking.Association{
		// 			{
		// 				Name: to.Ptr("as1"),
		// 				Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/associations"),
		// 				ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ServiceNetworking/trafficControllers/tc1/associations/as1"),
		// 				Location: to.Ptr("NorthCentralUS"),
		// 				Properties: &armservicenetworking.AssociationProperties{
		// 					AssociationType: to.Ptr(armservicenetworking.AssociationTypeSubnets),
		// 					Subnet: &armservicenetworking.AssociationSubnet{
		// 						ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet-tc/subnets/tc-subnet"),
		// 					},
		// 					ProvisioningState: to.Ptr(armservicenetworking.ProvisioningStateSucceeded),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
