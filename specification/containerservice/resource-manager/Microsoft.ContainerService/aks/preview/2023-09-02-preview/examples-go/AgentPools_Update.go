package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2023-09-02-preview/examples/AgentPools_Update.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_updateAgentPool() {
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
			Count:             to.Ptr[int32](3),
			EnableAutoScaling: to.Ptr(true),
			MaxCount:          to.Ptr[int32](2),
			MinCount:          to.Ptr[int32](2),
			NodeTaints: []*string{
				to.Ptr("Key1=Value1:NoSchedule")},
			OrchestratorVersion:    to.Ptr(""),
			OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
			ScaleSetEvictionPolicy: to.Ptr(armcontainerservice.ScaleSetEvictionPolicyDelete),
			ScaleSetPriority:       to.Ptr(armcontainerservice.ScaleSetPrioritySpot),
			VMSize:                 to.Ptr("Standard_DS1_v2"),
		},
	}, nil)
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
	// 		Count: to.Ptr[int32](3),
	// 		CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	// 		EnableAutoScaling: to.Ptr(true),
	// 		MaxCount: to.Ptr[int32](2),
	// 		MaxPods: to.Ptr[int32](110),
	// 		MinCount: to.Ptr[int32](2),
	// 		NodeTaints: []*string{
	// 			to.Ptr("Key1=Value1:NoSchedule")},
	// 			OrchestratorVersion: to.Ptr("1.9.6"),
	// 			OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ScaleSetEvictionPolicy: to.Ptr(armcontainerservice.ScaleSetEvictionPolicyDelete),
	// 			ScaleSetPriority: to.Ptr(armcontainerservice.ScaleSetPrioritySpot),
	// 			VMSize: to.Ptr("Standard_DS1_v2"),
	// 		},
	// 	}
}
