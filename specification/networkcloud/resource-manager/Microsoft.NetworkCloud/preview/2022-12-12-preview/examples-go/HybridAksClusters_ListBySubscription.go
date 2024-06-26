package armnetworkcloud_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/17aa6a1314de5aafef059d9aa2229901df506e75/specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/HybridAksClusters_ListBySubscription.json
func ExampleHybridAksClustersClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetworkcloud.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewHybridAksClustersClient().NewListBySubscriptionPager(nil)
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
		// page.HybridAksClusterList = armnetworkcloud.HybridAksClusterList{
		// 	Value: []*armnetworkcloud.HybridAksCluster{
		// 		{
		// 			Name: to.Ptr("HybridAksClusterName"),
		// 			Type: to.Ptr("Microsoft.NetworkCloud/hybridAksClusters"),
		// 			ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/hybridAksClusters/hybridAksClusterName"),
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
		// 				Name: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName"),
		// 				Type: to.Ptr("CustomLocation"),
		// 			},
		// 			Properties: &armnetworkcloud.HybridAksClusterProperties{
		// 				AssociatedNetworkIDs: []*string{
		// 					to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName")},
		// 					CloudServicesNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/cloudServicesNetworks/cloudServicesNetworkName"),
		// 					ClusterID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/clusters/clusterName"),
		// 					ControlPlaneCount: to.Ptr[int64](4),
		// 					ControlPlaneNodes: []*armnetworkcloud.NodeConfiguration{
		// 						{
		// 							CPUCores: to.Ptr[int64](2),
		// 							DiskSizeGB: to.Ptr[int64](128),
		// 							MemorySizeGB: to.Ptr[int64](8),
		// 							Nodes: []*armnetworkcloud.Node{
		// 								{
		// 									BareMetalMachineID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
		// 									ImageID: to.Ptr("myacr.azurecr.io/worker:latest"),
		// 									NetworkAttachments: []*armnetworkcloud.NetworkAttachment{
		// 										{
		// 											AttachedNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 											DefaultGateway: to.Ptr(armnetworkcloud.DefaultGatewayTrue),
		// 											IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
		// 											IPv4Address: to.Ptr("198.51.100.1"),
		// 											IPv6Address: to.Ptr("2001:0db8:0000:0000:0000:0000:0000:0001"),
		// 											MacAddress: to.Ptr("bf:1c:29:31:31:1f"),
		// 											NetworkAttachmentName: to.Ptr("netAttachName01"),
		// 									}},
		// 									NodeName: to.Ptr("control0"),
		// 									PowerState: to.Ptr(armnetworkcloud.HybridAksClusterMachinePowerStateOn),
		// 								},
		// 								{
		// 									BareMetalMachineID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
		// 									ImageID: to.Ptr("myacr.azurecr.io/worker:latest"),
		// 									NetworkAttachments: []*armnetworkcloud.NetworkAttachment{
		// 										{
		// 											AttachedNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 											DefaultGateway: to.Ptr(armnetworkcloud.DefaultGatewayTrue),
		// 											IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
		// 											IPv4Address: to.Ptr("198.51.100.1"),
		// 											IPv6Address: to.Ptr("2001:0db8:0000:0000:0000:0000:0000:0001"),
		// 											MacAddress: to.Ptr("bf:1c:29:31:31:1f"),
		// 											NetworkAttachmentName: to.Ptr("netAttachName01"),
		// 									}},
		// 									NodeName: to.Ptr("control1"),
		// 									PowerState: to.Ptr(armnetworkcloud.HybridAksClusterMachinePowerStateOn),
		// 								},
		// 								{
		// 									BareMetalMachineID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
		// 									ImageID: to.Ptr("myacr.azurecr.io/worker:latest"),
		// 									NetworkAttachments: []*armnetworkcloud.NetworkAttachment{
		// 										{
		// 											AttachedNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 											DefaultGateway: to.Ptr(armnetworkcloud.DefaultGatewayTrue),
		// 											IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
		// 											IPv4Address: to.Ptr("198.51.100.1"),
		// 											IPv6Address: to.Ptr("2001:0db8:0000:0000:0000:0000:0000:0001"),
		// 											MacAddress: to.Ptr("bf:1c:29:31:31:1f"),
		// 											NetworkAttachmentName: to.Ptr("netAttachName01"),
		// 									}},
		// 									NodeName: to.Ptr("control2"),
		// 									PowerState: to.Ptr(armnetworkcloud.HybridAksClusterMachinePowerStateOn),
		// 								},
		// 								{
		// 									BareMetalMachineID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
		// 									ImageID: to.Ptr("myacr.azurecr.io/worker:latest"),
		// 									NetworkAttachments: []*armnetworkcloud.NetworkAttachment{
		// 										{
		// 											AttachedNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 											DefaultGateway: to.Ptr(armnetworkcloud.DefaultGatewayTrue),
		// 											IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
		// 											IPv4Address: to.Ptr("198.51.100.1"),
		// 											IPv6Address: to.Ptr("2001:0db8:0000:0000:0000:0000:0000:0001"),
		// 											MacAddress: to.Ptr("bf:1c:29:31:31:1f"),
		// 											NetworkAttachmentName: to.Ptr("netAttachName01"),
		// 									}},
		// 									NodeName: to.Ptr("control3"),
		// 									PowerState: to.Ptr(armnetworkcloud.HybridAksClusterMachinePowerStateOn),
		// 							}},
		// 							VMCount: to.Ptr[int64](4),
		// 							VMSize: to.Ptr("NC_G2_v1"),
		// 					}},
		// 					DefaultCniNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/defaultCniNetworks/defaultCniNetworkName"),
		// 					DetailedStatus: to.Ptr(armnetworkcloud.HybridAksClusterDetailedStatusAvailable),
		// 					DetailedStatusMessage: to.Ptr("Hybrid AKS cluster is operational"),
		// 					HybridAksProvisionedClusterID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.HybridContainerService/provisionedClusters/hybridAksClusterName"),
		// 					ProvisioningState: to.Ptr(armnetworkcloud.HybridAksClusterProvisioningStateSucceeded),
		// 					Volumes: []*string{
		// 						to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/volumes/volumeName")},
		// 						WorkerCount: to.Ptr[int64](8),
		// 						WorkerNodes: []*armnetworkcloud.NodeConfiguration{
		// 							{
		// 								AgentPoolID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.HybridContainerService/provisionedClusters/provisionedClustersName/agentPools/regularWorkers"),
		// 								AgentPoolName: to.Ptr("regularWorkers"),
		// 								CPUCores: to.Ptr[int64](2),
		// 								DiskSizeGB: to.Ptr[int64](128),
		// 								MemorySizeGB: to.Ptr[int64](8),
		// 								NodePoolName: to.Ptr("regularWorkers"),
		// 								Nodes: []*armnetworkcloud.Node{
		// 									{
		// 										BareMetalMachineID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
		// 										ImageID: to.Ptr("myacr.azurecr.io/worker:latest"),
		// 										NetworkAttachments: []*armnetworkcloud.NetworkAttachment{
		// 											{
		// 												AttachedNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 												DefaultGateway: to.Ptr(armnetworkcloud.DefaultGatewayTrue),
		// 												IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
		// 												IPv4Address: to.Ptr("198.51.100.1"),
		// 												IPv6Address: to.Ptr("2001:0db8:0000:0000:0000:0000:0000:0001"),
		// 												MacAddress: to.Ptr("bf:1c:29:31:31:1f"),
		// 												NetworkAttachmentName: to.Ptr("netAttachName01"),
		// 										}},
		// 										NodeName: to.Ptr("worker0"),
		// 										PowerState: to.Ptr(armnetworkcloud.HybridAksClusterMachinePowerStateOn),
		// 									},
		// 									{
		// 										BareMetalMachineID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
		// 										ImageID: to.Ptr("myacr.azurecr.io/worker:latest"),
		// 										NetworkAttachments: []*armnetworkcloud.NetworkAttachment{
		// 											{
		// 												AttachedNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 												DefaultGateway: to.Ptr(armnetworkcloud.DefaultGatewayTrue),
		// 												IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
		// 												IPv4Address: to.Ptr("198.51.100.1"),
		// 												IPv6Address: to.Ptr("2001:0db8:0000:0000:0000:0000:0000:0001"),
		// 												MacAddress: to.Ptr("bf:1c:29:31:31:1f"),
		// 												NetworkAttachmentName: to.Ptr("netAttachName01"),
		// 										}},
		// 										NodeName: to.Ptr("worker1"),
		// 										PowerState: to.Ptr(armnetworkcloud.HybridAksClusterMachinePowerStateOn),
		// 									},
		// 									{
		// 										BareMetalMachineID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
		// 										ImageID: to.Ptr("myacr.azurecr.io/worker:latest"),
		// 										NetworkAttachments: []*armnetworkcloud.NetworkAttachment{
		// 											{
		// 												AttachedNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 												DefaultGateway: to.Ptr(armnetworkcloud.DefaultGatewayTrue),
		// 												IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
		// 												IPv4Address: to.Ptr("198.51.100.1"),
		// 												IPv6Address: to.Ptr("2001:0db8:0000:0000:0000:0000:0000:0001"),
		// 												MacAddress: to.Ptr("bf:1c:29:31:31:1f"),
		// 												NetworkAttachmentName: to.Ptr("netAttachName01"),
		// 										}},
		// 										NodeName: to.Ptr("worker2"),
		// 										PowerState: to.Ptr(armnetworkcloud.HybridAksClusterMachinePowerStateOn),
		// 									},
		// 									{
		// 										BareMetalMachineID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
		// 										ImageID: to.Ptr("myacr.azurecr.io/worker:latest"),
		// 										NetworkAttachments: []*armnetworkcloud.NetworkAttachment{
		// 											{
		// 												AttachedNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 												DefaultGateway: to.Ptr(armnetworkcloud.DefaultGatewayTrue),
		// 												IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
		// 												IPv4Address: to.Ptr("198.51.100.1"),
		// 												IPv6Address: to.Ptr("2001:0db8:0000:0000:0000:0000:0000:0001"),
		// 												MacAddress: to.Ptr("bf:1c:29:31:31:1f"),
		// 												NetworkAttachmentName: to.Ptr("netAttachName01"),
		// 										}},
		// 										NodeName: to.Ptr("worker3"),
		// 										PowerState: to.Ptr(armnetworkcloud.HybridAksClusterMachinePowerStateOn),
		// 									},
		// 									{
		// 										BareMetalMachineID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
		// 										ImageID: to.Ptr("myacr.azurecr.io/worker:latest"),
		// 										NetworkAttachments: []*armnetworkcloud.NetworkAttachment{
		// 											{
		// 												AttachedNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 												DefaultGateway: to.Ptr(armnetworkcloud.DefaultGatewayTrue),
		// 												IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
		// 												IPv4Address: to.Ptr("198.51.100.1"),
		// 												IPv6Address: to.Ptr("2001:0db8:0000:0000:0000:0000:0000:0001"),
		// 												MacAddress: to.Ptr("bf:1c:29:31:31:1f"),
		// 												NetworkAttachmentName: to.Ptr("netAttachName01"),
		// 										}},
		// 										NodeName: to.Ptr("worker4"),
		// 										PowerState: to.Ptr(armnetworkcloud.HybridAksClusterMachinePowerStateOn),
		// 									},
		// 									{
		// 										BareMetalMachineID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
		// 										ImageID: to.Ptr("myacr.azurecr.io/worker:latest"),
		// 										NetworkAttachments: []*armnetworkcloud.NetworkAttachment{
		// 											{
		// 												AttachedNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 												DefaultGateway: to.Ptr(armnetworkcloud.DefaultGatewayTrue),
		// 												IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
		// 												IPv4Address: to.Ptr("198.51.100.1"),
		// 												IPv6Address: to.Ptr("2001:0db8:0000:0000:0000:0000:0000:0001"),
		// 												MacAddress: to.Ptr("bf:1c:29:31:31:1f"),
		// 												NetworkAttachmentName: to.Ptr("netAttachName01"),
		// 										}},
		// 										NodeName: to.Ptr("worker5"),
		// 										PowerState: to.Ptr(armnetworkcloud.HybridAksClusterMachinePowerStateOn),
		// 								}},
		// 								VMCount: to.Ptr[int64](6),
		// 								VMSize: to.Ptr("NC_Gr_v1"),
		// 							},
		// 							{
		// 								AgentPoolID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.HybridContainerService/provisionedClusters/provisionedClustersName/agentPools/highMemory"),
		// 								AgentPoolName: to.Ptr("highMemory"),
		// 								CPUCores: to.Ptr[int64](32),
		// 								DiskSizeGB: to.Ptr[int64](128),
		// 								MemorySizeGB: to.Ptr[int64](128),
		// 								NodePoolName: to.Ptr("highMemory"),
		// 								Nodes: []*armnetworkcloud.Node{
		// 									{
		// 										BareMetalMachineID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
		// 										ImageID: to.Ptr("myacr.azurecr.io/worker:latest"),
		// 										NetworkAttachments: []*armnetworkcloud.NetworkAttachment{
		// 											{
		// 												AttachedNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 												DefaultGateway: to.Ptr(armnetworkcloud.DefaultGatewayTrue),
		// 												IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
		// 												IPv4Address: to.Ptr("198.51.100.1"),
		// 												IPv6Address: to.Ptr("2001:0db8:0000:0000:0000:0000:0000:0001"),
		// 												MacAddress: to.Ptr("bf:1c:29:31:31:1f"),
		// 												NetworkAttachmentName: to.Ptr("netAttachName01"),
		// 										}},
		// 										NodeName: to.Ptr("highmemworker0"),
		// 										PowerState: to.Ptr(armnetworkcloud.HybridAksClusterMachinePowerStateOn),
		// 									},
		// 									{
		// 										BareMetalMachineID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/bareMetalMachines/bareMetalMachineName"),
		// 										ImageID: to.Ptr("myacr.azurecr.io/worker:latest"),
		// 										NetworkAttachments: []*armnetworkcloud.NetworkAttachment{
		// 											{
		// 												AttachedNetworkID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.NetworkCloud/l3Networks/l3NetworkName"),
		// 												DefaultGateway: to.Ptr(armnetworkcloud.DefaultGatewayTrue),
		// 												IPAllocationMethod: to.Ptr(armnetworkcloud.VirtualMachineIPAllocationMethodDynamic),
		// 												IPv4Address: to.Ptr("198.51.100.1"),
		// 												IPv6Address: to.Ptr("2001:0db8:0000:0000:0000:0000:0000:0001"),
		// 												MacAddress: to.Ptr("bf:1c:29:31:31:1f"),
		// 												NetworkAttachmentName: to.Ptr("netAttachName01"),
		// 										}},
		// 										NodeName: to.Ptr("highmemworker1"),
		// 										PowerState: to.Ptr(armnetworkcloud.HybridAksClusterMachinePowerStateOn),
		// 								}},
		// 								VMCount: to.Ptr[int64](2),
		// 								VMSize: to.Ptr("NC_Gh_v1"),
		// 						}},
		// 					},
		// 			}},
		// 		}
	}
}
