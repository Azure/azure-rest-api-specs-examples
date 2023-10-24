package armredhatopenshift_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bce66ff64f0e9edc9ea6119d00324058413e81ed/specification/redhatopenshift/resource-manager/Microsoft.RedHatOpenShift/stable/2023-09-04/examples/OpenShiftClusters_Get.json
func ExampleOpenShiftClustersClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armredhatopenshift.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOpenShiftClustersClient().Get(ctx, "resourceGroup", "resourceName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.OpenShiftCluster = armredhatopenshift.OpenShiftCluster{
	// 	Name: to.Ptr("resourceName"),
	// 	Type: to.Ptr("Microsoft.RedHatOpenShift/OpenShiftClusters"),
	// 	ID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/Microsoft.RedHatOpenShift/OpenShiftClusters/resourceName"),
	// 	SystemData: &armredhatopenshift.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.1075056Z"); return t}()),
	// 		CreatedBy: to.Ptr("string"),
	// 		CreatedByType: to.Ptr(armredhatopenshift.CreatedByTypeApplication),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-02-03T01:01:01.1075056Z"); return t}()),
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
	// 		WorkerProfilesStatus: []*armredhatopenshift.WorkerProfile{
	// 			{
	// 				Name: to.Ptr("worker1"),
	// 				Count: to.Ptr[int32](1),
	// 				DiskSizeGB: to.Ptr[int32](128),
	// 				SubnetID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker"),
	// 				VMSize: to.Ptr("Standard_D2s_v3"),
	// 			},
	// 			{
	// 				Name: to.Ptr("worker2"),
	// 				Count: to.Ptr[int32](1),
	// 				DiskSizeGB: to.Ptr[int32](128),
	// 				SubnetID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker"),
	// 				VMSize: to.Ptr("Standard_D2s_v3"),
	// 			},
	// 			{
	// 				Name: to.Ptr("worker3"),
	// 				Count: to.Ptr[int32](1),
	// 				DiskSizeGB: to.Ptr[int32](128),
	// 				SubnetID: to.Ptr("/subscriptions/subscriptionId/resourceGroups/vnetResourceGroup/providers/Microsoft.Network/virtualNetworks/vnet/subnets/worker"),
	// 				VMSize: to.Ptr("Standard_D2s_v3"),
	// 		}},
	// 	},
	// }
}
