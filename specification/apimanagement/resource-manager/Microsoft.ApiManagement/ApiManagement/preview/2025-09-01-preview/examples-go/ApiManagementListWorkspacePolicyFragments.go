package armapimanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v4"
)

// Generated from example definition: 2025-09-01-preview/ApiManagementListWorkspacePolicyFragments.json
func ExampleWorkspacePolicyFragmentClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armapimanagement.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspacePolicyFragmentClient().NewListByServicePager("rg1", "apimService1", "wks1", nil)
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
		// page = armapimanagement.WorkspacePolicyFragmentClientListByServiceResponse{
		// 	PolicyFragmentCollection: armapimanagement.PolicyFragmentCollection{
		// 		Count: to.Ptr[int64](1),
		// 		NextLink: to.Ptr(""),
		// 		Value: []*armapimanagement.PolicyFragmentContract{
		// 			{
		// 				Name: to.Ptr("policyFragment1"),
		// 				Type: to.Ptr("Microsoft.ApiManagement/service/workspaces/policyFragments"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/policyFragments/policyFragment1"),
		// 				Properties: &armapimanagement.PolicyFragmentContractProperties{
		// 					Format: to.Ptr(armapimanagement.PolicyFragmentContentFormatXML),
		// 					Description: to.Ptr("A policy fragment example"),
		// 					Value: to.Ptr("<fragment><json-to-xml apply=\"always\" consider-accept-header=\"false\" /></fragment>"),
		// 				},
		// 			},
		// 		},
		// 	},
		// }
	}
}
