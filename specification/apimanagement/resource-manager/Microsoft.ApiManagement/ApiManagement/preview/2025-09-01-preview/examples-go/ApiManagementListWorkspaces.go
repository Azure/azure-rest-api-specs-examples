package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListWorkspaces.json
func ExampleWorkspaceClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceClient().NewListByServicePager("rg1", "apimService1", nil)
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
		// page = armapimanagement.WorkspaceClientListByServiceResponse{
		// 	WorkspaceCollection: armapimanagement.WorkspaceCollection{
		// 		Count: to.Ptr[int64](2),
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armapimanagement.WorkspaceContract{
		// 			{
		// 				Name: to.Ptr("wks1"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/workspaces"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1"),
		// 				Properties: &armapimanagement.WorkspaceContractProperties{
		// 					Description: to.Ptr("workspace 1"),
		// 					DisplayName: to.Ptr("my workspace"),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("wks1"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/workspaces"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks2"),
		// 				Properties: &armapimanagement.WorkspaceContractProperties{
		// 					Description: to.Ptr("workspace 2"),
		// 					DisplayName: to.Ptr("my workspace"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
