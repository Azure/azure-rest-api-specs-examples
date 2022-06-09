```go
package armcontainerservice_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
)

// x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/preview/2021-11-01-preview/examples/ManagedClustersAssociate_CRG.json
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
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"archv2": to.StringPtr(""),
				"tier":   to.StringPtr("production"),
			},
			Properties: &armcontainerservice.ManagedClusterProperties{
				AddonProfiles: map[string]*armcontainerservice.ManagedClusterAddonProfile{},
				AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
					{
						Type:                       armcontainerservice.AgentPoolType("VirtualMachineScaleSets").ToPtr(),
						CapacityReservationGroupID: to.StringPtr("<capacity-reservation-group-id>"),
						Count:                      to.Int32Ptr(3),
						EnableNodePublicIP:         to.BoolPtr(true),
						Mode:                       armcontainerservice.AgentPoolMode("System").ToPtr(),
						OSType:                     armcontainerservice.OSType("Linux").ToPtr(),
						VMSize:                     to.StringPtr("<vmsize>"),
						Name:                       to.StringPtr("<name>"),
					}},
				AutoScalerProfile: &armcontainerservice.ManagedClusterPropertiesAutoScalerProfile{
					ScaleDownDelayAfterAdd: to.StringPtr("<scale-down-delay-after-add>"),
					ScanInterval:           to.StringPtr("<scan-interval>"),
				},
				DiskEncryptionSetID:     to.StringPtr("<disk-encryption-set-id>"),
				DNSPrefix:               to.StringPtr("<dnsprefix>"),
				EnablePodSecurityPolicy: to.BoolPtr(true),
				EnableRBAC:              to.BoolPtr(true),
				KubernetesVersion:       to.StringPtr("<kubernetes-version>"),
				LinuxProfile: &armcontainerservice.LinuxProfile{
					AdminUsername: to.StringPtr("<admin-username>"),
					SSH: &armcontainerservice.SSHConfiguration{
						PublicKeys: []*armcontainerservice.SSHPublicKey{
							{
								KeyData: to.StringPtr("<key-data>"),
							}},
					},
				},
				NetworkProfile: &armcontainerservice.NetworkProfile{
					LoadBalancerProfile: &armcontainerservice.ManagedClusterLoadBalancerProfile{
						ManagedOutboundIPs: &armcontainerservice.ManagedClusterLoadBalancerProfileManagedOutboundIPs{
							Count: to.Int32Ptr(2),
						},
					},
					LoadBalancerSKU: armcontainerservice.LoadBalancerSKU("standard").ToPtr(),
					OutboundType:    armcontainerservice.OutboundType("loadBalancer").ToPtr(),
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
				Name: armcontainerservice.ManagedClusterSKUName("Basic").ToPtr(),
				Tier: armcontainerservice.ManagedClusterSKUTier("Free").ToPtr(),
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
	log.Printf("Response result: %#v\n", res.ManagedClustersClientCreateOrUpdateResult)
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-go/blob/sdk%2Fresourcemanager%2Fcontainerservice%2Farmcontainerservice%2Fv0.3.0/sdk/resourcemanager/containerservice/armcontainerservice/README.md) on how to add the SDK to your project and authenticate.
