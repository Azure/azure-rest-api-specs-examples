package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/RegisteredServers_ListByStorageSyncService.json
func ExampleRegisteredServersClient_NewListByStorageSyncServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRegisteredServersClient().NewListByStorageSyncServicePager("SampleResourceGroup_1", "SampleStorageSyncService_1", nil)
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
		// page.RegisteredServerArray = armstoragesync.RegisteredServerArray{
		// 	Value: []*armstoragesync.RegisteredServer{
		// 		{
		// 			Name: to.Ptr("SampleRegisteredServer_1.redmond.corp.microsoft.com"),
		// 			Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/registeredServers"),
		// 			ID: to.Ptr("/subscriptions/3a048283-338f-4002-a9dd-a50fdadcb392/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/080d4133-bdb5-40a0-96a0-71a6057bfe9a"),
		// 			Properties: &armstoragesync.RegisteredServerProperties{
		// 				AgentVersion: to.Ptr("3.1.5.0"),
		// 				AgentVersionExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T21:50:40.584Z"); return t}()),
		// 				AgentVersionStatus: to.Ptr(armstoragesync.RegisteredServerAgentVersionStatusOk),
		// 				ClusterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				ClusterName: to.Ptr(""),
		// 				DiscoveryEndpointURI: to.Ptr("https://kailanitest99.one.microsoft.com:443"),
		// 				FriendlyName: to.Ptr("afscv-2304-139"),
		// 				LastHeartBeat: to.Ptr("2018-06-11T21:50:40.5840543Z"),
		// 				LastOperationName: to.Ptr("ICreateRegisteredServerWorkflow"),
		// 				LastWorkflowID: to.Ptr("storageSyncServices/CV_FileStore_F1D485AA/workflows/4eecfbcf-9537-4b61-8fee-aaa3ace11c44"),
		// 				ManagementEndpointURI: to.Ptr("https://kailanitest99.one.microsoft.com:443/"),
		// 				MonitoringEndpointURI: to.Ptr("https://kailanitest99.one.microsoft.com:443/"),
		// 				ProvisioningState: to.Ptr("Succeeded"),
		// 				ResourceLocation: to.Ptr("westus"),
		// 				ServerID: to.Ptr("3635ca8b-5cc8-4a5c-bd43-c2de5ad8dc64"),
		// 				ServerManagementErrorCode: to.Ptr[int32](0),
		// 				ServerName: to.Ptr("afscv-2304-139"),
		// 				ServerOSVersion: to.Ptr("10.0.14393.0"),
		// 				ServerRole: to.Ptr("Standalone"),
		// 				ServiceLocation: to.Ptr("westus"),
		// 				StorageSyncServiceUID: to.Ptr("4aa14534-1c61-483b-b6a6-9630a76f109a"),
		// 			},
		// 	}},
		// }
	}
}
