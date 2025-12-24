package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/46c51b03d99b113ecc3b38883e3cb2d395fe94a4/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-09-01/examples/AgentPools_ListByKubernetesCluster.json
func ExampleAgentPoolsClient_NewListByKubernetesClusterPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAgentPoolsClient().NewListByKubernetesClusterPager("resourceGroupName", "kubernetesClusterName", &armnetworkcloud.AgentPoolsClientListByKubernetesClusterOptions{Top: nil,
		SkipToken: nil,
	})
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
		// page.AgentPoolList = armnetworkcloud.AgentPoolList{
		// 	Value: []*armnetworkcloud.AgentPool{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.NetworkCloud/kubernetesClusters/agentPools"),
		// 			ID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/kubernetesClusters/kubernetesClusterName/agentPools/agentPoolName"),
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
		// 			Properties: &armnetworkcloud.AgentPoolProperties{
		// 				AdministratorConfiguration: &armnetworkcloud.AdministratorConfiguration{
		// 					AdminUsername: to.Ptr("azure"),
		// 					SSHPublicKeys: []*armnetworkcloud.SSHPublicKey{
		// 						{
		// 							KeyData: to.Ptr("ssh-rsa AAtsE3njSONzDYRIZv/WLjVuMfrUSByHp+jfaaOLHTIIB4fJvo6dQUZxE20w2iDHV3tEkmnTo84eba97VMueQD6OzJPEyWZMRpz8UYWOd0IXeRqiFu1lawNblZhwNT/ojNZfpB3af/YDzwQCZgTcTRyNNhL4o/blKUmug0daSsSXISTRnIDpcf5qytjs1Xo+yYyJMvzLL59mhAyb3p/cD+Y3/s3WhAx+l0XOKpzXnblrv9d3q4c2tWmm/SyFqthaqd0= admin@vm"),
		// 					}},
		// 				},
		// 				AgentOptions: &armnetworkcloud.AgentOptions{
		// 					HugepagesCount: to.Ptr[int64](96),
		// 					HugepagesSize: to.Ptr(armnetworkcloud.HugepagesSizeOneG),
		// 				},
		// 				AttachedNetworkConfiguration: &armnetworkcloud.AttachedNetworkConfiguration{
		// 					L2Networks: []*armnetworkcloud.L2NetworkAttachmentConfiguration{
		// 						{
		// 							NetworkID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l2Networks/l2NetworkName"),
		// 							PluginType: to.Ptr(armnetworkcloud.KubernetesPluginTypeDPDK),
		// 					}},
		// 					L3Networks: []*armnetworkcloud.L3NetworkAttachmentConfiguration{
		// 						{
		// 							IpamEnabled: to.Ptr(armnetworkcloud.L3NetworkConfigurationIpamEnabledFalse),
		// 							NetworkID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 							PluginType: to.Ptr(armnetworkcloud.KubernetesPluginTypeSRIOV),
		// 					}},
		// 					TrunkedNetworks: []*armnetworkcloud.TrunkedNetworkAttachmentConfiguration{
		// 						{
		// 							NetworkID: to.Ptr("/subscriptions/123e4567-e89b-12d3-a456-426655440000/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/trunkedNetworks/trunkedNetworkName"),
		// 							PluginType: to.Ptr(armnetworkcloud.KubernetesPluginTypeMACVLAN),
		// 					}},
		// 				},
		// 				AvailabilityZones: []*string{
		// 					to.Ptr("1"),
		// 					to.Ptr("2"),
		// 					to.Ptr("3")},
		// 					Count: to.Ptr[int64](3),
		// 					DetailedStatus: to.Ptr(armnetworkcloud.AgentPoolDetailedStatusAvailable),
		// 					DetailedStatusMessage: to.Ptr("Agent pool is available"),
		// 					KubernetesVersion: to.Ptr("1.XX.Y"),
		// 					Labels: []*armnetworkcloud.KubernetesLabel{
		// 						{
		// 							Key: to.Ptr("kubernetes.label"),
		// 							Value: to.Ptr("true"),
		// 					}},
		// 					Mode: to.Ptr(armnetworkcloud.AgentPoolModeSystem),
		// 					ProvisioningState: to.Ptr(armnetworkcloud.AgentPoolProvisioningStateSucceeded),
		// 					Taints: []*armnetworkcloud.KubernetesLabel{
		// 						{
		// 							Key: to.Ptr("kubernetes.taint"),
		// 							Value: to.Ptr("true:NoSchedule"),
		// 					}},
		// 					UpgradeSettings: &armnetworkcloud.AgentPoolUpgradeSettings{
		// 						MaxSurge: to.Ptr("1"),
		// 					},
		// 					VMSKUName: to.Ptr("NC_P46_224_v1"),
		// 				},
		// 		}},
		// 	}
	}
}
