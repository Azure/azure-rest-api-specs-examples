package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/97f789aeb52adfc1e20c386005839f5276874d7d/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-05-01/examples/MachineList.json
func ExampleMachinesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMachinesClient().NewListPager("rg1", "clustername1", "agentpool1", nil)
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
		// page.MachineListResult = armcontainerservice.MachineListResult{
		// 	Value: []*armcontainerservice.Machine{
		// 		{
		// 			Name: to.Ptr("aks-nodepool1-25481572-vmss000000"),
		// 			Type: to.Ptr("Microsoft.ContainerService/managedClusters/agentPools/machines"),
		// 			ID: to.Ptr("/subscriptions/26fe00f8-9173-4872-9134-bb1d2e00343a/resourceGroups/dummyRG/providers/Microsoft.ContainerService/managedClusters/round/agentPools/nodepool1/machines/aks-nodepool1-25481572-vmss000000"),
		// 			Properties: &armcontainerservice.MachineProperties{
		// 				Network: &armcontainerservice.MachineNetworkProperties{
		// 					IPAddresses: []*armcontainerservice.MachineIPAddress{
		// 						{
		// 							Family: to.Ptr(armcontainerservice.IPFamilyIPv4),
		// 							IP: to.Ptr("172.20.2.4"),
		// 						},
		// 						{
		// 							Family: to.Ptr(armcontainerservice.IPFamilyIPv4),
		// 							IP: to.Ptr("10.0.0.1"),
		// 					}},
		// 				},
		// 				ResourceID: to.Ptr("/subscriptions/26fe00f8-9173-4872-9134-bb1d2e00343a/resourceGroups/dummyRG/providers/Microsoft.Compute/virtualMachineScaleSets/aks-nodepool1-25481572-vmss/virtualMachines/0"),
		// 			},
		// 			Zones: []*string{
		// 				to.Ptr("1")},
		// 		}},
		// 	}
	}
}
