package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/StorageAccountGetAllInDevice.json
func ExampleStorageAccountsClient_NewListByDataBoxEdgeDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewStorageAccountsClient().NewListByDataBoxEdgeDevicePager("testedgedevice", "GroupForEdgeAutomation", nil)
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
		// page.StorageAccountList = armdataboxedge.StorageAccountList{
		// 	Value: []*armdataboxedge.StorageAccount{
		// 		{
		// 			Name: to.Ptr("blobstorageaccount1"),
		// 			Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/storageAccounts"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForDataBoxEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccounts/blobstorageaccount1"),
		// 			Properties: &armdataboxedge.StorageAccountProperties{
		// 				Description: to.Ptr("It's an awesome storage account"),
		// 				BlobEndpoint: to.Ptr("https://blobstorageaccount1.blob.testedge.microsoftdatabox.com/"),
		// 				ContainerCount: to.Ptr[int32](0),
		// 				DataPolicy: to.Ptr(armdataboxedge.DataPolicyCloud),
		// 				StorageAccountCredentialID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForDataBoxEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccountCredentials/cisbvt"),
		// 				StorageAccountStatus: to.Ptr(armdataboxedge.StorageAccountStatusOK),
		// 			},
		// 	}},
		// }
	}
}
