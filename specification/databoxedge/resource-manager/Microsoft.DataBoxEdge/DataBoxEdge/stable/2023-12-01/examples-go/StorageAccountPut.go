package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge/v2"
)

// Generated from example definition: 2023-12-01/StorageAccountPut.json
func ExampleStorageAccountsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("4385cf00-2d3a-425a-832f-f4285b1c9dce", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewStorageAccountsClient().BeginCreateOrUpdate(ctx, "testedgedevice", "blobstorageaccount1", "GroupForEdgeAutomation", armdataboxedge.StorageAccount{
		Properties: &armdataboxedge.StorageAccountProperties{
			Description:                to.Ptr("It's an awesome storage account"),
			DataPolicy:                 to.Ptr(armdataboxedge.DataPolicyCloud),
			StorageAccountCredentialID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForDataBoxEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccountCredentials/cisbvt"),
			StorageAccountStatus:       to.Ptr(armdataboxedge.StorageAccountStatusOK),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to poll the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armdataboxedge.StorageAccountsClientCreateOrUpdateResponse{
	// 	StorageAccount: armdataboxedge.StorageAccount{
	// 		Name: to.Ptr("blobstorageaccount1"),
	// 		Type: to.Ptr("Microsoft.DataBoxEdge/dataBoxEdgeDevices/storageAccounts"),
	// 		ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForDataBoxEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccounts/blobstorageaccount1"),
	// 		Properties: &armdataboxedge.StorageAccountProperties{
	// 			Description: to.Ptr("It's an awesome storage account"),
	// 			BlobEndpoint: to.Ptr("https://blobstorageaccount1.blob.testedge.microsoftdatabox.com/"),
	// 			ContainerCount: to.Ptr[int32](0),
	// 			DataPolicy: to.Ptr(armdataboxedge.DataPolicyCloud),
	// 			StorageAccountCredentialID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForDataBoxEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccountCredentials/cisbvt"),
	// 			StorageAccountStatus: to.Ptr(armdataboxedge.StorageAccountStatusOK),
	// 		},
	// 	},
	// }
}
