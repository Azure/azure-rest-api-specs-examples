package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/6e34caed36815fc876c8e8c0371db76f809e52e8/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/OpenShiftClusters/stable/2025-07-25/examples/OpenShiftClusters_List.json
func ExampleOpenShiftClustersClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshift.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOpenShiftClustersClient().NewListPager(nil)
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
		// page.OpenShiftClusterList = armredhatopenshift.OpenShiftClusterList{
		// 	Value: []*armredhatopenshift.OpenShiftCluster{
		// 		{
		// 			Name: to.Ptr("resourceName"),
		// 			Type: to.Ptr("Microsoft.RedHatOpenShift/OpenShiftClusters"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup/providers/Microsoft.RedHatOpenShift/OpenShiftClusters/resourceName"),
		// 			SystemData: &armredhatopenshift.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armredhatopenshift.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armredhatopenshift.CreatedByTypeApplication),
		// 			},
		// 			Location: to.Ptr("location"),
		// 			Tags: map[string]*string{
		// 				"key": to.Ptr("value"),
		// 			},
		// 			Identity: &armredhatopenshift.ManagedServiceIdentity{
		// 				Type: to.Ptr(armredhatopenshift.ManagedServiceIdentityTypeUserAssigned),
		// 				UserAssignedIdentities: map[string]*armredhatopenshift.UserAssignedIdentity{
		// 					"": &armredhatopenshift.UserAssignedIdentity{
		// 					},
		// 				},
		// 			},
		// 			Properties: &armredhatopenshift.OpenShiftClusterProperties{
		// 				ApiserverProfile: &armredhatopenshift.APIServerProfile{
		// 					IP: to.Ptr("1.2.3.4"),
		// 					URL: to.Ptr("https://api.cluster.location.aroapp.io:6443/"),
		// 					Visibility: to.Ptr(armredhatopenshift.VisibilityPublic),
		// 				},
		// 				ClusterProfile: &armredhatopenshift.ClusterProfile{
		// 					Domain: to.Ptr("cluster.location.aroapp.io"),
		// 					ResourceGroupID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/clusterResourceGroup"),
		// 					Version: to.Ptr("4.11.0"),
		// 				},
		// 				ConsoleProfile: &armredhatopenshift.ConsoleProfile{
		// 					URL: to.Ptr("https://console-openshift-console.apps.cluster.location.aroapp.io/"),
		// 				},
		// 				IngressProfiles: []*armredhatopenshift.IngressProfile{
		// 					{
		// 						Name: to.Ptr("default"),
		// 						IP: to.Ptr("1.2.3.4"),
		// 						Visibility: to.Ptr(armredhatopenshift.VisibilityPublic),
		// 				}},
		// 				MasterProfile: &armredhatopenshift.MasterProfile{
		// 					SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/master"),
		// 					VMSize: to.Ptr("Standard_D8s_v3"),
		// 				},
		// 				NetworkProfile: &armredhatopenshift.NetworkProfile{
		// 					LoadBalancerProfile: &armredhatopenshift.LoadBalancerProfile{
		// 						EffectiveOutboundIPs: []*armredhatopenshift.EffectiveOutboundIP{
		// 							{
		// 								ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/clusterResourceGroup/providers/Microsoft.Network/publicIPAddresses/publicIPAddressName"),
		// 						}},
		// 						ManagedOutboundIPs: &armredhatopenshift.ManagedOutboundIPs{
		// 							Count: to.Ptr[int32](1),
		// 						},
		// 					},
		// 					PodCidr: to.Ptr("10.128.0.0/14"),
		// 					PreconfiguredNSG: to.Ptr(armredhatopenshift.PreconfiguredNSGDisabled),
		// 					ServiceCidr: to.Ptr("172.30.0.0/16"),
		// 				},
		// 				PlatformWorkloadIdentityProfile: &armredhatopenshift.PlatformWorkloadIdentityProfile{
		// 					PlatformWorkloadIdentities: map[string]*armredhatopenshift.PlatformWorkloadIdentity{
		// 						"": &armredhatopenshift.PlatformWorkloadIdentity{
		// 						},
		// 					},
		// 				},
		// 				ProvisioningState: to.Ptr(armredhatopenshift.ProvisioningStateSucceeded),
		// 				ServicePrincipalProfile: &armredhatopenshift.ServicePrincipalProfile{
		// 					ClientID: to.Ptr("clientId"),
		// 				},
		// 				WorkerProfiles: []*armredhatopenshift.WorkerProfile{
		// 					{
		// 						Name: to.Ptr("worker"),
		// 						Count: to.Ptr[int32](3),
		// 						DiskSizeGB: to.Ptr[int32](128),
		// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker"),
		// 						VMSize: to.Ptr("Standard_D2s_v3"),
		// 				}},
		// 				WorkerProfilesStatus: []*armredhatopenshift.WorkerProfile{
		// 					{
		// 						Name: to.Ptr("worker1"),
		// 						Count: to.Ptr[int32](1),
		// 						DiskSizeGB: to.Ptr[int32](128),
		// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker"),
		// 						VMSize: to.Ptr("Standard_D2s_v3"),
		// 					},
		// 					{
		// 						Name: to.Ptr("worker2"),
		// 						Count: to.Ptr[int32](1),
		// 						DiskSizeGB: to.Ptr[int32](128),
		// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker"),
		// 						VMSize: to.Ptr("Standard_D2s_v3"),
		// 					},
		// 					{
		// 						Name: to.Ptr("worker3"),
		// 						Count: to.Ptr[int32](1),
		// 						DiskSizeGB: to.Ptr[int32](128),
		// 						SubnetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker"),
		// 						VMSize: to.Ptr("Standard_D2s_v3"),
		// 				}},
		// 			},
		// 	}},
		// }
	}
}
