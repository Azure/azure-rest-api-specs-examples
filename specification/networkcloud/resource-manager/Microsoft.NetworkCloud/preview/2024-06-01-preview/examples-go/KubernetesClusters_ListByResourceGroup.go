package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/98d74b2db60e46ceb7e3b75755e51519cd500485/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2024-06-01-preview/examples/KubernetesClusters_ListByResourceGroup.json
func ExampleKubernetesClustersClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewKubernetesClustersClient().NewListByResourceGroupPager("resourceGroupName", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.KubernetesClusterList = armnetworkcloud.KubernetesClusterList{
		// 	Value: []*armnetworkcloud.KubernetesCluster{
		// 		{
		// 			Name: to.Ptr("KubernetesClusterName"),
		// 			Type: to.Ptr("Microsoft.NetworkCloud/kubernetesClusters"),
		// 			ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/kubernetesClusters/kubernetesClusterName"),
		// 			SystemData: &armnetworkcloud.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:27:03.008Z"); return t}()),
		// 				CreatedBy: to.Ptr("identityA"),
		// 				CreatedByType: to.Ptr(armnetworkcloud.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-22T13:29:03.001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("identityB"),
		// 				LastModifiedByType: to.Ptr(armnetworkcloud.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("location"),
		// 			Tags: map[string]*string{
		// 				"key1": to.Ptr("myvalue1"),
		// 				"key2": to.Ptr("myvalue2"),
		// 			},
		// 			ExtendedLocation: &armnetworkcloud.ExtendedLocation{
		// 				Name: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armnetworkcloud.KubernetesClusterProperties{
		// 				AADConfiguration: &armnetworkcloud.AADConfiguration{
		// 					AdminGroupObjectIDs: []*string{
		// 						to.Ptr("ffffffff-ffff-ffff-ffff-ffffffffffff")},
		// 					},
		// 					AdministratorConfiguration: &armnetworkcloud.AdministratorConfiguration{
		// 						AdminUsername: to.Ptr("azure"),
		// 						SSHPublicKeys: []*armnetworkcloud.SSHPublicKey{
		// 							{
		// 								KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
		// 						}},
		// 					},
		// 					AttachedNetworkIDs: []*string{
		// 						to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l2Networks/l2NetworkName"),
		// 						to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 						to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/trunkedNetworks/trunkedNetworkName")},
		// 						AvailableUpgrades: []*armnetworkcloud.AvailableUpgrade{
		// 							{
		// 								AvailabilityLifecycle: to.Ptr(armnetworkcloud.AvailabilityLifecycleGenerallyAvailable),
		// 								Version: to.Ptr("1.XX.Y"),
		// 						}},
		// 						ClusterID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName"),
		// 						ConnectedClusterID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.Kubernetes/connectedClusters/connectedClusterName"),
		// 						ControlPlaneKubernetesVersion: to.Ptr("1.XX.Y"),
		// 						ControlPlaneNodeConfiguration: &armnetworkcloud.ControlPlaneNodeConfiguration{
		// 							AdministratorConfiguration: &armnetworkcloud.AdministratorConfiguration{
		// 								AdminUsername: to.Ptr("azure"),
		// 								SSHPublicKeys: []*armnetworkcloud.SSHPublicKey{
		// 									{
		// 										KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
		// 								}},
		// 							},
		// 							AvailabilityZones: []*string{
		// 								to.Ptr("1"),
		// 								to.Ptr("2"),
		// 								to.Ptr("3")},
		// 								Count: to.Ptr[int64](3),
		// 								VMSKUName: to.Ptr("NC_XXXX"),
		// 							},
		// 							DetailedStatus: to.Ptr(armnetworkcloud.KubernetesClusterDetailedStatusAvailable),
		// 							DetailedStatusMessage: to.Ptr("Kubernetes cluster is operational"),
		// 							FeatureStatuses: []*armnetworkcloud.FeatureStatus{
		// 								{
		// 									Name: to.Ptr("Feature1"),
		// 									DetailedStatus: to.Ptr(armnetworkcloud.FeatureDetailedStatusRunning),
		// 									DetailedStatusMessage: to.Ptr("No issues detected"),
		// 									Version: to.Ptr("1"),
		// 							}},
		// 							InitialAgentPoolConfigurations: []*armnetworkcloud.InitialAgentPoolConfiguration{
		// 								{
		// 									Name: to.Ptr("SystemPool-1"),
		// 									AdministratorConfiguration: &armnetworkcloud.AdministratorConfiguration{
		// 										AdminUsername: to.Ptr("azure"),
		// 										SSHPublicKeys: []*armnetworkcloud.SSHPublicKey{
		// 											{
		// 												KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
		// 										}},
		// 									},
		// 									AgentOptions: &armnetworkcloud.AgentOptions{
		// 										HugepagesCount: to.Ptr[int64](96),
		// 										HugepagesSize: to.Ptr(armnetworkcloud.HugepagesSizeOneG),
		// 									},
		// 									AttachedNetworkConfiguration: &armnetworkcloud.AttachedNetworkConfiguration{
		// 										L2Networks: []*armnetworkcloud.L2NetworkAttachmentConfiguration{
		// 											{
		// 												NetworkID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l2Networks/l2NetworkName"),
		// 												PluginType: to.Ptr(armnetworkcloud.KubernetesPluginTypeDPDK),
		// 										}},
		// 										L3Networks: []*armnetworkcloud.L3NetworkAttachmentConfiguration{
		// 											{
		// 												IpamEnabled: to.Ptr(armnetworkcloud.L3NetworkConfigurationIpamEnabledFalse),
		// 												NetworkID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 												PluginType: to.Ptr(armnetworkcloud.KubernetesPluginTypeSRIOV),
		// 										}},
		// 										TrunkedNetworks: []*armnetworkcloud.TrunkedNetworkAttachmentConfiguration{
		// 											{
		// 												NetworkID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/trunkedNetworks/trunkedNetworkName"),
		// 												PluginType: to.Ptr(armnetworkcloud.KubernetesPluginTypeMACVLAN),
		// 										}},
		// 									},
		// 									AvailabilityZones: []*string{
		// 										to.Ptr("1"),
		// 										to.Ptr("2"),
		// 										to.Ptr("3")},
		// 										Count: to.Ptr[int64](3),
		// 										Labels: []*armnetworkcloud.KubernetesLabel{
		// 											{
		// 												Key: to.Ptr("kubernetes.label"),
		// 												Value: to.Ptr("true"),
		// 										}},
		// 										Mode: to.Ptr(armnetworkcloud.AgentPoolModeSystem),
		// 										Taints: []*armnetworkcloud.KubernetesLabel{
		// 											{
		// 												Key: to.Ptr("kubernetes.taint"),
		// 												Value: to.Ptr("true"),
		// 										}},
		// 										UpgradeSettings: &armnetworkcloud.AgentPoolUpgradeSettings{
		// 											MaxSurge: to.Ptr("1"),
		// 										},
		// 										VMSKUName: to.Ptr("NC_XXXX"),
		// 								}},
		// 								KubernetesVersion: to.Ptr("1.XX.Y"),
		// 								ManagedResourceGroupConfiguration: &armnetworkcloud.ManagedResourceGroupConfiguration{
		// 									Name: to.Ptr("my-managed-rg"),
		// 									Location: to.Ptr("East US"),
		// 								},
		// 								NetworkConfiguration: &armnetworkcloud.NetworkConfiguration{
		// 									AttachedNetworkConfiguration: &armnetworkcloud.AttachedNetworkConfiguration{
		// 										L2Networks: []*armnetworkcloud.L2NetworkAttachmentConfiguration{
		// 											{
		// 												NetworkID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l2Networks/l2NetworkName"),
		// 												PluginType: to.Ptr(armnetworkcloud.KubernetesPluginTypeDPDK),
		// 										}},
		// 										L3Networks: []*armnetworkcloud.L3NetworkAttachmentConfiguration{
		// 											{
		// 												IpamEnabled: to.Ptr(armnetworkcloud.L3NetworkConfigurationIpamEnabledFalse),
		// 												NetworkID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 												PluginType: to.Ptr(armnetworkcloud.KubernetesPluginTypeSRIOV),
		// 										}},
		// 										TrunkedNetworks: []*armnetworkcloud.TrunkedNetworkAttachmentConfiguration{
		// 											{
		// 												NetworkID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/trunkedNetworks/trunkedNetworkName"),
		// 												PluginType: to.Ptr(armnetworkcloud.KubernetesPluginTypeMACVLAN),
		// 										}},
		// 									},
		// 									BgpServiceLoadBalancerConfiguration: &armnetworkcloud.BgpServiceLoadBalancerConfiguration{
		// 										BgpAdvertisements: []*armnetworkcloud.BgpAdvertisement{
		// 											{
		// 												AdvertiseToFabric: to.Ptr(armnetworkcloud.AdvertiseToFabricTrue),
		// 												Communities: []*string{
		// 													to.Ptr("64512:100")},
		// 													IPAddressPools: []*string{
		// 														to.Ptr("pool1")},
		// 														Peers: []*string{
		// 															to.Ptr("peer1")},
		// 													}},
		// 													BgpPeers: []*armnetworkcloud.ServiceLoadBalancerBgpPeer{
		// 														{
		// 															Name: to.Ptr("peer1"),
		// 															BfdEnabled: to.Ptr(armnetworkcloud.BfdEnabledFalse),
		// 															BgpMultiHop: to.Ptr(armnetworkcloud.BgpMultiHopFalse),
		// 															HoldTime: to.Ptr("P300s"),
		// 															KeepAliveTime: to.Ptr("P300s"),
		// 															MyAsn: to.Ptr[int64](64512),
		// 															PeerAddress: to.Ptr("203.0.113.254"),
		// 															PeerAsn: to.Ptr[int64](64497),
		// 															PeerPort: to.Ptr[int64](179),
		// 													}},
		// 													FabricPeeringEnabled: to.Ptr(armnetworkcloud.FabricPeeringEnabledTrue),
		// 													IPAddressPools: []*armnetworkcloud.IPAddressPool{
		// 														{
		// 															Name: to.Ptr("pool1"),
		// 															Addresses: []*string{
		// 																to.Ptr("198.51.102.0/24")},
		// 																AutoAssign: to.Ptr(armnetworkcloud.BfdEnabledTrue),
		// 																OnlyUseHostIPs: to.Ptr(armnetworkcloud.BfdEnabledTrue),
		// 														}},
		// 													},
		// 													CloudServicesNetworkID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/cloudServicesNetworks/cloudServicesNetworkName"),
		// 													CniNetworkID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 													DNSServiceIP: to.Ptr("198.51.101.2"),
		// 													PodCidrs: []*string{
		// 														to.Ptr("198.51.100.0/24")},
		// 														ServiceCidrs: []*string{
		// 															to.Ptr("198.51.101.0/24")},
		// 														},
		// 														Nodes: []*armnetworkcloud.KubernetesClusterNode{
		// 															{
		// 																Name: to.Ptr("machine1"),
		// 																AgentPoolID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/kubernetesClusters/kubernetesClusterName/agentPools/agentPoolName"),
		// 																AvailabilityZone: to.Ptr("1"),
		// 																BareMetalMachineID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
		// 																CPUCores: to.Ptr[int64](20),
		// 																DetailedStatus: to.Ptr(armnetworkcloud.KubernetesClusterNodeDetailedStatusAvailable),
		// 																DetailedStatusMessage: to.Ptr("No issues detected"),
		// 																DiskSizeGB: to.Ptr[int64](120),
		// 																Image: to.Ptr("myacr.azurecr.io/foobar:latest"),
		// 																KubernetesVersion: to.Ptr("1.XX.Y"),
		// 																Labels: []*armnetworkcloud.KubernetesLabel{
		// 																	{
		// 																		Key: to.Ptr("kubernetes.label"),
		// 																		Value: to.Ptr("true"),
		// 																}},
		// 																MemorySizeGB: to.Ptr[int64](256),
		// 																Mode: to.Ptr(armnetworkcloud.AgentPoolModeSystem),
		// 																NetworkAttachments: []*armnetworkcloud.NetworkAttachment{
		// 																	{
		// 																		AttachedNetworkID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 																		DefaultGateway: to.Ptr(armnetworkcloud.DefaultGatewayTrue),
		// 																		IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
		// 																		IPv4Address: to.Ptr("198.51.100.1"),
		// 																		IPv6Address: to.Ptr("2001:0db8:0000:0000:0000:0000:0000:0000"),
		// 																		MacAddress: to.Ptr("bf:1c:29:31:31:1f"),
		// 																		NetworkAttachmentName: to.Ptr("netAttachName01"),
		// 																}},
		// 																PowerState: to.Ptr(armnetworkcloud.KubernetesNodePowerStateOn),
		// 																Role: to.Ptr(armnetworkcloud.KubernetesNodeRoleControlPlane),
		// 																Taints: []*armnetworkcloud.KubernetesLabel{
		// 																	{
		// 																		Key: to.Ptr("kubernetes.taint"),
		// 																		Value: to.Ptr("true"),
		// 																}},
		// 																VMSKUName: to.Ptr("NC_XXXX"),
		// 														}},
		// 														ProvisioningState: to.Ptr(armnetworkcloud.KubernetesClusterProvisioningStateSucceeded),
		// 													},
		// 											}},
		// 										}
	}
}
