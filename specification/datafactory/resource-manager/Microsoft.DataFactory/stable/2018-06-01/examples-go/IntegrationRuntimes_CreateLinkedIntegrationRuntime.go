package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v8"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/f6f50c6388fd5836fa142384641b8353a99874ef/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_CreateLinkedIntegrationRuntime.json
func ExampleIntegrationRuntimesClient_CreateLinkedIntegrationRuntime() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationRuntimesClient().CreateLinkedIntegrationRuntime(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleIntegrationRuntime", armdatafactory.CreateLinkedIntegrationRuntimeRequest{
		Name:                to.Ptr("bfa92911-9fb6-4fbe-8f23-beae87bc1c83"),
		DataFactoryLocation: to.Ptr("West US"),
		DataFactoryName:     to.Ptr("e9955d6d-56ea-4be3-841c-52a12c1a9981"),
		SubscriptionID:      to.Ptr("061774c7-4b5a-4159-a55b-365581830283"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationRuntimeStatusResponse = armdatafactory.IntegrationRuntimeStatusResponse{
	// 	Name: to.Ptr("exampleIntegrationRuntime"),
	// 	Properties: &armdatafactory.SelfHostedIntegrationRuntimeStatus{
	// 		Type: to.Ptr(armdatafactory.IntegrationRuntimeTypeSelfHosted),
	// 		DataFactoryName: to.Ptr("exampleFactoryName"),
	// 		State: to.Ptr(armdatafactory.IntegrationRuntimeStateOnline),
	// 		TypeProperties: &armdatafactory.SelfHostedIntegrationRuntimeStatusTypeProperties{
	// 			AutoUpdate: to.Ptr(armdatafactory.IntegrationRuntimeAutoUpdateOn),
	// 			AutoUpdateETA: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-20T19:00:00.000Z"); return t}()),
	// 			Capabilities: map[string]*string{
	// 				"connectedToResourceManager": to.Ptr("True"),
	// 				"credentialInSync": to.Ptr("True"),
	// 				"httpsPortEnabled": to.Ptr("True"),
	// 				"nodeEnabled": to.Ptr("True"),
	// 				"serviceBusConnected": to.Ptr("True"),
	// 			},
	// 			CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-17T03:43:25.705Z"); return t}()),
	// 			LatestVersion: to.Ptr("3.9.6774.1"),
	// 			Links: []*armdatafactory.LinkedIntegrationRuntime{
	// 				{
	// 					Name: to.Ptr("bfa92911-9fb6-4fbe-8f23-beae87bc1c83"),
	// 					CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-17T06:31:04.061Z"); return t}()),
	// 					DataFactoryLocation: to.Ptr("West US"),
	// 					DataFactoryName: to.Ptr("e9955d6d-56ea-4be3-841c-52a12c1a9981"),
	// 					SubscriptionID: to.Ptr("061774c7-4b5a-4159-a55b-365581830283"),
	// 			}},
	// 			LocalTimeZoneOffset: to.Ptr("PT8H"),
	// 			Nodes: []*armdatafactory.SelfHostedIntegrationRuntimeNode{
	// 				{
	// 					Capabilities: map[string]*string{
	// 						"connectedToResourceManager": to.Ptr("True"),
	// 						"credentialInSync": to.Ptr("True"),
	// 						"httpsPortEnabled": to.Ptr("True"),
	// 						"nodeEnabled": to.Ptr("True"),
	// 						"serviceBusConnected": to.Ptr("True"),
	// 					},
	// 					HostServiceURI: to.Ptr("https://yanzhang-dt.fareast.corp.microsoft.com:8050/HostServiceRemote.svc/"),
	// 					IsActiveDispatcher: to.Ptr(true),
	// 					LastConnectTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-17T06:30:46.626Z"); return t}()),
	// 					LastStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-17T03:45:30.849Z"); return t}()),
	// 					LastUpdateResult: to.Ptr(armdatafactory.IntegrationRuntimeUpdateResultNone),
	// 					MachineName: to.Ptr("YANZHANG-DT"),
	// 					MaxConcurrentJobs: to.Ptr[int32](20),
	// 					NodeName: to.Ptr("Node_1"),
	// 					RegisterTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-17T03:44:55.801Z"); return t}()),
	// 					Status: to.Ptr(armdatafactory.SelfHostedIntegrationRuntimeNodeStatusOnline),
	// 					Version: to.Ptr("3.8.6743.6"),
	// 					VersionStatus: to.Ptr("UpToDate"),
	// 			}},
	// 			PushedVersion: to.Ptr("3.9.6774.1"),
	// 			ScheduledUpdateDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-08-20T00:00:00.000Z"); return t}()),
	// 			ServiceUrls: []*string{
	// 				to.Ptr("wu.frontend.int.clouddatahub-int.net"),
	// 				to.Ptr("*.servicebus.windows.net")},
	// 				TaskQueueID: to.Ptr("823da112-f2d9-426b-a0d8-5f361b94f72a"),
	// 				UpdateDelayOffset: to.Ptr("PT19H"),
	// 				Version: to.Ptr("3.8.6743.6"),
	// 				VersionStatus: to.Ptr("UpdateAvailable"),
	// 			},
	// 		},
	// 	}
}
