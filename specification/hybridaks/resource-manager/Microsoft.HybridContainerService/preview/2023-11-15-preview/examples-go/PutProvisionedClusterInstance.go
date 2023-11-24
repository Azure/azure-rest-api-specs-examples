package armhybridcontainerservice_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcontainerservice/armhybridcontainerservice"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4bb583bcb67c2bf448712f2bd1593a64a7a8f139/specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2023-11-15-preview/examples/PutProvisionedClusterInstance.json
func ExampleProvisionedClusterInstancesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridcontainerservice.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewProvisionedClusterInstancesClient().BeginCreateOrUpdate(ctx, "subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster", armhybridcontainerservice.ProvisionedClusters{
		ExtendedLocation: &armhybridcontainerservice.ExtendedLocation{
			Name: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"),
			Type: to.Ptr(armhybridcontainerservice.ExtendedLocationTypesCustomLocation),
		},
		Properties: &armhybridcontainerservice.ProvisionedClusterProperties{
			AgentPoolProfiles: []*armhybridcontainerservice.NamedAgentPoolProfile{
				{
					Name:   to.Ptr("default-nodepool-1"),
					OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
					Count:  to.Ptr[int32](1),
					VMSize: to.Ptr("Standard_A4_v2"),
				}},
			CloudProviderProfile: &armhybridcontainerservice.CloudProviderProfile{
				InfraNetworkProfile: &armhybridcontainerservice.CloudProviderProfileInfraNetworkProfile{
					VnetSubnetIDs: []*string{
						to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.AzureStackHCI/logicalNetworks/test-vnet-static")},
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
				OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
				Count:  to.Ptr[int32](1),
				VMSize: to.Ptr("Standard_A4_v2"),
			},
			KubernetesVersion: to.Ptr("v1.20.5"),
			LicenseProfile: &armhybridcontainerservice.ProvisionedClusterLicenseProfile{
				AzureHybridBenefit: to.Ptr(armhybridcontainerservice.AzureHybridBenefitNotApplicable),
			},
			LinuxProfile: &armhybridcontainerservice.LinuxProfileProperties{
				SSH: &armhybridcontainerservice.LinuxProfilePropertiesSSH{
					PublicKeys: []*armhybridcontainerservice.LinuxProfilePropertiesSSHPublicKeysItem{
						{
							KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCY......."),
						}},
				},
			},
			NetworkProfile: &armhybridcontainerservice.NetworkProfile{
				NetworkPolicy: to.Ptr(armhybridcontainerservice.NetworkPolicyCalico),
				PodCidr:       to.Ptr("10.244.0.0/16"),
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
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ProvisionedClusters = armhybridcontainerservice.ProvisionedClusters{
	// 	Name: to.Ptr("test-hybridakscluster"),
	// 	Type: to.Ptr("Microsoft.HybridContainerService/provisionedClusterInstances"),
	// 	ID: to.Ptr("/subscriptions/fd3c3665-1729-4b7b-9a38-238e83b0f98b/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/test-hybridakscluster/providers/Microsoft.HybridContainerService/provisionedClusterInstances/default"),
	// 	ExtendedLocation: &armhybridcontainerservice.ExtendedLocation{
	// 		Name: to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"),
	// 		Type: to.Ptr(armhybridcontainerservice.ExtendedLocationTypesCustomLocation),
	// 	},
	// 	Properties: &armhybridcontainerservice.ProvisionedClusterProperties{
	// 		AgentPoolProfiles: []*armhybridcontainerservice.NamedAgentPoolProfile{
	// 			{
	// 				Name: to.Ptr("default-nodepool-1"),
	// 				OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
	// 				Count: to.Ptr[int32](1),
	// 				VMSize: to.Ptr("Standard_A4_v2"),
	// 		}},
	// 		CloudProviderProfile: &armhybridcontainerservice.CloudProviderProfile{
	// 			InfraNetworkProfile: &armhybridcontainerservice.CloudProviderProfileInfraNetworkProfile{
	// 				VnetSubnetIDs: []*string{
	// 					to.Ptr("/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourceGroups/test-arcappliance-resgrp/providers/Microsoft.AzureStackHCI/logicalNetworks/test-vnet-static")},
	// 				},
	// 			},
	// 			ControlPlane: &armhybridcontainerservice.ControlPlaneProfile{
	// 				LinuxProfile: &armhybridcontainerservice.LinuxProfileProperties{
	// 					SSH: &armhybridcontainerservice.LinuxProfilePropertiesSSH{
	// 						PublicKeys: []*armhybridcontainerservice.LinuxProfilePropertiesSSHPublicKeysItem{
	// 							{
	// 								KeyData: to.Ptr("ssh-rsa AAAAB1NzaC1yc2EAAAADAQABAAACAQCY......"),
	// 						}},
	// 					},
	// 				},
	// 				OSType: to.Ptr(armhybridcontainerservice.OsTypeLinux),
	// 				Count: to.Ptr[int32](1),
	// 				VMSize: to.Ptr("Standard_A4_v2"),
	// 			},
	// 			KubernetesVersion: to.Ptr("v1.20.5"),
	// 			LicenseProfile: &armhybridcontainerservice.ProvisionedClusterLicenseProfile{
	// 				AzureHybridBenefit: to.Ptr(armhybridcontainerservice.AzureHybridBenefitNotApplicable),
	// 			},
	// 			LinuxProfile: &armhybridcontainerservice.LinuxProfileProperties{
	// 				SSH: &armhybridcontainerservice.LinuxProfilePropertiesSSH{
	// 					PublicKeys: []*armhybridcontainerservice.LinuxProfilePropertiesSSHPublicKeysItem{
	// 						{
	// 							KeyData: to.Ptr("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCY......."),
	// 					}},
	// 				},
	// 			},
	// 			NetworkProfile: &armhybridcontainerservice.NetworkProfile{
	// 				NetworkPolicy: to.Ptr(armhybridcontainerservice.NetworkPolicyCalico),
	// 				PodCidr: to.Ptr("10.244.0.0/16"),
	// 			},
	// 			ProvisioningState: to.Ptr(armhybridcontainerservice.ResourceProvisioningStateSucceeded),
	// 		},
	// 	}
}
