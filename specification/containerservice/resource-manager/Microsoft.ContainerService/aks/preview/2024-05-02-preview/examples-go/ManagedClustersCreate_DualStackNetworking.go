package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v6"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c7b98b36e4023331545051284d8500adf98f02fe/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/preview/2024-05-02-preview/examples/ManagedClustersCreate_DualStackNetworking.json
func ExampleManagedClustersClient_BeginCreateOrUpdate_createUpdateManagedClusterWithDualStackNetworking() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewManagedClustersClient().BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
		Location: to.Ptr("location1"),
		Tags: map[string]*string{
			"archv2": to.Ptr(""),
			"tier":   to.Ptr("production"),
		},
		Identity: &armcontainerservice.ManagedClusterIdentity{
			Type: to.Ptr(armcontainerservice.ResourceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgName1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": {},
			},
		},
		Properties: &armcontainerservice.ManagedClusterProperties{
			AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
			AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
				{
					Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
					AvailabilityZones: []*string{
						to.Ptr("1"),
						to.Ptr("2"),
						to.Ptr("3")},
					Count:              to.Ptr[int32](3),
					EnableNodePublicIP: to.Ptr(true),
					Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
					OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
					ScaleDownMode:      to.Ptr(armcontainerservice.ScaleDownModeDeallocate),
					VMSize:             to.Ptr("Standard_DS1_v2"),
					Name:               to.Ptr("nodepool1"),
				}},
			AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
				BalanceSimilarNodeGroups: to.Ptr("true"),
				Expander:                 to.Ptr(armcontainerservice.ExpanderPriority),
				MaxNodeProvisionTime:     to.Ptr("15m"),
				NewPodScaleUpDelay:       to.Ptr("1m"),
				ScaleDownDelayAfterAdd:   to.Ptr("15m"),
				ScanInterval:             to.Ptr("20s"),
				SkipNodesWithSystemPods:  to.Ptr("false"),
			},
			DiskEncryptionSetID:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
			DNSPrefix:               to.Ptr("dnsprefix1"),
			EnablePodSecurityPolicy: to.Ptr(true),
			EnableRBAC:              to.Ptr(true),
			KubernetesVersion:       to.Ptr(""),
			LinuxProfile: &armcontainerservice.LinuxProfile{
				AdminUsername: to.Ptr("azureuser"),
				SSH: &armcontainerservice.SSHConfiguration{
					PublicKeys: []*armcontainerservice.SSHPublicKey{
						{
							KeyData: to.Ptr("keydata"),
						}},
				},
			},
			NetworkProfile: &armcontainerservice.NetworkProfile{
				IPFamilies: []*armcontainerservice.IPFamily{
					to.Ptr(armcontainerservice.IPFamilyIPv4),
					to.Ptr(armcontainerservice.IPFamilyIPv6)},
				LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
					ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
						Count: to.Ptr[int32](2),
					},
				},
				LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
				OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
			},
			ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
				ClientID: to.Ptr("clientid"),
				Secret:   to.Ptr("secret"),
			},
			WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
				AdminPassword: to.Ptr("replacePassword1234$"),
				AdminUsername: to.Ptr("azureuser"),
			},
		},
		SKU: &armcontainerservice.ManagedClusterSKU{
			Name: to.Ptr(armcontainerservice.ManagedClusterSKUName("Basic")),
			Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
		},
	}, &armcontainerservice.ManagedClustersClientBeginCreateOrUpdateOptions{IfMatch: nil,
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
	// res.ManagedCluster = armcontainerservice.ManagedCluster{
	// 	Name: to.Ptr("clustername1"),
	// 	Type: to.Ptr("Microsoft.ContainerService/ManagedClusters"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ContainerService/managedClusters/clustername1"),
	// 	Location: to.Ptr("location1"),
	// 	Tags: map[string]*string{
	// 		"archv2": to.Ptr(""),
	// 		"tier": to.Ptr("production"),
	// 	},
	// 	Identity: &armcontainerservice.ManagedClusterIdentity{
	// 		Type: to.Ptr(armcontainerservice.ResourceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rgName1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1": &armcontainerservice.ManagedServiceIdentityUserAssignedIdentitiesValue{
	// 				ClientID: to.Ptr("clientId1"),
	// 				PrincipalID: to.Ptr("principalId1"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armcontainerservice.ManagedClusterProperties{
	// 		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	// 			{
	// 				Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	// 				AvailabilityZones: []*string{
	// 					to.Ptr("1"),
	// 					to.Ptr("2"),
	// 					to.Ptr("3")},
	// 					Count: to.Ptr[int32](3),
	// 					CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	// 					EnableNodePublicIP: to.Ptr(true),
	// 					MaxPods: to.Ptr[int32](110),
	// 					Mode: to.Ptr(armcontainerservice.AgentPoolModeSystem),
	// 					NodeImageVersion: to.Ptr("AKSUbuntu:1604:2020.03.11"),
	// 					OrchestratorVersion: to.Ptr("1.9.6"),
	// 					OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					ScaleDownMode: to.Ptr(armcontainerservice.ScaleDownModeDeallocate),
	// 					VMSize: to.Ptr("Standard_DS1_v2"),
	// 					Name: to.Ptr("nodepool1"),
	// 			}},
	// 			AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
	// 				BalanceSimilarNodeGroups: to.Ptr("true"),
	// 				Expander: to.Ptr(armcontainerservice.ExpanderPriority),
	// 				MaxNodeProvisionTime: to.Ptr("15m"),
	// 				NewPodScaleUpDelay: to.Ptr("1m"),
	// 				ScaleDownDelayAfterAdd: to.Ptr("15m"),
	// 				ScanInterval: to.Ptr("20s"),
	// 				SkipNodesWithSystemPods: to.Ptr("false"),
	// 			},
	// 			DiskEncryptionSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	// 			DNSPrefix: to.Ptr("dnsprefix1"),
	// 			EnablePodSecurityPolicy: to.Ptr(true),
	// 			EnableRBAC: to.Ptr(true),
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
	// 					to.Ptr(armcontainerservice.IPFamilyIPv4),
	// 					to.Ptr(armcontainerservice.IPFamilyIPv6)},
	// 					LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	// 						AllocatedOutboundPorts: to.Ptr[int32](2000),
	// 						EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	// 							{
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	// 							},
	// 							{
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	// 							},
	// 							{
	// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip3-ipv6"),
	// 						}},
	// 						IdleTimeoutInMinutes: to.Ptr[int32](10),
	// 						ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	// 							Count: to.Ptr[int32](2),
	// 							CountIPv6: to.Ptr[int32](1),
	// 						},
	// 					},
	// 					LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	// 					NetworkPlugin: to.Ptr(armcontainerservice.NetworkPluginKubenet),
	// 					OutboundType: to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	// 					PodCidr: to.Ptr("10.244.0.0/16"),
	// 					PodCidrs: []*string{
	// 						to.Ptr("10.244.0.0/16"),
	// 						to.Ptr("fd11:1234::/64")},
	// 						ServiceCidr: to.Ptr("10.0.0.0/16"),
	// 						ServiceCidrs: []*string{
	// 							to.Ptr("10.0.0.0/16"),
	// 							to.Ptr("fd00:1234::/108")},
	// 						},
	// 						NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	// 						ProvisioningState: to.Ptr("Succeeded"),
	// 						ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	// 							ClientID: to.Ptr("clientid"),
	// 						},
	// 						WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	// 							AdminUsername: to.Ptr("azureuser"),
	// 						},
	// 					},
	// 				}
}
