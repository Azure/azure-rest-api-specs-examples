package armcontainerservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-03-01/examples/ManagedClustersCreate_Snapshot.json
func ExampleManagedClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcontainerservice.NewManagedClustersClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armcontainerservice.ManagedCluster{
			Location: to.Ptr("<location>"),
			Tags: map[string]*string{
				"archv2": to.Ptr(""),
				"tier":   to.Ptr("production"),
			},
			Properties: &armcontainerservice.ManagedClusterProperties{
				AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
				AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
					{
						Type:  to.Ptr(armcontainerservice.AgentPoolTypeVirtualMachineScaleSets),
						Count: to.Ptr[int32](3),
						CreationData: &armcontainerservice.CreationData{
							SourceResourceID: to.Ptr("<source-resource-id>"),
						},
						EnableFIPS:         to.Ptr(true),
						EnableNodePublicIP: to.Ptr(true),
						Mode:               to.Ptr(armcontainerservice.AgentPoolModeSystem),
						OSType:             to.Ptr(armcontainerservice.OSTypeLinux),
						VMSize:             to.Ptr("<vmsize>"),
						Name:               to.Ptr("<name>"),
					}},
				AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
					ScaleDownDelayAfterAdd: to.Ptr("<scale-down-delay-after-add>"),
					ScanInterval:           to.Ptr("<scan-interval>"),
				},
				DiskEncryptionSetID:     to.Ptr("<disk-encryption-set-id>"),
				DNSPrefix:               to.Ptr("<dnsprefix>"),
				EnablePodSecurityPolicy: to.Ptr(false),
				EnableRBAC:              to.Ptr(true),
				KubernetesVersion:       to.Ptr("<kubernetes-version>"),
				LinuxProfile: &armcontainerservice.LinuxProfile{
					AdminUsername: to.Ptr("<admin-username>"),
					SSH: &armcontainerservice.SSHConfiguration{
						PublicKeys: []*armcontainerservice.SSHPublicKey{
							{
								KeyData: to.Ptr("<key-data>"),
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
					ClientID: to.Ptr("<client-id>"),
					Secret:   to.Ptr("<secret>"),
				},
				WindowsProfile: &armcontainerservice.ManagedClusterWindowsProfile{
					AdminPassword: to.Ptr("<admin-password>"),
					AdminUsername: to.Ptr("<admin-username>"),
				},
			},
			SKU: &armcontainerservice.ManagedClusterSKU{
				Name: to.Ptr(armcontainerservice.ManagedClusterSKUNameBasic),
				Tier: to.Ptr(armcontainerservice.ManagedClusterSKUTierFree),
			},
		},
		&armcontainerservice.ManagedClustersClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
