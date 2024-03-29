package armstoragesync_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2020-09-01/examples/RegisteredServers_Create.json
func ExampleRegisteredServersClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstoragesync.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewRegisteredServersClient().BeginCreate(ctx, "SampleResourceGroup_1", "SampleStorageSyncService_1", "080d4133-bdb5-40a0-96a0-71a6057bfe9a", armstoragesync.RegisteredServerCreateParameters{
		Properties: &armstoragesync.RegisteredServerCreateParametersProperties{
			AgentVersion:      to.Ptr("1.0.277.0"),
			FriendlyName:      to.Ptr("afscv-2304-139"),
			ServerCertificate: to.Ptr("MIIDFjCCAf6gAwIBAgIQQS+DS8uhc4VNzUkTw7wbRjANBgkqhkiG9w0BAQ0FADAzMTEwLwYDVQQDEyhhbmt1c2hiLXByb2QzLnJlZG1vbmQuY29ycC5taWNyb3NvZnQuY29tMB4XDTE3MDgwMzE3MDQyNFoXDTE4MDgwNDE3MDQyNFowMzExMC8GA1UEAxMoYW5rdXNoYi1wcm9kMy5yZWRtb25kLmNvcnAubWljcm9zb2Z0LmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALDRvV4gmsIy6jGDPiHsXmvgVP749NNP7DopdlbHaNhjFmYINHl0uWylyaZmgJrROt2mnxN/zEyJtGnqYHlzUr4xvGq/qV5pqgdB9tag/sw9i22gfe9PRZ0FmSOZnXMbLYgLiDFqLtut5gHcOuWMj03YnkfoBEKlFBxWbagvW2yxz/Sxi9OVSJOKCaXra0RpcIHrO/KFl6ho2eE1/7Ykmfa8hZvSdoPd5gHdLiQcMB/pxq+mWp1fI6c8vFZoDu7Atn+NXTzYPKUxKzaisF12TsaKpohUsJpbB3Wocb0F5frn614D2pg14ERB5otjAMWw1m65csQWPI6dP8KIYe0+QPkCAwEAAaMmMCQwIgYDVR0lAQH/BBgwFgYIKwYBBQUHAwIGCisGAQQBgjcKAwwwDQYJKoZIhvcNAQENBQADggEBAA4RhVIBkw34M1RwakJgHvtjsOFxF1tVQA941NtLokx1l2Z8+GFQkcG4xpZSt+UN6wLerdCbnNhtkCErWUDeaT0jxk4g71Ofex7iM04crT4iHJr8mi96/XnhnkTUs+GDk12VgdeeNEczMZz+8Mxw9dJ5NCnYgTwO0SzGlclRsDvjzkLo8rh2ZG6n/jKrEyNXXo+hOqhupij0QbRP2Tvexdfw201kgN1jdZify8XzJ8Oi0bTS0KpJf2pNPOlooK2bjMUei9ANtEdXwwfVZGWvVh6tJjdv6k14wWWJ1L7zhA1IIVb1J+sQUzJji5iX0DrezjTz1Fg+gAzITaA/WsuujlM="),
			ServerID:          to.Ptr("080d4133-bdb5-40a0-96a0-71a6057bfe9a"),
			ServerOSVersion:   to.Ptr("10.0.14393.0"),
			ServerRole:        to.Ptr("Standalone"),
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
	// res.RegisteredServer = armstoragesync.RegisteredServer{
	// 	Name: to.Ptr("ankushb-prod3.redmond.corp.microsoft.com"),
	// 	Type: to.Ptr("Microsoft.StorageSync/storageSyncServices/registeredServers"),
	// 	ID: to.Ptr("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/530a0384-50ac-456d-8240-9d6621404151"),
	// 	Properties: &armstoragesync.RegisteredServerProperties{
	// 		AgentVersion: to.Ptr("3.1.5.0"),
	// 		AgentVersionExpirationDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-11T21:50:40.584Z"); return t}()),
	// 		AgentVersionStatus: to.Ptr(armstoragesync.RegisteredServerAgentVersionStatusOk),
	// 		ClusterID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		ClusterName: to.Ptr(""),
	// 		DiscoveryEndpointURI: to.Ptr("https://kailanitest99.one.microsoft.com:443"),
	// 		FriendlyName: to.Ptr("afscv-2304-139"),
	// 		LastHeartBeat: to.Ptr("2018-06-11T21:50:40.5840543Z"),
	// 		LastOperationName: to.Ptr("ICreateRegisteredServerWorkflow"),
	// 		LastWorkflowID: to.Ptr("storageSyncServices/CV_FileStore_F1D485AA/workflows/4eecfbcf-9537-4b61-8fee-aaa3ace11c44"),
	// 		ManagementEndpointURI: to.Ptr("https://kailanitest99.one.microsoft.com:443/"),
	// 		MonitoringEndpointURI: to.Ptr("https://kailanitest99.one.microsoft.com:443/"),
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		ResourceLocation: to.Ptr("westus"),
	// 		ServerID: to.Ptr("3635ca8b-5cc8-4a5c-bd43-c2de5ad8dc64"),
	// 		ServerManagementErrorCode: to.Ptr[int32](0),
	// 		ServerName: to.Ptr("afscv-2304-139"),
	// 		ServerOSVersion: to.Ptr("10.0.14393.0"),
	// 		ServerRole: to.Ptr("Standalone"),
	// 		ServiceLocation: to.Ptr("westus"),
	// 		StorageSyncServiceUID: to.Ptr("4aa14534-1c61-483b-b6a6-9630a76f109a"),
	// 	},
	// }
}
