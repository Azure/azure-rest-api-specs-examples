package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/c767823fdfd9d5e96bad245e3ea4d14d94a716bb/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListProductApis.json
func ExampleProductAPIClient_NewListByProductPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewProductAPIClient().NewListByProductPager("rg1", "apimService1", "5768181ea40f7eb6c49f6ac7", &armapimanagement.ProductAPIClientListByProductOptions{Filter: nil,
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
		// page.APICollection = armapimanagement.APICollection{
		// 	Count: to.Ptr[int64](1),
		// 	Value: []*armapimanagement.APIContract{
		// 		{
		// 			Name: to.Ptr("57681820a40f7eb6c49f6aca"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/products/apis"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/5768181ea40f7eb6c49f6ac7/apis/57681820a40f7eb6c49f6aca"),
		// 			Properties: &armapimanagement.APIContractProperties{
		// 				Description: to.Ptr("description_57681820a40f7eb6c49f6acc"),
		// 				APIRevision: to.Ptr("1"),
		// 				IsCurrent: to.Ptr(true),
		// 				Path: to.Ptr("suffix_57681820a40f7eb6c49f6ace"),
		// 				DisplayName: to.Ptr("api_57681820a40f7eb6c49f6acb"),
		// 				Protocols: []*armapimanagement.Protocol{
		// 					to.Ptr(armapimanagement.ProtocolHTTPS)},
		// 					ServiceURL: to.Ptr("http://contoso/57681820a40f7eb6c49f6acd"),
		// 				},
		// 		}},
		// 	}
	}
}
