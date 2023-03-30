package armdelegatednetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/delegatednetwork/armdelegatednetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/delegatedSubnetListByRG.json
func ExampleDelegatedSubnetServiceClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdelegatednetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDelegatedSubnetServiceClient().NewListByResourceGroupPager("testRG", nil)
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
		// page.DelegatedSubnets = armdelegatednetwork.DelegatedSubnets{
		// 	Value: []*armdelegatednetwork.DelegatedSubnet{
		// 		{
		// 			Name: to.Ptr("delegated1"),
		// 			Type: to.Ptr("Microsoft.DelegatedNetwork/delegatedSubnets"),
		// 			ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.DelegatedNetwork/delegatedSubnets/delegated1"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armdelegatednetwork.DelegatedSubnetProperties{
		// 				ControllerDetails: &armdelegatednetwork.ControllerDetails{
		// 					ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.DelegatedNetwork/controller/dnctestcontroller"),
		// 				},
		// 				ProvisioningState: to.Ptr(armdelegatednetwork.DelegatedSubnetStateSucceeded),
		// 				ResourceGUID: to.Ptr("5a82cbcf-e8ea-4175-ac2b-ad36a73f9801"),
		// 				SubnetDetails: &armdelegatednetwork.SubnetDetails{
		// 					ID: to.Ptr("/subscriptions/613192d7-503f-477a-9cfe-4efc3ee2bd60/resourceGroups/TestRG/providers/Microsoft.Network/virtualNetworks/testvnet/subnets/testsubnet"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
