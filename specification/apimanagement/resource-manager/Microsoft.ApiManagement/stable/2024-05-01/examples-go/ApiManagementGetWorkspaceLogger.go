package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementGetWorkspaceLogger.json
func ExampleWorkspaceLoggerClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceLoggerClient().Get(ctx, "rg1", "apimService1", "wks1", "templateLogger", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LoggerContract = armapimanagement.LoggerContract{
	// 	Name: to.Ptr("kloudapilogger1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/loggers"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/loggers/kloudapilogger1"),
	// 	Properties: &armapimanagement.LoggerContractProperties{
	// 		Description: to.Ptr("testeventhub3again"),
	// 		Credentials: map[string]*string{
	// 			"name": to.Ptr("testeventhub4"),
	// 			"connectionString": to.Ptr("Endpoint=sb://eventhubapim.servicebus.windows.net/;SharedAccessKeyName=Sender;SharedAccessKey=************"),
	// 		},
	// 		IsBuffered: to.Ptr(true),
	// 		LoggerType: to.Ptr(armapimanagement.LoggerTypeAzureEventHub),
	// 		ResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.EventHub/namespaces/eventhubapim"),
	// 	},
	// }
}
