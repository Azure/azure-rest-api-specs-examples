package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9cc7633f842575274f715cc02e37c5769ac2742d/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/ManagedClustersGet.json
func ExampleManagedClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewManagedClustersClient().Get(ctx, "rg1", "clustername1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
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
	// 		"archv2": to.Ptr(""),
	// 		"tier": to.Ptr("production"),
	// 	},
	// 	ETag: to.Ptr("beywbwei"),
	// 	Properties: &armcontainerservice.ManagedClusterProperties{
	// 		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	// 			{
	// 				AvailabilityZones: []*string{
	// 					to.Ptr("1"),
	// 					to.Ptr("2"),
	// 					to.Ptr("3")},
	// 					Count: to.Ptr[int32](3),
	// 					CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	// 					ETag: to.Ptr("nvewbvoi"),
	// 					MaxPods: to.Ptr[int32](110),
	// 					NodeImageVersion: to.Ptr("AKSUbuntu:1604:2020.03.11"),
	// 					OrchestratorVersion: to.Ptr("1.9.6"),
	// 					OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					UpgradeSettings: &armcontainerservice.AgentPoolUpgradeSettings{
	// 						MaxSurge: to.Ptr("33%"),
	// 					},
	// 					VMSize: to.Ptr("Standard_DS1_v2"),
	// 					Name: to.Ptr("nodepool1"),
	// 			}},
	// 			AzurePortalFQDN: to.Ptr("dnsprefix1-abcd1234.portal.hcp.eastus.azmk8s.io"),
	// 			CurrentKubernetesVersion: to.Ptr("1.9.6"),
	// 			DiskEncryptionSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	// 			DNSPrefix: to.Ptr("dnsprefix1"),
	// 			EnableRBAC: to.Ptr(false),
	// 			Fqdn: to.Ptr("dnsprefix1-abcd1234.hcp.eastus.azmk8s.io"),
	// 			KubernetesVersion: to.Ptr("1.9.6"),
	// 			LinuxProfile: &armcontainerservice.LinuxProfile{
	// 				AdminUsername: to.Ptr("azureuser"),
	// 				SSH: &armcontainerservice.SSHConfiguration{
	// 					PublicKeys: []*armcontainerservice.SSHPublicKey{
	// 						{
	// 							KeyData: to.Ptr("keydata"),
	// 					}},
	// 				},
	// 			},
	// 			MaxAgentPools: to.Ptr[int32](1),
	// 			NetworkProfile: &armcontainerservice.NetworkProfile{
	// 				DNSServiceIP: to.Ptr("10.0.0.10"),
	// 				IPFamilies: []*armcontainerservice.IPFamily{
	// 					to.Ptr(armcontainerservice.IPFamilyIPv4)},
	// 					LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	// 						AllocatedOutboundPorts: to.Ptr[int32](2000),
	// 						EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	// 							{
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	// 							},
	// 							{
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	// 						}},
	// 						IdleTimeoutInMinutes: to.Ptr[int32](10),
	// 						OutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileOutboundIPs{
	// 							PublicIPs: []*armcontainerservice.ResourceReference{
	// 								{
	// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.Network/publicIPAddresses/customeroutboundip1"),
	// 								},
	// 								{
	// 									ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.Network/publicIPAddresses/customeroutboundip2"),
	// 							}},
	// 						},
	// 					},
	// 					LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
	// 					NetworkPlugin: to.Ptr(armcontainerservice.NetworkPluginKubenet),
	// 					OutboundType: to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	// 					PodCidr: to.Ptr("10.244.0.0/16"),
	// 					PodCidrs: []*string{
	// 						to.Ptr("10.244.0.0/16")},
	// 						ServiceCidr: to.Ptr("10.0.0.0/16"),
	// 						ServiceCidrs: []*string{
	// 							to.Ptr("10.0.0.0/16")},
	// 						},
	// 						NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 						ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	// 							ClientID: to.Ptr("clientid"),
	// 						},
	// 						UpgradeSettings: &armcontainerservice.ClusterUpgradeSettings{
	// 							OverrideSettings: &armcontainerservice.UpgradeOverrideSettings{
	// 								ForceUpgrade: to.Ptr(true),
	// 								Until: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-11-01T13:00:00.000Z"); return t}()),
	// 							},
	// 						},
	// 					},
	// 				}
}
