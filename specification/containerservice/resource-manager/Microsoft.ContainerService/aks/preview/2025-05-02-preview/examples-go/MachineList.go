package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/97f789aeb52adfc1e20c386005839f5276874d7d/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-05-02-preview/examples/MachineList.json
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
		// 				Hardware: &armcontainerservice.MachineHardwareProfile{
		// 					VMSize: to.Ptr("Standard_DS1_v2"),
		// 				},
		// 				Kubernetes: &armcontainerservice.MachineKubernetesProfile{
		// 					CurrentOrchestratorVersion: to.Ptr("1.30.6"),
		// 					KubeletDiskType: to.Ptr(armcontainerservice.KubeletDiskTypeOS),
		// 					MaxPods: to.Ptr[int32](110),
		// 					NodeLabels: map[string]*string{
		// 						"key1": to.Ptr("val1"),
		// 					},
		// 					NodeTaints: []*string{
		// 						to.Ptr("Key1=Value1:NoSchedule")},
		// 						OrchestratorVersion: to.Ptr("1.30"),
		// 					},
		// 					Mode: to.Ptr(armcontainerservice.AgentPoolModeUser),
		// 					Network: &armcontainerservice.MachineNetworkProperties{
		// 						IPAddresses: []*armcontainerservice.MachineIPAddress{
		// 							{
		// 								Family: to.Ptr(armcontainerservice.IPFamilyIPv4),
		// 								IP: to.Ptr("172.20.2.4"),
		// 							},
		// 							{
		// 								Family: to.Ptr(armcontainerservice.IPFamilyIPv4),
		// 								IP: to.Ptr("10.0.0.1"),
		// 						}},
		// 					},
		// 					NodeImageVersion: to.Ptr("AKSUbuntu:1604:2023.03.11"),
		// 					Priority: to.Ptr(armcontainerservice.ScaleSetPrioritySpot),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ResourceID: to.Ptr("/subscriptions/26fe00f8-9173-4872-9134-bb1d2e00343a/resourceGroups/dummyRG/providers/Microsoft.Compute/virtualMachineScaleSets/aks-nodepool1-25481572-vmss/virtualMachines/0"),
		// 				},
		// 				Zones: []*string{
		// 					to.Ptr("1")},
		// 			}},
		// 		}
	}
}
