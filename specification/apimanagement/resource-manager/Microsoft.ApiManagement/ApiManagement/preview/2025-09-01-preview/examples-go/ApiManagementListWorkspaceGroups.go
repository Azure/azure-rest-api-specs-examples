package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListWorkspaceGroups.json
func ExampleWorkspaceGroupClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceGroupClient().NewListByServicePager("rg1", "apimService1", "wks1", nil)
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
		// page = armapimanagement.WorkspaceGroupClientListByServiceResponse{
		// 	GroupCollection: armapimanagement.GroupCollection{
		// 		Count: to.Ptr[int64](1),
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armapimanagement.GroupContract{
		// 			{
		// 				Name: to.Ptr("59306a29e4bbd510dc24e5f9"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/groups"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/groups/59306a29e4bbd510dc24e5f9"),
		// 				Properties: &armapimanagement.GroupContractProperties{
		// 					Type: to.Ptr(armapimanagement.GroupTypeExternal),
		// 					Description: to.Ptr("awesome group of people"),
		// 					BuiltIn: to.Ptr(false),
		// 					DisplayName: to.Ptr("AwesomeGroup (samiraad.onmicrosoft.com)"),
		// 					ExternalID: to.Ptr("aad://samiraad.onmicrosoft.com/groups/3773adf4-032e-4d25-9988-eaff9ca72eca"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
