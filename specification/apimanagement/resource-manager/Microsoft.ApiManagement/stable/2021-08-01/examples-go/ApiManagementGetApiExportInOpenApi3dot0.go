package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiExportInOpenApi3dot0.json
func ExampleAPIExportClient_Get_apiManagementGetApiExportInOpenApi3Dot0() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIExportClient().Get(ctx, "rg1", "apimService1", "aid9676", armapimanagement.ExportFormatOpenapi, armapimanagement.ExportAPITrue, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.APIExportResult = armapimanagement.APIExportResult{
	// 	ExportResultFormat: to.Ptr(armapimanagement.ExportResultFormatOpenAPI),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/aid9676"),
	// 	Value: &armapimanagement.APIExportResultValue{
	// 		Link: to.Ptr("https: //apimgmtstkjpszxxxxxxx.blob.core.windows.net/api-export/Swagger Petstore.yaml?sv=2015-07-08&sr=b&sig=qqtR1y5iTbz5P7USBduqB5vriIU4gmiGqe0lKVV8j9k%3D&se=2019-04-10T22:40:57Z&sp=r"),
	// 	},
	// }
}
