package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/4cd95123fb961c68740565a1efcaa5e43bd35802/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGatewayApis.json
func ExampleGatewayAPIClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewGatewayAPIClient().NewListByServicePager("rg1", "apimService1", "gw1", &armapimanagement.GatewayAPIClientListByServiceOptions{Filter: nil,
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
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/gateways/apis"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/gateways/gw1/apis/57681820a40f7eb6c49f6aca"),
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
