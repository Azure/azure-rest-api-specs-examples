package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListApiSchemas.json
func ExampleAPISchemaClient_NewListByAPIPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAPISchemaClient().NewListByAPIPager("rg1", "apimService1", "59d5b28d1f7fab116c282650", &armapimanagement.APISchemaClientListByAPIOptions{Filter: nil,
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
		// page.SchemaCollection = armapimanagement.SchemaCollection{
		// 	Count: to.Ptr[int64](1),
		// 	Value: []*armapimanagement.SchemaContract{
		// 		{
		// 			Name: to.Ptr("59d5b28e1f7fab116402044e"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/apis/schemas"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/59d5b28d1f7fab116c282650/schemas/59d5b28e1f7fab116402044e"),
		// 			Properties: &armapimanagement.SchemaContractProperties{
		// 				ContentType: to.Ptr("application/vnd.ms-azure-apim.swagger.definitions+json"),
		// 			},
		// 	}},
		// }
	}
}
