package armkoreaiartemis_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/koreaiartemis/armkoreaiartemis"
)

// Generated from example definition: 2026-07-01-preview/Organizations_ListWorkspaces.json
func ExampleOrganizationsClient_ListWorkspaces() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armkoreaiartemis.NewClientFactory("00000000-0000-0000-0000-000000000000", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewOrganizationsClient().ListWorkspaces(ctx, "contoso-rg", "contoso-org", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armkoreaiartemis.OrganizationsClientListWorkspacesResponse{
	// 	ListWorkspacesResponse: armkoreaiartemis.ListWorkspacesResponse{
	// 		Value: []*armkoreaiartemis.DiscoveredWorkspace{
	// 			{
	// 				PartnerResourceID: to.Ptr("workspace-0001"),
	// 				Name: to.Ptr("Contoso Production"),
	// 				OwnerEmail: to.Ptr("owner@contoso.com"),
	// 				IsDefault: to.Ptr(true),
	// 				Status: to.Ptr(armkoreaiartemis.WorkspaceLifecycleStatusActive),
	// 				IsManagedByLiftr: to.Ptr(true),
	// 				ArmResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso-rg/providers/Kore.AgentPlatform/organizations/contoso-org/workspaces/contoso-production"),
	// 			},
	// 			{
	// 				PartnerResourceID: to.Ptr("workspace-0002"),
	// 				Name: to.Ptr("Partner Managed Workspace"),
	// 				OwnerEmail: to.Ptr("admin@contoso.com"),
	// 				IsDefault: to.Ptr(false),
	// 				Status: to.Ptr(armkoreaiartemis.WorkspaceLifecycleStatusActive),
	// 				IsManagedByLiftr: to.Ptr(false),
	// 			},
	// 		},
	// 	},
	// }
}
