package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateWebsocketApi.json
func ExampleAPIClient_BeginCreateOrUpdate_apiManagementCreateWebSocketApi() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAPIClient().BeginCreateOrUpdate(ctx, "rg1", "apimService1", "tempgroup", armapimanagement.APICreateOrUpdateParameter{
		Properties: &armapimanagement.APICreateOrUpdateProperties{
			APIType:     to.Ptr(armapimanagement.APITypeWebsocket),
			Description: to.Ptr("apidescription5200"),
			Path:        to.Ptr("newapiPath"),
			DisplayName: to.Ptr("apiname1463"),
			Protocols: []*armapimanagement.Protocol{
				to.Ptr(armapimanagement.ProtocolWss),
				to.Ptr(armapimanagement.ProtocolWs)},
			ServiceURL: to.Ptr("wss://echo.websocket.org"),
		},
	}, &armapimanagement.APIClientBeginCreateOrUpdateOptions{IfMatch: nil})
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
	// res.APIContract = armapimanagement.APIContract{
	// 	Name: to.Ptr("apiid9419"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/apis"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/apiid9419"),
	// 	Properties: &armapimanagement.APIContractProperties{
	// 		APIType: to.Ptr(armapimanagement.APITypeWebsocket),
	// 		Description: to.Ptr("apidescription5200"),
	// 		APIRevision: to.Ptr("1"),
	// 		IsCurrent: to.Ptr(true),
	// 		IsOnline: to.Ptr(true),
	// 		Path: to.Ptr("newapiPath"),
	// 		DisplayName: to.Ptr("apiname1463"),
	// 		Protocols: []*armapimanagement.Protocol{
	// 			to.Ptr(armapimanagement.ProtocolWs),
	// 			to.Ptr(armapimanagement.ProtocolWss)},
	// 			ServiceURL: to.Ptr("wss://echo.websocket.org"),
	// 		},
	// 	}
}
