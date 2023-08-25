package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ad6484fee29013ab0b3619217970b106a1f44cf0/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2023-06-01/examples/AgentPoolsGet.json
func ExampleAgentPoolsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAgentPoolsClient().Get(ctx, "rg1", "clustername1", "agentpool1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgentPool = armcontainerservice.AgentPool{
	// 	Name: to.Ptr("agentpool1"),
	// 	ID: to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentPools/agentpool1"),
	// 	Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
	// 		Count: to.Ptr[int32](3),
	// 		CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	// 		MaxPods: to.Ptr[int32](110),
	// 		NodeImageVersion: to.Ptr("AKSUbuntu:1604:2020.03.11"),
	// 		OrchestratorVersion: to.Ptr("1.9.6"),
	// 		OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		UpgradeSettings: &armcontainerservice.AgentPoolUpgradeSettings{
	// 			MaxSurge: to.Ptr("33%"),
	// 		},
	// 		VMSize: to.Ptr("Standard_DS1_v2"),
	// 	},
	// }
}
