package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetApiExportInOpenApi2dot0.json
func ExampleAPIExportClient_Get_apiManagementGetApiExportInOpenApi2Dot0() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAPIExportClient().Get(ctx, "rg1", "apimService1", "echo-api", armapimanagement.ExportFormatSwagger, armapimanagement.ExportAPITrue, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.APIExportResult = armapimanagement.APIExportResult{
	// 	ExportResultFormat: to.Ptr(armapimanagement.ExportResultFormatSwagger),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echo-api"),
	// 	Value: &armapimanagement.APIExportResultValue{
	// 		Link: to.Ptr("https://apimgmtstkjpszxxxxxxx.blob.core.windows.net/api-export/Swagger Petstore Extensive.json?sv=2015-07-08&sr=b&sig=mxhLsFuOonu8EXIjyFPV%2FnDra0qTIoip7N7MuU%2BTFsA%3D&se=2019-04-10T22:41:31Z&sp=r"),
	// 	},
	// }
}
