package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/NodeGetAllInDevice.json
func ExampleNodesClient_NewListByDataBoxEdgeDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNodesClient().NewListByDataBoxEdgeDevicePager("testedgedevice", "GroupForEdgeAutomation", nil)
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
		// page.NodeList = armdataboxedge.NodeList{
		// 	Value: []*armdataboxedge.Node{
		// 		{
		// 			Name: to.Ptr("1D6QHQ2"),
		// 			Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/nodes"),
		// 			ID: to.Ptr("/subscriptions/db4e2fdb-6d80-4e6e-b7cd-736098270664/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/DataBoxEdgeDevices/testedgedevice/nodes/1D6QHQ2"),
		// 			Properties: &armdataboxedge.NodeProperties{
		// 				NodeChassisSerialNumber: to.Ptr("1D6QHQ2"),
		// 				NodeDisplayName: to.Ptr("Controller 1"),
		// 				NodeFriendlySoftwareVersion: to.Ptr("Data Box Edge 1908"),
		// 				NodeHcsVersion: to.Ptr("1.6.961.8307"),
		// 				NodeInstanceID: to.Ptr("ad051874-7d27-4ac4-a7b1-b6f231d8a035"),
		// 				NodeSerialNumber: to.Ptr("1D6QHQ2"),
		// 				NodeStatus: to.Ptr(armdataboxedge.NodeStatusUnknown),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("1DCNHQ2"),
		// 			Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/nodes"),
		// 			ID: to.Ptr("/subscriptions/db4e2fdb-6d80-4e6e-b7cd-736098270664/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/DataBoxEdgeDevices/testedgedevice/nodes/1DCNHQ2"),
		// 			Properties: &armdataboxedge.NodeProperties{
		// 				NodeChassisSerialNumber: to.Ptr("1D6QHQ2"),
		// 				NodeDisplayName: to.Ptr("Controller 1"),
		// 				NodeFriendlySoftwareVersion: to.Ptr("Data Box Edge 1908"),
		// 				NodeHcsVersion: to.Ptr("1.6.961.8307"),
		// 				NodeInstanceID: to.Ptr("ddf3a76d-ada4-44af-99c6-a69a0e21bcb6"),
		// 				NodeSerialNumber: to.Ptr("1DCNHQ2"),
		// 				NodeStatus: to.Ptr(armdataboxedge.NodeStatusUnknown),
		// 			},
		// 	}},
		// }
	}
}
