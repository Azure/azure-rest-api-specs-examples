package armcontainerservice_test

import (
	"context"
	"log"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-09-01/examples/ManagedClustersCreate_Snapshot.json
func ExampleManagedClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armcontainerservice.ManagedCluster{
			Resource: armcontainerservice.Resource{
				Location: to.StringPtr("<location>"),
				Tags: map[string]*string{
					"archv2": to.StringPtr(""),
					"tier":   to.StringPtr("production"),
				},
			},
			Properties: &armcontainerservice.ManagedClusterProperties{
				AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
				AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
					{
						ManagedClusterAgentPoolProfileProperties: armcontainerservice.ManagedClusterAgentPoolProfileProperties{
							Type:  armcontainerservice.AgentPoolTypeVirtualMachineScaleSets.ToPtr(),
							Count: to.Int32Ptr(3),
							CreationData: &armcontainerservice.CreationData{
								SourceResourceID: to.StringPtr("<source-resource-id>"),
							},
							EnableFIPS:         to.BoolPtr(true),
							EnableNodePublicIP: to.BoolPtr(true),
							Mode:               armcontainerservice.AgentPoolModeSystem.ToPtr(),
							OSType:             armcontainerservice.OSTypeLinux.ToPtr(),
							VMSize:             to.StringPtr("<vmsize>"),
						},
						Name: to.StringPtr("<name>"),
					}},
				AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
					ScaleDownDelayAfterAdd: to.StringPtr("<scale-down-delay-after-add>"),
					ScanInterval:           to.StringPtr("<scan-interval>"),
				},
				DiskEncryptionSetID:     to.StringPtr("<disk-encryption-set-id>"),
				DNSPrefix:               to.StringPtr("<dnsprefix>"),
				EnablePodSecurityPolicy: to.BoolPtr(false),
				EnableRBAC:              to.BoolPtr(true),
				KubernetesVersion:       to.StringPtr("<kubernetes-version>"),
				LinuxProfile: &armcontainerservice.ContainerServiceLinuxProfile{
					AdminUsername: to.StringPtr("<admin-username>"),
					SSH: &armcontainerservice.ContainerServiceSSHConfiguration{
						PublicKeys: []*armcontainerservice.ContainerServiceSSHPublicKey{
							{
								KeyData: to.StringPtr("<key-data>"),
							}},
					},
				},
				NetworkProfile: &armcontainerservice.ContainerServiceNetworkProfile{
					LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
						ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
							Count: to.Int32Ptr(2),
						},
					},
					LoadBalancerSKU: armcontainerservice.LoadBalancerSKUStandard.ToPtr(),
					OutboundType:    armcontainerservice.OutboundTypeLoadBalancer.ToPtr(),
				},
				ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
					ClientID: to.StringPtr("<client-id>"),
					Secret:   to.StringPtr("<secret>"),
				},
				WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
					AdminPassword: to.StringPtr("<admin-password>"),
					AdminUsername: to.StringPtr("<admin-username>"),
				},
			},
			SKU: &armcontainerservice.ManagedClusterSKU{
				Name: armcontainerservice.ManagedClusterSKUNameBasic.ToPtr(),
				Tier: armcontainerservice.ManagedClusterSKUTierFree.ToPtr(),
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ManagedCluster.ID: %s\n", *res.ID)
}
