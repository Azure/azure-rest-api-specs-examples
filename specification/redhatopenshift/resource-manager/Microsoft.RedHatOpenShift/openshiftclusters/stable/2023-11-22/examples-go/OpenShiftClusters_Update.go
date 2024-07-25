package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c1cea38fb7e5cec9afe223a2ed15cbe2fbeecbdb/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/openshiftclusters/stable/2023-11-22/examples/OpenShiftClusters_Update.json
func ExampleOpenShiftClustersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshift.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewOpenShiftClustersClient().BeginUpdate(ctx, "resourceGroup", "resourceName", armredhatopenshift.OpenShiftClusterUpdate{
		Properties: &armredhatopenshift.OpenShiftClusterProperties{
			ApiserverProfile: &armredhatopenshift.APIServerProfile{
				Visibility: to.Ptr(armredhatopenshift.VisibilityPublic),
			},
			ClusterProfile: &armredhatopenshift.ClusterProfile{
				Domain:               to.Ptr("cluster.location.aroapp.io"),
				FipsValidatedModules: to.Ptr(armredhatopenshift.FipsValidatedModulesEnabled),
				PullSecret:           to.Ptr("{\"auths\":{\"registry.connect.redhat.com\":{\"auth\":\"\"},\"registry.redhat.io\":{\"auth\":\"\"}}}"),
				ResourceGroupID:      to.Ptr("/subscriptions/subscriptionId/resourceGroups/clusterResourceGroup"),
			},
			ConsoleProfile: &armredhatopenshift.ConsoleProfile{},
			IngressProfiles: []*armredhatopenshift.IngressProfile{
				{
					Name:       to.Ptr("default"),
					Visibility: to.Ptr(armredhatopenshift.VisibilityPublic),
				}},
			MasterProfile: &armredhatopenshift.MasterProfile{
				EncryptionAtHost: to.Ptr(armredhatopenshift.EncryptionAtHostEnabled),
				SubnetID:         to.Ptr("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/master"),
				VMSize:           to.Ptr("Standard_D8s_v3"),
			},
			NetworkProfile: &armredhatopenshift.NetworkProfile{
				LoadBalancerProfile: &armredhatopenshift.LoadBalancerProfile{
					ManagedOutboundIPs: &armredhatopenshift.ManagedOutboundIPs{
						Count: to.Ptr[int32](1),
					},
				},
				PodCidr:          to.Ptr("10.128.0.0/14"),
				PreconfiguredNSG: to.Ptr(armredhatopenshift.PreconfiguredNSGDisabled),
				ServiceCidr:      to.Ptr("172.30.0.0/16"),
			},
			ServicePrincipalProfile: &armredhatopenshift.ServicePrincipalProfile{
				ClientID:     to.Ptr("clientId"),
				ClientSecret: to.Ptr("clientSecret"),
			},
			WorkerProfiles: []*armredhatopenshift.WorkerProfile{
				{
					Name:       to.Ptr("worker"),
					Count:      to.Ptr[int32](3),
					DiskSizeGB: to.Ptr[int32](128),
					SubnetID:   to.Ptr("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker"),
					VMSize:     to.Ptr("Standard_D2s_v3"),
				}},
		},
		Tags: map[string]*string{
			"key": to.Ptr("value"),
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
	// res.OpenShiftCluster = armredhatopenshift.OpenShiftCluster{
	// 	Name: to.Ptr("resourceName"),
	// 	Type: to.Ptr("Microsoft.RedHatOpenShift/OpenShiftClusters"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/Microsoft.RedHatOpenShift/OpenShiftClusters/resourceName"),
	// 	SystemData: &armredhatopenshift.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armredhatopenshift.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.107Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("string"),
	// 		LastModifiedByType: to.Ptr(armredhatopenshift.CreatedByTypeApplication),
	// 	},
	// 	Location: to.Ptr("location"),
	// 	Tags: map[string]*string{
	// 		"key": to.Ptr("value"),
	// 	},
	// 	Properties: &armredhatopenshift.OpenShiftClusterProperties{
	// 		ApiserverProfile: &armredhatopenshift.APIServerProfile{
	// 			IP: to.Ptr("1.2.3.4"),
	// 			URL: to.Ptr("https://api.cluster.location.aroapp.io:6443/"),
	// 			Visibility: to.Ptr(armredhatopenshift.VisibilityPublic),
	// 		},
	// 		ClusterProfile: &armredhatopenshift.ClusterProfile{
	// 			Domain: to.Ptr("cluster.location.aroapp.io"),
	// 			ResourceGroupID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/clusterResourceGroup"),
	// 			Version: to.Ptr("4.11.0"),
	// 		},
	// 		ConsoleProfile: &armredhatopenshift.ConsoleProfile{
	// 			URL: to.Ptr("https://console-openshift-console.apps.cluster.location.aroapp.io/"),
	// 		},
	// 		IngressProfiles: []*armredhatopenshift.IngressProfile{
	// 			{
	// 				Name: to.Ptr("default"),
	// 				IP: to.Ptr("1.2.3.4"),
	// 				Visibility: to.Ptr(armredhatopenshift.VisibilityPublic),
	// 		}},
	// 		MasterProfile: &armredhatopenshift.MasterProfile{
	// 			SubnetID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/master"),
	// 			VMSize: to.Ptr("Standard_D8s_v3"),
	// 		},
	// 		NetworkProfile: &armredhatopenshift.NetworkProfile{
	// 			PodCidr: to.Ptr("10.128.0.0/14"),
	// 			PreconfiguredNSG: to.Ptr(armredhatopenshift.PreconfiguredNSGDisabled),
	// 			ServiceCidr: to.Ptr("172.30.0.0/16"),
	// 		},
	// 		ProvisioningState: to.Ptr(armredhatopenshift.ProvisioningStateSucceeded),
	// 		ServicePrincipalProfile: &armredhatopenshift.ServicePrincipalProfile{
	// 			ClientID: to.Ptr("clientId"),
	// 		},
	// 		WorkerProfiles: []*armredhatopenshift.WorkerProfile{
	// 			{
	// 				Name: to.Ptr("worker"),
	// 				Count: to.Ptr[int32](3),
	// 				DiskSizeGB: to.Ptr[int32](128),
	// 				SubnetID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker"),
	// 				VMSize: to.Ptr("Standard_D2s_v3"),
	// 		}},
	// 	},
	// }
}
