package armchaos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
)

// Generated from example definition: 2026-05-01-preview/Workspaces_Get.json
func ExampleWorkspacesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armchaos.NewClientFactory("6b052e15-03d3-4f17-b2e1-be7f07588291", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWorkspacesClient().Get(ctx, "exampleRG", "exampleWorkspace", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res = armchaos.WorkspacesClientGetResponse{
	// 	Workspace: armchaos.Workspace{
	// 		Name: to.Ptr("exampleWorkspace"),
	// 		Type: to.Ptr("Microsoft.Chaos/workspaces"),
	// 		ID: to.Ptr("/subscriptions/6b052e15-03d3-4f17-b2e1-be7f07588291/resourceGroups/exampleRG/providers/Microsoft.Chaos/workspaces/exampleWorkspace"),
	// 		Identity: &armchaos.ManagedServiceIdentity{
	// 			Type: to.Ptr(armchaos.ManagedServiceIdentityTypeUserAssigned),
	// 			UserAssignedIdentities: map[string]*armchaos.UserAssignedIdentity{
	// 				"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/exampleResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/exampleIdentity": &armchaos.UserAssignedIdentity{
	// 				},
	// 			},
	// 			TenantID: to.Ptr("8c3e2fb2-fe7a-4bf1-b779-d73990782fe6"),
	// 		},
	// 		Tags: map[string]*string{
	// 			"environment": to.Ptr("production"),
	// 			"department": to.Ptr("engineering"),
	// 			"project": to.Ptr("chaos-testing"),
	// 		},
	// 		Location: to.Ptr("westus2"),
	// 		Properties: &armchaos.WorkspaceProperties{
	// 			ProvisioningState: to.Ptr(armchaos.ProvisioningStateUpdating),
	// 			Scopes: []*string{
	// 				to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/exampleResourceGroup"),
	// 			},
	// 		},
	// 		SystemData: &armchaos.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.0Z"); return t}()),
	// 			CreatedBy: to.Ptr("User"),
	// 			CreatedByType: to.Ptr(armchaos.CreatedByType("b3a41dba-4415-4d36-9ee8-e5eaa86db976")),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-07-01T00:00:00.0Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("User"),
	// 			LastModifiedByType: to.Ptr(armchaos.CreatedByType("b3a41dba-4415-4d36-9ee8-e5eaa86db976")),
	// 		},
	// 	},
	// }
}
