package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/366aaa13cdd218b9adac716680e49473673410c8/specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/LogicApps_ListCallbackURL.json
func ExampleLogicAppsClient_Invoke() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLogicAppsClient().Invoke(ctx, "testrg123", "testapp2", "testapp2", "<xms-logic-apps-proxy-path>", armappcontainers.LogicAppsProxyMethodPOST, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Interface = map[string]any{
	// 	"method": "POST",
	// 	"basePath": "https://testapp2.livelyriver-730e4fd9.eastasia.azurecontainerapps.io /api/test/triggers/When_a_HTTP_request_is_received/invoke",
	// 	"queries":map[string]any{
	// 		"api-version": "2022-05-01",
	// 		"sig": "IxEQ_ygZf6WNEQCbjV0Vs6p6Y4DyNEJVAa86U5B4xhk",
	// 		"sp": "/triggers/When_a_HTTP_request_is_received/run",
	// 		"sv": "1.0",
	// 	},
	// 	"value": "https://testapp2.livelyriver-730e4fd9.eastasia.azurecontainerapps.io:443/api/test/triggers/When_a_HTTP_request_is_received/invoke?api-version=2022-05-01&sp=%2Ftriggers%2FWhen_a_HTTP_request_is_received%2Frun&sv=1.0&sig=<sig> ",
	// }
}
