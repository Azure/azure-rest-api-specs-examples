package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/PutProvisionedCluster.json
func ExampleProvisionedClustersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armhybridcontainerservice.NewProvisionedClustersClient("a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx, "test-arcappliance-resgrp", "test-hybridakscluster", armhybridcontainerservice.ProvisionedClusters{
		Location: to.Ptr("westus"),
		ExtendedLocation: &armhybridcontainerservice.ProvisionedClustersExtendedLocation{
			Name: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"),
			Type: to.Ptr("CustomLocation"),
		},
		Properties: &armhybridcontainerservice.ProvisionedClustersAllProperties{
			AgentPoolProfiles: []*armhybridcontainerservice.NamedAgentPoolProfile{
				{
					Name:   to.Ptr("default-nodepool-1"),
					Count:  to.Ptr[int32](1),
					OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
					VMSize: to.Ptr("Standard_A4_v2"),
				}},
			CloudProviderProfile: &armhybridcontainerservice.CloudProviderProfile{
				InfraNetworkProfile: &armhybridcontainerservice.CloudProviderProfileInfraNetworkProfile{
					VnetSubnetIDs: []*string{
						to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.HybridContainerService/virtualNetworks/test-vnet-static")},
				},
				InfraStorageProfile: &armhybridcontainerservice.CloudProviderProfileInfraStorageProfile{
					StorageSpaceIDs: []*string{
						to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.HybridContainerService/storageSpaces/test-storage")},
				},
			},
			ControlPlane: &armhybridcontainerservice.ControlPlaneProfile{
				LinuxProfile: &armhybridcontainerservice.LinuxProfileProperties{
					SSH: &armhybridcontainerservice.LinuxProfilePropertiesSSH{
						PublicKeys: []*armhybridcontainerservice.LinuxProfilePropertiesSSHPublicKeysItem{
							{
								KeyData: to.Ptr("ssh-rsa AAAAB1NzaC1yc2EAAAADAQABAAACAQCY......"),
							}},
					},
				},
				Count:  to.Ptr[int32](1),
				OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
				VMSize: to.Ptr("Standard_A4_v2"),
			},
			KubernetesVersion: to.Ptr("v1.20.5"),
			LinuxProfile: &armhybridcontainerservice.LinuxProfileProperties{
				SSH: &armhybridcontainerservice.LinuxProfilePropertiesSSH{
					PublicKeys: []*armhybridcontainerservice.LinuxProfilePropertiesSSHPublicKeysItem{
						{
							KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCY......."),
						}},
				},
			},
			NetworkProfile: &armhybridcontainerservice.NetworkProfile{
				LoadBalancerProfile: &armhybridcontainerservice.LoadBalancerProfile{
					LinuxProfile: &armhybridcontainerservice.LinuxProfileProperties{
						SSH: &armhybridcontainerservice.LinuxProfilePropertiesSSH{
							PublicKeys: []*armhybridcontainerservice.LinuxProfilePropertiesSSHPublicKeysItem{
								{
									KeyData: to.Ptr("ssh-rsa AAAAB2NzaC1yc2EAAAADAQABAAACAQCY......"),
								}},
						},
					},
					Count:  to.Ptr[int32](1),
					OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
					VMSize: to.Ptr("Standard_K8S3_v1"),
				},
				LoadBalancerSKU: to.Ptr(armhybridcontainerservice.LoadBalancerSKUUnstackedHaproxy),
				NetworkPolicy:   to.Ptr(armhybridcontainerservice.NetworkPolicyCalico),
				PodCidr:         to.Ptr("10.244.0.0/16"),
			},
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
