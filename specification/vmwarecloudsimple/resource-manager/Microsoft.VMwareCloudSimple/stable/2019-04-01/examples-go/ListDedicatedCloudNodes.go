package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListDedicatedCloudNodes.json
func ExampleDedicatedCloudNodesClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDedicatedCloudNodesClient().NewListBySubscriptionPager(&armvmwarecloudsimple.DedicatedCloudNodesClientListBySubscriptionOptions{Filter: nil,
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
		// 			Name: to.Ptr("node-2"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/dedicatedCloudNodes"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup-1/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/node-2"),
		// 			Location: to.Ptr("westus2"),
		// 			Properties: &armvmwarecloudsimple.DedicatedCloudNodeProperties{
		// 				AvailabilityZoneID: to.Ptr("az1"),
		// 				AvailabilityZoneName: to.Ptr("Availability Zone 1"),
		// 				CloudRackName: to.Ptr("cloud_rack_1"),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-21T07:32:12.134Z"); return t}()),
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
		// 			Name: to.Ptr("node1"),
		// 			Type: to.Ptr("Microsoft.VMwareCloudSimple/dedicatedCloudNodes"),
		// 			ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup-2/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/node-1"),
		// 			Location: to.Ptr("westus2"),
		// 			Properties: &armvmwarecloudsimple.DedicatedCloudNodeProperties{
		// 				AvailabilityZoneID: to.Ptr("az1"),
		// 				AvailabilityZoneName: to.Ptr("Availability Zone 1"),
		// 				Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-22T11:57:43.891Z"); return t}()),
		// 				NodesCount: to.Ptr[int32](0),
		// 				PlacementGroupID: to.Ptr("n1"),
		// 				PlacementGroupName: to.Ptr("Placement Group 1"),
		// 				PurchaseID: to.Ptr("eb9d2e22-c787-4723-aa4f-3760b53a0a4d"),
		// 				SKUDescription: &armvmwarecloudsimple.SKUDescription{
		// 					Name: to.Ptr("CS28-Node"),
		// 					ID: to.Ptr("general"),
		// 				},
		// 				Status: to.Ptr(armvmwarecloudsimple.NodeStatusUnused),
		// 			},
		// 			SKU: &armvmwarecloudsimple.SKU{
		// 				Name: to.Ptr("VMware_CloudSimple_CS28"),
		// 			},
		// 	}},
		// }
	}
}
