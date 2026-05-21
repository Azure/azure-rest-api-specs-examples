package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync/v2"
)

// Generated from example definition: 2022-09-01/RegisteredServers_Update.json
func ExampleRegisteredServersClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("52b8da2f-61e0-4a1f-8dde-336911f367fb", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRegisteredServersClient().BeginUpdate(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", "080d4133-bdb5-40a0-96a0-71a6057bfe9a", armstoragesync.RegisteredServerUpdateParameters{
		Properties: &armstoragesync.RegisteredServerUpdateProperties{
			ApplicationID: to.Ptr("120d4132-bcd5-40a0-96a0-71a6057ebf0c"),
			Identity:      to.Ptr(true),
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
	// res = armstoragesync.RegisteredServersClientUpdateResponse{
	// 	RegisteredServer: armstoragesync.RegisteredServer{
	// 		Name: to.Ptr("ankushb-prod3.redmond.corp.microsoft.com"),
	// 		Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/registeredServers"),
	// 		ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/530a0384-50ac-456d-8240-9d6621404151"),
	// 		Properties: &armstoragesync.RegisteredServerProperties{
	// 			ActiveAuthType: to.Ptr(armstoragesync.ServerAuthTypeCertificate),
	// 			AgentVersion: to.Ptr("3.1.5.0"),
	// 			AgentVersionExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T21:50:40.5840543Z"); return t}()),
	// 			AgentVersionStatus: to.Ptr(armstoragesync.RegisteredServerAgentVersionStatusOk),
	// 			ApplicationID: to.Ptr("00000000-0000-0000-0000-000000000001"),
	// 			ClusterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			ClusterName: to.Ptr(""),
	// 			DiscoveryEndpointURI: to.Ptr("https://kailanitest99.one.microsoft.com:443"),
	// 			FriendlyName: to.Ptr("afscv-2304-139"),
	// 			Identity: to.Ptr(false),
	// 			LastHeartBeat: to.Ptr("2018-06-11T21:50:40.5840543Z"),
	// 			LastOperationName: to.Ptr("ICreateRegisteredServerWorkflow"),
	// 			LastWorkflowID: to.Ptr("storageSyncServices/CV_FileStore_F1D485AA/workflows/4eecfbcf-9537-4b61-8fee-aaa3ace11c44"),
	// 			LatestApplicationID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 			ManagementEndpointURI: to.Ptr("https://kailanitest99.one.microsoft.com:443/"),
	// 			MonitoringEndpointURI: to.Ptr("https://kailanitest99.one.microsoft.com:443/"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ResourceLocation: to.Ptr("westus"),
	// 			ServerID: to.Ptr("3635ca8b-5cc8-4a5c-bd43-c2de5ad8dc64"),
	// 			ServerManagementErrorCode: to.Ptr[int32](0),
	// 			ServerName: to.Ptr("afscv-2304-139"),
	// 			ServerOSVersion: to.Ptr("10.0.14393.0"),
	// 			ServerRole: to.Ptr("Standalone"),
	// 			ServiceLocation: to.Ptr("westus"),
	// 			StorageSyncServiceUID: to.Ptr("4aa14534-1c61-483b-b6a6-9630a76f109a"),
	// 		},
	// 	},
	// }
}
