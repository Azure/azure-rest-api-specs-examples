package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/ContainerListAllInDevice.json
func ExampleContainersClient_NewListByStorageAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewContainersClient().NewListByStorageAccountPager("testedgedevice", "storageaccount1", "GroupForEdgeAutomation", nil)
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
		// page.ContainerList = armdataboxedge.ContainerList{
		// 	Value: []*armdataboxedge.Container{
		// 		{
		// 			Name: to.Ptr("blobcontainer1"),
		// 			Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/storageAccounts/containers"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForDataBoxEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/EAuToWIN-6LJVNIFFB411021190437191/storageAccounts/storageaccount1/containers/blobcontainer1"),
		// 			Properties: &armdataboxedge.ContainerProperties{
		// 				ContainerStatus: to.Ptr(armdataboxedge.ContainerStatusOK),
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-20T23:14:32.301Z"); return t}()),
		// 				DataFormat: to.Ptr(armdataboxedge.AzureContainerDataFormatBlockBlob),
		// 				RefreshDetails: &armdataboxedge.RefreshDetails{
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("blobcontainer2"),
		// 			Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/storageAccounts/containers"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForDataBoxEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/EAuToWIN-6LJVNIFFB411021190437191/storageAccounts/storageaccount1/containers/blobcontainer2"),
		// 			Properties: &armdataboxedge.ContainerProperties{
		// 				ContainerStatus: to.Ptr(armdataboxedge.ContainerStatusOK),
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-20T23:15:38.007Z"); return t}()),
		// 				DataFormat: to.Ptr(armdataboxedge.AzureContainerDataFormatBlockBlob),
		// 				RefreshDetails: &armdataboxedge.RefreshDetails{
		// 				},
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("blobcontainer3"),
		// 			Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/storageAccounts/containers"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForDataBoxEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/EAuToWIN-6LJVNIFFB411021190437191/storageAccounts/storageaccount1/containers/blobcontainer3"),
		// 			Properties: &armdataboxedge.ContainerProperties{
		// 				ContainerStatus: to.Ptr(armdataboxedge.ContainerStatusOK),
		// 				CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-20T23:13:27.854Z"); return t}()),
		// 				DataFormat: to.Ptr(armdataboxedge.AzureContainerDataFormatBlockBlob),
		// 				RefreshDetails: &armdataboxedge.RefreshDetails{
		// 				},
		// 			},
		// 	}},
		// }
	}
}
