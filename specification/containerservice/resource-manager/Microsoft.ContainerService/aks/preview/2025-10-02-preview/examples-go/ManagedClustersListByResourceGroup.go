package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/455d20a5e76d8184f7cff960501a57e1f88986b7/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2025-10-02-preview/examples/ManagedClustersListByResourceGroup.json
func ExampleManagedClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewManagedClustersClient().NewListByResourceGroupPager("rg1", nil)
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
		// page.ManagedClusterListResult = armcontainerservice.ManagedClusterListResult{
		// 	Value: []*armcontainerservice.ManagedCluster{
		// 		{
		// 			Name: to.Ptr("clustername1"),
		// 			Type: to.Ptr("Microsoft.ContainerService/ManagedClusters"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
		// 			Location: to.Ptr("location1"),
		// 			Tags: map[string]*string{
		// 				"archv2": to.Ptr(""),
		// 				"tier": to.Ptr("production"),
		// 			},
		// 			Properties: &armcontainerservice.ManagedClusterProperties{
		// 				AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
		// 					{
		// 						Count: to.Ptr[int32](3),
		// 						CurrentOrchestratorVersion: to.Ptr("1.9.6"),
		// 						MaxPods: to.Ptr[int32](110),
		// 						OrchestratorVersion: to.Ptr("1.9.6"),
		// 						OSType: to.Ptr(armcontainerservice.OSTypeLinux),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 						VMSize: to.Ptr("Standard_DS1_v2"),
		// 						Name: to.Ptr("nodepool1"),
		// 				}},
		// 				DiskEncryptionSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
		// 				DNSPrefix: to.Ptr("dnsprefix1"),
		// 				EnableRBAC: to.Ptr(false),
		// 				Fqdn: to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
		// 				KubernetesVersion: to.Ptr("1.9.6"),
		// 				LinuxProfile: &armcontainerservice.LinuxProfile{
		// 					AdminUsername: to.Ptr("azureuser"),
		// 					SSH: &armcontainerservice.SSHConfiguration{
		// 						PublicKeys: []*armcontainerservice.SSHPublicKey{
		// 							{
		// 								KeyData: to.Ptr("keydata"),
		// 						}},
		// 					},
		// 				},
		// 				MaxAgentPools: to.Ptr[int32](1),
		// 				NetworkProfile: &armcontainerservice.NetworkProfile{
		// 					DNSServiceIP: to.Ptr("10.0.0.10"),
		// 					NetworkPlugin: to.Ptr(armcontainerservice.NetworkPluginKubenet),
		// 					PodCidr: to.Ptr("10.244.0.0/16"),
		// 					ServiceCidr: to.Ptr("10.0.0.0/16"),
		// 				},
		// 				NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
		// 					ClientID: to.Ptr("clientid"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}
