package armvmwarecloudsimple_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/vmwarecloudsimple/armvmwarecloudsimple"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/PatchDedicatedCloudNode.json
func ExampleDedicatedCloudNodesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armvmwarecloudsimple.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewDedicatedCloudNodesClient().Update(ctx, "myResourceGroup", "myNode", armvmwarecloudsimple.PatchPayload{
		Tags: map[string]*string{
			"myTag": to.Ptr("tagValue"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.DedicatedCloudNode = armvmwarecloudsimple.DedicatedCloudNode{
	// 	Name: to.Ptr("myNode"),
	// 	Type: to.Ptr("Microsoft.VMwareCloudSimple/dedicatedCloudNodes"),
	// 	ID: to.Ptr("/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.VMwareCloudSimple/dedicatedCloudNodes/myNode"),
	// 	Location: to.Ptr("westus"),
	// 	Properties: &armvmwarecloudsimple.DedicatedCloudNodeProperties{
	// 		AvailabilityZoneID: to.Ptr("az1"),
	// 		Created: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-21T13:43:26.441Z"); return t}()),
	// 		NodesCount: to.Ptr[int32](0),
	// 		PlacementGroupID: to.Ptr("n1"),
	// 		PurchaseID: to.Ptr("56acbd46-3d36-4bbf-9b08-57c30fdf6932"),
	// 		SKUDescription: &armvmwarecloudsimple.SKUDescription{
	// 			Name: to.Ptr("CS28-Node"),
	// 			ID: to.Ptr("general"),
	// 		},
	// 		Status: to.Ptr(armvmwarecloudsimple.NodeStatusUnused),
	// 	},
	// 	Tags: map[string]*string{
	// 		"myTag": to.Ptr("tagValue"),
	// 	},
	// }
}
