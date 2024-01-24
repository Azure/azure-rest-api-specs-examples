package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/8b5618d760532b69d6f7434f57172ea52e109c79/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-11-01/examples/AgentPoolsCreate_Spot.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createSpotAgentPool() {
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
			Count: to.Ptr[int32](3),
			NodeLabels: map[string]*string{
				"key1": to.Ptr("val1"),
			},
			NodeTaints: []*string{
				to.Ptr("Key1=Value1:NoSchedule")},
			OrchestratorVersion:    to.Ptr(""),
			OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
			ScaleSetEvictionPolicy: to.Ptr(armcontainerservice.ScaleSetEvictionPolicyDelete),
			ScaleSetPriority:       to.Ptr(armcontainerservice.ScaleSetPrioritySpot),
			Tags: map[string]*string{
				"name1": to.Ptr("val1"),
			},
			VMSize: to.Ptr("Standard_DS1_v2"),
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
	// 		MaxPods: to.Ptr[int32](110),
	// 		NodeLabels: map[string]*string{
	// 			"key1": to.Ptr("val1"),
	// 		},
	// 		NodeTaints: []*string{
	// 			to.Ptr("Key1=Value1:NoSchedule")},
	// 			OrchestratorVersion: to.Ptr("1.9.6"),
	// 			OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ScaleSetEvictionPolicy: to.Ptr(armcontainerservice.ScaleSetEvictionPolicyDelete),
	// 			ScaleSetPriority: to.Ptr(armcontainerservice.ScaleSetPrioritySpot),
	// 			SpotMaxPrice: to.Ptr[float32](-1),
	// 			Tags: map[string]*string{
	// 				"name1": to.Ptr("val1"),
	// 			},
	// 			VMSize: to.Ptr("Standard_DS1_v2"),
	// 		},
	// 	}
}
