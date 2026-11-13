package armkoreaiartemis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/koreaiartemis/armkoreaiartemis"
)

// Generated from example definition: 2026-07-01-preview/Workspaces_ListLinkedSaaS.json
func ExampleWorkspacesClient_ListLinkedSaaS() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkoreaiartemis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().ListLinkedSaaS(ctx, "contoso-rg", "contoso-org", "contoso-prod", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkoreaiartemis.WorkspacesClientListLinkedSaaSResponse{
	// 	WorkspaceLinkedSaaSListResponse: armkoreaiartemis.WorkspaceLinkedSaaSListResponse{
	// 		Value: []*armkoreaiartemis.WorkspaceLinkedSaaSItem{
	// 			{
	// 				SaaSResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Microsoft.SaaS/resources/kore-ai-saas-1"),
	// 				IsHiddenSaaS: to.Ptr(false),
	// 			},
	// 			{
	// 				SaaSResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Microsoft.SaaS/resources/kore-ai-saas-2"),
	// 				IsHiddenSaaS: to.Ptr(false),
	// 			},
	// 		},
	// 	},
	// }
}
