package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-04-02-preview/examples/AgentPoolsCreate_TypeVirtualMachines_Autoscale.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithVirtualMachinesPoolTypeWithAutoscalingEnabled() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAgentPoolsClient().BeginCreateOrUpdate(ctx, "rg1", "clustername1", "agentpool1", armcontainerservice.AgentPool{
		Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
			Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachines),
			NodeLabels: map[string]*string{
				"key1": to.Ptr("val1"),
			},
			NodeTaints: []*string{
				to.Ptr("Key1=Value1:NoSchedule")},
			OrchestratorVersion: to.Ptr("1.29.0"),
			OSType:              to.Ptr(armcontainerservice.OSTypeLinux),
			Tags: map[string]*string{
				"name1": to.Ptr("val1"),
			},
			VirtualMachinesProfile: &armcontainerservice.VirtualMachinesProfile{
				Scale: &armcontainerservice.ScaleProfile{
					Autoscale: []*armcontainerservice.AutoScaleProfile{
						{
							MaxCount: to.Ptr[int32](5),
							MinCount: to.Ptr[int32](1),
							Sizes: []*string{
								to.Ptr("Standard_D2_v2"),
								to.Ptr("Standard_D2_v3")},
						}},
				},
			},
		},
	}, &armcontainerservice.AgentPoolsClientBeginCreateOrUpdateOptions{IfMatch: nil,
		IfNoneMatch: nil,
	})
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
	// res.AgentPool = armcontainerservice.AgentPool{
	// 	Name: to.Ptr("agentpool1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/agentPools"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentPools/agentpool1"),
	// 	Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
	// 		Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachines),
	// 		CurrentOrchestratorVersion: to.Ptr("1.29.0"),
	// 		MaxPods: to.Ptr[int32](110),
	// 		NodeImageVersion: to.Ptr("AKSUbuntu-1804gen2containerd-2021.09.11"),
	// 		NodeLabels: map[string]*string{
	// 			"key1": to.Ptr("val1"),
	// 		},
	// 		NodeTaints: []*string{
	// 			to.Ptr("Key1=Value1:NoSchedule")},
	// 			OrchestratorVersion: to.Ptr("1.29.0"),
	// 			OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			Tags: map[string]*string{
	// 				"name1": to.Ptr("val1"),
	// 			},
	// 			VirtualMachineNodesStatus: []*armcontainerservice.VirtualMachineNodes{
	// 				{
	// 					Count: to.Ptr[int32](1),
	// 					Size: to.Ptr("Standard_D2_v2"),
	// 				},
	// 				{
	// 					Count: to.Ptr[int32](2),
	// 					Size: to.Ptr("Standard_D2_v3"),
	// 			}},
	// 			VirtualMachinesProfile: &armcontainerservice.VirtualMachinesProfile{
	// 				Scale: &armcontainerservice.ScaleProfile{
	// 					Autoscale: []*armcontainerservice.AutoScaleProfile{
	// 						{
	// 							MaxCount: to.Ptr[int32](5),
	// 							MinCount: to.Ptr[int32](1),
	// 							Sizes: []*string{
	// 								to.Ptr("Standard_D2_v2"),
	// 								to.Ptr("Standard_D2_v3")},
	// 						}},
	// 					},
	// 				},
	// 			},
	// 		}
}
