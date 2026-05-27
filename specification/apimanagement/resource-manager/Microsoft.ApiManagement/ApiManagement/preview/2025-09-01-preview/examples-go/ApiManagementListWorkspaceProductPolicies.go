package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListWorkspaceProductPolicies.json
func ExampleWorkspaceProductPolicyClient_ListByProduct() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspaceProductPolicyClient().ListByProduct(ctx, "rg1", "apimService1", "wks1", "armTemplateProduct4", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armapimanagement.WorkspaceProductPolicyClientListByProductResponse{
	// 	PolicyCollection: armapimanagement.PolicyCollection{
	// 		Count: to.Ptr[int64](1),
	// 		NextLink: to.Ptr(""),
	// 		Value: []*armapimanagement.PolicyContract{
	// 			{
	// 				Name: to.Ptr("policy"),
	// 				Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/products/policies"),
	// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/products/armTemplateProduct4/policies/policy"),
	// 				Properties: &armapimanagement.PolicyContractProperties{
	// 					Value: to.Ptr("<policies>\r\n  <inbound>\r\n    <base />\r\n  </inbound>\r\n  <backend>\r\n    <base />\r\n  </backend>\r\n  <outbound>\r\n    <base />\r\n  </outbound>\r\n  <on-error>\r\n    <base />\r\n  </on-error>\r\n</policies>"),
	// 				},
	// 			},
	// 		},
	// 	},
	// }
}
