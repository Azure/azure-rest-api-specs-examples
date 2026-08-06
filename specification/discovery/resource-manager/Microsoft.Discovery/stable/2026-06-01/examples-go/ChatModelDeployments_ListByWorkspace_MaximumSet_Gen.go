package armdiscovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/discovery/armdiscovery"
)

// Generated from example definition: 2026-06-01/ChatModelDeployments_ListByWorkspace_MaximumSet_Gen.json
func ExampleChatModelDeploymentsClient_NewListByWorkspacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdiscovery.NewClientFactory("A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewChatModelDeploymentsClient().NewListByWorkspacePager("rgdiscovery", "0f2d15df9509076ccf", nil)
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
		// page = armdiscovery.ChatModelDeploymentsClientListByWorkspaceResponse{
		// 	ChatModelDeploymentListResult: armdiscovery.ChatModelDeploymentListResult{
		// 		Value: []*armdiscovery.ChatModelDeployment{
		// 			{
		// 				ID: to.Ptr("/subscriptions/A54D43BD-2F5F-4BB1-95D4-9A8D23CC7DD4/resourceGroups/rgdiscovery/providers/Microsoft.Discovery/workspaces/0f2d15df9509076ccf/chatModelDeployments/eefcjuwvmdwzssuuknvzpgjjjpto"),
		// 				Name: to.Ptr("eefcjuwvmdwzssuuknvzpgjjjpto"),
		// 				Tags: map[string]*string{
		// 					"key984": to.Ptr("sqzgsgykyhltqwmpgvhlyp"),
		// 				},
		// 				Location: to.Ptr("uksouth"),
		// 				Type: to.Ptr("Microsoft.Discovery/workspaces/chatModelDeployments"),
		// 				SystemData: &armdiscovery.SystemData{
		// 					CreatedBy: to.Ptr("uymdmmhvojuqtvvxokgefohqpcjw"),
		// 					CreatedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("ucuttxilomgszapozsuit"),
		// 					LastModifiedByType: to.Ptr(armdiscovery.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2026-05-04T11:59:49.804Z"); return t}()),
		// 				},
		// 				Properties: &armdiscovery.ChatModelDeploymentProperties{
		// 					ProvisioningState: to.Ptr(armdiscovery.ProvisioningStateSucceeded),
		// 					ModelFormat: to.Ptr("zo"),
		// 					ModelName: to.Ptr("ijzwlirrkr"),
		// 					ModelVersion: to.Ptr("seiduxog"),
		// 					SKUName: to.Ptr("dymgademiauwwacz"),
		// 					Capacity: to.Ptr[int32](8),
		// 				},
		// 			},
		// 		},
		// 		NextLink: to.Ptr("https://microsoft.com/a"),
		// 	},
		// }
	}
}
