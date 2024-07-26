package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/81a4ee5a83ae38620c0e1404793caffe005d26e4/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-04-02-preview/examples/AgentPoolsCreate_EnableEncryptionAtHost.json
func ExampleAgentPoolsClient_BeginCreateOrUpdate_createAgentPoolWithEncryptionAtHostEnabled() {
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
			Count:                  to.Ptr[int32](3),
			EnableEncryptionAtHost: to.Ptr(true),
			OrchestratorVersion:    to.Ptr(""),
			OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
			VMSize:                 to.Ptr("Standard_DS2_v2"),
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
	// 		Count: to.Ptr[int32](3),
	// 		CurrentOrchestratorVersion: to.Ptr("1.17.13"),
	// 		EnableEncryptionAtHost: to.Ptr(true),
	// 		MaxPods: to.Ptr[int32](110),
	// 		OrchestratorVersion: to.Ptr("1.17.13"),
	// 		OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		VMSize: to.Ptr("Standard_DS2_v2"),
	// 	},
	// }
}
