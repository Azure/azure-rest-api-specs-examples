package armdatafactory_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v7"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/92de53a5f1e0e03c94b40475d2135d97148ed014/specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/IntegrationRuntimes_GetStatus.json
func ExampleIntegrationRuntimesClient_GetStatus() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatafactory.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationRuntimesClient().GetStatus(ctx, "exampleResourceGroup", "exampleFactoryName", "exampleIntegrationRuntime", nil)
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
	// 		State: to.Ptr(armdatafactory.IntegrationRuntimeStateOnline),
	// 		TypeProperties: &armdatafactory.SelfHostedIntegrationRuntimeStatusTypeProperties{
	// 			AutoUpdate: to.Ptr(armdatafactory.IntegrationRuntimeAutoUpdateOff),
	// 			Capabilities: map[string]*string{
	// 				"connectedToResourceManager": to.Ptr("True"),
	// 				"credentialInSync": to.Ptr("True"),
	// 				"httpsPortEnabled": to.Ptr("True"),
	// 				"nodeEnabled": to.Ptr("True"),
	// 				"serviceBusConnected": to.Ptr("True"),
	// 			},
	// 			CreateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-14T09:17:45.183Z"); return t}()),
	// 			LatestVersion: to.Ptr("3.7.6711.1"),
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
	// 					LastConnectTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-14T14:52:59.893Z"); return t}()),
	// 					LastStartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-14T14:52:59.893Z"); return t}()),
	// 					LastUpdateResult: to.Ptr(armdatafactory.IntegrationRuntimeUpdateResultNone),
	// 					MachineName: to.Ptr("YANZHANG-DT"),
	// 					MaxConcurrentJobs: to.Ptr[int32](56),
	// 					NodeName: to.Ptr("Node_1"),
	// 					RegisterTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-14T14:51:44.923Z"); return t}()),
	// 					Status: to.Ptr(armdatafactory.SelfHostedIntegrationRuntimeNodeStatusOnline),
	// 					Version: to.Ptr("3.8.6730.2"),
	// 					VersionStatus: to.Ptr("UpToDate"),
	// 			}},
	// 			ServiceUrls: []*string{
	// 				to.Ptr("wu.frontend.int.clouddatahub-int.net"),
	// 				to.Ptr("*.servicebus.windows.net")},
	// 				TaskQueueID: to.Ptr("1a6296ab-423c-4346-9bcc-85a78c2c0582"),
	// 				UpdateDelayOffset: to.Ptr("PT3H"),
	// 				Version: to.Ptr("3.8.6730.2"),
	// 				VersionStatus: to.Ptr("UpToDate"),
	// 			},
	// 		},
	// 	}
}
