package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9cc7633f842575274f715cc02e37c5769ac2742d/specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/ManagedClustersCreate_SecurityProfile.json
func ExampleManagedClustersClient_BeginCreateOrUpdate_createManagedClusterWithSecurityProfileConfigured() {
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
		Properties: &armcontainerservice.ManagedClusterProperties{
			AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
				{
					Type:               to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
					Count:              to.Ptr[int32](3),
					EnableNodePublicIP: to.Ptr(true),
					Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
					OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
					VMSize:             to.Ptr("Standard_DS2_v2"),
					Name:               to.Ptr("nodepool1"),
				}},
			DNSPrefix:         to.Ptr("dnsprefix1"),
			KubernetesVersion: to.Ptr(""),
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
				LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
					ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
						Count: to.Ptr[int32](2),
					},
				},
				LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUStandard),
				OutboundType:    to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
			},
			SecurityProfile: &armcontainerservice.ManagedClusterSecurityProfile{
				Defender: &armcontainerservice.ManagedClusterSecurityProfileDefender{
					LogAnalyticsWorkspaceResourceID: to.Ptr("/subscriptions/SUB_ID/resourcegroups/RG_NAME/providers/microsoft.operationalinsights/workspaces/WORKSPACE_NAME"),
					SecurityMonitoring: &armcontainerservice.ManagedClusterSecurityProfileDefenderSecurityMonitoring{
						Enabled: to.Ptr(true),
					},
				},
				WorkloadIdentity: &armcontainerservice.ManagedClusterSecurityProfileWorkloadIdentity{
					Enabled: to.Ptr(true),
				},
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
	// 	Properties: &armcontainerservice.ManagedClusterProperties{
	// 		AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
	// 			{
	// 				Type: to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
	// 				Count: to.Ptr[int32](3),
	// 				CurrentOrchestratorVersion: to.Ptr("1.9.6"),
	// 				EnableNodePublicIP: to.Ptr(true),
	// 				MaxPods: to.Ptr[int32](110),
	// 				Mode: to.Ptr(armcontainerservice.AgentPoolModeSystem),
	// 				NodeImageVersion: to.Ptr("AKSUbuntu:1604:2020.03.11"),
	// 				OrchestratorVersion: to.Ptr("1.9.6"),
	// 				OSType: to.Ptr(armcontainerservice.OSTypeLinux),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				VMSize: to.Ptr("Standard_DS2_v2"),
	// 				Name: to.Ptr("nodepool1"),
	// 		}},
	// 		CurrentKubernetesVersion: to.Ptr("1.9.6"),
	// 		DiskEncryptionSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.Compute/diskEncryptionSets/des"),
	// 		DNSPrefix: to.Ptr("dnsprefix1"),
	// 		EnableRBAC: to.Ptr(true),
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
	// 		MaxAgentPools: to.Ptr[int32](1),
	// 		NetworkProfile: &armcontainerservice.NetworkProfile{
	// 			DNSServiceIP: to.Ptr("10.0.0.10"),
	// 			IPFamilies: []*armcontainerservice.IPFamily{
	// 				to.Ptr(armcontainerservice.IPFamilyIPv4)},
	// 				LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
	// 					AllocatedOutboundPorts: to.Ptr[int32](2000),
	// 					EffectiveOutboundIPs: []*armcontainerservice.ResourceReference{
	// 						{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip1"),
	// 						},
	// 						{
	// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/MC_rg1/providers/Microsoft.Network/publicIPAddresses/mgdoutboundip2"),
	// 					}},
	// 					IdleTimeoutInMinutes: to.Ptr[int32](10),
	// 					ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
	// 						Count: to.Ptr[int32](2),
	// 					},
	// 				},
	// 				LoadBalancerSKU: to.Ptr(armcontainerservice.LoadBalancerSKUBasic),
	// 				NetworkPlugin: to.Ptr(armcontainerservice.NetworkPluginKubenet),
	// 				OutboundType: to.Ptr(armcontainerservice.OutboundTypeLoadBalancer),
	// 				PodCidr: to.Ptr("10.244.0.0/16"),
	// 				PodCidrs: []*string{
	// 					to.Ptr("10.244.0.0/16")},
	// 					ServiceCidr: to.Ptr("10.0.0.0/16"),
	// 					ServiceCidrs: []*string{
	// 						to.Ptr("10.0.0.0/16")},
	// 					},
	// 					NodeResourceGroup: to.Ptr("MC_rg1_clustername1_location1"),
	// 					ProvisioningState: to.Ptr("Succeeded"),
	// 					SecurityProfile: &armcontainerservice.ManagedClusterSecurityProfile{
	// 						Defender: &armcontainerservice.ManagedClusterSecurityProfileDefender{
	// 							LogAnalyticsWorkspaceResourceID: to.Ptr("/subscriptions/SUB_ID/resourcegroups/RG_NAME/providers/microsoft.operationalinsights/workspaces/WORKSPACE_NAME"),
	// 							SecurityMonitoring: &armcontainerservice.ManagedClusterSecurityProfileDefenderSecurityMonitoring{
	// 								Enabled: to.Ptr(true),
	// 							},
	// 						},
	// 						WorkloadIdentity: &armcontainerservice.ManagedClusterSecurityProfileWorkloadIdentity{
	// 							Enabled: to.Ptr(true),
	// 						},
	// 					},
	// 					ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
	// 						ClientID: to.Ptr("clientid"),
	// 					},
	// 					WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
	// 						AdminUsername: to.Ptr("azureuser"),
	// 					},
	// 				},
	// 			}
}
