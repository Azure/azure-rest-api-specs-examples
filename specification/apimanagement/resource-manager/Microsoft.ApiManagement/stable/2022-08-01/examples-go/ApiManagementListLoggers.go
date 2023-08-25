package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListLoggers.json
func ExampleLoggerClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewLoggerClient().NewListByServicePager("rg1", "apimService1", &armapimanagement.LoggerClientListByServiceOptions{Filter: nil,
		Top:  nil,
		Skip: nil,
	})
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
		// page.LoggerCollection = armapimanagement.LoggerCollection{
		// 	Count: to.Ptr[int64](3),
		// 	Value: []*armapimanagement.LoggerContract{
		// 		{
		// 			Name: to.Ptr("azuremonitor"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/loggers"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/loggers/azuremonitor"),
		// 			Properties: &armapimanagement.LoggerContractProperties{
		// 				IsBuffered: to.Ptr(true),
		// 				LoggerType: to.Ptr(armapimanagement.LoggerTypeAzureMonitor),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("vvktest"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/loggers"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/loggers/vvktest"),
		// 			Properties: &armapimanagement.LoggerContractProperties{
		// 				Credentials: map[string]*string{
		// 					"instrumentationKey": to.Ptr("{{Logger-Credentials-5b1a17ef2b3f91153004b10d}}"),
		// 				},
		// 				IsBuffered: to.Ptr(true),
		// 				LoggerType: to.Ptr(armapimanagement.LoggerTypeApplicationInsights),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("applicationinsights"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/loggers"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/loggers/applicationinsights"),
		// 			Properties: &armapimanagement.LoggerContractProperties{
		// 				Description: to.Ptr("miaoappinsight"),
		// 				Credentials: map[string]*string{
		// 					"instrumentationKey": to.Ptr("{{Logger-Credentials-5b2056062b3f911ae84a3069}}"),
		// 				},
		// 				IsBuffered: to.Ptr(true),
		// 				LoggerType: to.Ptr(armapimanagement.LoggerTypeApplicationInsights),
		// 			},
		// 	}},
		// }
	}
}
