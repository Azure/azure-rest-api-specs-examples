package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/Workflows_ListByStorageSyncService.json
func ExampleWorkflowsClient_NewListByStorageSyncServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkflowsClient().NewListByStorageSyncServicePager("SampleResourceGroup_1", "SampleStorageSyncService_1", nil)
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
		// page.WorkflowArray = armstoragesync.WorkflowArray{
		// 	Value: []*armstoragesync.Workflow{
		// 		{
		// 			Name: to.Ptr("828219ea-083e-48b5-89ea-8fd9991b2e75"),
		// 			Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/workflows"),
		// 			ID: to.Ptr("/subscriptions/3a048283-338f-4002-a9dd-a50fdadcb392/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/workflows/828219ea-083e-48b5-89ea-8fd9991b2e75"),
		// 			Properties: &armstoragesync.WorkflowProperties{
		// 				LastOperationID: to.Ptr("\"fe680c98-5725-49c8-b0dc-5e29745f752b\""),
		// 				LastStepName: to.Ptr("runServerJob"),
		// 				Operation: to.Ptr(armstoragesync.OperationDirectionDo),
		// 				Status: to.Ptr(armstoragesync.WorkflowStatusSucceeded),
		// 				Steps: to.Ptr("[{\"name\":\"validateInput\",\"friendlyName\":\"validateInput\",\"status\":\"Succeeded\",\"error\":null},{\"name\":\"newServerEndpoint\",\"friendlyName\":\"newServerEndpoint\",\"status\":\"Succeeded\",\"error\":null},{\"name\":\"updateReplicaGroupCertificates\",\"friendlyName\":\"updateReplicaGroupCertificates\",\"status\":\"Succeeded\",\"error\":null},{\"name\":\"runServerJob\",\"friendlyName\":\"runServerJob\",\"status\":\"Succeeded\",\"error\":null}]"),
		// 			},
		// 	}},
		// }
	}
}
