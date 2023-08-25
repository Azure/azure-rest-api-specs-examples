package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateLogger.json
func ExampleLoggerClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewLoggerClient().Update(ctx, "rg1", "apimService1", "eh1", "*", armapimanagement.LoggerUpdateContract{
		Properties: &armapimanagement.LoggerUpdateParameters{
			Description: to.Ptr("updating description"),
			LoggerType:  to.Ptr(armapimanagement.LoggerTypeAzureEventHub),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.LoggerContract = armapimanagement.LoggerContract{
	// 	Name: to.Ptr("eh1"),
	// 	Type: to.Ptr("Microsoft.ApiManagement/service/loggers"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/loggers/eh1"),
	// 	Properties: &armapimanagement.LoggerContractProperties{
	// 		Description: to.Ptr("updating description"),
	// 		Credentials: map[string]*string{
	// 			"connectionString": to.Ptr("{{Logger-Credentials-5f28745bbebeeb13cc3f7301}}"),
	// 		},
	// 		IsBuffered: to.Ptr(true),
	// 		LoggerType: to.Ptr(armapimanagement.LoggerTypeAzureEventHub),
	// 	},
	// }
}
