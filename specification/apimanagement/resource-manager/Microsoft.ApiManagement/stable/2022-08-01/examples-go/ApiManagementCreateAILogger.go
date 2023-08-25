package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateAILogger.json
func ExampleLoggerClient_CreateOrUpdate_apiManagementCreateAiLogger() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLoggerClient().CreateOrUpdate(ctx, "rg1", "apimService1", "loggerId", armapimanagement.LoggerContract{
		Properties: &armapimanagement.LoggerContractProperties{
			Description: to.Ptr("adding a new logger"),
			Credentials: map[string]*string{
				"instrumentationKey": to.Ptr("11................a1"),
			},
			LoggerType: to.Ptr(armapimanagement.LoggerTypeApplicationInsights),
		},
	}, &armapimanagement.LoggerClientCreateOrUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LoggerContract = armapimanagement.LoggerContract{
	// 	Name: to.Ptr("loggerId"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/loggers"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/loggers/loggerId"),
	// 	Properties: &armapimanagement.LoggerContractProperties{
	// 		Credentials: map[string]*string{
	// 			"instrumentationKey": to.Ptr("{{5a.......2a}}"),
	// 		},
	// 		IsBuffered: to.Ptr(false),
	// 		LoggerType: to.Ptr(armapimanagement.LoggerTypeApplicationInsights),
	// 	},
	// }
}
