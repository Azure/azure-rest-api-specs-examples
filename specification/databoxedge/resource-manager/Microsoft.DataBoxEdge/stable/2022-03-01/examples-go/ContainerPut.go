package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/ContainerPut.json
func ExampleContainersClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewContainersClient().BeginCreateOrUpdate(ctx, "testedgedevice", "storageaccount1", "blobcontainer1", "GroupForEdgeAutomation", armdataboxedge.Container{
		Properties: &armdataboxedge.ContainerProperties{
			DataFormat: to.Ptr(armdataboxedge.AzureContainerDataFormatBlockBlob),
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
	// res.Container = armdataboxedge.Container{
	// 	Name: to.Ptr("blobcontainer-5e155efe"),
	// 	Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/storageAccounts/containers"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForDataBoxEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccounts/storageaccount1/containers/blobcontainer1"),
	// 	Properties: &armdataboxedge.ContainerProperties{
	// 		ContainerStatus: to.Ptr(armdataboxedge.ContainerStatusOK),
	// 		CreatedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-10-20T23:13:27.8545799Z"); return t}()),
	// 		DataFormat: to.Ptr(armdataboxedge.AzureContainerDataFormatBlockBlob),
	// 		RefreshDetails: &armdataboxedge.RefreshDetails{
	// 		},
	// 	},
	// }
}
