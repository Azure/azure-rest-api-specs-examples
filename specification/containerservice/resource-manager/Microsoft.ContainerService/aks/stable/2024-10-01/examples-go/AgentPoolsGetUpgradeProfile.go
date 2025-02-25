package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f2826dd1e08d4625b29867c01a232d2dad423521/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2024-10-01/examples/AgentPoolsGetUpgradeProfile.json
func ExampleAgentPoolsClient_GetUpgradeProfile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAgentPoolsClient().GetUpgradeProfile(ctx, "rg1", "clustername1", "agentpool1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AgentPoolUpgradeProfile = armcontainerservice.AgentPoolUpgradeProfile{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.ContainerService/managedClusters/agentPools/upgradeProfiles"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1/agentPools/agentpool1/upgradeprofiles/default"),
	// 	Properties: &armcontainerservice.AgentPoolUpgradeProfileProperties{
	// 		KubernetesVersion: to.Ptr("1.12.8"),
	// 		LatestNodeImageVersion: to.Ptr("AKSUbuntu:1604:2020.03.11"),
	// 		OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 		Upgrades: []*armcontainerservice.AgentPoolUpgradeProfilePropertiesUpgradesItem{
	// 			{
	// 				KubernetesVersion: to.Ptr("1.13.5"),
	// 		}},
	// 	},
	// }
}
