package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListRGDedicatedCloudNodes.json
func ExampleDedicatedCloudNodesClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedCloudNodesClient().NewListByResourceGroupPager("myResourceGroup", &armvmwarecloudsimple.DedicatedCloudNodesClientListByResourceGroupOptions{Filter: nil,
		Top:       nil,
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
		// page.DedicatedCloudNodeListResponse = armvmwarecloudsimple.DedicatedCloudNodeListResponse{
		// 	Value: []*armvmwarecloudsimple.DedicatedCloudNode{
		// 		{
		// 			Name: to.Ptr("node-east-1"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/dedicatedCloudNodes"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourcegroups/myResourceGroup/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/node-east-1"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armvmwarecloudsimple.DedicatedCloudNodeProperties{
		// 				AvailabilityZoneID: to.Ptr("az1"),
		// 				AvailabilityZoneName: to.Ptr("Availability Zone 1"),
		// 				CloudRackName: to.Ptr("cloud_rack_1"),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-21T07:32:12.134Z"); return t}()),
		// 				NodesCount: to.Ptr[int32](0),
		// 				PlacementGroupID: to.Ptr("n2"),
		// 				PlacementGroupName: to.Ptr("Placement Group 2"),
		// 				PrivateCloudID: to.Ptr("private_cloud_id"),
		// 				PrivateCloudName: to.Ptr("private_cloud_name"),
		// 				PurchaseID: to.Ptr("225fadec-3bbe-4e61-a119-ff102da67d0d"),
		// 				SKUDescription: &armvmwarecloudsimple.SKUDescription{
		// 					Name: to.Ptr("CS28-Node"),
		// 					ID: to.Ptr("general"),
		// 				},
		// 				Status: to.Ptr(armvmwarecloudsimple.NodeStatusUsed),
		// 				VmwareClusterName: to.Ptr("Cluster"),
		// 			},
		// 			SKU: &armvmwarecloudsimple.SKU{
		// 				Name: to.Ptr("VMware_CloudSimple_CS28"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("node-east-2"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/dedicatedCloudNodes"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourcegroups/myResourceGroup/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/node-east-2"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armvmwarecloudsimple.DedicatedCloudNodeProperties{
		// 				AvailabilityZoneID: to.Ptr("az1"),
		// 				AvailabilityZoneName: to.Ptr("Availability Zone 1"),
		// 				CloudRackName: to.Ptr("cloud_rack_1"),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-21T07:32:12.114Z"); return t}()),
		// 				NodesCount: to.Ptr[int32](0),
		// 				PlacementGroupID: to.Ptr("n2"),
		// 				PlacementGroupName: to.Ptr("Placement Group 2"),
		// 				PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
		// 				PrivateCloudName: to.Ptr("private_cloud_name"),
		// 				PurchaseID: to.Ptr("225fadec-3bbe-4e61-a119-ff102da67d0d"),
		// 				SKUDescription: &armvmwarecloudsimple.SKUDescription{
		// 					Name: to.Ptr("CS28-Node"),
		// 					ID: to.Ptr("general"),
		// 				},
		// 				Status: to.Ptr(armvmwarecloudsimple.NodeStatusUsed),
		// 				VmwareClusterName: to.Ptr("Cluster"),
		// 			},
		// 			SKU: &armvmwarecloudsimple.SKU{
		// 				Name: to.Ptr("VMware_CloudSimple_CS28"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("myNode"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/dedicatedCloudNodes"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourcegroups/myResourceGroup/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/myNode"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armvmwarecloudsimple.DedicatedCloudNodeProperties{
		// 				AvailabilityZoneID: to.Ptr("az1"),
		// 				AvailabilityZoneName: to.Ptr("Availability Zone 1"),
		// 				CloudRackName: to.Ptr("cloud_rack_1"),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-08T09:35:47.961Z"); return t}()),
		// 				NodesCount: to.Ptr[int32](0),
		// 				PlacementGroupID: to.Ptr("n2"),
		// 				PlacementGroupName: to.Ptr("Placement Group 2"),
		// 				PrivateCloudID: to.Ptr("/subscriptions/{subscription-id}/providers/Microsoft.VMwareCloudSimple/locations/westus2/privateClouds/myPrivateCloud"),
		// 				PrivateCloudName: to.Ptr("myPrivateCloud"),
		// 				PurchaseID: to.Ptr("b3fcd958-f19c-4421-ab46-e4fa9cc8514e"),
		// 				SKUDescription: &armvmwarecloudsimple.SKUDescription{
		// 					Name: to.Ptr("CS28-Node"),
		// 					ID: to.Ptr("general"),
		// 				},
		// 				Status: to.Ptr(armvmwarecloudsimple.NodeStatusUsed),
		// 				VmwareClusterName: to.Ptr("Cluster"),
		// 			},
		// 			SKU: &armvmwarecloudsimple.SKU{
		// 				Name: to.Ptr("VMware_CloudSimple_CS28"),
		// 			},
		// 	}},
		// }
	}
}
