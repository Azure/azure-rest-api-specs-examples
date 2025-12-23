package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9cc7633f842575274f715cc02e37c5769ac2742d/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/ManagedClustersUpdateTags.json
func ExampleManagedClustersClient_BeginUpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedClustersClient().BeginUpdateTags(ctx, "rg1", "clustername1", armcontainerservice.TagsObject{
		Tags: map[string]*string{
			"archv3": to.Ptr(""),
			"tier":   to.Ptr("testing"),
		},
	}, &armcontainerservice.ManagedClustersClientBeginUpdateTagsOptions{IfMatch: nil})
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
	// res.ManagedCluster = armcontainerservice.ManagedCluster{
	// 	Name: to.Ptr("clustername1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	// 	Location: to.Ptr("location1"),
	// 	Tags: map[string]*string{
	// 		"archv3": to.Ptr(""),
	// 		"tier": to.Ptr("testing"),
	// 	},
	// 	Properties: &armcontainerservice.ManagedClusterProperties{
	// 		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	// 			{
	// 				Count: to.Ptr[int32](3),
	// 				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	// 				MaxPods: to.Ptr[int32](110),
	// 				OrchestratorVersion: to.Ptr("1.9.6"),
	// 				OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				VMSize: to.Ptr("Standard_DS1_v2"),
	// 				Name: to.Ptr("nodepool1"),
	// 		}},
	// 		DiskEncryptionSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	// 		DNSPrefix: to.Ptr("dnsprefix1"),
	// 		EnableRBAC: to.Ptr(false),
	// 		Fqdn: to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	// 		KubernetesVersion: to.Ptr("1.9.6"),
	// 		LinuxProfile: &armcontainerservice.LinuxProfile{
	// 			AdminUsername: to.Ptr("azureuser"),
	// 			SSH: &armcontainerservice.SSHConfiguration{
	// 				PublicKeys: []*armcontainerservice.SSHPublicKey{
	// 					{
	// 						KeyData: to.Ptr("keydata"),
	// 				}},
	// 			},
	// 		},
	// 		NetworkProfile: &armcontainerservice.NetworkProfile{
	// 			DNSServiceIP: to.Ptr("10.0.0.10"),
	// 			NetworkPlugin: to.Ptr(armcontainerservice.NetworkPluginKubenet),
	// 			PodCidr: to.Ptr("10.244.0.0/16"),
	// 			ServiceCidr: to.Ptr("10.0.0.0/16"),
	// 		},
	// 		NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	// 			ClientID: to.Ptr("clientid"),
	// 		},
	// 	},
	// }
}
