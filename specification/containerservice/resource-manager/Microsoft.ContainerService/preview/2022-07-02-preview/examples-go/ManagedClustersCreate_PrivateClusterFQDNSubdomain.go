package armcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2022-07-02-preview/examples/ManagedClustersCreate_PrivateClusterFQDNSubdomain.json
func ExampleManagedClustersClient_BeginCreateOrUpdate_createManagedPrivateClusterWithFqdnSubdomainSpecified() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewManagedClustersClient("subid1", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "rg1", "clustername1", armcontainerservice.ManagedCluster{
		Location: to.Ptr("location1"),
		Tags: map[string]*string{
			"archv2": to.Ptr(""),
			"tier":   to.Ptr("production"),
		},
		Properties: &armcontainerservice.ManagedClusterProperties{
			AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
			AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
				{
					Type:                   to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
					Count:                  to.Ptr[int32](3),
					EnableEncryptionAtHost: to.Ptr(true),
					EnableNodePublicIP:     to.Ptr(true),
					Mode:                   to.Ptr(armcontainerservice.AgentPoolModeSystem),
					OSType:                 to.Ptr(armcontainerservice.OSTypeLinux),
					VMSize:                 to.Ptr("Standard_DS2_v2"),
					Name:                   to.Ptr("nodepool1"),
				}},
			APIServerAccessProfile: &armcontainerservice.ManagedClusterAPIServerAccessProfile{
				EnablePrivateCluster: to.Ptr(true),
				PrivateDNSZone:       to.Ptr("/subscriptions/subid1/resourcegroups/rg1/providers/Microsoft.Network/privateDnsZones/privatelink.location1.azmk8s.io"),
			},
			AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
				ScaleDownDelayAfterAdd: to.Ptr("15m"),
				ScanInterval:           to.Ptr("20s"),
			},
			EnablePodSecurityPolicy: to.Ptr(true),
			EnableRBAC:              to.Ptr(true),
			FqdnSubdomain:           to.Ptr("domain1"),
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
			Name: to.Ptr(armcontainerservice.ManagedClusterSKUNameBasic),
			Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
