package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/e436160e64c0f8d7fb20d662be2712f71f0a7ef5/specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementListWorkspaceApis.json
func ExampleWorkspaceAPIClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceAPIClient().NewListByServicePager("rg1", "apimService1", "wks1", &armapimanagement.WorkspaceAPIClientListByServiceOptions{Filter: nil,
		Top:                 nil,
		Skip:                nil,
		Tags:                nil,
		ExpandAPIVersionSet: nil,
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
		// 	Count: to.Ptr[int64](4),
		// 	Value: []*armapimanagement.APIContract{
		// 		{
		// 			Name: to.Ptr("a1"),
		// 			Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/apis"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/a1"),
		// 			Properties: &armapimanagement.APIContractProperties{
		// 				APIRevision: to.Ptr("1"),
		// 				APIVersionSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apiVersionSets/c48f96c9-1385-4e2d-b410-5ab591ce0fc4"),
		// 				IsCurrent: to.Ptr(true),
		// 				Path: to.Ptr("api1"),
		// 				DisplayName: to.Ptr("api1"),
		// 				Protocols: []*armapimanagement.Protocol{
		// 					to.Ptr(armapimanagement.ProtocolHTTPS)},
		// 					ServiceURL: to.Ptr("http://echoapi.cloudapp.net/api"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("5a73933b8f27f7cc82a2d533"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/apis"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/5a73933b8f27f7cc82a2d533"),
		// 				Properties: &armapimanagement.APIContractProperties{
		// 					APIRevision: to.Ptr("1"),
		// 					APIVersion: to.Ptr("v1"),
		// 					APIVersionSetID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apiVersionSets/c48f96c9-1385-4e2d-b410-5ab591ce0fc4"),
		// 					IsCurrent: to.Ptr(true),
		// 					Path: to.Ptr("api1"),
		// 					DisplayName: to.Ptr("api1"),
		// 					Protocols: []*armapimanagement.Protocol{
		// 						to.Ptr(armapimanagement.ProtocolHTTPS)},
		// 						ServiceURL: to.Ptr("http://echoapi.cloudapp.net/api"),
		// 					},
		// 				},
		// 				{
		// 					Name: to.Ptr("echo-api"),
		// 					Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/apis"),
		// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/echo-api"),
		// 					Properties: &armapimanagement.APIContractProperties{
		// 						APIRevision: to.Ptr("1"),
		// 						IsCurrent: to.Ptr(true),
		// 						Path: to.Ptr("echo"),
		// 						DisplayName: to.Ptr("Echo API"),
		// 						Protocols: []*armapimanagement.Protocol{
		// 							to.Ptr(armapimanagement.ProtocolHTTPS)},
		// 							ServiceURL: to.Ptr("http://echoapi.cloudapp.net/api"),
		// 						},
		// 					},
		// 					{
		// 						Name: to.Ptr("5a7390baa5816a110435aee0"),
		// 						Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/apis"),
		// 						ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/apis/5a7390baa5816a110435aee0"),
		// 						Properties: &armapimanagement.APIContractProperties{
		// 							Description: to.Ptr("A sample API that uses a petstore as an example to demonstrate features in the swagger-2.0 specification"),
		// 							APIRevision: to.Ptr("1"),
		// 							IsCurrent: to.Ptr(true),
		// 							Path: to.Ptr("vvv"),
		// 							DisplayName: to.Ptr("Swagger Petstore Extensive"),
		// 							Protocols: []*armapimanagement.Protocol{
		// 								to.Ptr(armapimanagement.ProtocolHTTPS)},
		// 								ServiceURL: to.Ptr("http://petstore.swagger.wordnik.com/api"),
		// 							},
		// 					}},
		// 				}
	}
}
